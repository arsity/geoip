package lib

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

func TestRemoteDownloadRetries(t *testing.T) {
	for _, kind := range []string{"503", "429", "reset", "truncated"} {
		t.Run(kind, func(t *testing.T) {
			var requests atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if requests.Add(1) == 1 {
					switch kind {
					case "503":
						w.WriteHeader(503)
					case "429":
						w.Header().Set("Retry-After", "0")
						w.WriteHeader(429)
					case "reset":
						conn, _, err := w.(http.Hijacker).Hijack()
						if err != nil {
							t.Error(err)
							return
						}
						conn.Close()
					case "truncated":
						w.Header().Set("Content-Length", "100")
						fmt.Fprint(w, "partial")
					}
					return
				}
				fmt.Fprint(w, "complete\n")
			}))
			defer srv.Close()
			d := remoteDownloader{client: srv.Client(), attempts: 3, tempDir: t.TempDir()}
			reader, err := d.get(context.Background(), srv.URL)
			if err != nil {
				t.Fatal(err)
			}
			defer reader.Close()
			data, err := io.ReadAll(reader)
			if err != nil || string(data) != "complete\n" || requests.Load() != 2 {
				t.Fatalf("body=%q requests=%d err=%v", data, requests.Load(), err)
			}
			if err := reader.Close(); err != nil {
				t.Fatal(err)
			}
			files, err := os.ReadDir(d.tempDir)
			if err != nil || len(files) != 0 {
				t.Fatalf("temporary files leaked: %v %v", files, err)
			}
		})
	}
}

func TestRemoteDownloadFailure(t *testing.T) {
	for _, status := range []int{401, 403, 404, 500, 503} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var requests atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				w.WriteHeader(status)
			}))
			defer srv.Close()
			d := remoteDownloader{client: srv.Client(), attempts: 3, tempDir: t.TempDir()}
			reader, err := d.get(context.Background(), srv.URL+"?token=do-not-log")
			if reader != nil || err == nil {
				t.Fatalf("reader=%v err=%v", reader, err)
			}
			if strings.Contains(err.Error(), "do-not-log") {
				t.Fatal("credential leaked")
			}
			want := int32(1)
			if status >= 500 {
				want = 3
			}
			if requests.Load() != want {
				t.Fatalf("requests=%d, want %d", requests.Load(), want)
			}
		})
	}
}

func TestRemoteDownloadTimeoutAndCancellation(t *testing.T) {
	for _, bodyTimeout := range []bool{false, true} {
		t.Run(fmt.Sprint(bodyTimeout), func(t *testing.T) {
			var requests atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				if bodyTimeout {
					w.(http.Flusher).Flush()
				}
				<-r.Context().Done()
			}))
			defer srv.Close()
			client := srv.Client()
			client.Timeout = 100 * time.Millisecond
			d := remoteDownloader{client: client, attempts: 2, tempDir: t.TempDir()}
			if _, err := d.get(context.Background(), srv.URL); err == nil {
				t.Fatal("expected timeout")
			}
			if requests.Load() != 2 {
				t.Fatalf("requests=%d", requests.Load())
			}
			files, _ := os.ReadDir(d.tempDir)
			if len(files) != 0 {
				t.Fatal("partial response leaked")
			}
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			if _, err := d.get(ctx, srv.URL); !errors.Is(err, context.Canceled) {
				t.Fatalf("err=%v", err)
			}
			if requests.Load() != 2 {
				t.Fatal("canceled request was sent")
			}
		})
	}
}

func TestRemoteRetryPolicy(t *testing.T) {
	for _, err := range []error{io.EOF, io.ErrUnexpectedEOF, syscall.ECONNRESET, syscall.ECONNREFUSED, context.DeadlineExceeded} {
		if !transientDownloadError(err) {
			t.Errorf("should retry %v", err)
		}
	}
	for _, err := range []error{context.Canceled, x509.UnknownAuthorityError{}, os.ErrPermission} {
		if transientDownloadError(err) {
			t.Errorf("should not retry %v", err)
		}
	}
	now := time.Now().UTC().Truncate(time.Second)
	for _, tc := range []struct {
		header string
		want   time.Duration
	}{
		{"", time.Second}, {"invalid", time.Second}, {"-1", time.Second}, {"2", 2 * time.Second},
		{"9223372036854775807", 30 * time.Second},
		{now.Add(5 * time.Second).Format(http.TimeFormat), 5 * time.Second},
		{now.Add(time.Hour).Format(http.TimeFormat), 30 * time.Second},
	} {
		if got := retryDelay(tc.header, time.Second, now); got != tc.want {
			t.Errorf("%q: %s != %s", tc.header, got, tc.want)
		}
	}
	if got := safeRemoteURL("https://user:password@example.com/data?token=secret#secret"); got != "https://example.com/data" {
		t.Errorf("unsafe URL: %s", got)
	}
}
