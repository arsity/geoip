package maxmind

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/netip"
	"strconv"
	"strings"
)

// Names, not positions, define the schema. Unknown columns are permitted;
// missing/duplicate headers and malformed records fail the entire conversion.
type ipinfoLiteCSVReader struct {
	reader       *csv.Reader
	columns      map[string]int
	rows         int
	emptyCountry int
	emptyASN     int
}

type ipinfoLiteCSVRecord struct {
	Network, CountryCode, ASN, ASName, ASDomain string
}

func newIPInfoLiteCSVReader(r io.Reader, required ...string) (*ipinfoLiteCSVReader, error) {
	buffered := bufio.NewReader(r)
	if prefix, _ := buffered.Peek(3); string(prefix) == "\ufeff" {
		_, _ = buffered.Discard(3)
	}
	reader := csv.NewReader(buffered)
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("IPinfo CSV header: %w", err)
	}
	columns := make(map[string]int, len(header))
	for i, name := range header {
		name = strings.ToLower(strings.TrimSpace(strings.TrimPrefix(name, "\ufeff")))
		if name == "" {
			return nil, fmt.Errorf("IPinfo CSV: empty header at column %d", i+1)
		}
		if _, found := columns[name]; found {
			return nil, fmt.Errorf("IPinfo CSV: duplicate header %q", name)
		}
		columns[name] = i
	}
	for _, name := range required {
		if _, found := columns[name]; !found {
			return nil, fmt.Errorf("IPinfo CSV: missing required header %q", name)
		}
	}
	// Explicitly require every row to have the same number of fields as the header.
	reader.FieldsPerRecord = len(header)
	return &ipinfoLiteCSVReader{reader: reader, columns: columns}, nil
}

func (r *ipinfoLiteCSVReader) Read() (ipinfoLiteCSVRecord, error) {
	fields, err := r.reader.Read()
	if err != nil {
		if err == io.EOF {
			return ipinfoLiteCSVRecord{}, io.EOF
		}
		return ipinfoLiteCSVRecord{}, fmt.Errorf("IPinfo CSV record %d: %w", r.rows+1, err)
	}
	r.rows++
	field := func(name string) string {
		if i, ok := r.columns[name]; ok {
			return strings.TrimSpace(fields[i])
		}
		return ""
	}
	record := ipinfoLiteCSVRecord{
		Network: field("network"), CountryCode: strings.ToUpper(field("country_code")),
		ASN: field("asn"), ASName: field("as_name"), ASDomain: field("as_domain"),
	}
	if _, err := netip.ParsePrefix(record.Network); err != nil {
		// IPinfo's network field also allows a single IP address.
		if _, err := netip.ParseAddr(record.Network); err != nil {
			return record, fmt.Errorf("IPinfo CSV record %d: invalid network %q", r.rows, record.Network)
		}
	}
	if record.CountryCode == "" {
		r.emptyCountry++
	} else if code := record.CountryCode; len(code) != 2 || code[0] < 'A' || code[0] > 'Z' || code[1] < 'A' || code[1] > 'Z' {
		return record, fmt.Errorf("IPinfo CSV record %d: invalid country_code %q", r.rows, code)
	}
	if record.ASN == "" {
		r.emptyASN++
	} else {
		asn := strings.TrimPrefix(strings.ToUpper(record.ASN), "AS")
		if _, err := strconv.ParseUint(asn, 10, 32); err != nil {
			return record, fmt.Errorf("IPinfo CSV record %d: invalid ASN %q", r.rows, record.ASN)
		}
	}
	return record, nil
}

func (r *ipinfoLiteCSVReader) report(converter string) {
	log.Printf("IPinfo CSV [%s]: records=%d empty_country=%d empty_asn=%d", converter, r.rows, r.emptyCountry, r.emptyASN)
}
