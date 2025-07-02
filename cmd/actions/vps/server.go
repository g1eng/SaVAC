package vps_actions

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/urfave/cli/v3"
)

func (g *VpsActionGenerator) GenerateServerListAction(_ context.Context, cmd *cli.Command) error {
	servers, err := g.ApiClient.GetAllServers()
	if err != nil {
		return err
	}

	if cmd.Bool("l") && g.OutputType == core.OutputTypeText {
		t := helper.NewList()
		if len(servers) == 0 {
			fmt.Println("(no servers)")
		}
		if !g.ApiClient.NoHeader {
			t.SetHeader([]string{"id", "name", "status", "spec", "zone", "ipv4", "ipv6", "zone"})
		}

		for _, d := range servers {
			var mem string
			if d.MemoryMebibytes < 1000 {
				mem = "512M"
			} else {
				mem = fmt.Sprintf("%dc/%dGB", d.CpuCores, (d.MemoryMebibytes)/1024)
			}
			t.Append([]string{
				strconv.Itoa(int(d.Id)),
				d.Name,
				d.PowerStatus,
				mem,
				d.Zone.Code,
				d.Ipv4.Address,
				d.Ipv6.GetAddress(),
				d.GetZone().Code,
			})
		}
		t.Render()
		return nil
	}

	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(servers)
	case core.OutputTypeYaml:
		return helper.PrintYaml(servers)
	default:
		for _, d := range servers {
			switch d.PowerStatus {
			case "power_on":
				fmt.Fprintf(os.Stderr, "\x1b[32m")
				fmt.Fprintf(os.Stdout, "%d\n", d.Id)
				fmt.Fprintf(os.Stderr, "\x1b[0m")
			case "power_off":
				fmt.Fprintf(os.Stdout, "%d\n", d.Id)
			default:
				fmt.Fprintf(os.Stderr, "\x1b[33m")
				fmt.Fprintf(os.Stdout, "%d\n", d.Id)
				fmt.Fprintf(os.Stderr, "\x1b[0m")
			}
		}
		return nil
	}
}

func (g *VpsActionGenerator) GenerateServerInfoAction(_ context.Context, cmd *cli.Command) error {
	var (
		svList []*sakuravps.Server
		errs   []error
	)
	s := cmd.Args().First()
	_, err := strconv.Atoi(s)
	if err != nil {
		svList, err = g.ApiClient.GetAllServers()
		if err != nil {
			return fmt.Errorf("failed to get switch list: %w", err)
		}
		if cmd.IsSet("regex") {
			switches, err := core.SearchResourceWithRegex(svList, cmd.Args().First())
			if err != nil {
				return err
			}
			svList = switches.([]*sakuravps.Server)
		} else if cmd.IsSet("search") {
			switches, err := core.SearchResourceWithName(svList, cmd.Args().First())
			if err != nil {
				return err
			}
			svList = switches.([]*sakuravps.Server)
		} else {
			sv, err := core.MatchResourceWithName(svList, s)
			if err != nil {
				return err
			}
			svList = []*sakuravps.Server{sv}
		}
		for _, sv := range svList {
			powerRes, _, err := g.ApiClient.RawClient.ServerAPI.GetServerPowerStatus(context.Background(), sv.Id).Execute()
			if err != nil {
				errs = append(errs, err)
			} else {
				sv.PowerStatus = powerRes.Status
			}
			switch g.OutputType {
			case core.OutputTypeJson:
				helper.PrintJson(sv) // nolint
			case core.OutputTypeYaml:
				helper.PrintYaml(sv) // nolint
			case core.OutputTypeText:
				helper.PrintTableForServerInfo(sv)
			}
		}
	} else {
		for _, idString := range cmd.Args().Slice() {
			id, err := strconv.Atoi(idString)
			if err != nil {
				errs = append(errs, err)
			} else {
				sv, err := g.ApiClient.GetServerById(int32(id))
				if err != nil {
					errs = append(errs, err)
				} else {
					powerRes, _, err := g.ApiClient.RawClient.ServerAPI.GetServerPowerStatus(context.Background(), sv.Id).Execute()
					if err != nil {
						errs = append(errs, err)
					} else {
						sv.PowerStatus = powerRes.Status
					}
					switch g.OutputType {
					case core.OutputTypeJson:
						helper.PrintJson(sv) // nolint
					case core.OutputTypeYaml:
						helper.PrintYaml(sv) // nolint
					case core.OutputTypeText:
						helper.PrintTableForServerInfo(sv)
					}
				}
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerHostnameAction(_ context.Context, cmd *cli.Command) error {
	s := cmd.Args().First()
	id, err := strconv.Atoi(s)
	if err != nil {
		servers, err := g.ApiClient.GetAllServers()
		if err != nil {
			return err
		}
		for _, sv := range servers {
			if sv.Name == s {
				id = int(sv.Id)
				return g.ApiClient.SetHostName(int32(id), cmd.Args().Slice()[1])
			}
		}
		return fmt.Errorf("no such server: %s", s)
	} else if cmd.Args().Len() == 2 {
		return g.ApiClient.SetHostName(int32(id), cmd.Args().Slice()[1])
	} else {
		h, err := g.ApiClient.GetHostname(int32(id))
		if err != nil {
			return err
		}
		fmt.Println(h)
		return nil
	}
}

func (g *VpsActionGenerator) GenerateServerDescriptionAction(_ context.Context, cmd *cli.Command) error {
	s := cmd.Args().First()
	id, err := strconv.Atoi(s)
	if err != nil {
		id = -1
		servers, err := g.ApiClient.GetAllServers()
		if err != nil {
			return err
		}
		for _, sv := range servers {
			if sv.Name == s {
				id = int(sv.Id)
				break
			}
		}
		if id == -1 {
			return fmt.Errorf("first argument must be server id")
		}
	}
	if cmd.Args().Len() < 2 {
		desc, err := g.ApiClient.GetServerDescriptionById(int32(id))
		if err != nil {
			return err
		}
		switch g.OutputType {
		case core.OutputTypeJson:
			return helper.PrintJson(desc)
		case core.OutputTypeYaml:
			return helper.PrintYaml(desc)
		case core.OutputTypeText:
			for _, d := range desc {
				fmt.Println(d)
			}
		}
		return nil
	} else {
		return g.ApiClient.SetServerDescription(int32(id), cmd.Args().Slice()[1])
	}
}

func (g *VpsActionGenerator) GenerateServerTagAction(_ context.Context, cmd *cli.Command) error {
	s := cmd.Args().First()
	id, err := strconv.Atoi(s)
	if err != nil {
		id = -1
		servers, err := g.ApiClient.GetAllServers()
		if err != nil {
			return err
		}
		for _, sv := range servers {
			if sv.Name == s {
				id = int(sv.Id)
				break
			}
		}
		if id == -1 {
			return fmt.Errorf("server not found for the pattrern %s", s)
		}
	}
	if cmd.Args().Len() < 2 {
		tags, err := g.ApiClient.ListServerTags(int32(id))
		if err != nil {
			return err
		}
		for _, tag := range tags {
			println(tag)
		}
		return nil
	} else if cmd.Args().Len() == 2 {
		t, err := g.ApiClient.GetServerTag(int32(id), cmd.Args().Slice()[1])
		if err != nil {
			return err
		} else if t == nil {
			log.Fatalln(fmt.Errorf("tag not found: %s", cmd.Args().Slice()[1]))
		}
		fmt.Println(*t)
		return nil
	} else {
		return g.ApiClient.AddServerTag(int32(id), cmd.Args().Slice()[1], cmd.Args().Slice()[2])
	}
}

func (g *VpsActionGenerator) GenerateServerInterfaceAction(_ context.Context, cmd *cli.Command) error {
	var s string
	if !cmd.Args().Present() {
		s = "."
	} else {
		s = cmd.Args().First()
	}
	id, err := strconv.Atoi(s)
	if err != nil {
		if g.ApiClient.Debug {
			println("no interfaces id parsing")
		}
		if result, err := g.ApiClient.GetServerInterfacesWithPattern(s); err != nil {
			return err
		} else {
			switches, err := g.ApiClient.GetSwitchList()
			if err != nil {
				return err
			}
			if len(result) == 0 {
				return fmt.Errorf("[savac] no server found with the pattern %s", s)
			}
			t := helper.NewList()
			t.Append([]string{"server", "NIC id", "name", "address", "switchId", "switch"})
			switch g.OutputType {
			case core.OutputTypeJson:
				data := map[string][]sakuravps.ServerInterface{}
				for sv, interfaces := range result {
					data[sv.Name] = interfaces
				}
				return helper.PrintJson(data)
			case core.OutputTypeYaml:
				data := map[string][]sakuravps.ServerInterface{}
				for sv, interfaces := range result {
					data[sv.Name] = interfaces
				}
				return helper.PrintYaml(data)
			case core.OutputTypeText:
				for sv, interfaces := range result {
					svName := sv.Name
					for _, inf := range interfaces {
						s := []byte(inf.Mac)
						for i, c := range inf.Mac {
							switch c {
							case 'A':
								s[i] = 'a'
							case 'B':
								s[i] = 'b'
							case 'C':
								s[i] = 'c'
							case 'D':
								s[i] = 'd'
							case 'E':
								s[i] = 'e'
							case 'F':
								s[i] = 'f'
							}
						}
						var switchId, switchName string
						if inf.ConnectableToGlobalNetwork && inf.GetSwitchId() == 0 {
							switchId = "shared"
							switchName = "global"
						} else if inf.GetSwitchId() == 0 {
							switchId = "-"
							switchName = "-"
						} else {
							switchId = fmt.Sprintf("%d", inf.GetSwitchId())
							switchName = fmt.Sprintf("%d", inf.SwitchId.Get())
							for _, sw := range switches {
								if sw.Id == inf.GetSwitchId() {
									switchName = sw.Name
								}
							}
						}
						t.Append([]string{svName, fmt.Sprintf("%d", inf.Id), inf.DisplayName, string(s), switchId, switchName})
					}
				}
				t.Render()
			}
			return nil
		}
	} else {
		if g.ApiClient.Debug {
			println("interfaces id parsing")
		}
		res, err := g.ApiClient.GetServerInterfaces(int32(id))
		if err != nil {
			return err
		}
		switches, err := g.ApiClient.GetSwitchList()
		if err != nil {
			return err
		}
		switch g.OutputType {
		case core.OutputTypeJson:
			return helper.PrintJson(res)
		case core.OutputTypeYaml:
			return helper.PrintYaml(res)
		case core.OutputTypeText:
			helper.PrintTableForServerInterfaces(fmt.Sprintf("%d", id), res, switches)
		}
		return nil
	}
}

func (g *VpsActionGenerator) GenerateServerConnectAction(_ context.Context, cmd *cli.Command) error {
	nic := cmd.Args().First()

	nicId, err := strconv.Atoi(nic)
	if err != nil {
		return fmt.Errorf("failed to parse NIC id ")
	}

	if cmd.Args().Len() == 1 && cmd.Bool("disconnect") {
		return g.ApiClient.SetServerInterfaceConnection(int32(nicId), 0)
	} else if cmd.Args().Len() < 2 {
		return fmt.Errorf("switch id is not specified")
	}
	switchExpr := cmd.Args().Slice()[1]
	if switchId, err := strconv.Atoi(switchExpr); err != nil {
		if switchExpr == "global" || switchExpr == "shared" {
			return g.ApiClient.SetServerInterfaceConnection(int32(nicId), 1)
		}
		return fmt.Errorf("failed to parse switch id ")
	} else {
		if g.ApiClient.Debug {
			println("interfaces id parsing")
		}
		return g.ApiClient.SetServerInterfaceConnection(int32(nicId), int32(switchId))
	}
}

func (g *VpsActionGenerator) GenerateServerPtrAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() < 1 {
		return fmt.Errorf("at least one arguments needed")
	}
	s := cmd.Args().Get(0)
	id, err := strconv.Atoi(s)
	act := func(id int32) error {
		if cmd.Bool("6") {
			if cmd.Args().Len() != 2 {
				s, err := g.ApiClient.GetIpv6PtrWithServerId(id)
				if err != nil {
					return err
				}
				fmt.Println(s)
				return nil
			}
			p := cmd.Args().Get(1)
			return g.ApiClient.SetIpv6PtrWithServerId(id, p)
		} else {
			if cmd.Args().Len() != 2 {
				s, err := g.ApiClient.GetIpv4PtrWithServerId(id)
				if err != nil {
					return err
				}
				fmt.Println(s)
				return nil
			}
			p := cmd.Args().Get(1)
			fmt.Fprintf(os.Stderr, "[testsuite] p: %s\n", p)
			return g.ApiClient.SetIpv4PtrWithServerId(id, p)
		}
	}
	if err != nil {
		ids, err := g.ApiClient.GetServerIdsByNamePattern(s)
		if err != nil {
			return fmt.Errorf("no server found")
		}
		for i, id := range ids {
			if i != 0 {
				time.Sleep(500)
			}
			if err = act(id); err != nil {
				return err
			}
		}
		return nil
	} else {
		return act(int32(id))
	}
}

func (g *VpsActionGenerator) GenerateServerStartAction(_ context.Context, cmd *cli.Command) error {
	s := cmd.Args().First()
	var (
		targetServers []*sakuravps.Server
		errs          []error
	)
	startForEach := func() error {
		for _, sv := range targetServers {
			err := g.ApiClient.StartServerWithId(sv.Id)
			if err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			return fmt.Errorf("%v", errs)
		}
		return nil
	}
	id, err := strconv.Atoi(s) //nolint

	if err != nil {
		if g.ApiClient.Debug {
			fmt.Fprintf(os.Stderr, "pattrn %s\n", s)
		}
		allServer, err := g.ApiClient.GetAllServers()
		if err != nil {
			return err
		}

		if cmd.Bool("regex") {
			res, err := core.SearchResourceWithRegex(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]*sakuravps.Server)
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
			targetServers = res.([]*sakuravps.Server)
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
				targetServers = append(targetServers, &sakuravps.Server{Id: int32(id)})
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

func (g *VpsActionGenerator) GenerateServerStopAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Bool("force") {
		g.ApiClient.Forced = true
	}
	var (
		targetServers []*sakuravps.Server
		errs          []error
	)
	stopForEach := func() error {
		for _, sv := range targetServers {
			err := g.ApiClient.StopServerWithId(sv.Id)
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
		allServer, err := g.ApiClient.GetAllServers()
		if err != nil {
			return err
		}

		if cmd.Bool("regex") {
			res, err := core.SearchResourceWithRegex(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]*sakuravps.Server)
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
			targetServers = res.([]*sakuravps.Server)
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
				targetServers = append(targetServers, &sakuravps.Server{Id: int32(id)})
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

func (g *VpsActionGenerator) GenerateServerRebootAction(_ context.Context, cmd *cli.Command) error {
	var (
		targetServers []*sakuravps.Server
		errs          []error
	)
	rebootForEach := func() error {
		for _, sv := range targetServers {
			err := g.ApiClient.ForceRebootServerWithId(sv.Id)
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
		allServer, err := g.ApiClient.GetAllServers()
		if err != nil {
			return err
		}

		if cmd.Bool("regex") {
			res, err := core.SearchResourceWithRegex(allServer, s)
			if err != nil {
				return err
			}
			targetServers = res.([]*sakuravps.Server)
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
			targetServers = res.([]*sakuravps.Server)
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
				targetServers = append(targetServers, &sakuravps.Server{Id: int32(id)})
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
