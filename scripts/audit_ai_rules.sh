#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

required_files=(
  "docs/README.md"
  "docs/core/BLUEPRINT.md"
  "docs/core/STRUCTURE.md"
  "docs/core/AI_RULES.md"
  "docs/core/WORKFLOW.md"
  "docs/core/DOD.md"
  "docs/AI_RULES/00_INDEX.md"
  "docs/AI_RULES/01_DECISION_POLICY.md"
  "docs/AI_RULES/10_CORE/11_BLUEPRINT_FIRST.md"
  "docs/AI_RULES/10_CORE/12_STEP_BY_STEP_EXECUTION.md"
  "docs/AI_RULES/10_CORE/13_PROOF_AND_PROGRESS.md"
  "docs/AI_RULES/40_ARCHITECTURE/44_AUDIT_AND_DOD.md"
  "docs/AI_RULES/60_STACK/61_GO_RULES.md"
  "docs/adr/0001-foundation-raw-go-echo-postgres-hexagonal.md"
)

check_file() {
  local path="$1"
  if [[ ! -f "$path" ]]; then
    echo "[FAIL] missing file: $path"
    exit 1
  fi
  echo "[OK] file exists: $path"
}

check_contains() {
  local path="$1"
  local needle="$2"
  if ! grep -Fq "$needle" "$path"; then
    echo "[FAIL] missing text in $path :: $needle"
    exit 1
  fi
  echo "[OK] text found in $path :: $needle"
}

echo "== file existence =="
for f in "${required_files[@]}"; do
  check_file "$f"
done

echo
echo "== content checks =="
check_contains "docs/README.md" "Urutan rujukan"
check_contains "docs/core/BLUEPRINT.md" "Tujuan utama"
check_contains "docs/core/STRUCTURE.md" "Contracts antar layer"
check_contains "docs/core/AI_RULES.md" "Mandatory operational reference"
check_contains "docs/core/WORKFLOW.md" "Workflow default"
check_contains "docs/core/DOD.md" "Done minimum"
check_contains "docs/AI_RULES/00_INDEX.md" "Mandatory Read Order"
check_contains "docs/AI_RULES/00_INDEX.md" "Constitution Summary"
check_contains "docs/AI_RULES/01_DECISION_POLICY.md" "Mandatory Decision Sequence"
check_contains "docs/AI_RULES/01_DECISION_POLICY.md" "GAP Rule"
check_contains "docs/AI_RULES/10_CORE/11_BLUEPRINT_FIRST.md" "Implementation Gate"
check_contains "docs/AI_RULES/10_CORE/12_STEP_BY_STEP_EXECUTION.md" "Definition of Active Step"
check_contains "docs/AI_RULES/10_CORE/13_PROOF_AND_PROGRESS.md" "Accepted Proof"
check_contains "docs/AI_RULES/40_ARCHITECTURE/44_AUDIT_AND_DOD.md" "Typical DoD Components"
check_contains "docs/AI_RULES/60_STACK/61_GO_RULES.md" "Satu folder = satu package"
check_contains "docs/adr/0001-foundation-raw-go-echo-postgres-hexagonal.md" "## Decision"

echo
echo "[PASS] AI rules audit passed"
