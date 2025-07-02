package vps

import (
	"github.com/g1eng/savac/testutil/test_parameters"
	"time"
)

func (v *SavaClientSuite) TestSavaClient_ListRoles() {
	rls, err := v.apikeyCli.ListRoles()
	if err != nil {
		v.Fail("%v", err)
	}
	if rls == nil {
		v.Fail("ListRoles returned nil")
	}
}

func (v *SavaClientSuite) TestSavaClient_RoleOperations_FaultServer() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	rls, err := v.faultCli.ListRoles()
	if err == nil {
		v.Fail("ListRoles should have returned an error")
	}
	if len(rls) > 0 {
		v.Fail("ListRoles should return nil (actually: %v)", rls)
	}
	_, err = v.faultCli.GetRole(test_parameters.SampleRegisteredRoleId)
	if err == nil {
		v.Fail("GetRole should have returned an error")
	}
	dummyRole := test_parameters.DummyAnotherRole
	dummyRole.Id = 0
	_, err = v.faultCli.CreateRole(&dummyRole)
	if err == nil {
		v.Fail("CreateRole should have returned an error")
	}
	err = v.faultCli.DeleteRole(test_parameters.SampleRegisteredRoleId)
	if err == nil {
		v.Fail("DeleteRole should have returned an error")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetRole() {
	rls, err := v.apikeyCli.GetRole(test_parameters.SampleRegisteredRoleId)
	if err != nil {
		v.Fail("%v", err)
	}
	if rls.Id != test_parameters.SampleRegisteredRoleId {
		v.Fail("GetRole returned wrong role id")
	} else if rls.GetName() != test_parameters.SampleRegisteredRoleName {
		v.Fail("GetRole returned wrong role name")
	}

	rls, err = v.apikeyCli.GetRoleByName(test_parameters.SampleRegisteredRoleName)
	if err != nil {
		v.Fail("%v", err)
	}
	if rls.Id != test_parameters.SampleRegisteredRoleId {
		v.Fail("GetRole returned wrong role id")
	} else if rls.GetName() != test_parameters.SampleRegisteredRoleName {
		v.Fail("GetRole returned wrong role name")
	}
}

func (v *SavaClientSuite) TestSavaClientScenario_CreateDeleteRole() {
	dummyRole := test_parameters.DummyAnotherRole
	dummyRole.Id = 0
	role, err := v.apikeyCli.CreateRole(&dummyRole)
	if err != nil {
		v.Fail("%v", err)
	}
	err = v.apikeyCli.DeleteRole(role.Id)
	if err != nil {
		v.Fail("%v", err)
	}
}
