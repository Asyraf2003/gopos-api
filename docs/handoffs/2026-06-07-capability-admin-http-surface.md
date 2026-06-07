# Handoff: Capability Admin HTTP Surface

## Date

2026-06-07

## Active Scope

Add admin capability HTTP surface for `docs/blueprints/0010_capability_control_foundation.md` step 6.

## Current Branch Or Source Snapshot

Local workspace: `/home/asyraf/Code/go/pos-go`

Source of truth checked through GitHub connector on repository:

```text
Asyraf2003/gopos-api
```

## Files Included

- `docs/README.md`
- `docs/AGENTS.md`
- `docs/0001_index.md`
- `docs/0002_decision_policy.md`
- `docs/0003_session_start_protocol.md`
- `docs/blueprints/0010_capability_control_foundation.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`
- `docs/workflow/0071_handoff_protocol.md`
- `internal/app/bootstrap/app.go`
- `internal/modules/capability/domain/capability.go`
- `internal/modules/capability/usecase/list_capabilities.go`
- `internal/modules/capability/usecase/show_capability.go`
- `internal/modules/capability/usecase/enable_capability.go`
- `internal/modules/capability/usecase/disable_capability.go`
- `internal/modules/capability/usecase/check_capability.go`
- `internal/modules/capability/ports/capability_repository.go`
- `internal/platform/postgres/capability_repository.go`
- `internal/transport/http/middleware/capability.go`

## Files Changed

- `migrations/0008_seed_capability_manage_permission.up.sql`
- `migrations/0008_seed_capability_manage_permission.down.sql`
- `internal/presentation/http/id/capability/capability.go`
- `internal/modules/capability/transport/http/capability_handler.go`
- `internal/modules/capability/transport/http/capability_handler_read.go`
- `internal/modules/capability/transport/http/capability_handler_write.go`
- `internal/modules/capability/transport/http/capability_handler_response.go`
- `internal/modules/capability/transport/http/capability_handler_test.go`
- `internal/modules/capability/transport/http/capability_handler_read_test.go`
- `internal/modules/capability/transport/http/capability_handler_show_error_test.go`
- `internal/modules/capability/transport/http/capability_handler_write_test.go`
- `internal/modules/capability/transport/http/capability_handler_write_error_test.go`
- `internal/modules/capability/transport/http/capability_handler_test_assert_test.go`
- `internal/modules/capability/transport/http/capability_handler_test_context_test.go`
- `internal/modules/capability/transport/http/capability_handler_test_decode_test.go`
- `internal/modules/capability/transport/http/capability_handler_test_fake_test.go`
- `internal/modules/capability/transport/http/capability_handler_test_fake_show_test.go`
- `internal/modules/capability/transport/http/capability_handler_test_fake_write_test.go`
- `internal/app/bootstrap/app.go`
- `docs/handoffs/2026-06-07-capability-admin-http-surface.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`

## Files Forbidden To Touch

- POS domain CRUD
- `servicecatalog` implementation
- `productcatalog` implementation
- Route-to-capability audit script in this closed step
- Production secrets
- GitHub refs, branches, commits, PRs, issues by Web AI

## Blueprint Referenced

- `docs/blueprints/0010_capability_control_foundation.md`

## ADR/Rules Referenced

- `docs/workflow/0071_handoff_protocol.md`
- `docs/workflow/0072_transition_progress_ledger_protocol.md`
- `docs/architecture/0022_api_capability_control.md`
- `docs/testing/0060_test_and_quality_gates.md`

## Decisions Made

- Use one aggregate capability key: `capability.manage`.
- Use one permission key: `capability.manage`.
- Assign `capability.manage` only to role `admin`.
- Seed `api_capabilities.key = 'capability.manage'` for `/api/admin/capabilities`.
- Allow method `*` in `api_capabilities_method_check` for the aggregate admin capability gate.
- Split capability HTTP handler and tests to satisfy the 100-line file-size audit.
- Keep POS CRUD blocked.

## Implementation Facts

- Migration `0008` seeds permission `capability.manage`.
- Migration `0008` assigns `capability.manage` to role `admin`.
- Migration `0008` seeds `api_capabilities.key = 'capability.manage'`.
- Capability DTO mapping exists under `internal/presentation/http/id/capability/`.
- Admin capability list/show/enable/disable handler exists under `internal/modules/capability/transport/http/`.
- Bootstrap wires `/api/admin/capabilities...` behind authn, `capability.manage` authz, and runtime capability check.
- Existing account-role admin routes remain separate under `account.role.assign`.

## Proof Collected

```text
make dev
[APPLY] 0008_seed_capability_manage_permission.up.sql
[PASS] db migrate completed
```

SQL proof:

```text
permission capability.manage: 1 row
admin role permission capability.manage: 1 row
api_capabilities capability.manage: 1 row
```

```text
go test ./internal/modules/capability/transport/http/...
ok   pos-go/internal/modules/capability/transport/http
```

```text
go test ./internal/modules/capability/...
ok   pos-go/internal/modules/capability/domain
ok   pos-go/internal/modules/capability/transport/http
ok   pos-go/internal/modules/capability/usecase
```

```text
go test ./internal/app/bootstrap/...
ok   pos-go/internal/app/bootstrap
```

```text
make verify
[PASS] go vet audit passed
[PASS] format audit passed
[PASS] AI rules audit passed
[PASS] file size audit passed
[PASS] hexagonal import audit passed
[PASS] gosec audit passed
[PASS] aggregate audit passed
```

Gosec summary:

```text
Files: 97
Lines: 3960
Nosec: 0
Issues: 0
```

## Tests Or Commands Run

- `make dev`
- SQL permission proof
- SQL role-permission proof
- SQL `api_capabilities` proof
- `wc -l internal/modules/capability/transport/http/*.go`
- `go test ./internal/modules/capability/transport/http/...`
- `go test ./internal/modules/capability/...`
- `go test ./internal/app/bootstrap/...`
- `make verify`

## Gaps Still Open

- Route-to-capability audit script is not implemented yet.
- Route-level disabled protected endpoint proof remains open unless covered by later proof.
- POS CRUD remains blocked until capability-control foundation proof is complete.

## Next Valid Active Step

Add route-to-capability audit script for protected routes.

Do not start POS CRUD.

## Estimated Scope Progress Percentage

Admin capability HTTP surface step: 100%.

Capability-control foundation: 75%.

Overall Laravel-to-Go transition: 20%.

## Estimated Context-Window Status

Enough context remains for one focused follow-up step: route-to-capability audit script.
