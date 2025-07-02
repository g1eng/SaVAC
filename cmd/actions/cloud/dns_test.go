package cloud_actions

import (
	"log"
	"time"

	"github.com/g1eng/savac/pkg/core"
)

var dnsTestTable TestTable = []TestTableElement{
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsARecordName, "A", sampleUnregisteredDnsARecordValue}, //add
		"",
		"add",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsARecordName}, //del
		"delete command should delete rr with perfect matching of the name of rrdata",
		"del",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsARecordName, "A", sampleUnregisteredDnsARecordValue}, //add
		"",
		"add",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsARecordName, "AAAA", "::1"}, //add
		"",
		"add",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsCNAMERecordName, "CNAME", sampleUnregisteredDnsCNAMERecordValue}, //add
		"",
		"add",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsARecordName, "MX", sampleUnregisteredDnsMXRecordValue}, //add
		"",
		"add",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsARecordName, "TXT", "v=spf1 -all"}, //add
		"",
		"add",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, "CNAME"},
		"delete command should delete rr with rrtype",
		"del",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, "MX"},
		"delete command should delete rr with rrtype",
		"del",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, "A"},
		"delete command should delete rr with rrtype",
		"del",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, "AAAA"},
		"delete command should delete rr with rrtype",
		"del",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, "TXT"},
		"delete command should delete rr with rrtype",
		"del",
	},
	//{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsARecordName, "A", sampleUnregisteredDnsARecordValue},
	//{"--id", sampleRegisteredDnsApplianceId, "--regex", sampleUnregisteredDnsARecordName},
}

var dnsFaultTestTable = []TestTableElement{
	{
		[]string{sampleUnregisteredDnsARecordName, "A", sampleUnregisteredDnsARecordValue},
		"add command should fail without dns appliance id",
		"add",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, sampleUnregisteredDnsARecordName, "B", sampleUnregisteredDnsARecordValue},
		"add command should fail for invalid rrtype",
		"add",
	},
	{
		[]string{"A"},
		"delete command should fail without dns appliance id",
		"delete",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, "B"},
		"delete command should fail for invalid rrtype",
		"delete",
	},
	{
		[]string{"--id", sampleRegisteredDnsApplianceId, "--regex", "asa[]("},
		"delete command should fail for invalid regex pattern",
		"delete",
	},
}

func (s *CloudActionSuite) TestGenerateDnsApplianceListAction() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(s.Generator.GenerateDnsApplianceListAction, sampleRegisteredDnsApplianceId)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestGenerateDnsRecordListAction() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(s.Generator.GenerateDnsRecordListAction, sampleRegisteredDnsApplianceId)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateDnsRecordListAction)
	if err == nil {
		s.Fail("should have failed due to lacking appliance id")
	}
	err = testAction(s.Generator.GenerateDnsRecordListAction, "--id", sampleRegisteredDnsApplianceId)
	if err == nil {
		s.Fail("should have failed due to lacking appliance id (--id should not be used)")
	}
}

func (s *CloudActionSuite) TestGenerateDnsRecordAdd_And_Delete_Action() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	for _, tc := range dnsTestTable {
		if tc.Type == "add" {
			err = testAction(s.Generator.GenerateDnsRecordAddAction, tc.Args...)
		} else {
			err = testAction(s.Generator.GenerateDnsRecordDeleteAction, tc.Args...)
		}
		if err != nil {
			s.Failf("operation failure", "op %s\n%s\ndata: %v\nerr: %s", tc.Type, tc.ErrorMessage, tc, err.Error())
		}
		if isTestAcc {
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func (s *CloudActionSuite) TestGenerateDnsRecord_For_Invalid_Add_And_Delete_Action() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	for i, tc := range dnsFaultTestTable {
		log.Printf("test index: %d\n", i)
		err = testAction(s.Generator.GenerateDnsRecordAddAction, tc.Args...)
		if err == nil {
			s.Fail(tc.ErrorMessage)
		}
		if isTestAcc {
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *CloudActionSuite) TestGenerateDnsReadAction() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(s.Generator.GenerateDnsApplianceReadAction, sampleRegisteredDnsApplianceId)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateDnsApplianceReadAction, sampleRegisteredDnsZone)
	if err != nil {
		s.Fail(err.Error())
	}

	//with yaml
	s.Generator.OutputType = core.OutputTypeYaml
	err = testAction(s.Generator.GenerateDnsApplianceReadAction, sampleRegisteredDnsApplianceId)
	if err != nil {
		s.Fail(err.Error())
	}

	//export records in zonefile format
	s.Generator.OutputType = core.OutputTypeText
	err = testAction(s.Generator.GenerateDnsApplianceReadAction, sampleRegisteredDnsApplianceId)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestGenerateDnsRecordImport_Zonefile_Action() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err := testAction(s.Generator.GenerateDnsApplianceRecordImportAction, sampleRegisteredDnsApplianceId)
	if err == nil {
		s.Fail("should have failed due to lacking appliance id")
	}
	err = testAction(s.Generator.GenerateDnsApplianceRecordImportAction, "--file", "./fixtures/zonefile.txt", sampleRegisteredDnsApplianceId)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestGenerateDnsRecordImport_RespJson_Action() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err := testAction(s.Generator.GenerateDnsApplianceRecordImportAction, sampleRegisteredDnsApplianceId)
	if err == nil {
		s.Fail("should have failed due to lacking appliance id")
	}
	err = testAction(s.Generator.GenerateDnsApplianceRecordImportAction, "--file", "./fixtures/dnsresp.json", sampleRegisteredDnsApplianceId)
	if err != nil {
		s.Fail(err.Error())
	}
}
