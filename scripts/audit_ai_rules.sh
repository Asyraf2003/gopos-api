#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

required_files=(
  "docs/README.md"
  "docs/AGENTS.md"
  "docs/0001_index.md"
  "docs/0002_decision_policy.md"
  "docs/0003_session_start_protocol.md"
  "docs/core/0010_scope_and_facts.md"
  "docs/core/0011_blueprint_first.md"
  "docs/core/0012_step_by_step_execution.md"
  "docs/core/0013_proof_and_progress.md"
  "docs/architecture/0020_hexagonal_go_api.md"
  "docs/architecture/0021_package_boundaries.md"
  "docs/architecture/0022_api_capability_control.md"
  "docs/architecture/0023_public_contracts.md"
  "docs/architecture/0024_current_repo_layout.md"
  "docs/domain/0030_domain_contracts.md"
  "docs/db/0040_postgresql_policy.md"
  "docs/api/0050_echo_http_contract.md"
  "docs/testing/0060_test_and_quality_gates.md"
  "docs/workflow/0070_docs_go_workflow.md"
  "docs/workflow/0071_handoff_protocol.md"
  "docs/security/0080_security_baseline.md"
  "docs/scripts/0090_makefile_and_scripts.md"
  "docs/style/0100_go_style.md"
  "docs/templates/0110_domain_scope_packet.md"
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
check_contains "docs/README.md" "First Read Order"
check_contains "docs/AGENTS.md" "canonical AI_RULES package"
check_contains "docs/0001_index.md" "Mandatory Read Order"
check_contains "docs/0001_index.md" "Constitution Summary"
check_contains "docs/0002_decision_policy.md" "Mandatory Decision Sequence"
check_contains "docs/0002_decision_policy.md" "GAP Rule"
check_contains "docs/0003_session_start_protocol.md" "Mandatory Opening Flow"
check_contains "docs/core/0011_blueprint_first.md" "Implementation Gate"
check_contains "docs/core/0012_step_by_step_execution.md" "An active step must have"
check_contains "docs/core/0013_proof_and_progress.md" "Accepted Proof"
check_contains "docs/architecture/0021_package_boundaries.md" "One Folder One Package"
check_contains "docs/architecture/0024_current_repo_layout.md" "Protected Contracts"
check_contains "docs/scripts/0090_makefile_and_scripts.md" "Required Make Targets"
check_contains "docs/style/0100_go_style.md" "Forbidden Patterns"
check_contains "docs/adr/0001-foundation-raw-go-echo-postgres-hexagonal.md" "## Decision"

echo
echo "[PASS] AI rules audit passed"
