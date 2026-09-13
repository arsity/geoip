package lib

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Stage the complete response on disk before exposing it to a converter. A
// retry after a mid-body reset must not append a second response to data that
// a converter has already consumed. Disk staging also keeps large CSVs out of RAM.
var defaultRemoteDownloader = remoteDownloader{
	client:   &http.Client{Timeout: 2 * time.Minute},
	attempts: 4,
	backoff:  time.Second,
}

type remoteDownloader struct {
	client   *http.Client
	attempts int
	backoff  time.Duration
	tempDir  string
}

func (d remoteDownloader) get(ctx context.Context, rawURL string) (io.ReadCloser, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()
	var lastErr error
	for attempt := 0; attempt < d.attempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		file, retry, retryAfter, err := d.once(ctx, rawURL)
		if err == nil {
			return file, nil
		}
		lastErr = err
		if !retry || attempt+1 == d.attempts {
			break
		}
		delay := retryDelay(retryAfter, d.backoff<<attempt, time.Now())
		log.Printf("retrying download %s (attempt %d/%d) in %s: %v", safeRemoteURL(rawURL), attempt+2, d.attempts, delay, err)
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
	return nil, fmt.Errorf("download %s failed: %w", safeRemoteURL(rawURL), lastErr)
}

func (d remoteDownloader) once(ctx context.Context, rawURL string) (io.ReadCloser, bool, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, false, "", errors.New("invalid download URL")
	}
	resp, err := d.client.Do(req)
	if err != nil {
		// net/url errors include the entire URL, which may contain a token.
		for {
			var urlErr *url.Error
			if !errors.As(err, &urlErr) {
				break
			}
			err = urlErr.Err
		}
		return nil, transientDownloadError(err), "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		retry := resp.StatusCode == 408 || resp.StatusCode == 429 ||
			resp.StatusCode == 500 || resp.StatusCode == 502 ||
			resp.StatusCode == 503 || resp.StatusCode == 504
		return nil, retry, resp.Header.Get("Retry-After"), fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	file, err := os.CreateTemp(d.tempDir, "geoip-download-*")
	if err != nil {
		return nil, false, "", err
	}
	staged := &temporaryDownload{File: file}
	if _, err := io.Copy(file, resp.Body); err != nil {
		staged.Close()
		return nil, transientDownloadError(err), "", err
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		staged.Close()
		return nil, false, "", err
	}
	return staged, false, "", nil
}

func transientDownloadError(err error) bool {
	if errors.Is(err, context.Canceled) {
		return false
	}
	var networkError net.Error
	return (errors.As(err, &networkError) && (networkError.Timeout() || networkError.Temporary())) ||
		errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) ||
		errors.Is(err, syscall.ECONNRESET) || errors.Is(err, syscall.ECONNREFUSED) ||
		errors.Is(err, syscall.EPIPE)
}

func retryDelay(header string, fallback time.Duration, now time.Time) time.Duration {
	const maxDelay = 30 * time.Second
	delay := fallback
	if seconds, err := strconv.ParseInt(header, 10, 64); err == nil && seconds >= 0 {
		// Clamp before multiplying to avoid overflow on an untrusted header.
		if seconds > 30 {
			seconds = 30
		}
		delay = max(delay, time.Duration(seconds)*time.Second)
	} else if when, err := http.ParseTime(header); err == nil {
		delay = max(delay, when.Sub(now))
	}
	return min(maxDelay, max(time.Duration(0), delay))
}

func safeRemoteURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return "<invalid URL>"
	}
	u.User, u.RawQuery, u.Fragment = nil, "", ""
	return u.String()
}

type temporaryDownload struct {
	*os.File
	once sync.Once
	err  error
}

func (f *temporaryDownload) Close() error {
	f.once.Do(func() {
		f.err = errors.Join(f.File.Close(), os.Remove(f.Name()))
	})
	return f.err
}
