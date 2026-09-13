package maxmind

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestIPInfoCSVHeaderAndRows(t *testing.T) {
	input := "\ufeffextra,ASN, country_code ,as_domain,network,as_name\r\n" +
		"ignored,AS13335,au,cloudflare.com,1.0.0.0/24,\"Cloudflare, Inc\"\r\n" +
		"ignored,AS15169,US,google.com,2001:4860::/32,Google\r\n" +
		"ignored,,,,9.9.9.9,\r\n"
	r, err := newIPInfoLiteCSVReader(strings.NewReader(input), "network", "country_code", "asn", "as_domain")
	if err != nil {
		t.Fatal(err)
	}
	first, err := r.Read()
	if err != nil || first.Network != "1.0.0.0/24" || first.CountryCode != "AU" || first.ASName != "Cloudflare, Inc" {
		t.Fatalf("%+v %v", first, err)
	}
	second, err := r.Read()
	if err != nil || second.Network != "2001:4860::/32" || second.ASN != "AS15169" {
		t.Fatalf("%+v %v", second, err)
	}
	if _, err := r.Read(); err != nil {
		t.Fatal(err)
	}
	if _, err := r.Read(); !errors.Is(err, io.EOF) {
		t.Fatalf("want EOF, got %v", err)
	}
	if r.rows != 3 || r.emptyASN != 1 || r.emptyCountry != 1 {
		t.Fatalf("incorrect counters: %+v", r)
	}
}

func TestIPInfoCSVRejectsBadHeaders(t *testing.T) {
	for _, input := range []string{"", "\"unterminated", "network,asn\n", "network,country_code, NETWORK\n", "network,country_code,\n"} {
		if _, err := newIPInfoLiteCSVReader(strings.NewReader(input), "network", "country_code"); err == nil {
			t.Errorf("accepted bad header %q", input)
		}
	}
}

func TestIPInfoCSVRejectsBadRows(t *testing.T) {
	for _, row := range []string{
		"1.0.0.0/24,AU\n", "1.0.0.0/24,AU,AS13335,extra\n", "\"unterminated\n",
		"invalid,AU,AS13335\n", "1.0.0.0/33,AU,AS13335\n", ",AU,AS13335\n",
		"1.0.0.0/24,Australia,AS13335\n", "1.0.0.0/24,A1,AS13335\n",
		"1.0.0.0/24,AU,ASfoo\n", "1.0.0.0/24,AU,AS4294967296\n",
	} {
		r, err := newIPInfoLiteCSVReader(strings.NewReader("network,country_code,asn\n"+row), "network", "country_code")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := r.Read(); err == nil {
			t.Errorf("accepted bad row %q", row)
		}
	}
}

func TestIPInfoCSVQuotedHeaderWithBOM(t *testing.T) {
	r, err := newIPInfoLiteCSVReader(strings.NewReader("\ufeff\"network\",country_code\n1.0.0.0/24,AU\n"), "network", "country_code")
	if err != nil {
		t.Fatal(err)
	}
	if record, err := r.Read(); err != nil || record.CountryCode != "AU" {
		t.Fatalf("%+v %v", record, err)
	}
}
