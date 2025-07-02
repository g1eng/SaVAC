package vps

import (
	"fmt"
	"time"
)

func (v *SavaClientSuite) TestSavaClient_ListPermissions() {
	perm, err := v.apikeyCli.ListPermissions()
	if err != nil {
		v.Fail("%v", err)
	}
	if len(perm) == 0 {
		v.Fail("no DummyPermissions found")
	}
	fmt.Printf("%v", perm)
}

func (v *SavaClientSuite) TestSavaClient_ListPermissions_With_FaultResponse() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	rls, err := v.faultCli.ListPermissions()
	if err == nil {
		v.Fail("ListPermissions should return an error")
	}
	if rls != nil {
		v.Fail("ListPermissions should return nil, returned %#v", rls)
	}
}
