#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

echo "== format audit =="
echo "tool: gofmt -l"
echo

mapfile -t files < <(gofmt -l $(fd -e go .))

if (( ${#files[@]} > 0 )); then
  echo "[FAIL] go files need gofmt:"
  printf '  %s\n' "${files[@]}"
  exit 1
fi

echo "[PASS] format audit passed"
