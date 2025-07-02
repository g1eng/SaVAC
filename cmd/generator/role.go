package generator

import (
	"github.com/g1eng/savac/cmd/helper"

	"github.com/urfave/cli/v3"
)

func (g *VpsCommandGenerator) GenerateRoleSubcommand() *cli.Command {
	return &cli.Command{
		Name:    "role",
		Aliases: []string{"r"},
		Usage:   "Operate VPS IAM role",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list roles",
				Aliases: []string{"ls"},
				Flags: append(patternMatchingFlags, &cli.BoolFlag{
					Name:  "l",
					Usage: "list additional information",
					Value: false,
				}),
				Action: g.ag.GenerateRoleListAction,
			},
			{
				Name:    "create",
				Aliases: []string{"cr"},
				Usage:   "create a new role",
				Hidden:  true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "server",
						Aliases: []string{"sv"},
						Usage:   "server index or regex for resource filtering",
					},
					//&cli.StringSliceFlag{
					//	Name:    "tag",
					//	Aliases: []string{"T"},
					//	Hidden:  true,
					//	Usage:   "(not implemented) tag for resource filtering",
					//},
					&cli.StringFlag{
						Name:    "switch",
						Aliases: []string{"sw"},
						Usage:   "server index or regex for resource filtering",
					},
					&cli.StringFlag{
						Name:  "nfs",
						Usage: "NFS server index or regex for resource filtering",
					},
					&cli.StringFlag{
						Name:    "permissions",
						Aliases: []string{"p", "perm"},
						Usage:   "permission codes or regex for permission filtering",
					},
				},
				Before:    helper.CheckArgsExist,
				ArgsUsage: "<name>",
				Action:    g.ag.GenerateRoleCreateAction,
			},
			{
				Name:      "get",
				Aliases:   []string{"read"},
				Flags:     patternMatchingFlags,
				Usage:     "read the role information",
				Hidden:    true,
				ArgsUsage: "<role-id>",
				Action:    g.ag.GenerateRoleReadAction,
			},
			{
				Name:      "delete",
				Aliases:   []string{"del"},
				Usage:     "delete the role",
				Hidden:    true,
				ArgsUsage: "<role-id>",
				Before:    helper.CheckArgsExist,
				Action:    g.ag.GenerateRoleDeleteAction,
			},
			{
				Name:      "update",
				Usage:     "update the role",
				ArgsUsage: "<role-id>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "role name",
					},
					&cli.Int32SliceFlag{
						Name:    "server",
						Aliases: []string{"sv"},
						Usage:   "server indices for resource filtering",
					},
					&cli.Int32SliceFlag{
						Name:    "switch",
						Aliases: []string{"sw"},
						Usage:   "server indices for resource filtering",
					},
					&cli.Int32SliceFlag{
						Name:  "nfs",
						Usage: "NFS server indices for resource filtering",
					},
					&cli.StringSliceFlag{
						Name:    "permissions",
						Aliases: []string{"p", "perm"},
						Usage:   "permission codes for permissions filtering",
					},
				},
				Before: helper.CheckArgsExist,
				Action: g.ag.GenerateRoleUpdateAction,
			},
		},
	}
}
