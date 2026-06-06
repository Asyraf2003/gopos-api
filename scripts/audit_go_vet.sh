#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-/tmp/go-build-cache}"

echo "== go vet audit =="
echo "GOCACHE=$GOCACHE"
echo

go vet ./...

echo
echo "[PASS] go vet audit passed"
