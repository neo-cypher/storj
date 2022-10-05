// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleweb

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"mime"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"storj.io/common/errs2"
	"storj.io/common/memory"
	"storj.io/common/storj"
	"storj.io/common/uuid"
	"storj.io/storj/private/web"
	"storj.io/storj/satellite/analytics"
	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/console/consoleweb/consoleapi"
	"storj.io/storj/satellite/console/consoleweb/consoleql"
	"storj.io/storj/satellite/console/consoleweb/consolewebauth"
	"storj.io/storj/satellite/mailservice"
	"storj.io/storj/satellite/oidc"
	"storj.io/storj/satellite/payments/paymentsconfig"
	"storj.io/storj/satellite/rewards"
)

const (
	contentType = "Content-Type"

	applicationJSON    = "application/json"
	applicationGraphql = "application/graphql"
)

var (
	// Error is satellite console error type.
	Error = errs.Class("consoleweb")

	mon = monkit.Package()
)

// Config contains configuration for console web server.
type Config struct {
	Address         string `help:"server address of the graphql api gateway and frontend app" devDefault:"127.0.0.1:0" releaseDefault:":10100"`
	StaticDir       string `help:"path to static resources" default:""`
	Watch           bool   `help:"whether to load templates on each request" default:"false" devDefault:"true"`
	ExternalAddress string `help:"external endpoint of the satellite if hosted" default:""`

	AuthToken       string `help:"auth token needed for access to registration token creation endpoint" default:"" testDefault:"very-secret-token"`
	AuthTokenSecret string `help:"secret used to sign auth tokens" releaseDefault:"" devDefault:"my-suppa-secret-key"`

	ContactInfoURL                  string             `help:"url link to contacts page" default:"https://forum.storj.io"`
	FrameAncestors                  string             `help:"allow domains to embed the satellite in a frame, space separated" default:"tardigrade.io storj.io"`
	LetUsKnowURL                    string             `help:"url link to let us know page" default:"https://storjlabs.atlassian.net/servicedesk/customer/portals"`
	SEO                             string             `help:"used to communicate with web crawlers and other web robots" default:"User-agent: *\nDisallow: \nDisallow: /cgi-bin/"`
	SatelliteName                   string             `help:"used to display at web satellite console" default:"Storj"`
	SatelliteOperator               string             `help:"name of organization which set up satellite" default:"Storj Labs" `
	TermsAndConditionsURL           string             `help:"url link to terms and conditions page" default:"https://www.storj.io/terms-of-service/"`
	AccountActivationRedirectURL    string             `help:"url link for account activation redirect" default:""`
	PartneredSatellites             console.Satellites `help:"names and addresses of partnered satellites in JSON list format" default:"[{\"name\":\"US1\",\"address\":\"https://us1.storj.io\"},{\"name\":\"EU1\",\"address\":\"https://eu1.storj.io\"},{\"name\":\"AP1\",\"address\":\"https://ap1.storj.io\"}]"`
	GeneralRequestURL               string             `help:"url link to general request page" default:"https://supportdcs.storj.io/hc/en-us/requests/new?ticket_form_id=360000379291"`
	ProjectLimitsIncreaseRequestURL string             `help:"url link to project limit increase request page" default:"https://supportdcs.storj.io/hc/en-us/requests/new?ticket_form_id=360000683212"`
	GatewayCredentialsRequestURL    string             `help:"url link for gateway credentials requests" default:"https://auth.storjshare.io" devDefault:"http://localhost:8000"`
	IsBetaSatellite                 bool               `help:"indicates if satellite is in beta" default:"false"`
	BetaSatelliteFeedbackURL        string             `help:"url link for for beta satellite feedback" default:""`
	BetaSatelliteSupportURL         string             `help:"url link for for beta satellite support" default:""`
	DocumentationURL                string             `help:"url link to documentation" default:"https://docs.storj.io/"`
	CouponCodeBillingUIEnabled      bool               `help:"indicates if user is allowed to add coupon codes to account from billing" default:"false"`
	CouponCodeSignupUIEnabled       bool               `help:"indicates if user is allowed to add coupon codes to account from signup" default:"false"`
	FileBrowserFlowDisabled         bool               `help:"indicates if file browser flow is disabled" default:"false"`
	CSPEnabled                      bool               `help:"indicates if Content Security Policy is enabled" devDefault:"false" releaseDefault:"true"`
	LinksharingURL                  string             `help:"url link for linksharing requests" default:"https://link.storjshare.io" devDefault:"http://localhost:8001"`
	PathwayOverviewEnabled          bool               `help:"indicates if the overview onboarding step should render with pathways" default:"true"`
	NewProjectDashboard             bool               `help:"indicates if new project dashboard should be used" default:"true"`
	NewObjectsFlow                  bool               `help:"indicates if new objects flow should be used" default:"true"`
	NewAccessGrantFlow              bool               `help:"indicates if new access grant flow should be used" default:"true"`
	NewBillingScreen                bool               `help:"indicates if new billing screens should be used" default:"false"`
	GeneratedAPIEnabled             bool               `help:"indicates if generated console api should be used" default:"false"`
	OptionalSignupSuccessURL        string             `help:"optional url to external registration success page" default:""`
	HomepageURL                     string             `help:"url link to storj.io homepage" default:"https://www.storj.io"`
	NativeTokenPaymentsEnabled      bool               `help:"indicates if storj native token payments system is enabled" default:"false"`

	OauthCodeExpiry         time.Duration `help:"how long oauth authorization codes are issued for" default:"10m"`
	OauthAccessTokenExpiry  time.Duration `help:"how long oauth access tokens are issued for" default:"24h"`
	OauthRefreshTokenExpiry time.Duration `help:"how long oauth refresh tokens are issued for" default:"720h"`

	// RateLimit defines the configuration for the IP and userID rate limiters.
	RateLimit web.RateLimiterConfig

	console.Config
}

// Server represents console web server.
//
// architecture: Endpoint
type Server struct {
	log *zap.Logger

	config      Config
	service     *console.Service
	mailService *mailservice.Service
	partners    *rewards.PartnersService
	analytics   *analytics.Service

	listener          net.Listener
	server            http.Server
	cookieAuth        *consolewebauth.CookieAuth
	ipRateLimiter     *web.RateLimiter
	userIDRateLimiter *web.RateLimiter
	nodeURL           storj.NodeURL

	stripePublicKey string

	pricing paymentsconfig.PricingValues

	schema graphql.Schema

	templatesCache *templates
}

type templates struct {
	index               *template.Template
	notFound            *template.Template
	internalServerError *template.Template
	usageReport         *template.Template
}

// apiAuth exposes methods to control authentication process for each generated API endpoint.
type apiAuth struct {
	server *Server
}

// IsAuthenticated checks if request is performed with all needed authorization credentials.
func (a *apiAuth) IsAuthenticated(ctx context.Context, r *http.Request, isCookieAuth, isKeyAuth bool) (_ context.Context, err error) {
	if isCookieAuth && isKeyAuth {
		ctx, err = a.cookieAuth(ctx, r)
		if err != nil {
			ctx, err = a.keyAuth(ctx, r)
			if err != nil {
				return nil, err
			}
		}
	} else if isCookieAuth {
		ctx, err = a.cookieAuth(ctx, r)
		if err != nil {
			return nil, err
		}
	} else if isKeyAuth {
		ctx, err = a.keyAuth(ctx, r)
		if err != nil {
			return nil, err
		}
	}

	return ctx, nil
}

// cookieAuth returns an authenticated context by session cookie.
func (a *apiAuth) cookieAuth(ctx context.Context, r *http.Request) (context.Context, error) {
	tokenInfo, err := a.server.cookieAuth.GetToken(r)
	if err != nil {
		return nil, err
	}

	return a.server.service.TokenAuth(ctx, tokenInfo.Token, time.Now())
}

// cookieAuth returns an authenticated context by api key.
func (a *apiAuth) keyAuth(ctx context.Context, r *http.Request) (context.Context, error) {
	authToken := r.Header.Get("Authorization")
	split := strings.Split(authToken, "Bearer ")
	if len(split) != 2 {
		return ctx, errs.New("authorization key format is incorrect. Should be 'Bearer <key>'")
	}

	return a.server.service.KeyAuth(ctx, split[1], time.Now())
}

// RemoveAuthCookie indicates to the client that the authentication cookie should be removed.
func (a *apiAuth) RemoveAuthCookie(w http.ResponseWriter) {
	a.server.cookieAuth.RemoveTokenCookie(w)
}

// NewServer creates new instance of console server.
func NewServer(logger *zap.Logger, config Config, service *console.Service, oidcService *oidc.Service, mailService *mailservice.Service, partners *rewards.PartnersService, analytics *analytics.Service, listener net.Listener, stripePublicKey string, pricing paymentsconfig.PricingValues, nodeURL storj.NodeURL) *Server {
	server := Server{
		log:               logger,
		config:            config,
		listener:          listener,
		service:           service,
		mailService:       mailService,
		partners:          partners,
		analytics:         analytics,
		stripePublicKey:   stripePublicKey,
		ipRateLimiter:     web.NewIPRateLimiter(config.RateLimit),
		userIDRateLimiter: NewUserIDRateLimiter(config.RateLimit),
		nodeURL:           nodeURL,
		pricing:           pricing,
	}

	logger.Debug("Starting Satellite UI.", zap.Stringer("Address", server.listener.Addr()))

	server.cookieAuth = consolewebauth.NewCookieAuth(consolewebauth.CookieSettings{
		Name: "_tokenKey",
		Path: "/",
	})

	if server.config.ExternalAddress != "" {
		if !strings.HasSuffix(server.config.ExternalAddress, "/") {
			server.config.ExternalAddress += "/"
		}
	} else {
		server.config.ExternalAddress = "http://" + server.listener.Addr().String() + "/"
	}

	if server.config.AccountActivationRedirectURL == "" {
		server.config.AccountActivationRedirectURL = server.config.ExternalAddress + "login?activated=true"
	}

	router := mux.NewRouter()
	// N.B. This middleware has to be the first one because it has to be called
	// the earliest in the HTTP chain.
	router.Use(newTraceRequestMiddleware(logger, router))

	if server.config.GeneratedAPIEnabled {
		consoleapi.NewProjectManagement(logger, mon, server.service, router, &apiAuth{&server})
		consoleapi.NewAPIKeyManagement(logger, mon, server.service, router, &apiAuth{&server})
		consoleapi.NewUserManagement(logger, mon, server.service, router, &apiAuth{&server})
	}

	projectsController := consoleapi.NewProjects(logger, service)
	router.Handle(
		"/api/v0/projects/{id}/salt",
		server.withAuth(http.HandlerFunc(projectsController.GetSalt)),
	).Methods(http.MethodGet)

	router.HandleFunc("/registrationToken/", server.createRegistrationTokenHandler)
	router.HandleFunc("/robots.txt", server.seoHandler)

	router.Handle("/api/v0/graphql", server.withAuth(http.HandlerFunc(server.graphqlHandler)))

	usageLimitsController := consoleapi.NewUsageLimits(logger, service)
	router.Handle(
		"/api/v0/projects/{id}/usage-limits",
		server.withAuth(http.HandlerFunc(usageLimitsController.ProjectUsageLimits)),
	).Methods(http.MethodGet)
	router.Handle(
		"/api/v0/projects/usage-limits",
		server.withAuth(http.HandlerFunc(usageLimitsController.TotalUsageLimits)),
	).Methods(http.MethodGet)
	router.Handle(
		"/api/v0/projects/{id}/daily-usage",
		server.withAuth(http.HandlerFunc(usageLimitsController.DailyUsage)),
	).Methods(http.MethodGet)

	authController := consoleapi.NewAuth(logger, service, mailService, server.cookieAuth, partners, server.analytics, config.SatelliteName, server.config.ExternalAddress, config.LetUsKnowURL, config.TermsAndConditionsURL, config.ContactInfoURL, config.GeneralRequestURL)
	authRouter := router.PathPrefix("/api/v0/auth").Subrouter()
	authRouter.Handle("/account", server.withAuth(http.HandlerFunc(authController.GetAccount))).Methods(http.MethodGet)
	authRouter.Handle("/account", server.withAuth(http.HandlerFunc(authController.UpdateAccount))).Methods(http.MethodPatch)
	authRouter.Handle("/account/change-email", server.withAuth(http.HandlerFunc(authController.ChangeEmail))).Methods(http.MethodPost)
	authRouter.Handle("/account/change-password", server.withAuth(http.HandlerFunc(authController.ChangePassword))).Methods(http.MethodPost)
	authRouter.Handle("/account/delete", server.withAuth(http.HandlerFunc(authController.DeleteAccount))).Methods(http.MethodPost)
	authRouter.Handle("/mfa/enable", server.withAuth(http.HandlerFunc(authController.EnableUserMFA))).Methods(http.MethodPost)
	authRouter.Handle("/mfa/disable", server.withAuth(http.HandlerFunc(authController.DisableUserMFA))).Methods(http.MethodPost)
	authRouter.Handle("/mfa/generate-secret-key", server.withAuth(http.HandlerFunc(authController.GenerateMFASecretKey))).Methods(http.MethodPost)
	authRouter.Handle("/mfa/generate-recovery-codes", server.withAuth(http.HandlerFunc(authController.GenerateMFARecoveryCodes))).Methods(http.MethodPost)
	authRouter.Handle("/logout", server.withAuth(http.HandlerFunc(authController.Logout))).Methods(http.MethodPost)
	authRouter.Handle("/token", server.ipRateLimiter.Limit(http.HandlerFunc(authController.Token))).Methods(http.MethodPost)
	authRouter.Handle("/register", server.ipRateLimiter.Limit(http.HandlerFunc(authController.Register))).Methods(http.MethodPost, http.MethodOptions)
	authRouter.Handle("/forgot-password", server.ipRateLimiter.Limit(http.HandlerFunc(authController.ForgotPassword))).Methods(http.MethodPost)
	authRouter.Handle("/resend-email/{email}", server.ipRateLimiter.Limit(http.HandlerFunc(authController.ResendEmail))).Methods(http.MethodPost)
	authRouter.Handle("/reset-password", server.ipRateLimiter.Limit(http.HandlerFunc(authController.ResetPassword))).Methods(http.MethodPost)
	authRouter.Handle("/refresh-session", server.withAuth(http.HandlerFunc(authController.RefreshSession))).Methods(http.MethodPost)

	paymentController := consoleapi.NewPayments(logger, service)
	paymentsRouter := router.PathPrefix("/api/v0/payments").Subrouter()
	paymentsRouter.Use(server.withAuth)
	paymentsRouter.HandleFunc("/cards", paymentController.AddCreditCard).Methods(http.MethodPost)
	paymentsRouter.HandleFunc("/cards", paymentController.MakeCreditCardDefault).Methods(http.MethodPatch)
	paymentsRouter.HandleFunc("/cards", paymentController.ListCreditCards).Methods(http.MethodGet)
	paymentsRouter.HandleFunc("/cards/{cardId}", paymentController.RemoveCreditCard).Methods(http.MethodDelete)
	paymentsRouter.HandleFunc("/account/charges", paymentController.ProjectsCharges).Methods(http.MethodGet)
	paymentsRouter.HandleFunc("/account/balance", paymentController.AccountBalance).Methods(http.MethodGet)
	paymentsRouter.HandleFunc("/account", paymentController.SetupAccount).Methods(http.MethodPost)
	paymentsRouter.HandleFunc("/wallet", paymentController.GetWallet).Methods(http.MethodGet)
	paymentsRouter.HandleFunc("/wallet", paymentController.ClaimWallet).Methods(http.MethodPost)
	paymentsRouter.HandleFunc("/wallet/payments", paymentController.WalletPayments).Methods(http.MethodGet)
	paymentsRouter.HandleFunc("/billing-history", paymentController.BillingHistory).Methods(http.MethodGet)
	paymentsRouter.HandleFunc("/tokens/deposit", paymentController.TokenDeposit).Methods(http.MethodPost)
	paymentsRouter.Handle("/coupon/apply", server.userIDRateLimiter.Limit(http.HandlerFunc(paymentController.ApplyCouponCode))).Methods(http.MethodPatch)
	paymentsRouter.HandleFunc("/coupon", paymentController.GetCoupon).Methods(http.MethodGet)

	bucketsController := consoleapi.NewBuckets(logger, service)
	bucketsRouter := router.PathPrefix("/api/v0/buckets").Subrouter()
	bucketsRouter.Use(server.withAuth)
	bucketsRouter.HandleFunc("/bucket-names", bucketsController.AllBucketNames).Methods(http.MethodGet)

	apiKeysController := consoleapi.NewAPIKeys(logger, service)
	apiKeysRouter := router.PathPrefix("/api/v0/api-keys").Subrouter()
	apiKeysRouter.Use(server.withAuth)
	apiKeysRouter.HandleFunc("/delete-by-name", apiKeysController.DeleteByNameAndProjectID).Methods(http.MethodDelete)

	analyticsController := consoleapi.NewAnalytics(logger, service, server.analytics)
	analyticsRouter := router.PathPrefix("/api/v0/analytics").Subrouter()
	analyticsRouter.Use(server.withAuth)
	analyticsRouter.HandleFunc("/event", analyticsController.EventTriggered).Methods(http.MethodPost)
	analyticsRouter.HandleFunc("/page", analyticsController.PageEventTriggered).Methods(http.MethodPost)

	if server.config.StaticDir != "" {
		oidc := oidc.NewEndpoint(
			server.nodeURL, server.config.ExternalAddress,
			logger, oidcService, service,
			server.config.OauthCodeExpiry, server.config.OauthAccessTokenExpiry, server.config.OauthRefreshTokenExpiry,
		)

		router.HandleFunc("/.well-known/openid-configuration", oidc.WellKnownConfiguration)
		router.Handle("/oauth/v2/authorize", server.withAuth(http.HandlerFunc(oidc.AuthorizeUser))).Methods(http.MethodPost)
		router.Handle("/oauth/v2/tokens", server.ipRateLimiter.Limit(http.HandlerFunc(oidc.Tokens))).Methods(http.MethodPost)
		router.Handle("/oauth/v2/userinfo", server.ipRateLimiter.Limit(http.HandlerFunc(oidc.UserInfo))).Methods(http.MethodGet)
		router.Handle("/oauth/v2/clients/{id}", server.withAuth(http.HandlerFunc(oidc.GetClient))).Methods(http.MethodGet)

		fs := http.FileServer(http.Dir(server.config.StaticDir))
		router.PathPrefix("/static/").Handler(server.brotliMiddleware(http.StripPrefix("/static", fs)))

		router.HandleFunc("/activation/", server.accountActivationHandler)
		router.HandleFunc("/cancel-password-recovery/", server.cancelPasswordRecoveryHandler)
		router.HandleFunc("/usage-report", server.bucketUsageReportHandler)
		router.PathPrefix("/").Handler(http.HandlerFunc(server.appHandler))
	}

	server.server = http.Server{
		Handler:        server.withRequest(router),
		MaxHeaderBytes: ContentLengthLimit.Int(),
	}

	return &server
}

// Run starts the server that host webapp and api endpoint.
func (server *Server) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	server.schema, err = consoleql.CreateSchema(server.log, server.service, server.mailService)
	if err != nil {
		return Error.Wrap(err)
	}

	_, err = server.loadTemplates()
	if err != nil {
		// TODO: should it return error if some template can not be initialized or just log about it?
		return Error.Wrap(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group
	group.Go(func() error {
		<-ctx.Done()
		return server.server.Shutdown(context.Background())
	})
	group.Go(func() error {
		server.ipRateLimiter.Run(ctx)
		return nil
	})
	group.Go(func() error {
		defer cancel()
		err := server.server.Serve(server.listener)
		if errs2.IsCanceled(err) || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return err
	})

	return group.Wait()
}

// Close closes server and underlying listener.
func (server *Server) Close() error {
	return server.server.Close()
}

// appHandler is web app http handler function.
func (server *Server) appHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	if server.config.CSPEnabled {
		cspValues := []string{
			"default-src 'self'",
			"script-src 'sha256-wAqYV6m2PHGd1WDyFBnZmSoyfCK0jxFAns0vGbdiWUA=' 'self' *.stripe.com https://www.google.com/recaptcha/ https://www.gstatic.com/recaptcha/ https://hcaptcha.com *.hcaptcha.com",
			"connect-src 'self' *.tardigradeshare.io *.storjshare.io https://hcaptcha.com *.hcaptcha.com " + server.config.GatewayCredentialsRequestURL,
			"frame-ancestors " + server.config.FrameAncestors,
			"frame-src 'self' *.stripe.com https://www.google.com/recaptcha/ https://recaptcha.google.com/recaptcha/ https://hcaptcha.com *.hcaptcha.com",
			"img-src 'self' data: blob: *.tardigradeshare.io *.storjshare.io",
			// Those are hashes of charts custom tooltip inline styles. They have to be updated if styles are updated.
			"style-src 'unsafe-hashes' 'sha256-7mY2NKmZ4PuyjGUa4FYC5u36SxXdoUM/zxrlr3BEToo=' 'sha256-PRTMwLUW5ce9tdiUrVCGKqj6wPeuOwGogb1pmyuXhgI=' 'sha256-kwpt3lQZ21rs4cld7/uEm9qI5yAbjYzx+9FGm/XmwNU=' 'self' https://hcaptcha.com *.hcaptcha.com",
			"media-src 'self' blob: *.tardigradeshare.io *.storjshare.io",
		}

		header.Set("Content-Security-Policy", strings.Join(cspValues, "; "))
	}

	header.Set(contentType, "text/html; charset=UTF-8")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "same-origin") // Only expose the referring url when navigating around the satellite itself.

	var data struct {
		ExternalAddress                 string
		SatelliteName                   string
		SatelliteNodeURL                string
		StripePublicKey                 string
		PartneredSatellites             string
		DefaultProjectLimit             int
		GeneralRequestURL               string
		ProjectLimitsIncreaseRequestURL string
		GatewayCredentialsRequestURL    string
		IsBetaSatellite                 bool
		BetaSatelliteFeedbackURL        string
		BetaSatelliteSupportURL         string
		DocumentationURL                string
		CouponCodeBillingUIEnabled      bool
		CouponCodeSignupUIEnabled       bool
		FileBrowserFlowDisabled         bool
		LinksharingURL                  string
		PathwayOverviewEnabled          bool
		StorageTBPrice                  string
		EgressTBPrice                   string
		SegmentPrice                    string
		RegistrationRecaptchaEnabled    bool
		RegistrationRecaptchaSiteKey    string
		RegistrationHcaptchaEnabled     bool
		RegistrationHcaptchaSiteKey     string
		LoginRecaptchaEnabled           bool
		LoginRecaptchaSiteKey           string
		LoginHcaptchaEnabled            bool
		LoginHcaptchaSiteKey            string
		NewProjectDashboard             bool
		DefaultPaidStorageLimit         memory.Size
		DefaultPaidBandwidthLimit       memory.Size
		NewObjectsFlow                  bool
		NewAccessGrantFlow              bool
		NewBillingScreen                bool
		InactivityTimerEnabled          bool
		InactivityTimerDuration         int
		InactivityTimerViewerEnabled    bool
		OptionalSignupSuccessURL        string
		HomepageURL                     string
		NativeTokenPaymentsEnabled      bool
		PasswordMinimumLength           int
		PasswordMaximumLength           int
	}

	data.ExternalAddress = server.config.ExternalAddress
	data.SatelliteName = server.config.SatelliteName
	data.SatelliteNodeURL = server.nodeURL.String()
	data.StripePublicKey = server.stripePublicKey
	data.PartneredSatellites = server.config.PartneredSatellites.String()
	data.DefaultProjectLimit = server.config.DefaultProjectLimit
	data.GeneralRequestURL = server.config.GeneralRequestURL
	data.ProjectLimitsIncreaseRequestURL = server.config.ProjectLimitsIncreaseRequestURL
	data.GatewayCredentialsRequestURL = server.config.GatewayCredentialsRequestURL
	data.IsBetaSatellite = server.config.IsBetaSatellite
	data.BetaSatelliteFeedbackURL = server.config.BetaSatelliteFeedbackURL
	data.BetaSatelliteSupportURL = server.config.BetaSatelliteSupportURL
	data.DocumentationURL = server.config.DocumentationURL
	data.CouponCodeBillingUIEnabled = server.config.CouponCodeBillingUIEnabled
	data.CouponCodeSignupUIEnabled = server.config.CouponCodeSignupUIEnabled
	data.FileBrowserFlowDisabled = server.config.FileBrowserFlowDisabled
	data.LinksharingURL = server.config.LinksharingURL
	data.PathwayOverviewEnabled = server.config.PathwayOverviewEnabled
	data.DefaultPaidStorageLimit = server.config.UsageLimits.Storage.Paid
	data.DefaultPaidBandwidthLimit = server.config.UsageLimits.Bandwidth.Paid
	data.StorageTBPrice = server.pricing.StorageTBPrice
	data.EgressTBPrice = server.pricing.EgressTBPrice
	data.SegmentPrice = server.pricing.SegmentPrice
	data.RegistrationRecaptchaEnabled = server.config.Captcha.Registration.Recaptcha.Enabled
	data.RegistrationRecaptchaSiteKey = server.config.Captcha.Registration.Recaptcha.SiteKey
	data.RegistrationHcaptchaEnabled = server.config.Captcha.Registration.Hcaptcha.Enabled
	data.RegistrationHcaptchaSiteKey = server.config.Captcha.Registration.Hcaptcha.SiteKey
	data.LoginRecaptchaEnabled = server.config.Captcha.Login.Recaptcha.Enabled
	data.LoginRecaptchaSiteKey = server.config.Captcha.Login.Recaptcha.SiteKey
	data.LoginHcaptchaEnabled = server.config.Captcha.Login.Hcaptcha.Enabled
	data.LoginHcaptchaSiteKey = server.config.Captcha.Login.Hcaptcha.SiteKey
	data.NewProjectDashboard = server.config.NewProjectDashboard
	data.NewObjectsFlow = server.config.NewObjectsFlow
	data.NewAccessGrantFlow = server.config.NewAccessGrantFlow
	data.NewBillingScreen = server.config.NewBillingScreen
	data.InactivityTimerEnabled = server.config.Session.InactivityTimerEnabled
	data.InactivityTimerDuration = server.config.Session.InactivityTimerDuration
	data.InactivityTimerViewerEnabled = server.config.Session.InactivityTimerViewerEnabled
	data.OptionalSignupSuccessURL = server.config.OptionalSignupSuccessURL
	data.HomepageURL = server.config.HomepageURL
	data.NativeTokenPaymentsEnabled = server.config.NativeTokenPaymentsEnabled
	data.PasswordMinimumLength = console.PasswordMinimumLength
	data.PasswordMaximumLength = console.PasswordMaximumLength

	templates, err := server.loadTemplates()
	if err != nil || templates.index == nil {
		server.log.Error("unable to load templates", zap.Error(err))
		fmt.Fprintf(w, "Unable to load templates. See whether satellite UI has been built.")
		return
	}

	if err := templates.index.Execute(w, data); err != nil {
		server.log.Error("index template could not be executed", zap.Error(err))
		return
	}
}

// withAuth performs initial authorization before every request.
func (server *Server) withAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := r.Context()

		defer mon.Task()(&ctx)(&err)

		defer func() {
			if err != nil {
				consoleapi.ServeJSONError(server.log, w, http.StatusUnauthorized, console.ErrUnauthorized.Wrap(err))
				server.cookieAuth.RemoveTokenCookie(w)
			}
		}()

		tokenInfo, err := server.cookieAuth.GetToken(r)
		if err != nil {
			return
		}

		newCtx, err := server.service.TokenAuth(ctx, tokenInfo.Token, time.Now())
		if err != nil {
			return
		}
		ctx = newCtx

		handler.ServeHTTP(w, r.Clone(ctx))
	})
}

// withRequest ensures the http request itself is reachable from the context.
func (server *Server) withRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r.Clone(console.WithRequest(r.Context(), r)))
	})
}

// bucketUsageReportHandler generate bucket usage report page for project.
func (server *Server) bucketUsageReportHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	tokenInfo, err := server.cookieAuth.GetToken(r)
	if err != nil {
		server.serveError(w, http.StatusUnauthorized)
		return
	}

	ctx, err = server.service.TokenAuth(ctx, tokenInfo.Token, time.Now())
	if err != nil {
		server.serveError(w, http.StatusUnauthorized)
		return
	}

	// parse query params
	projectID, err := uuid.FromString(r.URL.Query().Get("projectID"))
	if err != nil {
		server.serveError(w, http.StatusBadRequest)
		return
	}
	sinceStamp, err := strconv.ParseInt(r.URL.Query().Get("since"), 10, 64)
	if err != nil {
		server.serveError(w, http.StatusBadRequest)
		return
	}
	beforeStamp, err := strconv.ParseInt(r.URL.Query().Get("before"), 10, 64)
	if err != nil {
		server.serveError(w, http.StatusBadRequest)
		return
	}

	since := time.Unix(sinceStamp, 0).UTC()
	before := time.Unix(beforeStamp, 0).UTC()

	server.log.Debug("querying bucket usage report",
		zap.Stringer("projectID", projectID),
		zap.Stringer("since", since),
		zap.Stringer("before", before))

	bucketRollups, err := server.service.GetBucketUsageRollups(ctx, projectID, since, before)
	if err != nil {
		server.log.Error("bucket usage report error", zap.Error(err))
		server.serveError(w, http.StatusInternalServerError)
		return
	}

	templates, err := server.loadTemplates()
	if err != nil {
		server.log.Error("unable to load templates", zap.Error(err))
		return
	}
	if err = templates.usageReport.Execute(w, bucketRollups); err != nil {
		server.log.Error("bucket usage report error", zap.Error(err))
	}
}

// createRegistrationTokenHandler is web app http handler function.
func (server *Server) createRegistrationTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)
	w.Header().Set(contentType, applicationJSON)

	var response struct {
		Secret string `json:"secret"`
		Error  string `json:"error,omitempty"`
	}

	defer func() {
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			server.log.Error(err.Error())
		}
	}()

	equality := subtle.ConstantTimeCompare(
		[]byte(r.Header.Get("Authorization")),
		[]byte(server.config.AuthToken),
	)
	if equality != 1 {
		w.WriteHeader(401)
		response.Error = "unauthorized"
		return
	}

	projectsLimitInput := r.URL.Query().Get("projectsLimit")

	projectsLimit, err := strconv.Atoi(projectsLimitInput)
	if err != nil {
		response.Error = err.Error()
		return
	}

	token, err := server.service.CreateRegToken(ctx, projectsLimit)
	if err != nil {
		response.Error = err.Error()
		return
	}

	response.Secret = token.Secret.String()
}

// accountActivationHandler is web app http handler function.
func (server *Server) accountActivationHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)
	activationToken := r.URL.Query().Get("token")

	user, err := server.service.ActivateAccount(ctx, activationToken)
	if err != nil {
		if console.ErrTokenInvalid.Has(err) {
			server.log.Debug("account activation",
				zap.String("token", activationToken),
				zap.Error(err),
			)
			server.serveError(w, http.StatusBadRequest)
			return
		}

		if console.ErrTokenExpiration.Has(err) {
			server.log.Debug("account activation",
				zap.String("token", activationToken),
				zap.Error(err),
			)
			server.serveError(w, http.StatusNotFound)
			return
		}

		if console.ErrEmailUsed.Has(err) {
			server.log.Debug("account activation",
				zap.String("token", activationToken),
				zap.Error(err),
			)
			http.Redirect(w, r, server.config.ExternalAddress+"login?activated=false", http.StatusTemporaryRedirect)
			return
		}

		if console.Error.Has(err) {
			server.log.Error("activation: failed to activate account with a valid token",
				zap.Error(err))
			server.serveError(w, http.StatusInternalServerError)
			return
		}

		server.log.Error(
			"activation: failed to activate account with a valid token and unknown error type. BUG: missed error type check",
			zap.Error(err))
		server.serveError(w, http.StatusInternalServerError)
		return
	}

	ip, err := web.GetRequestIP(r)
	if err != nil {
		server.serveError(w, http.StatusInternalServerError)
		return
	}

	tokenInfo, err := server.service.GenerateSessionToken(ctx, user.ID, user.Email, ip, r.UserAgent())
	if err != nil {
		server.serveError(w, http.StatusInternalServerError)
		return
	}

	server.cookieAuth.SetTokenCookie(w, *tokenInfo)

	http.Redirect(w, r, server.config.ExternalAddress, http.StatusTemporaryRedirect)
}

func (server *Server) cancelPasswordRecoveryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)
	recoveryToken := r.URL.Query().Get("token")

	// No need to check error as we anyway redirect user to support page
	_ = server.service.RevokeResetPasswordToken(ctx, recoveryToken)

	// TODO: Should place this link to config
	http.Redirect(w, r, "https://storjlabs.atlassian.net/servicedesk/customer/portals", http.StatusSeeOther)
}

// graphqlHandler is graphql endpoint http handler function.
func (server *Server) graphqlHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)

	handleError := func(code int, err error) {
		w.WriteHeader(code)

		var jsonError struct {
			Error string `json:"error"`
		}

		jsonError.Error = err.Error()

		if err := json.NewEncoder(w).Encode(jsonError); err != nil {
			server.log.Error("error graphql error", zap.Error(err))
		}
	}

	w.Header().Set(contentType, applicationJSON)

	query, err := getQuery(w, r)
	if err != nil {
		handleError(http.StatusBadRequest, err)
		return
	}

	rootObject := make(map[string]interface{})

	rootObject["origin"] = server.config.ExternalAddress
	rootObject[consoleql.ActivationPath] = "activation/?token="
	rootObject[consoleql.PasswordRecoveryPath] = "password-recovery/?token="
	rootObject[consoleql.CancelPasswordRecoveryPath] = "cancel-password-recovery/?token="
	rootObject[consoleql.SignInPath] = "login"
	rootObject[consoleql.LetUsKnowURL] = server.config.LetUsKnowURL
	rootObject[consoleql.ContactInfoURL] = server.config.ContactInfoURL
	rootObject[consoleql.TermsAndConditionsURL] = server.config.TermsAndConditionsURL

	result := graphql.Do(graphql.Params{
		Schema:         server.schema,
		Context:        ctx,
		RequestString:  query.Query,
		VariableValues: query.Variables,
		OperationName:  query.OperationName,
		RootObject:     rootObject,
	})

	getGqlError := func(err gqlerrors.FormattedError) error {
		var gerr *gqlerrors.Error
		if errors.As(err.OriginalError(), &gerr) {
			return gerr.OriginalError
		}
		return nil
	}

	parseConsoleError := func(err error) (int, error) {
		switch {
		case console.ErrUnauthorized.Has(err):
			return http.StatusUnauthorized, err
		case console.Error.Has(err):
			return http.StatusInternalServerError, err
		}

		return 0, nil
	}

	handleErrors := func(code int, errors gqlerrors.FormattedErrors) {
		w.WriteHeader(code)

		var jsonError struct {
			Errors []string `json:"errors"`
		}

		for _, err := range errors {
			jsonError.Errors = append(jsonError.Errors, err.Message)
		}

		if err := json.NewEncoder(w).Encode(jsonError); err != nil {
			server.log.Error("error graphql error", zap.Error(err))
		}
	}

	handleGraphqlErrors := func() {
		for _, err := range result.Errors {
			gqlErr := getGqlError(err)
			if gqlErr == nil {
				continue
			}

			code, err := parseConsoleError(gqlErr)
			if err != nil {
				handleError(code, err)
				return
			}
		}

		handleErrors(http.StatusOK, result.Errors)
	}

	if result.HasErrors() {
		handleGraphqlErrors()
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		server.log.Error("error encoding grapql result", zap.Error(err))
		return
	}

	server.log.Debug(fmt.Sprintf("%s", result))
}

// serveError serves error static pages.
func (server *Server) serveError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	switch status {
	case http.StatusInternalServerError:
		templates, err := server.loadTemplates()
		if err != nil {
			server.log.Error("unable to load templates", zap.Error(err))
			return
		}
		err = templates.internalServerError.Execute(w, nil)
		if err != nil {
			server.log.Error("cannot parse internalServerError template", zap.Error(err))
		}
	case http.StatusNotFound:
		templates, err := server.loadTemplates()
		if err != nil {
			server.log.Error("unable to load templates", zap.Error(err))
			return
		}
		err = templates.notFound.Execute(w, nil)
		if err != nil {
			server.log.Error("cannot parse pageNotFound template", zap.Error(err))
		}
	}
}

// seoHandler used to communicate with web crawlers and other web robots.
func (server *Server) seoHandler(w http.ResponseWriter, req *http.Request) {
	header := w.Header()

	header.Set(contentType, mime.TypeByExtension(".txt"))
	header.Set("X-Content-Type-Options", "nosniff")

	_, err := w.Write([]byte(server.config.SEO))
	if err != nil {
		server.log.Error(err.Error())
	}
}

// brotliMiddleware is used to compress static content using brotli to minify resources if browser support such decoding.
func (server *Server) brotliMiddleware(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		isBrotliSupported := strings.Contains(r.Header.Get("Accept-Encoding"), "br")
		if !isBrotliSupported {
			fn.ServeHTTP(w, r)
			return
		}

		info, err := os.Stat(server.config.StaticDir + strings.TrimPrefix(r.URL.Path, "/static") + ".br")
		if err != nil {
			fn.ServeHTTP(w, r)
			return
		}

		extension := filepath.Ext(info.Name()[:len(info.Name())-3])
		w.Header().Set(contentType, mime.TypeByExtension(extension))
		w.Header().Set("Content-Encoding", "br")

		newRequest := new(http.Request)
		*newRequest = *r
		newRequest.URL = new(url.URL)
		*newRequest.URL = *r.URL
		newRequest.URL.Path += ".br"

		fn.ServeHTTP(w, newRequest)
	})
}

// loadTemplates is used to initialize all templates.
func (server *Server) loadTemplates() (_ *templates, err error) {
	if server.config.Watch {
		return server.parseTemplates()
	}

	if server.templatesCache != nil {
		return server.templatesCache, nil
	}

	templates, err := server.parseTemplates()
	if err != nil {
		return nil, Error.Wrap(err)
	}

	server.templatesCache = templates
	return server.templatesCache, nil
}

func (server *Server) parseTemplates() (_ *templates, err error) {
	var t templates

	t.index, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "dist", "index.html"))
	if err != nil {
		server.log.Error("dist folder is not generated. use 'npm run build' command", zap.Error(err))
		// Loading index is optional.
	}

	t.usageReport, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "static", "reports", "usageReport.html"))
	if err != nil {
		return &t, Error.Wrap(err)
	}

	t.notFound, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "static", "errors", "404.html"))
	if err != nil {
		return &t, Error.Wrap(err)
	}

	t.internalServerError, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "static", "errors", "500.html"))
	if err != nil {
		return &t, Error.Wrap(err)
	}

	return &t, nil
}

// NewUserIDRateLimiter constructs a RateLimiter that limits based on user ID.
func NewUserIDRateLimiter(config web.RateLimiterConfig) *web.RateLimiter {
	return web.NewRateLimiter(config, func(r *http.Request) (string, error) {
		user, err := console.GetUser(r.Context())
		if err != nil {
			return "", err
		}
		return user.ID.String(), nil
	})
}

// responseWriterStatusCode is a wrapper of an http.ResponseWriter to track the
// response status code for having access to it after calling
// http.ResponseWriter.WriteHeader.
type responseWriterStatusCode struct {
	http.ResponseWriter
	code int
}

func (w *responseWriterStatusCode) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

// newTraceRequestMiddleware returns middleware for tracing each request to a
// registered endpoint through Monkit.
//
// It also log in INFO level each request.
func newTraceRequestMiddleware(log *zap.Logger, root *mux.Router) mux.MiddlewareFunc {
	log = log.Named("trace-request-middleware")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			begin := time.Now()
			ctx := r.Context()
			respWCode := responseWriterStatusCode{ResponseWriter: w, code: 0}
			defer func() {
				// Preallocate the maximum fields that we are going to use for avoiding
				// reallocations
				fields := make([]zapcore.Field, 0, 6)
				fields = append(fields,
					zap.String("method", r.Method),
					zap.String("URI", r.RequestURI),
					zap.String("IP", getClientIP(r)),
					zap.Int("response-code", respWCode.code),
					zap.Duration("elapse", time.Since(begin)),
				)

				span := monkit.SpanFromCtx(ctx)
				if span != nil {
					fields = append(fields, zap.Int64("trace-id", span.Trace().Id()))
				}

				log.Info("client HTTP request", fields...)
			}()

			match := mux.RouteMatch{}
			root.Match(r, &match)

			pathTpl, err := match.Route.GetPathTemplate()
			if err != nil {
				log.Warn("error when getting the route template path",
					zap.Error(err), zap.String("request-uri", r.RequestURI),
				)
				next.ServeHTTP(&respWCode, r)
				return
			}

			// Limit the values accepted as an HTTP method for avoiding to create an
			// unbounded amount of metrics.
			boundMethod := r.Method
			switch r.Method {
			case http.MethodDelete:
			case http.MethodGet:
			case http.MethodHead:
			case http.MethodOptions:
			case http.MethodPatch:
			case http.MethodPost:
			case http.MethodPut:
			default:
				boundMethod = "INVALID"
			}

			stop := mon.TaskNamed("visit_task", monkit.NewSeriesTag("path", pathTpl), monkit.NewSeriesTag("method", boundMethod))(&ctx)
			r = r.WithContext(ctx)

			defer func() {
				var err error
				if respWCode.code >= http.StatusBadRequest {
					err = fmt.Errorf("%d", respWCode.code)
				}

				stop(&err)
				// Count the status codes returned by each endpoint.
				mon.Event("visit_event_by_code",
					monkit.NewSeriesTag("path", pathTpl),
					monkit.NewSeriesTag("method", boundMethod),
					monkit.NewSeriesTag("code", strconv.Itoa(respWCode.code)),
				)
			}()

			// Count the requests to each endpoint.
			mon.Event("visit_event", monkit.NewSeriesTag("path", pathTpl), monkit.NewSeriesTag("method", boundMethod))

			next.ServeHTTP(&respWCode, r)
		})
	}
}
