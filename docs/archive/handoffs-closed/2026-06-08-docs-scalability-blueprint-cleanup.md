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

# Handoff: Docs Scalability And Blueprint Cleanup

## Date

2026-06-08

## Active Scope

Resolve docs scalability feedback before continuing ServiceCatalog persistence.

## Files Included

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/blueprints/0026_servicecatalog_postgres_persistence_slice.md
docs/blueprints/0026_servicecatalog_runtime_slice_2_plan.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/0004_adr_implementation_proof_index.md
docs/evidence/README.md
docs/evidence/2026-06-06-auth-runtime-local-dev.md
docs/handoffs/README.md
docs/archive/README.md
```

## Files Changed

```text
docs/blueprints/0026_servicecatalog_postgres_persistence_slice.md
docs/archive/blueprints-superseded/README.md
docs/archive/blueprints-superseded/2026-06-08-0026_servicecatalog_runtime_slice_2_plan.md
docs/archive/README.md
docs/evidence/README.md
docs/evidence/2026-06-06-auth-runtime-local-dev.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/0005_laravel_to_go_transition_history_2026_06_08.md
docs/handoffs/README.md
docs/handoffs/2026-06-08-docs-scalability-blueprint-cleanup.md
```

## Files Forbidden To Touch

```text
internal/**
cmd/**
migrations/**
scripts/**
production secrets
GitHub branches, commits, pull requests, issues, labels, reviewers, merges, refs, or CI
```

## Rules Referenced

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/workflow/0070_docs_go_workflow.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
```

## Decisions Made

- Keep `docs/blueprints/0026_servicecatalog_postgres_persistence_slice.md` as the only active `0026` blueprint.
- Archive broader runtime draft `0026_servicecatalog_runtime_slice_2_plan.md` as historical-only context.
- Keep ServiceCatalog persistence separate from HTTP transport, route registration, permission seeds, capability seeds, and route manifest changes.
- Require a later runtime/capability blueprint before any protected ServiceCatalog HTTP route is registered.
- Move completed-work history out of the active ledger into `docs/evidence/0005_laravel_to_go_transition_history_2026_06_08.md`.
- Add handoff archiving policy and current continuation index.
- Add closeout trigger for incomplete auth runtime evidence.
- Keep ADR `0009` and ADR `0012` visible as explicit closeout backlog items.

## Proof Collected

- `rg` proof found docs cleanup anchors for archived duplicate blueprint, active ledger current-state split, history evidence, handoff archiving policy, ADR closeout backlog, incomplete evidence review trigger, and required later runtime/capability blueprint ownership.
- `make verify` passed after docs cleanup.

## Tests Or Commands Run

```text
rg -n "0026_servicecatalog_runtime_slice_2_plan|docs/archive/blueprints-superseded|Current State Summary|0005_laravel_to_go_transition_history|Archiving Policy|ADR closeout backlog|2026-06-15|runtime/capability blueprint|ServiceCatalog PostgreSQL persistence slice is implemented|Create and accept the next ServiceCatalog runtime/capability blueprint|permission seed rows|capability seed rows" docs

make verify
```

## Gaps

- ADR `0009` debug auth lane still needs runtime closeout proof.
- ADR `0012` API output centralization still needs full response/error envelope proof.
- ServiceCatalog capability seed migration is not implemented and must be owned by a later accepted runtime/capability blueprint.

## Next Valid Active Step

Create and accept the next ServiceCatalog runtime/capability blueprint.

The next blueprint must own ServiceCatalog HTTP transport, route registration, request/response presenters, permission seed rows, capability seed rows, route capability manifest updates, and disabled-capability proof before protected ServiceCatalog routes are registered.

## Estimated Scope Progress Percentage

Docs scalability and blueprint cleanup scope: 100%.

Laravel-to-Go transition: unchanged by docs cleanup. Current active ledger estimate is 25%.

## Estimated Context-Window Status

Enough context remains for proof and final reporting.
