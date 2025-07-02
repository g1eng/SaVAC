package vps

import (
	"fmt"
	"time"

	"github.com/g1eng/savac/testutil/test_parameters"
)

func (v *SavaClientSuite) TestSavaClient_GetAllServers() {
	if v.isTestAcc {
		time.Sleep(time.Minute)
	}
	sv, err := v.serverCli.GetAllServers()
	if err != nil {
		v.Failf("server list should be given: %v", "server list should be given: %v", err)
	}
	if len(sv) == 0 {
		v.Failf("no server returned", "no server returned")
	}
}

//func (v *SavaClientSuite) TestSavaClient_GetAllServersWithPaging() {
//	sv, err := v.mockCli.GetAllServers()
//	if err != nil {
//		v.Failf("server list should be given: %v", "server list should be given: %v", err)
//	}
//	if len(sv) != len(test_parameters.DummyServers) {
//		v.Failf("server list should have returned %d DummyServers, but returned %d", "server list should have returned %d DummyServers, but returned %d", len(test_parameters.DummyServers), len(sv))
//	}
//}

func (v *SavaClientSuite) TestSavaClient_ServerRequest_With_Fault_Response() {
	_, err := v.faultCli.GetAllServers()
	if err == nil {
		v.Failf("exception shoould be raised for fault status such as 429", "exception shoould be raised for fault status such as 429")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerById() {
	sv, err := v.serverCli.GetServerById(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("server should be given: %v", "server should be given: %v", err)
	}
	if sv == nil {
		v.Failf("server should be not nil", "server should be not nil")
	}
}
func (v *SavaClientSuite) TestSavaClient_GetServerById_For_FaultResponse() {
	sv, err := v.faultCli.GetServerById(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("exception shoould be raised for fault status such as 429", "exception shoould be raised for fault status such as 429")
	}
	if sv != nil {
		v.Failf("server should be nil", "server should be nil")
	}
}
func (v *SavaClientSuite) TestSavaClient_GetServerById_For_Nonexistent_Server() {
	sv, err := v.serverCli.GetServerById(test_parameters.SampleUnregisteredServerId)
	if err == nil {
		v.Failf("exception should be raised", "exception should be raised")
	}
	if sv != nil {
		v.Failf("server should be nil", "server should be nil")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerIdsByNamePattern_Name() {
	svIds, err := v.serverCli.GetServerIdsByNamePattern(test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("server list should be given: %v", "server list should be given: %v", err)
	}
	if len(svIds) == 0 {
		v.Failf("no server returned", "no server returned")
	}
	if svIds[0] != test_parameters.SampleRegisteredServerId {
		v.Failf("server id should be %d, got %d", "server id should be %d, got %d", test_parameters.SampleRegisteredServerId, svIds[0])
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerIdsByNamePattern_Name_Unregistered() {
	svIds, err := v.serverCli.GetServerIdsByNamePattern(test_parameters.SampleUnregisteredServerName)
	if err == nil {
		v.Failf("exception should be raised", "exception should be raised")
	}
	if len(svIds) != 0 {
		v.Failf("no server should be returned", "no server should be returned")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerIdsByNamePattern_For_Id_Incompatible() {
	svIds, err := v.serverCli.GetServerIdsByNamePattern(fmt.Sprintf("%d", int(test_parameters.SampleRegisteredServerId)))
	if err == nil {
		v.Failf("server list should not be given for id", "server list should not be given for id")
	}
	if len(svIds) != 0 {
		v.Failf("no server returned", "no server returned")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerIdsByNamePattern_Id_Nonexistent() {
	if !v.isTestAcc {
		v.T().SkipNow()
	}
	svIds, err := v.serverCli.GetServerIdsByNamePattern(fmt.Sprintf("%d", int(test_parameters.SampleUnregisteredServerId)))
	if err == nil {
		v.Failf("exception should be raised", "exception should be raised")
	}
	if len(svIds) != 0 {
		v.Failf("no server should be returned", "no server should be returned")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetFilteredServerList() {
	svs, err := v.serverCli.GetFilteredServerList(test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("server list should be given: %v", "server list should be given: %v", err)
	}
	if len(svs) == 0 {
		v.Failf("no server returned", "no server returned")
	}
	if svs[0].Id != test_parameters.SampleRegisteredServerId {
		v.Failf("server id should be %d, got %d", "server id should be %d, got %d", test_parameters.SampleRegisteredServerId, svs[0].Id)
	} else if svs[0].Name != test_parameters.SampleRegisteredServerHostname {
		v.Failf("server name should be %v, got %v", "server name should be %v, got %v", test_parameters.SampleRegisteredServerHostname, svs[0].Name)
	}
}

func (v *SavaClientSuite) TestSavaClient_GetFilteredServerList_For_FaultResponse() {
	svs, err := v.faultCli.GetFilteredServerList(test_parameters.DefaultServer.Name)
	if err == nil {
		v.Failf("exception should be raised", "exception should be raised")
	}
	if len(svs) != 0 {
		v.Failf("DummyServers should not be returned", "DummyServers should not be returned")
	}
	//if err.Error() != Fixture429Text {
	//	v.Failf("error response should be dumped directly: %v, actual: %v","error response should be dumped directly: %v, actual: %v", Fixture429Text, err.Error())
	//}
}

func (v *SavaClientSuite) TestScenario_Start_Reboot_Stop_Server_With_Id() {
	if !v.isTestAcc {
		v.T().SkipNow()
	}
	err := v.serverCli.StartServerWithId(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("server should be started: %v", "server should be started: %v", err)
	}
	err = v.serverCli.StartServerWithId(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("a started server should fail to be started: %v", "a started server should fail to be started: %v", err)
	}
	if v.isTestAcc {
		time.Sleep(20 * time.Second)
	}
	err = v.serverCli.ForceRebootServerWithId(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("server should be forced to reboot: %v", "server should be forced to reboot: %v", err)
	}
	if v.isTestAcc {
		time.Sleep(30 * time.Second)
	}
	err = v.serverCli.StopServerWithId(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("server should be stopped: %v", "server should be stopped: %v", err)
	}
	if v.isTestAcc {
		time.Sleep(15 * time.Second)
	}
	err = v.serverCli.StopServerWithId(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("a stopped server should fail to be stopped: %v", "a stopped server should fail to be stopped: %v", err)
	}
}

func (v *SavaClientSuite) TestSavaClient_GetFilteredServerList_For_Nonexistent_Server() {
	svs, err := v.serverCli.GetFilteredServerList(test_parameters.SampleUnregisteredServerName)
	if err == nil {
		v.Failf("exception should be raised", "exception should be raised")
	}
	if len(svs) != 0 {
		v.Failf("server should not be returned", "server should not be returned")
	}
}

func (v *SavaClientSuite) TestSavaClient_ShowServerDescriptionById() {
	p, err := v.serverCli.GetServerDescriptionById(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("server description should be given: %v", "server description should be given: %v", err)
	}
	if p == nil {
		v.Failf("server description should be not nil", "server description should be not nil")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetServerHostName() {
	h, err := v.serverCli.GetHostname(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("server hostname should be given: %v", "server hostname should be given: %v", err)
	}
	if h != test_parameters.SampleRegisteredServerHostname {
		v.Failf("server hostname should be %v, got %v", "server hostname should be %v, got %v", test_parameters.SampleRegisteredServerHostname, h)
	}
}

func (v *SavaClientSuite) TestSavaClient_SetServerHostName() {
	err := v.serverCli.SetHostName(test_parameters.SampleRegisteredServerId, "hoge")
	if err != nil {
		v.Failf("server hostname should be set: %v", "server hostname should be set: %v", err)
	}
	err = v.serverCli.SetHostName(test_parameters.SampleRegisteredServerId, test_parameters.SampleRegisteredServerHostname)
	if err != nil {
		v.Failf("server hostname should be set: %v", "server hostname should be set: %v", err)
	}
}

func (v *SavaClientSuite) TestSavaClient_SetServerDescription() {
	err := v.serverCli.SetServerDescription(test_parameters.SampleRegisteredServerId, "hoge")
	if err != nil {
		v.Failf("server description should be set: %v", "server description should be set: %v", err)
	}
}

func (v *SavaClientSuite) TestScenario_Add_And_ListServerTag() {
	newKey, newVal := "key2", "value2"
	err := v.serverCli.AddServerTag(test_parameters.SampleRegisteredServerId, newKey, newVal)
	if err != nil {
		v.Failf("server tag should be added: %v", "server tag should be added: %v", err)
	}
	tags, err := v.serverCli.ListServerTags(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("list of server tags should be given: %v", "list of server tags should be given: %v", err)
	}
	if len(tags) == 0 {
		v.Failf("server tag should be added", "server tag should be added")
	}
	tagVal, err := v.serverCli.GetServerTag(test_parameters.SampleRegisteredServerId, newKey)
	if err != nil {
		v.Failf("no server tag set", "server tag value should be given: %v", err)
	}
	if tagVal == nil {
		v.Failf("null server tag", "server tag value should not be nil")
	} else if *tagVal != newVal {
		v.Failf("invalid server tag", "server tag value should be %v, got %v", *tagVal, "val")
	}
}

func (v *SavaClientSuite) TestSavaClient_SetIpv4PtrWithServerId() {
	err := v.serverCli.SetIpv4PtrWithServerId(test_parameters.SampleRegisteredServerId, test_parameters.SampleResolvableServerARecord)
	if err != nil {
		v.Failf("server PTR record should be set: %v", "server PTR record should be set: %v", err)
	}
}

func (v *SavaClientSuite) TestSavaClient_GetIpv4PtrWithServerId() {
	ptr, err := v.serverCli.GetIpv4PtrWithServerId(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("server PTR record should be given: %v", "server PTR record should be given: %v", err)
	}
	if ptr == "" {
		v.Failf("server PTR should not be empty", "server PTR should not be empty")
	}
}

func (v *SavaClientSuite) TestSavaClient_SetIpv6PtrWithServerId() {
	err := v.serverCli.SetIpv6PtrWithServerId(test_parameters.SampleRegisteredServerId, test_parameters.SampleResolvableServerARecord)
	if err != nil {
		v.Failf("server PTR record should be set: %v", "server PTR record should be set: %v", err)
	}
}

func (v *SavaClientSuite) TestSavaClient_GetIpv6PtrWithServerId() {
	ptr, err := v.serverCli.GetIpv6PtrWithServerId(test_parameters.SampleRegisteredServerId)
	if err != nil {
		v.Failf("server PTR record should be given: %v", "server PTR record should be given: %v", err)
	}
	if ptr == "" {
		v.Failf("server PTR should not be empty", "server PTR should not be empty")
	}
}

func (v *SavaClientSuite) TestSavaClient_MiscServer_FaultResponse() {
	_, err := v.faultCli.GetHostname(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
	_, err = v.faultCli.GetIpv4PtrWithServerId(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
	_, err = v.faultCli.GetIpv6PtrWithServerId(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
	_, err = v.faultCli.GetServerById(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
	_, err = v.faultCli.GetServerDescriptionById(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
	_, err = v.faultCli.ListServerTags(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
	_, err = v.faultCli.GetFilteredServerList(test_parameters.SampleRegisteredServerHostname)
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
	err = v.faultCli.SetServerDescription(test_parameters.SampleRegisteredServerId, "")
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
	err = v.faultCli.AddServerTag(test_parameters.SampleRegisteredServerId, "key", "val")
	if err == nil {
		v.Failf("an exception should be raised for server fault", "an exception should be raised for server fault")
	}
}
