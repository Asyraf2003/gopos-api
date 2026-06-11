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

# Handoff: ServiceCatalog Runtime Capability Implementation

## Date

2026-06-08

## Active Scope

Implement accepted blueprint 0027: ServiceCatalog HTTP runtime, route registration, presenters, permission/capability seed migration, route capability manifest coverage, audit coverage, and DB migration proof.

## Files Changed

```text
internal/app/bootstrap/app.go
internal/modules/servicecatalog/transport/http/
internal/presentation/http/id/servicecatalog/
migrations/0010_seed_service_catalog_permissions_capabilities.up.sql
migrations/0010_seed_service_catalog_permissions_capabilities.down.sql
scripts/audit_route_capabilities.sh
scripts/config/route_capabilities.tsv
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/handoffs/2026-06-08-servicecatalog-runtime-capability-implementation.md
```

## FACT

- Blueprint 0027 is accepted.
- ServiceCatalog HTTP transport package exists locally.
- ServiceCatalog Indonesian/public response presenters exist locally.
- Bootstrap wires ServiceCatalog routes behind authn, permission, and capability middleware.
- Route capability manifest includes seven ServiceCatalog protected routes.
- Route capability audit checks thirteen protected route rows and passes.
- Migration `0010_seed_service_catalog_permissions_capabilities.up.sql` seeds:
  - `service_catalog.read`
  - `service_catalog.manage`
  - cashier read permission
  - admin read/manage permissions
  - seven ServiceCatalog API capabilities
- Local database proof shows migration 0010 applied.
- Local full `make verify` proof passes.

## PROOF

Local DB proof:

```text
[APPLIED] 0010_seed_service_catalog_permissions_capabilities.up.sql
```

Local aggregate proof:

```text
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

## GAP

- Connector validation passed for implementation visibility through GitHub; focused handler and ServiceCatalog-specific disabled-capability proof now pass locally.
- Handler package now has focused local unit tests for query parsing, body validation, command mapping, envelope shape, and not-found mapping.
- ProductCatalog remains unstarted and blocked by its own accepted domain contract and owner decisions.

## PROGRESS

ServiceCatalog runtime/capability implementation: remote-visible through GitHub connector with local proof; focused handler and disabled-capability proof remote-visible through GitHub connector with local proof; connector validation passed for latest closeout proof files.

Business Phase 1: 35%.

Overall Laravel-to-Go transition: 30%.

## CONTEXT WINDOW STATUS

Enough context remains for next planning-only scope selection. ServiceCatalog runtime/capability implementation evidence is remote-visible through GitHub connector, and owner/local focused plus aggregate proof is recorded.

## NEXT

Execution channel: owner/local terminal.

Prepare the next scoped session prompt after ServiceCatalog runtime/capability closeout proof is recorded.
