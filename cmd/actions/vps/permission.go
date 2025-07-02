package vps_actions

import (
	"context"
	"fmt"
	sakuravps "github.com/g1eng/sakura_vps_client_go"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v3"
)

func (g *VpsActionGenerator) GeneratePermissionListAction(_ context.Context, cmd *cli.Command) error {
	perm, err := g.ApiClient.ListPermissions()
	if err != nil {
		return err
	}
	if cmd.IsSet("regex") {
		permList, err := core.SearchResourceWithRegex(perm, cmd.Args().First())
		if err != nil {
			return err
		}
		perm = permList.([]sakuravps.Permission)
	} else if cmd.IsSet("search") {
		permList, err := core.SearchResourceWithName(perm, cmd.Args().First())
		if err != nil {
			return err
		}
		perm = permList.([]sakuravps.Permission)
	}

	switch g.OutputType {
	case core.OutputTypeYaml:
		return helper.PrintYaml(perm)
	case core.OutputTypeJson:
		return helper.PrintJson(perm)
	default:
		if cmd.IsSet("l") {
			t := helper.NewList()
			t.SetHeader([]string{"name", "code", "category"})
			t.SetAlignment(tablewriter.ALIGN_LEFT)
			t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
			for _, p := range perm {
				t.Append([]string{p.Name, p.Code, p.Category})
			}
			t.Render()
		} else {
			for _, p := range perm {
				fmt.Println(p.Code)
			}
		}
	}
	return nil
}
