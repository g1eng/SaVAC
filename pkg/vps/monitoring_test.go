package vps

import (
	"fmt"
	"time"

	"github.com/g1eng/savac/pkg/core"
	"github.com/g1eng/savac/testutil/test_parameters"
)

func (v *SavaClientSuite) TestMapMonitoringToMetaMonitoring() {
	res, err := mapMonitoringToMetaMonitoring(test_parameters.DummyMonitoring)
	if err != nil {
		v.Fail("%v", err)
	}
	if len(res) != len(test_parameters.DummyMonitoring) {
		v.Fail("res length mismatch, expected %d, got %d", len(test_parameters.DummyMonitoring), len(res))
	}
}

func (v *SavaClientSuite) TestMapMonitoringToMetaMonitoring_For_InvalidPayload() {
	_, err := mapMonitoringToMetaMonitoring(test_parameters.DummyInvalidMonitoring)
	if err == nil {
		v.Fail("%v", "expected error")
	}
}

func (v *SavaClientSuite) TestSavaClient_AddPingMonitoringForServer() {
	if v.isTestAcc {
		time.Sleep(time.Minute)
	}
	err := v.monCli.AddPingMonitoringForServer(test_parameters.SampleRegisteredServerId, "test-ping", &test_parameters.DefaultNotification)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_AddTcpMonitoringForServer() {
	err := v.monCli.AddTcpMonitoringForServer(test_parameters.SampleRegisteredServerId, "test-tcp", 8080, &test_parameters.DefaultNotification)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_AddPop3MonitoringForServer() {
	err := v.monCli.AddPop3MonitoringForServer(test_parameters.SampleRegisteredServerId, "test-pop3", 8080, &test_parameters.DefaultNotification)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_DeleteAllMonitoringByServerId() {
	if v.isTestAcc {
		time.Sleep(180 * time.Second)
	}
	err := v.monCli.DeleteAllMonitoringByServerId(test_parameters.SampleRegisteredServerId)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_AddSmtpMonitoringForServer() {
	err := v.monCli.AddSmtpMonitoringForServer(test_parameters.SampleRegisteredServerId, "test-smtp", 8080, &test_parameters.DefaultNotification)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_AddHttpMonitoringForServer() {
	target := core.NewHttpMonitoringTarget(80, test_parameters.SampleResolvableServerARecord, "/")
	err := v.monCli.AddHttpMonitoringForServer(test_parameters.SampleRegisteredServerId, "test-http", target, &test_parameters.DefaultNotification)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_AddHttpsMonitoringForServer() {
	b := true
	target := core.NewHttpMonitoringTarget(443, test_parameters.SampleResolvableServerARecord, "/")
	target.Sni = &b
	err := v.monCli.AddHttpMonitoringForServer(test_parameters.SampleRegisteredServerId, "test-https", target, &test_parameters.DefaultNotification)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_GetMonitoringListByServerId() {
	_, err := v.monCli.GetMonitoringListByServerId(test_parameters.SampleRegisteredServerId)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_GetAllMonitoring() {
	_, err := v.monCli.GetAllMonitoring()
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_DeleteMonitoringByServerAndMonitoringId() {
	if v.isTestAcc {
		time.Sleep(180 * time.Second)
	}
	mon, err := v.monCli.GetMonitoringListByServerId(test_parameters.SampleRegisteredServerId)
	v.NoError(err)
	err = v.monCli.DeleteMonitoringByServerAndMonitoringId(test_parameters.SampleRegisteredServerId, mon[0].Id)
	v.NoError(err)
	if v.isTestAcc {
		err = v.monCli.DeleteMonitoringByServerAndMonitoringId(test_parameters.SampleRegisteredServerId, test_parameters.SampleUnregisteredSwitchId)
		v.Error(err)
	}
}

func (v *SavaClientSuite) TestSavaClient_DeleteAllMonitoringByUnknownServerName() {
	if v.isTestAcc {
		time.Sleep(30 * time.Second)
	}
	err := v.monCli.DeleteAllMonitoringByServerPattern(test_parameters.SampleUnregisteredServerName)
	v.Error(err)
	if v.isTestAcc {
		time.Sleep(30 * time.Second)
		id := fmt.Sprintf("%d", test_parameters.SampleUnregisteredServerId)
		err = v.monCli.DeleteAllMonitoringByServerPattern(id)
		v.Error(err)
	}
}

func (v *SavaClientSuite) TestSavaClient_DeleteAllMonitoringByServerPattern() {
	if v.isTestAcc {
		time.Sleep(180 * time.Second)
	}
	err := v.monCli.DeleteAllMonitoringByServerPattern(test_parameters.SampleRegisteredServerHostname)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_GetMonitoringListByServerId_With_Fixture() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	res, err := v.monCli.GetMonitoringListByServerId(test_parameters.DefaultServer.Id)
	if err != nil {
		v.Fail("%v", err)
	}
	v.Equal(len(res), len(test_parameters.DummyMonitoring))
}

func (v *SavaClientSuite) TestSavaClient_FaultMonitoringRequest() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	err := v.faultCli.AddPingMonitoringForServer(test_parameters.SampleRegisteredServerId, "ok", &test_parameters.DefaultNotification)
	v.Error(err)
	err = v.faultCli.AddPingMonitoringForServer(test_parameters.SampleRegisteredServerId, "ok", &test_parameters.DefaultNotification)
	v.Error(err)
	err = v.faultCli.AddTcpMonitoringForServer(test_parameters.SampleRegisteredServerId, "ok", 8080, &test_parameters.DefaultNotification)
	v.Error(err)
	err = v.faultCli.AddSmtpMonitoringForServer(test_parameters.SampleRegisteredServerId, "ok", 587, &test_parameters.DefaultNotification)
	v.Error(err)
	err = v.faultCli.AddPop3MonitoringForServer(test_parameters.SampleRegisteredServerId, "ok", 990, &test_parameters.DefaultNotification)
	v.Error(err)

	httpTarget := core.NewHttpMonitoringTarget(80, test_parameters.SampleResolvableServerARecord, "/")
	err = v.faultCli.AddHttpMonitoringForServer(test_parameters.SampleRegisteredServerId, "ok", httpTarget, &test_parameters.DefaultNotification)
	v.Error(err)
	httpTarget.Sni = nil
	err = v.faultCli.AddHttpsMonitoringForServer(test_parameters.SampleRegisteredServerId, "ok", httpTarget, &test_parameters.DefaultNotification)
	v.Error(err)
}
