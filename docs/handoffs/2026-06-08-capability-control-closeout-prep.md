# Handoff: Capability Control Closeout Prep

## Date

2026-06-08

## Active Scope

Capability-control foundation closeout prep for `docs/blueprints/0010_capability_control_foundation.md`.

## Current State

Capability-control foundation is at 95%.

Route-to-capability audit script exists, is wired into `make verify`, and passes.

Route-level disabled protected endpoint proof exists for current protected route capability keys.

Current protected routes now have runtime capability guards:

- `/api/me` -> `profile.self.show`
- `/api/authz/me` -> `authz.profile.self.show`
- `/api/auth/logout` -> `auth.session.logout`
- `POST /api/admin/accounts/:account_id/roles` -> `account.role.assign`
- `DELETE /api/admin/accounts/:account_id/roles/:role_key` -> `account.role.remove`
- `/api/admin/capabilities...` -> aggregate `capability.manage`

## Files Changed In Recent Work

```text
scripts/config/route_capabilities.tsv
scripts/audit_route_capabilities.sh
scripts/audit_all.sh
Makefile
internal/app/bootstrap/app.go
internal/app/bootstrap/app_capability_test.go
internal/modules/auth/transport/http/account_role_handler.go
internal/transport/http/middleware/capability_routes_test.go
docs/blueprints/0010_capability_control_foundation.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-08-capability-route-audit-script.md
docs/handoffs/2026-06-08-capability-route-disabled-proof.md
docs/handoffs/2026-06-08-capability-control-closeout-prep.md
```

## Proof Collected

Focused tests passed:

```text
go test ./internal/transport/http/middleware/... ./internal/modules/auth/transport/http/... ./internal/app/bootstrap/...
ok  	pos-go/internal/transport/http/middleware	0.008s
ok  	pos-go/internal/modules/auth/transport/http	0.006s
ok  	pos-go/internal/app/bootstrap	0.191s
```

Route capability audit passed:

```text
bash scripts/audit_route_capabilities.sh
checked route capability rows: 6
[PASS] route capability audit passed
```

`make verify` passed:

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
Files  : 97
Lines  : 3978
Nosec  : 0
Issues : 0
```

## Open Gaps

- Capability-control foundation needs closeout review before POS business blueprint work.
- POS CRUD remains blocked until explicit capability-control closeout.
- Future POS business capabilities must wait for accepted domain contracts.
- Overall Laravel-to-Go transition remains 20%.

## Next Valid Active Step

Close capability-control foundation proof and decide whether the first POS business-domain blueprint can start.

## Next Recommended Session

Target agent: Web AI

Template source: `docs/templates/0122_web_ai_session_prompts.md`

Execution channel: Web AI read-only analysis through GitHub connector, then owner/local terminal for any patch commands.

## Estimated Progress

Capability-control foundation: 95%.

Stage 1 Go quality foundation: 90%.

Overall Laravel-to-Go transition: 20%.

## Context Window Status

Current session context is large. Start a new Web AI session for closeout review if continuing carefully.
