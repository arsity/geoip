"""Offline download tests: only a local HTTP server and a fake token are used."""
import collections
import gzip
import hashlib
import http.server
import json
import os
from pathlib import Path
import subprocess
import tempfile
import threading
import unittest
import urllib.parse

SCRIPT = Path(__file__).with_name("download-ipinfo.sh").resolve()


class DownloadTests(unittest.TestCase):
    def exercise(self, mode):
        payloads = {
            "ipinfo_lite.csv.gz": gzip.compress(b"network,country_code\n1.0.0.0/24,AU\n"),
            "ipinfo_lite.mmdb": b"synthetic bytes; MMDB parsing is covered in Go tests",
        }
        if mode == "bad_gzip":
            payloads["ipinfo_lite.csv.gz"] = b"not gzip"
        counts = collections.Counter()

        class Handler(http.server.BaseHTTPRequestHandler):
            def do_GET(self):
                path = urllib.parse.urlsplit(self.path).path.removeprefix("/data/")
                counts[path] += 1
                if mode == "transient" and counts[path] == 1:
                    self.send_response(503)
                    self.end_headers()
                    return
                if path.endswith("/checksums"):
                    name = path.removesuffix("/checksums")
                    digest = hashlib.sha256(payloads[name]).hexdigest()
                    if mode == "mismatch" and name.endswith("mmdb"):
                        digest = "0" * 64
                    if mode == "rollover" and counts[path] == 1:
                        digest = "0" * 64
                    if mode == "bad_checksum":
                        digest = "invalid"
                    body = json.dumps({"checksums": {"sha256": digest}}).encode()
                else:
                    body = payloads[path]
                self.send_response(200)
                self.send_header("Content-Length", str(len(body)))
                self.end_headers()
                self.wfile.write(body)

            def log_message(self, *args):
                pass

        server = http.server.ThreadingHTTPServer(("127.0.0.1", 0), Handler)
        thread = threading.Thread(target=server.serve_forever, daemon=True)
        thread.start()
        try:
            with tempfile.TemporaryDirectory() as directory:
                output = Path(directory)
                for name in payloads:
                    (output / name).write_bytes(b"previous")
                env = dict(os.environ, IPINFO_TOKEN="geoip-test-token",
                           IPINFO_DOWNLOAD_BASE_URL=f"http://127.0.0.1:{server.server_port}/data")
                result = subprocess.run(["bash", str(SCRIPT), str(output)], env=env,
                                        capture_output=True, text=True, timeout=30)
                self.assertNotIn("geoip-test-token", result.stdout + result.stderr)
                success = mode in ("success", "rollover", "transient")
                self.assertEqual(result.returncode == 0, success, result.stdout + result.stderr)
                for name, payload in payloads.items():
                    self.assertEqual((output / name).read_bytes(), payload if success else b"previous")
                self.assertEqual(sorted(p.name for p in output.iterdir()), sorted(payloads))
                if mode == "rollover":
                    self.assertEqual(counts["ipinfo_lite.csv.gz"], 1)
                    self.assertEqual(counts["ipinfo_lite.mmdb"], 1)
        finally:
            server.shutdown()
            thread.join()
            server.server_close()

    def test_valid_download(self):
        self.exercise("success")

    def test_transient_http_error(self):
        self.exercise("transient")

    def test_rollover_does_not_redownload(self):
        self.exercise("rollover")

    def test_mismatch_fails_closed(self):
        self.exercise("mismatch")

    def test_invalid_checksum_fails_closed(self):
        self.exercise("bad_checksum")

    def test_bad_gzip_fails_closed(self):
        self.exercise("bad_gzip")

    def test_endpoint_override_rejected(self):
        env = dict(os.environ, IPINFO_TOKEN="not-a-real-token",
                   IPINFO_DOWNLOAD_BASE_URL="https://example.invalid/data")
        result = subprocess.run(["bash", str(SCRIPT)], env=env,
                                capture_output=True, text=True, timeout=5)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("loopback test server", result.stderr)
        self.assertNotIn("not-a-real-token", result.stderr)


if __name__ == "__main__":
    unittest.main()
