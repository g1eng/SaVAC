package generator

import (
	"github.com/g1eng/savac/cmd/helper"
	"github.com/urfave/cli/v3"
)

func (cg *CloudCommandGenerator) GenerateObjectStorageCommand() *cli.Command {
	permissionKeySubcommands := []*cli.Command{
		{
			Name:      "create",
			Usage:     "create a permission key",
			ArgsUsage: "<permission>",
			Before:    helper.CheckArgsExist,
			Action:    cg.ag.GenerateCreatePermissionKeyAction,
		}, {
			Name:      "delete",
			Usage:     "delete the permission key of the permission",
			Aliases:   []string{"del"},
			Before:    helper.CheckArgsExist,
			ArgsUsage: "<permission>",
			Action:    cg.ag.GenerateDeletePermissionKeyAction,
		}, {
			Name:      "read",
			Usage:     "read the permission key with permission id",
			Before:    helper.CheckArgsExist,
			ArgsUsage: "<permission>",
			Action:    cg.ag.GenerateListPermissionKeyAction,
		},
	}
	permissionSubcommands := []*cli.Command{
		{
			Name:   "create",
			Usage:  "create a new permission",
			Before: helper.CheckArgsExist,
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:  "rw",
					Usage: "specify comma-separated list of buckets to permit readwrite access",
				},
				&cli.StringSliceFlag{
					Name:  "ro",
					Usage: "specify comma-separated list of buckets to permit readonly access",
				},
				&cli.StringSliceFlag{
					Name:  "wo",
					Usage: "specify comma-separated list of buckets to permit writeonly access",
				},
				&cli.BoolFlag{
					Name:    "force",
					Usage:   "skip the check of bucket names and forcibly create the permission",
					Aliases: []string{"f"},
					Value:   false,
				},
			},
			Action: cg.ag.GenerateCreatePermissionAction,
		},
		{

			Name:      "update",
			Usage:     "update a permission",
			Before:    helper.CheckArgsExist,
			ArgsUsage: "permission-id",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:  "rw",
					Usage: "specify comma-separated list of buckets to permit readwrite access",
					Value: nil,
				},
				&cli.StringSliceFlag{
					Name:  "ro",
					Usage: "specify comma-separated list of buckets to permit readonly access",
					Value: nil,
				},
				&cli.StringSliceFlag{
					Name:  "wo",
					Usage: "specify comma-separated list of buckets to permit readonly access",
					Value: nil,
				},
				&cli.BoolFlag{
					Name:    "force",
					Aliases: []string{"f"},
					Usage:   "skip the check of bucket names and forcibly create the permission",
					Value:   false,
				},
			},
			Action: cg.ag.GenerateUpdatePermissionAction,
		}, {
			Name:      "delete",
			Usage:     "delete a permission",
			Aliases:   []string{"del"},
			Before:    helper.CheckArgsExist,
			ArgsUsage: "<permission>",
			Action:    cg.ag.GenerateDeletePermissionAction,
		}, {
			Name:    "list",
			Usage:   "list permissions",
			Aliases: []string{"ls"},
			Action:  cg.ag.GenerateListPermissionAction,
		}, {
			Name:     "key",
			Usage:    "manipulate permission keys",
			Commands: permissionKeySubcommands,
		},
	}

	accountSubcommands := []*cli.Command{
		{
			Name:  "site",
			Usage: "manipulate object account sites",
			Commands: []*cli.Command{
				{
					Name:   "create",
					Usage:  "create an account site (activate the service account)",
					Action: cg.ag.GenerateObjectStorageCreateSiteAction,
				},
				{
					Name:   "delete",
					Usage:  "delete the account site",
					Action: cg.ag.GenerateObjectStorageCreateSiteAction,
				},
				{
					Name:   "list",
					Usage:  "list account sites if exists",
					Action: cg.ag.GenerateObjectStorageListSiteAction,
				},
			},
		},
		{
			Name:  "key",
			Usage: "manipulate global account key",
			Commands: []*cli.Command{
				{
					Name:   "create",
					Usage:  "create a global access key for the site",
					Action: cg.ag.GenerateCreateAccountKeyAction,
				},
				{
					Name:   "list",
					Usage:  "list global access keys",
					Action: cg.ag.GenerateListAccountKeyAction,
				},
				{
					Name:    "delete",
					Usage:   "detele a global access key",
					Aliases: []string{"del"},
					Action:  cg.ag.GenerateDeleteAccountKeyAction,
				},
			},
		},
	}

	subCommands := []*cli.Command{
		{
			Name:      "ls",
			Usage:     "list buckets or objects",
			ArgsUsage: "[s3uri]",
			Action:    cg.ag.GenerateObjectStorageListActionMeta,
		}, {
			Name:      "mb",
			Usage:     "make a bucket",
			Before:    helper.CheckArgsExist,
			ArgsUsage: "bucket_name",
			Action:    cg.ag.GenerateCreateBucketAction,
		}, {
			Name:      "put",
			Usage:     "put an object",
			ArgsUsage: "<local-src> <s3uri>",
			UsageText: "This command upload one or more objects into the specified destination.\n" +
				"This command can accept stdin, using - as the first argument.\n" +
				"The destination must be valid S3 URI.\n" +
				"\n" +
				"If stdin or a large file (>=5GiB) is given, multipart upload is performed.",
			Before: helper.CheckArgsExist,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "recursive",
					Aliases: []string{"r"},
					Usage:   "upload objects in the directory recursively",
					Value:   false,
				},
			},
			Action: cg.ag.GeneratePutAction,
		}, {
			Name:      "get",
			Usage:     "get an object",
			Before:    helper.CheckArgsExist,
			ArgsUsage: "<remote-src> <dest>",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "recursive",
					Usage:   "copy object recursively",
					Aliases: []string{"r"},
					Value:   false,
				},
			},
			Action: cg.ag.GenerateGetAction,
		}, {
			Name:      "check",
			Usage:     "check the existence of an object",
			Before:    helper.CheckArgsExist,
			ArgsUsage: "<s3uri>",
			Action:    cg.ag.GenerateCheckAction,
		}, {
			Name:      "rm",
			Usage:     "remove an object",
			Before:    helper.CheckArgsExist,
			ArgsUsage: "<s3uri>",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "recursive",
					Aliases: []string{"r"},
					Value:   false,
				},
			},
			Action: cg.ag.GenerateRmAction,
		}, {
			Name:      "rb",
			Usage:     "remove a bucket",
			Before:    helper.CheckArgsExist,
			ArgsUsage: "bucket_name",
			Action:    cg.ag.GenerateDeleteBucketAction,
		}, {
			Name:     "permissions",
			Usage:    "manipulate permissions",
			Aliases:  []string{"perm"},
			Commands: permissionSubcommands,
		}, {
			Name:     "account",
			Usage:    "manipulate account-global resources",
			Commands: accountSubcommands,
		},
	}

	return &cli.Command{
		Name:                  "o12",
		Usage:                 "manipulate object storage",
		Aliases:               []string{"objs"},
		Hidden:                false,
		Commands:              subCommands,
		EnableShellCompletion: true,
	}
}
