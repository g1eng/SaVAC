package vps

import (
	"context"
	"encoding/json"
	"io"
	"strconv"
	"time"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/core"
)

func (s *SavaClient) SetMonitoringIntervalMinutes(interval int32) {
	s.monitoringIntervalMinutes = interval
}

func (s *SavaClient) GetAllMonitoring() ([]core.ServerMonitoringMeta, error) {
	servers, err := s.GetAllServers()
	if err != nil {
		return nil, err
	}
	var bulkResp []sakuravps.ServerMonitoring
	var metaResp []core.ServerMonitoringMeta
	for _, sv := range servers {
		res, rawResp, err := s.RawClient.ServerAPI.GetServerMonitoringList(context.Background(), sv.Id).Execute()
		if res != nil && res.Results != nil {
			bulkResp = append(bulkResp, res.Results...)
			m2, err := mapMonitoringToMetaMonitoring(bulkResp)
			if err != nil {
				return nil, err
			}
			for i := range m2 {
				m2[i].ServerId = sv.Id
			}
			metaResp = append(metaResp, m2...)
		} else if res != nil && err != nil {
			buf, _ := io.ReadAll(rawResp.Body)
			var meta core.MonitoringListResponseMeta
			err = json.Unmarshal(buf, &meta)
			if err != nil {
				return nil, err
			}
			//fmt.Fprintf(os.Stderr, "%m\n\n%m\n", string(buf), err.Error())
			//fmt.Fprintf(os.Stderr, "%m\n", err.Error())
			//err = m.ListMo(sv.Id)
			for i := range meta.Results {
				meta.Results[i].ServerId = sv.Id
			}
			metaResp = append(metaResp, meta.Results...)
		} else if err != nil {
			return nil, core.EncodeHttpError(rawResp, err)
		}
		time.Sleep(time.Millisecond * 350)
	}
	return metaResp, nil
}

func (s *SavaClient) GetMonitoringListByServerId(serverId int32) ([]core.ServerMonitoringMeta, error) {
	res, rawResp, err := s.RawClient.ServerAPI.GetServerMonitoringList(context.Background(), serverId).Execute()
	//res, rawResp, err = s.RawClient.ServerAPI.GetServerMonitoringList(context.Background(), serverId).Execute()
	var (
		meta   core.MonitoringListResponseMeta
		result []core.ServerMonitoringMeta
		page   int32 = 1
	)
	for {
		if res != nil && err != nil {
			buf, _ := io.ReadAll(rawResp.Body)
			//fmt.Fprintf(os.Stderr, "%m\n\n%m\n", string(buf), err.Error())
			err = json.Unmarshal(buf, &meta)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, core.EncodeHttpError(rawResp, err)
		} else {
			meta.Results, err = mapMonitoringToMetaMonitoring(res.Results)
		}
		result = append(result, meta.Results...)
		if !meta.Next.IsSet() || meta.Next.Get() == nil {
			break
		} else {
			//log.Printf("next: %m\n", *meta.Next.Get())
			page += 1
			res, rawResp, err = s.RawClient.ServerAPI.GetServerMonitoringList(context.Background(), serverId).Page(page).Execute()
		}
	}
	//if res != nil && res.Results != nil {
	//	m, err := mapMonitoringToMetaMonitoring(res.Results)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return m, nil
	//}
	//
	return result, err
}

func (s *SavaClient) AddPingMonitoringForServer(serverId int32, name string, notifier *sakuravps.ServerMonitoringSettingsNotification) error {
	//func (s *SavaClient) AddPingMonitoringForServer(serverId int32, name string, description string) error {
	healthCheckPing := sakuravps.NewHealthCheckPing("ping", s.monitoringIntervalMinutes)
	healthcheck := sakuravps.ServerMonitoringSettingsHealthCheck{
		HealthCheckPing: healthCheckPing,
	}

	settings := sakuravps.NewServerMonitoringSettings(true, healthcheck, *notifier)

	req := sakuravps.NewServerMonitoringWithDefaults()
	req.SetSettings(*settings)
	req.SetName(name)
	req.SetDescription(name)
	//req.SetDescription(description)

	_, rawResp, err := s.RawClient.ServerAPI.PostServerMonitoring(context.Background(), serverId).ServerMonitoring(*req).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}

	return nil
}

func (s *SavaClient) AddSshMonitoringForServer(serverId int32, name string, port int32, notifier *sakuravps.ServerMonitoringSettingsNotification) error {
	healthCheckSsh := sakuravps.NewHealthCheckSsh("ssh", port, s.monitoringIntervalMinutes)
	healthCheck := sakuravps.ServerMonitoringSettingsHealthCheck{
		HealthCheckSsh: healthCheckSsh,
	}

	settings := sakuravps.NewServerMonitoringSettings(true, healthCheck, *notifier)
	return s.createNewMonitoring(serverId, name, *settings)
}

func (s *SavaClient) AddTcpMonitoringForServer(serverId int32, name string, port int32, notifier *sakuravps.ServerMonitoringSettingsNotification) error {
	healthCheckTcp := sakuravps.NewHealthCheckTcp("tcp", port, s.monitoringIntervalMinutes)
	healthCheck := sakuravps.ServerMonitoringSettingsHealthCheck{
		HealthCheckTcp: healthCheckTcp,
	}

	settings := sakuravps.NewServerMonitoringSettings(true, healthCheck, *notifier)

	return s.createNewMonitoring(serverId, name, *settings)
}

func (s *SavaClient) AddSmtpMonitoringForServer(serverId int32, name string, port int32, notifier *sakuravps.ServerMonitoringSettingsNotification) error {
	healthCheckSmtp := sakuravps.NewHealthCheckSmtp("smtp", port, s.monitoringIntervalMinutes)
	healthCheck := sakuravps.ServerMonitoringSettingsHealthCheck{
		HealthCheckSmtp: healthCheckSmtp,
	}

	settings := sakuravps.NewServerMonitoringSettings(true, healthCheck, *notifier)

	return s.createNewMonitoring(serverId, name, *settings)
}

func (s *SavaClient) AddPop3MonitoringForServer(serverId int32, name string, port int32, notifier *sakuravps.ServerMonitoringSettingsNotification) error {
	healthCheckPop := sakuravps.NewHealthCheckPop3("pop3", port, s.monitoringIntervalMinutes)
	healthCheck := sakuravps.ServerMonitoringSettingsHealthCheck{
		HealthCheckPop3: healthCheckPop,
	}

	settings := sakuravps.NewServerMonitoringSettings(true, healthCheck, *notifier)
	return s.createNewMonitoring(serverId, name, *settings)
}

func (s *SavaClient) AddHttpMonitoringForServer(serverId int32, name string, target *core.HttpMonitoringTarget, notifier *sakuravps.ServerMonitoringSettingsNotification) error {
	healthCheckHttp := sakuravps.NewHealthCheckHttp(target.Port, *target.Host, target.Path, *target.BasicUserName, *target.BasicAuthPassword, target.Status, s.monitoringIntervalMinutes, "http")
	healthCheck := sakuravps.ServerMonitoringSettingsHealthCheck{
		HealthCheckHttp: healthCheckHttp,
	}
	settings := sakuravps.NewServerMonitoringSettings(true, healthCheck, *notifier)
	return s.createNewMonitoring(serverId, name, *settings)
}

func (s *SavaClient) AddHttpsMonitoringForServer(serverId int32, name string, target *core.HttpMonitoringTarget, notifier *sakuravps.ServerMonitoringSettingsNotification) error {
	var sni bool
	if target.Sni != nil && *target.Sni {
		sni = true
	}
	healthCheckHttps := sakuravps.NewHealthCheckHttps(target.Port, *target.Host, target.Path, *target.BasicUserName, *target.BasicAuthPassword, target.Status, s.monitoringIntervalMinutes, "https", sni)
	healthCheck := sakuravps.ServerMonitoringSettingsHealthCheck{
		HealthCheckHttps: healthCheckHttps,
	}
	settings := sakuravps.NewServerMonitoringSettings(true, healthCheck, *notifier)
	return s.createNewMonitoring(serverId, name, *settings)
}

func (s *SavaClient) createNewMonitoring(serverId int32, name string, setting sakuravps.ServerMonitoringSettings) error {
	req := sakuravps.NewServerMonitoringWithDefaults()
	req.SetSettings(setting)
	req.SetName(name)
	req.SetDescription(name)

	res, rawResp, err := s.RawClient.ServerAPI.PostServerMonitoring(context.Background(), serverId).ServerMonitoring(*req).Execute()
	if res != nil && err != nil {
		return nil
	} else if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

func (s *SavaClient) DeleteAllMonitoringByServerId(serverId int32) error {
	res, err := s.GetMonitoringListByServerId(serverId)
	if err != nil {
		return err
		//return encodeHttpError(rawResp, err)
	}
	for _, monitoring := range res {
		rawResp, err := s.RawClient.ServerAPI.DeleteServerMonitorings(context.Background(), serverId, monitoring.Id).Execute()
		if err != nil {
			return core.EncodeHttpError(rawResp, err)
		}
	}
	return nil
}

func (s *SavaClient) DeleteAllMonitoringByServerPattern(serverPat string) error {
	var serverIds []int32
	serverId, err := strconv.Atoi(serverPat)
	if err != nil {
		sv, err := s.GetServerIdsByNamePattern(serverPat)
		if err != nil {
			return err
		}
		serverIds = sv
	} else {
		serverIds = append(serverIds, int32(serverId))
	}
	for _, svId := range serverIds {
		err = s.DeleteAllMonitoringByServerId(svId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SavaClient) DeleteMonitoringByServerAndMonitoringId(serverId int32, monitoringId int32) error {
	rawResp, err := s.RawClient.ServerAPI.DeleteServerMonitorings(context.Background(), serverId, monitoringId).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

func mapMonitoringToMetaMonitoring(res []sakuravps.ServerMonitoring) (meta []core.ServerMonitoringMeta, err error) {
	for _, r := range res {
		var m core.ServerMonitoringMeta
		b, _ := json.Marshal(r)
		err = json.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}
		meta = append(meta, m)
	}
	return meta, nil
}
