package cloud_actions

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/search"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/urfave/cli/v3"
	"github.com/wpalmer/gozone"
)

func getDNSApplianceIdByString(s iaas.DNSAPI, arg string) (types.ID, error) {
	id, err := strconv.Atoi(arg)
	if err == nil {
		return types.ID(id), nil
	}
	res, err := s.Find(context.Background(), &iaas.FindCondition{
		Count: 100,
		From:  0,
		Sort:  nil,
		Filter: search.Filter{
			search.FilterKey{
				Field: "Name",
				Op:    "=",
			}: arg,
		},
	})
	if err != nil {
		return types.ID(-1), err
	} else if res.Count != 1 {
		return types.ID(-1), errors.New("duplicated names exist")
	}
	return res.DNS[0].ID, nil
}

func (g *CloudActionGenerator) GenerateDnsApplianceReadAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewDNSOp(*g.ApiClient.Caller)
	if cmd.NArg() != 1 {
		return fmt.Errorf("an appliance must be specified")
	}
	id, err := getDNSApplianceIdByString(s, cmd.Args().First())
	if err != nil {
		return err
	}
	res, err := s.Read(context.Background(), id)
	if err != nil {
		return err
	}
	if ctx.Value("output-type") == "text" {
		print(helper.FormatAsZoneFile(res))
		return nil
	}

	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(res)
	case core.OutputTypeYaml:
		return helper.PrintYaml(res)
	default:
		print(helper.FormatAsZoneFile(res))
		return nil
	}
}

func (g *CloudActionGenerator) GenerateDnsApplianceCreateAction(_ context.Context, cmd *cli.Command) error {
	s := iaas.NewDNSOp(*g.ApiClient.Caller)
	res, err := s.Create(context.Background(), &iaas.DNSCreateRequest{
		Name:        cmd.String("name"),
		Records:     nil,
		Description: cmd.String("description"),
		Tags:        cmd.StringSlice("tags"),
		IconID:      0,
	})
	if err != nil {
		return err
	}
	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(res)
	case core.OutputTypeYaml:
		return helper.PrintYaml(res)
	default:
		println(res.ID.String())
		return nil
	}
}

func (g *CloudActionGenerator) GenerateDnsApplianceDeleteAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewDNSOp(*g.ApiClient.Caller)
	if cmd.NArg() == 0 {
		return fmt.Errorf("DNS appliance id required")
	}
	for _, a := range cmd.Args().Slice() {
		id, err := getDNSApplianceIdByString(s, a)
		if err != nil {
			return err
		}
		err = s.Delete(context.Background(), id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *CloudActionGenerator) GenerateDnsApplianceListAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewDNSOp(*g.ApiClient.Caller)
	res, err := s.Find(context.Background(), nil)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateDnsRecordListAction(_ context.Context, cmd *cli.Command) error {
	s := iaas.NewDNSOp(*g.ApiClient.Caller)
	i, err := strconv.Atoi(cmd.Args().First())
	if err != nil {
		return err
	}
	res, err := s.Read(context.Background(), types.ID(i))
	if err != nil {
		return err
	}
	return helper.PrintJson(res.Records)
}

func (g *CloudActionGenerator) GenerateDnsRecordAddAction(_ context.Context, cmd *cli.Command) error {
	s := iaas.NewDNSOp(*g.ApiClient.Caller)
	ttl := cmd.Int("ttl")
	if cmd.NArg() < 3 {
		return fmt.Errorf("too few arguments")
	}
	name := cmd.Args().First()
	rtype := cmd.Args().Get(1)
	rdata := cmd.Args().Get(2)
	var id types.ID
	if cmd.String("id") == "" {
		return fmt.Errorf("DNS appliance ID must be specified")
	}
	i, err := strconv.Atoi(cmd.String("id"))
	if err != nil {
		id = types.StringID(cmd.String("id"))
	} else {
		id = types.ID(i)
	}
	origRRSet, err := s.Read(context.Background(), id)
	if err != nil {
		return err
	}
	rrSet := origRRSet.Records
	rrSet = append(rrSet, &iaas.DNSRecord{
		Name:  name,
		Type:  types.EDNSRecordType(rtype),
		RData: rdata,
		TTL:   ttl,
	})
	_, err = s.Update(context.Background(), id, &iaas.DNSUpdateRequest{
		Records: rrSet,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *CloudActionGenerator) GenerateDnsRecordDeleteAction(_ context.Context, cmd *cli.Command) error {
	s := iaas.NewDNSOp(*g.ApiClient.Caller)
	var id types.ID

	i, err := strconv.Atoi(cmd.String("id"))
	if err != nil {
		id = types.StringID(cmd.String("id"))
	} else {
		id = types.ID(i)
	}
	res, err := s.Read(context.Background(), id)
	if err != nil {
		return err
	}
	a := cmd.Args().First()
	reg := cmd.Bool("regex")
	regex, err := regexp.CompilePOSIX(a)
	if reg && err != nil {
		return err
	}
	matchRRType := false
	switch a {
	case "A":
		matchRRType = true
	case "AAAA":
		matchRRType = true
	case "MX":
		matchRRType = true
	case "CNAME":
		matchRRType = true
	case "TXT":
		matchRRType = true
	case "NS":
		matchRRType = true
	case "ALIAS":
		matchRRType = true
	case "CAA":
		matchRRType = true
	}
	var residualRR []*iaas.DNSRecord
	for _, rr := range res.Records {
		if matchRRType {
			if rr.Type.String() != a {
				residualRR = append(residualRR, rr)
			}
		} else if a == "*" {
			residualRR = append(residualRR, rr)
		} else if !reg && rr.Name != a {
			residualRR = append(residualRR, rr)
		} else if reg && !regex.MatchString(fmt.Sprintf("%s %s %s", rr.Name, rr.Type.String(), rr.RData)) {
			residualRR = append(residualRR, rr)
		}
	}
	if len(res.Records) != 0 && len(residualRR) == len(res.Records) {
		return errors.New("no records to delete")
	}
	_, err = s.Update(context.Background(), id, &iaas.DNSUpdateRequest{
		Records: residualRR,
	})
	return err
}

func (g *CloudActionGenerator) GenerateDnsApplianceRecordImportAction(_ context.Context, cmd *cli.Command) error {
	s := iaas.NewDNSOp(*g.ApiClient.Caller)
	if cmd.NArg() != 1 {
		return fmt.Errorf("an appliance-id must be specified")
	}
	id, err := getDNSApplianceIdByString(s, cmd.Args().First())
	if err != nil {
		return err
	}
	ap, err := s.Read(context.Background(), id)
	if err != nil {
		return err
	}
	var rrSets []*iaas.DNSRecord
	fileName := cmd.String("file")
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("could not open file %q: %w", fileName, err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var savedResponse iaas.DNS
	err = json.Unmarshal(data, &savedResponse)
	if err == nil {
		rrSets = savedResponse.Records
	} else {
		f, err = os.Open(fileName)
		if err != nil {
			return fmt.Errorf("could not open file: %q: %w", fileName, err)
		}
		//parsing zonefile
		scanner := gozone.NewScanner(f)
		var r gozone.Record
		for {
			err = scanner.Next(&r)
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			var name string
			if len(ap.DNSZone)+1 == len(r.DomainName) {
				name = "@"
			} else {
				name = r.DomainName[0 : len(r.DomainName)-len(ap.DNSZone)-2]
			}
			rrSets = append(rrSets, &iaas.DNSRecord{
				Name:  name,
				Type:  types.EDNSRecordType(r.Type.String()),
				RData: strings.Join(r.Data, " "),
				TTL:   int(r.TimeToLive),
			})
		}
	}
	_, err = s.Update(context.Background(), id, &iaas.DNSUpdateRequest{
		Records: rrSets,
	})
	return err
}
