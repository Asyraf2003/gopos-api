# Laravel To Go Transition Progress Ledger

## Status

Date updated: 2026-06-11

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
docs/blueprints/0027_servicecatalog_runtime_capability_slice.md
docs/blueprints/0030_productcatalog_postgres_persistence_slice.md
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
docs/handoffs/2026-06-06-auth-runtime-local-dev.md
docs/handoffs/2026-06-08-servicecatalog-runtime-capability-blueprint-accepted.md
docs/handoffs/2026-06-08-servicecatalog-runtime-capability-implementation.md
docs/handoffs/2026-06-09-productcatalog-domain-slice-1.md
docs/handoffs/2026-06-10-productcatalog-listproducts-skeleton-progress.md
docs/handoffs/2026-06-10-productcatalog-lookupproducts-skeleton-progress.md
docs/handoffs/2026-06-10-productcatalog-listproductversions-skeleton-progress.md
docs/handoffs/2026-06-10-productcatalog-listproductversions-behavior-progress.md
docs/handoffs/2026-06-11-productcatalog-postgres-create-find-update-progress.md
docs/archive/handoffs-closed/README.md
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
| Stage 2 | PostgreSQL target baseline for POS domains | Partial | 10% | ProductCatalog PostgreSQL migration 0011 has local DB apply proof; ProductCatalog PostgreSQL repository skeletons are remote-visible with compile-time port assertions; ProductRepository Create, FindByID, and Update behavior have focused local and integration proof; aggregate `make verify` proof pending |
| Stage 3 | API foundation and capability control | Closed | 100% | Auth/session foundation exists; capability contracts pass tests; PostgreSQL capability migration is applied; PostgreSQL adapter integration tests pass; runtime capability middleware tests pass; protected route seed migration exists; admin HTTP surface implementation and full `make verify` proof pass; route-to-capability audit script exists and is wired into `make verify`; route-level disabled protected endpoint proof passes for current protected route capability keys; final closeout proof passed on 2026-06-08 |
| Stage 4 | Cross-cutting modules | Not started | 0% | No audit/language/notification/idempotency transition implementation proof yet |
| Business Phase 1 | Service catalog and product catalog | Partial | 45% | ServiceCatalog domain/usecase, PostgreSQL persistence, and runtime/capability slice have local proof; ProductCatalog domain, ports, CreateProduct, UpdateProduct, SoftDeleteProduct, RestoreProduct, GetProductDetail, ListProducts, LookupProducts, ListProductVersions, and ProductRepository Create/FindByID/Update behavior have local focused proof; connector validation passed for the latest ProductCatalog behavior checkpoint; aggregate `make verify` proof for repository behavior checkpoint is pending |
| Overall Laravel-to-Go transition | POS API migration | Early foundation | 33% | Docs, auth debug lane, full verify gate, capability foundation, ServiceCatalog domain/usecase, PostgreSQL persistence, runtime/capability proof, ProductCatalog domain/usecase proof, and first ProductCatalog PostgreSQL repository behavior proof exist; ProductCatalog HTTP/runtime/capability/UI and broader POS APIs remain incomplete |

## Current State Summary
- Capability-control foundation is closed with proof.
- ServiceCatalog domain contract is accepted.
- ServiceCatalog slice 1 domain, ports, usecase contracts, and unit tests are implemented with proof.
- `make verify` passes, including Go tests, go vet, format, AI rules, file-size, hexagonal import, route capability audit, and gosec.
- ADR implementation proof index exists at `docs/evidence/0004_adr_implementation_proof_index.md`.
- Detailed completed-work history is archived in `docs/evidence/0005_laravel_to_go_transition_history_2026_06_08.md`.
- ServiceCatalog PostgreSQL persistence slice is implemented and closed with proof.
- ServiceCatalog runtime/capability blueprint 0027 is accepted and locally proven through focused handler tests, ServiceCatalog-specific disabled-capability proof, and aggregate audit.
- ServiceCatalog runtime/capability slice 0027 is locally implemented with proof.
- ServiceCatalog HTTP transport, route registration, request/response presenters, authorization/capability wiring, route capability manifest coverage, audit coverage, and capability/permission seed migration 0010 have local proof.
- Migration `0010_seed_service_catalog_permissions_capabilities.up.sql` is applied in local database proof.
- ProductCatalog ListProducts contract, constructor/skeleton, reader error propagation, query forwarding, success mapping, and empty-list behavior are remote-visible through GitHub connector with focused `go test ./internal/modules/productcatalog/...` proof and aggregate `make verify` proof.
- ProductCatalog LookupProducts contract, constructor/skeleton, reader error propagation, query forwarding, success mapping, and empty-list behavior are remote-visible through GitHub connector with focused `go test ./internal/modules/productcatalog/...` proof and aggregate `make verify` proof.
- ProductCatalog ListProductVersions behavior is locally implemented with focused `go test ./internal/modules/productcatalog/...` proof and aggregate `make verify` proof; connector validation passed through GitHub connector.

## Open Gaps
- Full Laravel source inventory is incomplete for many business domains.
- Laravel alter, foreign key, index, timestamp, and seed migrations are not fully inventoried.
- Product duplicate policy Option A is accepted locally with proof; final PostgreSQL indexes must preserve Laravel-compatible duplicate behavior.
- Runtime DB proof for manual auth login is still incomplete.
- ADR `0009` debug auth lane remains partial because manual auth runtime proof evidence is incomplete.
- ADR `0012` API output contract centralization remains partial because full response/error envelope coverage is not proven for every API surface.
- ADR closeout backlog:
  - ADR `0009`: close by 2026-06-15 or before the next auth runtime change, whichever comes first.
  - ADR `0012`: close before adding broad new HTTP surfaces beyond the next accepted ServiceCatalog slice.
- No POS domain PostgreSQL baseline has been accepted.
- ServiceCatalog domain contract is accepted.
- ServiceCatalog implementation slice 1 plan is accepted and implemented with proof.
- ProductCatalog domain, ports, CreateProduct, UpdateProduct, SoftDeleteProduct, RestoreProduct, GetProductDetail, ListProducts, LookupProducts, and ListProductVersions behavior are locally proven; connector validation passed for the latest behavior checkpoint.
- ProductCatalog PostgreSQL repository behavior is partial: ProductRepository Create, FindByID, and Update behavior have focused local/integration proof and connector validation; ProductReader, ProductVersionRepository, ProductDuplicateChecker, EXPLAIN/query-plan, runtime HTTP surface, route registration, presenter, capability seed, inventory mutation, and UI work are not started yet. ProductCatalog PostgreSQL migration 0011 has local DB apply proof and repository skeletons are remote-visible.
- ServiceCatalog runtime/capability implementation is remote-visible through GitHub connector with local proof; focused handler and disabled-capability proof are remote-visible through GitHub connector with local proof; connector validation passed for the latest closeout proof files.
- ProductCatalog domain contract blueprint `docs/blueprints/0028_productcatalog_domain_contract.md` is accepted locally with Option A duplicate policy and `make verify` proof; connector validation pending.
- ProductCatalog implementation slice 1 blueprint `docs/blueprints/0029_productcatalog_implementation_slice_1.md` is accepted locally with `make verify` proof; connector validation pending.
- ProductCatalog domain package, ports, CreateProduct, UpdateProduct, SoftDeleteProduct, RestoreProduct, GetProductDetail, ListProducts, LookupProducts, and ListProductVersions behavior are locally proven with focused `go test ./internal/modules/productcatalog/...` proof and aggregate `make verify` proof; connector validation passed for the latest behavior checkpoint.
- ProductCatalog ListProducts query forwarding, success item mapping, and empty-list behavior are remote-visible through GitHub connector with focused `go test ./internal/modules/productcatalog/...` proof and aggregate `make verify` proof.
- ProductCatalog implementation slice 1 is closed after ListProductVersions behavior connector validation.
- ProductCatalog PostgreSQL persistence blueprint `docs/blueprints/0030_productcatalog_postgres_persistence_slice.md` is accepted. ProductCatalog PostgreSQL persistence blueprint includes performance and flexibility standards for CRUD, show, list, lookup, duplicate guard, and version-list query paths. ProductCatalog migration `0011_create_product_catalog_tables.up.sql` has local DB apply proof; ProductCatalog PostgreSQL repository skeletons are remote-visible with compile-time port assertions; ProductRepository Create, FindByID, and Update behavior are implemented with focused PostgreSQL integration proof and connector validation. Remaining ProductReader, ProductVersionRepository, ProductDuplicateChecker, EXPLAIN/query-plan, HTTP/runtime/capability/UI work is not started yet.

## Next Valid Active Step

Run aggregate proof for ProductRepository Create, FindByID, and Update behavior before starting the next ProductCatalog PostgreSQL repository behavior step.

- Do not start ProductReader List/Lookup, ProductVersionRepository, ProductDuplicateChecker, EXPLAIN/query-plan, Echo/runtime, capability seed, inventory mutation, UI, or runtime HTTP work before aggregate proof and progress docs are validated.
- Do not start a new runtime slice while repository proof is not reflected in repository facts.
- Stay inside blueprint-allowed ProductCatalog PostgreSQL repository behavior only after aggregate proof passes.

## Handoff Requirement

Any Codex or web AI session that changes Laravel-to-Go transition docs, capability foundation, quality gates, or POS domain implementation must update this ledger or explicitly state why the ledger is unchanged.

The same session must create or update a handoff when durable work was done.

## Context Window Status

Current ledger update context status: updated after ProductRepository Create, FindByID, and Update behavior received focused local and integration proof; next step is aggregate proof before starting the next repository behavior step.
