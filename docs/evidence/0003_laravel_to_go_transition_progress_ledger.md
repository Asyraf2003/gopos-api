# Laravel To Go Transition Progress Ledger

## Status

Date updated: 2026-06-08

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
docs/blueprints/0024_servicecatalog_domain_contract.md
docs/blueprints/0025_servicecatalog_implementation_slice_1.md
docs/blueprints/0026_servicecatalog_postgres_persistence_slice.md
```

Related evidence:

```text
docs/evidence/0001_laravel_stage0_schema_and_route_inventory.md
docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md
docs/evidence/2026-06-06-auth-runtime-local-dev.md
docs/evidence/0004_adr_implementation_proof_index.md
docs/evidence/0005_laravel_to_go_transition_history_2026_06_08.md
```

Related handoffs:

```text
docs/handoffs/2026-06-06-manual-auth-login.md
docs/handoffs/2026-06-06-auth-runtime-local-dev.md
docs/handoffs/2026-06-07-capability-contracts.md
docs/handoffs/2026-06-07-capability-postgres-state.md
docs/handoffs/2026-06-07-capability-runtime-middleware.md
docs/handoffs/2026-06-07-capability-route-seeds.md
docs/handoffs/2026-06-07-capability-admin-http-surface.md
docs/handoffs/2026-06-07-prompt-template-selection-rule.md
docs/handoffs/2026-06-07-ai-execution-channel-boundaries.md
docs/handoffs/2026-06-07-web-ai-owner-terminal-output-test.md
docs/handoffs/2026-06-08-capability-route-audit-script.md
docs/handoffs/2026-06-08-capability-route-disabled-proof.md
docs/handoffs/2026-06-08-capability-control-closeout.md
docs/handoffs/2026-06-08-servicecatalog-domain-contract-blueprint.md
docs/handoffs/2026-06-08-servicecatalog-domain-contract-accepted.md
docs/handoffs/2026-06-08-servicecatalog-implementation-slice-1-plan.md
docs/handoffs/2026-06-08-servicecatalog-implementation-slice-1-accepted.md
docs/handoffs/2026-06-08-servicecatalog-implementation-slice-1.md
docs/handoffs/2026-06-08-servicecatalog-postgres-persistence-blueprint.md
docs/handoffs/2026-06-08-docs-quality-feedback-crosscheck.md
docs/handoffs/2026-06-08-docs-scalability-blueprint-cleanup.md
```

## Current Decision

Do not start broad Laravel code translation yet.

Capability-control foundation is closed with owner/local proof on 2026-06-08.

The valid order is:

1. Keep the progress ledger current.
2. Start the first POS business-domain blueprint/domain contract.
3. Repair or prove remaining runtime/security proof gaps as their own scopes.
4. Complete missing Stage 0 inventory batches before domain implementation.
5. Start `servicecatalog` only after an accepted domain contract and capability seed decision.
6. Start `productcatalog` only after an accepted domain contract and capability seed decision.

Protected POS CRUD implementation must wait for accepted domain contracts, POS PostgreSQL baseline decisions, authorization/capability mapping, transaction/audit decisions, and tests.

## Stage Progress

| Stage | Scope | Status | Estimate | Proof |
| --- | --- | --- | --- | --- |
| Stage 0 | Laravel source inventory and parity matrix | Partial | 40% | `0001_laravel_stage0_schema_and_route_inventory.md`, `0002_laravel_productcatalog_servicecatalog_inventory.md` |
| Stage 1 | Go quality foundation | Partial | 90% | `make verify` passes, including tests, vet, format, AI rules, file-size, hexagonal, route-to-capability audit, and gosec |
| Stage 2 | PostgreSQL target baseline for POS domains | Not started | 0% | No accepted POS PostgreSQL migration baseline proof yet |
| Stage 3 | API foundation and capability control | Closed | 100% | Auth/session foundation exists; capability contracts pass tests; PostgreSQL capability migration is applied; PostgreSQL adapter integration tests pass; runtime capability middleware tests pass; protected route seed migration exists; admin HTTP surface implementation and full `make verify` proof pass; route-to-capability audit script exists and is wired into `make verify`; route-level disabled protected endpoint proof passes for current protected route capability keys; final closeout proof passed on 2026-06-08 |
| Stage 4 | Cross-cutting modules | Not started | 0% | No audit/language/notification/idempotency transition implementation proof yet |
| Business Phase 1 | Service catalog and product catalog | Partial | 25% | ServiceCatalog domain/usecase slice and PostgreSQL persistence slice are implemented with proof; ServiceCatalog HTTP/capability integration and ProductCatalog are not implemented |
| Overall Laravel-to-Go transition | POS API migration | Early foundation | 25% | Docs, auth debug lane, full verify gate, capability foundation, and ServiceCatalog domain/usecase plus PostgreSQL persistence implementation proof exist; POS HTTP business endpoints and ProductCatalog are not implemented |

## Current State Summary

- Capability-control foundation is closed with proof.
- ServiceCatalog domain contract is accepted.
- ServiceCatalog slice 1 domain, ports, usecase contracts, and unit tests are implemented with proof.
- `make verify` passes, including Go tests, go vet, format, AI rules, file-size, hexagonal import, route capability audit, and gosec.
- ADR implementation proof index exists at `docs/evidence/0004_adr_implementation_proof_index.md`.
- Detailed completed-work history is archived in `docs/evidence/0005_laravel_to_go_transition_history_2026_06_08.md`.
- ServiceCatalog PostgreSQL persistence slice is implemented and closed with proof.

## Open Gaps

- Full Laravel source inventory is incomplete for many business domains.
- Laravel alter, foreign key, index, timestamp, and seed migrations are not fully inventoried.
- Product duplicate policy still needs an owner decision before final PostgreSQL indexes.
- Runtime DB proof for manual auth login is still incomplete.
- ADR `0009` debug auth lane remains partial because manual auth runtime proof evidence is incomplete.
- ADR `0012` API output contract centralization remains partial because full response/error envelope coverage is not proven for every API surface.
- ADR closeout backlog:
  - ADR `0009`: close by 2026-06-15 or before the next auth runtime change, whichever comes first.
  - ADR `0012`: close before adding broad new HTTP surfaces beyond the next accepted ServiceCatalog slice.
- No POS domain PostgreSQL baseline has been accepted.
- ServiceCatalog domain contract is accepted.
- ServiceCatalog implementation slice 1 plan is accepted and implemented with proof.
- ServiceCatalog HTTP transport, route registration, request/response presenters, authorization/capability wiring, and capability seeds are not implemented.
- ServiceCatalog capability seed migration is deferred, but must be owned by a later accepted runtime/capability blueprint before any ServiceCatalog protected HTTP route registration.
- ProductCatalog domain contract has not been accepted yet.
- No `productcatalog` Go business module has implementation proof.

## Next Valid Active Step

Create and accept the next ServiceCatalog runtime/capability blueprint.

- Candidate accepted scope: ServiceCatalog HTTP transport, route registration, request/response presenters, permission seed rows, capability seed rows, route capability manifest updates, and disabled-capability proof.
- Do not implement ProductCatalog or UI in this slice.
- Do not register protected ServiceCatalog HTTP routes before permission seeds, capability seeds, and route capability manifest coverage are included in the accepted blueprint.
- ProductCatalog remains blocked until its own accepted domain contract and duplicate policy decision exist.
## Handoff Requirement

Any Codex or web AI session that changes Laravel-to-Go transition docs, capability foundation, quality gates, or POS domain implementation must update this ledger or explicitly state why the ledger is unchanged.

The same session must create or update a handoff when durable work was done.

## Context Window Status

Current ledger update context status: updated after ServiceCatalog PostgreSQL persistence closeout and docs scalability cleanup; enough context to plan the next runtime/capability blueprint.
