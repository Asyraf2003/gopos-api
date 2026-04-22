#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

echo "== aggregate audit =="
echo

echo "-- go test ./... --"
go test ./...
echo

echo "-- audit ai rules --"
bash scripts/audit_ai_rules.sh
echo

echo "-- audit file size --"
bash scripts/audit_file_size.sh
echo

echo "-- audit security gosec --"
bash scripts/audit_security_gosec.sh
echo

echo "[PASS] aggregate audit passed"
