package vps

import (
	"github.com/g1eng/savac/testutil/test_parameters"
)

var (
	dummySwitchName = "okneko"
)

func (v *SavaClientSuite) TestScenario_Create_List_And_Delete_Switch() {
	err := v.swCli.CreateSwitch(dummySwitchName, "", "os3")
	if err != nil {
		v.Fail("failed to create dummy switch: %v", err)
	}
	sw, err := v.swCli.GetSwitchList()
	if err != nil {
		v.Fail("failed to get switch list: %v", err)
	}
	for _, s := range sw {
		if s.Name == dummySwitchName {
			err = v.swCli.DeleteSwitch(s.Id)
			if err != nil {
				v.Fail("failed to delete dummy switch: %v", err)
			} else {
				return
			}
		}
	}
	v.Fail("invalid condition")
}

func (v *SavaClientSuite) TestSavaPutSwitch() {
	err := v.swCli.PutSwitchName(test_parameters.SampleRegisteredSwitchId, dummySwitchName)
	if err != nil {
		v.Fail("failed to put dummy switch: %v", err)
	}
	err = v.swCli.PutSwitchName(test_parameters.SampleRegisteredSwitchId, test_parameters.SampleRegisteredSwitchName)
	if err != nil {
		v.Fail("failed to put dummy switch to the original one: %v", err)
	}
	newDesc := "sample-description"
	err = v.swCli.PutSwitchDescription(test_parameters.SampleRegisteredSwitchId, newDesc)
	if err != nil {
		v.Fail("failed to put dummy switch to the original one: %v", err)
	}
	//sw, err := v.swCli.GetSwitchById(test_parameters.SampleRegisteredSwitchId)
	//if err != nil {
	//	v.Fail("failed to get switch: %v", err)
	//}
	//if sw.Description != newDesc {
	//	v.Fail("failed to update switch description")
	//}
}

func (v *SavaClientSuite) TestSavaClient_GetSwitchById() {
	sw, err := v.swCli.GetSwitchById(test_parameters.SampleRegisteredSwitchId)
	if err != nil {
		v.Fail("failed to get a registered switch: %v", err)
	}
	if sw.Name != test_parameters.SampleRegisteredSwitchName {
		v.Fail("failed to get a registered switch: expected %v, got %v", test_parameters.SampleRegisteredSwitchName, sw.Name)
	}
	if sw.Id != test_parameters.SampleRegisteredSwitchId {
		v.Fail("failed to get a registered switch: expected %v, got %v", test_parameters.SampleRegisteredSwitchId, sw.Id)
	}
}

func (v *SavaClientSuite) TestSavaClient_PutSwitchName() {
	err := v.swCli.PutSwitchName(test_parameters.SampleRegisteredSwitchId, dummySwitchName)
	if err != nil {
		v.Fail("failed to put dummy switch name: %v", err)
	}
	err = v.swCli.PutSwitchName(test_parameters.SampleRegisteredSwitchId, test_parameters.SampleRegisteredSwitchName)
	if err != nil {
		v.Fail("failed to put dummy switch name to original one: %v", err)
	}
}

func (v *SavaClientSuite) TestScenario_Create_DeleteSwitch() {
	err := v.swCli.CreateSwitch(dummySwitchName, "", "os3")
	if err != nil {
		v.Fail("failed to create dummy switch: %v", err)
	}
	sw, err := v.swCli.GetSwitchList()
	if err != nil {
		v.Fail("failed to get switch list: %v", err)
	}
	for _, s := range sw {
		if s.Name == dummySwitchName {
			err = v.swCli.DeleteSwitch(s.Id)
			if err != nil {
				v.Fail("failed to delete dummy switch: %v", err)
			}
			if v.isTestAcc {
				err = v.swCli.DeleteSwitch(s.Id)
				if err == nil {
					v.Fail("delete should fail for a deleted switch")
				}
			}
			return
		}
	}
	v.Fail("failed to find a dummy switch")
}

func (v *SavaClientSuite) TestSavaClient_SwitchRequest_For_Fault_Server() {
	_, err := v.faultCli.GetSwitchById(test_parameters.SampleRegisteredSwitchId)
	if err == nil {
		v.Fail("an exception should be raised for server fault")
	}
	_, err = v.faultCli.GetSwitchList()
	if err == nil {
		v.Fail("an exception should be raised for server fault")
	}
	_, err = v.faultCli.GetSwitchById(test_parameters.SampleRegisteredSwitchId)
	if err == nil {
		v.Fail("an exception should be raised for server fault")
	}
	err = v.faultCli.CreateSwitch(dummySwitchName, "ok", "os3")
	if err == nil {
		v.Fail("an exception should be raised for server fault")
	}
	err = v.faultCli.DeleteSwitch(test_parameters.SampleRegisteredServerId)
	if err == nil {
		v.Fail("an exception should be raised for server fault")
	}
}
