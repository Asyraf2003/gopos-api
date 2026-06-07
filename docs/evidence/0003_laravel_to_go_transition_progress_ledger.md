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
docs/handoffs/2026-06-07-capability-runtime-middleware.md
docs/handoffs/2026-06-07-capability-route-seeds.md
docs/handoffs/2026-06-07-capability-admin-http-surface.md
docs/handoffs/2026-06-07-prompt-template-selection-rule.md
docs/handoffs/2026-06-07-ai-execution-channel-boundaries.md
docs/handoffs/2026-06-07-web-ai-owner-terminal-output-test.md
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
| Stage 3 | API foundation and capability control | Partial | 75% | Auth/session foundation exists; capability contracts pass tests; PostgreSQL capability migration is applied; PostgreSQL adapter integration tests pass; runtime capability middleware tests pass; protected route seed migration exists; admin HTTP surface implementation and full `make verify` proof pass; route audit remains missing |
| Stage 4 | Cross-cutting modules | Not started | 0% | No audit/language/notification/idempotency transition implementation proof yet |
| Business Phase 1 | Service catalog and product catalog | Not started | 0% | Catalog evidence and blueprint exist; Go business modules not implemented |
| Overall Laravel-to-Go transition | POS API migration | Early foundation | 20% | Docs, auth debug lane, full verify gate, capability contracts, capability PostgreSQL state, runtime capability middleware, protected route seeds, and admin capability HTTP surface exist; POS domains are not implemented |

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

## Open Gaps

- Full Laravel source inventory is incomplete for many business domains.
- Laravel alter, foreign key, index, timestamp, and seed migrations are not fully inventoried.
- Product duplicate policy still needs an owner decision before final PostgreSQL indexes.
- Capability-control foundation is partially implemented; route audit and disabled route-level API proof are still missing.
- Runtime DB proof for manual auth login is still incomplete.
- No POS domain PostgreSQL baseline has been accepted.
- No `servicecatalog` or `productcatalog` Go business module has implementation proof.

## Next Valid Active Step

Continue `docs/blueprints/0010_capability_control_foundation.md` before exposing protected POS endpoints.

- add route-to-capability audit script;
- preserve current admin capability HTTP proof;
- do not start POS CRUD before capability-control proof is complete.

## Handoff Requirement

Any Codex or web AI session that changes Laravel-to-Go transition docs, capability foundation, quality gates, or POS domain implementation must update this ledger or explicitly state why the ledger is unchanged.

The same session must create or update a handoff when durable work was done.

## Context Window Status

Current ledger update context status: enough context for the next capability-control foundation step after admin capability HTTP proof.
