package generator

import "github.com/urfave/cli/v3"

var (
	patternMatchingFlags []cli.Flag = []cli.Flag{
		&cli.BoolFlag{
			Name:    "regex",
			Aliases: []string{"E"},
		},
		&cli.BoolFlag{
			Name:    "search",
			Aliases: []string{"s"},
		},
		&cli.BoolFlag{
			Name:    "tags",
			Aliases: []string{"T"},
		},
	}
)
