# Handoff: Capability Admin HTTP Surface

## Date

2026-06-07

## Active Scope

Add admin capability HTTP surface for `docs/blueprints/0010_capability_control_foundation.md` step 6.

## Files Changed

```text
migrations/0008_seed_capability_manage_permission.up.sql
migrations/0008_seed_capability_manage_permission.down.sql
internal/presentation/http/id/capability/capability.go
internal/modules/capability/transport/http/capability_handler.go
internal/modules/capability/transport/http/capability_handler_test.go
internal/app/bootstrap/app.go
docs/handoffs/2026-06-07-capability-admin-http-surface.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

## Implementation Facts

- Migration `0008` seeds permission `capability.manage`.
- Migration `0008` assigns `capability.manage` only to role `admin`.
- Migration `0008` seeds `api_capabilities.key = 'capability.manage'` for `/api/admin/capabilities` with method `*`, required permission `capability.manage`, high risk, audit required, and owner package `internal/modules/capability/transport/http`.
- Migration `0008` extends `api_capabilities_method_check` to allow `*` because the admin capability control surface uses one aggregate capability key for list/show/enable/disable routes.
- Capability DTO mapping exists in `internal/presentation/http/id/capability/capability.go`.
- Admin capability handler exists in `internal/modules/capability/transport/http/capability_handler.go`.
- Registered handler routes are `GET /capabilities`, `GET /capabilities/:key`, `POST /capabilities/:key/enable`, and `POST /capabilities/:key/disable` on the provided Echo group.
- Handler tests use fake use cases and do not require PostgreSQL.
- Bootstrap wires the capability PostgreSQL repository, list/show/enable/disable use cases, check use case, and admin capability handler.
- Bootstrap protects `/api/admin/capabilities...` with `RequireAuth`, `RequirePermission("capability.manage")`, and `RequireCapability("capability.manage", checkCapabilityUsecase)`.
- Existing account-role route behavior remains on its separate `/api/admin` group with permission `account.role.assign`.

## Proof Placeholders

No proof command output was recorded in this handoff.

The user still needs to run and paste proof for:

```text
make dev
permission capability.manage SQL proof
admin role permission SQL proof
api_capabilities capability.manage SQL proof
go test ./internal/modules/capability/...
go test ./internal/modules/capability/transport/http/...
go test ./internal/app/bootstrap/...
make verify
```

## Remaining Gaps

- Route-to-capability audit script remains out of scope and not implemented.
- Route-level disabled protected endpoint proof remains open unless covered by later proof.
- POS CRUD remains blocked until capability-control foundation proof is complete.

## Next Valid Step

Collect proof for this admin capability HTTP surface step before moving to route-to-capability audit script work.
