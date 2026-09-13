#!/usr/bin/env bash
set -euo pipefail

: "${IPINFO_TOKEN:?IPINFO_TOKEN must be set}"
base="${IPINFO_DOWNLOAD_BASE_URL:-https://ipinfo.io/data}"
protocols='=https'
# The alternate endpoint exists only for offline tests. Never send a real
# credential to an arbitrary URL supplied through the environment.
if [[ "$base" != 'https://ipinfo.io/data' ]]; then
  if [[ "$IPINFO_TOKEN" != 'geoip-test-token' || ! "$base" =~ ^http://127\.0\.0\.1:[0-9]+/data$ ]]; then
    echo 'Only a loopback test server with a fake test token may override the download endpoint' >&2
    exit 1
  fi
  protocols='=http,https'
fi
output="${1:-./ipinfo}"
mkdir -p "$output"
staging="$(mktemp -d "$output/.download.XXXXXX")"
trap 'rm -rf "$staging"' EXIT

fetch() {
  curl --fail --silent --show-error --location \
    --proto "$protocols" --proto-redir "$protocols" \
    --retry 3 --retry-all-errors --retry-max-time 180 \
    --connect-timeout 15 --max-time 120 --remove-on-error \
    "$1" -o "$2"
}

checksum() {
  fetch "$base/$1/checksums?token=$IPINFO_TOKEN" "$staging/checksums.json"
  jq -er '.checksums.sha256 | strings | select(test("^[0-9a-fA-F]{64}$")) | ascii_downcase' "$staging/checksums.json"
}

# Checksums describe the exact downloaded bytes, including gzip compression.
# Recheck checksum metadata after a mismatch to handle a daily data rollover.
for name in ipinfo_lite.csv.gz ipinfo_lite.mmdb; do
  verified=false
  for attempt in 1 2; do
    expected="$(checksum "$name")"
    fetch "$base/$name?token=$IPINFO_TOKEN" "$staging/$name"
    actual="$(sha256sum "$staging/$name" | cut -d ' ' -f 1)"
    if [[ "$actual" == "$expected" ]]; then
      verified=true
      break
    fi
    expected="$(checksum "$name")"
    if [[ "$actual" == "$expected" ]]; then
      verified=true
      break
    fi
    echo "Checksum mismatch for $name (attempt $attempt/2)" >&2
  done
  if [[ "$verified" != true ]]; then
    echo "Refusing unverified IPinfo database: $name" >&2
    exit 1
  fi
  echo "Verified SHA-256: $name"
done

gzip -t "$staging/ipinfo_lite.csv.gz"
# Publish neither file until both files have passed validation. There is no
# stale-data fallback; any failure exits nonzero and stops the build.
mv -f "$staging/ipinfo_lite.csv.gz" "$output/ipinfo_lite.csv.gz"
mv -f "$staging/ipinfo_lite.mmdb" "$output/ipinfo_lite.mmdb"
