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

# Handoff: Decision And Data Request Protocol

## Date

2026-06-07

## Active Scope

Document the owner workflow rule for missing Laravel data and ADR-level decisions.

## Current Branch Or Source Snapshot

Local workspace: `/home/asyraf/Code/go/pos-go`

Previous transition ledger step was already clean and audit-passing.

## Files Included

```text
docs/0002_decision_policy.md
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/templates/0123_analysis_and_review_prompts.md
docs/templates/0125_data_capture_and_evidence_prompts.md
scripts/audit_ai_rules.sh
```

## Files Changed

```text
docs/0002_decision_policy.md
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/templates/0123_analysis_and_review_prompts.md
docs/templates/0125_data_capture_and_evidence_prompts.md
docs/handoffs/2026-06-07-decision-data-request-protocol.md
scripts/audit_ai_rules.sh
```

## Files Forbidden To Touch

No Go source, migrations, runtime config, or production secrets were in scope.

## Blueprint Referenced

```text
docs/blueprints/0012_laravel_to_go_api_transition_master_plan.md
docs/blueprints/0010_capability_control_foundation.md
```

## ADR And Rules Referenced

```text
docs/0002_decision_policy.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
```

## Decisions Made

- Missing Laravel data must be requested as the smallest specific source batch.
- ADR-level or owner decisions must be asked as concise questions with 2-3 viable options and tradeoffs.
- A recommended option should be first when there is a clear recommendation.
- Implementation must not continue across an unresolved ADR-level decision.

## Proof Collected

AI rules audit passed:

```text
bash scripts/audit_ai_rules.sh
```

Visible result:

```text
[PASS] AI rules audit passed
```

## Tests Or Commands Run

```text
bash scripts/audit_ai_rules.sh
```

Result:

```text
[PASS] AI rules audit passed
```

## Gaps Still Open

- Capability-control implementation proof is still not done.
- Laravel-to-Go overall transition remains at the progress recorded in `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`.

## Next Valid Active Step

Implement or complete `docs/blueprints/0010_capability_control_foundation.md`.

## Estimated Scope Progress Percentage

Decision/data request protocol scope: 100%.

Overall Laravel-to-Go transition: unchanged at 15%.

## Estimated Context-Window Status

Enough context remains for one focused implementation step after this docs protocol step.
