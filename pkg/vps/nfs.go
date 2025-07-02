package vps

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/core"
)

func (s *SavaClient) GetNFSInterface(id int32) (*sakuravps.NfsServerInterface, error) {
	res, resp, err := s.RawClient.NfsServerAPI.GetNfsServerInterface(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(resp, err)
	}
	return res, nil
}

func (s *SavaClient) GetNFSByName(name string) (*sakuravps.NfsServer, error) {
	nfs, err := s.GetAllNFS()
	if err != nil {
		return nil, err
	}
	for _, n := range nfs {
		if n.Name == name {
			return &n, nil
		}
	}
	return nil, fmt.Errorf("nfs %s not found", name)
}

func (s *SavaClient) GetNFSListByRegex(regex *regexp.Regexp) ([]*sakuravps.NfsServer, error) {
	var res []*sakuravps.NfsServer
	nfs, err := s.GetAllNFS()
	if err != nil {
		return nil, err
	}
	for _, n := range nfs {
		if regex.MatchString(n.Name) {
			res = append(res, &n)
		}
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("pattern %s not matched for NFS names", regex)
	}
	return res, nil
}

func (s *SavaClient) SetNFSName(id int32, name string) error {
	_, resp, err := s.RawClient.NfsServerAPI.PutNfsServer(context.Background(), id).PutServerRequest(sakuravps.PutServerRequest{
		Name: name,
	}).Execute()
	b, _ := io.ReadAll(resp.Body)
	log.Println("resp:", string(b))
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}
	return nil
}

func (s *SavaClient) SetNFSDescription(id int32, description string) error {
	_, resp, err := s.RawClient.NfsServerAPI.PutNfsServer(context.Background(), id).PutServerRequest(sakuravps.PutServerRequest{
		Description: description,
	}).Execute()
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}
	return nil
}

// GetAllNFS lists NFS DummyServers in table format
func (s *SavaClient) GetAllNFS() ([]sakuravps.NfsServer, error) {
	res, resp, err := s.RawClient.NfsServerAPI.GetNfsServerList(context.Background()).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(resp, err)
	}

	return res.Results, nil
}

func (s *SavaClient) GetNFSById(id int32) (*sakuravps.NfsServer, error) {
	res, resp, err := s.RawClient.NfsServerAPI.GetNfsServer(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(resp, err)
	}
	return res, nil
}

func (s *SavaClient) PutNFSConnection(nfsId int32, switchId int32) error {
	inf := *sakuravps.NewNfsServerInterfaceWithDefaults()
	_, resp, err := s.RawClient.NfsServerAPI.GetNfsServerInterface(context.Background(), nfsId).Execute()
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}
	req := s.RawClient.NfsServerAPI.PutNfsServerInterface(context.Background(), nfsId)
	if switchId == 0 {
		inf.SwitchId.Unset()
		inf.ConnectTo.Unset()
	} else {
		inf.SwitchId.Set(&switchId)
		s := "switch"
		inf.ConnectTo.Set(&s)
	}
	req = req.NfsServerInterface(inf)
	_, resp, err = req.Execute()
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}

	return nil
}

func (s *SavaClient) GetNfsStorageInfo(nfsId int32) (*sakuravps.NfsStorageInfo, error) {
	info, resp, err := s.RawClient.NfsServerAPI.GetNfsServerStorage(context.Background(), nfsId).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(resp, err)
	} else if resp.StatusCode != http.StatusOK {
		return info, core.EncodeHttpError(resp, fmt.Errorf("%s", resp.Status))
	}
	return info, nil
}

func (s *SavaClient) ShutdownNfs(nfsId int32, isForced ...bool) error {
	f := false
	if len(isForced) > 0 {
		f = isForced[0]
	}
	req := s.RawClient.NfsServerAPI.PostNfsServerShutdown(context.Background(), nfsId)
	resp, err := req.PostServerShutdownRequest(sakuravps.PostServerShutdownRequest{Force: &f}).Execute()
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}
	return nil
}

func (s *SavaClient) StartNfs(nfsId int32) error {
	res, err := s.RawClient.NfsServerAPI.PostNfsServerPowerOn(context.Background(), nfsId).Execute()
	if err != nil {
		return core.EncodeHttpError(res, err)
	}
	return nil
}

func (s *SavaClient) RebootNfs(nfsId int32) error {
	res, err := s.RawClient.NfsServerAPI.PostNfsServerForceReboot(context.Background(), nfsId).Execute()
	if err != nil {
		return core.EncodeHttpError(res, err)
	}
	return nil
}
