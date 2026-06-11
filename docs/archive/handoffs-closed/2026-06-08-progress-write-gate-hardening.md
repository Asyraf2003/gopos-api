// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

// audit:allow-oversize reason=bootstrap-wiring

# Handoff: Progress Write Gate Hardening

## Date

2026-06-08

## Active Scope

AI workflow guardrail hardening after Web AI progress/write-gate incident.

## FACT

- Web AI read-only and proof-before-progress rules already existed.
- The missing control was a mandatory output gate before `NEXT`.
- This scope changes docs and audit guardrails only.
- Runtime application behavior is unchanged.

## Files Included

```text
docs/templates/0122_web_ai_session_prompts.md
docs/workflow/0072_transition_progress_ledger_protocol.md
docs/core/0013_proof_and_progress.md
docs/workflow/0071_handoff_protocol.md
scripts/audit_ai_rules.sh
```

## Files Changed

```text
docs/templates/0122_web_ai_session_prompts.md
docs/workflow/0072_transition_progress_ledger_protocol.md
docs/core/0013_proof_and_progress.md
docs/workflow/0071_handoff_protocol.md
scripts/audit_ai_rules.sh
docs/handoffs/2026-06-08-progress-write-gate-hardening.md
```

## Files Forbidden To Touch

```text
internal/**
cmd/**
migrations/**
scripts/config/route_capabilities.tsv
production secrets
GitHub branches, commits, pull requests, issues, labels, reviewers, merges, refs, or CI
```

## Decisions Made

- Add `Progress Write Gate` to the Web AI template.
- Add `Progress Update Gate` to the transition ledger protocol.
- Require local proof and remote connector proof to be separated.
- Require local-only status wording: `locally implemented with proof; connector validation pending`.
- Block `NEXT` from skipping required ledger or handoff updates after durable proof.
- Keep Git mutation commands out of Web AI output unless the owner explicitly requests Git operations.
- Add audit anchors for the new guardrails.

## PROOF

```text
bash scripts/audit_ai_rules.sh
```

Result:

```text
[OK] text found in docs/templates/0122_web_ai_session_prompts.md :: Progress Write Gate
[OK] text found in docs/workflow/0072_transition_progress_ledger_protocol.md :: Progress Update Gate
[OK] text found in docs/templates/0122_web_ai_session_prompts.md :: locally implemented with proof; connector validation pending
[OK] text found in docs/templates/0122_web_ai_session_prompts.md :: Local proof and remote connector proof are not conflated
[OK] text found in docs/templates/0122_web_ai_session_prompts.md :: Git mutation instructions are absent unless the owner explicitly requested Git operations
[OK] text found in docs/templates/0122_web_ai_session_prompts.md :: NEXT does not skip required progress ledger or handoff updates
[PASS] AI rules audit passed
```

```text
make verify
```

Result:

```text
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed
```

## GAP

- The audit checks stable text anchors, not full semantic compliance of every future model answer.
- Git mutation safety still depends on the prompt and connector permission model outside the local repo.

## PROGRESS

Workflow guardrail hardening scope: 100%.

Laravel-to-Go implementation progress: unchanged.

## NEXT

Next execution channel: owner/local terminal.

Review proof output and continue only after ledger/handoff progress rules are satisfied.

## CONTEXT WINDOW STATUS

Enough context remains for proof and final reporting.
