# Handoff: Capability Route Seeds

## Date

2026-06-07

## Active Scope

Seed existing protected routes as capability records.

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
docs/handoffs/2026-06-07-capability-runtime-middleware.md
internal/app/bootstrap/app.go
internal/modules/system/transport/http/me_handler.go
internal/modules/auth/transport/http/logout_handler.go
internal/modules/auth/transport/http/account_role_handler.go
migrations/0006_capability_control.up.sql
```

## Files Changed

```text
migrations/0007_seed_existing_protected_capabilities.up.sql
migrations/0007_seed_existing_protected_capabilities.down.sql
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-07-capability-route-seeds.md
```

## Files Forbidden To Touch

```text
POS domain CRUD
servicecatalog implementation
productcatalog implementation
admin capability HTTP surface
route-to-capability audit script
production secrets
GitHub refs, branches, commits, PRs, issues
```

## Blueprint Referenced

```text
docs/blueprints/0010_capability_control_foundation.md
```

## Decisions Made

Seed only existing protected routes wired behind `RequireAuth` and `RequirePermission`.

Do not seed future POS business capabilities.

Do not seed `POST /api/auth/refresh` because current bootstrap does not protect it with `RequireAuth`.

Use idempotent `INSERT ... ON CONFLICT (key) DO UPDATE`.

Keep `capability.manage` for the future admin capability HTTP surface step.

## Seeded Capability Records

```text
profile.self.show                 GET     /api/me
authz.profile.self.show           GET     /api/authz/me
auth.session.logout               POST    /api/auth/logout
account.role.assign               POST    /api/admin/accounts/:account_id/roles
account.role.remove               DELETE  /api/admin/accounts/:account_id/roles/:role_key
```

## Proof Collected

Paste exact terminal output here:

```bash
git status --short
grep -R "sale\.order\|product\|inventory\|payment" -n migrations/0007_seed_existing_protected_capabilities.*.sql || true
grep -n "profile.self.show\|authz.profile.self.show\|auth.session.logout\|account.role.assign\|account.role.remove" migrations/0007_seed_existing_protected_capabilities.up.sql
env GOCACHE=/tmp/go-build go test ./internal/modules/capability/...
env GOCACHE=/tmp/go-build go test ./internal/platform/postgres/...
env GOCACHE=/tmp/go-build go test ./internal/transport/http/middleware -run Capability
make verify
```

## Gaps Still Open

`capability.manage` permission is not added yet.

Admin capability HTTP surface is not implemented yet.

Route-to-capability audit script is not implemented yet.

No route-level API proof yet shows a disabled protected endpoint stops before validation/usecase.

No POS domain PostgreSQL baseline or business module should start before capability-control proof is complete.

## Next Valid Active Step

Add admin capability HTTP surface.

Minimum proof for next step:

```text
capability.manage permission exists
admin list/show/enable/disable handlers exist
admin routes are protected by authn, authz, and capability check
disabled capability returns 403 before validation/usecase
go test ./...
make verify
handoff and ledger updated
```

## Estimated Scope Progress Percentage

Capability route seed active step: 100% only after proof commands pass.

Capability-control foundation: 72% after route seed proof.

Overall Laravel-to-Go transition: 19%, unchanged unless ledger owner decides route seed proof moves the aggregate.

## Estimated Context-Window Status

Enough context remains for one focused follow-up step after route seed proof.

## Ledger Update Patch Note

Update `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md` only after proof exists:

```md
- Existing protected route capability records are seeded idempotently in `migrations/0007_seed_existing_protected_capabilities.up.sql`.
- Seeded records cover only current protected routes:
  - `GET /api/me`
  - `GET /api/authz/me`
  - `POST /api/auth/logout`
  - `POST /api/admin/accounts/:account_id/roles`
  - `DELETE /api/admin/accounts/:account_id/roles/:role_key`
- No future POS business capabilities were seeded before domain contracts exist.
```

Update open gaps by removing:

```text
existing protected route capability records
```

Keep these gaps:

```text
capability.manage permission
admin HTTP surface
route-to-capability audit
route-level disabled protected endpoint proof
```
