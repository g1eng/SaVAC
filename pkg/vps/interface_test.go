package vps

import (
	"time"

	"github.com/g1eng/savac/testutil/test_parameters"
)

func (v *SavaClientSuite) TestSavaClient_GetAllServerInterface() {
	if v.isTestAcc {
		time.Sleep(time.Minute)
	}
	ifs, err := v.serverCli.GetAllServerInterface()
	if ifs == nil {
		v.Fail("DummyInterfaces should be listed")
	}
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_GetServerInterfaces() {
	ifs, err := v.serverCli.GetServerInterfaces(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Fail("%v", err)
	}
	if len(ifs) == 0 {
		v.Fail("server DummyInterfaces should be more then one for the VPS: real %d", len(ifs))
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerInterfaces_With_Nonexistent_ServerId() {
	ifs, err := v.serverCli.GetServerInterfaces(test_parameters.SampleUnregisteredServerId)
	if err == nil {
		v.Fail("an exception should be raised")
	}
	if len(ifs) != 0 {
		v.Fail("server DummyInterfaces should not be listed")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerInterfacesWithPattern() {
	ifs, err := v.serverCli.GetServerInterfacesWithPattern(test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Fail("%v", err)
	}
	if len(ifs) == 0 {
		v.Fail("server DummyInterfaces should be more than one for the VPS: real %d", len(ifs))
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerInterfacesWithPattern_With_Nonexistent_ServerName() {
	ifs, err := v.serverCli.GetServerInterfacesWithPattern(test_parameters.SampleUnregisteredServerName)
	if err == nil {
		v.Fail("an exception should be raised")
	}
	if len(ifs) != 0 {
		v.Fail("server DummyInterfaces should not be listed")
	}
}

func (v *SavaClientSuite) TestSavaClient_SetServerInterfaceConnection_ConnectTo_And_DisConnectFrom_Switch_And_ConnectTo_SharedNetwork_Finally() {
	err := v.serverCli.SetServerInterfaceConnection(test_parameters.SampleRegisteredServerInterfaceId, test_parameters.SampleRegisteredSwitchId)
	if err != nil {
		v.Failf("failed to connect to the switch", "switch %d: %v", test_parameters.SampleRegisteredSwitchId, err)
	}
	err = v.serverCli.SetServerInterfaceConnection(test_parameters.SampleRegisteredServerInterfaceId, 0)
	if err != nil {
		v.Failf("failed to disconnect the interface", "%v", err)
	}
	err = v.serverCli.SetServerInterfaceConnection(test_parameters.SampleRegisteredServerInterfaceId, 1)
	if err != nil {
		v.Failf("failed to connect to the internet", "%v", err)
	}
}

func (v *SavaClientSuite) TestSavaClient_SetServerInterfaceConnection_ConnectTo_Unknown_Switch_ShouldFail() {
	err := v.serverCli.SetServerInterfaceConnection(test_parameters.SampleRegisteredServerInterfaceId, test_parameters.SampleUnregisteredSwitchId)
	if err == nil {
		v.Fail("an exception should be raised")
	}
}

func (v *SavaClientSuite) TestSavacServerClient_TestInterfaceCommand_With_FaultServer() {
	_, err := v.faultCli.GetAllServerInterface()
	if err == nil {
		v.Fail("Should have gotten an error")
	}
	_, err = v.faultCli.GetServerInterfaces(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Fail("Should have gotten an error")
	}
	_, err = v.faultCli.GetServerInterfacesWithPattern(test_parameters.SampleRegisteredServerHostname)
	if err == nil {
		v.Fail("Should have gotten an error")
	}
	err = v.faultCli.SetServerInterfaceConnection(test_parameters.SampleRegisteredServerInterfaceId, test_parameters.SampleRegisteredSwitchId)
	if err == nil {
		v.Fail("Should have gotten an error")
	}
}
