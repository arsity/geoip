package maxmind

import (
	"encoding/json"
	"net/netip"
	"os"
	"path/filepath"
	"testing"

	"github.com/Loyalsoldier/geoip/lib"
)

// Deliberately synthetic ownership conflicts protect this fork's intentional
// metadata-first policy, including its non-matching-metadata ASN fallback.
const converterCSV = "as_domain,network,country_code,asn,as_name\n" +
	"google.com,1.0.0.0/24,AU,AS13335,Google\n" +
	"unrelated.example,2.0.0.0/24,FR,AS13335,Other\n" +
	"sub.spotify.com,3.0.0.0/24,US,AS8403,Spotify\n" +
	"google.com,2001:4860::/32,US,AS15169,Google\n" +
	"google.com,4.0.0.0/24,US,,No ASN\n" +
	"google.com,5.0.0.0/24,,AS15169,No country\n"

func csvFixture(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func entryContains(t *testing.T, c lib.Container, name, ip string) bool {
	t.Helper()
	entry, ok := c.GetEntry(name)
	if !ok {
		return false
	}
	prefixes, err := entry.MarshalText(nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, text := range prefixes {
		if netip.MustParsePrefix(text).Contains(netip.MustParseAddr(ip)) {
			return true
		}
	}
	return false
}

func TestIPInfoCountrySelectionAndFamilies(t *testing.T) {
	path := csvFixture(t, converterCSV)
	for _, family := range []string{"", "ipv4", "ipv6"} {
		t.Run(family, func(t *testing.T) {
			args, _ := json.Marshal(map[string]any{"uri": path, "wantedList": []string{" us "}, "onlyIPType": family})
			in, err := newIPInfoLiteCountryCSVIn(lib.ActionAdd, args)
			if err != nil {
				t.Fatal(err)
			}
			c, err := in.Input(lib.NewContainer())
			if err != nil {
				t.Fatal(err)
			}
			if entryContains(t, c, "US", "3.0.0.1") != (family != "ipv6") {
				t.Fatal("incorrect IPv4 filter")
			}
			if entryContains(t, c, "US", "2001:4860::1") != (family != "ipv4") {
				t.Fatal("incorrect IPv6 filter")
			}
			if entryContains(t, c, "AU", "1.0.0.1") {
				t.Fatal("wantedList ignored")
			}
			if entryContains(t, c, "US", "5.0.0.1") {
				t.Fatal("blank country was not skipped")
			}
		})
	}
}

func TestIPInfoASNMetadataPriorityAndFallback(t *testing.T) {
	args, _ := json.Marshal(map[string]any{
		"uri":          csvFixture(t, converterCSV),
		"wantedList":   map[string][]string{"cloudflare": {"AS13335"}, "google": {"AS15169"}, "spotify": {"AS8403"}},
		"metadataList": map[string][]string{"google": {"google.com"}, "spotify": {"spotify.com"}},
	})
	in, err := newIPInfoLiteASNCSVIn(lib.ActionAdd, args)
	if err != nil {
		t.Fatal(err)
	}
	c, err := in.Input(lib.NewContainer())
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		list, ip string
		want     bool
	}{
		{"GOOGLE", "1.0.0.1", true}, {"CLOUDFLARE", "1.0.0.1", false},
		{"CLOUDFLARE", "2.0.0.1", true}, {"SPOTIFY", "3.0.0.1", true},
		{"GOOGLE", "2001:4860::1", true}, {"GOOGLE", "4.0.0.1", false}, {"GOOGLE", "5.0.0.1", true},
	} {
		if got := entryContains(t, c, tc.list, tc.ip); got != tc.want {
			t.Errorf("%s/%s: got %v", tc.list, tc.ip, got)
		}
	}
	if matchIPInfoASNMetadata("google.com", "notgoogle.com", "Google") || matchIPInfoASNMetadata("google.com", "google.com.evil.example", "Google") {
		t.Fatal("domain boundary matching regressed")
	}
}

func TestIPInfoCSVFailureDoesNotMutateContainer(t *testing.T) {
	for _, asn := range []bool{false, true} {
		args, _ := json.Marshal(map[string]any{"uri": csvFixture(t, converterCSV+"google.com,INVALID,US,AS15169,Google\n")})
		var in lib.InputConverter
		var err error
		if asn {
			in, err = newIPInfoLiteASNCSVIn(lib.ActionAdd, args)
		} else {
			in, err = newIPInfoLiteCountryCSVIn(lib.ActionAdd, args)
		}
		if err != nil {
			t.Fatal(err)
		}
		c := lib.NewContainer()
		if _, err := in.Input(c); err == nil {
			t.Fatal("accepted invalid row")
		}
		for range c.Loop() {
			t.Fatal("partially imported failed input")
		}
	}
}

func TestIPInfoCSVEmptyAndMissingMetadata(t *testing.T) {
	for _, content := range []string{"network,country_code,asn\n", "network,country_code,asn\n1.0.0.0/24,,\n"} {
		args, _ := json.Marshal(map[string]any{"uri": csvFixture(t, content)})
		country, _ := newIPInfoLiteCountryCSVIn(lib.ActionAdd, args)
		asn, _ := newIPInfoLiteASNCSVIn(lib.ActionAdd, args)
		for _, in := range []lib.InputConverter{country, asn} {
			if _, err := in.Input(lib.NewContainer()); err == nil {
				t.Fatal("accepted empty input")
			}
		}
	}
	args, _ := json.Marshal(map[string]any{"uri": csvFixture(t, "network,asn\n1.0.0.0/24,AS13335\n"), "metadataList": map[string][]string{"cloudflare": {"cloudflare.com"}}})
	in, _ := newIPInfoLiteASNCSVIn(lib.ActionAdd, args)
	if _, err := in.Input(lib.NewContainer()); err == nil {
		t.Fatal("silently fell back with missing metadata columns")
	}
}
