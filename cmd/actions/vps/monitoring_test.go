package vps_actions

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/g1eng/savac/testutil/test_parameters"

	"github.com/urfave/cli/v3"
)

func (v *VpsActionSuite) TestGenerateServerMonitoringListAction() {
	err := testAction(actionGenerator.GenerateServerMonitoringListAction)
	if err != nil {
		v.Failf("fail", "should list all monitoring for all server: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerMonitoringListAction, "--server", test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should list monitoring for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}
}

func (v *VpsActionSuite) TestGenerateServerMonitoringList_WithServerName_Action() {
	err := testAction(actionGenerator.GenerateServerMonitoringListAction, "--server", test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should list monitoring for server %s: %v", test_parameters.SampleRegisteredServerHostname, err)
	}
}

func (v *VpsActionSuite) TestScenario_GenerateServerMonitoring_Create_Delete_Action() {
	// PING
	if v.isTestAcc {
		time.Sleep(60 * time.Second)
	} else {
		v.T().SkipNow()
	}
	err := testAction(actionGenerator.GenerateServerMonitoringPingAction)
	if err == nil {
		v.Failf("fail", "should fail without arguments")
	}
	err = testAction(actionGenerator.GenerateServerMonitoringPingAction, test_parameters.SampleRegisteredServerIdString, "test-monitoring-ping")
	if err != nil {
		v.Failf("fail", "should create server monitoring with email notification by default: server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}
	err = testAction(actionGenerator.GenerateServerMonitoringPingAction, test_parameters.SampleRegisteredServerIdString, "test-monitoring-ping", "email")
	if err != nil {
		v.Failf("fail", "should create server monitoring with email notification for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}

	err = testAction(actionGenerator.GenerateServerMonitoringPingAction, test_parameters.SampleRegisteredServerHostname, "test-monitoring-ping", "email")
	if err != nil {
		v.Failf("fail", "should create server monitoring with server hostname %s: %v", test_parameters.SampleRegisteredServerHostname, err)
	}

	err = testAction(actionGenerator.GenerateServerMonitoringPingAction, test_parameters.SampleRegisteredServerIdString, "test-monitoring-ping", "webhook", sampleWebhookUrl)
	if err != nil {
		v.Failf("fail", "should create server monitoring with webhook notification for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}
	err = testAction(actionGenerator.GenerateServerMonitoringPingAction, test_parameters.SampleRegisteredServerIdString, "test-monitoring-ping", "webhook", sampleWebhookUrl, "test-team", "test-channel")
	if err != nil {
		v.Failf("fail", "should create server monitoring with webhook notification for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}

	err = testAction(actionGenerator.GenerateServerMonitoringPingAction, test_parameters.SampleRegisteredServerIdString, "test-monitoring-ping", "webhook")
	if err == nil {
		v.Failf("fail", "should fail without arguments after `webhook`")
	}

	//TCP-like protocols

	testTargets := map[string]cli.ActionFunc{
		"ssh":  actionGenerator.GenerateServerMonitoringSshAction,
		"tcp":  actionGenerator.GenerateServerMonitoringTcpAction,
		"smtp": actionGenerator.GenerateServerMonitoringSmtpAction,
		"pop3": actionGenerator.GenerateServerMonitoringPop3Action,
	}

	testTcpLikeMonExec := func(protocol string, target cli.ActionFunc) {
		err := testAction(target)
		if err == nil {
			v.Failf("fail", "should fail without arguments")
		}
		err = testAction(target, test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+protocol, "email")
		if err != nil {
			v.Failf("fail", "should create server monitoring with email notification: server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
		}
		err = testAction(target, test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+protocol, "webhook")
		if err == nil {
			v.Failf("fail", "should fail without arguments after `webhook`")
		}
		err = testAction(actionGenerator.GenerateServerMonitoringPingAction, "--port", "2222", test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+protocol, "webhook", sampleWebhookUrl)
		if err != nil {
			v.Failf("fail", "should create server monitoring with webhook notification for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
		}
		err = testAction(actionGenerator.GenerateServerMonitoringPingAction, test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+protocol, "webhook", sampleWebhookUrl, "test-team", "test-channel")
		if err != nil {
			v.Failf("fail", "should create server monitoring with webhook notification for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
		}
		err = testAction(target, "--port", "2222", test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+protocol)
		if err != nil {
			v.Failf("fail", "should create server monitoring with changed port: server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
		}
	}

	if v.isTestAcc {
		time.Sleep(30 * time.Second)
	}
	for protocol, target := range testTargets {
		testTcpLikeMonExec(protocol, target)
		if v.isTestAcc {
			time.Sleep(30 * time.Second)
		}
	}

	if v.isTestAcc {
		time.Sleep(60 * time.Second)
	}

	mon, err := actionGenerator.ApiClient.GetMonitoringListByServerId(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("fail", "failed to get server monitoring for %s\n", test_parameters.SampleRegisteredServerIdString)
	} else if len(mon) < 3 {
		v.Failf("fail", "too few server monitoring set for %s\n", test_parameters.SampleRegisteredServerIdString)
	}
	err = testAction(actionGenerator.GenerateServerMonitoringInfoAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should list all monitoring for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}

	err = testAction(actionGenerator.GenerateServerMonitoringDeleteAction)
	if err == nil {
		v.Failf("fail", "should fail without arguments")
	}

	err = testAction(actionGenerator.GenerateServerMonitoringDeleteAction, test_parameters.SampleRegisteredServerHostname, fmt.Sprintf("%d", int(mon[0].Id)))
	if err != nil {
		v.Failf("fail", "should delete server monitoring for server %s with monitoring id %d: %v", test_parameters.SampleRegisteredServerHostname, mon[0].Id, err)
	}
	err = testAction(actionGenerator.GenerateServerMonitoringDeleteAction, test_parameters.SampleRegisteredServerIdString, fmt.Sprintf("%d", int(mon[1].Id)))
	if err != nil {
		v.Failf("fail", "should delete server monitoring for server id %s with monitoring id %d: %v", test_parameters.SampleRegisteredServerIdString, mon[1].Id, err)
	}
	//err = testAction(actionGenerator.GenerateServerMonitoringDeleteAction, SampleRegisteredServerIdString, fmt.Sprintf("%d", int(res.Results[0].Id)))
	//if err != nil {
	//	v.Failf("fail", "should delete server monitoring for monitoring %s on server %s: %v", fmt.Sprintf("%d", int(res.Results[0].Id)), SampleRegisteredServerIdString, err)
	//}
	res, _, err := actionGenerator.ApiClient.RawClient.ServerAPI.GetServerMonitoringList(context.Background(), test_parameters.SampleRegisteredServerId).Execute()
	if res == nil && err != nil {
		v.Failf("fail", "[testsuite] failed to get server monitoring list for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}
	if res != nil && err != nil {
		println("ok")
	}
	err = testAction(actionGenerator.GenerateServerMonitoringDeleteAction)
	if err == nil {
		v.Failf("fail", "should fail without arguments")
	}

	if v.isTestAcc {
		time.Sleep(180 * time.Second)
	}
	err = testAction(actionGenerator.GenerateServerMonitoringDeleteAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should delete all server monitoring for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}
}

func (v *VpsActionSuite) TestHttpAndHttpsMonitoringAction() {
	testTargets := map[string]func(context.Context, *cli.Command) error{
		"http":  actionGenerator.GenerateServerMonitoringHttpAction,
		"https": actionGenerator.GenerateServerMonitoringHttpsAction,
	}
	if v.isTestAcc {
		time.Sleep(30 * time.Second)
	}

	for k, target := range testTargets {
		if v.isTestAcc {
			time.Sleep(30 * time.Second)
		}
		err := testAction(target)
		if err == nil {
			v.Failf("fail", "should fail without arguments")
		}
		err = testAction(target, test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+k)
		if err == nil {
			v.Failf("fail", "should fail without --host parameter")
		}
		err = testAction(target, "--host", "testbed.savac.s0csec1.org", test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+k)
		if err != nil {
			v.Failf("fail", "should create server monitoring with email notification by default: server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
		}
		err = testAction(target, "--host", "testbed.savac.s0csec1.org", "--path", "/protected/resource", "--status", "401", test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+k, "email")
		if err != nil {
			v.Failf("fail", "should create server monitoring with email notification: server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
		}
		err = testAction(target, "--host", "testbed.savac.s0csec1.org", test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+k, "webhook")
		if err == nil {
			v.Failf("fail", "should fail without arguments after `webhook`")
		}
		err = testAction(actionGenerator.GenerateServerMonitoringPingAction, "--port", "2222", "--host", "testbed.savac.s0csec1.org", test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+k, "webhook", sampleWebhookUrl)
		if err != nil {
			v.Failf("fail", "should create server monitoring with webhook notification for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
		}
		err = testAction(actionGenerator.GenerateServerMonitoringPingAction, "--host", "testbed.savac.s0csec1.org", test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+k, "webhook", sampleWebhookUrl, "test-team", "test-channel")
		if err != nil {
			v.Failf("fail", "should create server monitoring with webhook notification for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
		}
		if k == "https" {
			err = testAction(actionGenerator.GenerateServerMonitoringPingAction, "--sni", "--host", "testbed.savac.s0csec1.org", test_parameters.SampleRegisteredServerIdString, "test-monitoring-"+k, "webhook", sampleWebhookUrl, "test-team", "test-channel")
			if err != nil {
				v.Failf("fail", "should create server monitoring with webhook notification for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
			}
		}
	}

	log.Println("waiting for server monitoring setup to be completed")
	if v.isTestAcc {
		time.Sleep(180 * time.Second)
	}

	err := testAction(actionGenerator.GenerateServerMonitoringDeleteAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should delete all server monitoring for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}
}

func (v *VpsActionSuite) TestGenerateServerMonitoringInfoAction() {
	err := testAction(actionGenerator.GenerateServerMonitoringInfoAction)
	if err == nil {
		v.Failf("fail", "should fail without arguments")
	}
	err = testAction(actionGenerator.GenerateServerMonitoringInfoAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should list all monitoring for server %s: %v", test_parameters.SampleRegisteredServerIdString, err)
	}
	err = testAction(actionGenerator.GenerateServerMonitoringInfoAction, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should list all monitoring for server %s: %v", test_parameters.SampleRegisteredServerHostname, err)
	}
	if v.isTestAcc {
		time.Sleep(5 * time.Second)
	}
}
