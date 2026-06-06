#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-/tmp/go-build-cache}"

echo "== aggregate audit =="
echo "GOCACHE=$GOCACHE"
echo

echo "-- go test ./... --"
go test ./...
echo

echo "-- go vet ./... --"
bash scripts/audit_go_vet.sh
echo

echo "-- format audit --"
bash scripts/audit_format.sh
echo

echo "-- audit ai rules --"
bash scripts/audit_ai_rules.sh
echo

echo "-- audit file size --"
bash scripts/audit_file_size.sh
echo

echo "-- audit hexagonal boundaries --"
bash scripts/audit_hexagonal.sh
echo

echo "-- audit security gosec --"
bash scripts/audit_security_gosec.sh
echo

echo "[PASS] aggregate audit passed"
