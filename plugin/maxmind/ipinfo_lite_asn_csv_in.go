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
	TypeIPInfoLiteASNCSVIn = "ipinfoLiteASNCSV"
	DescIPInfoLiteASNCSVIn = "Convert IPInfo Lite ASN CSV data to other formats"
)

var (
	defaultIPInfoLiteASNCSVFile = filepath.Join("./", "ipinfo", "ipinfo_lite.csv")
)

func init() {
	lib.RegisterInputConfigCreator(TypeIPInfoLiteASNCSVIn, func(action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
		return newIPInfoLiteASNCSVIn(action, data)
	})
	lib.RegisterInputConverter(TypeIPInfoLiteASNCSVIn, &IPInfoLiteASNCSVIn{
		Description: DescIPInfoLiteASNCSVIn,
	})
}

func newIPInfoLiteASNCSVIn(action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
	var tmp struct {
		URI        string                 `json:"uri"`
		Want       lib.WantedListExtended `json:"wantedList"`
		OnlyIPType lib.IPType             `json:"onlyIPType"`
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tmp); err != nil {
			return nil, err
		}
	}

	if tmp.URI == "" {
		tmp.URI = defaultIPInfoLiteASNCSVFile
	}

	// Filter want list
	wantList := make(map[string][]string) // map[asn][]listname or map[asn][]asn

	for list, asnList := range tmp.Want.TypeMap {
		list = strings.ToUpper(strings.TrimSpace(list))
		if list == "" {
			continue
		}

		for _, asn := range asnList {
			asn = strings.TrimPrefix(strings.ToLower(strings.TrimSpace(asn)), "as")
			if asn == "" {
				continue
			}

			if listArr, found := wantList[asn]; found {
				listArr = append(listArr, list)
				wantList[asn] = listArr
			} else {
				wantList[asn] = []string{list}
			}
		}
	}

	for _, asn := range tmp.Want.TypeSlice {
		asn = strings.TrimPrefix(strings.ToLower(strings.TrimSpace(asn)), "as")
		if asn == "" {
			continue
		}

		wantList[asn] = []string{"AS" + asn}
	}

	return &IPInfoLiteASNCSVIn{
		Type:        TypeIPInfoLiteASNCSVIn,
		Action:      action,
		Description: DescIPInfoLiteASNCSVIn,
		URI:         tmp.URI,
		Want:        wantList,
		OnlyIPType:  tmp.OnlyIPType,
	}, nil
}

type IPInfoLiteASNCSVIn struct {
	Type        string
	Action      lib.Action
	Description string
	URI         string
	Want        map[string][]string
	OnlyIPType  lib.IPType
}

func (g *IPInfoLiteASNCSVIn) GetType() string {
	return g.Type
}

func (g *IPInfoLiteASNCSVIn) GetAction() lib.Action {
	return g.Action
}

func (g *IPInfoLiteASNCSVIn) GetDescription() string {
	return g.Description
}

func (g *IPInfoLiteASNCSVIn) Input(container lib.Container) (lib.Container, error) {
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

func (g *IPInfoLiteASNCSVIn) process(file string, entries map[string]*lib.Entry) error {
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

		if len(record) < 6 {
			return fmt.Errorf("❌ [type %s | action %s] invalid record: %v", g.Type, g.Action, record)
		}

		// IPInfo already has "AS" prefix (e.g., "AS13335"); strip it for wantList lookup
		asnRaw := strings.TrimSpace(record[5])
		if asnRaw == "" {
			continue
		}
		asnBare := strings.TrimPrefix(strings.ToLower(asnRaw), "as")

		switch len(g.Want) {
		case 0: // it means user wants all ASNs
			asn := strings.ToUpper(asnRaw) // default list name is in "AS12345" format
			entry, got := entries[asn]
			if !got {
				entry = lib.NewEntry(asn)
			}
			if err := entry.AddPrefix(strings.TrimSpace(record[0])); err != nil {
				return err
			}
			entries[asn] = entry

		default: // it means user wants specific ASNs or customized lists with specific ASNs
			if listArr, found := g.Want[asnBare]; found {
				for _, listName := range listArr {
					entry, got := entries[listName]
					if !got {
						entry = lib.NewEntry(listName)
					}
					if err := entry.AddPrefix(strings.TrimSpace(record[0])); err != nil {
						return err
					}
					entries[listName] = entry
				}
			}
		}
	}

	return nil
}
