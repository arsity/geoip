package plaintext

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/Loyalsoldier/geoip/lib"
)

// A remote list such as the Tor exit list must survive a connection reset
// instead of failing the whole build.
func TestTextInRemoteRetriesConnectionReset(t *testing.T) {
	var requests atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requests.Add(1) == 1 {
			conn, _, err := w.(http.Hijacker).Hijack()
			if err != nil {
				t.Error(err)
				return
			}
			conn.Close()
			return
		}
		fmt.Fprint(w, "8.8.8.8\n")
	}))
	defer srv.Close()

	in, err := NewTextIn(lib.ActionAdd, WithNameAndURI("tor", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	container, err := in.Input(lib.NewContainer())
	if err != nil {
		t.Fatal(err)
	}
	entry, found := container.GetEntry("TOR")
	if !found {
		t.Fatal("entry TOR not found")
	}
	prefixes, err := entry.MarshalPrefix()
	if err != nil || len(prefixes) != 1 || prefixes[0].String() != "8.8.8.8/32" {
		t.Fatalf("prefixes=%v err=%v", prefixes, err)
	}
	if got := requests.Load(); got != 2 {
		t.Fatalf("requests=%d, want 2", got)
	}
}
