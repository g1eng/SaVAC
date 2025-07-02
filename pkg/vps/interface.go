package vps

import (
	"context"
	"fmt"
	"os"

	"github.com/g1eng/savac/pkg/core"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
)

func (s *SavaClient) GetAllServerInterface() (map[int32][]sakuravps.ServerInterface, error) {
	servers, err := s.GetAllServers()
	if err != nil {
		return nil, err
	}
	interfaces := map[int32][]sakuravps.ServerInterface{}
	for _, sv := range servers {
		sif, err := s.GetServerInterfaces(sv.Id)
		if err != nil {
			return nil, err
		}
		interfaces[sv.Id] = sif
	}
	return interfaces, nil
}

func (s *SavaClient) GetServerInterfaces(serverId int32) ([]sakuravps.ServerInterface, error) {
	var result []sakuravps.ServerInterface

	res, resp, err := s.RawClient.ServerAPI.GetServerInterfaceList(context.Background(), serverId).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(resp, err)
	}
	if res == nil {
		return nil, fmt.Errorf("[savac] no DummyInterfaces found")
	}
	result = res.Results
	if len(result) == 0 {
		return nil, fmt.Errorf("[savac] no DummyInterfaces found")
	}
	if s.Debug {
		for _, r := range result {
			fmt.Fprintf(os.Stderr, "%v\n", r)
		}
	}

	return result, nil
}

func (s *SavaClient) GetServerInterfacesWithPattern(pat string) (map[*sakuravps.Server][]sakuravps.ServerInterface, error) {
	result := map[*sakuravps.Server][]sakuravps.ServerInterface{}
	servers, err := s.GetFilteredServerList(pat)
	if err != nil {
		return nil, err
	}
	if len(servers) == 0 {
		return nil, fmt.Errorf("[savac] no server found with the pattern %s", pat)
	}

	var ifs []sakuravps.ServerInterface
	for _, sv := range servers {
		ifs, err = s.GetServerInterfaces(sv.Id)
		if err != nil {
			return result, err
		}
		result[sv] = ifs
	}
	return result, nil
}

func (s *SavaClient) SetServerInterfaceConnection(interfaceId int32, switchId int32) error {
	var serverId int32 = 0
	interfaces, err := s.GetAllServerInterface()
	var targetInterface sakuravps.ServerInterface
	if err != nil {
		return err
	}
	for svId, inf := range interfaces {
		for _, i := range inf {
			if i.Id == interfaceId {
				serverId = svId
				targetInterface = i
				break
			}
		}
	}
	if serverId == 0 {
		return fmt.Errorf("[savac] no server interface for interface id %d", interfaceId)
	}

	switch switchId {
	case 1:
		fmt.Fprintf(os.Stderr, "[warning] interface %d connects to the global network", targetInterface.Id)
		targetInterface.SwitchId = *sakuravps.NewNullableInt32(nil)
		targetInterface.SetConnectTo("global")
		targetInterface.SetConnectableToGlobalNetwork(true)
	case 0:
		targetInterface.ConnectTo.Unset()
		targetInterface.SwitchId.Unset()
	default:
		targetInterface.SetSwitchId(switchId)
	}

	req := s.RawClient.ServerAPI.PutServerInterface(context.Background(), serverId, targetInterface.Id)
	req = req.ServerInterface(targetInterface)
	_, resp, err := s.RawClient.ServerAPI.PutServerInterfaceExecute(req)
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}
	return nil
}
