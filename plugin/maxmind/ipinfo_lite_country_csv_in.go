package maxmind

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Loyalsoldier/geoip/lib"
)

const (
	TypeIPInfoLiteCountryCSVIn = "ipinfoLiteCountryCSV"
	DescIPInfoLiteCountryCSVIn = "Convert IPInfo Lite country CSV data to other formats"
)

var (
	defaultIPInfoLiteCountryCSVFile = filepath.Join("./", "ipinfo", "ipinfo_lite.csv")
)

func init() {
	lib.RegisterInputConfigCreator(TypeIPInfoLiteCountryCSVIn, func(action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
		return newIPInfoLiteCountryCSVIn(action, data)
	})
	lib.RegisterInputConverter(TypeIPInfoLiteCountryCSVIn, &IPInfoLiteCountryCSVIn{
		Description: DescIPInfoLiteCountryCSVIn,
	})
}

func newIPInfoLiteCountryCSVIn(action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
	var tmp struct {
		URI        string     `json:"uri"`
		Want       []string   `json:"wantedList"`
		OnlyIPType lib.IPType `json:"onlyIPType"`
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tmp); err != nil {
			return nil, err
		}
	}

	if tmp.URI == "" {
		tmp.URI = defaultIPInfoLiteCountryCSVFile
	}

	// Filter want list
	wantList := make(map[string]bool)
	for _, want := range tmp.Want {
		if want = strings.ToUpper(strings.TrimSpace(want)); want != "" {
			wantList[want] = true
		}
	}

	return &IPInfoLiteCountryCSVIn{
		Type:        TypeIPInfoLiteCountryCSVIn,
		Action:      action,
		Description: DescIPInfoLiteCountryCSVIn,
		URI:         tmp.URI,
		Want:        wantList,
		OnlyIPType:  tmp.OnlyIPType,
	}, nil
}

type IPInfoLiteCountryCSVIn struct {
	Type        string
	Action      lib.Action
	Description string
	URI         string
	Want        map[string]bool
	OnlyIPType  lib.IPType
}

func (g *IPInfoLiteCountryCSVIn) GetType() string {
	return g.Type
}

func (g *IPInfoLiteCountryCSVIn) GetAction() lib.Action {
	return g.Action
}

func (g *IPInfoLiteCountryCSVIn) GetDescription() string {
	return g.Description
}

func (g *IPInfoLiteCountryCSVIn) Input(container lib.Container) (lib.Container, error) {
	entries := make(map[string]*lib.Entry)

	if err := g.process(g.URI, entries); err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("❌ [type %s | action %s] no entry is generated", g.Type, g.Action)
	}

	ignoreIPType := lib.GetIgnoreIPType(g.OnlyIPType)

	for _, entry := range entries {
		switch g.Action {
		case lib.ActionAdd:
			if err := container.Add(entry, ignoreIPType); err != nil {
				return nil, err
			}
		case lib.ActionRemove:
			if err := container.Remove(entry, lib.CaseRemovePrefix, ignoreIPType); err != nil {
				return nil, err
			}
		default:
			return nil, lib.ErrUnknownAction
		}
	}

	return container, nil
}

func (g *IPInfoLiteCountryCSVIn) process(file string, entries map[string]*lib.Entry) error {
	if entries == nil {
		entries = make(map[string]*lib.Entry)
	}

	var f io.ReadCloser
	var err error
	switch {
	case strings.HasPrefix(strings.ToLower(file), "http://"), strings.HasPrefix(strings.ToLower(file), "https://"):
		f, err = lib.GetRemoteURLReader(file)
	default:
		f, err = os.Open(file)
	}

	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Read() // skip header

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// IPInfo Lite CSV reference:
		// network,country,country_code,continent,continent_code,asn,as_name,as_domain
		// 1.0.0.0/24,Australia,AU,Oceania,OC,AS13335,Cloudflare Inc,cloudflare.com

		if len(record) < 3 {
			return fmt.Errorf("❌ [type %s | action %s] invalid record: %v", g.Type, g.Action, record)
		}

		countryCode := strings.ToUpper(strings.TrimSpace(record[2]))
		if countryCode == "" {
			continue
		}

		if len(g.Want) > 0 && !g.Want[countryCode] {
			continue
		}

		cidrStr := strings.ToLower(strings.TrimSpace(record[0]))
		entry, got := entries[countryCode]
		if !got {
			entry = lib.NewEntry(countryCode)
		}

		if err := entry.AddPrefix(cidrStr); err != nil {
			return err
		}

		entries[countryCode] = entry
	}

	return nil
}
