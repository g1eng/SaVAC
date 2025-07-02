package vps

func (v *SavaClientSuite) TestSavaClient_Misc_ListZones() {
	zones, err := v.miscCli.GetAllZone()
	if err != nil {
		v.Fail("%v", err)
	}
	if len(zones) == 0 {
		v.Fail("no zones returned")
	}
}

func (v *SavaClientSuite) TestSavaClient_Misc_ListCDROMs() {
	discs, err := v.miscCli.ListCDROMs()
	if err != nil {
		v.Fail("%v", err)
	}
	if len(discs) == 0 {
		v.Fail("no cdroms returned")
	}
}

func (v *SavaClientSuite) TestSavaClient_Misc_FatalResponse() {
	_, err := v.faultCli.GetAllZone()
	if err == nil {
		v.Fail("expected an error")
	}
	_, err = v.faultCli.ListCDROMs()
	if err == nil {
		v.Fail("expected an error")
	}
}
