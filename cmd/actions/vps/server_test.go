package vps_actions

import (
	"fmt"
	"time"

	"github.com/g1eng/savac/testutil/test_parameters"

	"github.com/g1eng/savac/pkg/core"
)

func (v *VpsActionSuite) TestGenerateServerListAction() {
	var err error
	for i := 0; i <= core.OutputTypeText; i++ {
		actionGenerator.OutputType = i
		err = testAction(actionGenerator.GenerateServerListAction)
		if err != nil {
			v.Failf("fail", "%v", err)
		}
	}
	actionGenerator.OutputType = core.OutputTypeText
	err = testAction(actionGenerator.GenerateServerListAction, "-l")
	if err != nil {
		v.Failf("fail", "%v", err)
	}
}

func (v *VpsActionSuite) TestGenerateServerTagAction() {
	if v.isTestAcc {
		time.Sleep(60 * time.Second)
	}
	err := testAction(actionGenerator.GenerateServerTagAction)
	if err == nil {
		v.Failf("fail", "nonexistent pattern should return error")
	}
	err = testAction(actionGenerator.GenerateServerTagAction, test_parameters.SampleRegisteredServerIdString, test_parameters.SampleTagKey, test_parameters.SampleTagValue)
	if err != nil {
		v.Failf("fail", "should set the server tag: %v", err)
	}

	for i := 0; i <= core.OutputTypeText; i++ {
		actionGenerator.OutputType = i
		err = testAction(actionGenerator.GenerateServerTagAction, test_parameters.SampleRegisteredServerIdString, test_parameters.SampleTagKey)
		if err != nil {
			v.Failf("fail", "should get the server tag: %v", err)
		}
	}
	err = testAction(actionGenerator.GenerateServerTagAction, test_parameters.SampleRegisteredServerHostname, test_parameters.SampleTagKey)
	if err != nil {
		v.Failf("fail", "should get the server tag: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerTagAction, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should get all server tag: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateServerHostnameAction() {
	err := testAction(actionGenerator.GenerateServerHostnameAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should get the server hostname: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerHostnameAction, test_parameters.SampleRegisteredServerHostname, test_parameters.SampleRegisteredServerHostname+"-modified")
	if err != nil {
		v.Failf("fail", "should set the server hostname: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerHostnameAction, test_parameters.SampleRegisteredServerIdString, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should set the server hostname: %v", err)
	}
	//time.Sleep(10 * time.Second)
}

func (v *VpsActionSuite) TestGenerateServerInfoAction() {
	for i := 0; i <= core.OutputTypeText; i++ {
		actionGenerator.OutputType = i
		err := testAction(actionGenerator.GenerateServerInfoAction, test_parameters.SampleRegisteredServerIdString)
		if err != nil {
			v.Failf("fail", "%v", err)
		}
		err = testAction(actionGenerator.GenerateServerInfoAction, test_parameters.SampleRegisteredServerHostname)
		if err != nil {
			v.Failf("fail", "%v", err)
		}
	}
}

func (v *VpsActionSuite) TestGenerateServerInfoAction_With_Invalid_Pattern() {
	for i := 0; i <= core.OutputTypeText; i++ {
		actionGenerator.OutputType = i
		err := testAction(actionGenerator.GenerateServerInfoAction, fmt.Sprintf("%d", test_parameters.SampleUnregisteredServerId))
		if err == nil {
			v.Failf("fail", "nonexistent pattern should return error")
		}
	}
}

func (v *VpsActionSuite) TestScenario_GenerateServerStartRebootStopAction() {
	if !v.isTestAcc {
		v.T().SkipNow()
	}
	err := testAction(actionGenerator.GenerateServerStartAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should start the server: %v", err)
	}
	if v.isTestAcc {
		time.Sleep(20 * time.Second)
	}
	err = testAction(actionGenerator.GenerateServerStartAction, test_parameters.SampleRegisteredServerIdString)
	if err == nil {
		v.Failf("fail", "should fail for started server")
	}
	err = testAction(actionGenerator.GenerateServerRebootAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should reboot the server: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerRebootAction)
	if err == nil {
		v.Failf("fail", "should fail to reboot the server without arguments")
	}
	if v.isTestAcc {
		time.Sleep(20 * time.Second)
	}
	err = testAction(actionGenerator.GenerateServerStopAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should stop the server: %v", err)
	}

	err = testAction(actionGenerator.GenerateServerStartAction)
	if err == nil {
		v.Failf("fail", "should fail to start the server without arguments")
	}
	err = testAction(actionGenerator.GenerateServerStopAction)
	if err == nil {
		v.Failf("fail", "should fail to stop the server without arguments")
	}

	if v.isTestAcc {
		time.Sleep(5 * time.Second)
	}
	err = testAction(actionGenerator.GenerateServerStopAction, test_parameters.SampleRegisteredServerIdString)
	if err == nil {
		v.Failf("fail", "should fail to stop the stopped server")
	}
}

func (v *VpsActionSuite) TestGenerateServerDescriptionAction() {
	err := testAction(actionGenerator.GenerateServerDescriptionAction)
	if err == nil {
		v.Failf("fail", "nonexistent pattern should return error")
	}
	for i := 0; i <= core.OutputTypeText; i++ {
		actionGenerator.OutputType = i
		err = testAction(actionGenerator.GenerateServerDescriptionAction, test_parameters.SampleRegisteredServerIdString)
		if err != nil {
			v.Failf("fail", "should get the server description: %v", err)
		}
		err = testAction(actionGenerator.GenerateServerDescriptionAction, test_parameters.SampleRegisteredServerHostname)
		if err != nil {
			v.Failf("fail", "should get the server description: %v", err)
		}
		//time.Sleep(5 * time.Second)
	}
	err = testAction(actionGenerator.GenerateServerDescriptionAction, test_parameters.SampleRegisteredServerIdString, "test:neko")
	if err != nil {
		v.Failf("fail", "should set the server description: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateServerInterfaceAction() {
	if v.isTestAcc {
		time.Sleep(10 * time.Second)
	}
	for i := 0; i <= core.OutputTypeText; i++ {
		actionGenerator.OutputType = i
		err := testAction(actionGenerator.GenerateServerInterfaceAction, test_parameters.SampleRegisteredServerIdString)
		if err != nil {
			v.Failf("fail", "should get the server interfaces: %v", err)
		}
		err = testAction(actionGenerator.GenerateServerInterfaceAction, test_parameters.SampleRegisteredServerHostname)
		if err != nil {
			v.Failf("fail", "%v", err)
		}
		//time.Sleep(3 * time.Second)
	}
}

func (v *VpsActionSuite) TestScenario_GenerateServerStartRebootStopAction_With_Name() {
	if !v.isTestAcc {
		v.T().SkipNow()
	}
	err := testAction(actionGenerator.GenerateServerStartAction, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should start the server: %v", err)
	}
	if v.isTestAcc {
		time.Sleep(20 * time.Second)
	}
	err = testAction(actionGenerator.GenerateServerRebootAction, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should reboot the server: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerDescriptionAction, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should show the server info: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerListAction)
	if err != nil {
		v.Failf("fail", "should list the server: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerListAction, "-l")
	if err != nil {
		v.Failf("fail", "should list the server in detail: %v", err)
	}

	if v.isTestAcc {
		time.Sleep(20 * time.Second)
	}
	err = testAction(actionGenerator.GenerateServerStopAction, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should stop the server: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateServerPtrAction() {
	if v.isTestAcc {
		time.Sleep(30 * time.Second)
	}
	err := testAction(actionGenerator.GenerateServerPtrAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should get the server ptr: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerPtrAction, "--ipv6", test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should get the server ipv6 ptr: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerPtrAction, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("fail", "should set the server ptr with hostname: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerPtrAction, test_parameters.SampleRegisteredServerIdString, sampleResolvableServerARecord)
	if err != nil {
		v.Failf("fail", "should set the server ptr: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerPtrAction, test_parameters.SampleRegisteredServerHostname, sampleResolvableServerARecord)
	if err != nil {
		v.Failf("fail", "should set the server ptr: %v", err)
	}
	err = testAction(actionGenerator.GenerateServerPtrAction, "--ipv6", test_parameters.SampleRegisteredServerIdString, sampleResolvableServerARecord)
	if err != nil {
		v.Failf("fail", "should set the server ptr: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateServerConnectDisconnectAction() {
	err := testAction(actionGenerator.GenerateServerConnectAction, test_parameters.SampleRegisteredServerInterfaceIdText, test_parameters.SampleRegisteredSwitchIdText)
	if err != nil {
		v.Failf("fail", "should connect the server interface to the switch: %v", err)
	}
	//time.Sleep(1 * time.Second)
	err = testAction(actionGenerator.GenerateServerConnectAction, "--disconnect", test_parameters.SampleRegisteredServerInterfaceIdText)
	if err != nil {
		v.Failf("fail", "should disconnect the server interface from the switch: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateServerStartRebootStopAction() {
	err := testAction(actionGenerator.GenerateServerStartAction, test_parameters.SampleRegisteredServerIdString)
	if err != nil {
		v.Failf("fail", "should start the server: %v", err)
	}
	for _, args := range [][]string{
		{test_parameters.DefaultServer.Name},
		{test_parameters.DefaultServer.Name, "--regex"},
		{test_parameters.DefaultServer.Name, "--search"},
		{test_parameters.SampleRegisteredServerIdString},
	} {
		if v.isTestAcc {
			time.Sleep(20 * time.Second)
		}
		err = testAction(
			actionGenerator.GenerateServerStartAction,
			args...)
		if err == nil {
			v.Failf("fail", "should fail for started server")
		}
		if v.isTestAcc {
			time.Sleep(20 * time.Second)
		}
		err = testAction(actionGenerator.GenerateServerRebootAction, args...)
		if err != nil {
			v.Failf("fail", "should reboot the server: %v", err)
		}
		if v.isTestAcc {
			time.Sleep(20 * time.Second)
		}
		err = testAction(actionGenerator.GenerateServerStopAction, args...)
		if err != nil {
			v.Failf("fail", "should stop the server: %v", err)
		}
	}
	if v.isTestAcc {
		time.Sleep(5 * time.Second)
		err = testAction(actionGenerator.GenerateServerStopAction, test_parameters.SampleRegisteredServerIdString)
		if err == nil {
			v.Failf("fail", "should fail to stop the stopped server")
		}
	}
}

//func (v *VpsActionSuite) TestGenerateServerStartRebootStopAction_With_Name() {
//	targetServerName := test_parameters.SampleRegisteredServerHostname
//	if !v.isTestAcc {
//		targetServerName = test_parameters.SpecialServer.Name
//	}
//	err := testAction(actionGenerator.GenerateServerStartAction(), targetServerName)
//	if err != nil {
//		v.Failf("fail", "should start the server: %v", err)
//	}
//	if v.isTestAcc {
//		time.Sleep(20 * time.Second)
//	}
//	err = testAction(actionGenerator.GenerateServerRebootAction(), targetServerName)
//	if err != nil {
//		v.Failf("fail", "should reboot the server: %v", err)
//	}
//	err = testAction(actionGenerator.GenerateServerDescriptionAction(), targetServerName)
//	if err != nil {
//		v.Failf("fail", "should show the server info: %v", err)
//	}
//	err = testAction(actionGenerator.GenerateServerListAction())
//	if err != nil {
//		v.Failf("fail", "should list the server: %v", err)
//	}
//	err = testAction(actionGenerator.GenerateServerListAction(), "-l")
//	if err != nil {
//		v.Failf("fail", "should list the server in detail: %v", err)
//	}
//	if v.isTestAcc {
//		time.Sleep(20 * time.Second)
//	}
//	err = testAction(actionGenerator.GenerateServerStopAction(), targetServerName)
//	if err != nil {
//		v.Failf("fail", "should stop the server: %v", err)
//	}
//}
