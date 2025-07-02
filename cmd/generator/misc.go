package generator

import (
	"github.com/urfave/cli/v3"
)

func (g *VpsCommandGenerator) GenerateZoneCommand() *cli.Command {
	return &cli.Command{
		Name:    "zone",
		Aliases: []string{"zones"},
		Usage:   "List zones",
		Action:  g.ag.GenerateZoneAction,
	}
}

func (g *VpsCommandGenerator) GenerateDiscCommand() *cli.Command {
	return &cli.Command{
		Name:    "disc",
		Aliases: []string{"cdrom"},
		Usage:   "List CD-ROMs inserted in VM",
		Action:  g.ag.GenerateDiscAction,
	}
}
