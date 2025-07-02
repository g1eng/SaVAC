package vps

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/core"
)

// ListApiKeys lists ApiKey resources registered on the service.
func (s *SavaClient) ListApiKeys() ([]sakuravps.ApiKey, error) {
	var (
		k    []sakuravps.ApiKey
		req        = s.RawClient.ApiKeyAPI.GetApiKeyList(context.Background()).PerPage(100)
		page int32 = 1
	)

	for {
		res, rawResp, err := req.Execute()
		if err != nil {
			return nil, core.EncodeHttpError(rawResp, err)
		}
		if rawResp != nil && rawResp.StatusCode != http.StatusOK {
			msg, _ := io.ReadAll(rawResp.Body)
			return nil, fmt.Errorf("invalid response: %s\n%s", rawResp.Status, msg)
		}
		k = append(k, res.Results...)
		if res.Next.IsSet() && res.Next.Get() != nil {
			page += 1
			req = req.Page(page)
		} else {
			break
		}
	}
	return k, nil
}

func (s *SavaClient) CreateApiKey(name string, roleId int32) (*sakuravps.ApiKey, error) {
	req := sakuravps.ApiKey{
		Name: name,
		Role: roleId,
	}
	res, rawResp, err := s.RawClient.ApiKeyAPI.PostApiKey(context.Background()).ApiKey(req).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	return res, nil
}

func (s *SavaClient) GetApiKeyById(id int32) (*sakuravps.ApiKey, error) {
	res, rawResp, err := s.RawClient.ApiKeyAPI.GetApiKey(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	return res, nil
}

func (s *SavaClient) RotateApiKey(keyId int32) (string, error) {
	res, rawResp, err := s.RawClient.ApiKeyAPI.PostApiKeyRotate(context.Background(), keyId).Execute()
	if err != nil {
		return "", core.EncodeHttpError(rawResp, err)
	}
	if s.Debug {
		fmt.Fprintf(os.Stderr, "%v\n", res)
	}
	return res.Token, nil
}

func (s *SavaClient) DeleteApiKeyById(keyId int32) error {
	rawResp, err := s.RawClient.ApiKeyAPI.DeleteApiKey(context.Background(), keyId).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}
