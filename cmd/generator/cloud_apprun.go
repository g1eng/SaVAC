package generator

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func getAppRunApplicationDeployFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "username  for the private container registry",
		},
		&cli.StringFlag{
			Name:    "password",
			Usage:   "password for the private container registry",
			Aliases: []string{"P"},
		},
		&cli.Int32Flag{
			Name:     "port",
			Aliases:  []string{"p"},
			Required: true,
		},
		&cli.Float64Flag{
			Name:     "cpu",
			Usage:    "max cpu cores",
			Value:    0.1,
			Required: true,
		},
		&cli.Int32Flag{
			Name:     "memory",
			Usage:    "max memories in MB",
			Value:    256,
			Aliases:  []string{"m"},
			Required: true,
		},
		&cli.Int32Flag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Usage:   "timeout seconds for healthcheck",
			Value:   60,
		},
		&cli.StringFlag{
			Name:  "path",
			Usage: "http healthcheck path",
			Value: "/health",
		},
		&cli.IntFlag{
			Name:    "health-port",
			Aliases: []string{"H"},
			Usage:   "http healthcheck port",
			Value:   80,
		},
	}
}

func (cg *CloudCommandGenerator) GenerateAppRunCommand() *cli.Command {
	apprunSubCommends := []*cli.Command{
		{
			Name:    "ps",
			Aliases: []string{"ls", "list"},
			Flags: append([]cli.Flag{
				&cli.StringFlag{
					Name:    "versions",
					Aliases: []string{"ver", "version"},
				},
			}, patternMatchingFlags...),
			Usage:                 "list deployed apps",
			Hidden:                false,
			Action:                cg.ag.GenerateAppRunApplicationListAction,
			EnableShellCompletion: true,
		},
		{
			Name:    "get",
			Aliases: []string{"read"},
			Flags: append([]cli.Flag{
				&cli.StringFlag{
					Name:    "version",
					Aliases: []string{"ver"},
					Usage:   "read specific app deployment version",
				},
			}, patternMatchingFlags...),
			Usage: "get application deployment information",
			Before: func(ctx context.Context, command *cli.Command) (context.Context, error) {
				if command.Args().Len() == 0 {
					return ctx, fmt.Errorf("no application id specified")
				}
				return ctx, nil
			},
			Action: cg.ag.GenerateAppRunApplicationGetAction,
		},
		{
			Name: "run",
			Before: func(ctx context.Context, command *cli.Command) (context.Context, error) {
				if command.Args().Len() == 0 {
					return ctx, fmt.Errorf("no application image specified")
				}
				return ctx, nil
			},
			Flags:  getAppRunApplicationDeployFlags(),
			Usage:  "create app run",
			Action: cg.ag.GenerateAppRunApplicationCreateAction,
		},
		{
			Name: "update",
			Before: func(ctx context.Context, command *cli.Command) (context.Context, error) {
				if command.Args().Len() == 0 {
					return ctx, fmt.Errorf("no application id specified")
				}
				return ctx, nil
			},
			Flags: append(getAppRunApplicationDeployFlags(), &cli.StringFlag{
				Name:    "app-id",
				Aliases: []string{"id"},
				Usage:   "target app id or name",
			}),
			Usage:  "update the default app version",
			Action: cg.ag.GenerateAppRunApplicationCreateVersionAction,
		},
		{
			Name:    "delete",
			Aliases: []string{"del"},
			Flags: append([]cli.Flag{
				&cli.StringFlag{
					Name:    "version",
					Aliases: []string{"ver"},
					Usage:   "delete specific app deployment version",
				},
			}, patternMatchingFlags...),
			Usage: "delete an app",
			Before: func(ctx context.Context, command *cli.Command) (context.Context, error) {
				if command.Args().Len() == 0 {
					return ctx, fmt.Errorf("no application id specified")
				}
				return ctx, nil
			},
			Action: cg.ag.GenerateAppRunApplicationDeleteAction,
		},
	}
	return &cli.Command{
		Name:                  "apprun",
		Aliases:               []string{"app"},
		Usage:                 "Operate AppRun instances",
		Hidden:                false,
		Commands:              apprunSubCommends,
		EnableShellCompletion: true,
	}
}
