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

# Handoff: Capability Route Disabled Proof

## Date

2026-06-08

## Active Scope

Add route-level disabled protected endpoint proof for current protected routes in `docs/blueprints/0010_capability_control_foundation.md`.

## Current Branch Or Source Snapshot

Local workspace:

```text
/home/asyraf/Code/go/pos-go
```

GitHub repository:

```text
Asyraf2003/gopos-api
```

## Files Included

- `docs/blueprints/0010_capability_control_foundation.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`
- `docs/handoffs/2026-06-08-capability-route-audit-script.md`
- `internal/app/bootstrap/app.go`
- `internal/app/bootstrap/app_test.go`
- `internal/modules/auth/transport/http/account_role_handler.go`
- `internal/transport/http/middleware/capability.go`
- `internal/transport/http/middleware/capability_test.go`
- `internal/transport/http/middleware/capability_test_helpers_test.go`
- `scripts/audit_route_capabilities.sh`
- `scripts/audit_all.sh`
- `Makefile`

## Files Changed

- `internal/app/bootstrap/app.go`
- `internal/app/bootstrap/app_capability_test.go`
- `internal/modules/auth/transport/http/account_role_handler.go`
- `internal/transport/http/middleware/capability_routes_test.go`
- `docs/blueprints/0010_capability_control_foundation.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`
- `docs/handoffs/2026-06-08-capability-route-disabled-proof.md`

## Files Forbidden To Touch

- POS domain CRUD
- `servicecatalog` implementation
- `productcatalog` implementation
- Future POS business capability seeds
- Admin capability HTTP behavior
- Auth behavior redesign
- Production secrets
- GitHub refs, branches, commits, pull requests, issues, labels, reviewers, merges, or CI by Web AI

## Blueprint Referenced

- `docs/blueprints/0010_capability_control_foundation.md`

## Decisions Made

- Keep existing permission authorization as-is.
- Add runtime capability guards after permission guards for current protected routes.
- Split account-role handler route registration so assign/remove routes can use separate capability keys.
- Keep admin capability HTTP surface behavior unchanged.
- Add route-level disabled proof in middleware tests for current protected route capability keys.
- Add bootstrap source guard test to prevent capability guards from being silently removed.
- Do not start POS CRUD.
- Do not add future POS business capabilities before domain contracts exist.

## Implementation Facts

- `/api/me` is guarded by `profile.self.show`.
- `/api/authz/me` is guarded by `authz.profile.self.show`.
- `/api/auth/logout` is guarded by `auth.session.logout`.
- `POST /api/admin/accounts/:account_id/roles` is guarded by `account.role.assign`.
- `DELETE /api/admin/accounts/:account_id/roles/:role_key` is guarded by `account.role.remove`.
- `/api/admin/capabilities...` remains guarded by aggregate `capability.manage`.
- Route-level disabled proof covers current protected route capability keys and confirms disabled capability returns 403 before handler execution.
- Route capability audit still checks 6 seeded route capability rows.
- `make verify` passes.

## Proof Collected

Focused test proof:

```text
go test ./internal/transport/http/middleware/... ./internal/modules/auth/transport/http/... ./internal/app/bootstrap/...
ok  	pos-go/internal/transport/http/middleware	0.008s
ok  	pos-go/internal/modules/auth/transport/http	0.006s
ok  	pos-go/internal/app/bootstrap	0.191s
```

Route capability audit proof:

```text
bash scripts/audit_route_capabilities.sh
== route capability audit ==
manifest: scripts/config/route_capabilities.tsv

checked route capability rows: 6
[PASS] route capability audit passed
```

Aggregate proof summary:

```text
== aggregate audit summary ==
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit

[PASS] aggregate audit passed
```

Gosec summary:

```text
Summary:
  Gosec  : dev
  Files  : 97
  Lines  : 3978
  Nosec  : 0
  Issues : 0
```

## Tests Or Commands Run

- `go test ./internal/transport/http/middleware/... ./internal/modules/auth/transport/http/... ./internal/app/bootstrap/...`
- `bash scripts/audit_route_capabilities.sh`
- `make verify`

## Gaps Still Open

- Capability-control foundation needs closeout review before POS business blueprint work.
- POS CRUD remains blocked until explicit capability-control closeout.
- Future POS business capabilities must wait for accepted domain contracts.

## Next Valid Active Step

Close capability-control foundation proof and decide whether the first POS business-domain blueprint can start.

Do not start POS CRUD in this step.

## Estimated Scope Progress Percentage

Capability-control foundation: 95%.

Overall Laravel-to-Go transition: 20%.

Stage 1 Go quality foundation: 90%.

## Estimated Context-Window Status

Enough context remains for one focused closeout step. Prepare a next-session handoff if the next response becomes large.
