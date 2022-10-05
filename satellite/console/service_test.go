// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package console_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"storj.io/common/currency"
	"storj.io/common/macaroon"
	"storj.io/common/memory"
	"storj.io/common/storj"
	"storj.io/common/testcontext"
	"storj.io/common/testrand"
	"storj.io/common/uuid"
	"storj.io/storj/private/blockchain"
	"storj.io/storj/private/testplanet"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/payments"
	"storj.io/storj/satellite/payments/coinpayments"
	"storj.io/storj/satellite/payments/storjscan"
	"storj.io/storj/satellite/payments/storjscan/blockchaintest"
	"storj.io/storj/satellite/payments/stripecoinpayments"
)

func TestService(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 2,
	},
		func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
			sat := planet.Satellites[0]
			service := sat.API.Console.Service

			up1Pro1, err := sat.API.DB.Console().Projects().Get(ctx, planet.Uplinks[0].Projects[0].ID)
			require.NoError(t, err)
			up2Pro1, err := sat.API.DB.Console().Projects().Get(ctx, planet.Uplinks[1].Projects[0].ID)
			require.NoError(t, err)

			up2User, err := sat.API.DB.Console().Users().Get(ctx, up2Pro1.OwnerID)
			require.NoError(t, err)

			require.NotEqual(t, up1Pro1.ID, up2Pro1.ID)
			require.NotEqual(t, up1Pro1.OwnerID, up2Pro1.OwnerID)

			userCtx1, err := sat.UserContext(ctx, up1Pro1.OwnerID)
			require.NoError(t, err)

			userCtx2, err := sat.UserContext(ctx, up2Pro1.OwnerID)
			require.NoError(t, err)

			t.Run("GetProject", func(t *testing.T) {
				// Getting own project details should work
				project, err := service.GetProject(userCtx1, up1Pro1.ID)
				require.NoError(t, err)
				require.Equal(t, up1Pro1.ID, project.ID)

				// Getting someone else project details should not work
				project, err = service.GetProject(userCtx1, up2Pro1.ID)
				require.Error(t, err)
				require.Nil(t, project)
			})

			t.Run("GetSalt", func(t *testing.T) {
				// Getting project salt as a member should work
				salt, err := service.GetSalt(userCtx1, up1Pro1.ID)
				require.NoError(t, err)
				require.NotNil(t, salt)

				// Getting project salt as a non-member should not work
				salt, err = service.GetSalt(userCtx1, up2Pro1.ID)
				require.Error(t, err)
				require.Nil(t, salt)
			})

			t.Run("UpdateProject", func(t *testing.T) {
				updatedName := "newName"
				updatedDescription := "newDescription"
				updatedStorageLimit := memory.Size(100)
				updatedBandwidthLimit := memory.Size(100)

				// user should be in free tier
				user, err := service.GetUser(ctx, up1Pro1.OwnerID)
				require.NoError(t, err)
				require.False(t, user.PaidTier)
				// get context
				userCtx1, err := sat.UserContext(ctx, user.ID)
				require.NoError(t, err)
				// add a credit card to put the user in the paid tier
				err = service.Payments().AddCreditCard(userCtx1, "test-cc-token")
				require.NoError(t, err)
				// update auth ctx
				userCtx1, err = sat.UserContext(ctx, user.ID)
				require.NoError(t, err)

				// Updating own project should work
				updatedProject, err := service.UpdateProject(userCtx1, up1Pro1.ID, console.ProjectInfo{
					Name:           updatedName,
					Description:    updatedDescription,
					StorageLimit:   updatedStorageLimit,
					BandwidthLimit: updatedBandwidthLimit,
				})
				require.NoError(t, err)
				require.NotEqual(t, up1Pro1.Name, updatedProject.Name)
				require.Equal(t, updatedName, updatedProject.Name)
				require.NotEqual(t, up1Pro1.Description, updatedProject.Description)
				require.Equal(t, updatedDescription, updatedProject.Description)
				require.NotEqual(t, *up1Pro1.StorageLimit, *updatedProject.StorageLimit)
				require.Equal(t, updatedStorageLimit, *updatedProject.StorageLimit)
				require.NotEqual(t, *up1Pro1.BandwidthLimit, *updatedProject.BandwidthLimit)
				require.Equal(t, updatedBandwidthLimit, *updatedProject.BandwidthLimit)

				// Updating someone else project details should not work
				updatedProject, err = service.UpdateProject(userCtx1, up2Pro1.ID, console.ProjectInfo{
					Name:           "newName",
					Description:    "TestUpdate",
					StorageLimit:   memory.Size(100),
					BandwidthLimit: memory.Size(100),
				})
				require.Error(t, err)
				require.Nil(t, updatedProject)

				// attempting to update a project with bandwidth or storage limits set to 0 should fail
				size0 := new(memory.Size)
				*size0 = 0
				size100 := new(memory.Size)
				*size100 = memory.Size(100)

				up1Pro1.StorageLimit = size0
				err = sat.DB.Console().Projects().Update(ctx, up1Pro1)
				require.NoError(t, err)

				updateInfo := console.ProjectInfo{
					Name:           "a b c",
					Description:    "1 2 3",
					StorageLimit:   memory.Size(123),
					BandwidthLimit: memory.Size(123),
				}
				updatedProject, err = service.UpdateProject(userCtx1, up1Pro1.ID, updateInfo)
				require.Error(t, err)
				require.Nil(t, updatedProject)

				up1Pro1.StorageLimit = size100
				up1Pro1.BandwidthLimit = size0

				err = sat.DB.Console().Projects().Update(ctx, up1Pro1)
				require.NoError(t, err)

				updatedProject, err = service.UpdateProject(userCtx1, up1Pro1.ID, updateInfo)
				require.Error(t, err)
				require.Nil(t, updatedProject)

				up1Pro1.StorageLimit = size100
				up1Pro1.BandwidthLimit = size100
				err = sat.DB.Console().Projects().Update(ctx, up1Pro1)
				require.NoError(t, err)

				updatedProject, err = service.UpdateProject(userCtx1, up1Pro1.ID, updateInfo)
				require.NoError(t, err)
				require.Equal(t, updateInfo.Name, updatedProject.Name)
				require.Equal(t, updateInfo.Description, updatedProject.Description)
				require.NotNil(t, updatedProject.StorageLimit)
				require.NotNil(t, updatedProject.BandwidthLimit)
				require.Equal(t, updateInfo.StorageLimit, *updatedProject.StorageLimit)
				require.Equal(t, updateInfo.BandwidthLimit, *updatedProject.BandwidthLimit)

				project, err := service.GetProject(userCtx1, up1Pro1.ID)
				require.NoError(t, err)
				require.Equal(t, updateInfo.StorageLimit, *project.StorageLimit)
				require.Equal(t, updateInfo.BandwidthLimit, *project.BandwidthLimit)
			})

			t.Run("AddProjectMembers", func(t *testing.T) {
				// Adding members to own project should work
				addedUsers, err := service.AddProjectMembers(userCtx1, up1Pro1.ID, []string{up2User.Email})
				require.NoError(t, err)
				require.Len(t, addedUsers, 1)
				require.Contains(t, addedUsers, up2User)

				// Adding members to someone else project should not work
				addedUsers, err = service.AddProjectMembers(userCtx1, up2Pro1.ID, []string{up2User.Email})
				require.Error(t, err)
				require.Nil(t, addedUsers)
			})

			t.Run("GetProjectMembers", func(t *testing.T) {
				// Getting the project members of an own project that one is a part of should work
				userPage, err := service.GetProjectMembers(userCtx1, up1Pro1.ID, console.ProjectMembersCursor{Page: 1, Limit: 10})
				require.NoError(t, err)
				require.Len(t, userPage.ProjectMembers, 2)

				// Getting the project members of a foreign project that one is a part of should work
				userPage, err = service.GetProjectMembers(userCtx2, up1Pro1.ID, console.ProjectMembersCursor{Page: 1, Limit: 10})
				require.NoError(t, err)
				require.Len(t, userPage.ProjectMembers, 2)

				// Getting the project members of a foreign project that one is not a part of should not work
				userPage, err = service.GetProjectMembers(userCtx1, up2Pro1.ID, console.ProjectMembersCursor{Page: 1, Limit: 10})
				require.Error(t, err)
				require.Nil(t, userPage)
			})

			t.Run("DeleteProjectMembers", func(t *testing.T) {
				// Deleting project members of an own project should work
				err := service.DeleteProjectMembers(userCtx1, up1Pro1.ID, []string{up2User.Email})
				require.NoError(t, err)

				// Deleting Project members of someone else project should not work
				err = service.DeleteProjectMembers(userCtx1, up2Pro1.ID, []string{up2User.Email})
				require.Error(t, err)
			})

			t.Run("DeleteProject", func(t *testing.T) {
				// Deleting the own project should not work before deleting the API-Key
				err := service.DeleteProject(userCtx1, up1Pro1.ID)
				require.Error(t, err)

				keys, err := service.GetAPIKeys(userCtx1, up1Pro1.ID, console.APIKeyCursor{Page: 1, Limit: 10})
				require.NoError(t, err)
				require.Len(t, keys.APIKeys, 1)

				err = service.DeleteAPIKeys(userCtx1, []uuid.UUID{keys.APIKeys[0].ID})
				require.NoError(t, err)

				// Deleting the own project should now work
				err = service.DeleteProject(userCtx1, up1Pro1.ID)
				require.NoError(t, err)

				// Deleting someone else project should not work
				err = service.DeleteProject(userCtx1, up2Pro1.ID)
				require.Error(t, err)

				err = planet.Uplinks[1].CreateBucket(ctx, sat, "testbucket")
				require.NoError(t, err)

				// deleting a project with a bucket should fail
				err = service.DeleteProject(userCtx2, up2Pro1.ID)
				require.Error(t, err)
				require.Equal(t, "console service: project usage: some buckets still exist", err.Error())
			})

			t.Run("ChangeEmail", func(t *testing.T) {
				const newEmail = "newEmail@example.com"

				err = service.ChangeEmail(userCtx2, newEmail)
				require.NoError(t, err)

				user, _, err := service.GetUserByEmailWithUnverified(userCtx2, newEmail)
				require.NoError(t, err)
				require.Equal(t, newEmail, user.Email)

				err = service.ChangeEmail(userCtx2, newEmail)
				require.Error(t, err)
			})

			t.Run("GetAllBucketNames", func(t *testing.T) {
				bucket1 := storj.Bucket{
					ID:        testrand.UUID(),
					Name:      "testBucket1",
					ProjectID: up2Pro1.ID,
				}

				bucket2 := storj.Bucket{
					ID:        testrand.UUID(),
					Name:      "testBucket2",
					ProjectID: up2Pro1.ID,
				}

				_, err := sat.API.Buckets.Service.CreateBucket(userCtx2, bucket1)
				require.NoError(t, err)

				_, err = sat.API.Buckets.Service.CreateBucket(userCtx2, bucket2)
				require.NoError(t, err)

				bucketNames, err := service.GetAllBucketNames(userCtx2, up2Pro1.ID)
				require.NoError(t, err)
				require.Equal(t, bucket1.Name, bucketNames[0])
				require.Equal(t, bucket2.Name, bucketNames[1])

				// Getting someone else buckets should not work
				bucketsForUnauthorizedUser, err := service.GetAllBucketNames(userCtx1, up2Pro1.ID)
				require.Error(t, err)
				require.Nil(t, bucketsForUnauthorizedUser)
			})

			t.Run("DeleteAPIKeyByNameAndProjectID", func(t *testing.T) {
				secret, err := macaroon.NewSecret()
				require.NoError(t, err)

				key, err := macaroon.NewAPIKey(secret)
				require.NoError(t, err)

				apikey := console.APIKeyInfo{
					Name:      "test",
					ProjectID: up2Pro1.ID,
					Secret:    secret,
				}

				createdKey, err := sat.DB.Console().APIKeys().Create(ctx, key.Head(), apikey)
				require.NoError(t, err)

				info, err := sat.DB.Console().APIKeys().Get(ctx, createdKey.ID)
				require.NoError(t, err)
				require.NotNil(t, info)

				// Deleting someone else api keys should not work
				err = service.DeleteAPIKeyByNameAndProjectID(userCtx1, apikey.Name, up2Pro1.ID)
				require.Error(t, err)

				err = service.DeleteAPIKeyByNameAndProjectID(userCtx2, apikey.Name, up2Pro1.ID)
				require.NoError(t, err)

				info, err = sat.DB.Console().APIKeys().Get(ctx, createdKey.ID)
				require.Error(t, err)
				require.Nil(t, info)
			})
		})
}

func TestPaidTier(t *testing.T) {
	usageConfig := console.UsageLimitsConfig{
		Storage: console.StorageLimitConfig{
			Free: memory.GB,
			Paid: memory.TB,
		},
		Bandwidth: console.BandwidthLimitConfig{
			Free: 2 * memory.GB,
			Paid: 2 * memory.TB,
		},
		Segment: console.SegmentLimitConfig{
			Free: 10,
			Paid: 50,
		},
		Project: console.ProjectLimitConfig{
			Free: 1,
			Paid: 3,
		},
	}

	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.Console.UsageLimits = usageConfig
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		sat := planet.Satellites[0]
		service := sat.API.Console.Service

		// project should have free tier usage limits
		proj1, err := sat.API.DB.Console().Projects().Get(ctx, planet.Uplinks[0].Projects[0].ID)
		require.NoError(t, err)
		require.Equal(t, usageConfig.Storage.Free, *proj1.StorageLimit)
		require.Equal(t, usageConfig.Bandwidth.Free, *proj1.BandwidthLimit)
		require.Equal(t, usageConfig.Segment.Free, *proj1.SegmentLimit)

		// user should be in free tier
		user, err := service.GetUser(ctx, proj1.OwnerID)
		require.NoError(t, err)
		require.False(t, user.PaidTier)

		userCtx, err := sat.UserContext(ctx, user.ID)
		require.NoError(t, err)

		// add a credit card to the user
		err = service.Payments().AddCreditCard(userCtx, "test-cc-token")
		require.NoError(t, err)

		// expect user to be in paid tier
		user, err = service.GetUser(ctx, user.ID)
		require.NoError(t, err)
		require.True(t, user.PaidTier)
		require.Equal(t, usageConfig.Project.Paid, user.ProjectLimit)

		// update auth ctx
		userCtx, err = sat.UserContext(ctx, user.ID)
		require.NoError(t, err)

		// expect project to be migrated to paid tier usage limits
		proj1, err = service.GetProject(userCtx, proj1.ID)
		require.NoError(t, err)
		require.Equal(t, usageConfig.Storage.Paid, *proj1.StorageLimit)
		require.Equal(t, usageConfig.Bandwidth.Paid, *proj1.BandwidthLimit)
		require.Equal(t, usageConfig.Segment.Paid, *proj1.SegmentLimit)

		// expect new project to be created with paid tier usage limits
		proj2, err := service.CreateProject(userCtx, console.ProjectInfo{Name: "Project 2"})
		require.NoError(t, err)
		require.Equal(t, usageConfig.Storage.Paid, *proj2.StorageLimit)
	})
}

// TestUpdateProjectExceedsLimits ensures that a project with limits manually set above the defaults can be updated.
func TestUpdateProjectExceedsLimits(t *testing.T) {
	usageConfig := console.UsageLimitsConfig{
		Storage: console.StorageLimitConfig{
			Free: memory.GB,
			Paid: memory.TB,
		},
		Bandwidth: console.BandwidthLimitConfig{
			Free: 2 * memory.GB,
			Paid: 2 * memory.TB,
		},
		Segment: console.SegmentLimitConfig{
			Free: 10,
			Paid: 50,
		},
		Project: console.ProjectLimitConfig{
			Free: 1,
			Paid: 3,
		},
	}

	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.Console.UsageLimits = usageConfig
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		sat := planet.Satellites[0]
		service := sat.API.Console.Service
		projectID := planet.Uplinks[0].Projects[0].ID

		updatedName := "newName"
		updatedDescription := "newDescription"
		updatedStorageLimit := memory.Size(100) + memory.TB
		updatedBandwidthLimit := memory.Size(100) + memory.TB

		proj, err := sat.API.DB.Console().Projects().Get(ctx, projectID)
		require.NoError(t, err)

		userCtx1, err := sat.UserContext(ctx, proj.OwnerID)
		require.NoError(t, err)

		// project should have free tier usage limits
		require.Equal(t, usageConfig.Storage.Free, *proj.StorageLimit)
		require.Equal(t, usageConfig.Bandwidth.Free, *proj.BandwidthLimit)
		require.Equal(t, usageConfig.Segment.Free, *proj.SegmentLimit)

		// update project name should succeed
		_, err = service.UpdateProject(userCtx1, projectID, console.ProjectInfo{
			Name:        updatedName,
			Description: updatedDescription,
		})
		require.NoError(t, err)

		// manually set project limits above defaults
		proj1, err := sat.API.DB.Console().Projects().Get(ctx, projectID)
		require.NoError(t, err)
		proj1.StorageLimit = new(memory.Size)
		*proj1.StorageLimit = updatedStorageLimit
		proj1.BandwidthLimit = new(memory.Size)
		*proj1.BandwidthLimit = updatedBandwidthLimit
		err = sat.DB.Console().Projects().Update(ctx, proj1)
		require.NoError(t, err)

		// try to update project name should succeed
		_, err = service.UpdateProject(userCtx1, projectID, console.ProjectInfo{
			Name:        "updatedName",
			Description: "updatedDescription",
		})
		require.NoError(t, err)
	})
}

func TestMFA(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 0,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		sat := planet.Satellites[0]
		service := sat.API.Console.Service

		user, err := sat.AddUser(ctx, console.CreateUser{
			FullName: "MFA Test User",
			Email:    "mfauser@mail.test",
		}, 1)
		require.NoError(t, err)

		updateContext := func() (context.Context, *console.User) {
			userCtx, err := sat.UserContext(ctx, user.ID)
			require.NoError(t, err)
			user, err := console.GetUser(userCtx)
			require.NoError(t, err)
			return userCtx, user
		}
		userCtx, user := updateContext()

		var key string
		t.Run("ResetMFASecretKey", func(t *testing.T) {
			key, err = service.ResetMFASecretKey(userCtx)
			require.NoError(t, err)

			_, user := updateContext()
			require.NotEmpty(t, user.MFASecretKey)
		})

		t.Run("EnableUserMFABadPasscode", func(t *testing.T) {
			// Expect MFA-enabling attempt to be rejected when providing stale passcode.
			badCode, err := console.NewMFAPasscode(key, time.Time{}.Add(time.Hour))
			require.NoError(t, err)

			err = service.EnableUserMFA(userCtx, badCode, time.Time{})
			require.True(t, console.ErrValidation.Has(err))

			userCtx, _ = updateContext()
			_, err = service.ResetMFARecoveryCodes(userCtx)
			require.True(t, console.ErrUnauthorized.Has(err))

			_, user = updateContext()
			require.False(t, user.MFAEnabled)
		})

		t.Run("EnableUserMFAGoodPasscode", func(t *testing.T) {
			// Expect MFA-enabling attempt to succeed when providing valid passcode.
			goodCode, err := console.NewMFAPasscode(key, time.Time{})
			require.NoError(t, err)

			userCtx, _ = updateContext()
			err = service.EnableUserMFA(userCtx, goodCode, time.Time{})
			require.NoError(t, err)

			_, user = updateContext()
			require.True(t, user.MFAEnabled)
			require.Equal(t, user.MFASecretKey, key)
		})

		t.Run("MFAGetToken", func(t *testing.T) {
			request := console.AuthUser{Email: user.Email, Password: user.FullName}

			// Expect no token due to lack of MFA passcode.
			token, err := service.Token(ctx, request)
			require.True(t, console.ErrMFAMissing.Has(err))
			require.Empty(t, token)

			// Expect no token due to bad MFA passcode.
			wrongCode, err := console.NewMFAPasscode(key, time.Now().Add(time.Hour))
			require.NoError(t, err)

			request.MFAPasscode = wrongCode
			token, err = service.Token(ctx, request)
			require.True(t, console.ErrMFAPasscode.Has(err))
			require.Empty(t, token)

			// Expect token when providing valid passcode.
			goodCode, err := console.NewMFAPasscode(key, time.Now())
			require.NoError(t, err)

			request.MFAPasscode = goodCode
			token, err = service.Token(ctx, request)
			require.NoError(t, err)
			require.NotEmpty(t, token)
		})

		t.Run("MFARecoveryCodes", func(t *testing.T) {
			_, err = service.ResetMFARecoveryCodes(userCtx)
			require.NoError(t, err)

			_, user = updateContext()
			require.Len(t, user.MFARecoveryCodes, console.MFARecoveryCodeCount)

			for _, code := range user.MFARecoveryCodes {
				// Ensure code is of the form XXXX-XXXX-XXXX where X is A-Z or 0-9.
				require.Regexp(t, "^([A-Z0-9]{4})((-[A-Z0-9]{4})){2}$", code)

				// Expect token when providing valid recovery code.
				request := console.AuthUser{Email: user.Email, Password: user.FullName, MFARecoveryCode: code}
				token, err := service.Token(ctx, request)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				// Expect no token due to providing previously-used recovery code.
				token, err = service.Token(ctx, request)
				require.True(t, console.ErrMFARecoveryCode.Has(err))
				require.Empty(t, token)

				_, user = updateContext()
			}

			userCtx, _ = updateContext()
			_, err = service.ResetMFARecoveryCodes(userCtx)
			require.NoError(t, err)
		})

		t.Run("DisableUserMFABadPasscode", func(t *testing.T) {
			// Expect MFA-disabling attempt to fail when providing valid passcode.
			badCode, err := console.NewMFAPasscode(key, time.Time{}.Add(time.Hour))
			require.NoError(t, err)

			userCtx, _ = updateContext()
			err = service.DisableUserMFA(userCtx, badCode, time.Time{}, "")
			require.True(t, console.ErrValidation.Has(err))

			_, user = updateContext()
			require.True(t, user.MFAEnabled)
			require.NotEmpty(t, user.MFASecretKey)
			require.NotEmpty(t, user.MFARecoveryCodes)
		})

		t.Run("DisableUserMFAConflict", func(t *testing.T) {
			// Expect MFA-disabling attempt to fail when providing both recovery code and passcode.
			goodCode, err := console.NewMFAPasscode(key, time.Time{})
			require.NoError(t, err)

			userCtx, user = updateContext()
			err = service.DisableUserMFA(userCtx, goodCode, time.Time{}, user.MFARecoveryCodes[0])
			require.True(t, console.ErrMFAConflict.Has(err))

			_, user = updateContext()
			require.True(t, user.MFAEnabled)
			require.NotEmpty(t, user.MFASecretKey)
			require.NotEmpty(t, user.MFARecoveryCodes)
		})

		t.Run("DisableUserMFAGoodPasscode", func(t *testing.T) {
			// Expect MFA-disabling attempt to succeed when providing valid passcode.
			goodCode, err := console.NewMFAPasscode(key, time.Time{})
			require.NoError(t, err)

			userCtx, _ = updateContext()
			err = service.DisableUserMFA(userCtx, goodCode, time.Time{}, "")
			require.NoError(t, err)

			userCtx, user = updateContext()
			require.False(t, user.MFAEnabled)
			require.Empty(t, user.MFASecretKey)
			require.Empty(t, user.MFARecoveryCodes)
		})

		t.Run("DisableUserMFAGoodRecoveryCode", func(t *testing.T) {
			// Expect MFA-disabling attempt to succeed when providing valid recovery code.
			// Enable MFA
			key, err = service.ResetMFASecretKey(userCtx)
			require.NoError(t, err)

			goodCode, err := console.NewMFAPasscode(key, time.Time{})
			require.NoError(t, err)

			userCtx, _ = updateContext()
			err = service.EnableUserMFA(userCtx, goodCode, time.Time{})
			require.NoError(t, err)

			userCtx, _ = updateContext()
			_, err = service.ResetMFARecoveryCodes(userCtx)
			require.NoError(t, err)

			userCtx, user = updateContext()
			require.True(t, user.MFAEnabled)
			require.NotEmpty(t, user.MFASecretKey)
			require.NotEmpty(t, user.MFARecoveryCodes)

			// Disable MFA
			err = service.DisableUserMFA(userCtx, "", time.Time{}, user.MFARecoveryCodes[0])
			require.NoError(t, err)

			_, user = updateContext()
			require.False(t, user.MFAEnabled)
			require.Empty(t, user.MFASecretKey)
			require.Empty(t, user.MFARecoveryCodes)
		})
	})
}

func TestResetPassword(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 0,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		sat := planet.Satellites[0]
		service := sat.API.Console.Service

		user, err := sat.AddUser(ctx, console.CreateUser{
			FullName: "Test User",
			Email:    "test@mail.test",
		}, 1)
		require.NoError(t, err)

		newPass := user.FullName

		getNewResetToken := func() *console.ResetPasswordToken {
			token, err := sat.DB.Console().ResetPasswordTokens().Create(ctx, user.ID)
			require.NoError(t, err)
			require.NotNil(t, token)
			return token
		}
		token := getNewResetToken()

		// Expect error when providing bad token.
		err = service.ResetPassword(ctx, "badToken", newPass, "", "", token.CreatedAt)
		require.True(t, console.ErrRecoveryToken.Has(err))

		// Expect error when providing good but expired token.
		err = service.ResetPassword(ctx, token.Secret.String(), newPass, "", "", token.CreatedAt.Add(sat.Config.ConsoleAuth.TokenExpirationTime).Add(time.Second))
		require.True(t, console.ErrTokenExpiration.Has(err))

		// Expect error when providing good token with bad (too short) password.
		err = service.ResetPassword(ctx, token.Secret.String(), "bad", "", "", token.CreatedAt)
		require.True(t, console.ErrValidation.Has(err))

		// Expect success when providing good token and good password.
		err = service.ResetPassword(ctx, token.Secret.String(), newPass, "", "", token.CreatedAt)
		require.NoError(t, err)

		token = getNewResetToken()

		// Enable MFA.
		userCtx, err := sat.UserContext(ctx, user.ID)
		require.NoError(t, err)

		key, err := service.ResetMFASecretKey(userCtx)
		require.NoError(t, err)
		userCtx, err = sat.UserContext(ctx, user.ID)
		require.NoError(t, err)

		passcode, err := console.NewMFAPasscode(key, token.CreatedAt)
		require.NoError(t, err)

		err = service.EnableUserMFA(userCtx, passcode, token.CreatedAt)
		require.NoError(t, err)

		// Expect error when providing bad passcode.
		badPasscode, err := console.NewMFAPasscode(key, token.CreatedAt.Add(time.Hour))
		require.NoError(t, err)
		err = service.ResetPassword(ctx, token.Secret.String(), newPass, badPasscode, "", token.CreatedAt)
		require.True(t, console.ErrMFAPasscode.Has(err))

		for _, recoveryCode := range user.MFARecoveryCodes {
			// Expect success when providing bad passcode and good recovery code.
			err = service.ResetPassword(ctx, token.Secret.String(), newPass, badPasscode, recoveryCode, token.CreatedAt)
			require.NoError(t, err)
			token = getNewResetToken()

			// Expect error when providing bad passcode and already-used recovery code.
			err = service.ResetPassword(ctx, token.Secret.String(), newPass, badPasscode, recoveryCode, token.CreatedAt)
			require.True(t, console.ErrMFARecoveryCode.Has(err))
		}

		// Expect success when providing good passcode.
		err = service.ResetPassword(ctx, token.Secret.String(), newPass, passcode, "", token.CreatedAt)
		require.NoError(t, err)
	})
}

func TestRESTKeys(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		sat := planet.Satellites[0]
		service := sat.API.Console.Service

		proj1, err := sat.API.DB.Console().Projects().Get(ctx, planet.Uplinks[0].Projects[0].ID)
		require.NoError(t, err)

		user, err := service.GetUser(ctx, proj1.OwnerID)
		require.NoError(t, err)

		userCtx, err := sat.UserContext(ctx, user.ID)
		require.NoError(t, err)

		now := time.Now()
		expires := 5 * time.Hour
		apiKey, expiresAt, err := service.CreateRESTKey(userCtx, expires)
		require.NoError(t, err)
		require.NotEmpty(t, apiKey)
		require.True(t, expiresAt.After(now))
		require.True(t, expiresAt.Before(now.Add(expires+time.Hour)))

		// test revocation
		require.NoError(t, service.RevokeRESTKey(userCtx, apiKey))

		// test revoke non existent key
		nonexistent := testrand.UUID()
		err = service.RevokeRESTKey(userCtx, nonexistent.String())
		require.Error(t, err)
	})
}

// TestLockAccount ensures user's gets locked when incorrect credentials are provided.
func TestLockAccount(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 0,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		sat := planet.Satellites[0]
		service := sat.API.Console.Service
		usersDB := sat.DB.Console().Users()
		consoleConfig := sat.Config.Console

		newUser := console.CreateUser{
			FullName: "token test",
			Email:    "token_test@mail.test",
		}

		user, err := sat.AddUser(ctx, newUser, 1)
		require.NoError(t, err)

		updateContext := func() (context.Context, *console.User) {
			userCtx, err := sat.UserContext(ctx, user.ID)
			require.NoError(t, err)
			user, err := console.GetUser(userCtx)
			require.NoError(t, err)
			return userCtx, user
		}

		userCtx, _ := updateContext()
		secret, err := service.ResetMFASecretKey(userCtx)
		require.NoError(t, err)

		goodCode0, err := console.NewMFAPasscode(secret, time.Time{})
		require.NoError(t, err)

		userCtx, _ = updateContext()
		err = service.EnableUserMFA(userCtx, goodCode0, time.Time{})
		require.NoError(t, err)

		now := time.Now()

		goodCode1, err := console.NewMFAPasscode(secret, now)
		require.NoError(t, err)

		authUser := console.AuthUser{
			Email:       newUser.Email,
			Password:    newUser.FullName,
			MFAPasscode: goodCode1,
		}

		// successful login.
		token, err := service.Token(ctx, authUser)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		// check if user's account gets locked because of providing wrong password.
		authUser.Password = "qweQWE1@"
		for i := 1; i <= consoleConfig.LoginAttemptsWithoutPenalty; i++ {
			token, err = service.Token(ctx, authUser)
			require.Empty(t, token)
			require.True(t, console.ErrLoginPassword.Has(err))
		}

		lockedUser, err := service.GetUser(userCtx, user.ID)
		require.NoError(t, err)
		require.True(t, lockedUser.FailedLoginCount == consoleConfig.LoginAttemptsWithoutPenalty)
		require.True(t, lockedUser.LoginLockoutExpiration.After(now))

		// lock account once again and check if lockout expiration time increased.
		err = service.UpdateUsersFailedLoginState(userCtx, lockedUser)
		require.NoError(t, err)

		lockedUser, err = service.GetUser(userCtx, user.ID)
		require.NoError(t, err)
		require.True(t, lockedUser.FailedLoginCount == consoleConfig.LoginAttemptsWithoutPenalty+1)

		diff := lockedUser.LoginLockoutExpiration.Sub(now)
		require.Greater(t, diff, time.Duration(consoleConfig.FailedLoginPenalty)*time.Minute)

		// unlock account by successful login
		lockedUser.LoginLockoutExpiration = now.Add(-time.Second)
		lockoutExpirationPtr := &lockedUser.LoginLockoutExpiration
		err = usersDB.Update(userCtx, lockedUser.ID, console.UpdateUserRequest{
			LoginLockoutExpiration: &lockoutExpirationPtr,
		})
		require.NoError(t, err)

		authUser.Password = newUser.FullName
		token, err = service.Token(ctx, authUser)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		unlockedUser, err := service.GetUser(userCtx, user.ID)
		require.NoError(t, err)
		require.Zero(t, unlockedUser.FailedLoginCount)

		// check if user's account gets locked because of providing wrong mfa passcode.
		authUser.MFAPasscode = "000000"
		for i := 1; i <= consoleConfig.LoginAttemptsWithoutPenalty; i++ {
			token, err = service.Token(ctx, authUser)
			require.Empty(t, token)
			require.True(t, console.ErrMFAPasscode.Has(err))
		}

		lockedUser, err = service.GetUser(userCtx, user.ID)
		require.NoError(t, err)
		require.True(t, lockedUser.FailedLoginCount == consoleConfig.LoginAttemptsWithoutPenalty)
		require.True(t, lockedUser.LoginLockoutExpiration.After(now))

		// unlock account
		lockedUser.LoginLockoutExpiration = time.Time{}
		lockoutExpirationPtr = &lockedUser.LoginLockoutExpiration
		lockedUser.FailedLoginCount = 0
		err = usersDB.Update(userCtx, lockedUser.ID, console.UpdateUserRequest{
			LoginLockoutExpiration: &lockoutExpirationPtr,
			FailedLoginCount:       &lockedUser.FailedLoginCount,
		})
		require.NoError(t, err)

		// check if user's account gets locked because of providing wrong mfa recovery code.
		authUser.MFAPasscode = ""
		authUser.MFARecoveryCode = "000000"
		for i := 1; i <= consoleConfig.LoginAttemptsWithoutPenalty; i++ {
			token, err = service.Token(ctx, authUser)
			require.Empty(t, token)
			require.True(t, console.ErrMFARecoveryCode.Has(err))
		}

		lockedUser, err = service.GetUser(userCtx, user.ID)
		require.NoError(t, err)
		require.True(t, lockedUser.FailedLoginCount == consoleConfig.LoginAttemptsWithoutPenalty)
		require.True(t, lockedUser.LoginLockoutExpiration.After(now))
	})
}

func TestWalletJsonMarshall(t *testing.T) {
	wi := console.WalletInfo{
		Address: blockchain.Address{1, 2, 3},
		Balance: currency.AmountFromBaseUnits(10000, currency.USDollars),
	}

	out, err := json.Marshal(wi)
	require.NoError(t, err)
	require.Contains(t, string(out), "\"address\":\"0102030000000000000000000000000000000000\"")
	require.Contains(t, string(out), "\"balance\":{\"value\":\"100\",\"currency\":\"USD\"}")

}

func TestSessionExpiration(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 0,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.Console.Session.Duration = time.Hour
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		sat := planet.Satellites[0]
		service := sat.API.Console.Service

		user, err := sat.AddUser(ctx, console.CreateUser{
			FullName: "Test User",
			Email:    "test@mail.test",
		}, 1)
		require.NoError(t, err)

		// Session should be added to DB after token request
		tokenInfo, err := service.Token(ctx, console.AuthUser{Email: user.Email, Password: user.FullName})
		require.NoError(t, err)

		_, err = service.TokenAuth(ctx, tokenInfo.Token, time.Now())
		require.NoError(t, err)

		sessionID, err := uuid.FromBytes(tokenInfo.Token.Payload)
		require.NoError(t, err)

		_, err = sat.DB.Console().WebappSessions().GetBySessionID(ctx, sessionID)
		require.NoError(t, err)

		// Session should be removed from DB after it has expired
		_, err = service.TokenAuth(ctx, tokenInfo.Token, time.Now().Add(2*time.Hour))
		require.True(t, console.ErrTokenExpiration.Has(err))

		_, err = sat.DB.Console().WebappSessions().GetBySessionID(ctx, sessionID)
		require.ErrorIs(t, sql.ErrNoRows, err)
	})
}

func TestPaymentsWalletPayments(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 0, UplinkCount: 0,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		sat := planet.Satellites[0]
		now := time.Now().Truncate(time.Second).UTC()

		user, err := sat.AddUser(ctx, console.CreateUser{
			FullName: "Test User",
			Email:    "test@mail.test",
			Password: "example",
		}, 1)
		require.NoError(t, err)

		wallet := blockchaintest.NewAddress()
		err = sat.DB.Wallets().Add(ctx, user.ID, wallet)
		require.NoError(t, err)

		var transactions []stripecoinpayments.Transaction
		for i := 0; i < 5; i++ {
			tx := stripecoinpayments.Transaction{
				ID:        coinpayments.TransactionID(fmt.Sprintf("%d", i)),
				AccountID: user.ID,
				Address:   blockchaintest.NewAddress().Hex(),
				Amount:    currency.AmountFromBaseUnits(1000000000, currency.StorjToken),
				Received:  currency.AmountFromBaseUnits(1000000000, currency.StorjToken),
				Status:    coinpayments.StatusCompleted,
				Key:       "key",
				Timeout:   0,
			}

			createdAt, err := sat.DB.StripeCoinPayments().Transactions().Insert(ctx, tx)
			require.NoError(t, err)
			err = sat.DB.StripeCoinPayments().Transactions().LockRate(ctx, tx.ID, decimal.NewFromInt(1))
			require.NoError(t, err)

			tx.CreatedAt = createdAt.UTC()
			transactions = append(transactions, tx)
		}

		var cachedPayments []storjscan.CachedPayment
		for i := 0; i < 10; i++ {
			cachedPayments = append(cachedPayments, storjscan.CachedPayment{
				From:        blockchaintest.NewAddress(),
				To:          wallet,
				TokenValue:  currency.AmountFromBaseUnits(1000, currency.StorjToken),
				USDValue:    currency.AmountFromBaseUnits(1000, currency.USDollarsMicro),
				Status:      payments.PaymentStatusConfirmed,
				BlockHash:   blockchaintest.NewHash(),
				BlockNumber: int64(i),
				Transaction: blockchaintest.NewHash(),
				LogIndex:    0,
				Timestamp:   now.Add(time.Duration(i) * 15 * time.Second),
			})
		}
		err = sat.DB.StorjscanPayments().InsertBatch(ctx, cachedPayments)
		require.NoError(t, err)

		reqCtx := console.WithUser(ctx, user)

		sort.Slice(cachedPayments, func(i, j int) bool {
			return cachedPayments[i].BlockNumber > cachedPayments[j].BlockNumber
		})
		sort.Slice(transactions, func(i, j int) bool {
			return transactions[i].CreatedAt.After(transactions[j].CreatedAt)
		})

		var expected []console.PaymentInfo
		for _, pmnt := range cachedPayments {
			expected = append(expected, console.PaymentInfo{
				ID:        fmt.Sprintf("%s#%d", pmnt.Transaction.Hex(), pmnt.LogIndex),
				Type:      "storjscan",
				Wallet:    pmnt.To.Hex(),
				Amount:    pmnt.USDValue,
				Status:    string(pmnt.Status),
				Link:      console.EtherscanURL(pmnt.Transaction.Hex()),
				Timestamp: pmnt.Timestamp,
			})
		}
		for _, tx := range transactions {
			expected = append(expected, console.PaymentInfo{
				ID:        tx.ID.String(),
				Type:      "coinpayments",
				Wallet:    tx.Address,
				Amount:    currency.AmountFromBaseUnits(1000, currency.USDollars),
				Received:  currency.AmountFromBaseUnits(1000, currency.USDollars),
				Status:    tx.Status.String(),
				Link:      coinpayments.GetCheckoutURL(tx.Key, tx.ID),
				Timestamp: tx.CreatedAt,
			})
		}

		walletPayments, err := sat.API.Console.Service.Payments().WalletPayments(reqCtx)
		require.NoError(t, err)
		require.Equal(t, expected, walletPayments.Payments)
	})
}
