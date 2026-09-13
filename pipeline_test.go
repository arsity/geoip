package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/netip"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/Loyalsoldier/geoip/lib"
	"github.com/Loyalsoldier/geoip/plugin/v2ray"
	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"github.com/oschwald/geoip2-golang/v2"
	"google.golang.org/protobuf/proto"
)

type pipelineStep struct {
	Type   string         `json:"type"`
	Action string         `json:"action"`
	Args   map[string]any `json:"args,omitempty"`
}

type pipelineConfig struct {
	Input  []pipelineStep `json:"input"`
	Output []pipelineStep `json:"output"`
}

// Synthetic fixtures are not assertions about real-world geolocation. Keep the
// production config's policies and replace only its external input locations.
func TestFixturePipeline(t *testing.T) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		t.Fatal(err)
	}
	var config pipelineConfig
	if err := json.Unmarshal(data, &config); err != nil {
		t.Fatal(err)
	}
	t.Chdir(t.TempDir())
	if err := os.MkdirAll("ipinfo", 0700); err != nil {
		t.Fatal(err)
	}
	writeFixture(t, "ipinfo/ipinfo_lite.csv", "network,country,country_code,continent,continent_code,asn,as_name,as_domain\n"+
		"1.0.0.0/24,Australia,AU,Oceania,OC,AS13335,Cloudflare,cloudflare.com\n"+
		"8.8.8.0/24,United States,US,North America,NA,AS15169,Google,google.com\n"+
		"2001:4860::/32,United States,US,North America,NA,AS15169,Google,google.com\n"+
		"114.114.114.0/24,China,CN,Asia,AS,AS4134,Other,example.cn\n"+
		"2400:3200::/32,China,CN,Asia,AS,AS4134,Other,example.cn\n"+
		"223.5.5.0/24,United States,US,North America,NA,AS65535,Other,example.com\n"+
		"31.13.64.0/24,United States,US,North America,NA,AS32934,Meta,facebook.com\n"+
		"151.101.0.0/24,United States,US,North America,NA,AS54113,Fastly,fastly.com\n"+
		"23.246.0.0/24,United States,US,North America,NA,AS2906,Netflix,netflix.com\n"+
		"78.31.8.0/24,United States,US,North America,NA,AS8403,Spotify,spotify.com\n"+
		"149.154.160.0/24,United States,US,North America,NA,AS62041,Telegram,telegram.org\n"+
		"104.244.42.0/24,United States,US,North America,NA,AS13414,Twitter,x.com\n")
	writeMetadataFixture(t)
	for idx, input := range config.Input {
		uri, _ := input.Args["uri"].(string)
		if !strings.HasPrefix(uri, "https://") {
			continue
		}
		name, _ := input.Args["name"].(string)
		var content string
		switch input.Type + "/" + name {
		case "text/cn":
			content = "223.5.5.0/24\n"
			if input.Args["onlyIPType"] == "ipv6" {
				content = "240e::/32\n"
			}
		case "text/tor":
			content = "8.8.8.8\n"
		case "text/cloudflare":
			content = "1.1.1.0/24\n10.0.0.0/24\n"
			if strings.HasSuffix(uri, "ips-v6") {
				content = "2606:4700::/32\n"
			}
		case "text/telegram":
			content = "149.154.160.0/24\n"
		case "json/google":
			content = `{"prefixes":[{"ipv4Prefix":"8.8.8.0/24"},{"ipv6Prefix":"2001:4860::/32"}]}`
		case "json/fastly":
			content = `{"addresses":["151.101.0.0/24"],"ipv6_addresses":["2a04:4e42::/32"]}`
		case "json/cloudfront":
			content = `{"prefixes":[{"service":"CLOUDFRONT","ip_prefix":"13.32.0.0/24"}],"ipv6_prefixes":[{"service":"CLOUDFRONT","ipv6_prefix":"2600:9000::/32"}]}`
		default:
			t.Fatalf("new remote source needs a fixture: %s/%s", input.Type, name)
		}
		path := fmt.Sprintf("input-%d.txt", idx)
		writeFixture(t, path, content)
		input.Args["uri"] = path
	}
	data, err = json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	instance, err := lib.NewInstance()
	if err != nil {
		t.Fatal(err)
	}
	if err := instance.InitConfigFromBytes(data); err != nil {
		t.Fatal(err)
	}
	if err := instance.Run(); err != nil {
		t.Fatal(err)
	}
	verifyOutputs(t, "output")

	// Check overlapping country/service rules, CN replacement and private/Tor
	// precedence, including IPv6. Blank results must not become invented data.
	for file, cases := range map[string]map[string]string{
		"Country.mmdb":                 {"1.0.0.1": "CLOUDFLARE", "8.8.8.9": "GOOGLE", "8.8.8.8": "TOR", "2001:4860::1": "GOOGLE", "78.31.8.1": "SPOTIFY", "223.5.5.1": "CN", "240e::1": "CN", "10.0.0.1": "PRIVATE"},
		"Country-without-asn.mmdb":     {"1.0.0.1": "AU", "8.8.8.8": "US", "2001:4860::1": "US", "78.31.8.1": "US", "223.5.5.1": "CN", "240e::1": "CN", "114.114.114.1": "", "2400:3200::1": ""},
		"Country-only-cn-private.mmdb": {"223.5.5.1": "CN", "240e::1": "CN", "8.8.8.8": "", "10.0.0.1": "PRIVATE"},
		"Country-asn.mmdb":             {"8.8.8.8": "TOR", "78.31.8.1": "SPOTIFY", "10.0.0.1": "CLOUDFLARE", "223.5.5.1": ""},
	} {
		checkCountryLookups(t, filepath.Join("output", file), cases)
	}

	// Read generated rule files back through their supported input converters.
	for format, path := range map[string]string{
		"v2rayGeoIPDat": "cn.dat", "maxmindMMDB": "Country-only-cn-private.mmdb",
		"singboxSRS": "srs/cn.srs", "mihomoMRS": "mrs/cn.mrs", "text": "text/cn.txt",
		"clashRuleSet": "clash/ipcidr/cn.txt", "clashRuleSetClassical": "clash/classical/cn.txt", "surgeRuleSet": "surge/cn.txt",
	} {
		in := getInputForLookup(format, "cn", filepath.Join("output", path), "")
		c, err := in.Input(lib.NewContainer())
		if err != nil {
			t.Fatalf("round-trip %s: %v", format, err)
		}
		entry, found := c.GetEntry("CN")
		if !found {
			t.Fatalf("round-trip %s lost CN", format)
		}
		prefixes, err := entry.MarshalText(nil)
		if err != nil {
			t.Fatal(err)
		}
		slices.Sort(prefixes)
		if !slices.Equal(prefixes, []string{"223.5.5.0/24", "240e::/32"}) {
			t.Errorf("round-trip %s: %v", format, prefixes)
		}
	}
}

func writeFixture(t *testing.T, path, data string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(data), 0600); err != nil {
		t.Fatal(err)
	}
}

func writeMetadataFixture(t *testing.T) {
	t.Helper()
	writer, err := mmdbwriter.New(mmdbwriter.Options{DatabaseType: "IPinfo-Lite", IncludeReservedNetworks: true})
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range []struct{ prefix, code, country, continent, continentCode string }{
		{"1.0.0.0/24", "AU", "Australia", "Oceania", "OC"},
		{"8.8.8.0/24", "US", "United States", "North America", "NA"},
		{"114.114.114.0/24", "CN", "China", "Asia", "AS"},
	} {
		_, network, err := net.ParseCIDR(row.prefix)
		if err != nil {
			t.Fatal(err)
		}
		err = writer.Insert(network, mmdbtype.Map{"country_code": mmdbtype.String(row.code), "country": mmdbtype.String(row.country), "continent": mmdbtype.String(row.continent), "continent_code": mmdbtype.String(row.continentCode)})
		if err != nil {
			t.Fatal(err)
		}
	}
	file, err := os.Create("ipinfo/ipinfo_lite.mmdb")
	if err != nil {
		t.Fatal(err)
	}
	_, err = writer.WriteTo(file)
	closeErr := file.Close()
	if err != nil {
		t.Fatal(err)
	}
	if closeErr != nil {
		t.Fatal(closeErr)
	}
}

// Run separately against production output BEFORE publishing anything.
func TestReleaseOutputs(t *testing.T) {
	dir := os.Getenv("GEOIP_VERIFY_OUTPUT_DIR")
	if dir == "" {
		t.Skip("set GEOIP_VERIFY_OUTPUT_DIR to validate generated production files")
	}
	verifyOutputs(t, dir)
}

func verifyOutputs(t *testing.T, dir string) {
	t.Helper()
	all := readDAT(t, filepath.Join(dir, "geoip.dat"))
	services := []string{"CLOUDFLARE", "CLOUDFRONT", "FACEBOOK", "FASTLY", "GOOGLE", "NETFLIX", "SPOTIFY", "TELEGRAM", "TWITTER", "TOR"}
	for _, name := range append(append([]string{}, services...), "CN", "PRIVATE", "US", "AU") {
		if _, found := all[name]; !found {
			t.Fatalf("geoip.dat is missing required category %s", name)
		}
	}
	for file, want := range map[string][]string{
		"cn.dat": {"CN"}, "private.dat": {"PRIVATE"}, "geoip-only-cn-private.dat": {"CN", "PRIVATE"}, "geoip-asn.dat": services,
	} {
		entries := readDAT(t, filepath.Join(dir, file))
		if len(entries) != len(want) {
			t.Fatalf("%s: unexpected category count %d", file, len(entries))
		}
		for _, name := range want {
			if !proto.Equal(entries[name], all[name]) {
				t.Fatalf("%s disagrees with geoip.dat for %s", file, name)
			}
		}
	}
	for name := range all {
		for _, format := range []struct{ folder, ext string }{
			{"dat", ".dat"}, {"mrs", ".mrs"}, {"srs", ".srs"}, {"text", ".txt"},
			{"clash/classical", ".txt"}, {"clash/ipcidr", ".txt"}, {"surge", ".txt"},
			{"nginx/allow", ".conf"}, {"nginx/deny", ".conf"},
		} {
			requireOutputFile(t, filepath.Join(dir, format.folder, strings.ToLower(name)+format.ext))
		}
	}
	for _, file := range []string{"Country.mmdb", "Country-without-asn.mmdb", "Country-only-cn-private.mmdb", "Country-asn.mmdb"} {
		path := filepath.Join(dir, file)
		requireOutputFile(t, path)
		cases := map[string]string{}
		if file != "Country-asn.mmdb" {
			cases = map[string]string{"10.0.0.1": "PRIVATE", "127.0.0.1": "PRIVATE", "::1": "PRIVATE"}
		}
		checkCountryLookups(t, path, cases)
	}
	t.Logf("verified %d categories, aggregate DAT consistency, rule-format files and MMDB reader compatibility", len(all))
}

func requireOutputFile(t *testing.T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("missing output %s: %v", path, err)
	}
	if !info.Mode().IsRegular() || info.Size() == 0 {
		t.Fatalf("empty/non-file output %s", path)
	}
}

func readDAT(t *testing.T, path string) map[string]*v2ray.GeoIP {
	t.Helper()
	requireOutputFile(t, path)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var file v2ray.GeoIPList
	if err := proto.Unmarshal(data, &file); err != nil {
		t.Fatalf("invalid DAT %s: %v", path, err)
	}
	entries := make(map[string]*v2ray.GeoIP)
	for _, entry := range file.GetEntry() {
		name := entry.GetCountryCode()
		if name == "" || len(entry.GetCidr()) == 0 {
			t.Fatalf("empty category in %s", path)
		}
		if _, duplicate := entries[name]; duplicate {
			t.Fatalf("duplicate category %s in %s", name, path)
		}
		for _, cidr := range entry.GetCidr() {
			addr, ok := netip.AddrFromSlice(cidr.GetIp())
			if !ok || cidr.GetPrefix() > uint32(addr.BitLen()) {
				t.Fatalf("invalid CIDR in %s/%s", path, name)
			}
		}
		entries[name] = entry
	}
	if len(entries) == 0 {
		t.Fatalf("empty DAT %s", path)
	}
	return entries
}

func checkCountryLookups(t *testing.T, path string, cases map[string]string) {
	t.Helper()
	reader, err := geoip2.Open(path)
	if err != nil {
		t.Fatalf("MaxMind reader cannot open %s: %v", path, err)
	}
	defer reader.Close()
	for ip, want := range cases {
		record, err := reader.Country(netip.MustParseAddr(ip))
		if err != nil {
			t.Fatalf("MaxMind Country lookup in %s: %v", path, err)
		}
		if record.Country.ISOCode != want {
			t.Errorf("%s [%s]: got %q, want %q", path, ip, record.Country.ISOCode, want)
		}
	}
}
