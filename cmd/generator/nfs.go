package generator

import (
	"github.com/g1eng/savac/cmd/helper"
	"github.com/urfave/cli/v3"
)

func (g *VpsCommandGenerator) GenerateNfsCommand() *cli.Command {
	return &cli.Command{
		Name:  "nfs",
		Usage: "Operate nfs resources",
		Commands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "list NFS informations",
				Flags:  patternMatchingFlags,
				Action: g.ag.GenerateNfsListAction,
			},
			{
				Name:      "info",
				Usage:     "show NFS information with a service code",
				Aliases:   []string{"get", "show"},
				ArgsUsage: "service_code",
				Action:    g.ag.GenerateNfsInfoAction,
			},
			{
				Name:      "interfaces",
				Usage:     "show NFS interface informations with a service code",
				Aliases:   []string{"if", "interface"},
				Flags:     patternMatchingFlags,
				ArgsUsage: "service_code",
				Action:    g.ag.GenerateNfsInterfaceAction,
			},
			{
				Name:      "start",
				Usage:     "start NFS server",
				ArgsUsage: "<expr>",
				Flags:     patternMatchingFlags,
				Before:    helper.CheckArgsExist,
				Action:    g.ag.GenerateNfsStartAction,
			},
			{
				Name:      "stop",
				Usage:     "stop NFS server",
				ArgsUsage: "<expr>",
				Flags: append(patternMatchingFlags,
					&cli.BoolFlag{
						Name:  "force",
						Usage: "force stop NFS server",
						Value: false,
					},
				),
				Before: helper.CheckArgsExist,
				Action: g.ag.GenerateNfsStopAction,
			},
			{
				Name:      "reboot",
				Usage:     "reboot NFS server",
				ArgsUsage: "<expr>",
				Flags:     patternMatchingFlags,
				Before:    helper.CheckArgsExist,
				Action:    g.ag.GenerateNfsRebootAction,
			},
			{
				Name:      "connect",
				Usage:     "connect NFS to a switch",
				ArgsUsage: "nfs_id switch_id",
				Before:    helper.CheckArgsExist,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "disconnect",
						Aliases: []string{"d"},
					},
				},
				Action: g.ag.GenerateNfsConnectAction,
			},
		},
	}
}
