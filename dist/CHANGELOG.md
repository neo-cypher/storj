## Changelog
* 756df19a9  On branch main  Your branch is up to date with 'origin/main'.
* 2a4e99473 chore(deps): pin dependencies
* e7b8124d4 chore(deps): add renovate.json
* 7541f8465 Merge branch 'storj:main' into main
* b3e1be37f satellite/projectaccounting: query to get daily project usage by date range
* 4a26f0c4f cmd/storagenode: restore passing arguments
* 7f1dc74e3 cmd/storagenode: Change order load id in setup
* 11a35520c mod: bump common, quic, golang/x packages
* 3c8e41e66 web/satellite: get object map and preview by signed request.
* 778e7e100 web/satellite: logout after password reset added
* 3cd8e46e7 web/satellite: new project dashboard: buckets sections added
* 4d0dd737b satellite/admin: add endpoints for oauth clients
* 2e31ef3f2 satellite/metabase: better error message while move
* bc161794f satellite/metabase: drop DeleteObjectLatestVersion method
* 7e63afbef storagenode: docker image autoupdate binaries
* 5041ee0ab satellite/metabase: drop GetObjectLatestVersion method
* dffa7845f satellite/metainfo: split endpoint tests
* 02bbcd7fb satellite-ui: update wizard to instruct uplinkng usage
* 3a7fb4588 satellite/satellitedb: users/projects segment limit migration
* 2ae78db66 satellite/email: add delimiter to close the last part of Multipart emails
* 0d03473e0 cmd/partnerid-to-useragent-migration: LIMIT 1 in MigrateUsers id select query
* 1743a72f2 cmd/partnerid-to-useragent-migration: parallelize tests, use db test migration method
* b7ca2289e cmd/uplinkng: improvements to prepare for doc update
* 9d3614fc1 cmd/partnerid-to-useragent-migration: add value_attributions migration and tests
* 64c8de6ea mod: use vendored base58
* d7d941207 Update README.md
* 8204f244b Create main.yml
* af4594dad Merge branch 'storj:main' into main
* 6b6e9901e cmd/partnerid-to-useragent-migration: add bucket_metainfos migration and tests
* 1f1a054de cmd/partnerid-to-useragent-migration: add api_keys migration and tests
* 6d6d6776d cmd/partnerid-to-useragent-migration: add projects migration and tests
* b65eed53f Merge branch 'storj:main' into main
* d393f6094 satellite/accounting: adjust tests to expose new uplink errors
* a1aa9f24e private/testplanet: expose NewLogger for other tests
* b8c55fdd8 satellite/projectaccounting, satellite/console, web/satellite: implemented backend for bandwidth chart
* 5d6ee506b private/apigen: initial typings and methods
* 67c58fdd5 cmd/uplinkng: add setup command
* 17f56ab63 satellite/metainfo: split endpoint into smaller files (object/segment)
* 33b9d24c4 web/satellite: NotificationItem: support link added
* 9ce1a4e25 satellite/accounting: use AOST to get segment usage
* 345116258 web/console: Save timestamp when creating users to mark their first email verification notification.
* e6961e1fd cmd: add command for partner_id to user_agent satelliteDB migration
* c968bd0b7 docs/blueprints: Create API renovation blueprint
* 65c38c6ce satellite/metainfo: split metainfo endpoint into smaller files (buckets)
* 913718ef9 satellite/accounting: test bandwidth limit without cache
* be02aa9b1 satellite/overlay: add test for UpdateCheckIn panic
* 4be84c3a6 satellite/accounting: fix how we put segments usage into cache
* bfad42f29 satellite/metainfo: split metainfo endpoint into smaller files
* 53e96eb03 Merge branch 'storj:main' into main
* e7e39fe62 satellite/metainfo: reduce default user rate limit to 100
* 1c47163ee satellite/metainfo: add more logging for CreateBucket
* c30c7def4 satellite/{console,payments},web/satellite: save signup promo code to newly registered users
* 251e432e5 satellite/satellitedb: drop contained from nodes and reputations
* 464b4b02f satellite/console/consoleweb: Only replace some kinds of characters in name
* 2de6feefd satellite/contact: fix connection leak in PingMe endpoint
* 8b40a071a storagenode: check if QUIC is properly configured
* 51be14cea go-build.yml modifications
* 9163fc2c4 Typo in go build command: RACE was changed to race
* 1445a76e1 workflow
* 59648dc27 satellite/console: Filter special characters out of name
* 9f3c1f9cd satellite/contact: add PingMe endpoint
* 95750b9d3 satellite/console: verify limits while test
* 01287913e web/satellite: graphql error logging added
* 228f465d4 satellite/metainfo: сombine checks for storage and segment limit
* 270c6e24f satellite/console: fix default limit for free-tier segment usage limit
* a81af2829 web/satellite: error while adding credit card handled
* 309ac41bf satellite/satellitedb/dbx: add schema for oauth tokens
* 83f7d3e04 satellite/satellitedb/dbx: add schema for oauth clients
* 36c07720b satellite/metainfo: bound user agent metrics
* 1007e6300 web/satellite: Distinguish between file and folder upload buttons
* cf462fcf9 satellite/admin: update geofence endpoints to follow response conventions
* 6bbe2fc69 cmd/uplinkng: fix flag/error messages
* 4f96a8564 cmd/uplink/share: register access via DRPC
* 5e9643e1b satellite/metainfo: cleanup validation
* 18ab464b5 satellite/metainfo: UpdateProjectSegmentUsage enabled only with segment validation
* 3a987ac00 satellite/admin: Fix client API & docs
* 8e80a51b6 Makefile: Set NodeJS version to use in var
* 44a73b01d cmd/storagenode: use contact address instead of server address
* eb0d08d59 satellite/metainfo: collect versions of user tools
* e0b3c7152 satellite/nodestats: use overlay node for reputation status
* b6625cade satellite/metainfo: uncomment tests, renamed EncryptedPath into EncryptedObjectKey
* bc3392cda cmd/uplinkng: fix clingy argument usage
* 332b673a0 satellite/metainfo: comment tests for rename
* 15c2b2cc1 satellite/metainfo: comment tests for rename
* dc6128e9e satellite/admin: add update project segment limit on ui
* 8be655c48 satellite/metainfo: add segment limit validation, caching
* f0ada87a9 web/satellite: fix FileModal Share Link (#4341)
* 35113634f Makefile: do not remove files to keep state clean
* 1f8f7ebf0 satellite/{audit, reputation}: fix potential nodes reputation status inconsistency
* b21cbc85f satellite/orders: log level warn for remaining "bucketName or projectID not set" (#4326)
* 606a73a8e docs/testplan: Adding a testplan for Project Dashboard (#4318)
* 171f908d8 cmd/satellite: rename monkit stat with '_', not '-'
* ab37b65cf satellite/{accounting,orders,satellitedb}: group bucket bandwidth rollups by time window
* 5d3085f38 testuite/ui/satellite: cleanup and skip tests
* d06ba56bc web/satellite: fix invalid bucket name message
* a42b9d1a4 all: fix uses of email.com
* 04d04c030 web/satellite: remove old object browser
* 0c3a7b457 Update login_unverified_nonexistent_account_test.go
* 35712f4de Update login_unverified_nonexistent_account_test.go
* ba1bff536 Update login_reset_password_verified_account_test.go
* edf17754e Update login_reset_password_verified_account_test.go
* 14c64318a Update login_unverified_nonexistent_account_test.go
* bc7a9593a Update login_reset_password_unverified_account_test.go
* bde3a884c testsuite/ui/satellite: email verification UI tests
* 12cb733b1 satellite/satellitedb: add last_verifcation_reminder column to users table
* 6a059b7f9 web/satellite: remove all references to the old onboarding flow
* 1dce461f5 web/satellite: makes sure files with uppercase extensions are able to be previewed in the object browser
* c884f327c web/satellite: add new events to segment
* 5351c4375 satellite/analytics: track origin header and referer on signup
* 2a4ff6f46 satellite/satellitedb: select user_agent in apikeys method GetPagedByProjectID
* c58a81e7c testsuite/ui: disable flaky UI tests
* 1bd74658d satellite/console: Disable segment limit checking during project update
* f6a33ec11 Fix social media sharing links (#4321)
* 589c82f23 storagenode: display quic status on SNO dashboard
* ff653f6f2 satellite/metainfo: stop using in tests unused method
* 442647017 cmd/uplinkng: rename parallelism flags, small fixes
* da9e0188f Update README.md (#4320)
* 033446530 Makefile: fix variable used to trigger release drafting (#4319)
* 5f39f2076 web/: add check for change to eslint import
* e2481fb51 scripts,Jenkinsfile,Makefile: upload binaries to drafted github release
* db7985ca9 satellite/metainfo: set EncryptedMetadataEncryptedKey while listing objects
* 87a929212 satellite/metainfo/zombiedeletion: enable chore by default
* a6139f5a6 satellite/orders: log less decrypt order messages
* a04b9ce37 testsuite/ui/multinode: adding more UI tests (#4301)
* c10cc302b Makefile: Fix target for building Admin UI
* 185617326 satellite/email: Move MIME-Version header to the last element before Content-Type
* b7b010adc satellite/admin: add project limit on segment count
* f87ce0cad satellite/admin/ui: Fix component rendering
* 5573ece84 satellite/admin/ui: Migrate to SvelteKit
* db0bd38d9 web/satellite: Add new user events to segment for link clicks
* 60d03b9e4 satellite/metainfo: drop GetSegmentByLocation
* 3f0babffe satellite/{console,satellitedb}: add project segment limit to user
* 9d13c649a satellite/{console,satellitedb}: Forbid creating users with used email
* 137641f09 multinode/nodes: pass Node entity to Add method instead of set of a parameters
* 596124f0f docs: Adding a testplan folder with a README and TEMPLATE for testplan (#4306)
* ed9bde8f6 cmd/storagenode: add info command
* 60c828056 cmd/multinode: add add command to multinode
* 7fba79be3 cmd/uplinkng: add register command
* 75acd6109 cmd/{uplink,uplinkng}: set user agent for uplink binaries
* f83a801ef web/satellite: Fix subtitle in disable 2FA (#4285)
* 911540e76 satellite/orders: avoid logging "bucketName or projectID not set"
* 9615bd191 cmd/uplinkng: update clingy for better boolean support
* e6aa52cd6 cmd/uplinkng: reorganize iterator helpers
* 4762493e0 storagenode/reputation: fix missing space in suspension message (#4309)
* 9a3739934 satellite/metainfo: remove Gateway-ST library name from UserAgent string
* c8c4e8835 cmd/uplinkng: file level parallel copy
* b6f1efbbd web/satellite: fixing linting errors in the FileShareModal
* cd0d61f52 web/satellite: fixing linting errors in the FileEntry component
* e67b905b6 web/satellite: fixing linting errors in the FileBrowser component
* 2be56e335 web/satellite: fixing linting errors in the file browser
* 5c84243d9 Jenkinsfile: downgrade containers for version tests
* 679935474 ci: bump to go 1.17.5
* baaa96c20 cmd/uplinkng: introduce MultiWriteHandle
* 34890c919 cmd/uplinkng: introduce MultiReadHandle
* c33d2f58c satellite/metainfo: override StreamMeta only if needed
* a41758bba Track user events through HubSpot events API (#4300)
* 27c6c6aea ci: Golangci lint v1.43.0 changes (#4307)
* 000957f64 satellite/accounting: add functionality to get and cache project segment count
* af3edf2b7 scripts/release.sh: making clear if builing a production or development version (#4304)
* d94d8d177 satellite/accounting: expose project segment limit
* ace90fd84 mod: bump clingy version
* 3b51eea31 satellite/metainfo: store empty useragent in bucket attribution instead of throwing error
* 5013dd8a4 satellite/metainfo: remove error message prefix
* 15a7e7f35 web/satellite: fix delete bucket issue
* b5194762f cmd/uplinkng: share command added
* 73730b23e satellite/satellitedb: add segment_limit colum to projects table
* b3cea3d1b satellite/audit: account for piece size during audit reservoir sampling
* 82fb0fce0 web/satellite: added charts date range picker for new project dashboard
* be10ce84f web/satellite: implemented charts UI for new project dashboard
* b78f65e83 satellite/console, web/satellite: added object and segment count per project info to new project dashboard
* 16c8a0801 web/satellite: new project dashboard routing updates
* bde8fb0a5 satellite/admin: Move geofence endpoint to bucket info
* bae8276f3 satellite/admin/ui: Add missing geofencing endpoints
* eec5ad4f5 cmd/uplinkng: access inspect command added
* 984792fd1 satellite/satellitedb: Add GetByEmailWithUnverified to users table
* 9c1129b4c Revert "satellite/satellitedb: migrate partner_id db column to user_agent db column"
* 4daaa9503 satellite/console: enable object flow by default (#4298)
* 99deec47b satellite/metainfo: consistent error message for bucket not found
* 2c0a360a1 satellite/satellitedb: migrate partner_id db column to user_agent db column
* de8464e84 satellite/metabase: add GetProjectSegmentCount method
* 1d14e4ef6 satellite/metainfo: Check ctx in download Object & Segment
* 7d8448238 satellite/metainfo: prevent moving objects between geofenced buckets
* 76c2228fb satellite/metainfo: propagate geofencing between buckets and stream id
* 1fd14e291 testsuite/ui/uitest: fixes for uitest
* 788d775f6 web/satellite: fix browser progress-bar visibility
* 88904301b web/satellite/static: fix 404 and 500 page logo link
* 8f07abafe cmd/uplinkng/ultest: make possible to test access commands
* 24cf7e8ea cmd/uplinkng: add mv command
* 07fad7591 storagenode/piecestore: upload and download metrics for Grafana alerts (#4280)
* f9b630b0f Fix monthly earning estimation (#4282)
* 336500c04 satellite/repair: only record audit result if segment can be downloaded
* 3de7f8d5a satellite/console, web/satellite: feature flag for new project dashboard
* b7a5b1f3a web/satellites: fixes for new objects flow
* 89639199f satellite/repair/repairer: remove unused healthyMap
* 2e363707c satellite/buckets: tests for bucket placement config
* f773bb80c mod: bump common to fetch latest placement type changes
* a038fc1dc satellite/admin: add endpoints for managing a buckets geofence
* 814e3126f satellite/buckets: add new buckets service
* 09568b3e2 satellite/console: change default feature flag (#4274)
* dfd2977a0 testsuite/ui: fixes to setup
* debb82f59 Geofencing and advanced placement-constraint support (#4227)
* cb499cb7d Makefile: Fix target builds the satellite admin UI
* 07e07d3a2 satellite/metainfo: Add project ID to error logs
* ab425d424 build: Consider the Admin UI assets
* 289ae3325 satellite/metabase: always include encryption in listed object
* 5ed3846e1 testsuite/ui/satellite: add browser tests
* 6dffbd417 satellite/metainfo: Log in warn level limits exceeds
* 44b1ca6b9 satellite/{console,satellitedb}: move project limits from config file to DB to keep limits on a per user basis
* b41e140b2 web/satellite: fixed dropdowns weird behavior
* 96058be1a storagenode/storagenodedb: fix error message
* 58916d8ba ci: don't trigger UI tests when go.mod changes
* 4e67ea007 satellite/admin: Serve static UI assets
* af5b90ed3 storagenode/storagenodedb: include dbname in error
* bf51c286d satellite/geoip: update node check-in to associate a country code
* dda6720dd web/satellite: added 2 new steps for onb CLI flow
* 2ebdc1303 satellite/satellitedb: add disqualification reason to nodes table
* 12d4dc16a satellite/{metainfo,metabase}: make metadata optional for CommitObject
* fe3856ff9 web/satellite: remove ramda dependency
* 1d5b3e508 satellite/admin/ui: fix rollup building
* 8eebbf3d7 satellite/audit: fix TestReverify timeouts
* d5628740f satellite/overlay/straynodes: prevent disqualification in tests
* 583cdc6e0 satellite/admin/ui: Create Web UI v1
* 3122de456 web/satellite/browser: fix formatting
* 9bdcc415b satellite/nodeselection: add geofencing constraints to the node selection criteria
* 696b8f0d8 web/satellite: moved object's passphrase step to be after buckets
* 3c683998f satellite/console, web/satellite: feature flag for new objects flow
* 91158fd3b satellite: omit system metadata fields if not requested
* fb604be46 satellite/gracefulexit: use nodecache with gracefulexit
* 98b59fe30 web/satellite: reworked Encrypt your data component
* 4a530ccff cmd/storagenode: simplify windows service loop
* 609526f6a web/satellite: reworked new navigation sidebar
* 804304a75 web/satellite: integrate browser
* b314559a1 satellite/metainfo: TestBatch params updated
* e72108135 cmd/storagenode: fix service initialization code
* 1dd537953 satellite/metainfo: strip the uplink version from the UserAgent and bound its size
* 27091826b satellite/{satellitedb|metabase}: add SQL fields for geofencing
* 16a334020 cmd/uplinkng: add ranged download
* 6b7473e8b web/satellite/images: reduce svg size
* d043b9fac satellite/console: load index.html from disk when developing
* 6a69c6647 cmd/storj-sim: support printing node url
* e792727be satellite/metabase: allow setting metadata with begin object requests
* bd2448bc4 cmd/uplinkng: expand version information
* 1de8a695e satellite/{overlay,satellitedb}: fix stray nodes DQ bug
* 774ae017e satellite/{satellitedb, web}: display object count in satellite UI
* 187e50836 web/satellite/browser: auto-format code
* ea8782fcf web/satellite: move browser from storj/browser
* 460f8fcdd web/satellite: added satellite name to MFA label
* 4a955b8a7 web/satellite: account area stylings and status fixes
* 785eb93cd ci: install gateway@latest instead of @main
* 431f55d53 ci: cross-compile storagenode-updater
* 136899eee web/satellite: Trim whitespace around login credentials
* 4d023f773 cmd/storagenode-updater: update to correct version
* f94d8a200 cmd/uplinkng: mkdir config file folder
* 8c1a149de build: bump private dependency
* 9749f9756 cmd/uplinkng: port expanded option to show additional metadata (#4229)
* edb8d656d satellite/metainfo: adjust piecedeletion timeouts
* d3a0364f2 Makefile: add uplinkng to build binaries
* fcde92d38 web/satellite: display segment count in billing
* d441c8da1 satellite: use segment count for billing
* 52f8c8175 satellite/satellitedb: add segments column into invoiceprojectrecords table and drop not null constraint on objects column
* 1951450c5 cmd/satellite: add register-lost-segments command
* 5bbb69886 web/satellite: improved responsiveness of new nav sidebar
* 1409288c4 mod: bump semantic version to go 1.17
* 20d03bebd satellite/nodeselection: flexible interface to includes nodes in selection
* de38e2e7d satellite/metainfo: stop using sat StreamID Redundancy
* c7e257012 web/satellite: tab adaptation for new navigation sidebar
* eb68dbad4 satellite/satellitedb: fix ordering in listPendingTransitionShim
* d90a1b09a satellite/console,satellitedb: add signup promo code column to users
* 01cfbde56 satellite/console: increase free tier project bandwidth and storage from 50 GB to 150 GB and reduce free tier max projects from 3 to 1
* 3c7e28939 satellite/matainfo: handle metabase errors in metainfo endpoint in common way
* 65d139fb9 private/testplanet: create metabaseDB after reconfigure apply
* 6c2cb98d2 web/satellite: fix onb CLI setup uplink download link
* 85b49bb27 cmd/uplink: add ranged download to uplink cli
* a281adddd web/satellite: auto-create first demo bucket for the user
* 0ed3ef0fe web/satellite: registation success page moved
* df44e8152 satellite/admin: Move declaration before first usage
* 3e040767b testsuite/ui/satellite: fix tests to pass with new navigation
* df6cf60ff go.mod: update to minimum supported go version (#4239)
* f654c3ae4 satellite/gracefulexit: drop unused column in graceful_exit_progress
* 1b2252500 satellite/payments: Fix discount when listing invoices.
* 56f35e9ea mod: bump private repo
* 50baefa10 satellite/metabase: limit maximum number of parts and size
* 49492d28f testsuite/ui/satellite: test for new navigation sidebar
* f2d8e97d9 satellite/satellitedb: simplify select nodes query construction
* 398910703 web/satellite: add feature flag for new browser
* f82b69685 mod: bump common dependency
* b32fbc0f9 satellite/metainfo: speedup deletion tests
* 187941ff8 satellite/{metabase,metainfo}: add some missing monitoring
* bdadc4b46 private/testplanet: add stack traces related to the planet
* 9ed826eab cmd/uplink/mv: moving files between folders
* 987cb6ab1 cmd/uplinkng: allow removing pending objects
* 2df41028a storagenode/piecestore: add logs for restore trash endpoint
* 9773197c6 web/satellite: new nav account area and vertical layout fixes
* 0aab339d0 web/satellite: links and dropdowns for new navigation sidebar
* b619ce691 web/satellite: new nav project dropdown
* 000944777 satellite/console, web/satellite: feature flag for new navigation structure
* 9c8b6c3fc testsuite/ui: Add coupon tests
* 8d18033a9 satellite/rewards: adding SeaweedFS to partners list (#4230)
* d250f9dfe satellite/core: fix label for orders chore (#4228)
* 51f488fc4 satellite/satellitedb: fix selection query with AOST
* 4a146000c satellite/metainfo: read from DB only needed columns fro bucket
* 8293c9659 satellite/metabase: speedup iterator tests
* f5f239843 satellite/metabase: remove code related to part deletion
* 4ad7056bf satellite/payments: Add old invoice list functionality
* 35e4a87e6 satellite/repair: ignore expired segments at the beginning of the repair work
* 29599dd7c testsuite/ui/multinode: add first node test
* 4bbf667ad satellite/{satellitedb,attribution,console}: value attribution changes that add userAgent field to buckets table and all tables that have partner_id
* 8a0a23353 web/satellite: Fix typo in 2FA popup (#4221)
* 016ac7505 satellite/console: Enable new onboarding workflow (#4222)
* d7812a3b1 satellite/console,web/satellite: add slider and unit toggling to project limit updating
* 52c950e42 testsuite/ui/uitest: add Edge testing
* 96dc43fd4 satellite/metainfo: disable DeletePart method
* 56fe63612 satellite/{reputation/satellitedb}: remove references to contained column in reputations table
* bb21551a9 satellite/satellitedb: remove references to contained column in nodes table
* ac543714e add config file for git review usage
* 331132be9 web/satellite: Fix support link on registration page
* 8f4550d62 ci: fix using npm run build-wasm
* c54bcd63a Makefile: bump node to v16
* 4535da96e web/multinode: removed conversion for Expectations (#4216)
* 91ce70e4e satellite/metabase: allow overwriting segments in pending objects
* 5b729779a satellite/console: Automatically log a user in after verifying email
* b8dd35cea private/testplanet: extend testplanet with multinode instance
* b8e8110ca mod: bump common and uplink dependency
* aa858a3a2 testsuite/ui/uitest: separate func for browser
* 1fdb0eaa5 Revert "satellite/metabase: use storj.Nonce instead []byte"
* bc6d8c06e web: optimize builds
* 4020e9e2c private/testmonkit: move monkit test helper
* 5b4a9070b scripts/testversions: reduced the number of version to test
* 3b751a35c satellite/{payments,satellitedb}: Remove custom coupon implementation
* 1ef06fae9 satellite/metabase: use storj.Nonce instead []byte
* b7405db2a satellite/satellitedb: optimize migration test
* f77f61532 satellite/orders: use egress_dead for calculating allocated bandwidth
* 9fd091831 ci: optimize benchmarks
* f837551d7 storj: bump uplink to v1.7.0
* 1d8342ab1 web/satellite: fixed access grant web worker caching issue
* 05960b2cf satellite/admin: Response 404 when entity not found
* fb0d055a4 satellites/orders: populate egress_dead in project_bandwidth_daily_rollups
* 6ab2647ad ci: run UI tests
* 8a21a8cbc satellite/metainfo: don't enable connection pool if was setup earlier
* 535592831 satellite/payments/stripecoinpayments: fix nil panic during storj token deposit
* e5bb89736 satellite/internalpb: for StreamID rename EncryptedPath to EncryptedObjectKey
* f93dc5a16 satellite/metainfo/piecedeletion: enable connection pool
* bb55c3059 web: move eslint-storj to separate repo
* aa80d3705 satellite/metabase: add automatic conversion to DB value for ObjectKey fields
* 4b79f5ea8 satellite/repair: test if audit scores increases during repair
* 38ce0e154 satellite/gracefulexit: drop table graceful_exit_transfer_queue
* f2f8a9ca1 satellite/metabase: send metric about zombie segment deletion
* 5f444c1fe satellite/payments: fix currency rates acquisition
* 10a468fb5 web/satellite: make all header components visible on smaller screen sizes
* 854b66f1b testsuite/ui/satellite: refactored go-rod tests selectors a little bit
* bef4f5ec8 satellite/admin: Validate new user data earlier
* fed09316b satellite/admin: Use helper for sending JSON
* bb575ef73 satellite/admin: Send JSON content-type for errors
* 821a077f7 satellite/{admin,console}: Move tests
* 7ebd35337 private/testplanet: document env variables
* 5b40922ac satellite/metainfo: avoid endless loop when startup fails
* db2cb33db ci: use correct Go version for the image
* 6511bb91f private/testplanet: support writing monitoring spans
* 0209bc6ef cmd/uplink: add mv command
* 3dbd44347 web/satellite: don't ignore Vue errors and warnings
* 5b6613631 cmd/uplinkng: fix mb command
* ead310d31 satellite/satellitedb: avoid running migrate tests concurrently
* d5043a0f8 web/storagenode: dcs word removed from gui
* f829d64a3 web/satellite, testsuite/ui/satellite: updated onb Welcome screen's encryption info and added tests
* 46d1b4dfa testsuite/ui/satellite: added test for restart onb tour functionality
* 8e3d7d30e web/satellite, testsuite/ui/satellite: added tests for onb CLI os tabs switching
* 1def7b0ec web/satellite, testsuite/ui/satellite: added tests for invalid sign up credentials and satellites dropdown
* f52f5931d ci: don't fail when nothing to delete
* 0d58172c3 storagenode: add doc.go files for sno packages
* c053bdbd7 satellite/satellitedb: prepare to remove big.Float from db
* a16aecfa9 satellite/payments: specialized type for monetary amounts
* c911360eb satellite/metainfo: separate burst limit from rate limit config
* ab7e81c27 satellite/accounting: update GetProjectBandwidthTotals function to use the less expensive project_bandwidth_rollups table
* 6d3fd33ca satellite/metabase/segmentloop: start immediately on manual trigger
* 4db80773e satellite/satellitedb: add burst_limit for project
* cc9b845a8 ci: switch back to pulling gateway@main
* 9c232cebe ci: make deletion a separate step
* 512d0d61f satellite/metrics: speedup test
* 40084d51a satellite/metainfo: reduce number of testplanet runs
* d397f6e2b docs: move contribution guide to repository root
* 1ed5db146 satellite/metainfo: simplifying limits code
* 118e64fa1 jenkinsfile: bump build timeout
* 803664442 ci: try fix builds
* 7fd346692 ci: remove deleting workspace steps
* 0330871a3 satellite/metabase: add missing FinishMoveObject parameter
* 4fefa36a7 satellite/metabase: NewBucket added to metabase/metainfo FinishMoveObject methods
* 9da3de1a8 web/satellite, testsuite/satellite/ui: removed onb CLI flow's Generate AG step and updated tests
* 7d90770f8 .github: add issue templates
* 469ae72c1 satellite/repair: update audit records during repair
* 030ab669f web/satellite: Route all child routes to object browser
* 0e1c3cb81 docs: add contributing guide for SNO development
* b97479e36 web/satellite: Reset reCAPTCHA response token on error
* 8ef03b096 ci: cleanup ws before build, fix gateway install
* bd4017b2a ci: fix logical races between tools
* 99914dfc0 web/satellite: fix for button label uppercase issue
* cae08d816 satellite/metabase: FinishMoveObject segment query improved
* 71eb184ef storagenode/piecestore: simplify TestTooManyRequests
* 36911b44a satellite/accounting/tally: make tests faster
* df09e7d1f satellite/metainfo: ensure storagenodes finish work for test
* 8b91c55ec web/satellite, satellite/console: return old onboarding flow with feature flag
* 38366914b satellite/metainfo: add position to begin move results
* 32cee1e57 satellite/metabase/segmentloop: ensure we shutdown on ctx err
* 69d3a839f web/satellite: fix for onb CLI upload step image
* 91c69a8af web/satellite: fix for encrypt your data screen
* 5e4b196b1 satellite/metainfo: finish move object
* 9153b221f testsuite/ui/satellite: updated pnboarding CLI tests and improved selectors
* 853a7536f web/satellite: reworked onb CLI flow to correspond v1.1
* 0b500a30e satellite/audit: move audit metrics out of reporter
* 252b78580 satellite/console: add status check to user authorization to ensure deleted accounts cannot perform actions
* 09e1ff7fc web/satellite: fix promise usages
* 6e660cecd Jenkinsfile: test cross-compile and bump deps
* b2d724962 cmd/storagenode-updater: avoid depending on the storagenode code
* 6b3718f33 Makefile: skip darwin build for the time being
* f0b73b4f4 mod: bump drpc, uplink and common
* b160ec4c1 satellite/orders: bound RollupsWriteCache flushes
* 86845351c web/satellite: reworked upgrade to Pro account modal
* 3c3df0f36 storage/filestore: test interaction between VerifyDirReadable and VerifyDirWriteable
* 9dab10e1d web/satellite: Remove Advanced Report from GUI
* 2e6b4e90f web/satellite: add warning message to EncryptYourData component
* e805dce70 satellite/satellitedb: drop audit reputation score related columns from nodes table
* 07311cdd6 satellite/metainfo: begin move object added
* c00ecae75 satellite/gracefulexit: stop using gracefulexit_transfer_queue
* e1a4195a3 satellite/metainfo/endpoint: MaxObjectKey length validation for BeginObject and CommitObject added
* 1aec831d9 satellite/audit,storage: increase sleep delay in TestMaxVerifyCount
* b472925c1 web/satellite: hot fix for Onb CLI flow's Create Bucket component's back button
* 7f595445a satellite/metainfo: make subsequent auth validations not perform rate-limiting
* c06424cbf web/satellite: fixed inputs to have correct box-sizing
* 84a383638 satellite/metabase: increase maximum batch size for loop
* 495e53093 satellite/metainfo: drop metainfo.Service
* 67a33e4c0 satellite/metabase: drop unused IterateLoopStreams method
* c258f4bba private/testplanet: move Metabase outside Metainfo for satellite
* 9eaddee53 testsuite: tests for new CLI flow
* b86295c50 web/satellite: Final steps of onb CLI flow
* 54c60fab4 web/satellite: onb cli flow's access grant flow
* 97a5c58ce web/satellite: cli flow's Skip Flow
* 390c6b6fc web/satellite: cli flow's Encrypt Your Data component
* 7a08c19c5 web/satellite: initial setup for onboarding CLI flow
* cd973fcb7 Revert "Makefile: disable building storagenode for darwin" (#4186)
* d867a7007 Makefile: disable building storagenode for darwin
* f55af3125 web/satellite: Replace checkbox reCAPTCHA with invisible reCAPTCHA
* 03d638bbb cmd/uplink: add parallel upload
* 84539f579 web/: disallow warnings and errors in tests
* a7cda642a satellite/repair: add logging for irreparable segments in checker
* 6d876acfb satellite/console: UpdateProject changes
* e5977ec84 web/: add custom linter for requiring @vue/component
* ee4361fe0 satellite/audit: fix segment stripes length calculation
* 2c50210c9 satellite/admin: Don't update project desc when empty
* 7d5d55526 web/: disallow unregistered and unused components
* 0889866b1 web/: add @vue/component annotations
* b64d8084e satellite/audit: fix metric reporting when fail to complete an audit
* 5304f672a satellite/payments: fix invoice generation to include the last day of the month
* 2fafc0e16 satellite/console: Add CORS test
* 25971c581 Jenkinsfile: clarify modfile linting
* 3510bc036 testsuite: fix testsuite/go.mod and enable linters
* 238282b7f web/satellite: fix typo in  `Add a Payment method` popup
* 707c4a7b0 satellite/metabase: use implicit tx when batch deleting
* d0596c060 web/multinode: remove links to empty pages from navigation panel
* a7e151e97 web/satellite: added ability to restart onboarding tour
* ce6ae1b17 web/storagenode: add online score block for selected satellite
* 51fdceafe satellite/repair: increment repair_too_many_nodes_failed with 0 for redash alerting
* 28cb69061 satellite/audit: log error and increment metric if shares cannot be verified
* 26f839a44 satellite/repair/repairer: if not enough nodes for repair order limits, increment metric and log as irreparable segment
* 2340429ea satellite/metabase/zombiedeletion: add missing test case
* eb11899ac cmd/satellite: Remove billing 'check-paid-tier' command
* a14bfdd51 mod: bump common and downgrade sqlite3
* 24e02b635 satellite/{audit,orders}: if not enough nodes for audit order limits, increment metric and wrap error with ErrNotEnoughShares
* 58bd85cbf uplinkng: some windows test fixes
* 1549a6a68 web/satellite: Use OSP logo
* 0376fcdae integration/ui/satellite: fixed tests for onboarding flow
* 445ebf86c satellite/satellitedb: allow setting rate limit for a project to 0
* 211a63098 cmd/uplink: add parallelism flag for single object download
* 5c91ecd27 satellite/metainfo: endpoint cleanup
* c29734ef6 satellite/accounting/tally: remove commented metrics
* 7b4a09c1e satellite/console: Allow basic headers in CORS preflight
* 6daad6873 web/satellite: don't require explicit RootGetter type
* ad0b19fb0 web/satellite: fix lint issues
* 101bbc441 satellite/console: Fix flaky TestMFA test
* f0c3514b2 web/satellite: remove uses of v-html
* 6f2258d87 web/storagenode: fix use of v-html
* b7555980c satellite/metainfo: add zombiedeletion chore
* 6a6cc28fc satellite/console,private/web: Rate limit coupon code application
* 0344790c2 satellite/metabase: delete zombie object that has no new segments for specific period of time
* 1e02504be web/multinode: satellite, node, and usage request handling fix
* 8c8a3e217 testsuite/ui: move from integration
* 5a1a29a62 satellite/audit: fix containment bug where nodes not removed
* 70296c505 satellite/audit: change wording of audit worker error log
* e452f8516 satellite/metabase: sync batchSizeLimit and ListLimit constants
* ca5215c48 web/satellite: rework passphrase step of objects flow
* fafabaa50 web/satellite: redundant images deletion
* 2782e000a cmd/uplinkng: parallel rm
* d57583dfd cmd/uplinkng: initial setup
* 7eaa7490a cmd/uplinkng: secret prompting
* a5371353b satellite/metabase/segmentloop: fix rate.NewLimiter call
* 704cad8da satellite/console/consoleweb: add endpoint tests
* 96e39018c satellite/console,web/satellite: Allow disabling MFA with recovery code
* ef9a5210a satellite/payments: Add command to ensure free tier coupons
* 24b39b755 web/satellite: Warn user if number of MFA recovery codes is low
* cc083dbdc web/satellite,satellite/console: Allow paid tier users to edit limits
* 1fa0cfbfe satellite/console: Add CORS headers for register endpoint
* d22ecf9ec Jenkinsfile.public: pull latest ci image instead of building
* 2cd68bf4f private/lrucache: import from common
* 549e799bb satellite: wrap notfound on failed deletes as DRPC errors
* ea772a848 web/satellite: Allow users to see coupon in billing area
* 6c34ff64a satellite/satellitedb: remove referrence to audit information in nodes and audit_history tables
* 51ebc564d web/satellite,satellite/console: Overhaul password reset
* f80751847 ci: try fix go build scanning node_modules directory
* e43af5eb3 web/: enable more vue eslint rules
* 92c53afb8 satellite/{payments,console},web/satellite: Adds confirmation step if user already has coupon code applied and wants to replace it
* c5d522971 web/satellite: update Overview step of onboarding flow
* 9eb8dc0c7 web/satellite: Fix invalid font reference
* 2f7e67099 satellite/metainfo: allow per-project rate limit of zero
* b8280bd05 web/satellite: enable vue/recommended linting
* abf6bf950 web/storagenode: enable vue/recommended linting
* 5d8cbf087 web/multinode: enable vue/recommended linting
* 59b842231 cmd/uplinkng: implement revoke and normalize some language
* 08d860570 cmd/uplinkng: parallelsm and a ton of fixes
* e33f8d717 cmd/uplinkng: access creation/restriction and review fixes
* dc69e1b16 satellite/repair: use mutex instead of channel to collect download errors
* 077ec96d9 private/server: use quic implementation from storj.io/common
* c074a5666 satellite/satellitedb: improve Update query for reputation
* 6b153192a web/satellite: fix lint issues
* a8f125c67 satellite:{audit,repair}: log additional info when we can't download enough pieces
* 7b152cdde web/satellite: Fix Usage Report style
* 260a67332 satellite/metabase: don't delete pieces until query results are received
* 6613133a1 satellite/metabase: return metadata from IterateObjectsAllVersionsWithStatus on demand
* a60e836ba satellite/admin: Rename endpoints paths entities
* 8d75a35e5 satellite/admin: Add new endpoint get list of API Keys
* 1f353f323 segment/{metabase,repair}: change segment created_at column to not accept nulls
* e00b573f8 satellite/overlay: fix UpdateCheckIn comment
* 7f02e1f46 satellite/console: Remove config flag for MFA
* e4cc965c3 satellite/satellitedb: replace explicit transaction with dbx query for UpdateReputation
* 19852a767 web/storagenode: fix lint issues
* f87ad240e satellite/metainfo: add DeletePart method
* 65804801e all: fix mon.Task leak
* ae02f6ded satellite/reputation: return default reputation stats when node is not found
* cc590343c web/satellite: server side encryption acknowledge checkbox added
* a3e78491b web/multinode: fix lint issues
* f5ac00f90 web/satellite: switch to eslint, sass, bump deps
* f06e7c5f6 segment/{metabase,repair}: add dedicated methods on metabase.Pieces
* 80854ea92 web/storagenode: switch to eslint, sass, bump deps
* 5bb383631 satellite/metainfo: objectEntryToProtoListItem fills metadata on demand
* 0d8e7905c satellite/repair/checker: don't return error when joining loop
* 5e24060b2 web/multinode: switch to eslint
* df52e72fe web/multinode: fix package versions
* 92deef4f3 Revert "storagenode/payouts: historical payouts use satellitesDB instead of trustPool"
* 42a0b9240 web/satellite: set default NotAfter caveat for object browser access grant
* cc277806d web/satellite: regenerate recovery codes popup
* be92ff175 satellite/metabase: fix finishing bucket objects deletion
* 804d221d4 satellite/satellitedb: remove tables unused after metaloop refactoring
* 011b94438 satellite/metrics: fix metrics for total inline/remote bytes and segments
* fda316b46 Makefile: update channel name of #team-data
* 10d7a63ab Makefile: include arm64 (instead of arm) executables in arm64v8 images
* c6c2a1daf web/satellite: added MFA login step
* 2118c89e9 web/multinode: bump node-sass, webpack and lockfile
* 96291a399 web/{satellite,storagenode}: bump lockfile version
* 55d7bcc59 satellite/metabase/segmentloop: don't shutdown satellite on loop error
* 615aae6bd web/satellite: Remove activated account page
* ae345320f satellite/metabase: use db methods that pass context correctly
* d3e51353a satellite/satellitedb/reputation: replace audit_histories table with reputation
* b35b07081 satellite/nodestats: return default reputation when not found
* c8f4ff413 satellite/satellitedb: migrate existing data from overlaycache and audit_histories to reputations table
* 646ce5b8c satellite/overlay: remove reputation logic from overlay
* f8914ccce satellite/{repair, overlay}: use reputation store in repair
* 58238d850 satellite/{audit, accounting}: use reputation store in tests
* e91574cee satellite/{reputation, gracefulexit}: use reputation store in gracefulexit
* 55a77d04b satellite/satellitedb,private: add initial value on testplanet startup
* 6c7bf357c satellite/{reputation,audit,overlay}: replace overlay with reputation package in audit
* 6bbb5b5cc web/satellite: Static error page maintenance fixes
* 49fbb5010 web/satellite: Merge and update 'Forgot Password' page files
* 8f072bdee cmd/satellite: Skip non-existing users in paid tier conversion
* d73b9fff9 satellite/orders: set the expirationDate in CreatePutRepairOrderLimits
* a2b30e42a satellite/metainfo: log internal errors
* 3720afb81 web/satellite: Update emails to be consistent with deployed code
* 276ebcd4b satellite: Remove duplicate logo in error pages
* 149f6f262 satellite/payments: Implement coupon codes
* dae6ed7d0 satellite/console: Implement MFA backend
* 420d2f627 metabase-orphaned-segments: avoid processing recently created elements
* 42b113d92 web/storagenode: fix ingress chart tooltip
* 267a962c3 web/storagenode: fix egress chart tooltip
* a883d7f58 satellite/repair/checker: fix `remote_files_checked` metric
* b12d29935 satellite/metabase: remove metaloop package
* 1535bbe67 satellite/console: Forward friendly registration errors to client
* 237360e76 metabase-orphaned-segments: fix args binding for delete command
* 6d60d412f satellite/gracefulexit: use segment loop
* c4202b945 satellite/gracefulexit: use graceful_exit_segment_transfer_queue
* 449c87368 satellite/repair/repairer: attempt repair GETs using nodes' last IP and port first
* b0d98b1c1 satellite/gracefulexit: allow use of graceful_exit_segment_transfer_queue
* 73cdefbc4 munltinode/console: node online statuses extended with StatusUnauthorized and StatusStorageNodeInternal
* f73fcee89 cmd: add tool for orphaned segments
* 373ba8fd2 satellite/repair/repairer: metrics for repair bytes uploaded and downloaded
* adc0fbddf satellite/audit: don't fail nodes for audit if not enough pieces downloaded
* 27a714e8b satellite/accounting/tally: use objects iterator instead metaloop
* 4c7384ca2 web/satellite: removed redundant files splitting
* 2489b78d2 web/satellite: move user MFA secret generation to server-side
* 5a56021bc satellite/metabase: add intLimitRange
* bc2f81c2f web/satellite: Add beta checkbox to registration page
* 7624bdd09 web/satellite: Return server-side error message during registration
* 1e0a1b15d web/satellite: Repair reCAPTCHA resetting upon registration error
* 4a98dd40e storagenode/payouts: historical payouts use satellitesDB instead of trustPool
* 18c325202 cmd/uplink:  recommend notAfter expiration caveat be set
* 22e88c8f0 web/satellite: disable MFA functionality
* e463eb17a web/satellite: added enabling user MFA functionality to account settings
* 0d8010e35 satellite/satellitedb: Implement multi-factor authentication db
* 4c912c747 web/satellite: Update paid tier banner copy
* 6db6b76b2 multinode/console: embed web assets
* 30cd7d3da satellite/console: Update CSP for satellite UI
* 587050258 cmd/satellite: Add billing command for converting customers to paid tier
* 8855c0dff web/satellite: added MFA feature flag, updated client-side api and Vuex store module
* ee107fe8c satellite/rewards: adding partners to list (#4159)
* 76c8d5d2c satellite/metrics: move metrics to segment loop
* 8c052048b web/multinode: redirect fix before route enters dashboard
* 4d0fe3923 storagenode/satellites: address added, caching satellite's addresses from trust
* a767aed59 satellite/{metabase,metainfo}: require StreamID for UpdateObjectMetadata
* cbbbfca43 satellite/accounting: move storage node tally to separate service
* f8eebbc11 web/multinode: npm packages version audit & update
* a5f6bb9cc satellite/console: Add reCAPTCHA verification step to registration
* ec9ad5bd7 web/satellite: Don't require passphrase entry after generation
* 049570384 web/multinode: reputation domain and api created
* f79d69771 multinode/nodes: checking node availability before adding it
* cbb4cd3fc multinode/reputation: add vetted at timestamp
* a5fd90317 storagenode/reputation: add vetted at timestamp
* 2e9d3a737 satellite/metabase: fix listing prefixes with cursor set
* 8f4505f53 multinode: handling offline nodes added
* 73b922375 satellite/metainfo: implement UpdateObjectMetadata
* 818f6c6ea multinode/console: add summary to storage usage API
* b0ef9b9b6 Create MAINTAINERS.md
* ef7b89cc0 cmd/uplinkng: remove global flags
* f474bb617 cmd/uplinkng: update for breaking clingy changes
* bbd3efaee web/satellite: Paid Tier add payment modal implemented
* 4d418c13c web/multinode: routes restriction in no node added
* b8e6e5c6c web/multinode/nodes: status calculation placed to backend
* e36001b7c satellite/console: Implement paid tier
* 3b0b0ba3c multinode/nodes: node status extended with NotReachable
* ecf3de3c3 web/satellite: Add new intro section to register page
* 60717e356 UItests: added an onboarding wizard test
* 6c0225892 storagenode/peer: register multinode Payouts endpoint
* 74af44c12 satellite/metabase: make IteratePendingObjectsByKey recursive
* a6086d421 satellite/metabase: implement SetObjectMetadataLatestVersion
* 68627e7d8 multinode/console: add reputation satellite api
* bf5194d13 satellite/satellitedb: optimize ProjectAccounting.ArchiveRollupsBefore
* c248651f3 satellite/metabase/{meta,segment}loop: fix test flakiness
* 7cbff9509 cmd: delete metabase-expireat-migration
* bab43af6c web/satellite: added Upgrade to Paid Tier banner
* 1230cde31 web/multinode: disk space usage and stat charts
* b900f6b4f satellite/repair/checker: move checker to segment loop
* 8686267e0 web/multinode: storage api, service and store module created
* 55754df11 web/multinode: bandwidth related unit tests
* cece8e411 web/multinode: bandwidth charts added
* cc5de4288 web/multinode: bandwidth api and store module created
* 0ca758328 satellite/accounting: add total for bytes and segments to tallies
* e1379bea0 storagenode/piecestore: allow rejecting slow clients
* 4335b2133 satellite/metabase: add TotalEncryptedSize to LoopObjectEntry
* d53aacc05 satellite/repair: migrate to new repair_queue table
* 0ec3867ec satellitedb: add paid_tier column to users table
* aa7fd8c9c satellite/satellitedb: fix ProjectAccounting project bandwidth calculation
* 1c3688f1d storage/filestore: include satellite id in open_file_in_trash metrics
* cb18dc77f satellite/stripecoinpayments: version is the wrong name
* d999a963c satellite/console/.../consoleapi: Standardize serveJSONError
* 371517d93 satellite/accounting: fix as of system interval value for some tests
* 4e95d2703 web,satellite: Remove paywall-related functionality
* 4267a958d satellite/satellitedb: use AS OF SYSTEM TIME for GetProjectBandwidth query
* 70e6cdfd0 satellite/audit: move to segmentloop
* 8ce619706 satellite/audit: migrate to new segment_pending_audit table
* a89c0763a satellite/gracefulexit: create table graceful_exit_segment_transfer_queue
* d30fd7765 satellite/console: new endpoint to get total usage and limits for all the projects user owns
* 4031336cb private/lifecycle: monitor unexpected shutdowns
* 3c8e7e679 mod: bump storj.io/private and monkit
* 1f5fbbd24 satellite/metabase: while deleting bucket objects delete pieces in batches
* a59232bb1 satellite: return directly to avoid shadowed err variable
* 8c124c6fa satellite/{reputation,overlay,satellitedb}: create reputation service, DB, add overlay method UpdateReputation
* cff3da764 multinode/operators: using old api instead of new until nodes updated.
* e37df4491 Create a New User Business Account Test (#4148)
* 07f8cff08 satellite/satellitedb: Add reputation table
* b1201df82 satellite/console: add more tests for uplink access permission
* b9d7874dc multinode: old drpc api returned
* 8a6efa1f5 satellite/orders: query for node first before upsert/replace
* 70efbd983 jenkins UI fix (#4146)
* d9741491a web/satellite: added cancel upload warning popup
* 7d3b5b932 web/satellite: replace old fonts with new versions
* 59e3b586e satellite/{gracefulexit,overlay}: enable as of system time queries
* b582c974c satellite: remove irreparabledb leftovers from code
* f16bb4d19 Create a New User Personal Account Test (#4141)
* 4e645059b satellite: rename Endpoint2 to Endpoint
* 8a070e7c2 satellite/overlay: Ignore unnecessary check-ins
* 7815e647d cmd: add metabase-expiresat-migration tool
* 036687b58 docs/blueprints: Create coupon codes blueprint (#4024)
* ca64e5528 satellite/gc: remove skip first
* 8fec48f31 storagenode: use storj.io/dcs-satellites now
* 5f901c806 satellite/metabase/metaloop: add as of system interval to stats
* 9640cc2c0 satellite/metabase/segmentloop: verify processed count
* 341033dda satellite/metabase/metaloop: verify processed count
* a93e47514 satellite: remove irreparabledb
* 0df4a27bf satellite/metabase: add method to get table statistics
* 2ebbc9757 mod: bump storj.io/private
* 9b2607d6b satellite: remove garbage collection option from core
* 08c002533 satellite: remove contact service from core
* a89e20895 web/multinode/operator: wallets features type changed to support null
* d426d4f00 multinode/storage: disk space of concrete node and total disk space implemented
* 1a974c327 satellite/metabase: set Segment.ExpiresAt while getting segment
* 6949dc0ba satellite/metaloop: missing monitoring on observers
* ce87652a8 cmd/uplinkng: rm
* 173d1e638 cmd/uplinkng: meta get
* 1323e1a66 cmd/uplinkng: rb
* 7fae5654f cmd/uplinkng: access save: prompt for access
* d73287f04 cmd/uplinkng: tests for cp
* 98be54b9a cmd/uplinkng: refactor into some focused packages
* b24ea2ead cmd/uplinkng: test framework and ls tests
* f1a9b4559 mod: update quic for Go 1.17rc1
* 81dd7b2f3 storagenode/internalpb: fix protobuf generation for newer protobuf version
* 9a113da36 satellite/metabase: expose expires_at with loops
* 5d70b6abe multinode/bandwidth: added monthly bandwidth summaries
* f3a52d1da satellite/metabase/segmentloop: limit max interval
* bc79f01aa satellite/metabase: set expires_at while committing segment
* 157d3d980 multinode/console: storage usage and total storage usage
* ed28fa3ff web/satellite: added loaders across all the UI. Removed most of the requests from initial load
* c27da9574 Update index.html to be consistent with deployed code (#4140)
* b23a782c6 multinode/console: held amount summary use satellite address
* d674bc9c5 satellite/audit: include failing segment info in logs
* 944bceabc satellite/audit: fix reservoir sampling bias
* 1cd13adc1 ci: clearer initial build
* 072dd2469 integration/ui: add UI integration tests
* 962f81433 web/multinode: operator entity extended with nodeID and undistributed amount
* b445e7d39 multinode/operators: operator entity extended with nodeID and undistributed amount
* be16f4b68 web/multinode: wallets logic
* c0b170149 web/multinode: wallets page markup
* 21731ff8d satellite/metabase: add expires_at column to segments
* 347f5f87e satellite/metabase/metaloop: limit as of system time
* 3de9655b6 satellite/redis: used bandwidth key depends on day
* c9cfb5ed0 storagenode/retain: add more verbose monkit monitoring
* e6fe9d209 satellite/metabase: capture iterator errors
* 504433744 satellite/metabase: add monitoring for objects and segments
* b17d684f4 web/satellite: Update login area to use latest production code
* cbf593998 satellite/metabase/metaloop: fix objectsIterated metric
* 1d405f45e satellite/orders: add egress_dead to project_bandwidth_daily_rollups
* 8c0a0f019 web/multinode: payouts by node
* 556b80e06 web/: ignore node_modules while linting css
* 9c46490f6 satellite/accounting/projectbwcleanup: more realistic test value
* 49662e98c web/multinode: node selection component added
* 4469d229f satellite/metabase/{meta,segment}loop: avoid passing config
* 46a3242ed cmd/uplinkng: cp
* 136af8e63 web, satellite: allow registering business accounts to ask for contact from sales team
* 2cf10a7bf satellite/metabase/{segment,meta}loop: avoid no observers error
* a04c96b1b web/satellite: replace initial cloud loader with a new one
* c26d52a53 web/satellite: remove duplicated default project creation code
* 9e2d83747 testplanet: remove brittle config
* d987990c1 satellite/satellitedb: removing usage of project_bandwidth_rollups table
* da9ca0c65 testplanet/satellite: reduce the number of places default values need to be configured
* 75bf2ad58 satellite/analytics: Add analytics for "link shared" in objects view (#4130)
* 53322bb0a satellite/{audit,satellitedb}: release nodes from containment in Reverify rather than (Batch)UpdateStats
* e76cbc9bd satellite/gc: move GC to segments loop
* b07a39bfe satellite: log check in success node id
* 8675cd178 cmd/uplink: add flags to 'put' subcommand
* fbd11ba4e cmd/uplink: add expanded format for 'uplink ls'
* d8c11a79b multinode/payouts: paystub extended with disposed
* 4af0037a6 multinode/console: operators controller added
* 0b3ee4bda multinode/operators: domain entities and service created
* 053e58b68 satellite/metabase: add segmentloop service
* 6d9b91d43 private/multinodepb: drpc operators controller added
* 4c704ea9a mnd/db/nodes: paginated list added
* 10a0216af satellite/metainfo: use range for specifying download limit
* aa49c8c44 satellite/orders: fix TestProjectUsageBandwidth
* 91b7e24d5 multinode/payouts: naming refactoring
* 20b98d31b web/multinode: expectations block added
* 64e422572 multinode/console: server reorder payouts API routes
* 23f9beb63 multinode/console: held amount summary
* 0ef537a68 satellite/metabase/metaloop: limit max as of system time
* a5dbc544f satellite/console: project member deletion bugs fixed
* 32dd4315a web/satellite: improve AG's selection checkmark
* 3af91e7a9 satellite/metabase: add iteration over segments
* 37bc46c6f web/satellite: Fix for Payment Method area buttons on billing page not being aligned
* f536ee3aa satellite/satellitedb: Migrate non-expiring coupons to expire
* 63cfc8fbe satellite/orders: use project daily bandwidth rollups
* 59eabcca2 satellite/orders: populate project_bandwidth_daily_rollups table
* 02fc87e98 satellite/payments: Apply Stripe free tier coupon for new customers
* 92226d8dd storagenode/storagenodedb: fix fillInBlobAccess
* 099a31c69 multinode/payouts: paystubs for single period added
* 1bc4f8fcc multinode/payouts: paystubs for all time added
* 926bcd0f8 mod: bump private
* 6e05b240d private/lifecycle: log stack on slow shutdown
* 16022105f satellite/orders: create project_bandwidth_daily_rollups table
* c410eae17 web/satellite: better error message for duplicated bucket
* b33b0a38d web/satellite: remove redundant console errors on duplicated route changes
* 0ff3516f5 satellite/mailservice: Fix bug causing issues with test account creation
* e7cf369cd satellite/metabase: use pgxutil.Conn helper
* e460dc51f cmd/uplinkng: implement object listing
* cdcc67207 satellite/satellitedb: fix nil panic in UpdateCheckIn
* 79172777b mnd/payouts: estimations replaced with expectations
* 83e82eb47 mnd/payouts: estimated payouts per satellite changed to per node
* 515f9b54f mod: bump monkit
* f0dd7f373 mnd/payouts: undistributed payouts added
* 2896a2736 web/satellite: remove time caveats for object browser AGs
* f2842a27e web/multinode: by node payouts page markup
* c9d467485 cmd: remove metabase-createdat-migration tool
* f0d60a6ec Jenkins: add multinode UI stage
* 1ff2e1a47 web/multinode: update test snapshots
* 2d3c40fd7 cmd/internal/asset: delete unused code
* a9fcbf728 web/satellite: bump object browser version
* a1bf9ab6d cmd/uplinkng: initial commit with skeleton
* 10372afbe ci: fix lint errors
* 1852773e3 satellite/contact: rate limit node checkins
* 8f15f975a satellite/overlay: improve contended update checkin
* b42624893 scripts/tests: fix missing multinode binary
* f2812d76c multinode/payouts: satellite period/allTime summaries added
* dfe85beac satellite/satellitedb/dbx: fix duplicated lint line
* be87c80fe satellite/satellitedb: drop columns total_uptime_count and uptime_success_count
* 3ea7aa2c7 satellite/repair/repairer: log piece hash verification failures
* e0f316634 web/multinode: store and payouts components tests added
* 2ccbb32e1 web/multinode: payouts page markup
* 6ee221029 satellite/console: add test for time based permission
* ce075a1d5 web/multinode: payout api
* c08ca361d cmd/storj-sim: add multinode process
* ea7fbdf84 multinode/multinodedb: add db migration
* 910eec8ee satellite/metainfo: remove MetabaseDB interface
* d32ae0459 Revert "storagenode: docker image autoupdate binaries"
* 59f443e71 storagenode/contact: add authentication for PingNode endpoint
* 19561698b multinode/payouts: all satellites summaries added.
* 547a6e993 satellite/metabase: add DeletePart method
* 9ccebfa7f web/satellite: create select input component to capture employee count in registration
* 2ae80690c satellite/metabase/metabasetest: remove STORJ_TEST_DATABASES
* a11698f37 multinode/payouts: estimated payouts added
* 0858c3797 satellite/{metabase,satellitedb}: deduplicate AS OF SYSTEM TIME code
* 033006403 satellite/metainfo: fix setting object ZombieDeletionDeadline
* 7fb86617f satellite/satellitedb: Use CRDB AS OF SYSTEM & batch for GE
* 6161436d8 satellite/metabase/metabasetest: sort results
* 6e6051b17 satellite/metabase: drop alias migration code
* 2af7e4ef2 satellite/metabase/metaloop: use database time
* 60ff87a7d design doc: server-side move
* 205ea621c .github: remove noisy dependabot
* 38c41d094 satellite/admin: Fix typo in README
* a89a2b4b4 satellite/billing: make stripe invoice generation work with multiregion satellites
* 69b149a66 mod: bump uplink
* 2f8969ae3 satellite/admin: Add a ToC to README
* 150430336 satellite/admin: Add clarifications in README
* 0d3865950 satellite/metabase: rename delete_expired.go to delete_objects.go
* d2033c2f5 satellite/nodeselection/uploadselection: rename package
* 5269596c7 satellite/{metabase,metainfo}: use ObjectStream as argument
* d06206488 satellite/rolluparchive: archiveAge use 90d prod default for storj-sim
* dba932148 satellite/metabase: Remove pending_index
* 834b7cf6b private/server: close connection after each connector test
* efcb8e412 .github: enable dependabot for upgrading go dependencies
* e8ef68992 satellite/metabase: add DeleteZombieObjects method
* bb343d902 satellite/satellitedb: don't remove offline nodes from containment
* 0084164f3 Fix link to whitepaper in README.md (#4107)
* 222d9b7f0 web/satellite: add script to compile and compress wasm module
* a8533042a mod: bump uplink
* 1b736104c satellite/payments: Apply free tier coupons before preparing invoices
* dae9464ee storagenode/trust: missed one error test
* 396cd5a68 satellite/payments: reduce object fee to 0 (#4104)
* 2c657f594 satellite/metainfo: don't rely on exact error name
* 65b17df58 bump storj.io/monkit-jaeger to latest
* 5f29a2093 satellite/satellitedb: Add missing indexes
* 961e841bd all: fix error naming
* 30d2a2ea7 web/satellite: add total account balance to billing page
* 4c0817bcf satellite/payments: Populate new coupons during invoice generation
* 2e5e4cb64 web/satellite: access grant wizard refactored
* b317f28fd web/satellite: Update registration files
* bc9545513 web/satellite: delete bucket popup reworked
* 2936b6755 web/satellite: wording updates
* 75ba99b88 web/satellite: bump object browser component version with part size fix 2.0
* 5af898fcc satellite/metabase: add migration step for index on pending objects
* be7928d4b crashcollect: removed redundant structure
* 917e4b81d web/satellite: regular header reworked
* a5c1e4b4a crashcollect: process for receiving panics
* 02460fcc4 satellite/metainfo/expireddeletion: change chore interval to 24h
* 606a50144 satellite/console: add tardigradeshare.io domain to media-src CSP whitelist
* f2e06ce88 web/satellite: added skip button for the onboarding flow
* 64c82507c web/satellite: make project dashboard delay message more visible
* a84b8367d web/satellite: changed Token to API Key label for CLI step of AG flow
* c1fdf1ad7 satellite/satellitedb: add have_sales_contact column on users
* 7944df20d storj: use multipart API
* aa27a8f1d private/crashreportpb: add protobuf for crashreport and crashcollect
* 948529a60 satellite/metabase/metabasetest: move test utils to a separate package
* 307886ffe storagenode/piecestore: fix error handling in TestDownload
* 31c113786 web/satellite: add login path
* 08e1ce893 cmd/connecttest: display more helpful error message and fix tls config
* 7802ab714 pkg/,private/: merge with private package
* c641ddcb5 pkg/lrucache: rename package
* 4d3737812 storage/{cockroachkv,postgreskv}: delete unused code
* 8c62788b2 satellite/metabase/metaloop: stop timer properly
* d32515fa5 satellite/console: add storjshare.io domain to media-src CSP whitelist
* a2e20c93a private/dbutil: use dbutil and tagsql from storj.io/private
* dd5eb2616 satellite/metrics: join for monitoring
* c1fbecb96 satellite/metabase/metaloop: add Monitor
* fff21b330 cmd/metabase-verify: tool for verifying metabase state
* 630787520 monkit: fix monkit lock
* dcb73e46c cmd/connect-test: add connect-test tool
* a0c5da664 satellite/satellitedb: in stray nodes DQ, don't DQ nodes where last_contact_success = '0001-01-01 00:00:00+00'
* 4c9ed64f7 satellite/metabase/metaloop: move loop under metabase
* ebc981291 satellite/accounting: use correct error in tests
* 52c3d7d64 web/satellite: add disabled state to onboarding overview RClone link
* 43f588805 satellite/contact: only test PingBack failure case
* 244c488b1 storagenode: docker image autoupdate binaries
* a9c9f080b satellite/metaloop: Use AS OF SYSTEM TIME when querying the objects table
* 2537bbf54 satellite/gracefulexit: Try avoiding randomly test failure
* 8b2e18c91 pkg/quic: ignore successful session closed error
* 267506bb2 satellite/metabase: move package one level higher
* 2ee303027 all: remove code related to PointerDB
* 785adfb84 go.mod: bump to uplink v1.4.6
* a3c437a7b satellite/contact,storagenode/contact: try ping back to nodes through QUIC
* f6ec7f9bf web/satellite: remove paywall steps from onboarding flow
* f8aaf5c5f web/satellite: bump object browser version with part size fix
* a6b3a565d jenkins/build: fix file extension for msi package
* 546e2d2e8 web/satellite: remove details button on info bar for Dashboard page
* ebe581396 satellite/emails: update mailing address
* 5642a7a8d docker-compose: Delete deprecated file
* 613a95530 satellite/analytics: Add analytics for "path selected" in onboarding step (#4086)
* 75ca01e38 cmd/uplink: fix importing access in Windows
* 6a805b289 satellite/metainfo/metabase: use pgx.Batch to delete expired segments
* b2be1f162 satellite/metainfo/metabase: fix delete expiration
* 9cd17fd80 web/satellite: request passphrasse on project change. Small styling fixes
* b57819f59 satellite: take pricing from the config instead of hardcoding
* f9a6fbea5 web/satellite: Replace Tardigrade logos with Storj DCS
* f380bda81 web/satellite: add paths for register and login pages to allow for travel between both
* dfe4d09a8 web/satellite: Update login logo to DSC branding
* 35d8a840d storj/storj: more domain changes
* d7488924c web/satellite: bump file browser component version
* 1215500c8 web/satellite: replace spaces with %20 in linksharing url
* 63c3dd2b2 web/satellite: add bucket name restriction info to popup
* 68f67e535 web/satellite: trim passphrase for file browser flow
* 29e55d69b web/satellite: overview page copy and routing changed
* bd36a41a9 web/satellite/file_browser: open bucket on creation
* 3b09d6c30 storj/storj: update support request links
* f3c8cac1e storj/storj: more domain changes
* 569942c04 web/satellite/src/api: Absorb analytics errors
* 031206e45 web/satellite/tests: Fix snapshots
* 034519a12 web/satellite: Add security and encryption tags to overview onboarding step
* 32d76652c satellite/payments: Change customer pricing (#4085)
* e71052425 web/satellite: Update signup flow Storj logo
* bde367ae7 satellite/gc: check on bloom filter creation date
* 2f39fc0f8 web/satellite: updated file browser prep flow
* 240d57138 web/satellite: Add download button for access grant
* 56c872cc4 web/satelite: replaced duplicati with rclone for overview step of onboarding flow
* c3f8e0679 satellite/metainfo/metainfo: simplify expired objects deletion query
* afcc55fb4 storj/storj: some domain changes we can make now
* 0b59a165e web/satellite: remove google tag manager from satellite GUI
* 027554a15 satellite/metainfo: add tests for commit segment validation
* 3236f91ae satellite/metainfo: reenable piece hash verification
* 6e7992399 satellite/metainfo: bring back validation while committing segment
* 6ddd22e55 web/satellite: remove redundant FAQ link from no buckets view
* 3ade87eb2 satellite/console: added tardigradeshare.io and storjshare.io domains to CSP's whitelist
* 40b41e081 web/satellite: onboarding overview step updated
* d2705c114 satellite/analytics: Added analytics for "passphrase created", "account verified" and "external_link_clicked" (#4078)
* 037c7799b storj/storj: bump storj/common
* 6f5f2dd36 web/satellite: clear objects vuex state on destroy. Always fetch buckets list
* 2a4a70908 cmd/satellite: remove metrics from gc
* 6236deffb satellite/web: enable feature flag for new file browser and onboarding workflow (#4079)
* 9591e7557 satellite/console: add feature flag for onboarding overview step pathway rendering
* 8f13dcacb mod: bump uplink
* ec7af97a1 satellite/metainfo: document migrated objects
* 51d667a65 satellite/analytics: Add anonymous ID to TrackCreateUser
* b907aab0a web/satellite: use config value for linksharing
* a3e660488 satellite/console: config flag for enabling/disabling CSP
* c4a950a40 web/satellite: Remove client-side Segment analytics
* 16c98e1ec satellite/analytics: Add analytics for user signed in, project created, and access grant created (#4073)
* a264a4422 cmd: remove metainfo-migrator and metainfo-verifier
* 6b88a675c metabase-createdat-migration: add migration tool
* 372bb1140 mod: bump uplink and private
* 0a512108a web/satellite: Remove noPaywallInfoBar
* 3088d1411 Overview refactor (#4076)
* 6ae235138 web/satellite: import file browser component
* 444b1f475 storagenode/piecestore: update endpoint_test TestDownload to avoid index out of range runtime errors
* 71072eb59 storagenode/pieces: send piece deletions to trash
* 7e4e1040f satellite/console: Add endpoint for clientside analytics events
* 4330a4bc4 pkg/quic: set KeepAlive to true
* ffc8aca8b satellite/metainfo: update bandwidth usage in DownloadObject
* 98b7e1a99 satellite/metainfo: fix returning []{nil} in DownloadObject
* 47be50118 satellite/metainfo: add DownloadObject endpoint
* 17817d84e satellite/metainfo: add PlainOffset to SegmentDownloadResponse
* 288a7e8ee web/satellite: split AG wizard's result step into two
* f3df411b4 bump storj.io/common to latest
* d2101bd01 satellite/metainfo: add Position to SegmentDownloadResponse
* 04294e3e1 satellite/payments: Update coupon defaults for free tier
* e0c3a7a47 bump quic-go to latest
* 2462ac591 metainfo+satellite/satellitedb: use HasBucket when possible
* 523dd823d satellite/satellitedb: remove references to total_uptime_count and uptime_success_count columns
* a5224e7a6 satellite/metainfo/metaloop: use segment CreatedAt and RepairedAt
* 71d0adf90 satellite/metainfo: avoid error on inserting existing attribution
* 5c038c432 satellite/metainfo/metabase: add stream range for listing segments
* c334fd090 satellite/payments: Convert coupon duration to be nullable
* c4293f5e5 storagenode/payouts: DQ satellites removed from estimated payouts
* 0b7c28815 test-sim: improve integration test scripts
* f6c9b8d4f web/satellite: access grant wizard copy edit (#4066)
* 7dfb89a87 bump quic-go to latest release
* a5bc38b19 build: replace 'multipart-upload' branch occurences
* 86c41790c satellite/metainfo/metaloop: add observability
* f3176adbc private/migrate: Add missing validation table call
* 8d4df93d4 satellite/metainfo: remove TODO
* 794dc6d6d satellite/metainfo/metabase: add PlainOffset to segment listing
* c4cb79ff1 satellite/metainfo/metabase: update segment plain offset, when needed
* f2718734f web/storagenode: bug fixes
* b91ba7057 web/storagenode: added zksync wallet explorer button
* 5dbc91622 metainfo-migration: fix flaky TestMigrator_SingleSegmentObj test
* 23c8331d1 satellite/analytics: Add flag for disabling analytics (default disabled)
* a65aecfd9 compensation: always generate invoices for every node
* 035c393da satellite: update tests to pass etag.Reader to multipart.PutObjectPart
* 74b243d1c certificate: move protobuf to storj/storj repository
* 141444f6d satellite/repair/repairer: fix segmentAge metric
* 4b14839e1 make: bump-dependencies should update uplink from main branch
* d57873fd4 satellite/overlay: remove Inspector
* 86e698f57 pb: use *UnimplementedServer to avoid breaking API changes
* c8a3e0ff9 tests: fix calculation of previous month for billing integration test
* 05f8d2d0b satellite/satellitedb: filter offline suspended nodes from selection
* 1a51049ac satellite/{overlay,satellitedb}: add flag to toggle suspending nodes for offline audits
* eb44dc21b satellite/satellitedb: select stray nodes for DQ in separate tx from update
* 2ddbaf1eb satellite/wasm: support restricting full access grants to paths
* 527b7eb56 web/satellite: ui changes to registration and billing pages to enable user to add promo codes to their account - add Credit History table to billing acount page and set up ui for a user adding promo codes - implement promo codes ui in registration form - add feature flag to handle if coupon code ui should be rendered
* 9bc2a9a4b rollingupgrade: change how metainfo-migration tool is installed
* 3e91d22ce satellite/metainfo/metabase: update repair_at with UpdateSegmentPieces
* 2a68e0767 scripts/tests/testversions: Fix testversions Episode 2
* 63cd807c9 scripts/tests/testversions: Increase bandwidth/storage limits for testversions
* 3db52491e satellite/analytics: Add analytics service to satellite
* c4b2d76d1 web/satellite: buckets view for objects page
* ebf6bee0d web/satellite: import and setup AWS s3 client
* 346b85b66 web/satellite: generate gateway credentials for objects page
* 1c696168c satellite/metainfo/metabase: delete objects and segments in one query
* 0ef0a3f96 rollingupgrade: add migration to metabase
* 6ddcacbe7 web/satellite: initial setup of bucket's view of objects page
* 8e1aa4bb7 satellite/metainfo: etag support for committing and listing segments
* ec929ca1d satellite/metainfo/metabase: add support for encrypted ETag
* cc42dfcc7 satellite/metainfo: add EncryptionParameters to SatStreamID
* c5cb4dce4 redis: Rename functions prefixed with New by Open
* 903b1b5ff satellite ui: don't call the gateway mt the gateway mt
* 1156d8b4e mod: update drpc to v0.0.20
* 866a3e3ad storagenode/console: estimated payout tests fix
* 27bcb4671 satellite/console: change default project and usage limits
* 034eb2c21 storagenode: wallet features
* 3e37d1e71 satellite/console: delete api key by name and project id endpoint
* f19ef4afe satellite/metainfo/metaloop: move loop to a separate package
* 972d40df6 pkg/server: use drpcmigrate listenmux
* 54c2ace48 mod: update drpc to v0.0.19
* 50b08a5b0 metainfo-migrator: handle 'omit' conn string
* 9d5245146 satellite/metainfo/expireddeletion: add config value to set list limit
* 237782813 Merge remote-tracking branch 'origin/multipart-upload'
* d2b365bf9 satellite/payments: Disable paywall (free tier)
* 37602d86d satellite/satellitedb: drop columns uptime_reputation_alpha and uptime_reputation_beta
* 864ad70fe satellite/overlay/straynodes: set --stray-nodes.enable-dq release default to true
* 2607b1607 satellite/{overlay/straynodes,satellitedb}: rework DQNodesLastSeenBefore to return DQd node IDs and last contact successes
* a8b5f8f99 satellite/metainfo/metabase: select then delete expired objects
* 2525e1584 web/multinode: favicon added
* 536990dcf Jenkins: Run integration test Redis unavailability
* fa8b0a29e scripts: Do not use Docker for test-sim Redis test
* 27ae0d1f1 satellite/metainfo/metabase: add NewRedundancy parameter for UpdateSegmentPieces method
* 78e132cb7 satellite/satellitedb: Add coupon_codes table and allow optional link to coupons table
* e613c641d satellite/metainfo/metabase: benchmark expired deletion
* 1c0215862 satellite/metainfo/metabase: add RepairedAt to RawSegment
* c3ae122aa web/satellite: enter passphrase step for objects page
* bf4fdc771 cmd/uplink: add --client.enable-quic flag
* a04495713 satellite/audit: add missing logs for audit failure conditions
* be8a70309 satellite/metainfo/metabase: rename etag column to encrypted_etag
* 4c1098e57 Redis: Update Redis package to last major version
* 09adfc120 mod: bump monkit to v3.0.10
* 36bc17ef5 satellite/metainfo/metabase: add etag column to segments table
* cff1052bc web/satellite: save passphrase hash in local storage
* 9e20cfb6e web/satellite: create passphrase step for objects page
* b5f83f463 web/satellite: routes and initial components for objects page
* fa083a7f0 Merge remote-tracking branch 'origin/main' into multipart-upload
* 9491df76e satellite/satellitedb: optimize StoragenodeAccounting.SaveRollup
* 6a553ec9c web/satellite: change banner for beta satellites with URLs
* 1de6315b8 metainfo: Stub DownloadObject metainfo endpoint
* 65957c752 satellite/metainfo/expireddeletion: temporarily log errors instead failing whole system
* a36d854d8 bump monkit-jaeger to latest
* 0b060c829 storagenode/console: estimated payout flaky test skip
* c4f028990 satellite/metainfo: ListSegments returns CreatedAt in response
* 6e661da0a satellite/metainfo/metabase: read created_at from DB
* 487ad84b2 storagenode:consoleAPI flacky tests for estimatedPayout fixed
* 887f3b04e satellite/metainfo/metabase: pieces vs redundancy scheme validation while commiting segment
* dcdfa1037 jenkins/build: rename release binaries inside zip archive (#4060)
* 018b0684b satellite/metabase: different queries for delete_expired
* bdb33b381 satellite/metainfo/metabase: add created_at and repaired_at to segments
* 8138939a2 satellite/satellitedb: remove bad migration step
* e0f07ef04 metainfo-migrator: ignore invalid segment keys
* 80ae7787e metainfo-migration: add simple retry logic for queries
* 5616795f7 web/multinode: update node name modal
* 5e954ad48 satellite/metainfo/metabase: optimize ConvertAliasesToPieces
* 67e26aafc Merge remote-tracking branch 'origin/main' into multipart-upload
* 84b844a2a redis-server: Move testing type to specific testing pkg
* 1f59a08e8 pkg/server: eliminate `Retry` packet from quic-go hanshake phase
* 580eecf0b mod: bump monkit to v3.0.9
* 110c16ddc metainfo-loop-benchmark: run benchmark on the top of storj.io/private/process
* 4d08f0613 metainfo-loop-benchmark: add memory statistics
* 461d2cd23 satellite/metainfo{,/metabase}: add GetStreamPieceCountByNodeID
* ce4024a52 metainfo-migrator: handle objects with missing segments
* 981f1ca44 satellite/metainfo: use ListStreamPositions
* 5bc0f3bf2 satellite/metainfo/metabase: add ListStreamPositions
* 903cb172c satellite/metainfo/metabase: add iteration with prefix benchmark
* b5e95ec79 satellite/metainfo/metabase: add extra validation to IterateLoopSegments
* 150d75da2 satellite/metainfo/metabase: remove old TODO notes
* 703ca08b7 web/satellite: remove api keys related logic from client side
* 8de1116ee satelite/metainfo: remove unused method from metainfo loop
* 7af245c3e satellite/payments: De-shadow customer page var in stripe mock
* ec6741377 satellite/metainfo: add as of system time to object iteration
* b0b7b8110 satellite/metainfo/metabase: streams iterator
* 376547c33 satellite/compensation: smaller txns for RecordPeriod
* a44974a2f satellite/audit: fix pointless containment deletions
* a7ca78a51 web/satellite: Add user account tabs to signup ui - toggle logic to trigger visual tab change - toggling tabs renders appropriate fields - error handling for professional account inputs - successfully create professional user account
* c51ea68ad satellite/metainfo/metabase: reduce number of fields for LoopSegmentEntry
* c290e5ac9 satellite/orders: decrease FlushBatchSize default to 1000
* c223c2d84 satellite/metainfo/metabase: rename ListObjectsSegments to ListLoopSegmentEntries
* 8e9711151 satellite/metainfo/metabase: reduce fields number for LoopObjectEntry
* 7e72a231c satellite/metainfo/metabase: cast bucket_name to []byte
* 932039ebd blueprint: layer2 support for zksync
* ba0197a9b satellite/metainfo/metabase: avoid full table scan
* aeac6264c sallite/satellitedb: add metric stray_nodes_dq_count
* 646cf229a satellite/metainfo/metabase: rename FullIterateObjects method to IterateLoopObjects
* 856db68fd satellite/gracefulexit: extend GE data cleanup to include exit_progress
* 411a7ad0b satellite/satellitedb: drop uptime_reputation_alpha and uptime_reputation_beta from nodes table
* 994854dcd satellite/metainfo/metabase: use key with prefix as the next cursor
* 261a4c1c0 satellite/metainfo/metabase: fix iterator boundaries
* 0370d1553 storagenode/console: Fix TestStorageNodeApi flakiness
* 3a533640c Makefile: downgrade to go v1.15.7 (#4048)
* cf0c88c67 cmd/uplink: fix no result when listing single file
* 68605f32e satellite/metainfo: limit number of batched segments in metainfo loop
* 4a2c6f620 private/dbutil/pgutil: add UUIDArray helper
* 1d28dc314 private/dbutil/pgutil: move type helpers to separate file
* 4eb12eb70 web/satellite: create a project on AG name step
* 8bbabe57d cmd/uplink: allow deletion of pending objects
* f2be85686 storagenode/payouts: estimated payouts calculations upgraded
* 1af9400a2 satellite/satellitedb/dbx: remove unused methods
* e96ed289e cmd/uplink: add --pending for listing pending objects
* a0cc7e8c5 satellite/metainfo/metabase: use segment.Pieces to check if segment is inline
* badd6d12d web/satellite: distributed amount added to tables, content changed
* 300e88f9a web/satellite: config flag for satellites in beta
* f3e75c754 satellite/console: add new professional tab fields to create user function in service
* 846bb895e satellite/metainfo: loop cleanups
* c3e13b9b3 web/satellite: change info banner for US2 satellites
* 95b78e801 satellite/metainfo: select segments in batches for metainfo loop
* d995fb497 Merge remote-tracking branch 'origin/main' into multipart-upload
* 0178ec777 satellite/satellitedb: clearer testdata version parsing
* 32afbd0d1 satellite/satellitedb: faster test database setup
* a34da8531 ci: bump test timeout
* fce5c95eb web/satellite: logic for beta satellites terms confirmation
* 21a6dccd3 ci: set specific line coverage targets
* bc6889884 ci: enable coverage for multipart-upload branch
* 549033f2e satellite/satellitedb: don't include DQd and exited nodes in DQStrayNodes
* 3c49f0ff7 cmd/satellite: restore-trash, cleanup workers
* 79ecd80c0 satellite/metabase: Remove subquery on delete objects for CRDB
* 630718c39 satellite/gracefulexit: fix test DeleteAllFinishedTransferQueueItems
* 62a2d05ae satellite/accounting: try fix test ResetLimitsFirstDayOfNextMonth
* 1137620ba satellite/satellitedb: move tests to their domains
* e23f42310 uplink: add api key back to inspect
* e0f15130a satellite/metainfo: ensure list prefix is prepended in satStreamID
* 8b9da0181 satellite/metainfo/metabase: allow committing empty inline segments
* 4c0ea717e satellite/metainfo: remove unneeded dependencies from Loop
* a25e35f0b cmd/metainfo-loop-benchmark: add benchmark
* 1cb6376da satellite/metainfo: remove BucketsDB.ListAllBuckets
* 8e95e76c3 satellite/metainfo: fix metainfo loop
* c860b74a3 satellite/repair/checker: allow for multipart objects
* ac8c70e20 cmd/satellite: restore trash connection close
* 8093c666a satellite/metrics: allow for multipart objects
* b2ed7edd3 cmd/satellite: restore-trash parallel workers
* 69151965b cmd/satellite: add logic to send RestoreTrash to nodes
* 3ae3389dd cmd/satellite: restore-trash command
* adf687aeb satellite/metainfo/metabase: add full iterator
* 50239e66f web/storagenode: zksync receipt convertation
* 212baf614 web/storagenode: balance section markup, current month estimated payout frontend implementation
* 3eb321fd3 storagenode/console/consoleapi: fix flaky test
* 9e7e753fa sn/sndb/payouts: migrate distributed == paid for periods < 2020-12
* f969e2682 Jenkinsfile, Makefile: bump to go v1.16 (#4041)
* dcf7b884c storagenode/payouts: estimated payouts calculations fixed
* dd9ad0930 satellite/metainfo: return RS value with DownloadSegment response
* f1169ad7b satellite/metainfo: return RS value with DownloadSegment response
* 61f0fb67a satellite/metainfo/metabase: refresh alias cache only once
* 259b030b3 satellite/metabase: add tests for BeginObject
* 12402eb72 Merge remote-tracking branch 'origin/main' into multipart-upload
* 28433327c web/satellite: remove upload step from create AG flow, close webworker on logout, styling fixes
* 6ebe06cd1 satellite/metainfo: fix for getting real RS values
* 8685c6c9c private/testplanet: use better method in test
* 5dd76522a gracefulexit: use GetSegmentByLocation instead of GetObjectLatestVersion
* f7ad86521 satellite/overlay: fix data race in TestAuditHistoryBasic
* 79d6294db metainfo-migration: fix data race
* ed1caede9 web: size package fixed
* 4e3ac2c49 satellite/console: add new professional fields to console api
* 9bcffa3d0 cmd/metainfo-verifier: verifies if metainfo migration is done correctly
* 4b2e46a0c satellite/satellitedb: add employee size column to users
* b8513e2fc satellite/metainfo: get encryption from StreamMeta for old uplinks
* fe5b6e172 sat/console/consolewasm: restrict enc access to api key
* 341a4b7f5 metainfo-migration: extend test case
* 2e1455bc5 multinode/multinodedb: add sqlite3 support
* 5a23d9c16 metainfo-migrator: improve performance
* 2848bf488 satellite/metainfo/metabase: use alias pieces in segments table
* 25f81f353 satellite/metainfo/metabase: add AliasPiece
* 63c7f8b7f satellitedb/satellitedbtest: creating a database shouldn't auto-migrate
* 908a96ae3 Merge remote-tracking branch 'origin/main' into multipart-upload
* dc2bec9f8 satellite/metainfo: collect uplink versions with monkit
* 9cfaba2c5 private/tagsql: close Stmt debug tracker
* a86bd2751 storagenode/storagenodedb: paystubs distributed not null added
* 4661bc9b8 web/storagenode: notification createdAt bug fixed
* 49884b5aa multinode/console: payouts added
* 1ac17cfac uplink: link.tardigradeshare.io URLs should have an /s/ right after the domain
* d9d6edde7 go.mod: bump monkit-jaeger to replace dashes by underscore
* 4a797baa7 satellite/repair/repairer: a new set of rs_scheme tagged metrics
* 520da5b1e scripts/tests/rollingupgrade: fix rollingupgrade test
* babe01f22 scripts/tests/rollingupgrade: fix rollingupgrade test backwards compatibility
* 1e328f3c3 satellite/compensation: add wallet features to invoice csv
* 966535e9d {storagenode,satellite}/nodeoperator: add wallet features
* 570dcfad3 web/satellite: add segment event for Generate Gateway Credentials button
* fcd18ef48 web/satellite: rework signup verify redirect flow
* fd095e604 web/satellite: fix for access grants creation datepicker
* 89529237a web/satellite, web/marketing: remove referral program related UI code
* 3b49d3cdd satellite: remove referral program related code
* dc6a8cbcf storagenode/suspension: fixed timings of downtime suspension notifications
* e86041a5c web: stylint config moved from package.json
* cd5760620 web/storagenode: size package refactored
* 0c223f288 satellite/metainfo/metabase: don't use random node ID-s
* d0a7ab60e satellite/metainfo/metabase: add NodeAliasCache
* 3cc1a0b39 dbutil/cockroachutil: make crdb retry when tcp error
* 0260966de satellite/metainfo/metabase: add NodeAliasMap
* 8b49b210a Projects list view (#4035)
* 0ef696a15 cla: remove .clabot
* 12fde343c web/multinode: delete node option added to context menu
* 322c3a167 satellite/metainfo/metabase: add node alias table
* e114cfe86 satellite/satellitedb: fix broken query
* f011a5e8c Use WithNonce when creating caveats
* cdf58a1ff multinode/payouts: EarnedPerSatellite added private/multinodepb: multinode.proto updated
* db3a3088f satellite/satellitedb: add professional user fields to db interface (#4034)
* 615586a47 {satellite,storagnode}/internalpb: bump gogo to v1.3.2
* b92a86e52 storagenode/payouts: earned amount at satellite and list of paying satellites IDs added to DB
* 9a6001177 Merge remote-tracking branch 'origin/main' into multipart-upload
* 21dfd9988 bump storj.io/private to latest
* 9c9f48146 satellite/orders: Remove deprecated endpoint
* 0de3709f7 pkg/server: implement hybrid connector for tcp fallback
* 53d189d71 private/testplanet: add reconfiguration for disabling QUIC listener and TCP listener
* c8e42df2f satellitedb: reorder migrations 140-144
* 7e80badaf pkg/server,pkg/quic: accept an existing conn to create quic listener and allow disabling tcp/quic
* a754c3984 satellite/console/consolewasm: assert less about error
* 9506b67ca satellite/projectaccounting: Improve performance of ProjectAccounting.getBuckets
* f34519c96 metainfo/metabase: iterate through pending objects by key
* 8d25c4789 satellite/repair: fix comment in TestRepairExpiredSegment
* a17934cb5 satellite/satellitedb: remove reference to uptime counts
* 759bdd679 satellite/compensation: add total-paid and total-distributed to invoices
* 038bd0a4d satellite/repair/repairer: fix repair for pending objects
* 6f3d0c4ad Merge remote-tracking branch 'origin/main' into multipart-upload
* 339d1212c satellite/repair: don't remove expired segments from repair queue
* 22bc69ad6 cmd/uplink: Add more details to access inspect
* 0bbe2f1d6 satellite/metainfo: add unimplemented ListPendingObjectStreams
* e38372281 storagenode/payouts: Distributed added to paystubs
* cc0d88f9c satellite/satellitedb: Fix GE flaky test
* b519bb377 satellite/metainfo/metabase: GetSegmentByLocation
* d93944c57 satellite/orders: Delete unused methods & DB tables
* 076804eac cmd/satellite: Add command for GE data cleanup
* 1cf3d89a5 satellite/satellitedb: add distributed column and migration
* 6e2450ecf uplink cli: make uplink share --register output access keys again
* c489a70e6 storagenode/gracefulexit: omit finished exits from ListPendingExits
* 91bd4191d satellite/accounting: add rollup archiver chore
* 6f441960e satellite/metainfo/metabase: fix ParseSegmentKey to parse part numbers
* 74e293693 storagenode/gracefulexit: improve error message
* d0612199f Merge remote-tracking branch 'origin/main' into multipart-upload
* 5d895fb40 storagenode/{payouts,console}: use same time for all calculations
* ee1f67bb1 go.mod: bump quic to fix Go 1.16rc1
* 8a3db08f6 storagenode/payouts/estimatedpayouts: fix calculations
* c5ecca1e1 web/storagenode: remove uptime columns and references
* b7a073921 satellite/overlay: use DownloadSelectionCache for getting node IPs
* 54e01d37f satellite/overlay: add DownloadSelectionCache
* 89e682b4d satellite/repair/checker: add 29/80/130-52 to default repair overrides
* 52d6852e5 pkg/server: add retry logic for random port assignment
* fff10b041 quick update to readme (#4033)
* 19e3dc4ec satellite/overlay: rename NodeSelectionCache to UploadSelectionCache
* a97b5c855 go.mod: bump quic to master, such it compiles with Go 1.16
* 4d32bdaef satellite/satellitedb: drop bucket_metainfos_name_project_id_key index
* 824fd6f13 pkg/quic: add backward compatibility for qtls
* ca86820b8 satellite/snopayouts: use dbx + some refactorings
* 66e15fb7f satellite/compensation: remove ytd paid amounts
* 02845e7b8 pkg/server,private/testplanet: start to listen on quic
* 0f1961b7d satellite/metainfo/metabase: improve subquery comparision performance
* f18cb2452 .clabot: add harrymaurya05 and gregoirevda (#4032)
* da68f0cda cmd/metainfo-migration: add fastpb for faster unmarshaling
* d2148edcd cmd/uplink/cmd/setup.go: Access Grant vs API Key error messaging
* 93e841361 web/mnd: my nodes page ui
* 9820145e1 web/mnd: nodes domain, api and store
* 135d846af mnd/console/server: endpoint with index.html added
* 24d60384c satellite/satellitedb: add columns for professional users (#4028)
* c92bda7e7 tally monkit: change location to monitor piecesize
* 0b2568d71 satellite/overlay/straynodes: increase development duration without contact
* 3fc0d2a83 satellite/metainfo: add testing method from multipart-upload branch
* 50bea3ab1 satellite/metainfo: adjust tests to changes in uplink
* 8263f1832 satellite/console: Add graphql query for owned projects
* c139cbd76 storagenode/payouts: fix CurrentMonthExpectations timezone handling. Estimations based on node's join date.
* 49c8e9448 scripts: Add test Satellite working w/o Redis
* a700a1bda satellite/metainfo/metabase: add benchmark
* ba2907be1 satellite/metainfo: do not expire satStreamID
* b8fd8c775 satellite/metainfo/metabase: fix FixedSegmentSize for gapped parts
* 2ce829a8b satellite/metainfo/metabase: fix segment index in tests
* c44368c48 cmd/storj-sim: fix port assignment for multiple satellites
* ae03a0540 pkg/quic: add quic implementation
* b7d8dee5e satellite/console/wasm: add js tests
* 3c13aae61 satellite/metainfo: remove unused method
* 1f1d9fce5 satellite/console/wasm: add test to confirm wasm size isnt growing
* 6c3b7a233 web/satellite: add beta messaging for EU and US2 sats
* fa2dae0ac Pin Project Owner in Team tab to top of members list (#4021)
* ce675b770 satellite/metainfo/metabase: remove sql transaction while committing segments
* ff102d6ff Makefile: update to go v1.15.7 (#4026)
* 292caa904 satellite/metainfo/metabase: allow to set Encryption while committing object
* 2d087c54b satellite/metainfo/metabase: use PlainSize for fixed segment calculation
* 0a4807185 satellite/console: Add pagination fields for ListProjectsByOwnerID
* e6c038347 adding hypernet to the list (#4025)
* c24ada711 Merge remote-tracking branch 'origin/main' into multipart-upload
* d14607a5f satellite/{contact,nodestats,overlay,satellitedb}: remove references to total_uptime_count and uptime_success_count columns
* 1e0d57bf0 uplink: Add timeout to RegisterAccess
* 75d828200 private,satellite: add chore to dq stray nodes
* ea48322dd satellite/metainfo: use deterministic signing for satStreamID
* d54ae9f10 satellite/metainfo: close project in TestObjectOverrideOnUpload
* 2e34b631b cmd/satellite: Allow core & API without live accounting cache
* 0f70574e3 satellite/accounting/tally: Don't abort tally when cache is down
* 8e47adfad satellite/metainfo/metabase: optimize a query
* 6ba8f6c8a storanode, satellite: payout renamed to payouts, expected estimation payouts added, console api for audits reworked
* 1ed722752 versioncontrol: add process url resolver
* 678b07b31 satellite: Fix typos & code formatting
* c24f84914 satellite/console: Add ability to list projects by owner ID
* 424d2787e go.mod: bump deps to uplink v1.4.5
* 38beecc7a satellite/metainfo: Override pending object on upload
* 95320912b satellite/metainfo: simplify metainfo delete test
* 6dff40f5c Merge remote-tracking branch 'origin/main' into multipart-upload
* 85fb964af satellite/{metainfo,overlay}: improvements to GetObjectIPs
* 889d2eaae web/multinode: svg type error on build fix
* d11c2b709 go.mod: bump storj.io/common
* a4d06b9b1 satellite/metainfo: Don't response errors when Redis down
* a73c59bbd satellite/console/consoleweb: Change status codes usage limits
* a3b1059fe multinode/console: list node updated with total earned
* 0184d33e9 satellite/satellitedb: set default 0 on uptime columns
* ce2661664 satellite/accounting/live: Use Redis client directly
* ac058e5ec metainfo-migration: basic pointerdb->metabase migrator
* 1ad69b9f9 satellite: make external address calc more robust
* 0403e99a5 satellite/{overlay,satellitedb}: remove unused methods for old downtime tracking
* cf56071bb storagenode/payout: error message from empty receipt removed
* 0df0e980a satellite/metainfo/metabase: fix FixedSegmentSize calculation
* 24833465e satellite/metainfo/metabase: avoid magic constant
* ec88d21a3 Merge 'main' branch.
* 71ca18134 satellite/metainfo: reject downloading a segment from multipart object
* 1709117b0 satellite/console/wasm: add more unit tests
* 2d2359667 satellite/orders: remove unused satelliteAddress field
* ba5461562 satellite/orders: remove satellite address
* 51731db12 satellite/orders: use smaller encrypted metadata
* 6507f3ebc multinode/console: trusted satellites list api
* db6b64c2b Makefile: Custom gateway location for storj-sim
* da316c27a satellite/metainfo/metabase: disable PlainSize validation
* 35c61bd0e satellite: mark gateway credentials feature beta
* 8907180e8 satellite: pass contact.external-address config to web ui
* 4fc61f7af satellite/metainfo/metabase: allow commmitting objects > 2 GiB
* c0bc2123f uplink/cmd:  fix uplink access register ability to use saved accesses
* 590bc529d cmd/storj-sim: fix local binary lookup
* b33d7a318 cmd/storj-sim: add -failfast
* 1e44be290 go.mod: bump deps to uplink v1.4.4
* 2ad560e2d web/satellite: satellite address field was added to continue in CLI step
* 6f065a833 web/satellite: fix for all bucket names selecting
* f925d99d3 web/multinode: MyNodes page markup
* 857155f5a web/multinode: navigation panel added
* 69a491aa6 web/multinode: add new node modal added
* 5a43c86b6 multinode/console: list node satellite infos
* fb00d099c multinode/console: list node infos
* 16a8b1da3 web/multinode: onboarding pages and common components
* 7703b4b69 web/multinode: common dropdown component added
* 9cb4466eb cmd/storj-sim: use dev setup by default for consistency
* 710b86849 storagenode/{console, reputation}: Add audit history to storagenode API
* a90d6fcad satellite/repair/checker: Use segment health on checker insert
* 6e2ef3b9e Revert "satellite/satellitedb: Do not consider nodes with offline_suspended as reputable."
* ad3e3a38c Merge 'main' branch
* d4ebdba48 satellite/payments/stripecoinpayments: fix tests failing in 2021
* 39260b6fd storagenode/storagenodedb: Add audit_history to storagenode reputationdb
* edbee5388 satellite,storagenode: Pass audit history over GetStats endpoint
* 8b2e4bfa7 satellite/metainfo/piecedeletion Remove spaces from metrics.
* 7246368ca satellite/repair: clamp totalNodes to 100 or higher
* 641f01ed4 cmd/uplink: Actually overwrite accesses in uplink import
* f65a1e8c1 uplink: use output formats for 'uplink access register' instead of writing to AWS file
* 825dc7122 satellite/{overlay, satellitedb}: Refactor audit history
* 85ae13f11 satellite/satellitedb: Drop nodes_offline_times table.
* e24262c2c satellite/satellitedb: Do not consider nodes with offline_suspended as reputable.
* ca12e98d5 docs: change 'offline score' to 'online score' in downtime tracking doc
* 02d7638ed Jenkinsfile.public: specify main branch for docker build
* 7f1871b8f all: switch from master to main
* 222387743 storgenode/piecestore: Log upload size on Upload
* 73fb54fed web/multinode/app/store: fixed nodes mutation
* b0258f085 web/multinode/app: initial state design
* 6e68e8e00 web/storagenode: held progress typo fix
* ad5845919 satellite/admin: allow more than just "paid" invoice status during user deletion
* 7aea6feee web/satellite: access grants delete popup blinking bug fixed
* 95053531f web/satellite: access grant: result step back click bug fixed, small styling fixes
* f787aa464 go.mod: bump uplink to v1.4.3
* 7faaeed2b satellite/access grant wizard: don't hardcode the satellites
* efde103db accounting: rollup test is broken for the hour before midnight UTC
* 607001802 satellite/overlay: use AS OF SYSTEM TIME with Cockroach
* 7cccbdb76 cmd/uplink/share: add -- to flags referenced in help statements
* 1b0424cad uplink/cmd:  Export a RegisterAccess method.
* 9a8959d42 Merge 'master' branch
* 5934969dd satellite/orders: remove obsolete CreateDeleteOrderLimits method
* 777f6e583 satellite/metainfo: replace a call to PointerDB with call to Metabase
* c4578eb3e satellite/gc: add test for pending object
* 66d4d5eb4 satellite/metainfo/metabase: implement IterateObjectsAllVersions for pending/committed objects
* 876e1be3b Blueprint: Sparse Order Storage
* 563197c62 satellite/overlay: Add index on nodes table (#4012)
* dad8360b3 satellite/metainfo/metabase: rename IterateObjectsAllVersions to IterateObjectsAllVersionsWithStatus
* c78fdf472 web/multinode: base router implemented
* 9b5228357 satellite/accounting: Add index for project_id on bucket_storage_tallies (#4010)
* bafc6af99 ci: remove workaround for failing tests
* 6e501898c satellite/accounting: Performance improvements to getNodeIds used by GetBandwidthSince (#4009)
* d14f2e416 Makefile,scripts: move from aarch64 to arm64v8 (#4008)
* 529bae305 satellite/accounting: fix TestBilling_AuditRepairTraffic
* f3ef8088e satellite/metainfo/metabase: add Verify method for Pieces
* 18825d1e0 satellite/{metainfo,gracefulexit}: fix failing tests
* 4d37d1492 satellite/{metrics,repair}: adjust monitoring to new metainfo loop
* 866ce478b build: update node to v14.15.3 (#4007)
* 7eab85903 cmd: ensure proper arch is used for docker container
* d96143788 satellite/orders: remove the config IncludeEncryptedMetadata
* da0327c9b satellite/dbcleanup: remove expired serial chore
* 97a5e6c81 satellite/orders: stop inserting/reading from serial_numbers table
* aaa4a9f31 web/storagenode: overused disk space added to chart
* 4e1983966 satellite/accounting: fix TestProjectUsage_ tests
* f645654d2 web/satellite: access grant CLI step copy token fix
* 65b22be41 satellite/accounting/tally: use metabase
* f7a31308d satellite/repair: enable TestRemoveExpiredSegmentFromQueue test
* b3aa28cc0 satellite/gracefulexit: migrate to metabase
* 311b08283 satellite/metainfo: fix metainfo loop
* ce20db9f6 scripts/testdata: update satellite-config.yaml.lock
* 211174023 Merge 'master' branch
* e1456bc53 Makefile: build wasm files only once (#4005)
* 2437d5b17 satellite/access-grants: default auth service url (#4002)
* d3604a5e9 satellite/repair: use survivability model for segment health
* 3feee9f4f satellite/accounting: default project limits (#4001)
* 28eaae66a satellite/satellitedb: drop num_healthy_pieces column from injuredsegments
* 79a3a4780 build: added brotli compression for wasm bits
* f4bbd0f5d web/satellite: use brotli instead of gzip
* 50dd9fb11 web/satellite: move access grant web worker initialization to onlogin loading state
* b1712cc93 cmd/storj-sim: update default storj-sim access with real node id
* 9ea147d23 web/multinode: initial app and configs
* a3c2711b2 mnd/nodes: db interface methods updated
* 26e65eeef multinode/console: node list and get api
* 70ba4deea satellite/repair/checker: adjust irreparable part of repair checker
* cd2cdb616 web/satellite: create access grant: restricted key for CLI step, CLI step added to onboarding tour
* 9aa61245d satellite/audits: migrate to metabase
* 2381ca281 Merge 'master' branch
* c705237be satellite/inspector: migrate to metabase
* 2fd7809e5 storagenode/payout: stefanbenten satellite name added to payout history, satellites with no held history removed from list
* 8d3ea9c25 satellite/repair/repairer: implement SegmentRepairer with metabase
* 6206aa88e {satellite/metainfo,private/testplanet} use TestingAllSegments in tests
* 0e8323370 storj-sim: add node id to default access
* 1728e45e5 satellite/metainfo/metabase: optimize DeleteBucket
* 12055e786 all: minor cleanups
* 724b0f91e satellite/gc: update tests to use metabase
* 4706f0187 satellite/metainfo: add TestingAll{Segments,Objects}
* 8c52bb3a1 satellite/checker: use numHealthy as segment health in repair queue
* 9db28b92e satellite/console/wasm: fix issue from merge
* 187680b0c scripts: Fix typo in a comment
* 2ac72eaf1 satellite/repair/checker: add new monkit stats tagged with rs scheme
* 4fba9921f satellite/metainfo/metabase: define ErrSegmentNotFound error class
* 934ae32ca satellite/repair/checker: fix checker tests
* 57f374af2 Merge 'master' branch
* adb2c83e0 cmd/uplink: adds register, url, and dns flags to uplink share and replaces access grant with access
* 9fe477899 satellite/satellitedb: add lint ignore rule to support staticcheck 2020.2
* 12144a600 storagenode/console: payout tests and heldhistory joined_at rounding added
* c630037f3 multinodepb: diskspace, reputation and status added
* 3cc98de3e satellite/console/wasm: reduce size to <9MB
* d955946f1 satellite/compensation: don't abort entirely if a node isn't found
* 2dddcffe4 satellite/accounting/rollout: Remove unused variable
* 4a11ec282 multinode/nodes: package created, api tests added, small restructuring
* fc85179a1 satellite/metainfo: refactor SegmentLocation.Index to SegmentPosition
* 7d8f19e94 satellite/metainfo: metainfo loop should yield StreamID for segments
* e7e6985ae satellite/metainfo/metabase: add UpdateSegmentPieces method
* 2bb010e7c cmd: remove segment reaper
* 7e6e0d3e2 satellite/metainfo: metainfo loop implementation with metabase
* ca1e6b975 Adding Fastly (#3994)
* d559405ae Jenkinsfile: update to go v1.15.6 (#3993)
* 8fe829d5f build: add wasm bits to Dockerfile and bump to go v1.15.6 (#3992)
* 46e26fa47 Changed ChecksArea.vue audits/checks precision (#3987)
* 982c72913 .clabot: adding Doom4535 to approved list (#3989)
* 0649d2b93 satellite/repair: improve contention for injuredsegments table on CRDB
* b3acc1101 Merge 'master' branch
* c2a97aeb1 satellite/satellitedb: add ListAllBuckets method
* 588dff6d7 Upgrade to uplink v1.4.2
* 15add3c75 satellite/metainfo: enable deletion of a pending object.
* e322eac05 added myself to the clabot list. (#3988)
* 365410d10 satellite/metainfo/metabase: add commit with specifying a list of segments
* f077564bb satellite/metainfo: setup access to metabase in metainfo loop
* 5b6fcfb72 cmd/storj-sim: fix 32bit code
* 3c77825b1 cmd: move ca-certificates build step to be arch independent
* 1daadc11e web/satellite: dashboard lag message tooltip (#3982)
* 2f62cdf49 storagenode/console: diskSpaceInfo extended with overused diskspace, getDashboardData updated.
* d033ec3df private/testplanet: fix TestDownloadWithSomeNodesOffline
* 3f5e07fb1 satellite/metainfo: more tests fixed
* c8e8fea05 web/satellite: replacing api keys with access grants
* 70795c68f web/satellite: update onboarding flow: create access grant
* 55b495faa web/satellite: update onboarding flow: payment step update
* 701474acd web/satellite: update onboarding tour: overview step
* d5f1245cc web/satellite: update onboarding flow: removed redundant steps, routing applied
* aaec8bfc5 web/satellite: update onboarding tour: routing
* d0d0a192c satellite/metainfo: fix TestEndpoint_DeleteObjectPieces tests
* fb5244e8f satellite/metainfo: migrate GetObjectIP endpoint method to metabase
* a71c908c1 cmd: add ca-certificates to Docker images (#3986)
* bc3be9c41 Jenkinsfile: enable build failure on integration tests failure
* 326940151 deps: update to latest uplink from multipart-upload branch
* 218bbeaff Merge 'master' branch
* cdeea1c99 private/testplanet: add helper OpenProject method to testplanet uplink
* 7eccfccfd web/satellite: create access grant: bucket search feature
* 494bd5db8 all: golangci-lint v1.33.0 fixes (#3985)
* 746315672 scripts/tests/testversions: fix indentation
* 0faf7d529 scripts/tests/testversions: fix race in install_sim
* c1bdb88d3 Upgrade to uplink v1.4.1
* cc9e9ee1f storj-sim: use gateway from multipart-upload branch
* f90ea10a4 Allow for DB application names per process. (#3983)
* d75e4be11 satellite/{accounting, contact}: Remove periods and spaces from metrics.
* 13555f398 scripts/tests/testversions: enable concurrent installation for each version
* a7685f50c satellite/metainfo/metabase: set maxParts to MaxListLimit if greater
* fd7cc20a7 satellite/metainfo/metabase: add DeletePendingObject
* 3fc76f4ff satellite/downtime: Remove deprecated downtime tracking service.
* 6ba8dd58e Update to uplink v1.4.0
* bc25cc8dd satellite/metainfo: add encryption info in metadata of list items
* 83e7cd2a4 satellite/metainfo/metabase: replace SQL substring with go code
* 47e008b71 cmd/uplink:  Write to AWS Credentials file more safely from uplink access register
* 1728c3a99 satellite/dbx: standardize on assignment
* b26111035 satellite/orders: get bucketID from encrypted metadata in order instead of serial_numbers table
* e8378a87b satellite/metainfo: include redundancy in satStreamID of ListObjects
* d5c026416 satellite/metainfo: implement deleteBucketObjects with metabase objects iterator
* f08e34f15 satellite/metainfo: fix listing objects when cursor is specified
* e4c4ab06b satellite/metainfo: add testing methods for getting objects and segments
* 46839b992 satellite/metainfo: expired deletion service to use Metabase
* 70b91aac5 satellitedb: remove cruft caused by https://review.dev.storj.io/c/storj/storj/+/3223
* d8ba7b305 satellite/console: only allow project member to get all bucket names
* 76199db3c private/testplanet: expose Metabase to Test Planet.
* 65919f9f7 cmd/uplink:  add --aws-profile flag to uplink access register
* 5a7bc9657 Merge 'master' branch
* f456d7ce0 satellite: remove implementation detail from DB interface
* 28ea63be9 satellite/repair: avoid TestDBAccess
* 0771cdb0b web/satellite: create access grant: generate gateway credentials step
* bb7677a85 web/satellite: get gateway credentials request using url from config
* d3d6e0c67 web/satellite: create access grant: set duration restrictions
* 21602e049 satellite/metainfo: enable commented test
* 71e11b27f satellite/dbx: only retry with cockroach
* bd23d12bb satellite/dbx: add cockroach retries for other QueryContext operations
* ea2f39ca7 satellite/dbx: add retries for QueryRowContext-based operations
* d3b0691bb satellite/dbx: import dbx templates
* 5d8a67a4f satellitedb: retry GetBandwidthSince on cockroach
* 5dc013d3b satellite/overlay: Add retry to all selects in overlaycache
* 6bce907cb satellite: try to stream rollups to aggregation function to use less memory
* 6aae21541 satellitedb: do saverollup in batches
* 0ba516d40 satellite: support pointing db components at different databases
* 75f0f713a satellite/repair/checker/checker.go: Use number of healthy pieces instead of SegmentHealth for injured segments queue.
* 5c34b62bd Fixed typos in downtime tracking with audits doc (#3977)
* 7eb3b2d6d satellite/gc: Init map with an aprox size
* 319d2cad1 satellite: Fix typo in a comment
* cfb45a785 satellite/metainfo/metabase: add TotalPlainSize to objects table
* 8ceef9f35 satellite/metainfo: temporary disable one assertion in test
* 983b1737c satellite/metainfo: implement IsBucketEmpty with metabase
* 3792e2921 satellite/accounting/tally: make test less fragile
* c6626748d web/satellite: create acces grant: duration selection logic (#3978)
* 46102c194 satellite/metainfo: change ListSegments required permission to Read/Download
* dad36179c satellite/metainfo/metabase: fix segment listing with cursor.Part
* 257c8682d web/satellite: create access grant: result step
* 84fb8eee1 web/satellite: create access grant: generate access grant in web worker
* 67c210a64 satellite/metainfo: set list cursor to version 1
* 53b7fd7b0 satellite/{audit,gracefulexit}: remove logic for PieceHashesVerified
* efaba85c7 Merge 'master' branch
* 9de1617db satellite/orders: ensure encryption keys handles set twice
* 575f50df8 satellite/repair: Update repair override config to support multiple RS schemes.
* 55d5e1fd7 satellite/orders: ensure that expired deletion doesn't stall
* 5beb2f573 satellite/orders: add factory function to encryption key
* 2b92bba56 satellite/satellitedb/orders: Handle serial_numbers deletes in smaller increments on CRDB
* 5be4d413b satellite/metainfo: set PlainSize for inline segment
* 2a981b86d web/satellite: create access grant: enter passphrase step
* 4dddb6e66 web/satellite: create access grant: create passphrase step
* 005b089c0 web/satellite: create access grant: get restricted api key from web worker
* c22ae05bb web/satellite:access grant date picker (#3972)
* 05fe497a6 satellite/metainfo: add stream id when listing pending objects
* 54ae9b040 cmd/uplink: allow public access registration
* a8b66dce1 satellite/accounting: account for old orders that can be submitted in satellite rollup
* e19fabc88 scripts/tests/rollingupgrade: fix typo in flag
* 5fe3d2dea cmd/uplink  Allow use of named accesses in uplink register
* 146553922 satellite/metainfo: ignore empty inline segments for backward compatibility
* aeb801604 {satellite,storagenode}/orders: fix flaky tests
* 70b1c7aa1 satellite/metainfo: check if EncryptedMetadataNonce was set
* 41d86c098 storagenode/orders/ordersfile: Add reasonable size caps for orders/limits when detecting file corruption.
* 8da5e6a55 scripts/tests/rollingupgrade: use wait-for instead of sleep
* a30f5d7ec satellite/metainfo: use Object.Status from ObjectListRequest
* a17cd9aa3 storageode/apikey: added service, CLI issue api key
* 2b59640f1 cmd/satellite: ignore Canceled in exit from repair worker
* f5a5308cc web/satellite: upload data step
* 1c13065b0 web/satellite: create access grant: continue in CLI step
* 6f35ee98e web/satellite: create access grant: bucket names selection logic
* 4e49b00c6 web/satellite: create access grant permission step, regular permissions, buckets dropdown
* a332b3d81 web/satellite: create access grant name step
* 278e29c1c web/satellite: create access grant progress bar
* 6517315ff web/satellite: create access grant base container
* 6664a129b web/satellite: add all needed methods to access grant webworker
* b60939e48 web/satellite: delete access grant flow
* e16f02b70 web/satellite: access grant list page
* 9bdc8ebf7 satellite/metainfo: adjust GetObject to handle EncryptedMetadataEncryptedKey
* 72fed3b3f satellite/metainfo/metabase: add slow non-recursive listing
* fa95c6bbb storagenode/orders/ordersfile: Fix error message wrong var
* 5e7d47b9c satellite/metainfo: code cleaning
* 2e426ef0a satellite/metainfo/metabase: iterator compatibility
* 402a581b8 ci: fail build if metabase tests fail
* c9bbd83f9 satellite/metainfo/metabase: align object status to those already in use
* ac9433342 web/storagenode: notifications components unit tests
* 9740da650 storagenode/orders: Don't panic if size is over MaxInt32
* b7df41a0a satellite/metainfo/metabase: add EncryptedMetainfoEncryptedKey parameter to metabase
* c409194d4 storagenode/payouts: estimation payout heldamount rounding removed
* 0ec685b17 satellite/{satellitedb, repair/{queue, checker}}: Use new column "segmentHealth" instead of "numHealthy" in injured segments queue
* afc9545ff cmd/storj-sim: add "tool wait-for <address>"
* 9ab824d3e jenkins/rollingupgrade: sleep 5 seconds between old api startup and database migration (#3971)
* 48d8114b3 satellite/contact: treat pingback failure as error
* f31172285 multinode/db: nodes repository tests added
* 51fa52e63 web/satellite: access grant type, api, store module, mock
* 51a712f9e satellite/console: get all bucket names endpoint and service method
* 402cfcb7c satellite/metainfo/metabase: add prefix to iterator
* db480e6e1 storagenode/orders: Improve performance of handling corrupt orders.
* 7aba265db satellite/metainfo/metabase: add status option to iteration
* a015f4192 satellite/metainfo: ListObjects to use Metabase API
* 1726b39ed scripts: remove thrift mod replace
* 7c384c829 Merge 'master' branch
* cd7d911b2 satellite/metainfo: implement ListSegments with metabase
* c2ba5a990 Makefile: Update Go version security patch
* f558cc825 satellite/orders: add storagenode_bw_phase2 table and dont delete tallies for longer
* c0b5e7ce3 ci: ensure cockroach doesn't pollute repo
* 2e6ffd9af web/satellite:access grant empty state (#3970)
* ff3226831 satellite/metainfo/metabase: DeleteExpiredObjects
* a749ac9f4 satellite/metainfo/metabase: iterateObjectsAllVersions
* 1b4bfbb9d multinode/console: nodes addition and removal
* e6967720c cmd/multinode: create schema command added, run command bug fixed
* f8d3a977f web/storagenode: PayoutPeriodCalendar.vue unit tests
* 226e13e61 satellite/cosole: add tests for wasm access code
* 259f4ebcf web/storagenode: EstimationArea.vue unit tests
* 8182d8a72 satellite/metainfo/metabase: use txutil.WithTx in delete
* bc460cd62 satellite/metainfo/metabase: use txutil.WithTx in commit
* 9d25b3a7d satellite/metainfo/metabase: add ListSegments request
* 54c5d564a scripts/tests/testversions: fix older uplink setup
* 3e5640359 satellite/repair: add a repair health function
* fd707392e satellite/metainfo: set RedundancyScheme for begin segment response
* 31533ed1a satellite/console/wasm: remove storj.io/uplink deependency
* 592d0bd6b web/satellite: access grant routing (#3966)
* 5a337c48e {cmd,private,storagenode}: create storage dir verification during setup
* 4ce00c7ca cmd/multinode: run and setup commands added
* 5e0106f1f web/satellite: web worker for wasm
* 9dcd0936d satellite/metainfo: MakeInlineSegment to use metabase
* 07acf0e57 cmd/storagenode: add docker env variable to toggle running setup
* da9f1f061 satellite/repair: add monkit counter for segments below minimum required
* 3fe16f400 satellite/metainfo: upload/download with metabase
* 2ff7925e6 ci: set GOTRACEBACK=all
* d4d43f02b cmd/internal/wizard: select satellite
* e32466696 cmd/uplink: add access register command
* 8fd841b91 scripts/tests/testversions: fix installation during setup
* 2ce3170bb satellite/console/wasm: expose method to add caveats in the browser
* 5f840aae6 build: use go v1.15.4 (#3968)
* 2bd239bb7 satellite/metainfo/metabase: tests commit inline segment
* dbaf8b6c0 web/storagenode: sno types/store/api refactoring
* 48f1f7132 satellite/metainfo: set RS defaults
* 3ba52b25a satellite/rewards: update partners to include MAXN
* 7dde184cb Merge 'master' branch
* 92f925107 satellite/metainfo/metabase: add GetSegmentByOffset request
* b892a0014 mod: bump dependencies and reenable test
* 3ed4183e5 satellite/metainfo: delete object to use metabase
* a5e05a63b .clabot: clarify the cla message (#3964)
* db6bc6503 satellite/metainfo: Update metainfo RS config to more easily support multiple RS schemes.
* dc5a5df7f chore: fix typos in the documentation (#3959)
* e6dd3ecaa multinode/database: members repository created
* dbf9f121e ci: mark failures as unstable
* 8dc10e32a stefan benten satellited added to historical payout data
* 184e1ffa9 satellite/metainfo/metabase: commit object without proofs
* d63b7658e satellite/repair: fix lastSeenSegmentKey bug in IrreparableProcess
* 1e356f1c5 web/storagenode: held progress label changed 15 month -> 16 month
* 3b388c21c web/storagenode: notifications domain and app types separations
* 074784e1b .clabot: add hectorj2f (#3963)
* 754838f5d ci: convert test failures to unstable
* e1f37ece0 private/lifecycle: warn on slow service shutdown
* cbc192259 private/dbutil/pgtest: use round robin to pick databases
* f8c3848c7 satellite/console: change user's email endpoint/feature
* 60bb34a09 private/testblobs: fix data race in BadDB
* 0a77deec8 satellite/metainfo/metabase: basic migrate
* 56e6bc884 metabase: add DeleteObjectsAllVersions method
* 2398afe98 satellite/metainfo/metabase: add GetLatestObjectLastSegment request
* c4c29e370 wasm: add webassembly code for creating access grant in console web UI
* f986fdfe7 metabase: optimize delete segments
* 2dffaebc6 satellite/accounting: Fix and enhance code doc comments
* 8616fc146 satellite/orders: send IPs for graceful exit
* c55c23f81 private/testplanet: add STORJ_TESTPLANET_ABSTIME
* 4134100d8 satellite/metainfo/metabase: add UpdateObjectMetadata request
* 89cefa904 metabase: add Objects info to DeleteObjectResult
* f7aa5b4e3 metabase: CommitObject return metabase.Object
* 7e9ac2ec0 metabase: use storj.ErrObjectNotFound instead of custom error
* dc67ce74c satellite: remove IsUp field from overlay.UpdateRequest
* 0c23b1203 private/testplanet: use relative time logging
* ae17ae73d satellite/metainfo/metabase: add metadata and nonce to CommitObject request
* 7183dca6c all: fix defers in loop
* 716068a1e Merge branch 'master'.
* fd8e697ab {satellite,storagenode}/internalpb: use specific package name
* 9674ce40e mod: bump dependencies
* 0205f0d80 satellite/metainfo: fix usage of types from internalpb
* 77c4f99fa satellite/internalpb: move delegated_repair.proto
* 11338e9be satellite/internalpb: move audithistory.pb
* 1903b1547 storagenode/internalpb: move gracefulexit.proto
* 517212bfa satellite/metainfo: fix usage of types from internalpb
* cda67a659 storagenode/internalpb: move inspector.proto
* 7ce372c68 satellite/internalpb: add inspectors
* 004e610d0 satellite/internalpb: move datarepair.pb to internal
* b8c6fb764 satellite/metainfo: add metabase to metainfo service
* 995900e02 Merge branch 'master' into multipart-upload
* 8f26f66da internalpb: move satellite specific protobuf types storj/storj
* f5ba8b800 storagenode/suspensions: added offline-suspension notificatio chore + tests
* e0dca4042 all: add pprof labels for debugger
* 53b396b90 satellite/metainfo/metabase: move Encryption to Begin
* 624255e8b storagenode/secret: db tests added, small renaming fixes added
* caefde6b3 private/{dbutil,tagsql}: pass ctx to database opening
* e3985799a storage/{cockroachkv,postgreskv}: add ctx to opening
* 809eb14ac satellite/metainfo/metabase: move metainfo PoC into storj repo
* 89ce1fe62 storagenode/storagenodedb: add ctx to OpenNew and OpenExisting
* 096445bc1 certificate/authorization: add ctx to OpenDB
* d0beaa4a8 pkg/revocation: pass ctx into opening the database
* 9b2e00a38 satellite: pass ctx into satellitedb.Open
* ed1f6d797 satellite/config: move repair override from config to default (#3958)
* cb1fea87f satellite/metainfo: mark unused methods as 'not implemented'
* 1adb497a7 satellite/metainfo: remove unused code
* 76f4619a9 {satellite,storagenode}/gracefulexit: ensure client is closed
* 99c88efbb scripts/tests: fix gateway tests
* 92a2be2ab satellite/metainfo: get away from using pb.Pointer in Metainfo Loop
* bb7be2311 satellite/{audit,overlay,satellitedb}: enable reporting offline audits
* 2fbb4095b storagenode/orders/ordersfile: Handle remaining pb.Unmarshal errors
* 9adde49e1 satellite/gracefulexit: ensure test doesn't timeout on failure
* 22ec940f7 storage/filestore: defer closing
* 2f4e38399 web/satellite: project edit page: back button click area decreased
* 53ba01b1f storagenode/orders/ordersfile/v0.go: Return ErrEntryCorrupt on pb.Unmarshal failure
* f5880f683 satellite/orders: rollout phase3 of SettlementWithWindow endpoint
* 9a29ec5b3 Add index to graceful_exit_transfer_queue table
* 4b61ca638 satellite/console/consoleweb/consoleapi: Fix & add test DeleteAccount
* 76d4977b6 storagenode/gracefulexit: logic moved from worker to service
* 9abdcc05e satellite/console/consoleweb/consoleapi: report err to monkit
* 746cbfc65 scripts/tests/rollingupgrade: test current release version on master branch
* d6b9563e5 web/satellite: disposed removed from historical gross total, total+surge calculation changed
* 46b12c96b satellite/console/consoleweb/consoleql: Fix typo
* 1f386db56 cmd/satellite: remove metainfo commands (#3955)
* 1aeb14e65 satellite/audit: do not delete expired segments
* 89bdb20a6 storagenodedb/orders: select unsent satellite with expiration
* 360ab1786 satellite/audit: use LastIPAndPort preferentially
* 25df79a6b storagenode-updater: check binary version on self-update
* 979ee762b satellite/console/consoleweb: Fix typo in method name
* 77d54ff0a storagenode/bandwidthdb: Use existing indexes (#3949)
* 334ae5b16 satellite/admin: add apikey endpoints
* c6415406a docs/blueprints: graceful exit initial refactoring (#3938)
* d3805761a web/storagenode: accesible functional elements
* 9df74338a storagenode: secret db and service added
* 7c3afe164 satellite/overlay: uncomment dq for offline and disable with feature flag
* 7c275830a web/storagenode: gross total added to historical data, with surge moved
* 205c39d40 satellite/orders: upgrade to phase 2 rollout ordersWithWindow
* 139a7ee95 private/migrate: add ablity to create dbs during migration
* aa86c0889 storagenode/console: Add current storage used per satellite to storagenode api
* 0b43b9325 satellite/satellitedb: make limits per default NULL
* 20a50f090 cmd/metric-receiver: restore minimal metrics server
* 02cbf1e72 storagenode/orders: Add V1 orders file
* 59d85aab5 web/satellite: take project amount limit from db instead of config
* be84616e6 Update to uplink v1.3.1
* 126450f7d multinode/database: nodes repository
* e2d589f3c certificate/rpcerrs: move logging sanitizer into certificate
* c54dd4575 examples: remove rarely used code
* 2f4bb114d go.mod: bump common to remove sha256-simd
* f06ce1ef0 release: remove binary stripping
* 830817ec0 cmd/storj-sim: run gateway without --access flag
* cf2dd76db cmd/satellite: proper log usage
* 2268cc1df all: fix linter complaints
* 0bdb95226 all: use keyed special comment
* 3ff846787 satellite/projectlimit: Update limit increase link (#3950)
* 1d3b72876 satellite/{console/payments/satellitedb}: add validation for deletion of account and project
* 50756cb43 Adding dominickmarino to the CLA bot list (#3952)
* 4cbd4d52a satellite/orders: only hold the orders semaphore during database calls
* 0f0faf0a9 satellite/orders: do a better job limiting concurrent requests
* cf1748158 Bump Dependencies
* 7161506b6 Makefile: handle msi packages correctly
* c1ca470e7 storagenode/orders: fix import and cleanup go.mod and go.sum
* 7c303208f satellite/satellitedb: emergency temporary order processing semaphore
* ad8da61da cmd/satellite: Remove curl from Dockerfile
* 3209effeb storagenode/orders: Increase order sending interval from 5m to 1h
* e598876d7 cmd/storagenode-updater: trimm \n suffix on receiving service pid from systemctl
* a1488e53a web/storagenode: online score added
* b3cf12f56 satellite/console: Add more validation for console requests
* e7f2ec7dd satellite/audit: fix sanity check for verify-piece-hashes command
* 4280142b2 satellite/console: remove unnecessary Error.Wrap
* fbf2c0b24 storagenode/orders: Refactor orders store
* bd177bff0 cmd/storj-sim: cleanup gateway setup
* 14a2050b8 pkg/auth: move package to consoleauth
* 44bd65795 satellite/console: ensure only project members can remove other project members
* 9deea2ffe satellite/console: disable account deletion via API
* 338ed9087 Makefile: simplify binaries-upload
* 0aaad88a4 satellite/{admin, console}: add test for projectLimit increase and update README
* f27cee91b web/storagenode: 404 page rework
* b39a99bae satellite/{overlay,satellitedb}: always show node's real online score
* 0d7ae8b18 web/storagenode: held history table: disposed subtracted from held
* 664b8f682 storagenode/payout: estimation payout values switched from int64 to float64 to avoid incorrect rounding. float64 values rounding to 2nd sign after dot.
* f13084983 bump common to latest
* 406281fc1 web/storagenode: changed restriction to not show calendar in payuot history table [SG-516]
* c085a17a5 bump common and uplink to latest
* fa39274ff multinode/multinodedb: add dbx support
* c4d562596 cmd/storagenode-updater: fix service build tag, for platforms other than linux and windows
* 3a7eb10c7 web: tests fix
* 4e8d53c8f private/dbutil/pgutil: ensure storagenode doesn't depend on pgx
* 245986d52 negative space calculations fix removed
* 80eeaf85d web/satellite: sign up link updated
* a43afa7ea make: fix storagenode-updater binary target
* a840cb71e storagenode: check db version before run
* 139ff3427 go.mod: bump dependencies
* 000b1e601 satellite/admin/project_test.go: Update TestCheckUsageLastMonthUnappliedInvoice
* c2525ba2b satellite/{repair,satellitedb}: clean up healthy segments from repair queue at end of checker iteration
* c23a8e3b8 go.mod: update pgx to v4.9.0
* cd2a5484f storagenode/console: ignore untrusted satellite while returning dashboard data and calculating satellites data
* 081fdcc83 multinode: initial infrastructure
* 8783c2024 satellite/rewards: add partner videocoin (#3946)
* 8499323ac web/satellite: project summary added to dashboard page
* b409b53f7 cmd/satellite: command for verifying piece hashes
* d508c4c98 scripts: remove duplicate check config lock
* 2d27bc878 satellite/satellitedb: separate cockroach for migration tests
* 79eb682f9 satellite/console: allow coupons to be a valid payment option
* 8786e55a7 storagenode/storagenodedb: allow existing dbs on setup
* 4a2c66fa0 satellite/accounting: add cache for getting project storage and bw limits
* 888bfaae4 cmd/satellite: only add google profiler to satellite
* ff57b30ba Raise warning to ensure early notification (#3945)
* e9ffffcaa web/satellite: nav bar colour fixes
* ca1f0cfe4 web/satellite: dropdowns' working logic moved to local store
* 870abd867 storagenode/pieces: tidying trash log
* 9d0d0ad72 satellite/console: enable multiple projects all users
* 38108828a satellite/satellitedb: enable multiple projects existing users
* 5f6fccc6e satellite/satellitedb: makes limits nullable change backwards compatible
* 8614d49e6 web/satellite: account dropdown reworked
* 7710e083d cmd/uplink: update CLI format (#3773)
* 8287e3a32 storagenode/orders/store.go: combine writeLimit/writeOrder operations
* 1e5641aa5 web/satellite: unnecessary console error on no paywall user's login removed
* 54dd43004 storagenode/pieces: fix typo for satellite id and piece id
* 2f648fd98 satellite: make limits be nullable
* 34613e458 cmd/satellite: command for fixing old-style objects
* 2bdbe2124 web/satellite: stored selected project feature
* 96ec44ff1 storagenode/pieces: make log more legible
* 92a336cb5 web/storagenode: held returned block added
* 84b6b91ee build: Go 1.15.2
* f46161cf2 consoleweb: log index template failures
* 2668ec818 web/satellite: added ability to edit project name
* 8182fdad0 storagenode: heldamount renamed to payouts, renamed some methods and structs to more meaningful names. grouped estimated payout with pathouts satellite: heldamount renamed to SNOpayouts.
* 3bbf08917 web/satellite: node packages updated
* d199daa90 satellite/metainfo: replace local constants with metabase values
* 9a627b22e web/storagenode: frontend refactoring
* 7db5794c1 storagenode/orders/store: Do not lock order enqueues for entire duration of ListUnsentBySatellite
* 7d5e0259f satellite/projects: initial update project name functionality implemented
* e7c34a053 satellite/satellitedb: add column and index "updated_at" to injuredsegments
* 528aa76ae storagenode/payouts: payoutHistoryMonthly surge reworked, empty receipt now won't return error
* 5b9fd4e50 cmd/storj-sim: fix prefix writer infinite loop bug
* 789b07e22 storagenode/orders/store.go: Do not return error from ListUnsentBySatellite when order files are corrupted.
* 1a094fd6d storj: bump storj.io/private
* df5a6ebe3 web/satellite: navigation sidebar and dashboard header reworked
* ac29d8049 storagenode: heldamount GetPaystub refactored, estimationPayouts logic separated form console to separate service, storagenodeapi tests fixed.
* 1fecaed7d satellite/orders: don't version check old endpoint
* 2cf12b940 satellite/admin: use correct period end time to get results
* 27a9d14e2 satellite/repair: use metabase.SegmentKey type in repair package
* c753d17e8 satellite/metainfo/objectdeletion: replace objectdeletion.ObjectIdentifier with metabase.ObjectLocation
* c4d6f472f web/satellite: notification bar for reaching projects count limit
* 9fd97fa97 ci: add check-monkit
* 249f98f4c audit logging: log important user activity
* 9996aeca8 build: Go 1.15 (#3927)
* 179b5adad storagenode/orders: add missing mon.Task parameter
* 2111d9068 satellite/admin: add get handler for projects
* 8b4b44d42 private/web: fix ratelimter IP handling
* b45cad5ee bump common to latest
* 4e2413a99 satellite/satellitedb: uses vetted_at field to select for reputable nodes
* cccd62fff bump uplink version to v1.3.0
* 8649a0055 satellite/gracefulexit: replace `Path []byte` to `Key metabaseSegmentKey` TransferQueueItem
* 920229534 satellite/metainfo: replace ScopedPath with metabase.SegmentLocation
* 93898146f web/storagenode: net total payout displaying fix
* 2cc3c19d5 satellite/piecedeletion: add test for disproportionate requests/nodes ratio
* aa47e70f0 satellite/metainfo: use metabase.SegmentKey with metainfo.Service
* 3f0c21323 web/satellite: create new project button added to projects dropdown (#3936)
* 77b53bd21 private/lifecycle: log fatal ending to a runner
* c1e089b22 web/satellite: create project popup replaced with separate page
* dc48197bd satellite/orders: add bucket id to order limit
* 9247883c6 web/storagenode: list of audits and suspension scores
* af773ec8a cmd/uplink: use DeleteBucketWithObjects for bucket force deletion
* 36d752e92 storagenode/reputation: offline_under_review_at added
* 7d9897b7a storagenode/nodestats: online_score added
* 5ed0038ac web/satellite: estimation table surge added
* 3fbbe1847 web/satellite: scrolling with no paywall banner fixed
* 7c7334e9d satellite/metainfo/piecedeletion: fix deadlock when waiting for threshold to reach
* 4783c3e1d storagenode-updater/linux: restart systemd service
* 61b17f121 satellite/orders: add encryption keys flag to Service
* 1f711523d satellite/repair: switch to piecestore.UploadReader part 2
* 95ff29cce satellite/metainfo: reduce lookupLimit default to 2500
* b872fe52a satellite/repair: switch to piecestore.UploadReader
* 084e280a6 satellite/admin: fix time issues in tests
* 0604a672c satellite/metainfo: use metabase in loop
* 2d5444708 satellite/metainfo: Fix GetObjectIPs loop and test
* ca0c1a5f0 storagenode/{monitor,pieces}, storage/filestore: add loop to check storage directory writability
* 5d21e8552 satellite/audit/queue: Separate audit queue into two separate structs.
* ba1a113e2 satellite/metainfo: Add test for GetObjectIPs
* 351aa70eb satellite/metainfo: Implement GetObjectIPs
* 2d01dd973 satellite/satellitedb: Add online_score column to nodes table
* 328004c0e satellite/accounting: fix build - time rounding
* e072febbc Fixed typo in log for allocated space (#3934)
* c86c732fc satellite: simplify tests
* 4645805b1 private/dbutil: set connMaxLifetime to 30 minutes
* 60a95d0dc satellite/{satellitedb,overlay}: Enable offline suspension and review period
* 9225fc5ae satellite/accounting: simplify ExceedsBandwidthUsage call
* 3ca405aa9 satellite/orders: use metabase types as arguments
* 3bfb0a524 Adding Kesque, MSP360, Innovoedge, Taloflow, Restic partners (#3933)
* d7c6ca601 satellite/metainfo/metabase: add package for metainfo database
* 086a3d534 satellite/{payments,admin}: add deletion of user creditcards on account deletion
* 729079965 satellite/satellitedb : remove migation steps 69-102
* 3e343b683 cmd/segment-reaper: add metrics for zombie segments count
* dbb53151f private/testplanet:  Decrease metainfo MaxBuckets test value to speed testing.
* 4f28bf072 satellite/audit: Do not return errors from Verify or Reverify on segment modified, expired, or deleted
* c4a4745dd storagenode/console: audit per satellite now uses satelliteName instead of satelliteID
* f16cf5ccc storagenode/console & /inspector: added recalculation of disk space info
* 88ff8829a satellite/gracefulexit: RecvTimeout increased to 2h, so slow nodes stop receiving lot of fails and as a result DQ
* bd5213f68 satellite/metainfo: implement batch delete for DeleteBucket
* f0ef01de5 storagenode/gracefulexit: retry workers faster
* e6bea4108 Revert "gracefulexit: reconnect added"
* cff44fbd1 gracefulexit: reconnect added
* 68b67c83a storagenode/{orders,piecestore}: Always unlock unsent orders file, even with an empty order.
* 5729d087b web/satellite: dashboard template simplified, project selection moved to nav bar
* e5012fcb3 web/satellite: info bars for accounts with no paywall
* db57d76ee storagenode/gracefulexit: fix wrong error handling for corrupted pieces (#3930)
* 959cd5cd8 satellite/satellitedb: Update audit history from overlay.UpdateStats and overlay.BatchUpdateStats
* 5f0477ebe satellite/{overlay,satellitedb}: Create database functionality for updating audit history
* 91698207c storagenode: live tracking of order window usage
* 0155c21b4 private/testplanet, storagenode/{monitor,pieces}: write storage dir verification file on run and verify on loop
* 586e6f2f1 private/testblobs, storage, storage/filestore: add storage dir verification to filestore
* 14ad7a4f1 satellite/metainfo: add limiter for objectdeletion and piecedeletion services
* 708cb48aa storagenode/orders: implement orders filestore on storagenode
* 5445d595c storagenode/gracefulexit: Wait for the worker delete and transfer goroutines to finish before completing the exit
* 7f8df7407 private/testplanet: Use config with name set when empty
* be3fd0147 storagenode/storagenodedb: database name in all preflight errors
* b4c8e219c satellite/orders: calculate order expiration inside signer
* 189ab0784 satellite/orders: use Signer in CreateGetOrderLimits
* cd5e99ea6 satellite/orders: Signer for simplifying signing logic
* 01bb2bd17 satellite/audit: verifier checks if node made sucess GE before auditing
* 3383acc3c web/satellite: low balance banner removed for no paywall users
* 0518b1637 satellite/piecedeletion: move node info retrieval into the service
* c7b86a348 satellite/admin: add check project usage endpoint and fix some leftover http.Error handling
* 4cdba365e web/storagenode: payout history table
* c54847566 build: Go 1.14.7 (#3926)
* b71da59f8 satellite/metainfo: fix client context cancelation
* a14887e20 satellite/metainfo: Add GetObjectIPs method
* 94a09ce20 all: add missing dots
* 0525858b7 docs/blueprints: Fix typos horizontal scale GC
* c92171024 web/satellite: confirm saving API key modal implemented
* 00f9882ad bump to uplink v1.2.0
* ab1d0f097 satellite/storageusage: Group accounting rollups at_rest_total by day
* 7552ff26e satellite/db: drop project_invoice_stamps table
* 4ee1b2d45 storagenode/console: added list of all audits per satellite to sno dashboard/satellites
* 373934efb storagenode/heldamount: payout history: removed extra doubling with surge percent, added held percent
* 14bb00f71 web/storagenode: all stats held history
* 6ec7bc8b5 web/storagenode: last month estimated payout
* 32e1f16b4 web/storagenode: current month held amount for all satellites
* 6e1ff78cd satellite/metainfo: extend TestAttributionReport
* 2b2cb468b satellite/metainfo: implement DeleteObjectPieces with objectdeletion package
* ceb2eee49 satellite/metainfo: make overwrites impossible without delete permission
* 6e90ca8b0 cmd/storagenode: CLI dashboard last contact renamed to status
* 88dcc93f3 satellite/metainfo: use user PartnerID for bucket attribution
* 68a2726e2 web/storagenode: telemetry removed
* 2ef9f7f64 change doc link (#3924)
* d654ab5fa satellite/admin: change returned data to be json encoded
* f804f03b1 storagenode/heldamount: payuout history updated
* e02adfe5e satellite/overlay/config.go: Add AuditHistoryConfig to overlay
* 5dfe27f17 satellite/{repair,overlay}: Use overlay NodeSelectionCache for repair uploads
* f27285197 docs/blueprint: scale garbage collection
* 85a74b47e satellite/orders: 3-phase rollout
* 935f44ddb satellite/metainfo: Add Delete Service config
* 50ac53f76 satellite/metainfo/objectdeletion: Implement batching
* 53a5d18e1 storagenode: fixed logging about piece being moved to trash, and added logging when piece was actually deleted
* 20184d360 satellite/metainfo: move TestAttributionReport to attribution tests
* 4561d9bdb satellite/console/consoleweb: add support for `partner` parameter for signup page
* edfd3d766 satellite/payments: delete `credits` and `credits_spendings` db tables
* 52c651286 web/storagenode: node id copy
* 89dd0b475 web/storagenode: added ability to show and copy satellite id
* 7735b797c Bump dependencies for various fixes.
* bdb3a10d3 mod: bump drpc
* 14f49558a satellite/admin: fix test failing at the end of month
* dc971751a satellite/stripemock: Fix race condition on customers and transactions
* 76030a823 satellite/audit/{queue,chore}: Wait for audit queue to be finished before swapping
* c89746a38 storagenode: use hardcoded metric application name
* 635b051ae web/storagenode: extra division removed
* b265b7f55 satellite/console: make paywall optional
* d30cb1ada web/storagenode: extra division removed
* cb0caa2e2 satellites/payments: resolve data race in StripeMock
* 5ef4d756a storagenode: estimation usage_at_rest
* b4c9badab storagenode/console: estimation payout fix
* 5988ad664 storagenode/heldamount: payout history extended with satellite's id, url, surge percent
* 123aebd79 storagenode/version: version chore test fix
* f531bc863 storagenode/heldamount payout-history rout fix, usage_at_rest in estimation payout calculations fixed
* 36ed939b8 satellite/orders: add buckets db to service
* 44f919340 satellite/orders: make optimal threshold multiplier into an argument
* ba4c3d998 satellite/orders: remove unused node status logging flag
* 4d2a50578 storagenode/db: explicitly open and create dbs
* e982173d6 web/storagenode: payout page bug fix
* 92efffb48 storagenode/version: notification flow now based on cursor, chore_test added, versioncontrol added to reconfigure.
* 46bdb1dbf Revert "satellite/metainfo: log if replacing pieces in pointer goes wrong"
* e14f7a3fb satellite/repair: update healthyPieces and unhealthyPieces after CreateGetRepairOrderLimits
* 1f9ee4a53 mod: bump common
* 4bcf308a0 satellite/payments: fetch old deposit bonuses from Stripe metadata
* 20437b43a satellite/payments/stripecoinpayments: prevent data race
* cfca02183 satellite/accounting: Add chore to cleanup old project bandwidth rollups data
* 1f5d5235a storagenode/{monitor,piecestore}: if free disk < expected available space, return free disk
* 18b349524 satellite/metainfo: restore tests disabled to remove StreamID from BeginDeleteObject method on uplink side
* aa6afc387 error handling in heldamount cash and collector delete fixed
* 96c83eb0b satellite/payments: log returned error
* a20e85824 satellite/payments: add STORJ amount and rate to Stripe TX metadata
* 0949731ca storagenode/console: estimation payout held split from total payout, calculations fixed
* ce7adc8d8 satellite/console: remove need for nolint
* 65408db6e satellite/satellitedb: Coinpayments repeat insert bug fix
* 375d76638 satellite/metainfo/objectdeletion: Object deletion implementation
* 85aff88e7 ci: update to Go 1.14.6
* e69cec0b7 satellite/metainfo: stop generating StreamID for BeginDeleteObject request
* 67a292d13 satellite/satellitedb: Monitor node tallies
* 4684eb5b7 web/satellite: completed transaction state condition changed
* f1c73c127 web/storagenode: suspension tooltip content change
* fd7bfc94f private/dbutil: don't sort column names in an index
* b67d7ecbc cmd/storagenode,storage/cockroachkv: better error handling
* b719eecb2 satellite/admin: close http response body
* d8dcae307 all: fix error checking
* b84923558 satellite: fix scoping, formatting
* e70da5cd4 all: fix comments
* 1a6e25579 satellite/admin: check for existing user with given email
* 0534b9108 satellite/admin: setup payment account for user
* 080ba47a0 all: fix dots
* 0a800336c web/satellite: token deposit dropdown position fixed
* b33ef461e web/storagenode: available payout periods
* dfdf73282 storagenode/heldamount: db tests updated with payout.Receipt
* 9ace375ee satellite/{console,satellitedb}: change project limiting based on new users field
* 1944d734e satellite/orders: check and enforce node api version
* 1b807761b storagenode/orders: Update orders filestore to be compatible with new satellite endpoint
* 0209a2095 satellite/{console,satellitedb}: add project_limit column to users table
* 2c2d284f3 satellite/admin: add bucket limit handling endpoint
* 784a156ee satellite: prevents uplink from creating a bucket once it exceeds the max bucket allocation.
* 62fec2510 storagenode/heldamount: returns usage_at_rest in tbm instead of tbh
* 3a8766936 satellite/orders: remove race condition in new endpoint tests
* 581579e80 satellite/rewards: Add Duplicati to partners database.
* 257855b5d all: replace == comparison with errors.Is
* 7b4a8c4d6 storagenode/heldamount: payoutHistory added
* 1f1e3f060 satellite/orders: disable settlementwithwindow, skip flaky tests
* 7d6973b5a satellite: heldamount and nodestats not returning error node not found by rpc
* 114941761 satellite/admin: cleanup parameter handling
* 133fda4be satellite/admin: add user deletion
* 0a32ba0e6 satellite/admin: add project rename functionality
* f768302c9 satellite/admin: harden project deletion requirements
* 75ab3507d storagenode-updater: process storagenode config
* 8abb90701 satellite/orders: add settle orders with window
* d4bde7848 web/satellite: coupon status replaced with expiration date
* 02a39354c web/satellite: rel='noopener noreferrer' added to blank links
* 06a3510ae satellite/orders: add EncryptionKey encryption
* bbdb351e5 all: use jackc/pgx in place of lib/pq
* 262da1435 storagenode/console/consoleapi: disable flaky TestStorageNodeApi
* 78f5755d4 storagenode/nodestats: Add sat to heldamount error
* facde770d storagenode/heldadmount: removed logging errors node not found during getAllPaystubs/Payments from trusted satellites
* f73e92c26 storagenode/gracefulexit: added blobs clean
* 29ccda5d4 cmd/uplink/cmd: remove spf13/cast dependency
* f532d4b25 satellite/payments/stripecoinpayments: add test for invoice items values
* 24a1eac16 storage/postgreskv: Sort storage keys before delete (postgres)
* 569b49768 storage/cockroachkv: Sort storage keys before delete (cockroach)
* 9dc9cd8a1 tests: allow STORJ_TEST_POSTGRES
* 5bdcd86fa ci: test benchmarks
* 4869cfc9a satellite/vouchers: remove deprecated endpoint
* e17243fcd storagenode/console: estimation payour for current and previous month reworked
* cdb020028 satellite: collect hardware stats on satellites
* 885ef70c5 satellite/nodeapiversion: new table for tracking node api usage
* 6c5d948b8 web/storagenode: loading screen added
* 9dbd51139 private/dbutil: reduce db connection defaults (#3920)
* 0521435e0 storagenode/gracefulexit: added deletion of all files left in storage/blobs/satellite after successful GE
* dd9eaafce web/satellite: api keys component slightly refactored
* a26fab255 web/satellite: balance history dropdown implemented on billing page
* 2ec9de94b web/satellite: onboarding tour's payment wall updated
* a3c902ab8 storagenode/pieces: hours in a month should be 720
* e9dd5b284 storagenode/piecestore: Properly log/send metrics for all successful pieces
* 41497569c satellite/orders: add settled amount to rollups write cache
* 12a15e5a6 satellite/payments: remove migare-credits billing command
* fd740295e satellite/satellitedb: Add comment to revocation
* db55386fb satellite/rewards: Add Rclone to partners database.
* 78c0d9352 storage: support monkit traces of limit exceeded
* 00ae5ebba satellite/payemnts: Credit coin payments earlier
* ac716e151 storagenode/heldamount: payment receipt added to monthly paystub, heldamount.service separated for service and endpoint
* 5786595d1 clabot: change JTs username (#3917)
* e1a18ce97 satellite/accounting: bugfix infinit storage usage increase (#3916)
* c40d5043e cmd/uplink: pass useragent flag to library
* 35b709ba1 storagenode/storagenodedb: check if db is nil before closing
* 5a9c3a7e2 web/storagenode: blank target links rel added
* e3088d9ad satellite/satellitedb: add new DB table audit_histories
* 577f72cb9 storagenode/version: notifications added
* b639ec08d satellite/heldamount: payments added, endpoind for payments added
* b878fcc4b storagenode/heldamount: id removed from satellite name
* 66661c748 satellite/coinpayments: query status of >25 coinpayments
* 4a98c9514 private/date: fix MonthsCountSince build issue
* cadb435d2 {satellite/audit, private/testplanet}: remove ErrAlreadyExists, run 2 audit workers in testplanet
* 735dc6e16 satellite/orders: add encryption keys
* 9b90712aa storagenode/heldamount: payents added to db
* 9a0214965 storagenode/heladamount: held history extended with joined_at date, total_held and total_disposed amounts
* 3567b49ef satellite/metainfo/piecedeletion: fix int to string conversion
* a40f61dab cla: add alexottoboni (#3915)
* 4997fd55d satellite/repair: remove healthy from irreparabledb
* bdaabd611 satellite/metainfo/piecedeletion: add metrics for deletion requests
* d91cf5f4d satellite/satellitedb: add missing SeparateTx
* 66f536880 satellite/testing: Change testing to use PG 12.3 (#3913)
* 13a585453 satellite/satellitedb: clarify test migration merging
* df8cf8f58 satellite/orders: delete unused code
* db57e5cf7 scripts/tests/rollingupgrade: continue for v1.6.4 (#3912)
* 669d62098 web/storagenode: suspension score added
* 091b49b92 cmd/satellite: command to move unspent bonuses to Stripe balance
* 37cfc01ac cmd/storagenode: Add long description to ge command
* 51dfc6bf4 storagenode/gracefulexit: make minimum transfer speed to be 5KiB
* ed9816fd3 storage/filestore: Ignore IsNotExist error walking files
* 3b4b5f45c satellite: replace references to Suspended with UnknownAuditSuspended
* 5b3c8b2f1 web/satellite: google tag manager for signup pages
* cad21f11e cmd/uplink: Add revoke command
* 4ce1c099b satellite/metainfo: Add RevokeAPIKey to batch
* 34b019d16 satellite/metainfo: Handle revocation request
* 2d727bb14 satellite: Check macaroon revocation
* 433fc9105 cmd/satellite: clean up entrypoint (#3910)
* 98d477eff storagenode/collector: Fix comment doc
* 7e6f9cf5b storj: install latest gateway with `make install-sim`
* 2bd9067ad cmd/uplink: tighter linksharing security
* 6cc7fd5f3 satellite/metainfo: remove old endpoints
* 05a981bc1 satellite/metainfo: Add RevokeAPIKey stub endpoint
* f68e7b3fd satellite/overlay: replace pb.InfoResponse
* 6673125c0 satellite/metainfo: remove code for handling partner uuid
* 19d431ff3 satellite/metainfo: disable old endpoints
* d414f9825 storj: update to uplink v1.1.1
* 7a0778fac cmd/satellite: choose correct Stripe client also for commands
* b72bf4a6d satellite/admin: Make minor improvements in README
* adb5a25b6 cmd/satellite: revert 'organize subcommands'
* 0885ba564 satellite/satellitedb: add new columns for offline suspension
* 7b8e91ff2 satellite/satellitedb: no orders for exited nodes
* 958ea1b9d satellite/accounting: add download limit cache
* 3842118ba storagenode/heldamount: use correct field for repair usage
* 8ecf01ece web/satellite: credit history page implemented
* 2c3fe5597 storagenode/nodestats: unknown_audit_score added to Service
* c38485406 web/storagenode: inactive custom period selection state
* ddd261703 web/storagenode: held history monthly breakdown
* e48177ce8 satellite/payments/stripecoinpayments: add basic summary for invoice generation commands
* 410d89784 satellite: fix string(int) conversions
* d88394c3d satellite/metainfo: verify pointer in service
* 1ed5a1bac satellite/satellitedb/satellitedbtest: skip omitted database
* 0826c8d87 satellite/heldamount: fix dimension of usage_at_rest
* dc5502cb8 private: Prepare pkg for enabling gosec
* be5972779 storagenode/orders: Add archival functionality to orders filestore
* e8ef9b76b satellite/rewards: add CloudBloq
* c6c8b923a satellite/dbcleanup: run cleanup more frequently
* 26fce54b1 satellite/audit: increase MinDownloadTimeout in TestReverifySlowDownload
* bad299b54 satellite/satellitedb: serialize UpdateStats and BatchUpdateStats transactions
* 95a740380 web/storagenode: uptime section removed
* 44792f7b4 cmd/satellite: add finalize invoices command
* dfeb4dc9c satellite/payments: add finalize invoice command
* fdbb2c3ed cmd/satellite: organize subcommands
* dd1fabe2b web/satellite: added no payout data state reset
* e52809d53 cmd/storagenode: add check if satellites available to gracefulexit
* 0b109c32e storagenode/piecestore/usedserials: add monkit metric for serials that are randomly deleted
* a7211badc web/satellite: validate methods replaced into class Validator
* cc0fca850 web/satellite: charges container on billing page reworked
* 47a766328 web/satellite: deposit and billing history splitted to be shown separately
* 1c30efd3a private/testplanet: allow setting "omit" as database to reduce output
* 90e2e3c8c satellite/payments/stripecoinpayments: display more details about invoice item generation
* d9b1c4520 storagenode-updater: fix logic to use minimum version correctly
* 4dd47a460 storj: add bump-dependencies target
* a974db7ee Update automatic updater blueprint to catch missing requirement
* 36c461bd5 private/tagsql: track proper closing of rows and statements
* c78f88beb scripts/tests: use the same version for storj-sim as satellites
* 943eb872d satellite/audit: depend less on details of some error message
* 10f8b5492 Revert "private/tagsql: add finalizer based leak checks during dev"
* 175f048aa scripts;cmd/storj-sim: include satellite node id in satellite address
* 34db4a80f ci: fix staticcheck failures
* 0de1b5fa9 ci: fix golangci-lint failures
* 09ca382ab storagenode/db: preflight improve index discovery
* 9a04ca052 scripts/tests/rollingupgrade: skip v1.6.3
* 5ed1985b2 docs/blueprints: trusted-delegated-repair.md
* 7f8e55302 console/dashboard: added pieces headers size to calculations
* 056ae7ffa scripts/test/rollingupgrade: set correct satellite address for imported access
* 34ee3ef97 bump uplink to v1.0.7
* 2b2efcc66 satellite/payments/stripecoinpayments: move Coupons expiration date sorting directly to listing method
* 96286fde4 satellite/payments/stripecoinpayments: list projects by owner, not by project members
* 9d7713cdd script/testdata: update tracing agent default address
* c8c0a4226 storagenode/orders: begin implementation of file store for order limits
* 6b447a415 web/storagenode: free space calculation for chart fix
* c6310b34d private/tagsql: add finalizer based leak checks during dev
* 254b42ff6 satellite/satellitedb: fix leaked rows from repairQueue.Insert
* eced8dcfe build: Go 1.14.4 (#3904)
* b20ced951 satellite/satellitedb: drop project_id column from coupons table
* fbfc3e5d1 satellite/payments: adjust label for bonus
* 6a60e1e96 satellite/satellitedb: inclusive interval_start in GetAllocatedBandwidthTotal
* 3aa3732b0 jenkins: Enable Cobertura failing for code coverage
* 4d0cb1af7 storagenode/piecestore: verify before checking free disk
* fca4f43a0 storage/filestore: benchmark diskInfoFromPath
* 18914e4d1 web/storagenode: months online calculations changed
* c272872d5 satellite/payments: available coupon value feature
* e8c401031 satellite/metainfo: add missing error check
* 2b3545c0c satellite/satellitedb: use delete returning to query pending_serial_queue
* b82d04e61 satellite/metainfo: limit size of uplink-provided metadata to 2KiB
* 44433f38b satellite/satellitedb: remove ORDER BY when reading from queue
* 163c027a6 satellite/satellitedb: remove monkit trace from convertDBNode
* d66e646b5 satellite/payments: add deposit bonus to stripe balance
* 948cc7333 satellite/payments: mark coupon as expired on its last month
* 89c9672ce storagenode/piecestore: available storage check added in Upload
* df0ef7e0c docs: update access-revocation blueprint
* e79e83b61 storage/cockroachkv: handle retry errors for GetAll and DeleteMultiple
* 36286a062 docs/blueprints: installation and auto-update on Linux (#3463)
* a9f648966 satellite/payments/stripecoinpayments: remove ProjectID from Coupon struct
* 1b69d20eb satellite/payments/stripecoinpayments: apply coupons in the order of their expiration date
* 21518bcc9 private/testuplink: move tests to uplink
* ee9bb0d68 satellite/payments: coupon expiration bug fixed
* 79a562c3e cmd/uplink/cmd: improve uplink remove bucket command with force flag
* 839a6f15d stripecoinpayments: use fixed date for test
* f6452e9c0 satellite/payments/stripecoinpayments: process customers instead of projects
* 801a3ab90 satellite/coinpayments: Reduce update interval to 2 minutes (#3897)
* c9b9c686f web/satellite: logic for new signup/login flow
* 0148f421c web/satellite: onboarding flow's vertical adaptation fixed
* b1bb665c7 satellite/metainfo: Handle "server is not accepting clients" error during CRDB node rejoins
* d5c86aa98 docs/blueprints: tracking downtime with audits (#3834)
* e7e69f383 satellite/repair/repairer/ec.go: Add monkit tracing for ec repairer
* d6c90b7ab pkg/macaroon: remove aliased package
* 2cf4f4d1a bump uplink to v1.0.6
* ff7c726e9 ci: fix coverage flags
* d9fec04f5 lib: remove the old implementation
* 59ca5a1c5 scripts: remove lib/uplink from update-access
* 07050eea2 all: use common/storj
* 48e67a29a cmd/storj-sim: remove old libuplink from storj-sim
* acf8b72cd satellite/repair/repairer: cut off long tail when minimum number of required uploads is met
* 1e065fb45 satellite: migration to fix bad imported payment history
* dc57640d9 storagenode/piecestore: switch usedserials db for in-memory usedserials store
* 909d6d966 storagenode/piecestore/usedserials: add in-memory store for used serials
* 7395dd1e6 storagenode/gracefulexit: revalidate existing pieces
* 0c9a4a5e5 satellite/gracefulexit: better validation error messages
* 84892631c private/testplanet: remove old libuplink from testplanet
* 73214c6d1 storagenode/heldamount: heldhistory reworked to all satellites
* 4ad163de2 storage/cockroachkv: rename opi to oci
* d5dfeca84 satellite/payments: Return empty iterator instead of nil from StripeMock List methods
* e0d71dcae storagenode/windows-installer: ignore set firewall exception error
* 2be93df56 scripts: reduce file size for backward compatibility tests
* f2a0c6442 storage/filestore: log potential disk corruption
* 8db848791 storagenode/console: added estimated payout for current month and estimated pay stub for previous month (until there's real data in satellite's table) + heldback percentage rate for previous month.
* ee7de0424 satellite/payments: use quantities in invoices
* 75b3db542 satellite/payments/stripecoinpayments: test invoice user with more than 1 project
* 290c006a1 satellite/repair/{checker,queue}: add metric for new segments added to repair queue
* ed857f0fb satellite/admin: Add the proper documentation for bandwidth limits (#3900)
* c5d4a1315 satellite/nodeselection: use NodeURL
* 9edb36355 satellite/metainfo: log if replacing pieces in pointer goes wrong
* 427bfc13f web/storagenode: disk stat chrt added
* b6771d0c5 web/storagenode: period selection separation
* d13b693b0 web/storagenode: total payout info area reworked
* 5364a9f74 docs/blueprints: Add access-revocation.md
* 8bd4d7b43 storage/cockroachkv: add check if retry is needed during iteration
* 88a256131 satellite/payments/stripecoinpayments: don't request anything from Coinpayments if credentials are missing
* 7014cf208 satellite/gracefulexit: Change handleFailed to return nil if we can't get the pending transfer
* a4c19b335 cmd/uplink/cmd: remove port detection when selecting satellite while setup
* 2fbb34c3e nodeselection: Increase minimum free space to 500MB (#3898)
* 3da100c24 cmd/satellite:  Install curl per default to docker image (#3899)
* 074649835 satellite/satellitedb: add some docs and improve some snapshots
* 5c1096404 satellite/payments/stripecoinpayments: add test for listing issues while invoice generation
* 340700eda cmd/uplink/cmd: remove old libuplink from Uplink CLI code
* bcb867f1d cmd/uplink:  add UserAgent flag for partner value attribution
* 3d332de22 private/testplanet: remove StorageNodeCount from testplanet uplink definion
* f5ac678b0 storagenode/satellitesdb: added FK constraint to satelliteID
* 45ccf5906 web/satellite: onboarding tour adapted to multiple project state
* f0619c6db web/satellite: error handling for login/signup rate limit exceeding
* 7d00b3c79 web/satellite: billing history item's expiration countdown fixed
* 83cc80ef0 satellite/payments/stripecoinpayments: fix listing issue for coupon usages and credit spendings
* bef84a5f9 storagenode: remove dependency to overlay.NodeDossier
* b42778c42 private/testplanet: remove some additional Local-s
* 7f323754a metainfo/piecedeletion: use NodeURL-s
* 03e5f922c satellite/overlay: updates node with a vetted_at timestamp if they meet the vetting criteria
* 2514d6328 dbutil/cockroachutil: add monkit to QueryContext
* f43cb1688 private/tagsql: verify SQL connection with ping
* fe6a6f063 private/testplanet: cleanup predefined data generation
* 2c9afe7f1 storagenode/console/api/helamount: periods with heldamount data endpoint added
* 49ad90dcd storagenode/reputation: unknown_score (unknown_alpha / unknown_alpha+unknown_beta) added to reputation stats, https://storjlabs.atlassian.net/browse/SG-326
* 5d016425f satellite/{contact,downtime,overlay}: use NodeURL
* 941d10cbc private/testplanet: remove Peer.Local()
* 94b2b315f storagenode/trust: refactor GetAddress to GetNodeURL
* ed627144e all: use DialNodeURL throughout the codebase
* 963db8c73 cmd/segment-reaper: Remove unneeded test cases
* 1a4ab3df4 satellite/admin: add user update endpoint (#3896)
* 8ec64f3da satellite/overlay: enable node selection cache on all satellites (#3895)
* c7cbdcbcc cmd/internal: request lowercase access name
* 6fac1ed13 satellite/contact: rm unhelpful verbose logging (#3894)
* 705e82ea9 private/testplanet: add AddUser and AddProject to satellite functionality
* 0a26c4af9 satellite/admin: add coupon deletion (#3893)
* 671aca56b satellite/admin: add coupon creation and listing (#3891)
* cbdd3bfd7 cmd/segment-reaper: Rename test name to what it tests
* ac375d37b satellite/payments: remove mockpayments and add Stripe client mock instead
* 16abf02b3 satellite/{nodeselection,overlay}: use the new package
* 17d3cb155 satellite/admin: add user creation endpoint (#3889)
* 1de813d22 cmd/storj-admin: initial commit of new tool (#3890)
* 5ea82dfb2 cmd/segment-reaper: Report zombie segment's size
* 08692aef9 satellite/nodeselection: node selection with proper bias
* 1b06e2a92 satellite/metainfo: remove deprecated error validation (#3876)
* 70502658c satellite/admin: add delete project endpoint (#3888)
* 1ecbfc453 satellite/overlay: add random test for node selection from cache
* d27cb60f8 cmd/segment-reaper: Remove unneeded test logic
* aac1e3c45 satellite/payments: move inspector commands to satellite cli
* 47def0209 satellite: account balance divided into Free Credits and Coins
* 9fb423ac1 satellite/payments/stripecoinpayments: fix listing issue in InvoiceApplyProjectRecords
* 4544f2dff web/satellite: navigation side bar reworked
* 8d87a6efc cockroachutil/driver: handle retryable errors returned from Next
* ef2671927 storagenode/piecestore: move queue size defaults (#3881)
* ffbaaff17 web/storagenode: node module store tests added
* 26f36d590 cmd/satellite: ensure we only create stripe customers for activated users (#3884)
* f6301612a satellite/orders: use serial SerialNumber
* 2eb2c25e5 satellite/payments/stripecoinpayments: add StripeClient interface to be able to cover more testing scenarios
* 1ec5eb06b web/satellite: account dropdown reworked
* a3459ec6b web/satellite: project dashboard page UI reworked
* f03b23d2d storage/postgreskv: monitor calls to sql.Next
* 28a886f39 satellite/admin: obvious project name in documentation
* 49571f1a2 satellite/payments: all invoice commands require period
* ef8719212 storagenode/notifications: tests fixed: added time interval between inserts so created_at fields are different when running tests on Windows
* 46df8c197 satellite/gracefulexit: add log message when node fails validation for piece transfer
* 6352d4610 satellite/satellitedb: do better ::date conversions
* 75ddecdf5 satellite/metainfo: Add test to verify payment report for attributed bucket
* 5a7a4d2e9 satellite: add Go test version of satellite-config-lock tests
* ff68a3c2b web/storagenode: payout page scroll area extended
* 159df8b2e Add logging listener for retrieving and setting log levels
* 4b4222c2f web/storagenode: disposed held added to total held calculations
* 1eab5e298 satellite/console: Increase default webUI rate limit to 5
* 0e3be60b7 satellite/satellitedb: simplify migrate step
* 0e263f009 satellite/admin: add API endpoint for adding projects
* c58dccb7a cmd/segment-reaper: Remove segments of skipped objects
* 931409284 cmd/segment-reaper: Rename Bitmask to BitArray
* 50e831427 cmd/segment-reaper: Fix bitArray.IsSequence method
* e23bd806b satellite/accounting: separate usage and bandwidth limit (#3878)
* 0619f97e2 satellite/console: remove unused arguments
* 65f3e26f8 satellite: Change Default Project Limits and minimum STORJ Payment (#3877)
* 073f9ed49 satellite/payments: deduct stripe discount first
* 93f1fe49e mod: update to latest common
* b397dfc82 satellite/payments: better labels for discounts in invoices
* e0cf9ae88 storagenode/orders: set devDefault sender interval to 30s
* 22fbe804e satellite/accounting: test if project bandwidth limits reset with billing cycle
* ec589a828 all: fix comments about grpc
* 7d29f2e0d all: remove drpc wrappers
* e6d5ce6b7 all: remove grpc
* 4b612a790 cmd/segment-reaper: add support for >65 segments
* bcd93ee37 private/testplanet: add StopNodeAndUpdate
* 12fa50569 satellite/payments: fix invoice amount in billing history
* 94503db5a cmd/satellite: Enable admin service like other services, also sort environment variables (#3874)
* a549d523d bump drpc and common
* 678b85917 satellite/overlay: remove MinimumRequiredNodes
* 2af713892 web/satellite: change invoice link title to 'Invoice PDF'
* 9837230f1 web/satellite: project dropdown reworked
* ce6a500b0 satellite/overlay: support DistinctIP=false in selection cache
* d8bdc60e1 storagenode/console: remove last ping id and address
* 90d859fbb private/testplanet: use drpc piecestore mock for testing
* c5452a87e private/testplanet: use drpc referral manager server
* 2955c50bc satellite/overlay: fix data-race in node selection cache
* 693f224ff storagenode/web: Segment.io integration
* 73b2cd3fd docs/blueprints: blueprint nodes cache (#3838)
* 4e94da3fd satellite/overlay: add feature flag for node selection cache
* e4da7d65f pkg/macaroons: fix action types comments
* efee3563a storagenode/storagenodedb: fix joined_at migration
* b754806b2 storagenode/reputation: unknown_audit_reputation_alpha and beta added to db, and reputation endpoint
* 8a443bc9f storagenode/console: api notifications tests added
* b5433322b cmd/storj-sim: enable open user registration
* 16cd9b06e storagenode/heldamount: added api for heldamount history separated by periods
* 0ae0cea93 web/satellite: billing history UX improved
* 18078bf7e satellite/audit: increases audit worker concurrency to 2
* c5871b83b web/storagenode: no data state added to payout table
* b7495d414 mod: update sqlite to v3.31.1
* d98b8f6e2 satellite/metainfo,storage: use different limit for metainfo loop
* 13baacc47 lib/uplink: better diagnostics with invalid encryption access
* 324fd410a updating Aha link on our readme
* 8f60cfc4f satellite/overlay: Add flag for enabling/disabling disqualification from suspension mode
* f93d76499 cmd/uplink: properly path escape shares
* 31b322f07 pkg/server: log apikey head
* acf53bea4 satellite/orders;accounting: Add monthly project download bandwidth rollup
* 22691bd47 satellite/accounting: test project limits reset after file delete
* 0b80d6049 cmd/uplink: disable tlemetry if user opt out
* 2284008b8 web/satellite: onboarding tour: api keys and upload data steps
* 7a83473f0 storagenode/console: API /dashboard disk-space data separated for UsedForPieces, Trash, Available
* 16081f53d storj: update to uplink 1.0.5
* 2fad15095 storagenode/piecestore: remove piecestore.Client.Delete
* 57eb8a17e storagenode: allow configuring database path independently
* c630cf249 storagenode/pieces: implement buffering for writing
* 01e0ba2e0 satellite/metainfo: move logging to debug for piece deletion (#3873)
* 4ad01e817 storagenode/storagenodedb: backfill reputation.joined_at
* 6c4d3f133 storagenode/dashboard: trash added to avaliable space calculations
* 202e65e40 storagenode/storagenodedb: do correct generalized alter table procedure
* 516b8cf2b storagenode-updater: add recovery for windows service restart
* bb2885196 test-sim-backwards: refactor, reduce node count to 9 in mixed mode
* 5a85e8d74 satellite/audit: Fix flaky TestVerifierSlowDownload
* 2b00d78ad satellite/payments: add documentation
* 4a95ad0b2 storagenode/piecestore: remove error from info log for download canceled (#3869)
* 0dfbdae61 private/date: MonthsCountSince removed, being unused anymore
* 2c47a9b1d pkg/server: log api key in debug traffic
* 4e5e6c8e5 web/satellite: onboarding tour: create project step
* f54a4960a private/date: TestMonthCountSince temporary fix
* 00b2dfa31 storagenode/orders: fix TestDB and TestCleanArchive
* 418a4cc92 satellite/console: fix flaky TestProjectsList
* 8928399d0 all: rename CreateTables to MigrateToLatest
* 6a6427526 satellite/overlay: remove old updateaddress method
* 4f492a1ca storagenode/piecestore: add wait to test
* e26a5598e satellite/overlay: fix UpdateCheckIn test
* d225e2de4 all: add missing ctx.Cleanup calls in tests
* 51f69d53d private/testplanet: fix closablePeer
* 237d9da47 storagenode/pieces: Deleter can handle multiple tests
* 131654b08 storagenode/piecestore: make tests use DeletePieces directly
* 19df9beb3 web/satellite: pagination position unified
* 518946fab satellite/metainfo/piecedeletion: add metrics for unhandled pieces
* 70fe974a9 satellite/payments/stripecoinpayments: fix test flakyness
* 588826157 satellite/metainfo: reduce RateLimit test flakyness
* 39d77c31e certificate/authorization: fix test name
* 621b4e7f9 satellite: remove non-important debug lines
* 3785098f1 versioncontrol: remove per request debug
* b2acd93a7 storagenode/inspector: adjust Uptime type change in protobuf
* 004d82987 satellite/metainfo: fix TestEndpoint_DeleteObjectPieces_ObjectWithoutLastSegment test
* de366537a satellite/satellitedb/overlaycache: fix behavior around gracefully exited nodes
* baccfd36b private/testplanet: Mark sn peer deleter test mode
* c5309a3f9 cmd/uplink: set sample rate for tracing to be 1 when tracing is enabled
* 42f63c653 satellite/compensation: add offline status tracking
* cafa7a5f0 satellite/admin: add a readme about endpoints
* 85c45cd56 private/dbutil/pgtest: support multiple databases for testing
* 6f84be133 satellite/metainfo: add MigrateToLatest to PointerDB
* 13bf0c62a satellite/pieces: Fix race in piece deleter
* a3eeab291 web/satellite: UI unit tests for billing history
* 849326efe satellite/console: cleanup rate limiter
* 51bf2b615 web/storagenode: dependencies for testing and payout store tests added
* 9b4a3f8fc cmd/uplink: use tracing.enabled flag
* ef913be23 satellite/satellitedb/satellitedbtest: don't use subtest naming
* db0371703 storagenode/pieces: Return UnhandledCount to satellite
* edda8d73b storagenode/pieces: Piece deleter monitor queue
* befe7574e web/satellite: onboarding tour: adding payment methods step
* 03871d17c satellite/satellitedb: Update ticket ref
* c52fc964d Upgrade storj.io/uplink to v1.0.4
* 533a65a29 web/storagenode: disq and suspended text color fixed in dark mode
* 16d9d8683 web/storagenode: added held amount in table for current period
* 895eac171 web/storagenode: api calls ungrouped, removed extra current period call
* d73630fd4 satellite/satellitedb: Ensure we just return bucket usage for buckets that exist (#3863)
* 341aecfe0 satellite/console: add rate limiter to login, register, password recovery
* 825226c98 satellite/overlay: use node selection cache for uploads (#3859)
* 7a4dcd61f satellite/overlay: add changes to selected node benchmarks (#3862)
* a785d3715 storagenode/pieces: Process deletes asynchronously
* 1336070fe storagenode/piecestore: add missing log message about audit errors (#3861)
* 720e26d23 satellite/satellitedb/overlaycache: update unknown alpha/beta values properly
* a0692d0db private/migrate: enhance docs in some funcs
* 72b93f312 satellite/satellitedb: disqualify suspended nodes when the grace period passes
* 676f3e851 satellite/metainfo/piecedeletion: try to make batches larger
* 0bdcf123c bump monkit, monkit-jaeger, and private to latest
* 60e07f0a8 Revert "satellite/accounting: Remove unnecessary index bucket_bandwidth_rollups_project_id_action_interval_index"
* 178aa8b5e satellite/{metainfo,repair}: Delete expired segments from metainfo
* e34937317 Call the right make target for the segment-reaper (#3860)
* 0b8699bcb cmd: add prompt for enabling tracing during uplink cli setup
* 805e328c4 storagenode/heldamount payments removed
* 105dc7acc satellite/accounting: Remove unnecessary index bucket_bandwidth_rollups_project_id_action_interval_index
* 7365fc434 web/satellite: available credits disabled due to coupon expiration bug
* 89c877f46 web/storagenode: payout calculation fix
* 3d56efc82 storagenode/console/service: Satellites EarliestJoinDate calculation ignores empty date
* 1e295a48e add container image for segment-reaper (#3855)
* e655e160d private/testuplink: delete delete
* 030007668 web/storagenode: audit checks based on score
* a23751212 web/storagenode: added division on price multiplier
* ed701c196 web/storagenode: disk space displayed by hour
* e999f24e5 storagenode/nodestats/cache: storagenodeDB/heldamount sync with satelliteDB/storagenode_paystub
* b9dbd8051 web/storagenode: all fetched paystub data treated as list
* 8375a09c8 cmd: remove InitTracing from satellite and storagenode main.go file
* 5f38f8f1f satellite/gc: use hostname for metric instance ids instead of node id
* 1af70703e storagenode/console/service: SatellitePayStubMonthly returns array of objects
* 884638a29 scripts: reduce segment size for integration tests
* c021b3587 private/testplanet: migrate testplanet to new libuplink
* ac1cc9eed web/storagenode: payout fetch error fix, held row added
* 6e3585e39 satellite/heldamount/endpoint : GetAllPaystubs added
* 31c9cc322 web/storagenode: dark theme
* aa6cff405 web/storagenode: payout page bug fix
* 52cd33a63 scripts/deploy-storagenode: tag docker image as latest (#3858)
* 4cd86ff78 satellite/accounting: Add index on bucket_bandwidth_rollups for action, interval_start, and project_id
* 7403e9a8a satellite web ui: fix spelling mistake
* 30369b027 storagenode/storagenodedb/reputation: add joined_at
* 7e0e74c65 satellite/metainfo: attribution based on useragent
* a6f0be204 cmd/inspector: decode path
* c498e52cf satellite/nodestats/endoint: reputation stats updated
* 9052085f7 private/testplanet: simplify uplink usage
* a129a8bd3 all: separate err check for http
* 2dce4c232 web/satellite: redirect to verification page on sign up if inside iframe
* 93e39a6b3 satellite/metainfo: support co-joining the loop
* 45d1ca87f satellite/metainfo/piecedeletion: handle zero case
* dc78cd963 storagenode/piecestore: fix annoying info log for upload canceled (#3857)
* 75b9a5971 satellite: update log levels (#3851)
* d3ce845f8 satellite: revert log lines used to figure out node id
* 90bac5154 cmd/uplink/cmd: improve error message when config.yaml or --access flag is missing
* a2ce83676 remove sugar logging
* 158013a86 satellite/console: redirect on account activation
* 2c0d61b18 satellite/metainfo: avoid temporary list
* 921b57455 docs/blueprints: slow down and retry (#3826)
* 5ea1602ca satellite/overlay: add selected node cache (#3846)
* d7794a485 satellite/overlay: hardcode default values for audit alpha/beta
* e2d5b403e cmd/uplink: support --force (like awscli) for rb
* 6109fa685 web/satellite: allow users who have balance to create projects
* bb4b7a919 storagenode/console/service satellites extends with data of oldest join to satellite
* 743b3fb22 storagenode/nodestats: add pricing model, storagenode/cache: add paystub history storing
* d392c6e65 web/satellite: copy api key container reworked
* 3ee6c14f5 satellite/downtime: add concurrency to downtime estimation
* 17ccf36c1 web/satellite: payment methods component divided into smaller ones
* 68946c86b cmd/storagenode: update colored storagenode dashboard to have correct indendation
* c97131ae7 satellite/metainfo: organize attribution methods
* b9dc72c8b update links in README.md
* 67a19ec95 docs/blueprints: update link to uplink docs
* d08b2de55 lib/uplink: deprecate docs
* e35cbfb6e bump storj.io/common
* 1cad686e9 storage: reduce default lookup limit to 500
* 83bbc7a37 satellite/accounting/tally: ignore expired pointers
* 02613407a satellite/satellitedb: only suspend node if not already suspended
* 4e2a101fc web/storagenode: audit checks depending on score
* e33da9087 private/dbutil/cockroachutil: stop checking for jackc/pgx
* d86cce202 satellite/satellitedb: use arrays for arguments in node selection
* ccf4f9ed2 satellite/satellitedb: node selection code cleanup
* cf80b3caf satellite/overlay: combine SelectStorageNodes and SelectNewStorageNodes (#3831)
* d658a6a6e private/dbutil/txutil: fix logic in transaction retries
* 14b3704f5 storagenode: add suspended status to storagenode dashboard/api
* f36e8548f satellite/metainfo: adjust max inline segment size validation to potential encryption overhead.
* 11a44cdd8 all: don't depend on gogo/proto directly
* 0a1bbc982 satellite/overlay: add TestEnsureMinimumRequested
* 3a9422cc9 satellite/nodestats: add pricing model to endpoint
* cf26951a5 satellite/satellitedb/pbold: remove dead code
* 1eb501b5c satellite/accounting: test zombie segments not billed
* 96e58d21b cmd;pkg/server: init tracing collector in all processes
* 43cf03611 installer/windows: don't require ca.key file (#3847)
* 80ee7321c satellite/accounting: add test billing without expansion factor
* ef6c9d89b web/satellite: dashboard blur step after adding payment method removed
* 2ded64ba2 satellite/compensation: more fixes to get prod running smoothly
* d23e03f8a lib/uplink-gomobile: extend encryption key capabilities
* a4c554f2e satellite/admin: support user query by email
* 0f71e60a5 storanode/version/chore notifications temporary disabled
* 61b6ff518 web/satellite: available credits amount added on billing page
* 2d1a6968b satellite/metainfo: revert returning segment size and inline segment size with BeginObject
* 1547e791a satellitedb: remove free_bandwidth column from nodes table
* b17dc656b lib/uplink-gomobile: update to latest master
* 4ab553cf3 scripts: only latest releases have version file
* 42be4bdc0 satellite/contact: add timeout to PingBack method
* 9200efc61 satellite/satellitedb: fix selecting a nullable string
* 0c8c11b25 satellite/audit: add not_enough_shares_for_audit counter
* dbd036de8 web/satellite: api key and proj member name's length visibility extended
* 4b01a8dd1 Descellate Usage Error to Debug (#3780)
* 06cddb053 cmd/identity: improve CA key message (#3819)
* 0a4d25399 storagenode/storagenodedb: Improve preflight schema error message (#3844)
* b3939bc1b scripts/tests: fix test-version and rolling-upgrade test installation
* 4f8cf53c9 Remove VOLUME /root/.local/share/storj/storagenode (#3820)
* a409bd5de satellite/orders: check for expired orders first
* 6492b13d8 all: remove old uuid
* e72553587 cmd/satellite: add gc to entrypoint (#3842)
* 1024bf9ce all: simplify uuid usage
* fe2340285 storagenode/console/consoleserver: fix TestConsole tests
* c178a08cb satellite/metainfo: add max segment size and max inline size to BeginObject response
* 4a79b609e satellite/metainfo: fix panic when we batch BeginObjectDelete without all permissions
* 8f73fb7a3 all: simplify uuid usage
* f66390635 storagenode/contact: call return value from mon.Task() on function finish
* 90319dbec scripts: fix test-sim-backwards
* ffe7a3c21 satellite/compensation: make surge percent an int64 for strictcsv
* 644df8dcd private/version: minimal fix for tag-release.sh
* 9a72b5fde web/storagenode: byte*h to byte*month conversion
* d444fbade scripts: cleanup rolling upgrade test
* 473cd4695 bump storj.io/uplink to v1.0.3
* a416b0394 satellite/accounting: fix TestProjectBandwidthTotal
* 9bd0bd0c2 private/currency: add strictcsv support to microunit
* 0a69da4ff all: switch to storj.io/common/uuid
* 1d79228ed satellite/metainfo: support uplink useragent
* 1fb37915c heldamount date picking hot fix
* dc32f1da5 storagenode/cache/heldamount added, errNoRows ignored
* 0374e3067 scripts: fix storj-sim installation function for rolling upgrade tests
* e2ff2ce67 satellite: compensation package and commands
* 23e5a0471 satellite/audit: clean up logging (#3832)
* a73416e20 revert to use the config for storagesync (#3837)
* d77f3b878 satellitedb/migrate: set vetted_at backfill to now.day
* 439aba922 satellite/overlay: reduce overhead of GetNodes
* cb781d66c satellite/overlay: optimize FindStorageNodes
* 028ce1602 cla: Adding sixcorners to the CLA approved list (#3836)
* a6540dc3e satellite/overlay: remove unused KeyLock
* de903a652 satellite/accounting: Test no repair traffic in billing
* 048ca4558 satellite/repair: clean up logging (#3833)
* 831668478 scripts/tests: fix gateway installation
* 4122be154 web/storagenode: payout page adaptation
* 5c6c79773 web/storagenode: payout info logic implementation
* 6cac75591 web/storagenode: notifications page adaptations
* f28100b73 bump storj.io/private
* c97096950 satellite/overlay: add benchmark for node selection
* 480ea1e4b satellite/repair/repairer: fix temporary file handling
* e1a443b04 private/testplanet: allow modifying created database
* a933bcc99 satellite/repair/repairer/ec.go: add option for downloading pieces onto disk instead of in memory during repair
* e8f18a2cf private/testplanet: expose storagenode and satellite Config
* b47381138 web/storagenode: added hash to bundle name
* 9d3d411ca satellite/accounting: Add test billing inline segments
* 75961fe8c cmd/gateway: remove gateway command from repository and adjust `make install-sim`
* 551182766 satellite/orders: don't log expired order limits
* df462d726 satellite/accounting: Add index on bucket_bandwidth_rollups to minimize full table scans
* 8e0ca0e6f satellite/gc: update release default for gc to run separately (#3830)
* 23da9228b satellite/console: email used error handling for registration
* 506582e20 web/satellite: project overview page removed, dashboard page implemented
* 97e980cd8 private/dbutil: add database name to configure as a tag
* f879bfcf7 storagenode/console/server/heldamount - endpoint tests added
* fecd36f6b bump uplink version to latest
* 4d91533a3 Update README.md (#3828)
* b75cbc8e2 satellite,storagenode: remove references to free bandwidth
* c715c75fe pkg/server: add counters for grpc calls
* b7b19289d bump storj.io/common to latest
* 0a50d6bdc bump storj.io/uplink to v1.0.1
* a73147249 bump storj.io/common to latest and storj.io/drpc to v0.0.11
* 2a482e8bc private/version/checker: remove code which was moved to `storj/private/debug` package
* bd9a998ab pkg/server: remove dead code
* 3a3648c80 scripts: add script for running versions tests locally
* fdf40a752 storj: remove `storj/private/version` package which was moved to `storj/private` repo
* bdec51658 cmd/storj-sim: Increase storj-sim max-alpha-usage (#3824)
* aed8dea62 satellite/accounting: Add test not billing after deletion
* 173cb1e48 Changing LogLevel to Warn (#3822)
* a9fb7b769 storagenode/trust: fix go 1.14 failures
* 326c0cebd storage/boltdb: update to etcd/bbolt v1.3.4
* f0aeda309 storj: remove from `storj/pkg` packages moved to `storj/private` repo
* ddf87286c cmd/statreceiver: Remove statreceiver code
* 6a7571f73 cmd/s3-benchmark: move to storj.io/benchmark
* d7558db5e web/satellite: succesful registration flow reworked
* fe39845a8 web/satellite: refreshing billing history on entering billing page implemented
* f61b77ea3 web/satellite: request limit link added on project details page
* 859fce3af web/satellite: browser caching bug fixed
* e258de518 cmd/storj-sim: remove identity dir from gateway configuration
* 6c44512c1 satellite/metainfo: increase stream ID expiration to 48h
* 06c499ccf ci: remove .golangci.yml
* 98ac16fb5 cmd/wizard: improve prompt for satellite
* 378173381 pkg/process: replace storj/pkg/telemetry with common/telemetry
* 3b66ba6f0 scripts/tests: uplink no longer respects --client.segment-size
* aeab599d2 satellitedb: removed unused id on storagenode_storage_tallies table, add index on node_id
* 05da31a97 statreceiver: Add name of buffer for better logging/debugging
* 1b6ab173a private/context2: moved to storj.io/common/context2
* d994584e8 cmd/storj-sim: remove removed flags from storj-sim gateway setup
* 9b78473c0 satellitedb: adds vetted_at nullable timestamp to nodes table
* b2590cf28 bump uplink to 1.0.0 (#3816)
* 699b635e5 satellite/overlay: rename newNodePercentage to newNodeFraction
* 0781a91ff linksharing: move to storj.io/linksharing
* eb1d8aab9 satellite/metainfo/pointerverification: service for verifying pointers
* 8597e6b51 storagenode/console/api period payment api extended
* 0df586c3a satellitedb/heldamount updated, tests added + storagenode console updated
* fde5c3542 storagenode/console/api: period payStub api extended
* 10b032e48 libuplink: return deleted bucket/object (step 4)
* 78b253c77 libuplink: return deleted bucket/object (step 2)
* 514287414 satellite/gc: move garbage collection to its own process
* 19685ad81 go.mod: bump common and uplink versions
* 09e0f3de6 satellite/metainfo/piecedeletion: add Service
* 115f4559e satellite/orders: more efficient processing of orders
* bd4982e24 test/backwards-compatibility: Exclude rc tags from testing (#3787)
* b10b69d9c test/rollingupgrade: fix stage 1 release version (#3810)
* ba5991dc8 satellite/repair: add monitoring for remote_segments_healthy_percentage
* 2f991b6c5 satellite/{overlay, satellitedb}: account for `suspended` field in overlay cache
* 80acf33ab script/release: fix error in regex (#3809)
* 81afbcc12 satellite/metainfo: check bucket existence on upload and listing
* b30a9cb1f web/satellite: readme structure section
* 7baa59753 satellite/orders: add tests for double sending the same order
* 2cd5eb7da web/satellite: navigation side bar reworked
* ad9cac308 satellite/metainfo: reduce test flakiness
* 22ea0c7c1 satellite/metainfo/piecedeletion: add Dialer
* bdbf764b8 satellite/orders;overlay: Consolidate order limit storage node lookups into 1 query.
* 49a30ce4a satellite/payments: Set proper defaults for the release (#3806)
* 8b72181a1 satellite/{audit,overlay,satellitedb}: implement unknown audit reputation and suspension
* 52590197c satellite/payments: More Cleanup and Satellite command to ensure we have stripe customers (#3805)
* 3d6518081 satellite/metainfo/piecedeletion: add Combiner
* 97df9b570 web/satellite: new project popup reworked
* dedb87a6d web/satellite: billing periods logic implemented on billing page
* 5dc1b7db9 storagenode/cache heldamount disabled
* 54dbc5173 web/satellite: limits info bar added to billing page, free credit amount changed
* 27f811a9e metainfo: delete methods return the deleted item
* 4f0bf3fe1 build: cleanup more gateway targets from Makefile (#3802)
* 44ade00e8 Add inspector to satellite image (#3801)
* 135fa7500 web/storagenode: pages align fix
* f12b6dc27 storagenode/console/api: monthly payStub api extended
* 89374e260 storagenode/console/consoleapi: using cached data in heldamount api
* 9f84261c3 storagenode/cache heldamount added
* 3a1b71a75 web/storagenode: payout info month range selection logic
* 0c18ecf32 storagenode/api refactored
* 45507d285 storagenode/storagenodedb: heldamount tests added
* 94c4d1e73 satellite/satellitedb/heldamount added, endpoint added
* 988bb5285 storagenode/heldamount GetPayment added, console/server updated
* 7b0371e9e storagenode/heldamount/service added, console/heldamountapi added, console/server updated
* 5ccce0433 storagenode/storagenodedb: heldamount added
* bd603c075 satellite/payments: Improve Invoice Generation  (#3800)
* 94c11c521 satellite: remove some unnecessary UTC() calls
* 6e79d1c20 web/storagenode: charts swap due to bandwidth bar info deletion
* 41887883f satellite/satellitedb: check indexes on migration
* 85b6316ce statreceiver: Update to filter on packet headers
* 39cb82119 satellite/overlay: rm combinedcache, fix IP naming to be network (#3798)
* 3cf05b24d scripts: add benchmark test for delete operation
* 176d64336 web/storagenode: payments markup
* 7f0a6c9ee web/satellite: project creation tests extended
* 8e274a5ce cmd/uplink: Enable telemetry on ctx
* 202b12d21 web/storagenode: dashboard styles adaptation
* 02aee17cd accounting/projectlimit: reset at the beginning of the month (#3796)
* ce1ae134a build: don't push gateway docker images
* 1b5b153fb web/satellite: adding payment method flow reworked, animations added
* 051569c69 satellite: enable open registration (and add flag that disables it) SM-441
* 803e2930f satellite: use IP for all uplink operations, use hostname for audit and repairs
* 140e2f004 web/storagenode/src: removes references to available and remaining bandwidth
* 1a875baa1 scripts/tests: fix arguments orders passed to test-versions.sh
* 520b16e82 satellite/console: allow for project limits even with open registration
* 7aa30d2f0 accounting/projectlimit: remove expansion factor (#3795)
* eda8778f4 golangci-lint: Disable ifElseChain go-critic rule
* 46b04a38c scripts/tests/: fix uplink access to contain satellite id
* de59e8e12 web/satellite: project creation restricted before adding payment method
* 1baf1bd24 satellite/satellitedb: Add index on num_healthy_pieces column in injuredsegments table
* 56c33f519 satellite/payments: project charges api extended to show usage and period
* 4abbf3198 web/satellite: login loading screen removed
* 208e88f2c web/satellite: no api keys state implemented
* 16878a22e satellite/metainfo: stops hiding real validateAuth
* e4d5addb0 cmd/uplink: add url-based link sharing
* 79553059c satellite/repair: put irreparable segments in irreparableDB
* 178dbb468 storagenode/storagenodedb: allow storagenodes to start test_table exists
* 20e96d417 satellite/metainfo: fix data race in test
* d7b5df70d cmd/uplink: remove unused flag
* c20cf25f3 cmd: migrate uplink CLI to new API
* 4e9cd77d5 web/satellite: update links to APIKeys documentation
* 0675413f7 satellite/satellitedb: increase migrate test timeout
* 842c8d8ed scripts/tests/rollingupgrade: fix installation for current commit
* e4da7bd9c satellite/repair/checker: use repair override if available in checker and irreparable
* e99e675fb satellite/satellitedb:  use time zones with all timestamps
* 0d60c1a4b satellite/audit: fix checkSegmentAltered to detect segments that have changed during an audit
* e6d452dec satellite/accounting: Billing tests wait for SNs
* f4d5d89b6 private/testplanet: add WaitForStorageNodeEndpoints
* 9f390f37d satellite/metainfo: return default ciphers (path and encryption) for old uplinks
* a7f927df9 satellite/accounting: Disable billing test
* 2af71f346 satellite/orders: add monkit to looking up node addr
* 5c9becb9b satellite/orders: billing partial download
* 1c1750e6b removes bandwidth limiting
* 5f2ca0338 satellite/satellitedb: fix err and close order
* bcb0453db upgrade dependencies for trace db debug endpoint
* 8fa8178f0 release/rollingupgrade: on release tags run rolling upgrade against previous release (#3792)
* 7244a6a84 storagenode/{contact, piecestore}: implement low disk notification with cooldown
* d384e48ad private/testplanet: set rollout seed to avoid warnings in logs
* decb2ec69 private/processgroup: moved to storj.io/common/processgroup
* a02424a22 pkg/server: use common implementation for user timeouts
* 443aa08a0 private/dbutil/txutil: remove the individual retry events
* f495544c5 satellite/satellitedb/dbx: add fields to node table for placing nodes into suspended mode for too many unknown-error audits
* 484ec7463 storagenode: notifications on outdated software version
* 4d3db6828 cmd/gateway: fix go.mod formatting
* df88f416c satellite/accounting: Add test billing download traffic post deletion
* d64ef3d89 satellite/accounting: Test billing donwload/upload traffic
* 4deab5ac6 satellite/metainfo: combine CommitSegment and CommitObject in batch v2
* 1db087cfb satellite/satellitedb: migration to create tables for compensation
* 6043d01c9 satellite/audit/verifier: add metric for number of successfully downloaded shares
* 1f7c3be8f private/testplanet: add option to run testplanet databases non-parallel
* fb2711d05 scripts: update postgres helper script to set password
* b0d2cf0e4 web/storagenode: on logo click action added
* 2b9f28b02 satellite/accounting/reportedrollup: remove expiration check
* f85606b5a private/grpctlsopts: grpc related tlsopts
* d5540c89a satellite/repair/checker: add monkit metrics for segments immediately above repair threshold
* 46228fee9 cmd/gateway: use proper module name
* 64330c55b all: use pbgrpc
* 8822e98c1 cmd/gateway: simplify module handling
* 89e5c77d8 satellite/metainfo: track observer timing
* 4e5a7f13c satellite/repair/queue: Prioritize selection of items off repair queue by segment health
* ac34485f5 scripts/tests: install correct version of gateway
* 9a8db0583 web/satellite: updating billing history after render added
* fc105af0e web/satellite: user select text restricted
* 594d6e03a docs/blueprints: Add design doc for distributed tracing.
* e486a073c docs: Add uplink telemetry doc
* e19e3c110 pkg/process:
* b387c6b90 web/satellite: overview page implemented
* 29452d82a go.mod: unlock graphql dependency and bump to latest
* 9752d0188 private/prompt: remove dependency to go-prompt
* 50a21de9d traces: fix memory leak for long running traces that aren't being collected
* d57810267 storagenode/piecestore: add workgroup to endpoint to prevent stray goroutine after shutdown
* 92d86fa04 satellite/repair: fix repair concurrency
* f22bddf12 {storagenode/contact, private/testplanet}: remove ErrFailureToStart and panic in testplanet.Start
* 8ea620b3c satellite/console: redirecting to login after activation implemented
* f671eb2be satellite/satellitedb: use queue for orders to get back fast billing
* dca6fcbe2 satellite/payments/stripecoinpayments: credits added to invoice calculations
* 985c3ef89 satellite/console: handling graphql errors bug fix
* e30f7b35b cmd/gateway: use a separate repository
* 5011e7831 storagenode/piecestore: remove unused DeletePiece endpoint
* a645e52ed satellite/metainfo: remove DeletePieces_node_id metric
* 5132d285d cmd/statreceiver: Add instance tag to influx metric
* f185adcf7 satellite/payments: fix projects list pagination
* 2601f25c9 web/storagenode: notification logic implementation
* 54e38b898 pkg/miniogw: gateway implementation with new libuplink
* 5342dd9fe go.mod: update uplink
* 0a8f268a7 go.mod: Update golang.org/x/crypto to fix vulnerability
* 4044b8eee storagenode/pieces: ensure chore is stopped before test ends
* fd5611fb5 private/testplanet: ensure server is closed in test
* ea970e45c satellite/payments: remove unused code
* 77f67a808 satellite/metainfo: add timeout for delete request
* a5afa4834 docs/blueprints: Add blueprint for creating a new state for storagenodes, suspension (#3778)
* 154ebdea6 satellite/api;repair: Have monkit pick the hostname as it's unique identifier
* e6da8d024 satellite/metainfo: use global limiter for DeletePieces Service
* dccca518d web/satellite: shortened token payments dropdown bug fix
* 4fb70b438 lib/uplink: use empty encryption access for restricting
* 5c508d143 web/satellite: new project button removed after creating one
* f4472b0b8 web/satellite: iframe checking added for login/register/forgotpass views
* 07da9d17c web/satellite: successful project creation popup reworked
* 4e1c00d37 jenkins: Update Jenkins to use -e POSTGRES_HOST_AUTH_METHOD=trust when running postgres.
* 47a52d3ce cmd/statreceiver: Change v2 known metrics to toml config file
* c3c69a088 cmd/statreceiver: downgrade tally metrics correctly
* 3e70a893d storagenode/{piecestore, contact}: report capacity to satellites if below specific threshold
* 85a2c7fda cmd/statreceiver: downgrade repair metrics from v3 to v2
* 1a84a00cc satellite/orders: Fix doc comments
* dbe8428f9 satelite/metainfo: return NotFound on delete non existing bucket
* 8bef560ab storage: delete unused code and lower visibility of static iterator
* 948589d38 private/dbutil/txutil: include details about retry attempts in error
* 892b190db satellite/admin: add project limit modification and authorization token
* 1b68d4548 jenkins: use go 1.13.8 (#3781)
* 1c1e62be0 web/satellite: comments update
* ef2f10149 satellite/metainfo: don't allow deleting non-empty bucket
* 8f2008568 storagenode/piecestore: clearer client cancellation error message
* 827da1ae2 satellite/payments: fail when trying to consume consumed transactions
* da58dc4a7 satellite/payments: increase batch size for transactions and account balance loops
* 6c6e2eb8b satellite/peyments: fix potential infinite loop in update account balance cycle
* 4e8695116 satellite/accounting: iterate over projects from tally rather than live accounting projects
* f9189f8d9 satellite/console: only create user with registration token
* 2ae997830 satellite/gc: skip first gc run
* 900cc4777 traces: fix memory leak for long running traces that aren't being collected
* 05a240050 storagenode: monitor available space and bandwidth
* ab660ccc8 statreceiver: Add http scheme to influx url
* ddbaaf451 statreceiver: Add influx v3 output
* 4e6a01f79 statreceiver: Rename metric_handlers to v2_metric_handlers
* 33e74f97d statreceiver: Add prod lua file to source control
* eeaaa8aa9 satellite/payments/stripecoinpayments: added ApplyInvoiceCredits
* c4fd84ad3 satellite/metainfo: Add metrics and traces DeletePices
* cea4c25f5 mod: bump common and uplink version
* 247255482 uplinkc: add ability to not set encryption restrictions with restrict_scope
* 76849558c satellite/gracefulexit: increase performance and tolerate higher error rate
* 37cf42a9a satellite/metainfo: overwrite zombie segments
* 5d6cb68cd storage/{cockroachkv,postgreskv}: detailed monitoring for list
* dbf46c4aa satellite/admin: administrative endpoint
* 2d2f5e1a7 satellite/satellitedb/dbx: remove typo in dbx file and format it
* f10b22eae accounting/tally: if delta < 0, delta = 0
* 33d696b09 storage/redis/redisserver: simplify redisserver creation
* b22bf16b3 satellite/overlay: add config flag for node selection free disk requirement
* dba647199 web/satellite: multiple storj coin transactions bug fix
* c59938479 Adding Monty Anderson to the CLA
* 961944f24 satellite/orders: Resolve storage node addresses to IP addresses.
* 429f08b4f satellite: add Admin peer
* 7d62b1cf3 cmd/storj-sim: don't wait for process to start indefinitely
* 426c8eb31 private/testplanet: add DeleteBucket method for uplink
* bd9cebda5 satellite/payments: fix transaction list pagination
* 208c05e3d Add metrics to track rate limit.
* ccd8b7f10 satellite/satellitedb: add benchmark for satellitedb setup and close
* d2fca7614 cmd/uplink/cmd: set exact argument counts
* 6cd86b214 web/storagenode: npm audit fix
* 984ed2673 satellite/payments: fix invoice project records pagination
* dc075eaa9 satellite/payments : deposit bonuses (credits) added
* 3331b443e satellite/metainfo: Delete all the piece of a storage node in one single request
* 3900dadaf satellite/overlay: find new nodes with ExcludedIPs
* c4a9a5d48 satellite/downtime: update detection and estimation downtime chores for more trustworthy downtime tracking
* 99c3ba5bb testplanet: log stack trace for error during creation
* bde302fdb jenkins: disable npm audit as its failing everything
* a90955ece web/satellite: billing history refreshing bug fix
* 68d5b1d6e build/jenkins: allow -rc release tags
* 55a3a9039 web/satellite: uplink CLI docs link behaviour on API keys page reworked
* 6679036ac web/satellite: unauthorize error handled
* 13903449c satellite/accounting: fix flaky TestProjectUsageStorage
* f151a0b9e installer/windows: add service dependency for storagenode
* 75355547c satellite/satellitedb: don't include GET_AUDIT and GET_REPAIR with chargeable BW
* 34f38bf6c mod: upgrade miniredis to latest
* 8147d6ccc web/storagenode: chart date bug fixed
* 7999d24f8 all: use monkit v3
* 17580fdf5 storagenode/pieces: Add test to cache store
* efa0f6d44 storagenode/monitor: set MinimumDiskSpace default to 500GB.
* 384bdf1f5 web/storagenode: code refactoring pt.1
* 91a480f5a Jenkins: add storagenode npm checks
* 109d733dd web/storagenode: npm dependencies updated
* dd9d18f15 upgrade drpc so that we have the monkit metric capability
* 9e5679fda storagenode/console/consoleserver: set content-type manually
* bd7894511 statreceiver: add v2/v3 splitter and downgrade
* eaf3318a5 diagrams: satellite graph per process
* d20db90cf private/dbutil/txutil: create new transactions for retries
* 71ff044ed storagenode/bandwidth: fix tests to not fail for 10 hours near the end of the month
* 03166d6be storagenode/piecestore: log available bandwidth and space on uploads
* 97d360afd satellite/satellitedb: use correct type
* ae3eae31d web/storagenode: chart point's hit radius extended
* d3fe122d9 web/satellite: project/billing dropdown bug fix
* 2cd233e1f web/storagenode: dashboard content header reworked
* 78d0868bc storagenode/pieces: Log error if cannot calculate piece size
* d0b427246 storagenode: fix global logger in tests
* 2968857e2 storagenode/pieces: Prevent recalculate from having negative numbers
* a181e0b62 libuplink: adjust tests to changes in encryption store
* 81d44f19e storage/filestore: ensure we bail on deleted folder without error
* 157b8c4d7 storagenode/pieces: accumulate errors in traversal
* 5a053483b storagenode/pieces: read trash from blobstore
* 4dafd03f1 storagenode: Prevent negative values in piece_space_used, migrate negatives to 0
* 00fc192f6 storagenode/pieces: Explicitly walk satellite pieces in SpaceUsedTotalAndBySatellite
* 21b65ca3b storagenode/storagenodedb: migrate to set total to content_size
* 006a2824b satellite/repair: lock monkit stats in checker and repairer
* 8dea4f52d satellite: add control panel
* 4e2bf8171 pkg/debug: add better title
* 81eddaa2c storagenode/monitor: reduce space requirement to 0
* 8c1985587 scripts/tests/rollingupgrade: explicitly set debug port for old satellite api during rolling upgrade test
* d10d6fd15 storagenode,satellite: ignore error on listening debug port
* 0b898c48a satellite/payments: coupons expiration logic fix
* 10be53860 storagenode: add pkg/debug support
* a2b2bc676 pkg/debug: implement control panel
* f237d7009 storagenode,satellite: use pkg/debug
* f833289e1 pkg/debug: separate debug endpoints to a server
* e0cb8037c satellite/projectusage: reduce usage limit from 5GB to 0GB
* 9bb7ceb65 satellite\payments: amount for coupons increased
* da5e408af storage: add DeleteMultiple method
* f4667426b satellite\payments: project limits for coupons increased
* 1ad4110d6 web/storagenode: notification red dot hide
* e641ff45a web/satellite: logout fix
* 7cbd83a39 jenkins: Remove os_arch from file inside zip file artifact
* 149273c63 satellite/metainfo: add cache expiration for project level rate limiting
* d30d2d920 satellite/metainfo: Adding Monkit Meters to the Request Logs
* e87886696 satellite/metainfo: Too many requests should have RPC status ResourceExhaused
* e549e3297 satellite/payments: fix promotional coupons
* f6641aa43 jenkins: use go 1.13.7
* 6b72bf92c satellite/payments: convert egress price to per byte basis
* 083b396c1 satellite/payments: allow floating point numbers for pricing
* a0c9f7f3b satellite/projectusage: reduce usage limit from 25GB to 5GB
* e66b3c9be satellite: remove repair worker from core
* e319660f7 private/lifecycle: implement Group
* a1948ed33 satellite/orders: add old method for CreateGetOrderLimitsOld to maintain compatibility with old versions of the uplink
* 3a0f6de66 cmd/uplink: list available accesses
* 54dbaaece satellite/orders: create as many orderLimits as needed to download a file
* 8ce9ce7f0 satellite/gracefulexit: wait for errgroup to return
* d8e3556a2 storagenode/preflight: wait for server to shutdown when tests are finished
* 90fc1922d satellite/metainfo: override bucket RS values with satellite config
* 3abb8c8ed Dont require an IP address being set
* 4e819c196 web/satellite: uncaught router exception fixed
* 2209924d4 satellite/satellitedb: use arrays and batch inserts for SaveTallies query
* 227e03dea satellite/satellitedb: insert using arrays
* d09bd4a74 satellite/satellitedb/dbx: regenerate with paged composite key fixes
* ae3f47147 cmd/satellite: close pointerDB properly
* f3fcbe256 satellite/metainfo: revert combine CommitSegment and CommitObject in batch
* 9bcb81108 web/satellite: verification email change
* f4317d257 mod: bump uplink and common
* ca32ffbfc satellite/metainfo: move deletion before upload to satellite
* 913a80c20 web/satellite: Fixing typos in website pages
* 9ea32016c storagenode/orders: fix typos in log messages (#3760)
* a0a94a9ac satellite/satellitedb: insert into reported_serials w/ arrays
* 90cf78e6f satellite/coinpayments: fix migration
* 5c68f4fc7 storagenode/gracefulexit: higher concurrency and shorter timeouts
* a6c6440ab satellite/order: decrease expire time from 7 days to 2 days
* 226bc4de3 storagenode/preflightcheck: enable database check by default
* e4cff1c93 storagenode/preflight: update allowed time  difference for preflight clock sync
* 26e33e7e0 satellite/gracefulexit: make orders with right bucket id and action
* 494fead7a satellitedb/orders: fix comma bug in SQL stmt
* f5c9597d2 golangci: Enable new linter added to last release
* f917fecc6 cmd/uplink: remove non-interactive mode from setup
* 524da4d2d web/satellite: word secret added to api keys
* b691a1756 web/satellite: payment banner second state
* 5e831f1f4 web/satellite: bakeoff UI prettifying
* d5a60aec5 satellite/metainfo: Delete segments in reverse order
* 416e5053a cmd/uplink: add note about share not-after/not-before flags formats
* 746c07cb9 cmd/uplink: improve usage for uplink subcommands
* 5a1838bc2 private/dbutil: retry single statements on cockroachdb
* 665ed3b6b satellite/satellitedb: fix issue with shared memory on range for bucket rollups
* 2f77ce48f private/testplanet: Add databases to testplanet.databases near creation
* a9e4c6f66 satellite/satellitedb/dbx: Remove bashism from gen.sh
* 16bb374de storagenode/piecestore: add large timeouts to read/write operations
* 89a148047 private/testplanet: shutdown databases in reverse order
* 44de90ecc storagenode/pieces: Rename vars and update comments
* 14fd6a9ef storagenode/pieces: Track total piece size
* 826fef0fa web/satellite: hyperlink on copy api key popup added
* 6bc989ae6 web/satellite: banner action area extended
* 81d4c05d9 web/satellite: hyperlinks updated
* 9414b8eb4 web/satellite: redundant scrolls fixed
* 5151dc981 web/satellite: popup's images changed
* fd84fa631 private/dbutil: rollback pending transactions on panic
* cadd727df cmd/uplink: just a silly spelling fix
* 62d378392 storagenode/peer: ensure contact.external-address and server.address is valid
* 40a890639 satellite/orders: Flush all pending bandwidth rollup writes on shutdown
* 960e10308 satellite/orders: Rename orders_write_cache to rollups_write_cache
* 0548c3f6b satellite/orders: RollupsWriteCache has a single method to reset cache
* fab58e9c1 cmd/uplink: hide advanced flags from output
* c6f94ce9e satellite/metainfo: remove support for boltdb based pointerDB
* 5a4745edd all: remove usages of testplanet.New
* 76fdb5d86 storage: add configurable lookup limits
* 38f707c0d storage/redis: Limit should not be applied as count.
* 5de4f6655 scripts/tests: change multisegment file to be 128kb
* 818242f45 storage/postgreskv: Reverting back to the venerable PG CAS
* b2c454dbd cmd/uplink: print user-friendly FATAL errors.
* 3b86917cc private/dbutil/pgutil: faster cockroach constraint finding
* fc2766eef private/testplanet: flatten migration for running tests
* d0041c94d pkg/process: increase default log level to warn
* 8b3db7032 private/testplanet: increase metainfo rate limit
* 80a6219c4 cmd/uplink: Create dir before saving config file
* 650245494 satellite/metainfo: move RS configuration to satellite
* 75314a436 satellite/satellitedb: fix roundToNextDay to handle timezones appropriately
* 0c0b47823 satellite: use require.WithinDuration
* d32626fe8 scripts/tests: update uplink config migration for test versions
* c353e8b10 run backwards compat tests with cockroachdb
* d289cdc6e web/satellite: project limits blinking fixed
* 21a5d70a8 satellite/metainfo: Rate limiting - API requests
* 38196b480 web/satellite: validation to reset password page added
* 6b6517aac web/satellite: date view on billing history improved
* 28d302c6f web/satellite: redundant text removed from api keys page
* cb827cbe3 cmd/{uplink,gateway}: update cli copy "http" -> "https"
* 877286785 satellite/metainfo: combine CommitSegment and CommitObject in batch
* 86f194769 uplink: adjust to changes in storj/uplink
* c636b0619 satellite/console: use cookie based auth scheme
* b678b55f8 satellite/metainfo: improve metainfo logging
* c1c878efc all: fix import groupings
* 0def7a9d2 scripts/tests/testversions;scripts/tests/rollingupgrade: update test versions script
* 33790e0f7 satellite/console: handle graphql errors properly
* 21f53e38d storagenode/storagenodedb/storagenodedbtest: pass ctx as an argument
* f3b4bf2b7 satellite/satellitedb/satellitedbtest: pass ctx as an argument
* 1279eeae3 private/tagsql,storage: fixes to context cancellation
* ba2fce814 satellite/satellitedb: better coupons query
* 10d932fd6 lib/uplinkc: fix test flakiness by setting MaxTimeSkew
* c4cbc6ff2 satellite/payments: promotional coupons generation functional added
* a4026f97b satellite: fix test time comparisons
* 3b55b50ea cmd/uplink: Add ability to generate named accesses
* 034f9845b storage: Plumb limit through storage backends.
* d5438036b {satellite,storagnode}/gracefulexit: reduce logging
* ee0293c21 private/dbutil/sqliteutil: add missing err check
* cf7b22c46 satellite/satellitedb: add missing err check
* 3cd584c00 storagenode/gracefulexit: move database test
* ca8ceff5e pkg/cache: fix expiration test
* 48303bc0a satellite/console/limits: add content-type to response, fix error
* c207cd08f satellite/satellitedb: gracefulexit, add missing Errs check
* 7bc76624c storagenode/storagenodedb: fix closing in-use database
* 1abfe4214 satellite: use tagsql
* 25b76fe63 storagenode/storagenodedb: use tagsql
* 59d06644b private/migrate: switch to tagsql
* 5fd833b10 private/dbutil: remove basic Query
* 23463e1dc Catch the Stage properly for Alerts (#3753)
* 14b43b7e9 storage/postgreskv/schema/data.go: Regenerate migrations that failed to update.
* f4097d518 satellite: reduce logging of node status
* 76a6b28b3 ci: use tmpfs for running tests
* 273eb66fa cmd/storagenode,storagenode/preflight: add config flag to disable storagenode database preflight check.
* d8368d0b3 satellite/payments: coinpayments add completed status, treat received status as pending, add balance for completed transactions only
* 614e04d05 storagenode/pieces: Cache inits trash info from db
* 6f2f97b31 storagenode\gracefulexit: broke worker deleteOnePieceOrAll into deleteOnePiece and deleteAllPieces and deletePiece
* 0c660f549 satellite/payments: fixed test so that it passes on non-UTC systems, simplified date comparison
* 5d80e22af private/tagsql: implement wrapper for sql.DB
* 8bbb9083f web/satellite: error on logout from team page removed
* 22af78b62 cmd/uplink: fix 'must specify access' for old uplink configurations
* 491cd8d8a scripts: automated test for testing uplink share command (#3736)
* 955abd929 satellite/satellitedb/orders: add multi row upserts to process orders
* 409d4123b Add proper Pathdata Index (#3750)
* cd48dc369 satellite/satellitedb: Remove unused indexes
* 47bb7a7a8 satellite/satellitedb/dbx: regenerate with default support
* 696d98a23 satellite/satellitedb: fix nitpicks and timestamp issue found in review
* e115bc190 cmd/storagenode;storagenode/storagenodedb: add preflight database check for storagenode
* b6f1a91c6 scripts/testversions,rollingupgrade: remove encrytion key
* 4424697d7 satellite/accounting: refactor live accounting to hold current estimated totals
* 81d53b809 storagenode/storagenodedb: fixes to row handling
* 7d79aab14 satellite/satellitedb: fixes to row handling
* 0c365d157 scripts/testversions: replace apikey with access
* 1e77cb88e scripts/rollingupgrade: replace apikey with access
* f42851b1a satellite/satellitedb: remove the big honkin mutex
* 78c6d5bb3 satellite/satellitedb: reported_serials table for processing orders
* 9da16b1d9 satellite/satellitedb/dbx: name the package dbx
* dc1db75ab web/satellite: forgot password error message removed
* db8aee080 satellite/contact; storagenode/preflight: add clock check on startup for storagenode
* 07c2824d9 storagenode/gracefulexit: fix exit-status command output
* 6dc948da4 satellite/metainfo: Create service for deleting pieces
* 02ed2b5a1 satellite/metainfo: disallow cancelling deletion request
* 08f63614b private/context2: add WithoutCancellation
* 19d318ea9 web/satellite: password strength check added to change password popup
* f9b2af934 cmd/gateway: Fix entrypoint setup command
* 64fb2d3d2 Revert "dbutil: statically require all databases accesses to use contexts"
* c01cbe013 satellitedb: save out all db-touching traces
* 8e242cd01 dbutil: statically require all databases accesses to use contexts
* 5af1f9e6d storagenode/{piecestore,storagenodedb}: use context in queries
* 64f056bee private/dbutil/sqlutil: use context in queries
* df9e53ea0 private: ensure we don't eat the underlying error
* d80cfeb4a storagenode: ensure we don't eat the underlying error
* 23e266432 storagenode/inspector: return rpcstatus
* 86093d094 postgreskv: drop not null on buckets
* 41d5e8630 satellite/payments: coupon addition removed
* cd4ff0722 private/testplanet: use defaultInterval
* 6c4e3b64d cmd/uplink: 'uplink access inspect' for displaying access fields
* c51f9a22e web/satellite: recovery password fonts and styles fixed
* 224025d71 web/satellite: password recovery page title fixed
* c8ccd26e0 cmd/uplink: import imports 'access' into existing configuration
* a57ce18f5 satellite/payments: coupons, coupons usage, invoice generation with pricing model applied
* ee87846f0 satellite/contact: add placeholder for GetTime endpoint
* 3b99f0378 satellite/orders: add monitoring to bucket bandwidth cache operations
* 4950d7106 satellite/orders: Add write cache for bw rollups
* e1ba3931e postgres2: use cockroachkv impl against postgres
* dd234aad4 jenkins: bump go to 1.13.6
* 71ec0ad37 satellite/satellitedb: add big honkin mutex to ProcessOrders
* 6bf40f3e5 web/satellite: npm packages updated
* 172918596 cmd/uplink: output cleanup of uplink share command
* b9740f0c0 storage/cockroachkv: add ctx argument
* ff267168c private/migrate: add ctx argument
* 24958bd7d satellite: add ctx to DB.CreateTables
* 0835b9024 private/dbutil/pgutil: add ctx argument
* c7b846589 private/dbutil/sqliteutil: add ctx argument
* bcc23f686 Satellite/orders: remove allocated bandwith from storagenode_bandwidth_rollups
* 36db00b2b cmd/uplink: don't require setup or import if --access is set
* 4aef0e382 satellite/satellitedb: only reject orders if row not found
* 5a1b2f49f storage/cockroachkv: add application name to the db connection string.
* 67b5528dc web/satellite: tab name on referral page fixed
* b579c260a cmd: rename "scope" flag to "access"
* 57a31ac4e web/satellite: resend email timer fix
* 131c4d94f {cmd/uplink, lib/uplink}: change RS total to 110
* 455e14add Jenkins: only run versions test on master commit
* 77fd41a02 satellite: add an expiring lru cache around api keys
* 6b1829f3c satellite/downtime: new chore estimates downtime
* cf19e141e storagenode/notifications: return unread count and fix json id, list-notifications method fix
* b252fcdc0 mod: bump uplink dependency
* a4e5c1887 satellite/payments: mock methods added to endpoint to match pb PaymentsServer
* d3d75a597 satellite,storage: clean global ctx usage in tests
* 5cd605831 satellite/metainfo: Add back-pressure mechanism DeleteObjectPieces
* 8d8d57c3b mod: update sqlite module to v2.0.2
* 76ee8a1b4 satellite: remove UptimeReputation configs from codebase
* ebeee5800 storagenode/gracefulexit: remove satellite entry when node fail precondition
* 05e4a8665 satellite/metainfo: Close client DeleteObjectPieces
* 922c43f92 satellite/metainfo: Close client DeleteObjectPieces
* 082ec8171 uplink: move to storj.io/uplink (#3746)
* cf2128d3b uplink: avoid cyclic dependency to storj.io
* 00c0c51b1 cmd/uplink: fix TestSetGetMeta flakiness
* 6a4e4e030 web/satellite: dates to utc; limits bug fix;
* c740b82e6 satellitedb/dbx: remove sed usage for bash script
* 6b21334c4 satellite/satellitedb: use txutil.ExecuteInTx in dbx WithTx()
* 3f38a7ed1 lib/uplink: revert library move
* 0c88a7b47 private/migrate: use transactional helpers and not Begin()
* 0c5e38143 satellite/console: use transaction helpers in consoledb
* 22b6e9220 satellite/satellitedb: use transaction helpers in irreparabledb
* 93f3650cf ci: add build number to postgres and redis docker containers for versions and rolling upgrade tests
* 723ed2329 satellite/satellitedb: use transaction helpers in orders
* 05e2b0ca8 mod: bump common and some other packages
* 6607dacb1 uplink/metainfo: move listing to kvmetainfo
* 027e3d4f6 satellte/metainfo: Avoid a noisy warning
* e232042e8 satellite/metainfo: move old API tests to separate file
* fb4b11d13 scripts: remove old scripts (#3742)
* 6861f28bb release/script: allow RC release tags
* 9618959f1 satellite/metainfo: Use errs2 IsRPC helper function
* 1cb0f80a8 satellite/gracefulexit: dq node on exit fail
* 4a26fb5bd satellite/satellitedb: don't use crdb.ExecuteTx with postgres
* 0135852a0 storage/postgreskv: use transactional helper
* f3aee1b75 satellite/satellitedb: use transaction helpers in containment
* 2549c601e all: bump storj.io/common dependency
* 6c2e4cc0c satellite/overlay: Return NodeLastContact instead of a node dossier from overlay.GetOfflineNodesLimited
* 4203e25c5 satellite/satellitedb: use transaction helpers in overlaycache
* b072e16ff satellite/satellitedb: use transaction helpers in peeridentities
* eb81879d4 satellite/satellitedb: use transaction helpers in usercredits
* 623184242 private/dbutil: add WithTx transaction helpers
* f41d44094 all: reduce number of log messages
* a33734bee satellite/satellitedb/dbx: add cockroach driver type
* ea84af578 scripts/tests/rollingupgrade: create new test files for final upload stage
* 91947311f ci: always try to pull latest image
* 07a1702f4 scripts/tests/rollingupgrade: fix test-versions.sh path referrence
* 71c5c2213 scripts/tests/testversions: make binary installation and upload/download running in parallel
* 80b41af8f satellite/metainfo: Fixed bug that discarded context cancellation errors
* fc4ea2869 satellite/metainfo: Return ErrObjectNotFound
* 29fe206b9 satellite/gc: add timeout to retain requests
* 828d0b998 pkg/server: set TCP_USER_TIMEOUT and monitor leaked conns
* 0038abb51 private/testplanet: use redis for live accounting
* e1e7cebe4 satellite/metainfo: added rate limiting support to the metainfo loop.
* 05b406e99 satellite:{downtime,overlay}: Implement offline node detection chore
* 5aac77c2a Slack the build team instead of everyone (#3739)
* 389567fc9 satellite/console: add credit card charges to billing history
* 325790703 installer/windows: batch file improvements (#3441)
* 0cc7056a9 satellite/console: convert dates to UTC in advanced usage reports
* 38eff6069 satellite/metainfo: adjust old API test to new API
* e34ac3ef3 ci,scripts/tests/rolling-upgrade: run rolling upgrade test on private jenkins
* aecea820f scripts: add rolling upgrade test script
* 8859c3623 satellite/{downtime,contact}: Add CheckNodeAvailability for use within the downtime tracking chores.
* 6098f606a lib/uplink: move to uplink
* ff74b44c5 satellite/overlay: Add ability for overlay to get offline nodes ordered by last checked time
* c3b58f165 satellte/metainfo: Make BeginDeleteObject to delete pieces
* 71ea2b0dc uplinkc: object test fix
* 105a9a484 web/satellite/tests/unit/common: Fix hardcoded year check
* 3e873158f ci: increase cockroach max-memory to reduce flakiness
* 3528c56c6 satellite/satellitedb/satellitedbtest: skip unconfigured db
* e03d3fb57 uplink: move configs to cmd/uplink/cmd
* 2680bae88 private/testplanet: remove dependency to uplink
* a35eb8027 lib/uplinkc: do some checks around conversions and allocations
* aa3e183c2 satellite/gracefulexit: add ge eligibility check
* 726615537 satellite/metainfo: List segments manually limit check
* bb3baf5a4 satellite/satellitedb: Add nodes_offline_times table for downtime tracking
* 758fe35ab storagenode/orders: adding jitter to sending (#3725)
* 82ee13b00 Update Coverage URL (#3737)
* e99bdac94 web/satellite: ux bugs fixes
* 090603b8e web/storagenode: wording on DQ info message updated
* 059537c16 satellite/metainfo: Add new TODOs & remove old ones
* 6615ecc9b common: separate repository
* 115b8b0fc storagenode/piecestore: delete several pieces in a single request
* 7df3c9efc cmd/uplink: use arguments in share command as allowed path prefixes
* 7d1e28ea3 storagenode: Include trash space when calculating space used
* 3849b9a21 pkr/rpc: remove RawConn
* d55288cf6 pkg/rpc: replace methods with direct calls to pb
* 006baa9ca pkg/rpc: remove drpc aliases
* 6f1eaef8d cmd/uplink: Pass -- in tests to avoid treating generated arg strings as flags.
* acb4435a6 satellite/satellitedb: improve Cockroach migrate test
* 31fbdcc8f pkg/encryption: better EncodingDecodingStress
* 98243360d cmd/storagenode-updater: faster update test
* 6e71591b9 satellitedb;storagenodedb: remove unnecessary use of DB transactions in graceful exit
* 6fc009f6e uplink/eestream: move Pad to encryption package to break dependency to eestream
* b959ccbae satellite/gracefulexit: Use proper rpc status codes for disqualified nodes and too many connections
* e47ec84de storagenode notification service and api added
* ea455b6df all: remove code to default to grpc
* 62c58f4a9 satellite: consistent report range arguments
* 9e4d83317 private/testplanet: use default interval
* c6854accd scripts: add test-versions stage to private Jenkins
* 46c8d2e9c private/testplanet: Wait until peer ends when closing it
* 858f349a1 lib/uplinkc: add user agent
* 9f1de67f6 web/storagenode: notification markup
* ef8bc8832 ci: use external repository
* 389d1821e uplink/paths/encryption: support commandline argument to override path cipher to be urlsafe base64 for lists and deletes (#2855)
* ffbc43d17 satellite/satellitedb/dbx: monitor the database calls
* d5d0c442a satellite/accounting/rollup: Use lastRollup as zero-value
* 298692324 build: update golang to 1.13.5
* 2daf24a1e private/testcontext: remove version dependency
* 9ed9d3516 pkg/peertls: move tests requiring redis or bolt
* 1eaf9e9ed pkg/storj: move non-common types
* 71d6eef0d pkg/storj: remove unnecessary interfaces
* de0aefaa6 web/satellite: card stars align fixed
* d5c5b57fa satellite/db: enable DeleteTallies
* f8d086463 satellite/metainfo: use KnownReliable in DeleteObjectPieces
* af24581ac satellite/audit: do not report offline to overlay (#3547)
* 366f4b949 satellite: Create method for deleting pieces of SNs
* 5ee1a0085 satellite/overlay: filter reliable nodes from a list
* efa08d408 storage/cockroachkv: use different batch size for recursive iteration
* 1a625887e scripts: Add script to automate testing against all highest release points from major releases starting from v0.15.4 for uplink
* 67892b4ad private/testidentity: clone identities for each version test
* 69dc514fc lib/uplink: fix test data race
* afe05edff {storagenode,satellite}/gracefulexit: ensure workers finish their work
* 7a36507a0 private/testcontext: ensure we call cleanup everywhere
* 7455ab771 pkg/peertls/tlsopts: move test that requires testplanet
* b04f9996c pkg/rpc: move test that needs testplanet
* 04d459a4a web/satellite: buckets page adaptation fixed
* 0e12c2343 uplinkc: add restrict_scope function (#3724)
* 08947e177 storagenode/garbagecollection: enable in production
* a4f9865b4 satellite: adds and enables cockroachdb compatibility for tests
* a47d7ac89 scripts: Add script that filters postgres plaintext backup to cockroachdb compat.
* 8242eecea Adding benchmarking script that reports response times.
* 0008aebf8 pkg/rpc: Change drpcheader to save a packet
* 2f7465c29 private/dbutil: register "cockroach" as sql.DB driver
* b5ddfc6fa satellite/satellitedb: unexport satellitedb.DB
* 6fb9b4b23 ci: remove Jenkinsfile.drpc
* 6dc1249c0 web/satellite: billing banner changed
* c5116cb2a satellitedb: fix migration cockroach test
* 53d9bc453 storagenode/notifications: db created (#3707)
* 11db70906 web/satellite: project limits
* ab777e823 do not update pointer for failed audits
* 45f8bb508 uplink/storage: remove unused Meta methods
* c2ea75208 storagenode/orderdb: fix db lock
* cb8949656 storagenode/trust: wire up list into pool
* 9ac2c4d81 cmd/segment-reaper: Add test cases from/to processSegment
* fdd86d17a lib/uplinkc: use SliceHeader instead of [1<<30]
* 770123de1 satellite/metainfo: Improve docs & use common funcs
* ff854b395 Use Postgres 9.6 for Jenkins builds (#3729)
* f94cc6498 web/satellite: cutted storj logo fix
* 2867b6a46 storagenode/trust: list implementation
* 086b00d55 Remove tasks from tight download read loops
* 767d8eb77 web/satellite: deposit link opens directly
* 03fa0150b satellite/metainfo: Trace endpoint getPointer method
* f5d26a178 web/satellite: billing page behaviour  (#3711)
* f659d98a4 satellite/payments/tokens: return checkout url on new deposit (#3696)
* 77839dd41 satellite/console: project usage limits api (#3702)
* 9d8f6a6d3 web/storagenode: charts titles aligned
* e1d99522b cmd/segment-reaper: test processSegment for single project
* 0253eff2e web/satellite: add tracking event for segment.io (#3641)
* 6ce22be74 cmd/storj-sim: make initial provisioning nicer
* 820e109cd adding Igor Gaidaenko our new team member! (#3726)
* b29456984 segment-reaper: fix for not analyzing last project in detect command
* 0d1ba7bc0 satellite/metainfo: Return error misusing func
* 958772467 cmd/storagenode-updater, pkg/process: Fix logging timestamp
* fb8e78132 storagenodedb: reenable utccheck in tests
* 5ed9373db storagenode/trust: source entry cache
* 715d97e3d storagenode/trust: rule and excluders
* 5d8d9cd89 projectLimit error message changed (#3718)
* 23df647a1 pkg/rpc/rpcpool: add idle expiration to connections
* 94651921c storage/testsuite: pass ctx in to bulk setup methods
* 72d407559 satellite/metainfo: don't leak error implementation detail (#3722)
* 218f436f6 uplink: remove dependency to storage/ (#3720)
* 4f282921c jenkins: run storj-sim integration tests with cockraochdb  (#3723)
* d8a8f92e3 private/dbutil/cockroachutil: keep crdb connstr for tests
* 6fae361c3 replace planet.Start in tests with planet.Run
* 8cf1aa6e4 satellite/accounting: fix project limits migration (#3717)
* f56107fc6 uplinkc: add download_range function (#3704)
* 011bb5d55 segment-reaper: test from/to detect parameters
* e961b1840 cmd/segment-reaper: Implement unit test for observer (#3658)
* de30d232c uplinkc: add more functions to manage Scope (#3694)
* 71b58edb2 satellite/repair: decrease repair interval
* 48da8baab storj-sim: work with cockroach:// urls for satellite databases
* eb52ac623 storagenode/trust: source implementations
* 6ab72a6e7 satellite/gracefulexit: enable graceful exit in production
* 7d0aadfec storagenode/trust: satellite URL implementation
* c3c02bec3 satellite/satellitedb: reset storage node reputations to re-enable disqualification (#3693)
* 9d1faeee5 storagenode/garbagecollection: increase MaxTimeSkew to be higher than satellite MaxCommitInterval
* 56a3b62be satellite/satellitedb: ensure migration tests run (#3706)
* 6ef359aba web/storagenode: blurred checks hint added for all satellites state (#3709)
* fe5cfeb61 web/storagenode: disk space chart's tooltip size fixed (#3684)
* c6776ae6b error messages fixed (#3712)
* fa5288c25 satellitedb: bucket search fixed (#3594)
* ea92c6860 web/satelllite: default cc delete disabling (#3695)
* ffd570eb0 uplink/metainfo: remove additional GetObject from download
* 27462f68e cmd/segment-reaper: add detecting and printing zombie segments (#3598)
* 8d49a99ad uplink/metainfo: Fix doc comments
* b17d40ffd segment-reaper: delete command logic (#3660)
* c400471bb satellite/metainfo: Fix some docs comments
* 1df7b360d satellite/metainfo: Use cockroachdb client for metainfo db
* f15192ea4 storage/cockroachkv: initial client implementation
* 8e9532dbd satellitedb/repairqueue: use errs.New() instead of fmt.Errorf() to retain stack trace
* 378b863b2 private,satellite: unite all the "temp db schema" things
* 97fa000ca web/satellite: usage data converting implemented (#3673)
* 28a66747c satellite/satellitedb: cockroachDB compatible transaction in offersDB.Create (#3686)
* 7c5f777a4 satellitedb/repairqueue: runs a different implementation of the query within Select() for postgres vs cockroach
* 7af42e3c1 satellite/metainfo, satellite/repair, uplink/eestream: add metric for download failed due to not enough pieces available (#3665)
* 850c35808 private/dbutil/pgutil: make QuerySchema work on crdb
* e9eff5447 satellitedb/attribution: updates valueAttrQuery to work for both postgres and cockroach
* ecb960f50 private/dbutil: distinguishes between db drivers and implementations to allow for different implementations of SQL queries.
* b77982619 uplinkc: basic support for Scope (#3689)
* c00f68893 format db name in query to match actual db name (#3688)
* 9420fa9fc satellite/gracefulexit: Add graceful exit completed/failed receipt verification to satellite CLI (#3679)
* 2461ccd46 pkg/private/fpath: subsume AtomicWriteFile
* 4d2881b71 satellite/satellitedb: return latest coupon time as UTC
* e0b9b5b31 satellitedb: fixes for cockroachdb compatibility (#3682)
* d69482e93 satellite/metainfo: Improve piece hash validation (#3671)
* 52851026a satellite/console : remove code for creating user credit in CreateUser, skip TestUserCredits (#3680)
* b1fa7cdfb satellite/satellitedb/satellitedbtest: flag to run cockroach tests (#3674)
* 90b631c8b web/satellite: referral links and registration (#3678)
* 526a126a5 Satellite: log the project information for upload, download, and list activities (#3651)
* 24cacc8ad installer/windows: fix minor typo (#3676)
* 42c61138e storage: Improve doc comments delete methods (#3591)
* ae60b6533 web/satellite: titles styling fixed (#3662)
* ae05fa087 satellitedb/accounting: fixed query (#3672)
* d9a23b872 web/satellite: password strength implemented on register page (#3669)
* 2a13a7764 web/satellite: bundles size reduced (#3667)
* 0f523e82a missing line returned (#3670)
* b10eee6e7 satellitedb/accounting: project usage optimization (#3668)
* 7e9b633dd satellite/mailservice: move logging to send rendered async to cover template parsing (#3654)
* 18a5e614d satellite/web: add segmentio plugin (#3405)
* f83837bb0 web/satellite: billing banners (#3649)
* 0bf7ac5b2 cmd/segment-reaper: Use bitmask type in observer (#3661)
* 7288d0178 cmd/segment-reaper: several refactorings (#3637)
* a6235d396 storage/filestore: Monitor when we open files in trash
* 756b9b9e2 satellite/payments: coupons and coupon usage (#3648)
* bf97ef06f storagenode: Add new endpoint to receive satellite requests for… (#3590)
* 66f1a1680 add completion receipt to exit-status cli command on storage node (#3650)
* bba92911c fix calcuation of durability ration (#3656)
* cf13cf6e0 web/satellite: referral links (#3655)
* 854e5507a crdb uses namespaced db for each test (#3646)
* 56f8fd2dd storagenode/pieces: Add EmptyTrash functionality (#3640)
* 038ac5860 web/storagenode: minimal allowed version view implemented (#3583)
* 2502c2615 Jenkinsfile: remove whitespace (#3244)
* 7abad3c6b refactor sql to be compatible with pq and cockroach (#3647)
* ce49c1041 web/storagenode: info containers of remaining bandwidth/diskSpace replaced (#3653)
* b763c5319 web/storagenode: disk space chart's tooltip reworked (#3652)
* 1e33995a8 cmd/segment-reaper: skeleton command for zombie deletion (#3599)
* 36fead009 satellite/metainfo: add UserAgent support to endpoints (#3548)
* 59385aff6 web/storagenode: ingress chart implemented (#3618)
* 79a4fff6c satellite/referrals: set up referrals service and http endpoints (#3566)
* 17b057b33 satellite/audit: monitor worker function
* 01895d8bd lib/uplink: explain safe versions (#3644)
* 8a002e8c8 satellite/accounting: separate project limit from project entity (#3632)
* 8234e24d1 web/satellite:  token payments logic (#3581)
* 1aa2bc0a8 satellite/metainfo: reduce pointerDB access for CommitObject (#3589)
* 48a557eb2 satellite/metainfo: Fix misspelling in comment (#3636)
* 031ba86de argon2: choose a steady parallelism value (#3630)
* 388f33b84 satellitedb: add support to testplanet for cockroachdb (#3634)
* e5934cc92 satellite/console/auth: return in error handle added (#3639)
* 9af97d366 Make sed a little more cross platformable (#3629)
* 68cd6f4a6 web: ms edge support bug fixed (#3638)
* ded7f6e2e web/satellite: registration/welcome message fixed, usage-report url fixed, storj-sim fixed (#3622)
* 37c379f7d web/satellite: fonts changed to Inter (#3620)
* bb16c07dc storagenode/updater: read identity location from storagenode's config.yaml (#3607)
* 77d4add10 cmd/segment-reaper: Implement bitmask type (#3626)
* 8842b0c25 storagenode/gracefulexit: improve logging (#3633)
* 63e51df9a private/testplanet: add a mock referral manager server into testplanet (#3631)
* 1339252cb satellite/gracefulexit: refactor concurrency (#3624)
* b7a8ffcdf pkg/pb/referralmanager: update to add satellite ID to Get Tokens request (#3625)
* d96df2691 satellite/metainfo: improve Loop comments (#3595)
* 273977176 storagenode: add bandwidth metrics (#3623)
* 976881f72 satellite/console: Add security headers (#3615)
* 87c7a2ff4 satellite/payments: token deposit accept cents (#3628)
* c72c44356 satellite/payments: add cents values to transaction info (#3614)
* 2030d6765 web/satellite: charges summary fix (#3619)
* 6a4389d3e satellite/console: apiKeys case-insensitive search added (#3621)
* b995406ff satellite/satellitedb: separate uuid creation from db layer (#3600)
* 40012e579 satellite/metainfo: continue instead of return (fixing my bad advice)
* c2e605e81 satellite/metainfo: Don't return error in loop when path has less than 4 parts (#3616)
* 6aeddf2f5 storagenode/pieces: Add Trash and RestoreTrash to piecestore (#3575)
* 6d728d6ea storagenode/collect: delete piece 24 hours after expiration (#3613)
* b5707d1e5 scripts: make update-tools.sh more verbose (#3572)
* 61c8bcc9a web/storagenode: egress chart implemented (#3574)
* 49694b2c7 web/satellite: successful reset password page styling bug fixed (#3612)
* 9ca547acb web/satellite: project charges (#3611)
* 16f0e998b web/satellite: project selection dropdown bug fixed (#3605)
* fe8d556b4 add sourcerer hall-of-fame to README
* da39c71d3 storagenode: add new metric satellite.request (#3610)
* e9c3194c8 satellitedb: merge migration into one step (#3551)
* 3d94c3fc9 satellite/payments: stripe client use satellite logger, give credits when applying transaction to balance (#3603)
* 8653dda2b satellite/audit: do not contain nodes for unknown errors (#3592)
* 8e1e4cc34 piecestore: Fix invalid comment and typos (#3604)
* 46ad27c5f docs/blueprints: Change adapt endpoint by creating (#3601)
* 5964502ce uplink/metainfo: remove GetObject from download Batch (#3596)
* 24318d74b storagenode/console: show satellite url in satellite selection (#3602)
* b5b4543df docs/blueprints: Deletion performance (#3520)
* 53176dcb0 pkg/rpc/rpcstatus: do not depend on grpc/drpc build mode
* f3b20215b pkg/{rpc,server,tlsopts}: pick larger defaults for buffer sizes
* 3a0b71ae5 pkg/pb/referralmanager: update RedeemTokenRequest, rm ReserveToken (#3593)
* c52c7275a satellite/repair: reduce upload timeout (#3597)
* 8ea09d555 cmd/segment-reaper: add iteration over DB (#3576)
* ec41a51bb metainfo/loop: refactor iterator to be reusable (#3578)
* fdcf32846 short name field removed from registration page (#3584)
* 3fe518d54 satellite: added ability to inject stripe public key post build (#3560)
* f4a626bbc storagenode-updater: re-enable auto-update storagenode-updater (#3588)
* 8b3444e08 satellite/nodeselection: don't select nodes that haven't checked in for a while (#3567)
* ecd2ef4a2 all: build release fully dprc and test in mixed mode
* 2524cc5d4 pkg/pb: remove unneeded proto imports (#3585)
* b79fad659 web/satellite: redirecting condition for universal login page added (#3550)
* 91859f1cc satellite/metainfo: fix broken test (#3580)
* a36e9b504 satellite/payments: switch to using STORJ tokens (#3582)
* 53c6741ba  satellite/payments: add API for retrieving conversion ratio, convert tokens to USD before applying to balance (#3530)
* ead8af3c1 storagenode/updater: recover command in Service Manager (#3425)
* a8e4e9cb0 satellite/payments: project usage charges (#3512)
* aa7b5b7c5 satellite/metainfo: avoid large conditional block (#3579)
* 0d35505fe SNOboard/console: router changed for gorillaMux, caching added (#3577)
* bc16cb5d2 libuplink: remove additional GetBucket for upload/download (#3568)
* c193dee9a uplink/storage/streams: Fix upload error clean up (#3555)
* 2166c2a21 storage/filestore: Add Trash and RestoreTrash to Blobs (#3529)
* ee6c1cac8 private: rename internal to private (#3573)
* f1d0d0d68 web/satellite: static pages styling fixed (#3549)
* d2a8ab5d7 pkg/pb: add referral manager protobuf definition (#3561)
* 491292263 cmd: skeleton for segment-reaper detect command (#3562)
* 1a54007f1 storagenode/storagenodedb: dont log opening of each database (#3571)
* b2a7a9f4c scripts: add script to update tools (#3570)
* b56cc2171 cmd/storagenode-updater: make updater test windows compatible (#3542)
* 77ed04742 installer/windows: unit test mocking file system (#3543)
* 1e64006e3 lint: add staticcheck as a separate step (#3569)
* d5963d81d storage/postgreskv: fix ultra-slow nonrecursive list (#3564)
* bd89f51c6 Keep v0pieceinfo database isolated (#3364)
* db8294cfb storagenode/gracefulexit: get hash and limit using original piece ID (#3557)
* 0ff1547d0 certificate: close authDB in tests (#3559)
* 1a9757a7f satellite/gracefulexit: add count for order limits sent from satellite to exiting node (#3544)
* 4c822b630 {certificates,pkg/rpcstatus}: improve error logging (#3475)
* a72bf6c25 pkg/rpc: generate drpc/grpc tags correctly (#3556)
* 0d0b5a449 storage/filestore: monkit event for delete queuing (#3507)
* 2f19b0a80 Unhook jenkins from the staging cluster while we sort it out (#3553)
* 0c025fa93 storage/: remove reverse-key-listing feature
* 28a70200d Add stage name to slack error message (#3554)
* 168f72d37 Initialize math/rand default source (#3552)
* 39021b964 lib/uplinkc: add additional fields to UplinkConfig (#3526)
* 013e0d94b pkg/rpc: ensure connections are quickly closed
* a01c48f4a satellite/rewards: use base32 instead of base64+json (#3522)
* 9797f8c49 satellite/console: service error types added, error handling fixed (#3538)
* 70c7ee842 satellite: adding proper headers to api responses (#3489)
* b2d2377d2 web/satellite no buckets page reworked (#3419)
* 7cc4217fe cmd/storagenode-updater: simplify and reorder update sqeuence (#3487)
* b7a04eb88 web/satellite: saving selected project in local storage implemented (#3470)
* e99af1822 web/storagenode: date and data formating fixed (#3519)
* c1ae8c332 satellite/console: auth API error handling refactored (#3540)
* 7355065dc pkg/{cfgstruct,identity}: replace seperator in default values when path tag set (#3498)
* 36311a3a0 satellite/console: add token deposit API, populate billing history with transactions (#3500)
* 89efd17f4 docs/design: zombie segments cleaner (#3461)
* 6516471cb uplink/storage/streams: Upload loop operations reorganization (#3429)
* 4b85d3d73 internal/testplanet: better error message when postgres is not defined (#3539)
* f3e803203 lib/uplinkc: add clarifying comments to download_read (#3525)
* 20eef5a20 sorting header on api keys page styles fixed (#3537)
* e065ad001 detailed usage charges info markup added (#3528)
* 20623fdc9 Increase min required difficulty to 36 in signing service (#3535)
* 9dce3dc94 installer/windows: unit tests for C# custom actions (part 1) (#3521)
* 69b0ae02b satellite/gracefulexit:  separate functional code in endpoint (#3476)
* 06a52e75b web/satellite: name hint (#3515)
* 9ce6dad31 web/satellite: usage report date selected date range formatted (#3518)
* 7ef0bbe00 credit cards icons selection added (#3527)
* 6331f839a satellite/gracefulexit: not allow disqualified node to graceful exit (#3493)
* f3dccb56b satellite/gracefulexit: Check if pointer has been overwritten or deleted before sending transfer message. (#3481)
* e4a220347 uplink: Suppress one metainfo call on delete (#3511)
* 53a060cdf cmd/storagenode-updater: TODO cleanup (#3486)
* f62107d3e pkg/rpc: fix grpc dial timeouts (#3517)
* 68a779006 satellite/gracefulexit: select new node filtered by Distinct IP (#3435)
* aaead2a29 makefile: bump drpc and build storj-sim binaries like release (#3506)
* 23c556ae1 satellite/rewards: fixes from review comments (#3495)
* ebcd37c57 storagenode/contact: fix connection leak with contact checkin
* f37f28020 bug with not compressing to small css file fixed (#3516)
* cc032d315 satellite/metainfo: fix some uses of metainfo.Delete (#3513)
* 3b78addb2 Metadata Access from Uplink CLI (#3310)
* fd9f860fd token error code fixed (#3514)
* 2f13c0adf storj-sim: add rollout seed to avoid warnings (#3508)
* 994a69cfd jenkins: use lower segment size for back comp test (#3097)
* 8a3b0ccbe added ability to close menu on small windows by clicking outside (#3484)
* e7a188e3e web/storagenode: satellite dropdown enhanced (#3490)
* e6c7b8180 web/satellite: tabs title change (#3496)
* df923c2a8 internal/dbutil/pgutil: retry create schema on a concurrent call (#3510)
* 3b18c864d test/backwards-compatibility: fix port change (#3509)
* 0885eaa18 Storage Node Satellite Selection design (#3353)
* f32ff4ec7 release: use /usr/bin/env echo (#3505)
* 7eb6724c9 logging: unify logging around satellite ID, node ID and piece ID (#3491)
* 12cead8f7 release: opt more binaries in to grpc (#3503)
* 0b32690d0 satellite/peer: add payments config (#3488)
* 841674e0e add goversioninfo to build-dev-deps make target for windows version info (#3501)
* 7cdc1b351 satellite/audit: do not audit expired segments (#3497)
* 2eb0cc56f satellite/gracefulexit: Check if node already has a piece in the pointer (#3434)
* ce89c44c0 Rebuild faster by removing weird make calls (#3502)
* 0e3f0eeb1 web/satellite: project description enhanced (#3494)
* 466cc8ab0 web/storagenode: text selection on specific elements disabled (#3492)
* 257d3946d storagenode/gracefulexit: allow storagenodes to concurrently transfer pieces for graceful exit (#3478)
* 78fedf5db satellite/gracefulexit: handle piece not found messages from storagenode (#3456)
* 19a59c9d5 Fix hyperlink's background color on Welcome dialog (#3483)
* d6b5d49ff Change "Ethereum wallet address" label to "STORJ payout address" (#3482)
* ab5c623ac cli: should return non-zero code for error (#3469)
* 35edc2bcc satellite/payments: invoice creation (#3468)
* 9c59efd33 satellite/rewards: ensure that partner information is asked from a service (#3275)
* bee1acef4 web/satellite: info button added on PM page (#3449)
* 1a0757ebf web/satellite: auth error messages fixed (#3426)
* c27958427 web/satellite: blinking before fetching PMs, APIkeys, buckets removed (#3432)
* 169d9d7bf web/satellite scrollbars to navigation and billing page added (#3442)
* 2dcfa18f3 web/satellite: api keys copied state removed (#3445)
* 21f3a68de web/satellite: removed redundant functionality from billing page (#3444)
* c81f90b1e web/satellite: added functionality to remove yourself from someone else's project (#3462)
* 4ef98b05e datepicker tests issue resolved (#3473)
* 75db47f13 web/satellite: Pages lazy loading added (#3467)
* def3dcbaa satellite/audit: increase timeout to 5 minutes (#3480)
* e0c2dfcb8 satellite/metainfo: don't allow uplink to commit same piece multiple times (#3460)
* 11f0ea325 5s (#3477)
* aa7d15a36 storagenode/contact: exponential backoff retries for pinging Satellites (#3372)
* f6a4155c4 certificates: move db test to separate file (#3439)
* 5abb91afc satellite: change the Peer name to Core (#3472)
* 281b8b696 update to the CLA file adding Rafael Gomes (#3474)
* e8f9a182d Bump go version to 1.13.4 (#3450)
* 4d26d0a6a storagenode/pieces: Add migration from v0 piece to v1 piece (#3401)
* 0c2e498f0 satellite/satellitedb: console tables archview comments updated (#3465)
* 0621830f0 cmd/storj-sim: allow user to designate which redis db they want to start at (#3458)
* 761cec5ea satellite/payments: archview comments updated (#3464)
* 5cb46d2ce satellite/payments: mock payment service created, api calls from frontend returned (#3448)
* 3a842bf53 change MaxClockOffset (renamed) to 15 min and use duration type (#3438)
* 84fea5820 Fix error when storagenode-updater service is stopped (#3440)
* 857c7f3cd storagenode/updater: disable self-autoupdate (#3459)
* 8d92c288e satellitedb: separate migration into subcommand (#3436)
* abd27d496 satellite: make outgoing connections with grpc (#3457)
* d820b5abf drpc: bump version
* 17e9044c0 pkg/rpc/rpcpeer: check both drpc and grpc for peers on a context
* 41c0093e5 drpc: enable by default (#3452)
* f9df0ea59 satellite/gracefulexit: check for unknown error in graceful exit disable test
* be2dd1ca7 cmd/storj-sim: add --redis flag (#3451)
* 76b64b79b cmd/identity: allow using redis for RevocationDB (#3259)
* aa761700a satellite/satellitedb: update nodes in sorted order (#3446)
* e0ef01d7a transparency in dialog image (#3443)
* a59d07d4d web/satellite: text selection on specific elements disabled (#3424)
* 43c2b7291 mobile: build gomobile in .build on Jenkins (#3421)
* 87687938d storagenode/contact: fix panic in ping satellites (#3447)
* 590312970 satellite/gracefulexit: add flag for enabling/disabling graceful exit on the satellite (#3437)
* 0d2c03a12 installer/windows: custom upgrade validation dialog (#3415)
* d9bb25b4b satellite/metainfo: support a wider range of values for RS.Total in satellite metainfo validation (#3431)
* 8ce09700a web/satellite: internal server error template added, errors separated (#3430)
* ecde507e4 satellite/payments: list invoices (#3389)
* 3c35339f0 Increase default difficulty in identity CLI to 36 (#3428)
* 8601f54f3 pr template: migrations run concurrently with api servers now
* 43103ae13 lower storage node counts in tests (#3427)
* acc7b116a scripts: use postgres script with all tests (#3404)
* 015350e23 storagenode-updater: add autoupdating (#3422)
* 7e7de8a29 web/satellite: date picker unit tests temporary removed (#3423)
* 59f81a4a0 groupcancel/ec delete: add a timeout based on completion times
* 44dc2c8c6 Ensure we dont return already claimed token again (#3420)
* bfa6699e2 satellite/repair: add timeout for repair download from a single node(#3418)
* e96d61501 satellite: remove satellite API code from peer (#3414)
* 4d85b1157 satellite/contact: improve errors in contact endpoints (#3356)
* 487813506 satellite/gracefulexit, storagenode/gracefulexit: add timeouts (#3407)
* 545388623 satellite/repair, uplink/ecclient: remove unused expiration arg from ec.Repair and ec.putPiece (#3416)
* 3ee0b89f8 storagenode/gracefulexit: delete pieces when receive Delete or Completed message from satellite (#3406)
* 00df37677 mobile: add build stage to Jenkins (#3377)
* 65a8e0bcb {satellite,storagenode}/gracefulexit: clearer log messages (#3413)
* 54594e79c satellite/gracefulexit: add metrics on satellite for graceful exit (#3355)
* 1828caf69 installer/windows: make upgrade possible (#3345)
* cd0940724 satellite/gracefulexit: use sync2 cycle inside satellite graceful exit endpoint (#3394)
* 506d9715d drpc: version bump (#3386)
* cd3d3850f satellite/gracefulexit: only allow one connection per node to graceful exit endpoint (#3357)
* 30a320574 satellite/payments: update account balance (#3379)
* 8786a37f8 uplink/storage: use Batch to optimize upload requests (#3408)
* 95cd3fc75 cmd/storj-sim: use less ports (#3412)
* b2ff13f1f {cmd/satellite, storj/satellite}: create command to run repair process in isolation (#3341)
* 14e661a1b uplink/storage/segments: non-functional improvements (#3400)
* 6354b3884 web/satellite: auth graphql api replaced with REST API (#3396)
* b0a943813 hot fix (#3410)
* 2e44a9fa6 satellite/satellitedb/dbx: remove sqlite.sql file (#3409)
* ec690929d satellite/metainfo: fix index out of range error for validate pointer (#3398)
* 1defd4dbf storagenode/piecestore: Respect config.MaxConcurrentRequests for drpc (#3402)
* 7c2daa4dd Use original pointer when calling UpdatePieces (#3397)
* 5b0398a71 storagenode/gracefulexit: Exclude finished exits from chore/worker processing. Fix update status bug (#3399)
* fef0c51c1 web/satellite: notification plugin (#3352)
* 016be4525 internal/textcontext: Give proper name compiled bin (#3395)
* 56f8b2d62 uplink/storage: remove bucket store (#3376)
* 4f281ef34 uplink: Refactor segments Store Get for metainfo Batch (#3362)
* 04b16c8b6 stylelint added, lint errors fixed (#3360)
* 9905f2c61 add piece num to transfer queue PK (#3390)
* 93353df4d internal/sync2: make Fence accept context (#3393)
* 366da85e0 web/satellite: ts default exports removed (#3382)
* 2105dfbbd web/storagenode: stylelint applied (#3363)
* c8d64bd59 Make windows version more dynamic (#3169)
* f17b34d02 cmd/storagenode-updater: add "...up to date" log line if shoudn't update (#3392)
* b10d44456 Rename --log-path back to --log for storagenode-updater (#3391)
* da2eaa708 scripts: dev script to start postgres before tests (#3344)
* 292e64ee2 satellite/gracefulexit: check duplicate node id before update pointer (#3380)
* a4e618fd1 handle context cancelled in satellite graceful exit endpoint (#3388)
* 279cd5625 change uuid conversion (#3384)
* 1469f7f41 storagenode/contact: wait for UpdateSelf before start (#3332)
* ed48e74e2 gracefulexit: fix build for drpc (#3387)
* e54d290d2 satellite/gracefulexit: Add signatures for success/failed exit finished messages. (#3368)
* 6df4d7bc7 storagenode/gracefulexit + satellite/gracefulexit: add storagenode-side transfer validation (#3371)
* 65abc95f7 Fix args passed to the storagenode-updater service (#3383)
* aca18109a Use new wallet, email and externl-address args in storagenode setup (#3381)
* 472672d52 docs/storage-node-downtime-tracking: add open issue regarding satellite outages (#3370)
* f0caa6ce5 satellite:gracefulexit: Update pointer after success (#3369)
* f4f142bf3 satellite/console: get, update auth  api endpoints added (#3375)
* d61b3688f internal/version: fix `OldMinimum` to use old semver type (#3373)
* 03f399857 web/satellite: removed ability to check project owner for deletion (#3350)
* 3880f7a60 web/storagenode: vue-svg-loader implemented (#3347)
* d9d82b033 uplink: Reduce satellite request using Batch when possible (#3351)
* 7f6893ea8 satellite/metainfo: fix broken object listing (#3348)
* 0b678c23c {internal/version,versioncontrol}: fix old versions (#3359)
* 696c567e8 satellite/gracefulexit: add piece hash validation for successful transfer (#3313)
* fa1ac24e1 satellite/gracefulexit: add failure threshold check (#3329)
* bff5c19de web/satellite: temporar removing of payments api calls (#3361)
* 75412e54e storagenode/piecestore: Rename liveGRPCRequests back to liveRequests (#3354)
* e82245e10 satellite/payments: credit card selection (#3304)
* 14c764853 storagenode/piecestore: Only limit grpc requests (#3342)
* 655e2b342 sattelite/console: sattellite name API endpoint added (#3349)
* 99ccdc1bc web/satellite: long email bug fixed (#3294)
* 810dc80d4 web/satellite: vue-svg-loader implemented (#3307)
* 1a304f5ef satellite/payments: add payments loop, update pending transactions (#3318)
* 521c39bda uplink/metainfo: cleanup method names (#3315)
* abb567f6a cmd/satellite: add graceful exit reports command to satellite CLI (#3300)
* 51d5d8656 pkg/rpc: drpc connection pooling
* 2c6fa3c5f pkg/rpc: remove read/write deadlines as a mechanism for request timeouts (#3335)
* 3e0d12354 storagenode/gracefulexit: Implement storage node graceful exit worker - part 1 (#3322)
* 3eec4e907 satellite/console: add REST delete API endpoint (#3337)
* 9d84e6347 installer/windows: custom welcome dialog (#3228)
* 4489c8653 web/storagenode: node version bug fixed (#3339)
* 97018d6d6 Expose range download in gomobile (#3336)
* 867771787 web/satellite: project owner status added on team page (#3160)
* 04c2454c7 satellite/metainfo: pass streamID/segmentID between Batch request/response (#3311)
* 5e78f4000 storagenode/pieces: remove old comment (#3334)
* 78000c019 Makefile: handle empty/default $GOPATH (#3333)
* 94cec7e3b add config dir arg to windows installer (#3328)
* 6dd478e43 satellite/console: forgot password, resend email endpoints added, default http route replaced with gorilla mux (#3327)
* 1014d5a7d cmd/satellite: add API run command to satellite dockerfile entrypoint (#3319)
* 244fea774 gomobile Reader.Read should take offset and length params (#3326)
* 1814fbfa8 satellite/console: new passwordChange API endpoint (#3308)
* f468816f1 {internal/version,versioncontrol,cmd/storagenode-updater}: add rollout to storagenode updater (#3276)
* 91872fcbd README: Replace reference Rocket chat to forum (#3324)
* 243ba1cb1 {versioncontrol,internal/version,cmd/*}: refactor version control (#3253)
* f65801309 Updating Support Link in README (#3314)
* 3c438f31b satellite/satellitedb: remove sqlite support (#3296)
* 89ed99770 satellite/satellitedb: switch to postgres only (#3320)
* c0ffd5988 Go 1.13.3 (#3321)
* 44bf98ee9 cmd/storj-sim: make postgres default to STORJ_SIM_POSTGRES (#3317)
* 2a5526fcc satellite/repair: reduce timeout (#3302)
* 02b930342 Adding cpustejovsky to the CLA (#3316)
* f929310ad pkg/rpc/rpcstatus: fix drpc grpc compatibilty (#3306)
* 071d1c431 upload: Add more info to returned error response & to logs (#3218)
* f4162bd33 satellite/console: add REST registration API endpoint (#3303)
* ee55d14f3 mobile: create release script (#3197)
* 45c35d7c3 satellite/satellitedb: add exit_status column to nodes table (#3301)
* 26cc625dc satellite/console: payments api  (#3297)
* 22d0f8994 satellite/gracefulexit: use zap.Stringer instead of zap.String (#3299)
* 24e72f35d satellite/payments: token deposit (#3283)
* afdd34058 name removed from header (#3209)
* 774758c65 web/satellite: bucket name bug fixed (#3240)
* 41dc5ef7f web/satellite: styling bugs fixed (#3277)
* 25dbf46bb web/satellite: validation added on delete account popup (#3290)
* edf8318c0 jenkins: update to golangci-lint 1.21.0 (#3289)
* 863e4cacb examples/scope: small scope exploder command line tool (#3266)
* 34764e5c9 cmd/satellite: create API subcommand (#3280)
* 855fca003 satellite/metrics: create a metrics chore (#3263)
* e5099f31f add context.Clean and correct rpc error code (#3295)
* 76ad83f12 satellite/accounting: add redis support to live accounting (#3213)
* 21c6737b1 uplink/ecclient: clarify defer logic in putPiece (#3247)
* 6e7607239 satellite/repair: improve logging (#3287)
* c1033fc1e satellite/console: move token to separate endpoint (#3292)
* bbe89ec9d add siedov mykola to clabot (#3293)
* 2301a8287 Satellite/PieceHashValidation: Increase time window from 2h to 24h to avoid timezone issues (#3291)
* eeb38245f satellite/audit: improve logging (#3285)
* 951c2891b {versioncontrol,internal/version}: add rollout to versioncontrol server (#3176)
* ed6b88a12 piecestore: update usage before completing upload (#3286)
* 87e376439 storagenode/cmd: add exit-status command for graceful exit (#3264)
* 37ab84355 satellite/gracefulexit: protobuf field name updates (#3284)
* a5f4bbee2 satellite/payments: dbx scheme renamed, userID placed on Account level (#3281)
* 4962c6843 piecestore: fix test flakiness around upload/download usage tracking (#3282)
* cf430d2d7 scripts: add check-monitoring script to detect changes to monkit calls (#3114)
* abb5b6c49 storagenode/piecestore: Fix to ignore both gRPC and dRPC EOF errors. (#3274)
* 1ad2ba7e3 storagenode/gracefulexit: Add graceful exit chore and worker. (#3262)
* 4d43b67ff add credit card to payment account (#3279)
* 9caa3181d uplink/piecestore: Check SN piece hash timestamp (#3246)
* 875c7dfe7 web/satellite: project members frontend selection caching added (#3238)
* f1867a954 web/satellite: project members sorting fixed (#3231)
* a34730a60 web/satellite: api keys selection caching added (#3239)
* 501816770 satellite/payments: receive credit cards (#3249)
* 57ff0af6b mobile: add EncryptionRestrictions (#3260)
* 5b20c716e satellite/repair: dont error on deleted segments (#3252)
* 93d5eeda3 Update dial.go (#3261)
* 694177e21 pkg/pb: regen gracefulexit.pb.go (#3270)
* 87a426f22 internal/testplanet: add satellite.API to testplanet (#3237)
* 1b6651766 contact: small typo
* b185dbbee satellite/discovery: remove discovery related code (#3175)
* 96aeedcde OrderLimit/GracePeriod: Increase time window from 1h to 24h (#3255)
* 6ede140df pkg/rpc: defeat MITM attacks in most cases (#3215)
* dd21953fd Change Error Handling to return more clear messages to the user/client (#3254)
* 78ccf1483 fix interface and EOF check (#3251)
* e567f2763 storagenode/piecestore: Change test to use ioutil.ReadAll to attempt to reduce test flake (#3250)
* a1275746b satellite/gracefulexit: Implement the 'process' endpoint on the satellite (#3223)
* d17be5823 remove random sleep in storagenode contact (#3243)
* a5d177653 audits: missing continue
* 84156d691 docs/blueprints: updating term to held amount (#3248)
* 0560b2bac web/storagenode: chart y and x axis reworked (#3179)
* 78a71ad3b web/storagenode: node status updated (#3220)
* 535742d37 satellite/payments: add coinpayments HTTP client (#3181)
* 451909b3e satellite/payments: account balance (#3242)
* 925639987 CI: test drpc and grpc (#3163)
* 743a0fc38 storagenode/cmd: create start graceful exit CLI (#3202)
* 91fcabd83 web/satellite: notification message text breaking fixed (#3221)
* 9f5d81b5b web/satellite: prevent multiple request on login and registration (#3234)
* 8170eb8d2 web/satellite: api key validation bug fixed (#3227)
* 4ba4d03fb web/satellite: labels and titles style following problems solved (#3226)
* 5ac531863 logo click action changed to redirection to Overview page (#3210)
* 89c68890b Bump drpc version (#3235)
* e9c36d560 satellite: make PointerDB an argument to satellite.New (#3233)
* 0cc23add5 satellite/payments - payment account setup (#3187)
* f60a7baf1 internal/testplanet: ensure that metainfo schema gets dropped (#3229)
* cf6f8fb53 Installer: identity directory defaults to %APPDATA%\Storj\Identity\Storagenode (#3216)
* 5f775b9e4 satellite/console: Added error for adding api key with existing name attempt. (#3185)
* ca0f74942 web/storagenode: charts on first day of the month updated (#3219)
* 9cbb0c437 cmd/storj-sim: expose GATEWAY_X_API_KEY (#3225)
* e82e61ad6 Fix typo in libuplink docs. (#3222)
* 784ca1582 satellite/audit: fix audit panic (#3217)
* 2e4d8630b installer/windows: checkbox for opening SNOboard after installation (#3212)
* ed3a42485 satellite/console: Removed 'user not found' message on password reset request (#3184)
* 3a3d576d9 satellite/audit: add mutex to pieceHashesVerified map (#3214)
* f75893c1b satellite/overlay: do not include gracefully exiting nodes in node selection (#3211)
* 10b364a2d cmd/storagenode: enable migration of configs of different types (#3189)
* 447c219d9 satellite/gracefulexit: Add protobuf definitions for communication between storage node and satellite (#3201)
* a744fdfef satellite/metainfo: remove Iterate from service (#3196)
* 0ea0d8c3d satellite/overlay: remove overlay.IsVetted (#3203)
* 7485a1d0f Fix jenkins builds of storagenode console. (#3208)
* 616b2f521 installer/windows: add to private Jenkins build (#3206)
* d5b2e1ef8 storagenode/signature: Reject uploads with a timestamp too far in the future (#3194)
* c00954323 satellite/audit: Add piece hash verified to log messages (#3204)
* 37491d0d3 storagenode: embed the console into the binary and makefile (#3164)
* 1d2808eeb Update links to documentation (#3198)
* 828b0af0d Storage Node Dashboard shortcut opens the SNOboard (#3171)
* 73d4d8346 satellite/accounting: implement tally as an observer (#2992)
* 4c4519f0b satellite/gracefulexit: add transfer queue for pieces (#3174)
* e1b7d0116 satellite/audit: do not fail or contain nodes for audited segments that are not piece-hash-verified (#3161)
* a75e3e6b8 satellite/repairer: fix segment_time_until_repair metric (#3199)
* 2a4a4b60d web/satellite: scroll added for project description container (#3166)
* ef2615fcf web/satellite: DatePicker rework (#3157)
* d29946d3d docs/blueprints: referral manager v1 (#3038)
* f81821785 units added to buckets table header (#3158)
* 79f3ef496 Add the tempDir as parameter (#3065)
* dfd76a93f web/satellite: Payments methods removed (#3180)
* a1b2527b2 satellite/web: Fixed Login Button overlapping. (#3156)
* 6770a9ce1 mobile: move to lib/uplink-gomobile (#3177)
* b908a09c8 satellite/repair: remove deprecated error message (#3193)
* eb5413ae5 defer close piecestore in downloadAndVerifyPiece (#3192)
* c1fbfea7f drpc: bump to latest version
* 7ceaabb18 Delete Bootstrap and Kademlia (#2974)
* 4fab22d69 pkg/rpc: don't leak goroutines during a drpc dial
* 64e43e555 pkg/rpc: return context error if ready after DialContext fails
* 6afa4dd9c satellite/accounting: refactor code and remove unused fields (#3178)
* 4a86906e1 web/storagenode: memory conversions extended (#3188)
* a11619e7f storagenode/console: use bandwidth monthly summary (#3183)
* 4824ecdb8 storagenode/console: use bytes for remaining info (#3186)
* 0911b7d1f uplink/gateway: wizard message update (#3150)
* b25e0154c internal/testplanet: use postgres for pointerDB (#3139)
* 51003dcad web/satellie: added missing alt attribute to img tags (#3172)
* c169993b3 web/satellite: sorting in api keys fixed (#3159)
* 284b75d86 web/satellite: fix registration popup ref check (#3173)
* 394a9c82c satellite/{accounting/tally,repair/checker}: create a valid test pointer (#3167)
* 2f5ede65f mobile: API updates (#3155)
* 794254b0b mobile: build libuplink aar (#3165)
* f05a2eea5 satellite/console: remove domain prefix from token cookie key (#3170)
* c9e0aa7c7 pkg/kademlia: make tests run/work with drpc
* b2e328f11 storagenode/dashboard: update online status (#3168)
* 6e6d0ad9b split satellite: create separate API process (#3152)
* 94c7df0d6 pkg/rpc/rpcstatus: Fix return type (#3162)
* 1db425123 Satellite/repair: Add Repair Threshold Override to allow earlier repair (#3151)
* ef5e0dce2 scripts: ignore .build directory for size checks (#3153)
* b7d4a3b09 web/storagenode: vue style guide applying (#3120)
* 2dfea75e8 web/storagenode: logic removed from template (#3138)
* 6a1203ea4 web/satellite: logic removed from template (#3136)
* 4f2f8ae11 satellite/overlay: add UpdateExitStatus and GetExitingNodes for graceful exit (#3087)
* 08ed50bca satellite/metainfo: add commit interval to prevent long delays between order limit creation and segment commit (#3149)
* 89c59d06f storagenode/storagenodedb: add SQL receiver logic for graceful exit (#3067)
* c7d05ebbf web/satellite: editing of project description bug fixed (#3145)
* 755cbd4dc storagenode/main: map aliases for kademlia config values (#3118)
* 09c3efa51 web/storagenode: vue files renamed (#3121)
* bd8807e43 Fix locale to "en-US" when parsing numbers (#3143)
* f94bc8e46 update to go 1.13 (#3140)
* 9f082c61b Add AlexeyALeonov (#3141)
* edadf4600 satellite/audit: delete nodes from containment when segment has changed (#3115)
* 8989c0ae3 Add support for freebsd on amd64 (#3116)
* d1109f17f storagenode/dashboard: change etherscan link to open ERC-20 transactions (#3131)
* b85addb65 web/satellite: Payment methods markup  (#3054)
* 423d35fb3 satellite/console: Added support URLs and other fields to config file (#3090)
* 29b96a666 internal/testplanet: fix conn leak (#3132)
* 02f68d68d Put -s and -w in the right spot (#3135)
* 93349f247 pkg/rpc: add WithInsecure when doing non-tls dials
* fd54cc80d web/satellite routing updated, tests added (#3113)
* 2c5e16988 storagenode/storagenodedb: Vacuum info.db to prepare for splitting storagenodedbs (#3134)
* acbe44943 satellite/console: remove payments (#3074)
* c8aa821cc pkg/certificates: move certificate package to root (#3107)
* 52497c253 Bump go version to 1.13.1 (#3130)
* 0afbb1e0c web/satellite: vue style guide applying for UI (#3112)
* c874dae59 internal/testplanet: ensure monitor chore is finished before contacting satellite (#3124)
* 9962f0c3f build: make binaries smaller with -s and -w ldflags (#3125)
* fd72de211 satellite/satellitedb: update node version columns in UpdateCheckIn (#3129)
* 098cbc9c6 all: use pkg/rpc instead of pkg/transport
* c71f3a3f4 internal/version: Change default endpoint to query (#3126)
* 9edfb6efe satellite/satellitedb: Initial GE Satellite DB Implementation (#3049)
* c1fb791a4 satellite/accounting/rollup: pause in tests (#3122)
* e0d5cbbbd jenkins: update golangci to v1.19.1 (#3119)
* e82666c76 satellite/accounting/tally: ensure we have a root piece id (#3123)
* 94bbb9563 internal/testplanet: set intervals to 15s by default (#3103)
* cd187db0a jenkins: disable flaky test-certificates (#3111)
* 366a0be6a storagenode: avoid starting command on service stop (#3105)
* 8d16349b4 storagenode/winsvc: avoid the 15s sleep when stopping the service (#3098)
* 515799267 fix certificates auth export command (#3110)
* ab3e3f827 satellite/repair/checker: set erasure share size in tests (#3101)
* a7040647a run certificate authorization endpoint (#3108)
* 607da4ab4 metainfo: move FinishDeleteSegment logic to BeginDeleteSegment (#3104)
* eb73a3d35 Windows installer for storage nodes (#2921)
* 580e511b4 storagenode/storagenodedb: Migrate to separate dbs (#3081)
* 4bd1ce868 satellite/metainfo: close loop separately to avoid logical races (#3100)
* 3647fd7b3 Adding our new employees (#3106)
* cfe058d19 web: using kebab notation for props names in html (#3030)
* d32d85a71 pkg/listenmux: resolve deadlock in test
* a20a7db79 pkg/rpc: build tag based selection of rpc details
* 9ceff9f9c satellite/overlay: move CheckIn benchmark to overlay (#3095)
* d2502bb51 Adds tests for kad replacement and restores kad operator configs (#3094)
* 69aa0c6cc satellite/console: GraphQL input length limitation. (#3045)
* 1ed724b7a internal/migrate: make rebind optional (#3071)
* ddfcbfcd3 jenkins: increase timeouts, so master has time to complete (#3096)
* 68d281db4 autoupdater: use blang/semver (#3063)
* 756d0ad2b web/satellite: new navigation area adaptation implemented (#3093)
* 74ccc1b86 web/satellite: eye icon fits lastpass (#3091)
* 02fb891d2 internal/testplanet: reenable TestUplinksParallel (#2338)
* 8b358ef36 internal/dbutil/sqliteutil: Fix error handling, ensure connections are closed (#3078)
* 53db51715 satellite/overlay: don't use transport observers (#2989)
* 724bb4472 Remove Kademlia dependencies from Satellite and Storagenode (#2966)
* 93788e521 remove kademlia: create upsert query to update uptime  (#2999)
* 45df0c534 storagenode/process: respond to Windows Service events (#3025)
* 946ec201e metainfo: move api keys to part of the request (#3069)
* d22987ea1  satellite/audit: Fix flakiness in TestReverifyDifferentShare
* 870e70e0e web/satellite: copy api key popup title changed (#3080)
* 149a2f8d3 go.mod: update to grpc 1.23.1 and update golang.org/x (#3072)
* 26526de54 add auto-updater to Makefile (#3076)
* 021dba6c9 web/storagenode: UI node ID container fixed (#3079)
* d86035360 storagenode-updater: respond to Windows Service events (#3077)
* 43846f207 cmd: rename auto-updater to storagenode-updater (#3089)
* b07a490f9 satellite/metainfo: fix GetObject method (#3088)
* fae2c2c9f satellite/contact: return status codes from endpoint (#3086)
* a4048fd52 satellite/audit: fix containment mode (#3085)
* 1c72e80e4 uplink/satellite: fix for case when inline segment is last one (#3062)
* ccdd43561 defer client.close() (#3084)
* 695de9dcd rm noisy debug logs that we dont need (#3083)
* ce3203e91 update NodeSelectionConfig.OnlineWindow to 4hr default (#3082)
* a2b1e9fa9 storagenode/storagenodedb: refactor both data access objects and migrations to support multiple DB connections (#3057)
* 186e67e05 pkg/transport: set default timeout to 10 minutes (#3075)
* 574c96c35 satellite/metainfo: Verify storagenode signature on satellite upload (#2985)
* fd20fa38c internal/dbutil/sqliteutil: add MigrateTablesToDatabase (#3064)
* 7c203b488 add satelliteSystem to testplanet and update tests (#3066)
* cc70cd232 satellite/repair: add metric trackers for segment age before repair (#3056)
* 684b07b2c scripts/protobuf.go: update drpc version for protobuf generation (#3059)
* d65386f69 auto-updater: unpack/check binary, restart service (#2968)
* 4aaf525bd satellite/audit: set devDefaults for ChoreInterval and QueueInterval to 1m (#3058)
* febe32bc7 pkg/miniogw: Add a missed stack trace error (#3035)
* 91d54af70 Add satellites database business objects. (#3055)
* b37ea864b satellite/repair: delete pieces that failed piece hashes verification from pointer (#3051)
* 7ada5d415 satellite/metainfo: make piece hashes nil before storing pointer in metainfo.UpdatePieces() (#3050)
* 0c2ae7786 storagenode/dashboard: Enable storagenode dashboard in docker images (#3024)
* d3ef574b2 pkg/pb: minor changes to contact.proto (#3048)
* 5a50042c7 uplink/storage/streams: Add test for interrupted deletes (#3040)
* f550ab5d1 Uplink "import" command (#2981)
* 886041e0b satellite/satellitedb: add new graceful exit tables and add graceful exit fields to nodes (#3033)
* 95aa33c96 satellite/repair/repairer: update audit status as failed after failing piece hash verification (#2997)
* 46ada7b6b web/storagenode: error processing added (#3046)
* 9d1e19cd9 satellite/console/server: fix usage rollup report (#3047)
* 8b54e329f web/satellite: added trailing coma rule to tslint (#3044)
* 7240e6cbb satellite: remove remote/inline file from BucketTally (#3041)
* ca058e606 storagenode/orders: fix data race in settle (#3042)
* 44bcdd222 storagenode/contact: test node info is updated (#3039)
* ccbf73ecc uplink/ecclient: Remove unneeded atomic operation (#3036)
* 77e7f182d notification height fixed (#3020)
* 8a48500ba uplink/ecclient: Report success in debug level (#3037)
* 4266f8c36 web/satellite: linter updated, analytics artifacts removed (#3043)
* a342c3aba web/satellite: analytics removed, buckets header foreground improved (#3027)
* 0dcbd3dc0 bootstrap/satellite/certificate/storagenode: register drpc services
* 007662d49 pkg/server: serve drpc listeners along with grpc
* 477b47f55 pkg/server: use a listenmux with nothing registered
* 731016cd8 Increase file size limit to 650 KB (#3034)
* 8ef57a2af satellite/satellitedb: use noreturn (#3022)
* a085b05ec satellitePeer -> satellite rename consistency in repair test (#3032)
* ab1147afb storagenode/pieces: fix race condition in cache service (#2972)
* 82a651ac3 satellite/contact: Populate PeerIdentities table in satellitedb (#2998)
* cc8a47324 scripts: Fix warn message update sat config lock (#3029)
* 500ed27e0 web/satellite: error while project creation fixed (#3023)
* 489047a96 web/satellite: notifications for billing page implemented (#3004)
* 208327835 Script for deploying the Docker manifest for watchtower (#3015)
* bb6086aea web/satellite: API keys paged backend (#2839)
* 0c1798407 web/satellite: API Keys Pagination implementation (#3019)
* 5383b8be8 web/storagenode: frontend bug fixes (#3021)
* 1d8cd526e storj-sim: correct storagenode dashboard config (#3010)
* 6c0d21046 web/satellite: minor improvements (#3016)
* 8b668ab1f satellite/metainfo.Loop: use a parsed path for observers (#3003)
* e5ac95b6e storagenode/inspector: fix TestInspectorStats flakyiness by waiting for requests to be handled (#3018)
* 9f2f1527c storagenode/storagenodedb: add new tables for graceful exit (#3008)
* aa3567187 satellite/audit: worker now verifies and reverifies (#2965)
* dbe90926c  internal/testplanet: reduce coalesce duration (#3009)
* 3d410add4 satellite/overlay: avoid large statement for piece counts (#3001)
* c45d53e08 storagenode/dashboard: change disqualification text (#3007)
* c139ed8ea storagenode/console: remove kademlia (#2942)
* 7718802f0 storagenode/storagenodedb: prepare for multiple databases (#3005)
* 708c95d04 pkg/listenmux: multiplex listener based on first bytes
* 40ff56f6c versioncontrol: add new version schema (#2991)
* b9e533119 internal/migrate: remove duplicate logging (#3000)
* 289cfe8ff satellite/repair: do not log "retrieved segment" if repair queue empty (#2995)
* eac20e977 docs/design: Graceful Exit (#2866)
* 4ddfffce4 web/satellite: markup for billing history page (#2993)
* 6c80f01bf pkg/certificates: add authorization endpoint and refactor (#2971)
* 0b32572ae migrate: Allow work on separate dbs (#2996)
* 7cf565056 pkg/certificate: properly close certificateclient.Client (#2986)
* 2fc4d6161 implement contact.checkin method (#2952)
* 64c467ffe uplink: integrate new Metainfo calls (#2640)
* 2c7813d40 satellite/console: Added email normalization to users table (#2586)
* 09447e24c jenkins: update to go1.13 (#2990)
* a801fab66 all: add archview annotations (#2964)
* 86f4b41a7 docs: add reputation documentation (#2982)
* 0ccae6b06 cmd: windows log file workaround (#2979)
* 0d4fd6cee cmd/uplink: fix progressbar data race (#2987)
* e03e6844f pkg/process: reduce noise in storj-sim (#2988)
* 40ca660c0 all: use min tls 1.2 for grpc (#2967)
* 7589ca796 cmd/storj-sim: allow overriding executables (#2976)
* 1e22f5804 satellite/server diagnostics and http headers improved (#2978)
* 646f290ff satellite/accounting: use sync2.Cycle instead of ticker (#2977)
* 3b72cb672 web/satellite: side bar markup updated (#2937)
* c5658fa73 web/storagenode: markup and logic for node operator dashboard (#2906)
* c81e4fcb9 design/blueprints (#2927)
* d6f50ec70 notification to buckets page added (#2949)
* 2b9fcd119 web/satellite: tslint update (#2962)
* 64602c300 fix flaky repair tests (#2973)
* 60eba990e use-drpc: use protoc-gen-drpc to generate protobufs
* fb1081522 Repair with hashes (#2925)
* 338775028 storagenode/contact: create chore for nodes to ping satellites (#2877)
* a3e0955e1 satellite/satellitedb: ensure that we process orders in order (#2950)
* c35ad5cbf storagenode/console: update api (#2969)
* 28353823e jenkins: remove Jenkinsfile.gerrit (#2953)
* 587be8f20 auto-updater: download versions and archive with binary (#2922)
* 6d363fb75 satellite/audit: create the audit queue, chore, and worker (#2888)
* 1fc0c63a1 {cmd,pkg}/certificates: service refactor (#2938)
* b222936d1 cmd/certificates: separate subcommand configs (#2932)
* c4ee7eb51 jenkins: build and test web/satellite (#2908)
* 15e52e098 web/satellite: Fixed usage tests. (#2963)
* 4bcaa5912 add list direction option to bucket_test.c (#2961)
* 61168493d uplink: don't stop deleting segments on first error (#2943)
* d6bbc2a65 modules versions fixed (#2926)
* bc8a531a2 dashboard refactored, tests added (#2929)
* e47f08afa lib/uplinkc: use ListDirection in struct BucketListOptions (#2946)
* 9821a21e5 satellite,storagenode,bootstrap: add contact service to peer (#2951)
* adfa16188 pkg/contact: bare-bones service and endpoint (#2941)
* 62df8ddb0 cmd/certificates: auth export improvements & certificates test script (#2897)
* af5fb8e9c satellite/vouchers: deprecate voucher endpoint, return 'please upgrade' error (#2940)
* f7bae57e5 pk/pb: add initial proto for the new contact endpoints (#2948)
* 05e9916b9 storagenode/buildscript: Overwrite existing docker manifest file (#2934)
* dadd7327d satellite/nodestats: return storage usage in Byte*hours (#2858)
* a4dddc65b web/satellite: Fixed Project Members deletion. (#2891)
* 534a74dac web/satellite: account billing page initial markup (#2928)
* f7403f97b storagenode/storageusage: add summary, rename timestamp to interval_start (#2911)
* 758f7cb3d storagenode/bandwidth: remove bandwidth concerns from console, add satellite summary (#2923)
* 83815ee7b satellitedb: always release savepoint processing orders (#2936)
* 52cdca01a Add TopperDEL to the CLA list (#2947)
* bf5a16b82 jenkins: work around git.apache.org failures (#2944)
* ee614bf03 storagenode: add custom dial timeout for orders sending (#2939)
* 4d65fdaea docs/design: Adapt SN Downtime tracking to be a blueprint (#2931)
* d7d6e23a3 pkg/server: Don't use Sugared logger (#2935)
* 3fbe31aad storagenode: Increase order sending request timeout (#2930)
* 10a896bf7 web/marketing: static asset path (#2872)
* a6721ba92 satellite/metainfo: Improve metainfo ListSegments (#2882)
* 9a1b9f843 uplink/ecclient: change delete logs from err to debug level (#2917)
* 7e702aa10 web/satellite: navigation, button and project members unit tests (#2907)
* c680970e0 Storage node auto-updater skeleton (#2902)
* 48ba45eb5 pkg/process prometheus: forgot digits (#2920)
* 539f3857f web/satellite: downgrade typescript and pin stuff (#2918)
* b20dcfd64 add type to prometheus metrics (#2916)
* 9d1910cb2 storagenode/orderarchive: Reduce TTL from 45 to 7 days (#2915)
* b3f9a8813 pkg/process: remove prometheus help (#2914)
* 537769d7f storagenode/orders: Don't return error Archiving unsent (#2903)
* 243d28059 Clean up jenkins message to slack (#2912)
* 8a5db77e0 storagenode/retain: add comment (#2910)
* 7e37452ab open port for storage node dashboard API (#2899)
* aa37334e7 Send failures to slack for added visibility (#2909)
* 4ede12a2a satellite/orders: Fix for V3-2529: Release v0.19.0 storage nodes can't submit orders, duplicate key value violates unique constraint (#2900)
* 24a36999b Revert "web/satellite: navigation, button and project members unit tests (#2904)" (#2905)
* 5bb51c987 web/satellite: navigation, button and project members unit tests (#2904)
* 368f6cc32 web/satellite: account route redirect fix (#2895)
* 8eda360ad add segment path into logs (#2898)
* 07d6019a1 web/satellite: project members UI slightly reworked, bugs and tests fixed (#2896)
* 5fb823843 Fix downloading non encrypted segments (#2870)
* b4d7d6778 storagenode/reputation: add disqualified flag (#2862)
* 62e3bf5b3 storagenode/retain: fix concurrency issues (#2828)
* 842c7118c docs/design: update audit v2 (#2901)
* c7bd75b03 docs/design: Storage Node downtime tracking (#2857)
* 3e121b13e docs/design: satellite service separation (#2815)
* 7b874db8c web/satellite project related bugs fixed (#2894)
* 66ec727a3 jenkins: backwards-compatibility, don't overwrite installed binaries (#2892)
* 499c4d0c2 web/satellite: navigation bugs fixed (#2893)
* 46a495fba docs/design: Fix typos & remove not applicable title (#2879)
* ec715cba9 docs/design: Remove subtitle part not applicable (#2880)
* 1b2a2d045 [build]: Adjust arm63v6 and aarch64 images to match convention (#2845)
* 8c2439943 web/satellite: usage api refactored (#2864)
* 8fbb25f3b web/satellite: ProjectMembers unit-tests refactoring. (#2865)
* f404aad87 Add deploy script for storagenodes (#2889)
* c8405f6c2 docs/design/archive: moves archive directory (#2885)
* b587c93f4 satellite/gc: Service run must call mon.Task (#2887)
* 49303ea3a satellite/audit: mv ReservoirService into its own file (#2886)
* 599324c36 satellite/dbcleanup: delete expired serials from satellite (#2867)
* c309bd3fe lint: add linting for errs package (#2881)
* 9abc3e9d6 moves audit gating design draft to doc archive (#2883)
* 106a21ebe docs/design: Use WiX toolset directly instead of go-msi for Windows installer (#2851)
* a33106df1 satellite/satellitedb: persist piece counts to/from db (#2803)
* d0ab3c03e cmd/*: Change loglevel from error to warn (#2876)
* a250551b6 storagenode/piecestore + uplink/piecestore: return `PieceHash` and original `OrderLimit` during GET_REPAIR (#2775)
* 977472ed3 all: use fewer storage nodes to improve test memory usage (#2875)
* 33aff7195 satellitedb/overlay: add database for storing peer identities (#2764)
* 1f3537d4a storagenode/vouchers: remove storagenode vouchers (#2873)
* 462640a9f docs/design: automatic updater (#2789)
* d5667fbe3 lib/uplinkc: ensure it compiles on 32bit arch (#2835)
* 36c9d569f design/docs: add successful pingback to kademlia removal document (#2837)
* 051052307 satellite/rewards: add mongodb into partner info (#2800)
* 6ff94caf2 satellite/satellitedb: move tests near the interface (#2863)
* 4e16a5c59 satellite/marketingweb: fix broken pipe error (#2853)
* 65e2d2e71 storagenode/piecestore: ignore canceled errors on download (#2822)
* 2ae4129d0 satellite/nodestats: add disqualified flag #2856
* 12d50ebb9 streams: don't encrypt segment count (#2859)
* 357a4cdf5 projects tests fixed (#2854)
* 9981bbda2 web/satellite: Projects api refactoring (#2844)
* 3d9441999 storagenode/orders: add archive cleanup to orders service (#2821)
* df2969964 satellite/audit: Improve code comment in reporter (#2838)
* 00b2e1a7d all: enable staticcheck (#2849)
* 7c6fdd09d web/satellite: router refactoring (#2846)
* 1b9d163c9 satellite/marketingweb: WriteHeader should only be called once (#2850)
* cff6dc2b2 Rename design-doc-process.md to README.md (#2848)
* 14e36e4d6 storagenode/nodestats: fix issue on 32 bit platforms (#2841)
* 2d69d4765 all: fix Error.New formatting (#2840)
* 243cedb62 satellite/audit: implement reservoir struct and RemoteSegment observer method (#2744)
* 476fbf919 storagenode/storagenodedb: refactor SQLite3 database connection initialization. (#2732)
* 87ef5e339 web/satellite: buckets api and store refactored (#2830)
* 1e099839d web/satellite: api keys component updated (#2819)
* 0f5d53378 web/satellite: Notification rework (#2781)
* 9df2ec6a3 docs/design: Kademlia Removal Design Proposal (#2686)
* 513563eff TabNavigation component fixed, unit tests added (#2834)
* 9ec0ceddf pkg/revocation: ensure we close revocation databases (#2825)
* 95080643b satellite/repair: fix data race (#2833)
* 87f3b6c70 satellite/audit: Improve comments in verifier (#2829)
* 25f0b1398 pkg/peertls: extension handling refactor (#2831)
* 77f1555cf web/satellite: Moved endpoint url to config file. (#2832)
* 56383de29 satellite/metainfo: use status.Error and fix error codes (#2827)
* 2e3ff4587 web/satellite: ProjectMembers store and graphQL queries refactoring (#2820)
* 661535018 initialize used space table with sum over pieceinfo (#2818)
* 3c4179442 pkg/kademlia: clearer error message (#2824)
* 95c88b803 satellite/nodestats: use proper error codes (#2826)
* 2e1fc4bc5 docs/design: describe process in more detail (#2817)
* 25154720b lib/uplink: remove redis and bolt dependencies (#2812)
* 8832a393e formatting fix (#2823)
* 863675605 scripts/test-sim-backwards: usability improvements (#2816)
* d83a96513 storagenode/piecestore: Add retain service on storagenode (#2785)
* bf681f32a web/satellite/credits: api/store/component and tests refactored (#2811)
* 1a69ec831 satellite/orders: document protocol and fix typos (#2813)
* b77f582b2 web/satellite: Project members web client refactoring. (#2783)
* 6400d63a6 satellite/satellitedb: Add piece count column to nodes table (#2795)
* af41039ea web/satellite: projects api key refactored (#2807)
* ac5858961 Add aarch64 support and use go 1.12.9 (#2814)
* 309ad2cb2 jenkins: Backwards Compatibility use Postgres (#2810)
* 057d30152 uplink/storage/segments: seed download permuatation with timestamp (#2809)
* 546d099cf storagenode/orders: An invalid one don't have to stop all (#2804)
* e47b8ed13 storagenode: No FATAL error when unsent orders aren't found (#2801)
* b1abbe5ce web/satellite: add team member placeholder changed (#2794)
* 5fe05df77 web/satellite: Buckets page rework (#2763)
* 8df683a26 Update satellite settlement endpoint to batch order processing into transactions. (#2762)
* 989051328 web/satellite: api key header updated, redundant components removed (#2770)
* 497f10d7b add method CleanArchive to delete archived orders (#2796)
* 189b26889 uplink/piecestore: Change where ignore cancel happens for closing downloads (#2786)
* 9c4af157f docs/design: windows installer (#2780)
* 83461a7c0 satellite/metainfo: fix storing enc key/nonce (#2759)
* e22c0bff9 lib/uplinkc: fix flaky download test (#2791)
* 2dab0ab46 pkg/process: only propagate missing keys to flags (#2784)
* 7db685172 satellite/rewards: update current reward to be finished once redemption cap has reached (#2745)
* 1915b59af satellite/repair: monkit improvements (#2773)
* 5c61c79ab Start building binary for the link sharing service (#2790)
* b3a3d8df0 go.mod: update grpc-go (#2778)
* 012775f87 web/satellite user api simplification (#2787)
* 2c769fe9d satellite/overlaycache: add missing audit and uptime success count (#2788)
* d28c8aad8 design doc for design doc process (#2782)
* e1d8503f5 jenkins: add backwards compatibility test (#2774)
* 786828ea0 satellite/console: bypass activation token logic for open source partner referral link (#2736)
* fec10ccbd cmd/certificates: ensure we can bind config values (#2779)
* 3a82b6397 uplink/ecclient: performance - close connections faster (#2757)
* a496e26e4 satellite/rewards: hardcoded partner ids (#2776)
* 141af7e2f storagenode/console: refactor service and api (#2751)
* 7fcbec48f go.mod: update grpc-go and go (#2777)
* 48211daa9 uplink/piecestore: handle Download errors better (#2771)
* 6ad7ca769 satellite/satellitedb: save tallies in single transaction (#2758)
* 26fb99247 storagenode: Add more test assertions (#2772)
* 9eba5ac63 lib/uplink: remove Seek method (#2768)
* c57aaee66 docs/design: update kademlia audit gating
* 43cadc65e skip flaky test (#2769)
* 89a8d3273 storagenode/pieces: Restore lost test case (#2767)
* ce5c45b33 satellite/console/server: parse html files as templates (#2750)
* f3db40910 lib/uplinkc: make .so smaller by removing symbol information (#2730)
* 022f5d2e1 storagenode: add space used cache for pieces (#2753)
* 878a3c802 satellite/console: store partner id on api key and project creation (#2743)
* 3e41767f2 satellie/gc: enable garbage collection on the satellite (#2765)
* 32f95a14f satellite/certdb: remove certdb that was used to store uplink certificates (#2760)
* 13900140e remove identity setup from uplink instructions (#2761)
* 0decd1419 satellite/console: project members refactoring (#2752)
* 260f64827 Update repair-with-hashes.md (#2756)
* 7af05177e pkg/identity: support encode and decode functionality of Peer Identity information
* 1f837c53e uplink/ecclient: read concurrently with dials during download (#2711)
* 6290d8c64 web/satellite: title removed from header component (#2749)
* fc4c675ff web/satellite: pagination component implemented (#2722)
* 02b7be74f web/satellite Generic List component (#2712)
* cccc43862 web/satellite: team area header updated (#2740)
* 4cf2b9673 storagenode/nodestats: fix test duration (#2748)
* 9dccf59e8 Restrict node info only for trusted satellites (#2737)
* 05e99a78d docs/design: use piece hashes for repair (#2723)
* af69dcc6d docs/design: auditing 2.0 (#2703)
* 9c4d2eb5b go.mod: use pseudo version (#2741)
* ad5a1d4a0 go.mod: use custom grpc that doesn't have tracing support (#2739)
* 28a7778e9 storagenode/nodestats: cache node stats (#2543)
* 34c928514 Uplink setup: respect tls args in libuplink config (#2738)
* 4576c4f41 Satellite console server gzip compression (#2665)
* 17bdb5e9e move piece info into files (#2629)
* 271f9ea72 architecture-owner codeowners for design docs (#2735)
* f236fc3d9 Remove the use of in memory SQLite3 tables for storage nodes. (#2726)
* 8e60ba269 libuplink upload/download err handling improvements (#2725)
* c899b2542 codeowners: use proxy team for assignment (#2729)
* f66e9052c codeowners: setup random reviewers (#2727)
* b0e596471 satellite/console: create owner_id column for project table (#2706)
* 2a464bfdb Jg/1622 add backwards compatibility test (#2656)
* e34b2c553 Reduce UpdateAddress calls with address cache (#2681)
* bdcb40fbc storagenode/storagenodedb: Add cursor to pieceInfo.GetPieceIDs (#2724)
* c8edeb025 satellite/overlay: rename overlay.Cache to overlay.Service (#2717)
* de7dddbe5 metainfo: Batch request (#2694)
* c32fc283c make display order consistent (#2705)
* e5f30bddc Add  Karl Mozurkewich (#2720)
* e5379d4f6 libuplink: Add size to `ObjectInfo` (#2714)
* 538e8168a pkg/process: support json trace output (#2713)
* 21a3bf89e cmd/uplink: use scopes to open (#2501)
* 53017a0f0 Fix graphql query escaping (#2695)
* fcbc9d71d satellite/repair: add shouldDelete (#2702)
* e4c10f331 uplink/ecclient: add more monkit for segment piece info (#2701)
* 25823d433 add new line to end of require macro print statements (#2710)
* a89b3a745 Web/satellite search component test build fix (#2709)
* 688d932d9  Make one implementation for SetAttribution/SetBucketAttribution (#2683)
* 4a827cd92 web/satellite color pallet of buttons on create/login account updated (#2697)
* 28156d357 storagenode: more live request tracking (#2699)
* 65932ad69 Updating the certdb to support storage of multiple public keys for same node ID (#2692)
* 1b051ef3c web/satellite common search component created (#2696)
* 8532625f8 Add 24 hours to account for the entire rest of the end day (#2698)
* 471b46c0c adding more info the error logs in kademlia (#2668)
* a13008198 docs/design: Remove "Title" word from title header (#2693)
* ebbf0e146 uplink/storage: don't import mock in production code (#2687)
* 4547084f2 Implement checker observer (#2620)
* 51833d065 satellite/satellitedb: get active offer for partners (#2664)
* b74d4198f Use UTC date in TestCachedBandwidthMonthRollover (#2684)
* 194451064 internal/sync2: Fix typo in doc comment (#2685)
* dbfde60f7 Enable architectural constraints (#2650)
* 369a51ed0 lib/uplink: ensure it's silent by default (#2676)
* 287fdf993 Integrate new Metainfo calls (server side) (#2682)
* 238e264a8 satellite confirms success threshold reached (#2657)
* c9b46f2fe V3-1987: Optimize audits stats persistence (#2632)
* 26a2fbb71 storagenode: batch archiving unsent_orders (#2507)
* 4adafd056 satellite/rewards: add generate referral link logic (#2655)
* 4f0d39cc6 don't use global loggers (#2675)
* 3cd477454 storagenode/piecestore: Make method unexported (#2674)
* cc7f5d2f8 check nil against Bandwidth service not DB (#2673)
* ec3d5c0bd don't use global loggers (#2671)
* 9ba8b53ed pkg/auth: use grpc.WithPerRPCCredentials (#2670)
* 9496a8adb Changed forgot password button label. (#2660)
* 338c1b226 adds comments (#2666)
* 4b8820230 Initialize correctly libuplink for `uplink setup` and `gateway setup` (#2662)
* abef20930 storagenode: Report gRPC error when satellite is untrusted (#2658)
* 4cc697239 lib/uplinkc: avoid allocating too much on stack and nicer names (#2659)
* 6dee5a391 added more logging info (#2667)
* 55e3a37d1 Adds more detailed error logging information V3-2239 (#2654)
* 4632ab0a6 Delete irreparable segments (#2642)
* 159ad439b Add count to repair queue (#2661)
* d85934131 web/satellite Master fix (#2663)
* efa2c776b satellite/satellitedb: add updateEarnedCredits method for user_credits table (#2609)
* 2f15c0c5e Update libuplink-android tests (#2631)
* 6b40866ed Re-implement libstorj API (V2) using libuplink (V3) (#2573)
* 518a19822 CSP implementation for satellite console (#2644)
* 2b2a01fed web/satellite cursor:pointer removed from new project popup (#2604)
* 8f8b13abb Re-enable SN bandwidth rollups. Fix SN bandwidth rollup unique constraint issue. Re-organize service code (#2617)
* 5ab32a494 web/satellite project payment methods code refactored, dialogs repositioned, bugs fixed (#2643)
* 64199d37b Improve Node Selection Query (#2648)
* e75813d09 satellite/repair: move segment repairer to satellite and simplify (#2651)
* 175c30048 Use SaltedKeyFromPassphrase in Uplink CLI and Gateway (#2637)
* dd7c8610b satellite/repair: move test files (#2649)
* 5d0816430 rename all the things (#2531)
* 906c77b55 Add RetainStatus to storagenode config (#2633)
* 3b6c69c59 fix mutation (#2641)
* 2144181c9 satellite/marketingweb: Add Partner Offer Type (#2540)
* e260f9f1c Remove unnecessary UNION and Selects (#2624)
* 734c793de Add UpdatePieces method to Metainfo service (#2572)
* 5b6f3020f swapping noop to standard zap logger (#2634)
* 0cdeae192 add missing error handling (#2630)
* b9a17913f storagenode/pieces: remove buffering from reading/writing and fix io.EOF bug (#2554)
* d84b98771 Disable CAWhitelist config when signing identities(#2628)
* 6e5e26124 Fix missing Version config parameter (#2627)
* f11413bc8 Implement garbage collection on satellite (#2577)
* b8fe34981 Support empty path components (#2574)
* 5710dc3a3 Metainfo RPC segment methods (part 2) (#2616)
* 4de815ae5 web/satellite short name updated to nickname (#2610)
* 3ebeda333 storage/redis: make sure to return after put() (#2625)
* 0e1cb7bfb CompareAndSwap in KeyValueStore (#2602)
* b24e60a33 satellite/satellitedb: Use var block for single variable declar… (#2622)
* aaf3283b4 satellite/console: referral link for individual users (#2614)
* 353b08992 update testplanet with libuplink (#2618)
* 3c8f1370d [v3 2137] - Add more info to find out repair failures (#2623)
* 79e9b62f6 pkg/storage/segments: Clarify logic in Repair method (#2621)
* cba008d7d Add GetObject method to Metainfo (#2611)
* 53d96be44 Stylistic Go Cleanup (#2524)
* 9204c57c0 markup for add stripe card popup (#2570)
* 6f2b85603 Metainfo RPC segment methods (part 1) (#2567)
* 30eeb6481 web/satellite add card button on account payment methods disabled (#2613)
* 499b5db83 web/satellite info button on create account screen fixed (tooltip created) (#2612)
* 537d1cf09 Add test to ensure correct order of repair queue selection (#2551)
* 6c1c3fb4a Add metainfo loop service (#2563)
* 6143be591 V3-1839 web satellite auth cookie subdomain differentiation (#2601)
* 5b73180f9 Disable bandwidth rollups until duplicate data issue is resolved (#2606)
* 6778caf84 satellite/console: add referral link logic (#2576)
* 13dd50104 storagenode/storagenodedb: move tests near the interface rather than the implementation (#2596)
* 63c1a050f pkg/dht: remove (#2599)
* 665a1e386 close test project/uplink (#2605)
* 5fe95fe8f kademlia: restrict maximum find near query (#2598)
* ff31f865a web/satellite hover to cross buttons added (#2603)
* 29b576961 value attribution merge fix and more test cases (#2588)
* 5f096a3ea Remove GetValid, add GetAll to vouchers DB (#2594)
* 17a2e81af web/satellite buckets empty state and search bugs fixed (#2597)
* af93366fb V3-2076 Do not allow users to type more than the allowed number of ch… (#2600)
* a7cc94077 Nodes should not be able to fail the same audit multiple times (#2404)
* 3e9f90b7e satellite/satellitedb: fix nil value from offers table (#2587)
* 2e3ad0ce2 Use trusted pool to get satellite addr in storagenode orders send (#2590)
* b4c51e974 update the uplink CLI copy (#2593)
* 3af925065 update irreparableDB.GetLimited query to use where instead of offset (#2585)
* 91f0adef1 Add the ability to set dial and request timeouts from the cmd/uplink CLI to libuplink. (#2439)
* 848eb8c02 remove kademlia from vouchers package (#2589)
* 8fb70aed1 Satellite frontend refactoring (#2537)
* 416fa80e8 Link Sharing Service (#2431)
* af7ffb807 v3-2156: Add partner_id on user creation (#2571)
* 0931f0e71 Log fatal errors with respective severity (#2581)
* f6f65a80d storagenode/trust: implement fetching peer identity without kademlia and endpoint (#2584)
* 20a7748fb web/satellite close cross symbol bug fixed (#2583)
* 0ca6abc81 scripts/atomicalign: check for proper alignment for 64bit atomics (#2580)
* 38a40088c Remove orphaned tmp data from Storagenodes (#2582)
* 5f194b453 SNO Dashboard API  (#2427)
* 4ebaae26d Disabling Dashboard Link, as its localhost and not working yet (#2579)
* 89afe3ee3 remove struct to ensure 64bit alignment for atomics (#2578)
* 0b827f137 Speedup metainfo tests (#2556)
* de300e923 Network Wipe (Pre Beta) (#2566)
* d04461367 SN DB Optimization: Add rollups to bandwidth usage (#2541)
* d453cd148 Consider encryption overhead when validating max segment size (#2569)
* 260d9c49a Metainfo RPC objects methods (#2534)
* b51d3a69d enable utccheck in all tests by default (#2565)
* c12a4aed3 satellite/console: add redeem reward method (#2484)
* 8b2d46a97 add vouchers to QueryRequest (#2559)
* 3a34a0df7 repair: fix data race in reliability cache (#2561)
* b9d8ddaad storagenode: remove datetime calls in favor of UTC (#2557)
* aa99482fa Jg/add tests (#2547)
* 0d1dce508 ensure uplink is sending correct size with PieceHash (#2555)
* 5bec82014 Fix monkit leaking (#2553)
* 002d9748e signing: ensure we don't break signatures (#2542)
* a2418b22a storagenodedb: optimize index usage and queries (#2545)
* 3a01106ab quick fix (#2548)
* daa3b32ee Add Attribution Columns to appropriate tables for OSPP referral link  (#2516)
* b590e53d6 Order by attempted time in injured segments select (#2533)
* 3cc45a02f Update vouchers.go (#2544)
* 07b23b7cc update tally to rm bucket count (#2539)
* 64b2769de discovery: parallelize refresh (#2535)
* d887ffec6 satellite/satellitedb: add default offer for offers table (#2522)
* 003828c34 web/marketing/pages: partner UI table (#2523)
* 6e57b102c satellite/satellitedb: remove num_redeemed in offers table (#2510)
* f11bf46a1 Jg/1967 mv bucket metadata uplink (#2505)
* f420b29d3 [V3-1927] Repairer uploads to max threshold instead of success… (#2423)
* 8a9c63809 tests for PiecePublicKey/PiecePrivateKey (#2526)
* d52f764e5 protocol: implement new piece signing and verification (#2525)
* 8b507f3d7 Address concerns with storagenode Retain endpoint (#2527)
* 268c629ba Replace base64 encoding for path segments (#2345)
* 0f36b160f Higher dial timeout for TestDownloadSharesDownloadTimeout (#2500)
* de85d1706 Add checker metrics (#2487)
* 0e463dccf 7 day validity window for order limits (#2520)
* 32e0227c4 Project Payment methods (#2037)
* 02565db73 storagenode: migration to drop unused index and used_serials data (#2508)
* f06aec06f Move int64s to top of struct to resolve alignment issue on ARM (#2521)
* 0d294103e satellite/rewards: nicer offers handling (#2390)
* 94eeb58b4 mobile: add a way to get an encryption access at some path root (#2519)
* 10547cc1e segments: send in the object path to the initial CreateSegment… (#2518)
* 7886a4d7b storagenodedb: use datetime functions in sqlite queries (#2512)
* c29f033e7 Move kademlia dialer into separate package (#2466)
* fa1f5c8d7 garbage collection endpoint on storage node (#2424)
* 8f024b3db scss beautification (#2499)
* 13e917f9e [Mobile] Add method for EncryptionAccessWithDefaultKey  (#2479)
* a79c7d77f overlay cache: slight modification of node-is-online rules (#2490)
* 1c5db71fa Change protobuf expirations to use time.Time (#2509)
* bbc25a2bf Drop SN `certifiates` table from DB (#2498)
* d616be8ae storagenode: use minimum time in the order for expiration (#2504)
* 910e4744a satellite/metainfo: optimize pointerdb size (#2506)
* 4faed2098 method name changer per PR 2469 (#2494)
* 6aefa1ae8 adds metainfo rpc doc (#1957)
* 203f36a41 rename order.go to order.pb.go (#2496)
* c5de36ed1 update to go 1.12.7 (#2493)
* 7212423c5 Link Sharing service design (#2360)
* aefb77506 frontend template tabs fixed (#2497)
* e81e25c27 Web satellite master fix (#2495)
* 671648012 Update password based key derivation design doc (#2367)
* 50d601ab0 pkg/identity: Use identity error class (#2488)
* 93cd05d60 error icon changed (#2475)
* 887b784f5 discovery: use fetch info to ping (#2491)
* ff6f1d1b3 storagenode: add in-memory tracking for bandwidth and disk usage (#2469)
* f9696d6c5 satellite/metainfo: add buckets RPC and database (#2460)
* 16156e3b3 Ensure we force a segment size and account storage before committing them (#2473)
* 3587e1a57 Change pointerdb pointer to use time.Time for Creation date (#2483)
* 674742d1a satellite/datarepair: use reliability cache (#1976)
* a786e4249 pkg: Align errs Class messages (#2485)
* e36f43c47 satellite/rewards: replace iota with hardcoded OfferType value (#2432)
* 19ab9852f Update node.proto to use time.Time instead of timestamp (#2482)
* a554752ce adding expiration date check on satellite's CreateSegment (#2476)
* dcf8e2936 Update vouchers to use time.Time instead of timestamp (#2478)
* 8d15d774b Removing trace of encodedPiece.Read() because it's fast. (#2477)
* ccef5eee4 Add proper Version Handling to Identity, Gateway and Uplink Binary (#2471)
* 524eb24c8 storagenode/nodestats: combine stats into single RPC call (#2455)
* 65aa8f227 piecestore: pipeline chunks with orders (#2451)
* 88732188c Update inspector timestamp to time.time (#2464)
* dbc07fa86 Satellite frontend tabs fixed (#2465)
* 094e1b8b9 Add 'Old' suffix to some metainfo methods/messages (#2462)
* 963e1b971 Rename voucher service to endpoint (#2467)
* 3f4662598 storagenode/piecestore: add piece_creation field (#2441)
* e3a9c2db9 fix ordering of protoc includes (#2308)
* 7aca0eb28 storagenode/dashboard: show console address (#2456)
* ea9a8b37b pkg/datarepair/checker: Fix typo in doc comment (#2461)
* 681f502b0 lib/uplink: ensure tests are compiled separately (#2440)
* f9ed0dc1a Improve stability of TestDownloadSharesDownloadTimeout (#2210)
* b5f8a536e Satellite console frontend data reload (#2457)
* e0a8937c6 pkg/storage/ec: Make minor improvements in client.Repair (#2454)
* f275f2a0f uplink mobile: add passphrase generation helpers (#2428)
* 8b07df37f refactor some store encryption stuff (#2434)
* 537c6021d fixes issue where both 36 byte and 16 uuids have duplicate time entries after deploying new version with 16 byte uuid (#2450)
* 945513517 scripts/protobuf: use command-line diff (#2458)
* 42c125d69 Satellite nodestats endpoint add audit check (#2426)
* ce4b99762 storagenode/nodestats: connection leak (#2443)
* 0587dd79d change bucket table to fix conflict with pointerDB (#2452)
* 5176e4f09 Rename BucketItem into Bucket (#2442)
* 5557d557f storagenode/consoledb: fix daily bandwidth query (#2446)
* 56fcf228f Enable bucket name validation (#2449)
* ae8b9698f Rename/remove EncryptionScheme -> EncryptionParame… (#2363)
* d499d162f implement storj.NodeURL in trusted satellites (#2388)
* c0097bbc3 design doc: path component encoding (#2386)
* 2ee5bada2 Add pieceNum to PieceID derivation function (#2193)
* 61dfa61e3 Add timestamp and piece size to piece hash (#2198)
* 2f6e193c3 Remove long tail timer in ecclient (#2433)
* ca0058c9f Set MinDownloadTimeout to 5s in testplanet (#2447)
* 46b5c30f3 Fix test-sim debug-addr error (#2420)
* 38f3d860a storagenode: decline uploads when there are too many live requests (#2397)
* f1e670c9b go.mod: update grpc to 1.22.0 (#2444)
* cded0b412 Properly format Log Message (#2438)
* 47e4584fb V3-1989: Storage node database is locked for several minutes while submiting orders (#2410)
* 42e124cd0 monitor optimal wait fraction (#2435)
* 9e26149a2 pkg/cfgstruct: add pflag support to bind (#2425)
* 6d55bbdb5 OrderLimit creation date time limit (#2412)
* 0de4cf31c libuplink scope type (#2382)
* 3583c65f5 Move from allowed range to minimum for Version Control (#2421)
* d32c90744 overlay.UpdateStats removes node from containment mode (#2419)
* f56cb1c61 satellite/console: add get current reward offer to console (#2341)
* 9e8ecb630 Storagenode nodestats at daily space usage (#2422)
* f4dfb8fb9 Add daily node space usage to satellite nodestats endpoint (#2361)
* 52e5a4eee pass logger into repairer and ecclient (#2365)
* 16cd1fde8 Storagenode add daily bandwidth usage query for SNO (#2348)
* 3337087a9 Add storj.StreamID (#2407)
* 13fcbb655 Update Readme to actual needed Go Version (#2415)
* a2362f92d Rollback uptime disqualification (#2417)
* a7aa854df Routing table antechamber should expire old nodes that failed connection (#2387)
* 9e62423f4 reduce vetting requirement (#2416)
* ea55ae63e add bucket metainfo rpc (#2383)
* d3b9c77d6 removing Jira in readme (#2413)
* 699ccea19 Creates Routing Table Antechamber (#2318)
* 0158b1eb5 add bucket metadata table in SA masterDB (#2369)
* e8605d312 satellite/rewards: use USD type (#2384)
* 74e12a0bb Fixed typo in disqualification.md (#2396)
* 314837dd6 Fix smtp sender return error on close (#2162)
* fb6686785 fix repair queue (#2405)
* 4ba212107 add back testrand (#2408)
* b6ae6a4e2 pkg/datarepair/checker: Fix doc comments (#2401)
* 385c04672 pkg/pb: rename Order2 to Order, OrderLimit2 to OrderLimit (#2406)
* 3ffb42483 account for disqualified nodes in data repair tests (#2315)
* 3f643551e remove flakiness in TestDataRepair and TestSegmentStoreRepair (#2335)
* 2b68a7242 internal/testplanet: ensure that metainfo connections get closed (#2381)
* 0c52f40d9 cmd/internal/assets: package for embedding resources (#2374)
* 4f0e43796 Audit tests cleanup (#2402)
* fd2e70858 pkg/storage/segments: Abort repaired test unmet requirements (#2403)
* 9a1e1d84f go.mod: enforce go 1.12 (#2399)
* 8a5999953 Revert "miscommit add debug info"
* 512f3fa93 add debug info
* 477ac876a Fix typo in sender.go (#2395)
* 725f8174b Add nerdatwork to CLA List (#2394)
* ff92a5bcf Improve RS validation message (#2377)
* b750bc2d0 Keep some unhealthy pieces in the event of a partial repair (#2380)
* 7e0c5d51d pkg/storj: add NodeURL type for verified url-s (#2379)
* ebd976ec2 satellite/marketing: Create New Offer (#2186)
* 53e550f7d uplink docs (#2372)
* 827fb92b4 satellite/console: nicer error handling in tests (#2378)
* f3dc86688 Set correct default erasure share size for libuplink (#2376)
* 7bbd3bcbb jenkins: adjust go test parallelism (#2364)
* 1a65e42d3 rename EncryptionCtx back to EncryptionAccess (#2368)
* 811168e2c Uplink bucket attribution extension (#2353)
* e83ebd7cd jenkins: avoid using goimports and distribute load better (#2359)
* 261750252 edit voucher denied message (#2362)
* efcdaa43a lib/uplink: encryption context (#2349)
* 27c92ffc1 jenkins: Stop docker container after 18m (#2358)
* 2128b460b cmd/uplink/cmd: don't create benchmark data on init (#2351)
* a6db2dd33 Storagenode add nodestats client for SNO console (#2287)
* 043a5b6fa lib/uplink: unexport bucket attribution check (#2352)
* 7facda442 value attribution integration with libuplink (#2297)
* 615bfca13 Fix TestGetSignee flakiness (#2350)
* 8bf7c5c67 SNO Dashboard http status codes updated (#2333)
* fbe9696e9 pkg/kademlia: clean up peer discovery (#2252)
* 3925e8458 pkg/eestream: plumb contexts through (#2187)
* 7b66e0cd7 Use dial to clarify that it's internally closing the connection. (#2347)
* b6ad3e9c9 internal/testrand: new package for random data (#2282)
* 6b23380d6 Make RS validation always enabled (#2336)
* 11016b506 Moving verbose info logging to debug (#2346)
* ae3697980 Take advantage of Is and IsFunc from zeebo/errs (#2310)
* caa2fcf62 satellite/orders: don't panic (#2331)
* 05a283f33 jenkins: disallow files over 600KB (#2323)
* 0e528bc56 Add attribution report to the satellite CLI (#2288)
* c7679b9b3 Fix some leaks and add notes about close handling (#2334)
* 605e3fb73 v3-2023: Add migration for project_id change in V3-2010 (#2332)
* 57ef352b3 Update Wizard to allow more easily addition of satellites and move package to cmd folder (#2340)
* 2cc01c789 Rename Satellites in the Wizard (#2339)
* 20de18fee [v3-1952 test 6] Disqualification is a one way ticket. Storage Node can't get around it. (#2328)
* c28f80009 Skip TestDataRepair and TestUplinksParallel, because they are flaky (#2337)
* e5c48fab7 fix ordersDB methods to take correct args (#2314)
* 414648d66 Fix some metainfo.Client leaks (#2327)
* 35f2ab5de Enable node selection tests (#2316)
* c42046601 Dockerfile changes for marketing resources. (#2329)
* 13ce28bf3 Add Alexander Bender (#2330)
* 841bb83a6 fix ts lint build error (#2325)
* 1f5870891 Delete all Tardigrade Satellite Data from SNO's (#2324)
* eb5dc4ba2 Add node stats GRPC endpoint to the satellite (#2281)
* 6502143e7 fix import ordering (#2322)
* fa802dc42 pkg/audit: Improve Reporter.ReporterAudits doc (#2266)
* c35c8e4c2 allow reading bucket metadata with restricted keys (#2321)
* 8c57434de pkg/process/metrics: add an instance prefix (#2190)
* b3da72c21 Fix TestVoucherService (#2317)
* 7bf5a0963 Use ErrNodeOffline when returning error for offline node (#2312)
* 70f28ae41 [v3-1952 tests 4 & 5] DQ nodes should not be used for download or upload  (#2272)
* bbedff12a satellite: rearrange marketing package (#2268)
* 30f790a04 Create and use an encryption.Store (#2293)
* eb1b1c434 scripts/tag-release.sh: libuplink release tagging (#2256)
* 5fb2e0191 web/satellite: create ReferralStats component (#2309)
* e17908a26 metainfo: add project info rpc for getting a project salt using an api key (#2311)
* fd6a4d96f change uptime dq threshold to 0.4 (#2313)
* 20041d43b hovers on icons, apikey name errormessage added (#2307)
* 01beaa289 Mask IP Addresses to subnets (#2305)
* 96bc0ccfa SNO Dshboard initial api endpoint added (#2284)
* 8226024ca Do not use disqualified nodes when asking for get order limits (#2303)
* fc11a17df initial design doc for SN graceful exit (#2077)
* 75d71f70c Stop pinging disqualified nodes (#2306)
* fdeb83480 Bucket name validation (#2244)
* e9e68c842 kademlia audit gating design draft (#1941)
* e285fe199 Don't require encryption keys for project or bucket management (#2291)
* 894845916 do ip filtering in a more correct way? (#2301)
* d1e9829e8 preserve reputation migration (#2295)
* 1283036e3 add storage node voucher request service (#2158)
* 81f1bc19d add functions for password/root key derivation (#2294)
* d8e62bc06 support value attribution endpoint (#2231)
* 930481792 Uplink C bindings part 4 (#2260)
* 043d603cb satellite rs config check with validation check set to false default (#2229)
* 8f47fca5d Remove audit / uptime ratio fields (#2247)
* 69c2cbacd internal/testplanet: improve test speed (#2280)
* 3d6b25a04 [v3-1952 test 1 & 3] pkg/audit: Add DQ test for too many failed audits (#2265)
* 23e081f0c storagenode: delete piece when upload is cancelled (#2286)
* 4f2e893e6 Fix the way project_id is stored in bucket_storage_tallies and bucket_bandwidth_rollups (#2283)
* bfcfe3931 Enable and fix statdb tests (#2270)
* 169fc9594 pkg/audit: fix maxRetries bug (#2271)
* 86ec9b957 jenkins: capability to use leakcheck (#2179)
* d103dd2c4 Include information about disqualified nodes in the storage node payment report (#2239)
* 09e55ca28 jenkins timeout workaround (#2285)
* 5b030062c internal/testplanet: split planet.go file to avoid package naming conflicts (#2279)
* 5f47b7028 Uplink C bindings part 3 (#2258)
* a5baebfa6 reject invalid orders (#2262)
* aa25c4458 kvmetainfo: merge with storage/buckets (#2277)
* 568b000e9 satellite: make order expiration configurable (#2251)
* 76b54458e satellite: send external address in order limits (#2278)
* 24918e072 Include Node ID and Piece ID in piecestore errors (#2261)
* 06006effb lib/uplink: minor fixes (#2257)
* d435d4859 internal/testplanet: add missing consoleserver.Config (#2269)
* daf166d42 fix tally test flakiness (#2250)
* edb3d1cbf pkg/overlay: update node selection config values for reputation (#2264)
* 5b3108675 pkg/audit: Move test helper funcs to separated file (#2259)
* 9386187fe add disqualification and new reputation system into overlay cache (#2227)
* 1c2067a09 Create graphql query for getting user credit usage  (#2255)
* e5fd0287e V3-1819 Storage node operator server and service started (#2112)
* 1726a6c07 pkg/bloomfilter: implementation and benchmark results (#2113)
* d583ab707 add user credit usage method into console service (#2240)
* 213bcacc8 Adding Andrew Harding (#2254)
* 964c87c47 Fix checks around repair threshold (#2246)
* 6ba110c90 Bind Jenkins Builds to specific nodes (#2253)
* 35e71e5f9 cmd/storj-sim: ensure prefix writer can be used concurrently (#2243)
* 402902563 Update pr template (#2237)
* 3706bd78e gerrit: add cla
* 2b3088745 gerrit integration
* 9d670b33b reduce expected disk size for storj-sim and testplanet (#2249)
* d5f89afac fix merge resolution fail (#2248)
* ddcf4fc2a add support to hide config settings (#2241)
* 0eb8ac58c Create protocol buffer for garbage collection (#2222)
* 7b450927f code updates with comments (#2245)
* 8f4a6afc8 Fixes for value attribution (#2238)
* ac65d1975 Avoid reporting audit failures for files deleted during audit (#2233)
* f0f59a557 Always encrypt inline segments without padding (#2183)
* ad8cad490 Expand the inspector tool to provide node id's for each segment, rather than just numeric totals (#2205)
* b30c35d30 change ReputationAuditOmega (et al.) to AuditReputationWeight (#2232)
* 119a8fd3c removed fields (#2234)
* 89cd4d7bc Use a persistent cache and pkg directory (#2235)
* 1334933cd Console: Remove payments checks during project creation and member deletion (#2236)
* 954ca3c6e add db implementation for user_credits table (#2169)
* c481e071b Rework gateway container image to build like the rest (#2230)
* 81134e97c nodes db - adds alpha and beta reputation fields (#2215)
* 503b951bc modify build scripts for satellite ui (#2228)
* 09940d4e0 value attribution DB interface Insert & Get method support (#2200)
* 8398fae9b Add node churn and containment/reverify monkit stats (#2217)
* 78db9d278 Fix project email sign in and project links and comments (#2166)
* 35c864833 [v3-1914] Storage node disqualification: Change type from bool to timestamp (#2212)
* 0e39b93dc Design Doc Notification System (#2101)
* d177761f5 Adds test coverage to CalculateBucketAtRestData (#2203)
* 7795ba816 update UI to reflect final mockups (#2175)
* f4f776d09 Use mail.test as domain in emails (#2224)
* 8e29ef8a6 Use zap.Stringer instead of zap.String (#2223)
* b1e5cf120 add index on pieceinfo expireation for faster GetExpired calls (#2220)
* b8bced690 improve logging (#2219)
* f378125c8 draft of removing stats from inspector (#2226)
* e58a06bd0 config: update release values to match prod (#2192)
* 76061c50c Uplink C bindings part 2 (#2209)
* cc48a6366 docs/design: Fix typos in disqualification (#2211)
* 1a5483221 add fake id so there's always a segment instance created on Vue (#2214)
* 7d5e425d8 monkit: start reporting 10th percentile, remove init.ializers from everything (#2173)
* 7a68b1090 fix satellite gui build (#2204)
* 319cc77a3 increase audit timeout (#2208)
* 0aae26d78 gitignore: cleanup old things (#2197)
* ebd9b375f Repair should not corrupt files (#2194)
* 6932dc92f NodeSelectionConfig inits and test cleanup (#2202)
* 4e1cc37a4 Fix bucket_storage_tally data (#2165)
* df2fad15d pkg/process/logging: different defaults for release/dev (#2191)
* fa09a7894 pkg/process: prometheus support (#2176)
* 2fdcd8c63 Give the uplink image the same business (#2168)
* 703e5fee4 mostly removed things (#2199)
* bc3396472 Uplink C bindings part 1 (#2196)
* 1ae5654eb kademlia/routing: add contexts to more places so monkit works (#2188)
* ec63204ab fix export name for DatePicker (#2185)
* c0fe061a6 add user credits table (#2085)
* 1a1a08447 testcontext: sanitize folder name (#2195)
* 380684e76 Removing bucket_id and adding project_id and bucket_name as primary key (#2182)
* 8e8a5eb89 initial value attribution check development (#2091)
* ff7a9030e lib/uplink: expose restrict on api keys (#2189)
* 474d9e749 only parse templates for the first requests (#2180)
* e85c56518 Add Isaac (#2184)
* 6a826c53c fix satellite gui build issue (#2181)
* a4d6e8c8c split update methods into two for offer data (#2133)
* 14486b188 satellite/orders: TODOs about transactions (#2134)
* fdddaa2a4 Fix marketingweb code (#2177)
* 5314c950c pkg/audit: add more reverify tests (#2144)
* df1401d95 add segment.io into satellite GUI (#1946)
* d84cd719f add const stripesPerBlock const to calc blocksize (#2163)
* af66d9c6e Open a new port on satellite for admin GUI (#1901)
* 252c8ac18 Add email to self node info (#2171)
* f624b213a reputation: Add configuration parameters (#2150)
* f0880b9b3 Audit service should download erasure shares in parallel (#2164)
* 5e7fed754 improve logging (#2170)
* 61c8c3fb4 satellite/orders: batch update storage node allocations (#2146)
* 7a28decfc Always use difficulty larger than 8 for identites (#2160)
* 74484fc57 Enforce our Minimum Requirements for Node Operators and sanity check them (#2155)
* 749846b42 Update golangci-lint (#2159)
* 8f2dca843 Re-enabling and fixing repairer tests (#2099)
* 469d485f6 dbutil: reduce defaults so sum is < 100 (#2157)
* 3a7ebea76 Moved monkit tracing after the limit nil check & fix node id formatting. (#2156)
* 16acb5854 rename verifier test files (#2154)
* 23587bba0 Storagenode vouchers table (#2121)
* f5227abd3 uplink: enc.encryption-key flag is only available for setup command (#2090)
* 2919aa2f7 Reputation / node selection doc (#2135)
* 29b3cec68 update to the reporting section to calculate bandwidth (#2151)
* 44359e8bb pkg/kademlia: Fix bug using ref of var range loop (#2149)
* 8ddb1dee0 Improve error analysis in audit service (#2125)
* cac463767 satellite: Fix Peer.Run method doc comment (#2147)
* 30784484c Fixed error notification on fail project member deletion. (#2148)
* 43d4f3daf discovery: remove graveyard (#2145)
* a55df84bf Adding RS values to the info log and NodeID to tracing. (#2130)
* 03fece56d Ensure Storage Nodes collect expired used serial numbers (#2143)
* bf3d168cf makes sure all uplink cli configs get passed to libuplink, add stripeSize (#2103)
* f69f7b2d5 frontend bug fixes (#2127)
* 51db703b2 Console add payments to service (#2100)
* 206c3a20f Web/reload issue (#2139)
* 0fc8fd533 Web/register success popup close issue (#2141)
* 91a0ece95 Remove bwagreement leftovers (#2140)
* 0bf50c682 sync2: prioritize Cycle stopping messages (#2137)
* 28a120159 pkg/bwagreements: remove service (#2138)
* 503fe0e5e pkg/audit: Fix code style named result params (#2136)
* 2d19475d6 monitor macaroon parse/serialization perf (#2132)
* 0bda2f535 node disqualification design draft (#2079)
* 25d7dda13 add disqualified check to node selection queries (#2102)
* 099cf921d Fix flaky audit test (#2093)
* cc47bf2df postgreskv: actually monitor the iterate call (#2129)
* b5ac4f3ea Better metainfo Create/Commit request validation (#2088)
* b9d586901 Value Attribution DBX Model (#2116)
* f360097b2 satellite/orders: monitor order settlement more (#2128)
* d7f3a5f81 internal,lib,uplink: add monkit task to missing places (#2118)
* f1641af80 storage: add monkit task to missing places (#2122)
* 68c8ce63a Fix storing struct in context (#2126)
* ccb158c99 pkg/auth: add monkit task to missing places (#2123)
* 679cdfda9 storage/filestore: add monkit task to missing places (#2124)
* 7bd676722 api keys text bug fixed (#2114)
* a913f8303 Update Notification Text to new process (#2120)
* 23213a7a2 Fix Config Comments in DButil Package (#2119)
* d68c99377 Change for free credit design doc (#2046)
* d02427e41 db: set max open conns, conn max lifetime, add db stat monitoring (#2117)
* 09b0c2a63 create db implementation for offer table (#2031)
* 9ed8c9605 jenkins: check that lock file is updated properly (#1977)
* d6c02fc65 storagenode/piecestore: add meters to upload/download rates (#2106)
* fbdff0fad nobuckets state implemented (#2110)
* d15eaed58 add capability of logging all GRPC calls/payloads (#2067)
* 3b02607c1 go.mod: update grpc to latest version (#2096)
* c4bb84f20 storagenode: add monkit task to missing places (#2107)
* 3fe8343b6 repairer: fix config comments (#2105)
* e2c5d75b4 Add VitaliiShpital (#2111)
* 29d16b4d6 satellite: add monkit task to missing places (#2108)
* 9c5708da3 pkg/*: add monkit task to missing places (#2109)
* bc3463b9d Don't convert to float to avoid losing precision. (#2092)
* e077b0d38 rename VetNode to IsVetted (#2097)
* b8e0ac637 satellitedb/overlaycache: avoid tx leak in case of error (#2095)
* 6809129e6 Console add stripe service (#2080)
* 8c0c51862 pkg/storj: use proper unquoting of json data in NodeID and PieceID (#2094)
* 79a008510 Fix storage and object_count calculations on console (#2081)
* 506243684 Account settings page rework (#2007)
* 2ab95b533 Check errors for possible outcomes from audit's DownloadShares (#2072)
* e60ff9dcb process/metrics: have metrics suffix default to dev/release status (#2073)
* 140251882 fix bug for setting flag only values in process setup (#2089)
* 294fddcec Add audit monkit stats (#2087)
* 5b730e307 Make maxReverifyCount configurable (#2071)
* 4ad512092 Checker service refactor (v3-1871) (#2082)
* 6db938808 add disqualified column to nodes table (#2086)
* 590b1a5a1 Satellite voucher service (#2043)
* 24c813297 rewrites encode methods for readability (#2084)
* 4b75752d6 Relocate IP Lookup to ensure it is always set (#2061)
* 04c20b0ac Add monkit stats to piecestore upload/download (#2078)
* 934ebf9cb Added the irreparable repair functionality (#1955)
* 006509b75 Value Attribution design doc (#1996)
* e74cac52a Command line flags features and cleanup (#2068)
* 771271e7b updating help message on storage node for AllocatedBandwidth (#2074)
* 8918a14bb add payment popup added (#2032)
* 16e3b77cf Enable Scopelint Linter (#2049)
* 268dc6b7e Enable gocritic linter (#2051)
* 20782d75b Fix dates source on UsageReport page (#2070)
* 03bae48ad Remove unused parameter (#2069)
* fb86238ac aws s3 performance tests (#2060)
* dc2e6b937 Changing 64bit atomics to 32bit for alignment on ARM. (#2066)
* f731267e8 Per-project usage limiting (#2036)
* c07162bee address potential divide by 0` (#2065)
* 8912d7149 Fixes Auth Issue (#2064)
* 5a4ff2c85 add repair monkit stats (#2045)
* aa6ff17b7 add Reverify to auditing (#2041)
