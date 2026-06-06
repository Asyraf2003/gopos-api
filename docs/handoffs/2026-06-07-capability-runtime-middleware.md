# Handoff: Capability Runtime Middleware

## Date

2026-06-07

## Active Scope

Add runtime capability check middleware/policy.

## Current Branch Or Source Snapshot

Local workspace: `/home/asyraf/Code/go/pos-go`

## Files Included

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/blueprints/0010_capability_control_foundation.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-capability-postgres-state.md
internal/modules/capability/**
internal/transport/http/middleware/**
```

## Files Changed

```text
internal/transport/http/middleware/capability.go
internal/transport/http/middleware/capability_test.go
internal/transport/http/middleware/capability_test_helpers_test.go
docs/blueprints/0010_capability_control_foundation.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/README.md
docs/workflow/README.md
docs/handoffs/2026-06-07-capability-runtime-middleware.md
```

## Files Forbidden To Touch

No POS domain code, admin capability HTTP surface, protected route wiring, route-to-capability audit script, production secrets, or Git operations were in scope.

## Blueprint Referenced

```text
docs/blueprints/0010_capability_control_foundation.md
```

## ADR And Rules Referenced

```text
docs/architecture/0021_package_boundaries.md
docs/architecture/0022_api_capability_control.md
docs/api/0050_echo_http_contract.md
docs/core/0013_proof_and_progress.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
```

## Decisions Made

- Runtime capability check lives in `internal/transport/http/middleware`.
- Middleware depends on a small `CapabilityChecker` interface instead of importing usecase concrete types.
- Disabled capability maps to `403`.
- Repository or checker failures map to `500`.
- Empty capability key or nil checker maps to `500` as guard misconfiguration.
- The next active step is route capability seeding, not POS domain implementation.

## Proof Collected

```text
env GOCACHE=/tmp/go-build go test ./internal/transport/http/middleware -run Capability
```

Result:

```text
ok   pos-go/internal/transport/http/middleware
```

```text
env GOCACHE=/tmp/go-build go test ./internal/modules/capability/...
```

Result:

```text
ok   pos-go/internal/modules/capability/domain
ok   pos-go/internal/modules/capability/usecase
```

```text
bash scripts/audit_hexagonal.sh
```

Result:

```text
[PASS] hexagonal audit passed
```

```text
make verify
```

Result:

```text
[PASS] gosec audit passed
[PASS] aggregate audit passed
```

## Tests Or Commands Run

```text
gofmt -w internal/transport/http/middleware/capability.go internal/transport/http/middleware/capability_test.go internal/transport/http/middleware/capability_test_helpers_test.go
env GOCACHE=/tmp/go-build go test ./internal/transport/http/middleware -run Capability
env GOCACHE=/tmp/go-build go test ./internal/modules/capability/...
bash scripts/audit_hexagonal.sh
make verify
```

## Gaps Still Open

- Existing protected routes are not seeded as capability records yet.
- `capability.manage` permission is not added yet.
- No admin capability HTTP surface yet.
- No route-to-capability audit script yet.
- No route-level API proof shows a disabled protected endpoint stops before validation/usecase yet.
- No POS domain PostgreSQL baseline or business module should start before capability-control proof is complete.

## Next Valid Active Step

Seed existing protected routes as capability records.

Minimum proof for the next step:

```text
seed records cover existing protected routes only
seed path is idempotent
no future POS capability is seeded before domain contracts exist
repository/integration or migration proof verifies seeded records
make verify remains passing
handoff and progress ledger are updated
```

## Web AI Continuation Prompt

Use this when continuing in GPT web or another browser AI with a GitHub connector:

```text
You are Web AI helping on the Go POS API repository.

Use the GitHub connector as read-only source of truth for committed repository files. Do not mutate GitHub, create commits, open PRs, edit issues, or change refs. If you need local-only proof, ask me for exact terminal output.

Read first:
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/blueprints/0010_capability_control_foundation.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-capability-runtime-middleware.md

Active step:
Seed existing protected routes as capability records.

Stay inside this step. Do not start POS domain CRUD.

Expected output:
FACT, GAP, proposed file changes, exact terminal commands for local Codex/user execution, proof commands, and a draft handoff update. If a decision is needed, ask one concise question with 2-3 options and plus/minus tradeoffs.
```

## Estimated Scope Progress Percentage

Capability runtime middleware active step: 100%.

Capability-control foundation: 65%.

Overall Laravel-to-Go transition: 19%.

## Estimated Context-Window Status

This session has enough context for one focused follow-up step, but Web AI continuation should start from the ledger and this handoff.
