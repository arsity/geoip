package lib

import (
	"context"
	"encoding/json"
	"io"
)

func GetRemoteURLContent(url string) ([]byte, error) {
	reader, err := GetRemoteURLReader(url)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

// GetRemoteURLReader returns a complete, retried download. Callers must close
// the reader to release and remove its temporary backing file.
func GetRemoteURLReader(url string) (io.ReadCloser, error) {
	return defaultRemoteDownloader.get(context.Background(), url)
}

func GetIgnoreIPType(onlyIPType IPType) IgnoreIPOption {
	switch onlyIPType {
	case IPv4:
		return IgnoreIPv6
	case IPv6:
		return IgnoreIPv4
	}

	return nil
}

type WantedListExtended struct {
	TypeSlice []string
	TypeMap   map[string][]string
}

func (w *WantedListExtended) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	slice := make([]string, 0)
	mapMap := make(map[string][]string, 0)

	err := json.Unmarshal(data, &slice)
	if err != nil {
		err2 := json.Unmarshal(data, &mapMap)
		if err2 != nil {
			return err2
		}
	}

	w.TypeSlice = slice
	w.TypeMap = mapMap

	return nil
}
