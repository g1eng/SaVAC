package vps_actions

func (v *VpsActionSuite) TestGenerateDiscAction() {
	err := testAction(actionGenerator.GenerateDiscAction)
	if err != nil {
		v.Failf("fail", "should not raise error: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateZoneAction() {
	err := testAction(actionGenerator.GenerateZoneAction)
	if err != nil {
		v.Failf("fail", "should not raise error: %v", err)
	}
}
