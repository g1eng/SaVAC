package helper

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/g1eng/savac/pkg/core"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/webaccel-api-go"
	"github.com/urfave/cli/v3"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/ghodss/yaml"
	"github.com/hokaccha/go-prettyjson"
	"github.com/olekukonko/tablewriter"
)

func GetOutputDigitByName(name string) (int, error) {
	switch name {
	case "json":
		return core.OutputTypeJson, nil
	case "yaml":
		return core.OutputTypeYaml, nil
	case "table":
		return core.OutputTypeText, nil
	default:
		return -1, fmt.Errorf("unknown output: %s", name)
	}
}

func PrintJson(data interface{}) error {
	s, err := prettyjson.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println(string(s))
	return nil
}

func PrintYaml(data interface{}) error {
	s, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println(string(s))
	return nil
}

func PrintTableForServerInfo(res *sakuravps.Server) {
	t := NewTable()

	t.Append([]string{"id", "", strconv.Itoa(int(res.Id))})
	t.Append([]string{"name", "", res.Name})
	if res.Description != "" {
		t.Append([]string{"description", "", res.Description})
	}
	t.Append([]string{"status", "", res.PowerStatus})
	t.Append([]string{"spec", "core", strconv.Itoa(int(res.CpuCores))})
	t.Append([]string{"spec", "memory", strconv.Itoa(int(res.MemoryMebibytes))})
	t.Append([]string{"zone", "", fmt.Sprintf("%s (%s)", res.Zone.Code, res.Zone.Name)})

	t.Append([]string{"ipv4", "hostname", res.Ipv4.Hostname})
	t.Append([]string{"ipv4", "ptr", res.Ipv4.Ptr})
	t.Append([]string{"ipv4", "address", res.Ipv4.Address})
	t.Append([]string{"ipv4", "netmask", res.Ipv4.Netmask})
	t.Append([]string{"ipv4", "gateway", res.Ipv4.Gateway})
	for _, n := range res.Ipv4.Nameservers {
		t.Append([]string{"ipv4", "nameservers", n})
	}
	t.Append([]string{"ipv6", "hostname", res.Ipv6.GetHostname()})
	t.Append([]string{"ipv6", "ptr", res.Ipv6.GetPtr()})
	t.Append([]string{"ipv6", "address", res.Ipv6.GetAddress()})
	t.Append([]string{"ipv6", "gateway", res.Ipv6.GetGateway()})
	for _, n := range res.Ipv6.Nameservers {
		t.Append([]string{"ipv6", "nameservers", n})
	}

	t.Render()
}

func PrintTableForServerInterfaces(svName string, res []sakuravps.ServerInterface, switches []sakuravps.Switch) {
	t := NewList()
	//if !s.NoHeader {
	//	t.SetHeader([]string{"server", "NIC id", "name", "address", "switchId", "switch"})
	//}
	for _, inf := range res {
		s := []byte(inf.Mac)
		for i, c := range inf.Mac {
			switch c {
			case 'A':
				s[i] = 'a'
			case 'B':
				s[i] = 'b'
			case 'C':
				s[i] = 'c'
			case 'D':
				s[i] = 'd'
			case 'E':
				s[i] = 'e'
			case 'F':
				s[i] = 'f'
			}
		}
		var switchId, switchName string
		if inf.ConnectableToGlobalNetwork && inf.GetSwitchId() == 0 {
			switchId = "shared"
			switchName = "global"
		} else if inf.GetSwitchId() == 0 {
			switchId = "-"
			switchName = "-"
		} else {
			switchId = fmt.Sprintf("%d", inf.GetSwitchId())
			switchName = fmt.Sprintf("%d", inf.SwitchId.Get())
			for _, sw := range switches {
				if sw.Id == inf.GetSwitchId() {
					switchName = sw.Name
				}
			}
		}
		t.Append([]string{svName, fmt.Sprintf("%d", inf.Id), inf.DisplayName, string(s), switchId, switchName})
	}
	t.Render()
}

func PrintMonitoringList(res []core.ServerMonitoringMeta) {
	t := NewList()
	if len(res) == 0 {
		return
	}
	//if !s.NoHeader {
	//	t.SetHeader([]string{"id", "name", "resource_id", "monitored", "protocol", "notification"})
	//}
	for _, d := range res {
		status := "yes"
		if !d.Settings.Enabled {
			status = "no"
		}
		_, eo := d.Settings.Notification.GetEmailOk()
		w := d.Settings.Notification.GetIncomingWebhook()
		wo := w.GetWebhooksUrl() != ""

		notification := []string{}
		if eo {
			notification = append(notification, "email")
		}
		if wo {
			notification = append(notification, "webhook")
		}

		protocol := d.Settings.HealthCheck.Protocol
		t.SetHeader([]string{
			"server",
			"mon-id",
			"name",
			"resource-id",
			"enabled",
			"protocol",
			"notification",
		})

		t.Append([]string{
			strconv.Itoa(int(d.ServerId)),
			strconv.Itoa(int(d.Id)),
			d.Name,
			d.MonitoringResourceId,
			status,
			protocol,
			strings.Join(notification, "/"),
		})
	}
	t.Render()
}

func PrintMonitoringDetailTable(res map[int32][]core.ServerMonitoringMeta, noHeader ...bool) {
	t := NewList()
	for _, d := range res {
		if len(d) == 0 {
			return
		}
		break
	}
	if len(noHeader) == 0 || (len(noHeader) != 0 && !noHeader[0]) {
		t.SetHeader([]string{"server", "mon_id", "name", "resource_id", "monitored", "protocol", "notification", "port", "host", "path"})
	}
	for serverId, m := range res {
		for _, d := range m {
			status := "yes"
			if !d.Settings.Enabled {
				status = "no"
			}
			_, eo := d.Settings.Notification.GetEmailOk()
			w := d.Settings.Notification.GetIncomingWebhook()
			wo := w.GetWebhooksUrl() != ""

			notification := []string{}
			if eo {
				notification = append(notification, "email")
			}
			if wo {
				notification = append(notification, "webhook")
			}

			protocol := ""
			port := "-"
			host := "-"
			path := "-"
			protocol = d.Settings.HealthCheck.Protocol
			switch protocol {
			case "ping":
			case "ssh":
				port = fmt.Sprintf("%d", *d.Settings.HealthCheck.Port.Get())
			case "http":
				port = fmt.Sprintf("%d", *d.Settings.HealthCheck.Port.Get())
				host = *d.Settings.HealthCheck.Host.Get()
				path = *d.Settings.HealthCheck.Path.Get()
			case "https":
				port = fmt.Sprintf("%d", *d.Settings.HealthCheck.Port.Get())
				host = *d.Settings.HealthCheck.Host.Get()
				path = *d.Settings.HealthCheck.Path.Get()
				//sni = *d.Settings.HealthCheck.Sni.Get()
			case "smtp":
			case "pop3":
			case "imap":
			case "tcp":
				port = fmt.Sprintf("%d", *d.Settings.HealthCheck.Port.Get())
			default:
				fmt.Fprintf(os.Stderr, "Unknown health check protocol: %s\n", d.Settings.HealthCheck.Protocol)
			}

			t.Append([]string{
				strconv.Itoa(int(serverId)),
				strconv.Itoa(int(d.Id)),
				d.Name,
				d.MonitoringResourceId,
				status,
				protocol,
				strings.Join(notification, "/"),
				port,
				host,
				path,
			})
		}
	}
	t.Render()
}

func PrintRolesDetail(roles []sakuravps.Role) error {
	if roles == nil {
		return errors.New("given roles is nil")
	}
	t := NewTable()
	t.SetAutoMergeCells(false)
	t.Append([]string{"id", "name", "description", "permissions", "server", "switch", "nfs"})
	for _, role := range roles {
		var tbody []string
		resource := role.AllowedResources.Get()
		permString := strings.Join(role.AllowedPermissions, "\n")
		if role.PermissionFiltering != "enabled" {
			permString = "*"
		}
		if resource != nil {
			if role.ResourceFiltering != "enabled" {
				tbody = []string{
					fmt.Sprintf("%d", role.Id),
					role.Name,
					role.Description,
					permString,
					"*",
					"*",
					"*",
				}
			} else {
				tbody = []string{
					fmt.Sprintf("%d", role.Id),
					role.Name,
					role.Description,
					permString,
					ConcatResourceIds(resource.Servers),
					ConcatResourceIds(resource.Switches),
					ConcatResourceIds(resource.NfsServers),
				}
			}
		}
		t.Append(tbody)
	}
	t.Render()
	return nil
}

func PrintTableForWebaccelCertificateRevisions(res *webaccel.Certificates) error {
	t := NewList()
	t.SetHeader([]string{
		"rev",
		"serial",
		"dns_names",
		"O",
		"ISS",
		"not_before",
		"not_after",
		"fingerprint",
		"ca_bundles",
	})

	certSepIndex := strings.LastIndex(res.Current.CertificateChain, "-----BEGIN CERTIFICATE-----")
	parentCertObj, err := parseCert(res.Current.CertificateChain[certSepIndex:])
	if err != nil {
		return err
	}

	t.Append([]string{
		fmt.Sprintf("%d", len(res.Old)+1),
		res.Current.SHA256Fingerprint[:14],
		strings.Join(res.Current.DNSNames, "\n"),
		res.Current.Subject.Organization,
		res.Current.Issuer.CommonName,
		time.Unix(res.Current.NotBefore/1000, 0).Local().Format(time.RFC822),
		time.Unix(res.Current.NotAfter/1000, 0).Local().Format(time.RFC822),
		res.Current.SerialNumber[:14],
		parentCertObj.Subject.Names[0].Value.(string),
	})
	for i, cert := range res.Old {
		rev := len(res.Old) - i
		certSepIndex = strings.LastIndex(cert.CertificateChain, "-----BEGIN CERTIFICATE-----")

		//separate the intermediate CA certificate
		certPem, parentCert := cert.CertificateChain, cert.CertificateChain[certSepIndex:]
		certPem = certPem[:len(certPem)-1]
		crtObj, err := parseCert(certPem)
		if err != nil {
			return err
		}
		serial := formatHex(fmt.Sprintf("%x", crtObj.SerialNumber))[:14]
		fingerprint, err := getFingerprintHex(certPem)
		if err != nil {
			return err
		}
		parentCertObj, err := parseCert(parentCert)
		if err != nil {
			return err
		}

		t.Append([]string{
			fmt.Sprintf("%d", rev),
			fingerprint[:14],
			strings.Join(crtObj.DNSNames, "\n"),
			strings.Join(crtObj.Subject.Organization, ","),
			crtObj.Issuer.CommonName,
			crtObj.NotBefore.Local().Format(time.RFC822),
			crtObj.NotAfter.Local().Format(time.RFC822),
			serial,
			parentCertObj.Subject.Names[0].Value.(string),
		})
	}
	t.Render()
	return nil
}

func FormatAsZoneFile(res *iaas.DNS) string {
	var out string
	out += fmt.Sprintf("; Zone: %s\n", res.DNSZone)
	out += fmt.Sprintf("$ORIGIN %s.\n", res.DNSZone)
	out += "$TTL 3600\n\n"
	for _, rr := range res.Records {
		out += fmt.Sprintf("%s %d IN %s %s\n", rr.Name, rr.TTL, rr.Type, rr.RData)
	}
	return out
}

func NewList() *tablewriter.Table {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetNoWhiteSpace(true)
	// t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	t.SetTablePadding("  ")
	t.SetRowLine(false)
	t.SetHeaderLine(false)
	t.SetBorder(false)
	return t
}

func NewTable() *tablewriter.Table {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetRowLine(true)
	t.SetHeaderLine(true)
	t.SetBorder(true)
	t.SetNoWhiteSpace(false)
	t.SetAutoMergeCells(true)
	t.SetAlignment(tablewriter.ALIGN_LEFT)
	return t
}

func GenRandomString(length int) string {
	s := make([]byte, length)
	for i := range s {
		r := rand.Intn(62) // nolint
		if r < 10 {
			r += 0x30
		} else if r < 36 {
			r += 0x41 - 10
		} else {
			r += 0x61 - 36
		}
		s[i] = byte(r)
	}
	return string(s)
}

func CheckArgsExist(ctx context.Context, command *cli.Command) (context.Context, error) {
	if command.Args().Len() == 0 {
		return ctx, fmt.Errorf("no arguments")
	}
	return ctx, nil
}

func CheckTwoArgsExist(ctx context.Context, command *cli.Command) (context.Context, error) {
	if command.Args().Len() != 2 {
		return ctx, fmt.Errorf("just two arguments required")
	}
	return ctx, nil
}

// calculate hex value for
func getFingerprintHex(certPem string) (string, error) {
	b, _ := pem.Decode([]byte(certPem))
	if b == nil {
		return "", fmt.Errorf("failed to read PEM encoded certificate")
	}
	fingerprint := sha256.Sum256(b.Bytes)
	fingerprintHex := hex.EncodeToString(fingerprint[:])

	return formatHex(fingerprintHex), nil
}

func formatHex(hexVal string) string {
	l := len(hexVal)
	for p := range hexVal {
		if p > 0 && p%2 == 0 {
			hexVal = hexVal[:l-p] + ":" + hexVal[l-p:]
		}
	}
	return strings.ToUpper(hexVal)
}

func parseCert(certPem string) (*x509.Certificate, error) {
	b, _ := pem.Decode([]byte(certPem))
	if b == nil {
		return nil, fmt.Errorf("failed to parse PEM-encoded certificate")
	}
	return x509.ParseCertificate(b.Bytes)
}

func ConcatResourceIds(a []int32) (res string) {
	var buf []string
	for _, v := range a {
		buf = append(buf, fmt.Sprintf("%d", v))
	}
	return strings.Join(buf, ",")
}
