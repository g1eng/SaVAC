package generator

import (
	"github.com/urfave/cli/v3"
)

func (g *VpsCommandGenerator) GeneratePermissionCommand() *cli.Command {
	return &cli.Command{
		Name:    "permissions",
		Aliases: []string{"perm", "permission"},
		Usage:   "list readonly permissions",
		Flags: append(patternMatchingFlags,
			&cli.BoolFlag{
				Name:  "l",
				Usage: "list additional information",
				Value: false,
			},
		),
		Action: g.ag.GeneratePermissionListAction,
	}
}
