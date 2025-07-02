package vps_actions

func (v *VpsActionSuite) TestGeneratePermissionListAction() {
	err := testAction(actionGenerator.GeneratePermissionListAction, "-l")
	if err != nil {
		v.Failf("fail", "%v", err)
	}
}
