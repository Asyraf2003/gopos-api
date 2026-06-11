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

# Handoff: Capability Route Audit Script

## Date

2026-06-08

## Active Scope

Implement route-to-capability audit script only for `docs/blueprints/0010_capability_control_foundation.md` step 7.

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

- `docs/README.md`
- `docs/AGENTS.md`
- `docs/0001_index.md`
- `docs/0002_decision_policy.md`
- `docs/0003_session_start_protocol.md`
- `docs/templates/0120_prompt_authoring_rules.md`
- `docs/templates/0122_web_ai_session_prompts.md`
- `docs/workflow/0071_handoff_protocol.md`
- `docs/workflow/0072_transition_progress_ledger_protocol.md`
- `docs/blueprints/0010_capability_control_foundation.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`
- `docs/handoffs/2026-06-07-capability-admin-http-surface.md`
- `internal/app/bootstrap/app.go`
- `internal/modules/system/transport/http/me_handler.go`
- `internal/modules/auth/transport/http/logout_handler.go`
- `internal/modules/auth/transport/http/account_role_handler.go`
- `internal/modules/capability/transport/http/capability_handler.go`
- `migrations/0007_seed_existing_protected_capabilities.up.sql`
- `migrations/0008_seed_capability_manage_permission.up.sql`
- `scripts/audit_all.sh`
- `Makefile`

## Files Changed

- `scripts/config/route_capabilities.tsv`
- `scripts/audit_route_capabilities.sh`
- `scripts/audit_all.sh`
- `Makefile`
- `docs/blueprints/0010_capability_control_foundation.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`
- `docs/handoffs/2026-06-08-capability-route-audit-script.md`

## Files Forbidden To Touch

- POS domain CRUD
- `servicecatalog` implementation
- `productcatalog` implementation
- Admin capability HTTP behavior
- Auth behavior
- Future POS business capability seeds
- Production secrets
- GitHub refs, branches, commits, pull requests, issues, labels, reviewers, merges, or CI by Web AI

## Blueprint Referenced

- `docs/blueprints/0010_capability_control_foundation.md`

## Decisions Made

- Use a static manifest for current protected route-to-capability coverage.
- Keep aggregate admin capability coverage as `capability.manage` with wildcard method and prefix match.
- Wire route capability audit into aggregate `make verify`.
- Add aggregate audit summary so `make verify` shows every passed audit gate at the end.
- Do not start POS CRUD.
- Do not change admin capability HTTP behavior.
- Do not change auth behavior.
- Do not add future POS capabilities before domain contracts exist.

## Implementation Facts

- `scripts/config/route_capabilities.tsv` records 6 current protected route capability rows.
- `scripts/audit_route_capabilities.sh` validates manifest rows against migration seed files and route source patterns.
- `Makefile` exposes `make audit-route-capabilities`.
- `scripts/audit_all.sh` runs route capability audit and prints a final aggregate summary.
- `make verify` includes route capability audit through audit-all.

## Proof Collected

Standalone script proof:

```text
bash scripts/audit_route_capabilities.sh
== route capability audit ==
manifest: scripts/config/route_capabilities.tsv

checked route capability rows: 6
[PASS] route capability audit passed
```

Make target proof:

```text
make audit-route-capabilities
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

## Tests Or Commands Run

- `bash scripts/audit_route_capabilities.sh`
- `make audit-route-capabilities`
- `make verify`

## Gaps Still Open

- Route-level disabled protected endpoint proof remains open.
- POS CRUD remains blocked until capability-control foundation proof is complete.
- Future POS business capabilities must wait for accepted domain contracts.

## Next Valid Active Step

Add route-level disabled protected endpoint proof for current protected routes.

Do not start POS CRUD.

## Estimated Scope Progress Percentage

Capability-control foundation: 85%.

Overall Laravel-to-Go transition: 20%.

## Estimated Context-Window Status

Enough context remains for one focused follow-up step: route-level disabled protected endpoint proof.
