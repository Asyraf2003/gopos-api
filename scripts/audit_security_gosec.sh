#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

GOSEC_BIN="${GOSEC_BIN:-/home/asyraf/go/bin/gosec}"

if [[ ! -x "$GOSEC_BIN" ]]; then
  echo "[FAIL] gosec binary not found or not executable: $GOSEC_BIN"
  exit 1
fi

echo "== security audit: gosec =="
echo "binary: $GOSEC_BIN"
echo

"$GOSEC_BIN" ./...

echo
echo "[PASS] gosec audit passed"
