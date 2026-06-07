#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-/tmp/go-build-cache}"

passed_steps=()

run_step() {
  local name="$1"
  shift

  echo "-- $name --"
  "$@"
  passed_steps+=("$name")
  echo
}

echo "== aggregate audit =="
echo "GOCACHE=$GOCACHE"
echo

run_step "go test ./..." go test ./...
run_step "go vet audit" bash scripts/audit_go_vet.sh
run_step "format audit" bash scripts/audit_format.sh
run_step "AI rules audit" bash scripts/audit_ai_rules.sh
run_step "file size audit" bash scripts/audit_file_size.sh
run_step "hexagonal import audit" bash scripts/audit_hexagonal.sh
run_step "route capability audit" bash scripts/audit_route_capabilities.sh
run_step "security gosec audit" bash scripts/audit_security_gosec.sh

echo "== aggregate audit summary =="
for step in "${passed_steps[@]}"; do
  printf '[PASS] %s\n' "$step"
done
echo
echo "[PASS] aggregate audit passed"
