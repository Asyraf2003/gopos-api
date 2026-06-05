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
- There is no first-class capability registry, storage, middleware, or admin control surface yet.
- There is no active business-domain blueprint for products, sale orders, payments, or inventory.
- There is no domain contract yet for POS business CRUD or transaction lifecycle.
- There is no route-to-capability audit script yet.
- There is no API contract test proving disabled capabilities stop before use cases.

## DECISION
- Capability-control foundation is the next active product blueprint before POS business APIs.
- Existing permissions remain the authorization source of truth.
- Capability keys are separate API exposure controls and must reference required permission keys.
- Capability checks must not replace authn or authz.
- No product, order, payment, inventory, or transaction CRUD implementation is active in this blueprint.
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
- Add or extend script proof later for route-to-capability coverage, aligned with `docs/scripts/0090_makefile_and_scripts.md`.
- Full local proof target should remain `make audit-all` or future `make verify` when available.

## STEP ORDER
1. Add this active blueprint.
2. Add capability domain and usecase contracts without HTTP wiring.
3. Add PostgreSQL migration and adapter for capability state.
4. Add runtime capability check middleware/policy.
5. Seed existing protected routes as capability records.
6. Add admin capability HTTP surface.
7. Add route-to-capability audit script.
8. Only after capability-control proof, create the first POS business-domain blueprint.

## DOD
- Blueprint exists in `docs/blueprints/`.
- Blueprint does not activate POS business CRUD.
- Blueprint references current auth, roles, permissions, and admin route state.
- Blueprint defines architecture, DB, API, admin, security, test, and script plans.
- Blueprint leaves exactly one next active implementation step.
- Proof includes file inspection and diff inspection.

## NEXT ACTIVE STEP
Add capability domain and usecase contracts without HTTP wiring.
