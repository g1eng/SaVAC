package generator

import (
	"github.com/g1eng/savac/cmd/helper"
	"github.com/urfave/cli/v3"
)

func (cg *CloudCommandGenerator) GenerateContainerRegistryCommand() *cli.Command {
	subCmd := []*cli.Command{
		{
			Name:      "create",
			Usage:     "create a container registry",
			ArgsUsage: "name",
			Before:    helper.CheckArgsExist,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "description",
					Usage: "description of the container registry",
				},
				&cli.StringFlag{
					Name:  "permission",
					Usage: "access permission for the registry [readwrite|readonly|none]",
					Value: "none",
				},
				&cli.StringFlag{
					Name:  "domain",
					Usage: "alternative domain name of the container registry",
				},
				&cli.StringSliceFlag{
					Name:  "tags",
					Usage: "tags of the container registry",
				},
			},
			Action: cg.ag.GenerateContainerRegistryCreateAction,
		},
		{
			Name:    "delete",
			Usage:   "delete a container registry",
			Aliases: []string{"del"},
			Action:  cg.ag.GenerateContainerRegistryDeleteAction,
		},
		{
			Name:   "list",
			Usage:  "list container registry",
			Action: cg.ag.GenerateContainerRegistryListAction,
		},
		{
			Name:  "user",
			Usage: "manage container registry user",
			Commands: []*cli.Command{
				{
					Name:      "add",
					Usage:     "add an user to the container registry",
					Aliases:   []string{"create"},
					ArgsUsage: "registry-id -u username -p password",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "user",
							Aliases:  []string{"u"},
							Usage:    "user name to add to the container registry",
							Required: true,
						},
						&cli.StringFlag{
							Name:     "password",
							Aliases:  []string{"p"},
							Usage:    "password for the user",
							Required: true,
						},
						&cli.StringFlag{
							Name:     "permission",
							Aliases:  []string{"perm"},
							Usage:    "permission for the user",
							Required: true,
							Value:    "readwrite",
						},
					},
					Before: helper.CheckArgsExist,
					Action: cg.ag.GenerateContainerRegistryUserCreateAction,
				},
				{
					Name:      "list",
					Usage:     "list users for the container registry",
					ArgsUsage: "registry-id",
					Before:    helper.CheckArgsExist,
					Action:    cg.ag.GenerateContainerRegistryUserListAction,
				},
				{
					Name:      "delete",
					Usage:     "remove an user from the container registry",
					ArgsUsage: "registry-id",
					Before:    helper.CheckArgsExist,
					Aliases:   []string{"del"},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "user",
							Aliases:  []string{"u"},
							Usage:    "user name to remove from the container registry",
							Required: true,
						},
					},
					Action: cg.ag.GenerateContainerRegistryUserDeleteAction,
				},
			},
		},
	}
	return &cli.Command{
		Name:                  "container-registry",
		Usage:                 "manipulate container registries",
		Aliases:               []string{"co"},
		Hidden:                false,
		Commands:              subCmd,
		EnableShellCompletion: true,
	}
}
