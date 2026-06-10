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
  "docs/workflow/0072_transition_progress_ledger_protocol.md"
  "docs/security/0080_security_baseline.md"
  "docs/scripts/0090_makefile_and_scripts.md"
  "docs/style/0100_go_style.md"
  "docs/templates/0110_domain_scope_packet.md"
  "docs/templates/0120_prompt_authoring_rules.md"
  "docs/templates/0121_codex_session_prompts.md"
  "docs/templates/0122_web_ai_session_prompts.md"
  "docs/templates/0123_analysis_and_review_prompts.md"
  "docs/templates/0124_testing_and_proof_prompts.md"
  "docs/templates/0125_data_capture_and_evidence_prompts.md"
  "docs/templates/0126_resume_after_pause_prompts.md"
  "docs/adr/0001-foundation-raw-go-echo-postgres-hexagonal.md"
  "docs/evidence/0003_laravel_to_go_transition_progress_ledger.md"
  "scripts/audit_format.sh"
  "scripts/audit_go_vet.sh"
  "scripts/audit_hexagonal.sh"
)

required_readme_dirs=(
  "docs"
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
echo "== folder readmes =="
for d in "${required_readme_dirs[@]}"; do
  check_file "$d/README.md"
done

while IFS= read -r d; do
  check_file "${d%/}/README.md"
done < <(fd . docs -t d)

echo
echo "== content checks =="
check_contains "docs/README.md" "First Read Order"
check_contains "docs/AGENTS.md" "canonical AI_RULES package"
check_contains "docs/0001_index.md" "Mandatory Read Order"
check_contains "docs/0001_index.md" "Constitution Summary"
check_contains "docs/0002_decision_policy.md" "Mandatory Decision Sequence"
check_contains "docs/0002_decision_policy.md" "GAP Rule"
check_contains "docs/0002_decision_policy.md" "Decision And Data Request Rule"
check_contains "docs/0003_session_start_protocol.md" "Mandatory Opening Flow"
check_contains "docs/core/0011_blueprint_first.md" "Implementation Gate"
check_contains "docs/core/0012_step_by_step_execution.md" "An active step must have"
check_contains "docs/core/0013_proof_and_progress.md" "Accepted Proof"
check_contains "docs/architecture/0021_package_boundaries.md" "One Folder One Package"
check_contains "docs/architecture/0024_current_repo_layout.md" "Protected Contracts"
check_contains "docs/workflow/0070_docs_go_workflow.md" 'Every folder under `docs/` should have a `README.md`'
check_contains "docs/scripts/0090_makefile_and_scripts.md" "Required Make Targets"
check_contains "docs/style/0100_go_style.md" "Forbidden Patterns"
check_contains "docs/templates/0120_prompt_authoring_rules.md" "Web AI Source Rules"
check_contains "docs/templates/0120_prompt_authoring_rules.md" "Working Style Rule"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "GitHub connector access is read-only by default."
check_contains "docs/templates/0122_web_ai_session_prompts.md" "PROOF THE TERMINAL AGENT MUST RUN"
check_contains "docs/templates/0123_analysis_and_review_prompts.md" "2-3 owner decision options with tradeoffs"
check_contains "docs/templates/0125_data_capture_and_evidence_prompts.md" "smallest specific missing source batch"
check_contains "docs/templates/0120_prompt_authoring_rules.md" "Treat \"write docs/...\", \"update docs/...\", \"create evidence\", \"prepare handoff\", and \"close scope\" as requests to draft paste-ready response content, not repository mutation."
check_contains "docs/README.md" "Web AI connector access is read-only by default"
check_contains "docs/workflow/0071_handoff_protocol.md" "Automatic Handoff Triggers"
check_contains "docs/workflow/0072_transition_progress_ledger_protocol.md" "Progress may increase only from proof."
check_contains "docs/templates/0122_web_ai_session_prompts.md" "Progress Write Gate"
check_contains "docs/workflow/0072_transition_progress_ledger_protocol.md" "Progress Update Gate"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "locally implemented with proof; connector validation pending"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "Local proof and remote connector proof are not conflated"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "Git mutation instructions are absent unless the owner explicitly requested Git operations"
check_contains "docs/templates/0123_cli_command_formatter_rules.md" "CLI Command Formatter Rules"
check_contains "docs/templates/0123_cli_command_formatter_rules.md" 'Do not use `set -euo pipefail`'
check_contains "docs/templates/0123_cli_command_formatter_rules.md" 'Convert `grep` commands to `rg`'
check_contains "docs/templates/0123_cli_command_formatter_rules.md" 'Convert `find` commands to `fd`'
check_contains "docs/templates/0123_cli_command_formatter_rules.md" "Do not prepend this unless explicitly needed"
check_contains "docs/AGENTS.md" "docs/templates/0123_cli_command_formatter_rules.md"
check_contains "docs/0003_session_start_protocol.md" "docs/templates/0123_cli_command_formatter_rules.md"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "docs/templates/0123_cli_command_formatter_rules.md"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "NEXT does not skip required progress ledger or handoff updates"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "Web AI file update requests require COMMAND PLAN FOR OWNER / LOCAL TERMINAL"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "Paste-ready text must not replace the command plan"
check_contains "docs/templates/0122_web_ai_session_prompts.md" "PROGRESS and CONTEXT WINDOW STATUS are present for non-trivial work"
check_contains "docs/workflow/0071_handoff_protocol.md" "Do not let paste-ready handoff text replace a required owner/local command plan"
check_contains "docs/templates/0120_prompt_authoring_rules.md" "Web AI file update requests require COMMAND PLAN FOR OWNER / LOCAL TERMINAL"
check_contains "docs/README.md" "Progress Write Gate"
check_contains "docs/README.md" "locally implemented with proof; connector validation pending"
check_contains "docs/AGENTS.md" "Before NEXT, apply the Progress Write Gate"
check_contains "docs/AGENTS.md" "Local proof and remote connector proof must not be conflated"
check_contains "docs/0003_session_start_protocol.md" "Before naming NEXT, apply the Progress Write Gate"
check_contains "docs/0003_session_start_protocol.md" "durable proof changes progress"
check_contains "docs/evidence/0003_laravel_to_go_transition_progress_ledger.md" "Overall Laravel-to-Go transition"
check_contains "docs/README.md" "Documentation Cascade Rule"
check_contains "docs/AGENTS.md" "context-window status"
check_contains "scripts/audit_hexagonal.sh" "hexagonal import audit"
check_contains "scripts/audit_format.sh" "format audit"
check_contains "scripts/audit_go_vet.sh" "go vet audit"
check_contains "docs/adr/0001-foundation-raw-go-echo-postgres-hexagonal.md" "## Decision"

echo
echo "[PASS] AI rules audit passed"
