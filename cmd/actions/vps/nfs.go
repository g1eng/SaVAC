package vps_actions

import (
	"context"
	"fmt"
	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/urfave/cli/v3"
	"os"
	"strconv"
	"strings"
)

func (g *VpsActionGenerator) GenerateNfsListAction(_ context.Context, cmd *cli.Command) error {
	res, err := g.ApiClient.GetAllNFS()
	if err != nil {
		return err
	}
	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(res)
	case core.OutputTypeYaml:
		return helper.PrintYaml(res)
	default:
		t := helper.NewList()
		if len(res) == 0 {
			fmt.Println("(no servers)")
		}
		if !g.ApiClient.NoHeader {
			t.SetHeader([]string{"id", "name", "status", "spec", "zone", "ipv4"})
		}
		for _, n := range res {
			t.Append([]string{
				fmt.Sprintf("%d", int(n.Id)),
				n.Name,
				n.PowerStatus,
				func() string {
					var s []string
					for _, val := range n.Storage {
						s = append(s, fmt.Sprintf("%dGiB", val.SizeGibibytes))
					}
					return strings.Join(s, ",")
				}(),
				n.Zone.Code,
				n.Ipv4.Address,
			})
		}
		t.Render()
	}
	return nil
}

func (g *VpsActionGenerator) GenerateNfsInfoAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return fmt.Errorf("no argument")
	}
	code, err := strconv.Atoi(cmd.Args().First())
	showNfsInfo := func() error {
		res, err := g.ApiClient.GetNFSById(int32(code))
		if err != nil {
			return err
		}
		storageInfo, err := g.ApiClient.GetNfsStorageInfo(res.Id)
		var capacity, usage, percentage string
		if err != nil {
			capacity, usage, percentage = "unknown", "unknown", "unknown"
		} else {
			capacity, usage, percentage =
				fmt.Sprintf("%d GiB", storageInfo.GetCapacityKib()/1000000),
				fmt.Sprintf("%d GiB", storageInfo.GetUsageKib()/1000000),
				fmt.Sprintf("%d%s", storageInfo.GetUsagePercentage(), "%")
		}

		switch g.OutputType {
		case core.OutputTypeJson:
			return helper.PrintJson(res)
		case core.OutputTypeYaml:
			return helper.PrintYaml(res)
		default:
			t := helper.NewList()
			t.SetHeader([]string{"id", "name", "zone", "status", "ipv4", "capacity", "usage", "percentage"})
			code := fmt.Sprintf("%d", res.Id)
			t.Append([]string{code, res.Name, res.Zone.Name, res.PowerStatus, res.Ipv4.Address, capacity, usage, percentage})
			t.Render()
		}
		return nil
	}
	if err != nil {
		res, err := g.ApiClient.GetNFSByName(cmd.Args().First())
		if err != nil {
			return err
		}
		code = int(res.Id)
		err = showNfsInfo()
		if err != nil {
			return err
		}
		return nil
	}

	return showNfsInfo()
}

func (g *VpsActionGenerator) GenerateNfsInterfaceAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return fmt.Errorf("no argument")
	}
	code, err := strconv.Atoi(cmd.Args().First())
	if err != nil {
		return err
	}
	res, err := g.ApiClient.GetNFSInterface(int32(code))
	if err != nil {
		return err
	}
	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(res)
	case core.OutputTypeYaml:
		return helper.PrintYaml(res)
	default:
		t := helper.NewList()
		if !g.ApiClient.NoHeader {
			t.SetHeader([]string{"id", "hwaddr", "switch-id"})
		}
		swId := "0"
		if _, ok := res.GetSwitchIdOk(); ok {
			swId = fmt.Sprintf("%d", res.GetSwitchId())
		}
		t.Append([]string{
			fmt.Sprintf("%d", code),
			res.Mac,
			swId,

			func() string {
				if res.SwitchId.Get() == nil {
					return "-"
				} else {
					return fmt.Sprintf("%d", *res.SwitchId.Get())
				}
			}(),
		})
		t.Render()
	}
	return nil
}

func (g *VpsActionGenerator) GenerateNfsConnectAction(_ context.Context, cmd *cli.Command) error {
	var (
		switchId int
		err      error
	)
	if cmd.Args().Len() < 1 {
		return fmt.Errorf("no argument")
	} else if cmd.Args().Len() == 1 && cmd.Bool("disconnect") {
		switchId = 0
	} else if cmd.Args().Len() < 2 {
		return fmt.Errorf("two arguments required")
	} else {
		switchId, err = strconv.Atoi(cmd.Args().Get(1))
		if err != nil {
			return err
		}
	}
	nfsId, err := strconv.Atoi(cmd.Args().First())
	if err != nil {
		res, err := g.ApiClient.GetNFSByName(cmd.Args().First())
		if err != nil {
			return err
		}
		err = g.ApiClient.PutNFSConnection(res.Id, int32(switchId))
		if err != nil {
			return err
		}
		return nil
	}
	//nfsInterface, err := g.ApiClient.GetNFSInterface(int32(nfsId))
	//if err != nil {
	//	return err
	//}
	return g.ApiClient.PutNFSConnection(int32(nfsId), int32(switchId))
}

func (g *VpsActionGenerator) GenerateNfsStartAction(ctx context.Context, cmd *cli.Command) error {
	var (
		targetServers []sakuravps.NfsServer
		errs          []error
	)
	startForEach := func() error {
		for _, sv := range targetServers {
			err := g.ApiClient.StartNfs(sv.Id)
			if err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			return fmt.Errorf("%v", errs)
		}
		return nil
	}
	s := cmd.Args().First()
	id, err := strconv.Atoi(s) //nolint
	if err != nil {
		if g.ApiClient.Debug {
			fmt.Fprintf(os.Stderr, "pattrn %s\n", s)
		}
		allServer, err := g.ApiClient.GetAllNFS()
		if err != nil {
			return err
		}

		if cmd.Bool("regex") {
			res, err := core.SearchResourceWithRegex(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]sakuravps.NfsServer)
			err = startForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		} else if cmd.Bool("search") {
			res, err := core.SearchResourceWithName(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]sakuravps.NfsServer)
			err = startForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		} else {
			for _, name := range cmd.Args().Slice() {
				sv, err := core.MatchResourceWithName(allServer, name)
				if err != nil {
					errs = append(errs, err)
				} else {
					targetServers = append(targetServers, sv)
				}
			}
			err = startForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		}
	} else {
		ids := cmd.Args().Slice()
		for _, i := range ids {
			id, err = strconv.Atoi(i)
			if err != nil {
				errs = append(errs, err)
			} else {
				targetServers = append(targetServers, sakuravps.NfsServer{Id: int32(id)})
			}
		}
		err = startForEach()
		if err != nil {
			errs = append(errs, err)
		}
		if len(errs) > 0 {
			return fmt.Errorf("%v", errs)
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateNfsStopAction(ctx context.Context, cmd *cli.Command) error {
	var (
		targetServers []sakuravps.NfsServer
		errs          []error
		err           error
	)
	isForced := cmd.Bool("force")
	stopForEach := func() error {
		for _, sv := range targetServers {
			if isForced {
				err = g.ApiClient.ShutdownNfs(sv.Id, isForced)
			} else {
				err = g.ApiClient.ShutdownNfs(sv.Id)
			}
			if err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			return fmt.Errorf("%v", errs)
		}
		return nil
	}
	s := cmd.Args().First()
	id, err := strconv.Atoi(s) // nolint
	if err != nil {
		if g.ApiClient.Debug {
			fmt.Fprintf(os.Stderr, "pattrn %s\n", s)
		}
		allServer, err := g.ApiClient.GetAllNFS()
		if err != nil {
			return err
		}

		if cmd.Bool("regex") {
			res, err := core.SearchResourceWithRegex(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]sakuravps.NfsServer)
			err = stopForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		} else if cmd.Bool("search") {
			res, err := core.SearchResourceWithName(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]sakuravps.NfsServer)
			err = stopForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		} else {
			for _, name := range cmd.Args().Slice() {
				sv, err := core.MatchResourceWithName(allServer, name)
				if err != nil {
					errs = append(errs, err)
				} else {
					targetServers = append(targetServers, sv)
				}
			}
			err = stopForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		}
	} else {
		ids := cmd.Args().Slice()
		for _, i := range ids {
			id, err = strconv.Atoi(i)
			if err != nil {
				errs = append(errs, err)
			} else {
				targetServers = append(targetServers, sakuravps.NfsServer{Id: int32(id)})
			}
		}
		err = stopForEach()
		if err != nil {
			errs = append(errs, err)
		}
		if len(errs) > 0 {
			return fmt.Errorf("%v", errs)
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateNfsRebootAction(_ context.Context, cmd *cli.Command) error {
	var (
		targetServers []sakuravps.NfsServer
		errs          []error
	)
	rebootForEach := func() error {
		for _, sv := range targetServers {
			err := g.ApiClient.RebootNfs(sv.Id)
			if err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			return fmt.Errorf("%v", errs)
		}
		return nil
	}
	s := cmd.Args().First()
	id, err := strconv.Atoi(s) // nolint
	if err != nil {
		if g.ApiClient.Debug {
			fmt.Fprintf(os.Stderr, "pattrn %s\n", s)
		}
		allServer, err := g.ApiClient.GetAllNFS()
		if err != nil {
			return err
		}

		if cmd.Bool("regex") {
			res, err := core.SearchResourceWithRegex(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]sakuravps.NfsServer)
			err = rebootForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		} else if cmd.Bool("search") {
			res, err := core.SearchResourceWithName(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]sakuravps.NfsServer)
			err = rebootForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		} else {
			for _, name := range cmd.Args().Slice() {
				sv, err := core.MatchResourceWithName(allServer, name)
				if err != nil {
					errs = append(errs, err)
				} else {
					targetServers = append(targetServers, sv)
				}
			}
			err = rebootForEach()
			if err != nil {
				errs = append(errs, err)
			}
			if len(errs) > 0 {
				return fmt.Errorf("%v", errs)
			}
		}
	} else {
		ids := cmd.Args().Slice()
		for _, i := range ids {
			id, err = strconv.Atoi(i)
			if err != nil {
				errs = append(errs, err)
			} else {
				targetServers = append(targetServers, sakuravps.NfsServer{Id: int32(id)})
			}
		}
		err = rebootForEach()
		if err != nil {
			errs = append(errs, err)
		}
		if len(errs) > 0 {
			return fmt.Errorf("%v", errs)
		}
	}
	return nil
}
