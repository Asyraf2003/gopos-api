# Laravel To Go Transition Progress Ledger

## Status

Date updated: 2026-06-07

Active blueprint:

```text
docs/blueprints/0012_laravel_to_go_api_transition_master_plan.md
```

Related blueprints:

```text
docs/blueprints/0010_capability_control_foundation.md
docs/blueprints/0020_catalog_foundation_migration.md
docs/blueprints/0022_manual_auth_login_foundation.md
docs/blueprints/0023_quality_security_hex_audit_gates.md
```

Related evidence:

```text
docs/evidence/0001_laravel_stage0_schema_and_route_inventory.md
docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md
docs/evidence/2026-06-06-auth-runtime-local-dev.md
```

Related handoffs:

```text
docs/handoffs/2026-06-06-manual-auth-login.md
docs/handoffs/2026-06-06-auth-runtime-local-dev.md
docs/handoffs/2026-06-07-capability-contracts.md
docs/handoffs/2026-06-07-capability-postgres-state.md
```

## Current Decision

Do not start broad Laravel code translation yet.

The valid order is:

1. Keep the progress ledger current.
2. Complete capability-control foundation proof.
3. Repair or prove security audit gate.
4. Complete missing Stage 0 inventory batches.
5. Start `servicecatalog`.
6. Start `productcatalog`.

Protected POS endpoints must wait for capability-control proof.

## Stage Progress

| Stage | Scope | Status | Estimate | Proof |
| --- | --- | --- | --- | --- |
| Stage 0 | Laravel source inventory and parity matrix | Partial | 40% | `0001_laravel_stage0_schema_and_route_inventory.md`, `0002_laravel_productcatalog_servicecatalog_inventory.md` |
| Stage 1 | Go quality foundation | Partial | 85% | `make verify` passes, including tests, vet, format, AI rules, file-size, hexagonal, and gosec; route-to-capability audit is still pending |
| Stage 2 | PostgreSQL target baseline for POS domains | Not started | 0% | No accepted POS PostgreSQL migration baseline proof yet |
| Stage 3 | API foundation and capability control | Partial | 45% | Auth/session foundation exists; capability contracts pass tests; PostgreSQL capability migration is applied; PostgreSQL adapter integration tests pass; middleware, HTTP surface, and route audit remain missing |
| Stage 4 | Cross-cutting modules | Not started | 0% | No audit/language/notification/idempotency transition implementation proof yet |
| Business Phase 1 | Service catalog and product catalog | Not started | 0% | Catalog evidence and blueprint exist; Go business modules not implemented |
| Overall Laravel-to-Go transition | POS API migration | Early foundation | 18% | Docs, auth debug lane, full verify gate, capability contracts, and capability PostgreSQL state exist; POS domains are not implemented |

## Completed Work With Proof

- Docs consolidation and AI workflow rules exist under `docs/`.
- Codex, web AI, analysis, testing, evidence, and resume templates exist under `docs/templates/`.
- Web AI GitHub connector rules are documented as read-only by default.
- Manual debug login foundation is documented in `docs/handoffs/2026-06-06-manual-auth-login.md`.
- Manual debug accounts are documented as:
  - `admin@example.com` with password `12345678`;
  - `kasir@example.com` with password `12345678`.
- Quality and architecture audit scripts are documented as wired.
- Capability domain, port, and usecase contracts exist under `internal/modules/capability/`.
- Capability contracts have unit test proof in domain and usecase packages.
- Full `make verify` proof passes, including gosec.
- Capability PostgreSQL migration `0006_capability_control.up.sql` is applied locally.
- Capability PostgreSQL repository integration tests pass.

## Open Gaps

- Full Laravel source inventory is incomplete for many business domains.
- Laravel alter, foreign key, index, timestamp, and seed migrations are not fully inventoried.
- Product duplicate policy still needs an owner decision before final PostgreSQL indexes.
- Capability-control foundation is partially implemented; runtime middleware, admin HTTP surface, route audit, and disabled-endpoint API proof are still missing.
- Runtime DB proof for manual auth login is still incomplete.
- No POS domain PostgreSQL baseline has been accepted.
- No `servicecatalog` or `productcatalog` Go business module has implementation proof.

## Next Valid Active Step

Continue `docs/blueprints/0010_capability_control_foundation.md` before exposing protected POS endpoints.

Minimum next-step proof:

- runtime capability middleware tests;
- disabled capability stops before handler/usecase execution;
- enabled capability allows handler execution;
- `make verify` remains passing;
- updated handoff and progress ledger.

## Handoff Requirement

Any Codex or web AI session that changes Laravel-to-Go transition docs, capability foundation, quality gates, or POS domain implementation must update this ledger or explicitly state why the ledger is unchanged.

The same session must create or update a handoff when durable work was done.

## Context Window Status

Current ledger update context status: enough context for one focused next implementation step after capability PostgreSQL state.
