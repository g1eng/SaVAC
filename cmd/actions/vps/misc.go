package vps_actions

import (
	"context"
	"regexp"
	"strconv"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/urfave/cli/v3"
)

func (g *VpsActionGenerator) GenerateDiscAction(_ context.Context, _ *cli.Command) error {
	if res, err := g.ApiClient.ListCDROMs(); err != nil {
		return err
	} else {
		switch g.OutputType {
		case core.OutputTypeJson:
			return helper.PrintJson(res)
		case core.OutputTypeYaml:
			return helper.PrintYaml(res)
		default:
			t := helper.NewList()
			t.SetHeader([]string{"id", "name", "license required", "description"})
			pat := regexp.MustCompilePOSIX("[、。\n]")
			for _, d := range res {
				licenseMsg := func() string {
					if d.LicenseRequired {
						return "yes"
					} else {
						return "no"
					}
				}()
				t.Append([]string{strconv.Itoa(int(d.Id)), d.Name, licenseMsg, pat.ReplaceAllString(d.Description, ", ")})
			}
			t.Render()
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateZoneAction(ctx context.Context, cmd *cli.Command) error {
	if zones, err := g.ApiClient.GetAllZone(); err != nil {
		return err
	} else {
		switch g.OutputType {
		case core.OutputTypeJson:
			return helper.PrintJson(zones)
		case core.OutputTypeYaml:
			return helper.PrintYaml(zones)
		default:
			t := helper.NewList()
			if !g.ApiClient.NoHeader {
				t.SetHeader([]string{"code", "name"})
			}
			for _, z := range zones {
				t.Append([]string{z.Code, z.Name})
			}
			t.Render()
		}
	}
	return nil
}
