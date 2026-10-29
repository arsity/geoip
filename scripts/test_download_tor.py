"""Offline Tor list download tests: only a local HTTP server is used."""
import collections
import http.server
import os
from pathlib import Path
import subprocess
import tempfile
import threading
import unittest

SCRIPT = Path(__file__).with_name("download-tor.sh").resolve()
PRIMARY = "".join(f"10.0.{i // 256}.{i % 256}\n" for i in range(150)).encode() + b"2001:db8::1\n"
FALLBACK = "".join(f"192.0.2.{i}/32\n" for i in range(120)).encode()


class DownloadTests(unittest.TestCase):
    def exercise(self, primary, fallback=FALLBACK, transient=False):
        counts = collections.Counter()

        class Handler(http.server.BaseHTTPRequestHandler):
            def do_GET(self):
                counts[self.path] += 1
                body = primary if self.path == "/primary" else fallback
                if body is None or (transient and self.path == "/primary" and counts[self.path] == 1):
                    self.send_response(503)
                    self.end_headers()
                    return
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
                (output / "torbulkexitlist").write_bytes(b"previous")
                base = f"http://127.0.0.1:{server.server_port}"
                env = dict(os.environ, TOR_EXIT_LIST_URL=f"{base}/primary",
                           TOR_FALLBACK_URL=f"{base}/fallback", TOR_RETRY_DELAY="1")
                result = subprocess.run(["bash", str(SCRIPT), str(output)], env=env,
                                        capture_output=True, text=True, timeout=60)
                content = (output / "torbulkexitlist").read_bytes()
                self.assertEqual(sorted(p.name for p in output.iterdir()), ["torbulkexitlist"])
                return result, content, counts
        finally:
            server.shutdown()
            thread.join()
            server.server_close()

    def test_primary(self):
        result, content, counts = self.exercise(PRIMARY)
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual(content, PRIMARY)
        self.assertEqual(counts["/fallback"], 0)

    def test_transient_primary_error_is_retried(self):
        result, content, counts = self.exercise(PRIMARY, transient=True)
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual(content, PRIMARY)
        self.assertEqual(counts["/primary"], 2)

    def test_unreachable_primary_uses_fallback(self):
        result, content, _ = self.exercise(None)
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual(content, FALLBACK)
        self.assertIn("::warning::", result.stdout)

    def test_invalid_primary_uses_fallback(self):
        for body in (b"<html>blocked</html>\n" * 200, b"10.0.0.1\n"):
            result, content, _ = self.exercise(body)
            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertEqual(content, FALLBACK)

    def test_both_fail_keeps_previous_file_and_fails(self):
        for fallback in (None, b"<html>\n" * 200):
            result, content, _ = self.exercise(None, fallback)
            self.assertNotEqual(result.returncode, 0)
            self.assertEqual(content, b"previous")

    def test_non_https_url_rejected(self):
        env = dict(os.environ, TOR_EXIT_LIST_URL="http://example.invalid/list")
        with tempfile.TemporaryDirectory() as directory:
            result = subprocess.run(["bash", str(SCRIPT), directory], env=env,
                                    capture_output=True, text=True, timeout=5)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("non-HTTPS", result.stderr)


if __name__ == "__main__":
    unittest.main()
