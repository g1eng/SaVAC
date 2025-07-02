package vps

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/g1eng/savac/pkg/core"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
)

func (s *SavaClient) GetAllServers() ([]*sakuravps.Server, error) {
	var (
		tmpResult []sakuravps.Server
		result    []*sakuravps.Server
		page      int32 = 1
	)
	req := s.RawClient.ServerAPI.GetServerList(context.Background())
	req = req.PerPage(10)
	req = req.Page(page)

	res, rawResp, err := req.Execute()
	for {
		if err != nil {
			return nil, core.EncodeHttpError(rawResp, err)
		} else if rawResp.StatusCode != http.StatusOK {
			return nil, core.EncodeHttpError(rawResp, errors.New(rawResp.Status))
		}
		tmpResult = append(tmpResult, res.Results...)
		if !res.Next.IsSet() || res.Next.Get() == nil {
			break
		} else {
			page += 1
			res, rawResp, err = req.Page(page).Execute()
		}
	}

	if len(tmpResult) == 0 {
		return nil, fmt.Errorf("[savac] not found")
	}
	for _, r := range tmpResult {
		//s, _, err := s.RawClient.ServerAPI.GetServer(context.Background(), r.Id).Execute()
		//if err != nil {
		//	return nil, err
		//}
		result = append(result, &r)
	}
	return result, nil
}

func (s *SavaClient) GetServerIdsByNamePattern(pattern string) ([]int32, error) {
	sv, err := s.GetFilteredServerList(pattern)
	if err != nil {
		return nil, err
	}
	var ids []int32
	for _, s := range sv {
		ids = append(ids, s.Id)
	}
	return ids, nil
}

// GetFilteredServerList gets server list with given regex pattern and returns an array
// of opanapi.Server objects.
func (s *SavaClient) GetFilteredServerList(pat string) ([]*sakuravps.Server, error) {
	regex, err := regexp.Compile(pat)
	if err != nil {
		return nil, err
	}
	if s.Debug {
		fmt.Fprintf(os.Stderr, "%s\n", pat)
	}

	res, err := s.GetAllServers()
	if err != nil {
		return res, err
	}

	var result []*sakuravps.Server

	for _, r := range res {
		if regex.MatchString(r.Name) {
			result = append(result, r)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("[savac] not found")
	}
	if s.Debug {
		for _, r := range result {
			fmt.Fprintf(os.Stderr, "%v\n", r)
		}
	}
	return result, nil
}

// GetServerById shows server detail information with server ID.
func (s *SavaClient) GetServerById(id int32) (*sakuravps.Server, error) {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	powerRes, _, err := s.RawClient.ServerAPI.GetServerPowerStatus(context.Background(), res.Id).Execute()
	if err != nil || powerRes == nil {
		fmt.Fprintf(os.Stderr, "[no response] ")
		return nil, err
	}
	res.PowerStatus = powerRes.Status

	return res, nil
}

// GetServerDescriptionById shows server detail information with server ID.
func (s *SavaClient) GetServerDescriptionById(id int32) ([]string, error) {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	desc := strings.Split(res.Description, "\n")
	return desc, nil
}

// GetHostname describe the name of the specified server
func (s *SavaClient) GetHostname(id int32) (string, error) {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return "", core.EncodeHttpError(rawResp, err)
	}
	return res.Name, nil
}

// SetHostName set the unique host name for the server. (not the hostname inside OS)
func (s *SavaClient) SetHostName(id int32, hostName string) error {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}

	req := sakuravps.NewPutServerRequest(hostName, res.Description)
	_, rawResp, err = s.RawClient.ServerAPI.PutServer(context.Background(), id).PutServerRequest(*req).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}

	return nil
}

// SetServerDescription set the description for the server.
func (s *SavaClient) SetServerDescription(id int32, desc string) error {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}

	req := sakuravps.NewPutServerRequest(res.Name, desc)
	_, rawResp, err = s.RawClient.ServerAPI.PutServer(context.Background(), id).PutServerRequest(*req).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

// ListServerTags shows tags for the server on description field.
func (s *SavaClient) ListServerTags(id int32) ([]string, error) {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	desc := strings.Split(res.Description, "\n")
	var tags []string
	for _, d := range desc {
		if strings.Contains(d, ":") {
			if p := strings.SplitN(d, ":", 1); len(p[0]) != 0 {
				tags = append(tags, d)
			}
		}
	}
	return tags, nil
}

// GetServerTag shows tags for the server on description field.
func (s *SavaClient) GetServerTag(id int32, tagName string) (*string, error) {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	desc := strings.Split(res.Description, "\n")
	cmp := tagName + ":"
	for _, d := range desc {
		if strings.Contains(d, cmp) {
			t := strings.TrimLeft(d, cmp)
			return &t, nil
		}
	}
	return nil, fmt.Errorf("no tag %s found", tagName)
}

// AddServerTag set a tag for the server. using description field.
func (s *SavaClient) AddServerTag(id int32, tagName string, tagValue string) error {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	desc := strings.Split(res.Description, "\n")
	found := false
	i := 0
	d := ""
	for i, d = range desc {
		if t := strings.Split(d, ":"); t[0] == tagName {
			found = true
			break
		}
	}
	if found {
		desc[i] = tagName + ":" + tagValue
	} else {
		desc = append(desc, tagName+":"+tagValue)
	}
	req := sakuravps.NewPutServerRequest(res.Name, strings.Join(desc, "\n"))
	_, rawResp, err = s.RawClient.ServerAPI.PutServer(context.Background(), id).PutServerRequest(*req).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

// StartServerWithId boots server, which is specified with server ID.
func (s *SavaClient) StartServerWithId(id int32) error {
	rawResp, err := s.RawClient.ServerAPI.PostServerPowerOn(context.Background(), id).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

// StopServerWithId shutdown server, which is specified  with server ID.
func (s *SavaClient) StopServerWithId(id int32) error {
	cond := false
	a := sakuravps.PostServerShutdownRequest{
		Force: &cond,
	}
	req := s.RawClient.ServerAPI.PostServerShutdown(context.Background(), id)
	rawResp, err := req.PostServerShutdownRequest(a).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

// ForceRebootServerWithId reboot server (forcibly) with specified server ID.
func (s *SavaClient) ForceRebootServerWithId(id int32) error {
	rawResp, err := s.RawClient.ServerAPI.PostServerForceReboot(context.Background(), id).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

// SetIpv4PtrWithServerId sets ipv4 PTR record for the server
func (s *SavaClient) SetIpv4PtrWithServerId(id int32, hostname string) error {
	req := sakuravps.NewPutServerIpv4PtrRequest(hostname)
	_, rawResp, err := s.RawClient.ServerAPI.PutServerIpv4Ptr(context.Background(), id).PutServerIpv4PtrRequest(*req).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

// GetIpv4PtrWithServerId gets IPv4 PTR record for the server
func (s *SavaClient) GetIpv4PtrWithServerId(id int32) (string, error) {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return "", core.EncodeHttpError(rawResp, err)
	}
	return res.Ipv4.Ptr, nil
}

// SetIpv6PtrWithServerId sets ipv6 PTR record for the server (now upstream API seems to provide incompatible argument name `ipv4`)
func (s *SavaClient) SetIpv6PtrWithServerId(id int32, hostname string) error {
	req := sakuravps.NewPutServerIpv4PtrRequest(hostname)
	_, rawResp, err := s.RawClient.ServerAPI.PutServerIpv6Ptr(context.Background(), id).PutServerIpv4PtrRequest(*req).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

// GetIpv6PtrWithServerId gets IPv6 PTR record for the server
func (s *SavaClient) GetIpv6PtrWithServerId(id int32) (string, error) {
	res, rawResp, err := s.RawClient.ServerAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		return "", core.EncodeHttpError(rawResp, err)
	}
	if res.Ipv6.Ptr.IsSet() {
		return *res.Ipv6.Ptr.Get(), nil
	} else {
		return "", nil
	}
}
