package core

import (
	"net/http"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
)

type RawApi struct {
	raw *sakuravps.APIClient
}

func NewApiClient(client *sakuravps.APIClient) *RawApi {
	return &RawApi{
		client,
	}
}

func NewHttpMonitoringTarget(port int32, host string, path string) *HttpMonitoringTarget {
	return &HttpMonitoringTarget{
		Port:              port,
		Path:              path,
		Status:            http.StatusOK,
		Host:              sakuravps.NewNullableString(&host),
		BasicUserName:     sakuravps.NewNullableString(nil),
		BasicAuthPassword: sakuravps.NewNullableString(nil),
		Sni:               nil,
	}
}

type HttpMonitoringTarget struct {
	Port              int32
	Host              *sakuravps.NullableString
	BasicUserName     *sakuravps.NullableString
	BasicAuthPassword *sakuravps.NullableString
	Path              string
	Status            int32
	Sni               *bool
}

type HealthCheckMeta struct {
	Port              sakuravps.NullableInt    `json:"port"`
	Host              sakuravps.NullableString `json:"host"`
	Path              sakuravps.NullableString `json:"path"`
	BasicAuthUsername sakuravps.NullableString `json:"basic_auth_username"`
	BasicAuthPassword sakuravps.NullableString `json:"basic_auth_password"`
	Status            sakuravps.NullableInt    `json:"status"`
	IntervalMinutes   int32                    `json:"interval_minutes"`
	Protocol          string                   `json:"protocol"`
	Sni               sakuravps.NullableBool   `json:"sni"`
}

type ServerMonitoringMeta struct {
	Id                   int32                        `json:"id"`
	ServerId             int32                        `json:"server_id,omitempty"`
	Name                 string                       `json:"name"`
	Description          string                       `json:"description"`
	MonitoringResourceId string                       `json:"monitoring_resource_id"`
	UpdateStatus         string                       `json:"update_status"`
	Settings             ServerMonitoringSettingsMeta `json:"settings"`
}

type ServerMonitoringSettingsMeta struct {
	Enabled      bool                                           `json:"enabled"`
	HealthCheck  HealthCheckMeta                                `json:"health_check"`
	Notification sakuravps.ServerMonitoringSettingsNotification `json:"notification"`
}

type MonitoringListResponseMeta struct {
	Count    int32                    `json:"count"`
	Next     sakuravps.NullableString `json:"next"`
	Previous sakuravps.NullableString `json:"previous"`
	Results  []ServerMonitoringMeta   `json:"results"`
}

type FilteredResource interface {
	GetName() string
}

type FilterableResource *FilteredResource
