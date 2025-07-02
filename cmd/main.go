package main

import (
	"context"

	"github.com/g1eng/savac/pkg/core"

	// "bufio"
	// "bytes"
	"fmt"
	"os"

	// "github.com/g1eng/savac/pkg/client"
	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/vps"
)

func throw(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(1)
}

func main() {
	conf := sakuravps.NewConfiguration()
	conf.UserAgent = core.USER_AGENT
	cli := vps.NewClient(sakuravps.NewAPIClient(conf))

	command := Generate(cli)

	if err := command.Run(context.Background(), os.Args); err != nil {
		throw(err)
	}
	os.Exit(0)
}
