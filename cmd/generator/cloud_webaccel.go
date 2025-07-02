package generator

import (
	"github.com/g1eng/savac/cmd/helper"
	"github.com/urfave/cli/v3"
)

func (cg *CloudCommandGenerator) GenerateWebAccelCommand() *cli.Command {
	updateSiteFlag := []cli.Flag{
		&cli.StringFlag{Name: "request-protocol", Usage: "request protocol for the site", Aliases: []string{"p"}, Value: "http+https"},
		&cli.IntFlag{Name: "default-cache-ttl", Usage: "default cache TTL (on your own risk)", Aliases: []string{"ttl"}},
		&cli.StringSliceFlag{Name: "cors", Usage: "CORS allowed origins (can be multiply used)"},
		&cli.BoolFlag{Name: "vary", Usage: "enable vary support", Value: false},
		&cli.StringFlag{Name: "accept-encoding", Usage: "accept encoding (gzip/brotli)", Aliases: []string{"ae"}},
		&cli.StringFlag{Name: "origin-protocol", Usage: "origin protocol (http/https)", Aliases: []string{"op"}, Value: "https"},

		&cli.StringFlag{Name: "origin-type", Usage: "origin type for the site (web/bucket)", Aliases: []string{"ot"}, Value: "web"},
		&cli.StringFlag{Name: "origin", Aliases: []string{"o"}, Usage: "origin destination (option for a `web` origin)"},
		&cli.StringFlag{Name: "host-header", Usage: "host header for the access to the origin (option for a `web` origin)"},

		&cli.StringFlag{Name: "bucket", Aliases: []string{"b"}, Usage: "origin bucket name (option for a `bucket` origin)"},
		&cli.StringFlag{Name: "endpoint", Usage: "S3 endpoint with TLS support (option for a `bucket` origin)", Sources: cli.EnvVars("SAKURASTORAGE_ENDPOINT"), Value: "s3.isk01.sakurastorage.jp"},
		&cli.StringFlag{Name: "region", Usage: "S3 region (option for a `bucket` origin)", Sources: cli.EnvVars("SAKURASTORAGE_REGION")},
		&cli.BoolFlag{Name: "docindex", Usage: "whether document index is enabled/disabled (option for a `bucket` origin)", Value: false},
		&cli.StringFlag{Name: "access-key", Usage: "access key for the S3 endpoint", Sources: cli.EnvVars("SAKURASTORAGE_ACCESS_KEY")},
		&cli.StringFlag{Name: "access-secret", Usage: "access secret for the S3 endpoint", Sources: cli.EnvVars("SAKURASTORAGE_ACCESS_SECRET")},
	}
	return &cli.Command{
		Name:    "webaccel",
		Aliases: []string{"wa"},
		Usage:   "Configure Web Accellerator",
		Commands: []*cli.Command{
			{
				Name:      "create",
				Usage:     "create a new site",
				Before:    helper.CheckArgsExist,
				ArgsUsage: "<name>",
				Flags: append([]cli.Flag{
					&cli.StringFlag{Name: "domain-type", Aliases: []string{"t"}, Value: "subdomain"},
					&cli.StringFlag{Name: "domain", Aliases: []string{"d"}, Required: false},
				}, updateSiteFlag...),
				Action: cg.ag.GenerateWebAccelSiteCreateAction,
			},
			{
				Name:   "list",
				Usage:  "list sites",
				Flags:  patternMatchingFlags,
				Action: cg.ag.GenerateWebAccelSiteListAction,
			},
			{
				Name:      "read",
				Usage:     "get site information",
				Before:    helper.CheckArgsExist,
				Flags:     patternMatchingFlags,
				ArgsUsage: "<siteId>",
				Action:    cg.ag.GenerateWebAccelSiteReadAction,
			},
			{
				Name:      "update",
				Usage:     "update the site configuration",
				Before:    helper.CheckArgsExist,
				ArgsUsage: "<siteId>",
				Flags:     updateSiteFlag,
				Action:    cg.ag.GenerateWebAccelSiteUpdateAction,
			},
			{
				Name:      "enable",
				Usage:     "enable the site",
				Before:    helper.CheckArgsExist,
				ArgsUsage: "<siteId>",
				Action:    cg.ag.GenerateWebAccelSiteUpdateStatusAction(true),
			},
			{
				Name:      "disable",
				Usage:     "disable the site",
				Before:    helper.CheckArgsExist,
				ArgsUsage: "<siteId>",
				Action:    cg.ag.GenerateWebAccelSiteUpdateStatusAction(false),
			},
			{
				Name:    "origin-guard",
				Aliases: []string{"og"},
				Usage:   "manipulate origin guard token",
				Commands: []*cli.Command{
					{
						Name:      "read",
						Usage:     "read origin guard token for the site",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelReadOriginGuardTokenAction,
					},
					{
						Name:      "create",
						Usage:     "create the new origin guard token",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						UsageText: "create the new origin guard token\n" +
							"" +
							"Be careful not to shut out current traffic to the non-cached contents.\n" +
							"Use --next flag to create next token and ensure access to the origin with\n" +
							"the new token at first.\n" +
							"And don't forget to finalize the next token after that to run `create` again." +
							"",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "next",
								Usage: "create the next origin guard token to prepare a migration",
								Value: false,
							},
						},
						Action: cg.ag.GenerateWebAccelCreateOriginGuardTokenAction,
					},
					{
						Name:      "delete",
						Usage:     "delete origin guard tokens for the site",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "next",
								Usage: "delete the next origin guard token to cancel the migration",
								Value: false,
							},
						},
						Action: cg.ag.GenerateWebAccelDeleteOriginGuardTokenAction,
					},
				},
			},
			{
				Name:      "one-time-url",
				Aliases:   []string{"otu"},
				Usage:     "configure a secret or get the one-time URL with it",
				Before:    helper.CheckArgsExist,
				ArgsUsage: "<siteId>",
				Commands: []*cli.Command{
					{
						Name:  "secret",
						Usage: "get or set site-wide one-time secret",
						UsageText: "set site-wide secret for the one-time URL.\n" +
							"In web accelerator service, one or more one-time secrets result site-wide protection.\n" +
							"By default, a site holds two secrets set past and the an old secret is not purged automatically.\n" +
							"So you need to purge or overwrite all of them to setup the protection with single secret.\n" +
							"\n" +
							"If you do not need the protection at now, purge secrets using --purge flag." +
							"",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "purge",
								Usage: "purge all preset secret for the site",
								Value: false,
							},
							&cli.BoolFlag{
								Name:    "random",
								Aliases: []string{"r"},
								Usage:   "set random secret string for the site",
								Value:   false,
							},
						},
						Action: cg.ag.GenerateWebAccelSiteWideOnetimeSecretAction,
					},
					{
						Name:      "generate",
						Aliases:   []string{"g"},
						Usage:     "generate one-time url for a path or an URL",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<url|siteId>",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "path",
								Aliases: []string{"p"},
								Usage:   "path for the one-time url (only required without url)",
							},
							&cli.StringFlag{
								Name:    "expired",
								Aliases: []string{"e"},
								Usage:   "expiration date of the one-time URL",
								Value:   "3min",
							},
						},
						Action: cg.ag.GenerateWebAccelOneTimeUrlAction,
					},
				},
			},
			{
				Name:      "delete",
				Usage:     "delete the site",
				Before:    helper.CheckArgsExist,
				Flags:     patternMatchingFlags,
				ArgsUsage: "<siteId>",
				Action:    cg.ag.GenerateWebAccelDeleteAction,
			},
			{
				Name:      "certificate",
				Usage:     "update the certificate configuration",
				Before:    helper.CheckArgsExist,
				ArgsUsage: "<siteId>",
				Aliases:   []string{"cert"},
				Commands: []*cli.Command{
					{
						Name:  "import",
						Usage: "import the endpoint certificate chain (and key, optionally)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "file",
								Usage:    "specify the new certificate chain to be imported",
								Required: true,
							},
							&cli.StringFlag{
								Name:  "key",
								Usage: "specify the new private key to be imported",
							},
						},
						Action: cg.ag.GenerateWebAccelCertificateImportAction,
					},
					{
						Name:      "read",
						Usage:     "read the certificate chain for the site",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "revision", Value: 0, Aliases: []string{"r"}},
						},
						Action: cg.ag.GenerateWebAccelCertificateReadAction,
					},
					{
						Name:      "revisions",
						Usage:     "list revisions of the certificate chain for the site",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelRevisionsAction,
					},
					{
						Name:      "auto-renewal",
						Aliases:   []string{"auto"},
						Usage:     "configure the autorenewal of certificates",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "lets-encrypt", Value: true},
							&cli.BoolFlag{Name: "disable", Value: false},
						},
						Action: cg.ag.GenerateWebAccelCertificateAutoRenewalAction,
					},
					{
						Name:      "delete",
						Usage:     "delete the certificate chain from the site",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelCertificateDeleteAction,
					},
				},
			},
			{
				Name:      "purge-cache",
				Usage:     "purge cache for sites",
				Aliases:   []string{"pc"},
				Before:    helper.CheckArgsExist,
				ArgsUsage: "[url...]",
				UsageText: "purges cache of the specific (or all) site.\n" +
					"\n" +
					"This command accepts valid URLs for registered sites.\n" +
					"If non-registred hostname is specified, the process is terminated.\n" +
					"",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "all",
						Usage: "clear all cache for the site",
						Value: false,
					},
				},
				Action: cg.ag.GenerateWebAccelPurgeCacheAction,
			},
			{
				Name:  "acl",
				Usage: "configure site ACL",
				Commands: []*cli.Command{
					{
						Name:      "read",
						Usage:     "read ACL rules for the site",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelACLReadAction,
					},
					{
						Name:    "apply",
						Aliases: []string{"upsert"},
						Usage:   "apply ACL rule",
						UsageText: "confiugre the access control for the site.\n" +
							"--allow or --deny flag must be specified to set up ACL.\n" +
							"If both of them are specified, deny-last rule is applied.",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{
								Name:  "allow",
								Usage: "allowed prefixes, e.g. `192.0.2.0/26,192.0.2.128/26` ",
							},
							&cli.StringSliceFlag{
								Name:  "deny",
								Usage: "denied prefixes, e.g. `192.0.2.0/26,192.0.2.128/26` ",
							},
						},
						Action: cg.ag.GenerateWebAccelAclApplyAction,
					},
					{
						Name:      "clear",
						Aliases:   []string{"delete", "del"},
						Usage:     "clear ACL",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelAclFlushAction,
					},
				},
			},
			{
				Name:    "logging",
				Aliases: []string{"log"},
				Usage:   "manipulate log upload config",
				Commands: []*cli.Command{
					{
						Name:      "apply",
						Usage:     "create/update the log upload configuration for the site",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:   "endpoint",
								Hidden: true,
								Value:  "https://s3.isk01.sakurastorage.jp",
							},
							&cli.StringFlag{
								Name:   "region",
								Hidden: true,
								Value:  "jp-north-1",
							},
							&cli.StringFlag{
								Name:     "bucket",
								Aliases:  []string{"b"},
								Usage:    "bucket name to store access logs",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "access-key",
								Aliases:  []string{"key"},
								Sources:  cli.EnvVars("SAKURASTORAGE_ACCESS_KEY"),
								Usage:    "IAM access key for the bucket",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "access-secret",
								Aliases:  []string{"secret"},
								Sources:  cli.EnvVars("SAKURASTORAGE_ACCESS_SECRET"),
								Usage:    "IAM access secret for the bucket",
								Required: true,
							},
						},
						Action: cg.ag.GenerateWebAccelAccessLogApplyAction,
					},
					{
						Name:      "read",
						Usage:     "read the access log configuration",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelAccessLogReadAction,
					},
					{
						Name:      "enable",
						Usage:     "enable the access log configuration",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelAccessLogUpdateStatusAction(true),
					},
					{
						Name:      "disable",
						Usage:     "disable the access log configuration",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelAccessLogUpdateStatusAction(false),
					},
					{
						Name:      "clear",
						Aliases:   []string{"detach"},
						Usage:     "clear the log upload configuration for the iste",
						Before:    helper.CheckArgsExist,
						ArgsUsage: "<siteId>",
						Action:    cg.ag.GenerateWebAccelAccessLogDeleteAction,
					},
				},
			},
			{
				Name:  "usage",
				Usage: "show the site monthly traffic report",
				UsageText: "Show the service traffic statics.\n" +
					"If year and month option are omitted, the command shows report for the current month.\n" +
					"If the month option is omitted, the command reports for each month of the year.",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "year",
						Aliases: []string{"Y"},
						Usage:   "specify year to be reported",
					},
					&cli.IntFlag{
						Name:    "month",
						Aliases: []string{"M"},
						Usage:   "specify the month to be reported",
					},
				},
				Action: cg.ag.GenerateWebAccelUsageReadAction,
			},
		},
		EnableShellCompletion: true,
	}
}
