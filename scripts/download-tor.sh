#!/usr/bin/env bash
set -euo pipefail

# check.torproject.org intermittently resets every connection from some GitHub
# Actions runners for the whole job, so retrying in-process is not enough.
# Fall back to the list this repository published in its last successful build.
# A list that is a day old is far better than skipping the whole release.
primary="${TOR_EXIT_LIST_URL:-https://check.torproject.org/torbulkexitlist}"
fallback="${TOR_FALLBACK_URL:-https://raw.githubusercontent.com/${GITHUB_REPOSITORY:-arsity/geoip}/release/text/tor.txt}"
retry_delay="${TOR_RETRY_DELAY:-10}"
min_entries="${TOR_MIN_ENTRIES:-100}"
protocols='=https'
# Plain HTTP exists only for offline tests against a loopback server.
for url in "$primary" "$fallback"; do
  if [[ "$url" =~ ^http://127\.0\.0\.1:[0-9]+/ ]]; then
    protocols='=http,https'
  elif [[ "$url" != https://* ]]; then
    echo "Refusing non-HTTPS Tor list URL" >&2
    exit 1
  fi
done
output="${1:-./tor}"
mkdir -p "$output"
staging="$(mktemp -d "$output/.download.XXXXXX")"
trap 'rm -rf "$staging"' EXIT

fetch() {
  curl --fail --silent --show-error --location \
    --proto "$protocols" --proto-redir "$protocols" \
    --retry 5 --retry-all-errors --retry-delay "$retry_delay" --retry-max-time 180 \
    --connect-timeout 15 --max-time 60 --remove-on-error \
    "$1" -o "$2"
}

# Reject error pages and truncated bodies: every line must look like an IP or
# CIDR, and a real exit list has well over a hundred entries.
valid() {
  local count
  count="$(grep -cv '^[[:space:]]*$' "$1" || true)"
  if ((count < min_entries)); then
    echo "Tor list has only $count entries" >&2
    return 1
  fi
  if grep -Evq '^[[:space:]]*$|^([0-9]{1,3}(\.[0-9]{1,3}){3}|[0-9A-Fa-f]*:[0-9A-Fa-f:.]*)(/[0-9]{1,3})?$' "$1"; then
    echo "Tor list contains a line that is not an IP or CIDR" >&2
    return 1
  fi
}

if fetch "$primary" "$staging/list" && valid "$staging/list"; then
  echo "Downloaded Tor exit list: $(grep -cv '^[[:space:]]*$' "$staging/list") entries"
else
  echo "::warning::Tor exit list download failed; reusing the list from the last release"
  rm -f "$staging/list"
  fetch "$fallback" "$staging/list"
  valid "$staging/list"
  echo "Reused previous Tor exit list: $(grep -cv '^[[:space:]]*$' "$staging/list") entries"
fi
mv -f "$staging/list" "$output/torbulkexitlist"
