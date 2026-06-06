# Handoff: Capability PostgreSQL State

## Date

2026-06-07

## Active Scope

Add PostgreSQL migration and adapter for capability state.

## Current Branch Or Source Snapshot

Local workspace: `/home/asyraf/Code/go/pos-go`

## Files Included

```text
docs/blueprints/0010_capability_control_foundation.md
docs/db/0040_postgresql_policy.md
internal/modules/capability/domain/**
internal/modules/capability/ports/**
internal/modules/capability/usecase/**
internal/platform/postgres/tx.go
internal/platform/postgres/account_role_assigner.go
internal/platform/postgres/principal_resolver.go
migrations/0002_authorization_minimum.up.sql
scripts/db_migrate.sh
```

## Files Changed

```text
migrations/0006_capability_control.up.sql
migrations/0006_capability_control.down.sql
internal/platform/postgres/capability_repository.go
internal/platform/postgres/capability_repository_query.go
internal/platform/postgres/capability_repository_row.go
internal/platform/postgres/capability_repository_integration_test.go
docs/blueprints/0010_capability_control_foundation.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-capability-postgres-state.md
```

## Files Forbidden To Touch

No HTTP route wiring, middleware, admin HTTP surface, POS domain code, runtime config, or production secrets were in scope.

## Blueprint Referenced

```text
docs/blueprints/0010_capability_control_foundation.md
```

## ADR And Rules Referenced

```text
docs/db/0040_postgresql_policy.md
docs/architecture/0021_package_boundaries.md
docs/architecture/0022_api_capability_control.md
docs/core/0013_proof_and_progress.md
```

## Decisions Made

- Capability state table is `api_capabilities`.
- Table owns risk-level and HTTP method check constraints.
- Capability repository lives under `internal/platform/postgres`.
- Repository implements list, get, and save/upsert against `ports.CapabilityRepository`.
- Repository honors existing transaction context via `TxFromContext`.
- Existing protected route capability seeding and `capability.manage` permission are deferred to HTTP/route wiring.

## Proof Collected

```text
make db-migrate
```

Result:

```text
[APPLY] 0006_capability_control.up.sql
[PASS] db migrate completed
```

```text
make db-status
```

Result:

```text
[APPLIED] 0006_capability_control.up.sql
```

```text
env GOCACHE=/tmp/go-build go test -tags=integration ./internal/platform/postgres
```

Result:

```text
ok   pos-go/internal/platform/postgres
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
gofmt -w internal/platform/postgres/capability_repository*.go
env GOCACHE=/tmp/go-build go test ./internal/platform/postgres -run '^$'
env GOCACHE=/tmp/go-build go test ./internal/modules/capability/...
env GOCACHE=/tmp/go-build go test ./internal/modules/...
make db-migrate
env GOCACHE=/tmp/go-build go test -tags=integration ./internal/platform/postgres -run CapabilityRepository
make verify
make db-status
env GOCACHE=/tmp/go-build go test -tags=integration ./internal/platform/postgres
```

`make db-migrate` and `make db-status` needed unsandboxed local PostgreSQL access. Both passed after approval.

## Gaps Still Open

- No runtime capability middleware yet.
- No admin capability HTTP surface yet.
- No route-to-capability audit script yet.
- No disabled endpoint API proof returning `403` before validation/usecase yet.
- Existing protected routes are not seeded as capability records yet.
- `capability.manage` permission is not added yet.

## Next Valid Active Step

Add runtime capability check middleware/policy.

## Estimated Scope Progress Percentage

Capability PostgreSQL state active step: 100%.

Capability-control foundation: 50%.

Overall Laravel-to-Go transition: 18%.

## Estimated Context-Window Status

Enough context remains for one focused follow-up step if needed.
