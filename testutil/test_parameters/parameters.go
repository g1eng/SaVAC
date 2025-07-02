package test_parameters

import (
	"fmt"
	"os"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
)

var (
	SampleResolvableServerARecord = os.Getenv("SAKURA_VPS_API_TESTING_HOST")
	//sampleWebhookUrl                          = os.Getenv("SAVAC_WEBHOOK_URL")
	//mockserverPort                            = 8080

	SampleUnregisteredServerId           int32  = 123
	SampleUnregisteredServerName         string = "kamagaya-123-tokoro10-geogia"
	SampleUnregisteredServer1InterfaceId int32  = 19
	SampleUnregisteredSwitchId           int32  = 213
)

const (
	Fixture429Text = "{\"code\": 429, \"message\": \"Exceeded limit\"}"

	SampleRegisteredServerIdString              = "4312339"
	SampleRegisteredServerId              int32 = 4312339
	SampleRegisteredServerHostname              = "test-server"
	SampleTagKey                                = "key1"
	SampleTagValue                              = "value1"
	SampleRegisteredServerInterfaceIdText       = "43058438"
	SampleRegisteredServerInterfaceId     int32 = 43058438
	SampleRegisteredSwitchIdText                = "102006233"
	SampleRegisteredSwitchId              int32 = 102006233
	SampleRegisteredSwitchName                  = "registered-switch"
	SampleNewSwitchName                         = "test-switch"
	SampleRegisteredNfsIdText                   = "2343"
	SampleRegisteredNfsId                 int32 = 2343
	SampleRegisteredNfsName                     = "t2fs-test"
	DefaultNfsServerInterfaceId           int32 = 43193803
	SampleRegisteredNfsSwitchIdText             = "102006233"
	SampleRegisteredNfsSwitchId           int32 = 102006233
	SampleRegisteredApiKeyName                  = "sample-key-testing"
	SampleRegisteredRoleId                int32 = 388
	SampleRegisteredRoleName                    = "test-role-dummy"
	SampleUnregisteredRoleId              int32 = 999999
	SampleUnregisteredRoleName                  = "test-role-this-is-not-exist"
)

var (
	FakeServerEndpoint = map[string]string{
		"cmd0":           "127.0.0.1:17380",
		"cmd":            "127.0.0.1:17382",
		"pkg/server":     "127.0.0.1:18080",
		"pkg/switch":     "127.0.0.1:18081",
		"pkg/nfs":        "127.0.0.1:18082",
		"pkg/monitoring": "127.0.0.1:18083",
		"pkg/apikey":     "127.0.0.1:18084",
		"pkg/misc":       "127.0.0.1:18085",
	}
	FaultEndpoint = map[string]string{
		"cmd0":           "127.0.0.1:17381",
		"cmd":            "127.0.0.1:17482",
		"pkg/server":     "127.0.0.1:18090",
		"pkg/switch":     "127.0.0.1:18091",
		"pkg/nfs":        "127.0.0.1:18092",
		"pkg/monitoring": "127.0.0.1:18093",
		"pkg/apikey":     "127.0.0.1:18094",
		"pkg/misc":       "127.0.0.1:18095",
	}

	SampleRegisteredApiKeyId int32 = 359
	sampleHostname                 = "server1.example.com"
	ipv6SampleAddress              = "db02::2"
	ipv6SamplePrefixLength   int32 = 96
	DefaultIpv4                    = sakuravps.ServerIpv4{
		Address:     "192.0.2.2",
		Netmask:     "255.255.255.0",
		Gateway:     "192.0.2.1",
		Nameservers: nil,
		Hostname:    sampleHostname,
		Ptr:         sampleHostname,
	}
	DefaultIpv6 = sakuravps.ServerIpv6{
		Address:     *sakuravps.NewNullableString(&ipv6SampleAddress),
		Prefixlen:   *sakuravps.NewNullableInt32(&ipv6SamplePrefixLength),
		Gateway:     sakuravps.NullableString{},
		Nameservers: nil,
		Hostname:    *sakuravps.NewNullableString(&sampleHostname),
		Ptr:         *sakuravps.NewNullableString(&sampleHostname),
	}
	DefaultNfsIpv4         = sakuravps.NewNfsServerIpv4WithDefaults()
	DefaultZone            = sakuravps.NewServerZoneWithDefaults()
	DefaultServer          = &sakuravps.Server{Id: SampleRegisteredServerId, Name: SampleRegisteredServerHostname, Description: fmt.Sprintf("%s:%s", SampleTagKey, SampleTagValue), Ipv4: DefaultIpv4, Ipv6: DefaultIpv6, CpuCores: 2, MemoryMebibytes: 512, Zone: *DefaultZone, PowerStatus: "power_off"}
	DefaultSwitch          = &sakuravps.Switch{Id: SampleRegisteredSwitchId, Name: SampleRegisteredSwitchName, Description: "hoge"}
	SpecialServer          = &sakuravps.Server{Id: 12345, Name: "test-server-2", Description: "", Ipv4: DefaultIpv4, Ipv6: DefaultIpv6, CpuCores: 2, MemoryMebibytes: 512, Zone: *DefaultZone, PowerStatus: "power_off"}
	DefaultInterface       = &sakuravps.ServerInterface{Id: SampleRegisteredServerInterfaceId, DisplayName: "", ConnectableToGlobalNetwork: true, ConnectTo: *nullString, SwitchId: *nullInt}
	AnotherInterface       = &sakuravps.ServerInterface{Id: SampleRegisteredServerInterfaceId + 1, DisplayName: "", ConnectableToGlobalNetwork: false, ConnectTo: *nullString, SwitchId: *anInt}
	DefaultNfsStorageInner = &[]sakuravps.NfsServerStorageInner{{
		Type:          "ssd",
		SizeGibibytes: 100,
	}}
	DefaultNfsServer                = &sakuravps.NfsServer{Id: SampleRegisteredNfsId, Name: SampleRegisteredNfsName, Description: fmt.Sprintf("%s:%s", SampleTagKey, SampleTagValue), Ipv4: *DefaultNfsIpv4, SettingStatus: "unknown", Storage: *DefaultNfsStorageInner, PowerStatus: "on", Zone: *DefaultZone}
	connectToString                 = "switch"
	connectedSwitchId         int32 = 38291
	DefaultNfsServerInterface       = &sakuravps.NfsServerInterface{
		ConnectTo: *sakuravps.NewNullableString(&connectToString),
		Mac:       "1F:1F:1F:1F:1F:FF",
		SwitchId:  *sakuravps.NewNullableInt32(&connectedSwitchId),
	}
	DefaultNotification = sakuravps.ServerMonitoringSettingsNotification{
		Email:         sakuravps.ServerMonitoringSettingsNotificationEmail{Enabled: true},
		IntervalHours: 1,
	}
	WebhookNotification = sakuravps.ServerMonitoringSettingsNotification{
		Email:           sakuravps.ServerMonitoringSettingsNotificationEmail{Enabled: true},
		IncomingWebhook: sakuravps.ServerMonitoringSettingsNotificationIncomingWebhook{Enabled: false, WebhooksUrl: *sakuravps.NewNullableString(nil)},
		IntervalHours:   1,
	}
	DefaultTcpMonitoring = sakuravps.ServerMonitoring{
		Id: 1, Name: "test", Description: "", MonitoringResourceId: "TEST", UpdateStatus: "updating", Settings: sakuravps.ServerMonitoringSettings{
			Enabled: false,
			HealthCheck: sakuravps.ServerMonitoringSettingsHealthCheck{
				HealthCheckTcp: &sakuravps.HealthCheckTcp{
					Protocol:        "tcp",
					Port:            8080,
					IntervalMinutes: 1,
				},
			},
			Notification: DefaultNotification,
		},
	}
	DefaultPingMonitoring = sakuravps.ServerMonitoring{
		Id: 2, Name: "test", Description: "", MonitoringResourceId: "TEST", UpdateStatus: "updating", Settings: sakuravps.ServerMonitoringSettings{
			Enabled: false,
			HealthCheck: sakuravps.ServerMonitoringSettingsHealthCheck{
				HealthCheckPing: &sakuravps.HealthCheckPing{
					Protocol:        "ping",
					IntervalMinutes: 1,
				},
			},
			Notification: WebhookNotification,
		},
	}
	nullString   = sakuravps.NewNullableString(nil)
	a            = "ok"
	aString      = sakuravps.NewNullableString(&a)
	nullInt      = sakuravps.NewNullableInt32(nil)
	aIntBody     = int32(1)
	anInt        = sakuravps.NewNullableInt32(&aIntBody)
	DummyServers = []sakuravps.Server{
		*DefaultServer, *DefaultServer, *DefaultServer, *DefaultServer, *DefaultServer,
		*DefaultServer, *DefaultServer, *DefaultServer, *DefaultServer, *DefaultServer,
		*DefaultServer, *DefaultServer, *DefaultServer, *DefaultServer, *DefaultServer,
		*DefaultServer, *DefaultServer, *DefaultServer, *DefaultServer, *DefaultServer,
		*SpecialServer, *DefaultServer,
	}
	DummySwitches = []sakuravps.Switch{
		*DefaultSwitch, *DefaultSwitch, *DefaultSwitch,
	}
	DummyInterfaces = []sakuravps.ServerInterface{
		*DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface,
		*DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface,
		*DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface,
		*DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface,
		*DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface,
		*DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface,
		*DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface,
		*DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface, *DefaultInterface, *AnotherInterface,
	}
	DummyMonitoring = []sakuravps.ServerMonitoring{
		DefaultTcpMonitoring, DefaultPingMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring,
		DefaultTcpMonitoring, DefaultPingMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring,
		DefaultTcpMonitoring, DefaultPingMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring, DefaultTcpMonitoring,
	}
	DummyInvalidMonitoring       = append(DummyMonitoring, sakuravps.ServerMonitoring{})
	DefaultApiKeyName            = "test-api-key"
	SampleUnregisteredApiKeyName = "savac-test-key-not-registered"
	DefaultApiRoleName           = "sample-role"
	DefaultApiKey                = sakuravps.ApiKey{
		Id:    SampleRegisteredApiKeyId,
		Name:  DefaultApiKeyName,
		Role:  SampleRegisteredRoleId,
		Token: "okokokokokokokokokokokokokokoko",
	}
	DefaultRole = sakuravps.Role{
		Id:                  SampleRegisteredRoleId,
		Name:                SampleRegisteredRoleName,
		Description:         "ok",
		PermissionFiltering: "disabled",
		AllowedPermissions:  nil,
		ResourceFiltering:   "disabled",
		AllowedResources:    *sakuravps.NewNullableRoleAllowedResources(nil),
	}
	DefaultPermission = sakuravps.Permission{
		Code:     "DEFAULT_PERM",
		Name:     "default-permission-dummy",
		Category: "dummy",
	}
	DummyAllowedResources = sakuravps.RoleAllowedResources{
		Servers:    []int32{123, 345, 567},
		Switches:   nil,
		NfsServers: []int32{12357},
	}
	DummyAnotherRole = sakuravps.Role{
		Id:                  234,
		Name:                "another-role-dummy",
		Description:         "nok",
		PermissionFiltering: "enabled",
		AllowedPermissions:  []string{"this-is-dummy"},
		ResourceFiltering:   "enabled",
		AllowedResources:    *sakuravps.NewNullableRoleAllowedResources(&DummyAllowedResources),
	}
	DummyApiKeys     = []sakuravps.ApiKey{DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey, DefaultApiKey}
	DummyRoles       = []sakuravps.Role{DefaultRole, DummyAnotherRole, DefaultRole, DummyAnotherRole, DefaultRole, DummyAnotherRole, DefaultRole, DummyAnotherRole, DummyAnotherRole, DefaultRole, DummyAnotherRole, DefaultRole, DummyAnotherRole}
	DummyPermissions = []sakuravps.Permission{DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission, DefaultPermission}
)
