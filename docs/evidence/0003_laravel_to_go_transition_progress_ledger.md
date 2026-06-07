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
| Business Phase 1 | Service catalog and product catalog | Partial | 15% | ServiceCatalog slice 1 domain, ports, usecase contracts, and unit tests exist; ProductCatalog not implemented |
| Overall Laravel-to-Go transition | POS API migration | Early foundation | 22% | Docs, auth debug lane, full verify gate, capability contracts, capability PostgreSQL state, runtime capability middleware, protected route seeds, admin capability HTTP surface, route-to-capability audit, and ServiceCatalog slice 1 domain/usecase implementation proof exist; POS HTTP and PostgreSQL business-domain adapters are not implemented |

## Completed Work With Proof

- Docs consolidation and AI workflow rules exist under `docs/`.
- Codex, web AI, analysis, testing, evidence, and resume templates exist under `docs/templates/`.
- Web AI GitHub connector rules are documented as read-only by default.
- Prompt template selection rule exists so next-session prompts must select exactly one target agent and one matching template source.
- Hybrid Web AI/Codex next-session prompts are forbidden unless explicitly requested as a collaboration packet.
- AI execution channel boundaries are clarified: Web AI no longer defaults to Codex as executor, owner/local terminal command-plan loop is documented, and collaboration packet remains special-case only.
- A Web AI output test found residual Codex-default behavior: normal analysis still emitted `HANDOFF TEXT FOR CODEX`.
- Templates are tightened so normal Web AI analysis must prefer owner/local terminal command plans and omit Codex handoff unless explicitly requested.
- Prompt template hardening is a workflow/docs quality improvement and does not increase POS implementation progress.
- Prompt template hardening proof is recorded in `docs/handoffs/2026-06-07-prompt-template-selection-rule.md`; `make verify` passed after the docs change.
- AI execution-channel boundary proof is recorded in `docs/handoffs/2026-06-07-ai-execution-channel-boundaries.md`; `make verify` passed after the docs change.
- Web AI owner/local terminal output proof is recorded in `docs/handoffs/2026-06-07-web-ai-owner-terminal-output-test.md`; `make verify` passed after the docs change.
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
- Runtime capability middleware exists under `internal/transport/http/middleware`.
- Runtime capability middleware tests prove enabled capability allows handler execution, disabled capability returns `403` before handler execution, checker errors return `500`, and misconfigured guards return `500`.
- Protected route capability seed migration `migrations/0007_seed_existing_protected_capabilities.up.sql` exists for current protected routes.
- Migration `0008_seed_capability_manage_permission` adds `capability.manage`, assigns it to `admin`, and seeds `api_capabilities.key = 'capability.manage'`.
- Capability response DTO mapping exists under `internal/presentation/http/id/capability/`.
- Admin capability list/show/enable/disable handler exists under `internal/modules/capability/transport/http/`.
- Bootstrap wires `/api/admin/capabilities...` behind authn, `capability.manage` authorization, and runtime capability check.
- User-provided SQL proof confirmed `capability.manage` permission, admin role assignment, and `api_capabilities` row.
- Local proof confirmed capability handler files pass file-size audit, focused capability tests pass, bootstrap tests pass, and `make verify` passes.
- Route-to-capability audit script exists with manifest coverage for 6 current protected route capability rows.
- `make verify` summary now includes route capability audit and passes aggregate audit.
- Route-level disabled protected endpoint proof passes for current protected route capability keys.
- Capability-control foundation closeout proof passed on 2026-06-08 with focused tests, route capability audit, DB migration status, and `make verify`.
- DB migration status confirmed `0006_capability_control.up.sql`, `0007_seed_existing_protected_capabilities.up.sql`, and `0008_seed_capability_manage_permission.up.sql` are applied.
- Final closeout `make verify` passed with gosec reporting 97 files, 3978 lines, 0 nosec, and 0 issues.
- Docs quality feedback crosscheck added a 5-minute quick reference, evidence status index, incomplete auth runtime evidence gap marker, blueprint/log boundary rule, and concrete ServiceCatalog scope packet example.
- This docs quality improvement does not increase Laravel-to-Go implementation progress.
- ADR implementation proof index exists at `docs/evidence/0004_adr_implementation_proof_index.md`; ADR proof tracking is now explicit and does not increase POS implementation progress.
- ServiceCatalog slice 1 implementation exists under `internal/modules/servicecatalog/domain`, `internal/modules/servicecatalog/ports`, and `internal/modules/servicecatalog/usecase`.
- ServiceCatalog domain rules prove name normalization, blank-name rejection, positive default price validation, and active-by-default creation.
- ServiceCatalog ports define repository contracts for create, update, find by ID, find by normalized name, list, lookup, and set active.
- ServiceCatalog usecase tests prove create, duplicate rejection, update, lifecycle activation/deactivation, missing show not-found, list status filtering, lookup active-only default behavior, and lookup max-limit enforcement.
- Focused proof passed: `go test ./internal/modules/servicecatalog/...`.
- Full proof passed: `make verify`, including `go test ./...`, go vet, format, AI rules, file-size, hexagonal import, route capability, and gosec audits.
- Final `make verify` gosec summary reported 112 files, 4659 lines, 0 nosec, and 0 issues.
- ServiceCatalog PostgreSQL persistence slice blueprint exists at `docs/blueprints/0026_servicecatalog_postgres_persistence_slice.md`; it is proposed only and does not increase implementation progress.

## Open Gaps

- Full Laravel source inventory is incomplete for many business domains.
- Laravel alter, foreign key, index, timestamp, and seed migrations are not fully inventoried.
- Product duplicate policy still needs an owner decision before final PostgreSQL indexes.
- Runtime DB proof for manual auth login is still incomplete.
- ADR `0009` debug auth lane remains partial because manual auth runtime proof evidence is incomplete.
- ADR `0012` API output contract centralization remains partial because full response/error envelope coverage is not proven for every API surface.
- No POS domain PostgreSQL baseline has been accepted.
- ServiceCatalog domain contract is accepted.
- ServiceCatalog implementation slice 1 plan is accepted and implemented with proof.
- ServiceCatalog HTTP transport, PostgreSQL adapter, migrations, route registration, and capability seeds are not implemented.
- ServiceCatalog PostgreSQL persistence slice blueprint `0026` is proposed but not accepted or implemented yet.
- ProductCatalog domain contract has not been accepted yet.
- No `productcatalog` Go business module has implementation proof.

## Next Valid Active Step

Accept or revise `docs/blueprints/0026_servicecatalog_postgres_persistence_slice.md`.

- Candidate accepted scope: ServiceCatalog PostgreSQL migration, PostgreSQL repository adapter, and repository/integration tests only.
- Do not implement HTTP transport, route registration, capability seed migrations, ProductCatalog, or UI in this slice.
- ProductCatalog remains blocked until its own accepted domain contract and duplicate policy decision exist.
## Handoff Requirement

Any Codex or web AI session that changes Laravel-to-Go transition docs, capability foundation, quality gates, or POS domain implementation must update this ledger or explicitly state why the ledger is unchanged.

The same session must create or update a handoff when durable work was done.

## Context Window Status

Current ledger update context status: updated after ServiceCatalog PostgreSQL persistence blueprint proposal; enough context to accept or revise blueprint `0026` before implementation.
