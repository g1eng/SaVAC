package helper

import (
	"context"
	"testing"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/core"
	"github.com/g1eng/savac/testutil/test_parameters"
	"github.com/urfave/cli/v3"
)

var (
	mon = test_parameters.DummyMonitoring
	m   = test_parameters.DefaultTcpMonitoring
	p   = 8080
	m1  = core.ServerMonitoringMeta{
		Id:                   m.Id,
		Name:                 m.Name,
		ServerId:             test_parameters.SampleRegisteredServerId,
		Description:          m.Description,
		MonitoringResourceId: m.MonitoringResourceId,
		UpdateStatus:         m.UpdateStatus,
		Settings: core.ServerMonitoringSettingsMeta{
			Enabled: m.Settings.Enabled,
			HealthCheck: core.HealthCheckMeta{
				Port:              sakuravps.NullableInt{},
				Host:              sakuravps.NullableString{},
				Path:              sakuravps.NullableString{},
				BasicAuthUsername: sakuravps.NullableString{},
				BasicAuthPassword: sakuravps.NullableString{},
				Status:            sakuravps.NullableInt{},
				IntervalMinutes:   5,
				Protocol:          "ping",
				Sni:               sakuravps.NullableBool{},
			},
			Notification: m.Settings.Notification,
		},
	}
)

func Test_PrintYaml(t *testing.T) {
	err := PrintYaml(test_parameters.DummyServers)
	if err != nil {
		t.Error(err)
	}
}

func Test_PrintJson(t *testing.T) {
	err := PrintJson(test_parameters.DummyServers)
	if err != nil {
		t.Error(err)
	}
}

func Test_PrintTableForServerInfo(t *testing.T) {
	PrintTableForServerInfo(test_parameters.DefaultServer)
}

func Test_PrintMonitoringList(t *testing.T) {
	monMeta := []core.ServerMonitoringMeta{
		m1, m1, m1, m1, m1,
	}
	mon0 := []core.ServerMonitoringMeta{}
	PrintMonitoringList(monMeta)
	PrintMonitoringList(mon0)
}
func Test_PrintMonitoringDetailTable(t *testing.T) {
	p := 8080
	h := "example.com"
	pa := "/some/resource"
	m2 := m1
	m2.Settings.HealthCheck.Protocol = "tcp"
	m2.Settings.HealthCheck.Port = *sakuravps.NewNullableInt(&p)
	m3 := m1
	m3.Settings.HealthCheck.Protocol = "ssh"
	m3.Settings.HealthCheck.Port = *sakuravps.NewNullableInt(&p)
	m4 := m1
	m4.Settings.HealthCheck.Protocol = "smtp"
	m4.Settings.HealthCheck.Port = *sakuravps.NewNullableInt(&p)
	m5 := m1
	m5.Settings.HealthCheck.Protocol = "pop3"
	m5.Settings.HealthCheck.Port = *sakuravps.NewNullableInt(&p)
	m6 := m1
	m6.Settings.HealthCheck.Protocol = "imap"
	m6.Settings.HealthCheck.Port = *sakuravps.NewNullableInt(&p)
	m7 := m1
	m7.Settings.HealthCheck.Protocol = "http"
	m7.Settings.HealthCheck.Port = *sakuravps.NewNullableInt(&p)
	m7.Settings.HealthCheck.Host = *sakuravps.NewNullableString(&h)
	m7.Settings.HealthCheck.Path = *sakuravps.NewNullableString(&pa)
	m8 := m1
	m8.Settings.HealthCheck.Protocol = "https"
	m8.Settings.HealthCheck.Port = *sakuravps.NewNullableInt(&p)
	m8.Settings.HealthCheck.Host = *sakuravps.NewNullableString(&h)
	m8.Settings.HealthCheck.Path = *sakuravps.NewNullableString(&pa)
	b := false
	m8.Settings.HealthCheck.Sni = *sakuravps.NewNullableBool(&b)
	monMeta := []core.ServerMonitoringMeta{
		m1, m2, m3, m4, m5, m6, m7, m8,
		m1, m2, m3, m4, m5, m6, m7, m8,
	}
	PrintMonitoringDetailTable(map[int32][]core.ServerMonitoringMeta{
		32: monMeta,
		33: monMeta,
	})
}

func Test_PrintTableForServerInterface(t *testing.T) {
	ifs := test_parameters.DummyInterfaces
	var ids []int32
	for _, i := range ifs {
		ids = append(ids, i.Id)
	}
	s := sakuravps.Switch{
		Id:          0,
		Name:        "neko",
		Description: "nekoneko",
		SwitchCode:  "neekkon",
		Zone: sakuravps.SwitchZone{
			Code: "nk1",
			Name: "neko-special",
		},
		ServerInterfaces:    ids,
		NfsServerInterfaces: nil,
		ExternalConnection:  sakuravps.NullableSwitchExternalConnection{},
	}
	sw := []sakuravps.Switch{
		s, s, s, s, s,
	}
	PrintTableForServerInterfaces("hoge", ifs, sw)
}

func Test_CheckArgsExist(t *testing.T) {
	tables := []struct {
		name      string
		args      []string
		wantError bool
	}{
		{
			name:      "no args should return error",
			args:      []string{"progname0"},
			wantError: true,
		},
		{
			name:      "no error for an argument",
			args:      []string{"progname0", "ok"},
			wantError: false,
		},
		{
			name:      "no error for several arguments",
			args:      []string{"progname0", "foo", "bar", "baz"},
			wantError: false,
		},
	}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			c := cli.Command{
				Name:   tt.name,
				Before: CheckArgsExist,
				Action: func(ctx context.Context, command *cli.Command) error {
					return nil
				},
			}
			err := c.Run(context.Background(), tt.args)
			if tt.wantError {
				if err == nil {
					t.Errorf("CheckArgsExist() error = %v, wantErr %v", err, tt.wantError)
				}
			} else if err != nil {
				t.Errorf("CheckArgsExist() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
