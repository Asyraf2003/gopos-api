# Handoff: Capability Contracts

## Date

2026-06-07

## Active Scope

Implement capability domain and usecase contracts without HTTP wiring.

## Current Branch Or Source Snapshot

Local workspace: `/home/asyraf/Code/go/pos-go`

## Files Included

```text
docs/blueprints/0010_capability_control_foundation.md
docs/architecture/0022_api_capability_control.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
internal/modules/auth/domain/principal.go
internal/modules/auth/usecase/assign_account_role.go
internal/modules/auth/usecase/remove_account_role.go
internal/modules/auth/ports/account_role_assigner.go
scripts/audit_hexagonal.sh
```

## Files Changed

```text
internal/modules/capability/domain/capability.go
internal/modules/capability/domain/errors.go
internal/modules/capability/domain/capability_test.go
internal/modules/capability/ports/capability_repository.go
internal/modules/capability/usecase/check_capability.go
internal/modules/capability/usecase/disable_capability.go
internal/modules/capability/usecase/enable_capability.go
internal/modules/capability/usecase/list_capabilities.go
internal/modules/capability/usecase/show_capability.go
internal/modules/capability/usecase/capability_test.go
docs/blueprints/0010_capability_control_foundation.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-capability-contracts.md
```

## Files Forbidden To Touch

No HTTP routes, PostgreSQL adapter, migration, runtime config, or production secrets were in scope for this step.

## Blueprint Referenced

```text
docs/blueprints/0010_capability_control_foundation.md
```

## ADR And Rules Referenced

```text
docs/architecture/0021_package_boundaries.md
docs/architecture/0022_api_capability_control.md
docs/core/0012_step_by_step_execution.md
docs/core/0013_proof_and_progress.md
```

## Decisions Made

- Capability contracts are a dedicated module under `internal/modules/capability/`.
- Domain owns capability metadata, risk-level validation, enable/disable state behavior, and disabled error.
- Ports expose a repository contract with list, get, and save operations.
- Use cases added: list, show, check, enable, and disable capability.
- No HTTP, SQL, migration, route registration, or middleware wiring was added in this step.

## Proof Collected

```text
env GOCACHE=/tmp/go-build go test ./internal/modules/capability/...
```

Result:

```text
ok   pos-go/internal/modules/capability/domain
?    pos-go/internal/modules/capability/ports [no test files]
ok   pos-go/internal/modules/capability/usecase
```

```text
env GOCACHE=/tmp/go-build go test ./internal/modules/...
```

Result:

```text
ok   pos-go/internal/modules/auth/domain
?    pos-go/internal/modules/auth/ports [no test files]
ok   pos-go/internal/modules/auth/transport/http
ok   pos-go/internal/modules/auth/usecase
ok   pos-go/internal/modules/capability/domain
?    pos-go/internal/modules/capability/ports [no test files]
ok   pos-go/internal/modules/capability/usecase
?    pos-go/internal/modules/system/ports [no test files]
?    pos-go/internal/modules/system/transport/http [no test files]
```

```text
bash scripts/audit_hexagonal.sh
```

Result:

```text
[PASS] hexagonal import audit passed
```

```text
bash scripts/audit_format.sh
```

Result:

```text
[PASS] format audit passed
```

```text
bash scripts/audit_go_vet.sh
```

Result:

```text
[PASS] go vet audit passed
```

```text
bash scripts/audit_ai_rules.sh
```

Result:

```text
[PASS] AI rules audit passed
```

## Tests Or Commands Run

```text
gofmt -w internal/modules/capability/...
env GOCACHE=/tmp/go-build go test ./internal/modules/capability/...
env GOCACHE=/tmp/go-build go test ./internal/modules/auth/domain ./internal/modules/auth/usecase ./internal/modules/system/...
env GOCACHE=/tmp/go-build go test ./internal/modules/...
bash scripts/audit_hexagonal.sh
bash scripts/audit_format.sh
bash scripts/audit_go_vet.sh
bash scripts/audit_ai_rules.sh
```

Initial `go test ./internal/modules/capability/...` without `GOCACHE` failed because `/home/asyraf/.cache/go-build` was read-only in the sandbox. The tests passed after using `GOCACHE=/tmp/go-build`.

## Gaps Still Open

- No PostgreSQL migration or adapter for capability state yet.
- No runtime capability middleware yet.
- No admin capability HTTP surface yet.
- No route-to-capability audit script yet.
- No disabled endpoint API proof returning `403` before validation/usecase yet.
- Security gosec proof remains outside this step.

## Next Valid Active Step

Add PostgreSQL migration and adapter for capability state.

## Estimated Scope Progress Percentage

Capability contracts active step: 100%.

Capability-control foundation: 35%.

Overall Laravel-to-Go transition: 16%.

## Estimated Context-Window Status

Enough context remains for one focused follow-up step if needed.
