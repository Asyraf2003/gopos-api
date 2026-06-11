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

# Capability Control Foundation Blueprint

## FACT
- `docs/` is the canonical documentation root; standards live beside active product blueprints, ADRs, evidence, handoffs, and archive.
- Active implementation blueprints live in `docs/blueprints/`.
- Current repo stack is Go, Echo, PostgreSQL, and hexagonal ports/adapters.
- Current runtime already has auth, JWT verification, DB-backed principal resolution, roles, permissions, refresh/logout, `/api/me`, `/api/authz/me`, and admin account-role routes.
- Current authorization uses stable permission keys resolved from PostgreSQL.
- Accepted ADRs already define roles, permissions, request principal, admin authz minimum, and centralized API output direction.
- `docs/architecture/0022_api_capability_control.md` requires protected API operations to have capability metadata and a runtime capability check before request validation and use case execution.

## GAP
- No capability-control foundation blocker remains after 2026-06-08 closeout proof.
- There is no active business-domain blueprint for products, sale orders, payments, or inventory.
- There is no domain contract yet for POS business CRUD or transaction lifecycle.
- Future POS business capability seeds must wait for accepted domain contracts.

## DECISION
- Capability-control foundation is closed with proof before POS business APIs.
- Existing permissions remain the authorization source of truth.
- Capability keys are separate API exposure controls and must reference required permission keys.
- Capability checks must not replace authn or authz.
- No product, order, payment, inventory, or transaction CRUD implementation is active in this blueprint.
- The first POS business-domain blueprint/domain contract may start after this closeout.
- Capability-control implementation must follow the current repo layout and hexagonal boundaries.

## SCOPE-IN
- Define capability metadata shape.
- Add a runtime capability check path for protected endpoints.
- Add storage or registry design for capability state.
- Add admin-only capability list/show/enable/disable plan.
- Add script/test proof requirements for route-to-capability coverage.
- Cover existing protected routes first as bootstrap data.

## SCOPE-OUT
- Products CRUD.
- Sale order workflow.
- Payment/refund flow.
- Inventory stock movement.
- Report APIs.
- UI implementation.
- Replacing roles/permissions authorization.
- Full policy engine.
- Audit sink beyond the explicit capability-control audit decision.

## TARGET SHAPE
Every protected endpoint follows:

```text
request id -> recovery -> authn -> authz -> capability check -> validation -> use case -> presenter
```

Every capability declares:

```text
key
domain
operation
method
path
default_state
required_permission
risk_level
audit_required
idempotency_required
owner_package
test_proof
disabled_reason
```

## CURRENT STATE
- Authn middleware exists in `internal/transport/http/middleware`.
- Permission guard exists and checks `domain.Principal.HasPermission`.
- PostgreSQL migrations seed roles and permissions.
- Admin account-role routes are protected by auth and `account.role.assign`.
- Current protected route wiring is centralized in `internal/app/bootstrap/app.go`.
- Health endpoint is public infrastructure and may stay outside capability control.
- Capability-control foundation closeout proof passed on 2026-06-08.
- Current protected routes have route-to-capability audit coverage and route-level disabled proof.

## PACKAGE/ARCHITECTURE PLAN
- Add a dedicated capability module under `internal/modules/capability/`.
- Suggested module roles:
  - `domain`: capability metadata, state, and validation invariants.
  - `ports`: capability repository/query interfaces.
  - `usecase`: list, show, enable, disable, and check capability orchestration.
  - `transport/http`: admin HTTP handlers for list/show/enable/disable.
- Add platform PostgreSQL adapter under `internal/platform/postgres` for capability persistence.
- Add middleware or policy adapter under `internal/transport/http/middleware` for runtime capability checks.
- Keep Echo imports out of domain/usecase.
- Keep SQL out of handlers and use cases.
- Keep public response DTOs in `internal/presentation/http/id/...` when user-facing output is added.

## DB/MIGRATION PLAN
- Add a small PostgreSQL migration for capability state.
- Candidate table: `api_capabilities`.
- Minimum columns:
  - `key text primary key`
  - `domain text not null`
  - `operation text not null`
  - `method text not null`
  - `path text not null`
  - `default_enabled boolean not null`
  - `enabled boolean not null`
  - `required_permission text not null`
  - `risk_level text not null`
  - `audit_required boolean not null`
  - `idempotency_required boolean not null`
  - `owner_package text not null`
  - `disabled_reason text null`
  - `created_at timestamptz not null default now()`
  - `updated_at timestamptz not null default now()`
- Seed existing protected routes first.
- Do not seed future POS business capabilities before their domain contracts exist.
- Rollback must drop only capability-control tables created by this slice.

## API CONTRACT PLAN
- Admin capability endpoints are protected API contracts.
- Initial proposed endpoints:
  - `GET /api/admin/capabilities`
  - `GET /api/admin/capabilities/:key`
  - `POST /api/admin/capabilities/:key/enable`
  - `POST /api/admin/capabilities/:key/disable`
- Responses use the project envelope family:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

- Disabled protected endpoint returns `403` before request validation and before use case execution.
- Error responses must use the centralized public error envelope direction.

## ADMIN CONTROL SURFACE PLAN
- Admin users can list capability metadata.
- Admin users can inspect capability state and disabled reason.
- Admin users can enable a capability.
- Admin users can disable a capability with an optional reason.
- Capability metadata must be readable by future UI so the UI can hide, disable, or explain unavailable operations without becoming the authority source.

## SECURITY/AUTHZ PLAN
- Authn remains mandatory for protected admin capability routes.
- Authz remains permission-based.
- Add a new permission key for capability administration, for example `capability.manage`.
- Runtime capability check happens after authz and before validation/use case.
- Disabled capability must stop before the use case runs.
- Capability control must not become the only authorization layer.
- Capability enable/disable should be auditable by decision, even if the first implementation records only state and reason.

## TEST/SCRIPT PROOF PLAN
- Unit tests for capability domain validation and use cases.
- HTTP tests for admin list/show/enable/disable route shape, authz, validation, and envelope.
- Middleware tests proving disabled capability returns `403` before handler/usecase execution.
- Repository/integration tests for PostgreSQL capability state when `DATABASE_URL` is available.
- `go test ./...` remains mandatory for Go changes.
- Route-to-capability audit script proof is wired into `make verify`, aligned with `docs/scripts/0090_makefile_and_scripts.md`.
- Full local proof target should remain `make audit-all` or future `make verify` when available.

## STEP ORDER
1. Add this active blueprint.
2. Add capability domain and usecase contracts without HTTP wiring. Done with proof in `docs/archive/handoffs-closed/2026-06-07-capability-contracts.md`.
3. Add PostgreSQL migration and adapter for capability state. Done with proof in `docs/archive/handoffs-closed/2026-06-07-capability-postgres-state.md`.
4. Add runtime capability check middleware/policy. Done with proof in `docs/archive/handoffs-closed/2026-06-07-capability-runtime-middleware.md`.
5. Seed existing protected routes as capability records. Done with proof in `docs/archive/handoffs-closed/2026-06-07-capability-route-seeds.md`.
6. Add admin capability HTTP surface. Done with proof in `docs/archive/handoffs-closed/2026-06-07-capability-admin-http-surface.md`.
7. Add route-to-capability audit script. Done with proof in `docs/archive/handoffs-closed/2026-06-08-capability-route-audit-script.md`.
8. Add route-level disabled protected endpoint proof. Done with proof in `docs/archive/handoffs-closed/2026-06-08-capability-route-disabled-proof.md`.
9. Close capability-control foundation proof. Done with proof in `docs/archive/handoffs-closed/2026-06-08-capability-control-closeout.md`.
10. Only after capability-control proof closeout, create the first POS business-domain blueprint/domain contract.

## DOD
- Blueprint exists in `docs/blueprints/`.
- Blueprint does not activate POS business CRUD.
- Blueprint references current auth, roles, permissions, and admin route state.
- Blueprint defines architecture, DB, API, admin, security, test, and script plans.
- Blueprint leaves exactly one next active implementation step.
- Proof includes file inspection and diff inspection.

## NEXT ACTIVE STEP

Capability-control foundation is closed.

Start the first POS business-domain blueprint/domain contract.

Do not start POS CRUD implementation.

Do not add future POS business capability seeds before accepted domain contracts.

## CLOSEOUT

Capability-control foundation closed on 2026-06-08.

Owner/local terminal proof:

```text
go test ./internal/transport/http/middleware/... ./internal/modules/auth/transport/http/... ./internal/app/bootstrap/...
ok  pos-go/internal/transport/http/middleware  (cached)
ok  pos-go/internal/modules/auth/transport/http  (cached)
ok  pos-go/internal/app/bootstrap  (cached)

bash scripts/audit_route_capabilities.sh
checked route capability rows: 6
[PASS] route capability audit passed

make db-status
[APPLIED] 0006_capability_control.up.sql
[APPLIED] 0007_seed_existing_protected_capabilities.up.sql
[APPLIED] 0008_seed_capability_manage_permission.up.sql

make verify
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed

Gosec:
Files: 97
Lines: 3978
Nosec: 0
Issues: 0

Closeout decision:

Capability-control foundation: 100%
Stage 1 Go quality foundation: 90%
Overall Laravel-to-Go transition: 20%

Next valid active step:

Start the first POS business-domain blueprint/domain contract.
Do not start POS CRUD implementation.
Do not add future POS business capability seeds before accepted domain contracts.
```
