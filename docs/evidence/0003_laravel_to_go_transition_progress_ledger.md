<!--
Copyright (C) 2026 Asyraf Mubarak

This file is part of gopos-api.

gopos-api is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, version 3 only.

gopos-api is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

audit:allow-oversize reason=bootstrap-wiring
-->

# Laravel To Go Transition Progress Ledger

## Status

Date updated: 2026-06-14

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
docs/blueprints/0028_productcatalog_domain_contract.md
docs/blueprints/0029_productcatalog_implementation_slice_1.md
docs/blueprints/0030_productcatalog_postgres_persistence_slice.md
docs/blueprints/0031_productcatalog_runtime_capability_slice.md
docs/blueprints/0032_api_docs_error_envelope_slice.md
docs/blueprints/0033_productcatalog_runtime_smoke_proof_slice.md
docs/blueprints/0037_supplier_domain_contract.md
docs/blueprints/0038_supplier_implementation_slice_1.md
docs/blueprints/0039_supplier_postgres_persistence_slice.md
```

Related evidence:

```text
docs/evidence/0001_laravel_stage0_schema_and_route_inventory.md
docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md
docs/evidence/2026-06-06-auth-runtime-local-dev.md
docs/evidence/0004_adr_implementation_proof_index.md
docs/evidence/0005_laravel_to_go_transition_history_2026_06_08.md
docs/evidence/2026-06-13_api_architecture_product_status_review.md
docs/evidence/2026-06-14_productcatalog_runtime_smoke_proof.md
```

Related handoffs:

```text
docs/handoffs/2026-06-06-auth-runtime-local-dev.md
docs/handoffs/2026-06-08-servicecatalog-runtime-capability-blueprint-accepted.md
docs/handoffs/2026-06-08-servicecatalog-runtime-capability-implementation.md
docs/handoffs/2026-06-14-productcatalog-runtime-smoke-proof-closeout.md
docs/handoffs/2026-06-14-supplier-implementation-slice-1-closeout.md
docs/handoffs/2026-06-14-supplier-postgres-persistence-migration-only.md
docs/archive/handoffs-closed/README.md
```

## Current Decision

Do not start broad Laravel code translation yet.

Capability-control foundation is closed with owner/local proof on 2026-06-08.

The valid order is:

1. Keep the progress ledger current.
2. Select exactly one accepted slice before implementation.
3. Repair or prove remaining runtime/security proof gaps as their own scopes.
4. Complete missing Stage 0 inventory batches before domain implementation.
5. Do not re-open ProductCatalog catalog API, persistence, runtime/capability, API docs, error envelope, or runtime smoke proof work unless a bug is found.
6. Do not treat ProductCatalog catalog API closeout as Product/inventory area closeout.

Protected POS CRUD implementation must wait for accepted domain contracts, POS PostgreSQL baseline decisions, authorization/capability mapping, transaction/audit decisions, and tests.

## Stage Progress

| Stage | Scope | Status | Estimate | Proof |
| --- | --- | --- | --- | --- |
| Stage 0 | Laravel source inventory and parity matrix | Partial | 40% | `0001_laravel_stage0_schema_and_route_inventory.md`, `0002_laravel_productcatalog_servicecatalog_inventory.md` |
| Stage 1 | Go quality foundation | Partial | 90% | `make verify` passes, including tests, vet, format, AI rules, file-size, hexagonal, route-to-capability audit, and gosec |
| Stage 2 | PostgreSQL target baseline for POS domains | Partial | 10% | ProductCatalog PostgreSQL migration 0011 has local DB apply proof; ProductCatalog PostgreSQL repository skeletons are remote-visible with compile-time port assertions; ProductRepository Create, FindByID, and Update behavior have focused local, integration, and aggregate `make verify` proof; ProductReader GetByID behavior has focused local, integration, and aggregate `make verify` proof; ProductReader List behavior has focused local, integration, and aggregate `make verify` proof; ProductReader Lookup behavior has focused integration, aggregate `make verify`, and connector proof; ProductVersionRepository Append/ListByProductID behavior has focused integration, aggregate `make verify`, and connector proof; ProductDuplicateChecker behavior has focused integration, aggregate `make verify`, and connector proof; ProductCatalog PostgreSQL EXPLAIN/query-plan proof passed for key read paths; ProductCatalog PostgreSQL EXPLAIN/query-plan proof passed for key read paths |
| Stage 3 | API foundation and capability control | Closed | 100% | Auth/session foundation exists; capability contracts pass tests; PostgreSQL capability migration is applied; PostgreSQL adapter integration tests pass; runtime capability middleware tests pass; protected route seed migration exists; admin HTTP surface implementation and full `make verify` proof pass; route-to-capability audit script exists and is wired into `make verify`; route-level disabled protected endpoint proof passes for current protected route capability keys; final closeout proof passed on 2026-06-08 |
| Stage 4 | Cross-cutting modules | Not started | 0% | No audit/language/notification/idempotency transition implementation proof yet |
| Business Phase 1 | Service catalog and product catalog | Partial | 64% | ServiceCatalog domain/usecase, PostgreSQL persistence, and runtime/capability slice have local proof; ProductCatalog catalog API has local proof through domain/usecase, PostgreSQL persistence, protected runtime/capability routes, developer API docs, standardized ProductCatalog error envelope coverage, shared success envelope coverage for ProductCatalog/ServiceCatalog/Capability, Product API readiness for Supplier/Faktur planning, Supplier domain contract acceptance, and real local HTTP/auth/DB runtime smoke proof; Product inventory/stock mutation, stock adjustment create/reverse, ProductCatalog audit/outbox persistence, runtime language switch, extended filters, Supplier implementation, and Faktur remain incomplete |
| Overall Laravel-to-Go transition | POS API migration | Early foundation | 41% | Docs, auth debug lane, full verify gate, capability foundation, ServiceCatalog domain/usecase, PostgreSQL persistence, runtime/capability proof, ProductCatalog domain/usecase proof, ProductCatalog PostgreSQL persistence closeout proof, ProductCatalog runtime/capability closeout proof, ProductCatalog API docs/error envelope proof, and ProductCatalog runtime smoke proof exist; Product/inventory behavior, audit/outbox, language switch, extended filters, ADR 0012 full output centralization, and broader POS APIs remain incomplete |

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
- ProductCatalog catalog API is locally closed for pure API scope through domain/usecase, PostgreSQL persistence, protected runtime/capability routes, developer API docs, and standardized ProductCatalog error envelope proof.
- ProductCatalog runtime smoke proof is locally proven through local PostgreSQL, migrations through `0013_seed_product_catalog_permissions_capabilities.up.sql`, Echo server on port `18081`, manual auth token presence, protected ProductCatalog list/lookup/create/show routes, direct `products` table row proof, and reversible `product_catalog.list` capability disable/re-enable proof.
- Superseded ProductCatalog implementation handoffs have been archived under `docs/archive/handoffs-closed/`; current ProductCatalog runtime smoke closeout context is recorded in `docs/handoffs/2026-06-14-productcatalog-runtime-smoke-proof-closeout.md`.

## Open Gaps
- Full Laravel source inventory is incomplete for many business domains.
- Laravel alter, foreign key, index, timestamp, and seed migrations are not fully inventoried.
- Product duplicate policy Option A is accepted locally with proof; final PostgreSQL indexes must preserve Laravel-compatible duplicate behavior.
- ADR `0009` full debug auth lane remains partial; ProductCatalog runtime smoke proof includes manual login token issuance only for the ProductCatalog runtime path and does not prove the full auth runtime surface.
- ADR `0012` API output contract centralization remains partial because Auth/System output contract centralization is deferred by owner decision and full response/error envelope coverage is not proven for every API surface.
- ADR closeout backlog:
  - ADR `0009`: close by 2026-06-15 or before the next auth runtime change, whichever comes first.
  - ADR `0012`: close before adding broad new HTTP surfaces beyond the next accepted ServiceCatalog slice.
- No POS domain PostgreSQL baseline has been accepted.
- ServiceCatalog domain contract is accepted.
- ServiceCatalog implementation slice 1 plan is accepted and implemented with proof.
- ProductCatalog domain, ports, CreateProduct, UpdateProduct, SoftDeleteProduct, RestoreProduct, GetProductDetail, ListProducts, LookupProducts, and ListProductVersions behavior are locally proven; connector validation passed for the latest behavior checkpoint.
- ProductCatalog catalog API path is locally proven through domain/usecase, PostgreSQL persistence, runtime/capability, developer API docs, standardized ProductCatalog error envelope slices, shared success envelope foundation for ProductCatalog/ServiceCatalog/Capability, Product API readiness for Supplier/Faktur planning, Supplier domain contract acceptance, and end-to-end local HTTP/auth/DB smoke proof. Product inventory/stock mutation, stock adjustment create/reverse, ProductCatalog audit/outbox persistence, runtime language switch, extended Laravel filters, Supplier implementation, Faktur, and deferred broader ADR `0012` Auth/System output contract centralization remain incomplete.
- ServiceCatalog runtime/capability implementation is remote-visible through GitHub connector with local proof; focused handler and disabled-capability proof are remote-visible through GitHub connector with local proof; connector validation passed for the latest closeout proof files.
- ProductCatalog domain contract blueprint `docs/blueprints/0028_productcatalog_domain_contract.md` is accepted locally with Option A duplicate policy and `make verify` proof; connector validation pending.
- ProductCatalog implementation slice 1 blueprint `docs/blueprints/0029_productcatalog_implementation_slice_1.md` is accepted locally with `make verify` proof; connector validation pending.
- ProductCatalog domain package, ports, CreateProduct, UpdateProduct, SoftDeleteProduct, RestoreProduct, GetProductDetail, ListProducts, LookupProducts, and ListProductVersions behavior are locally proven with focused `go test ./internal/modules/productcatalog/...` proof and aggregate `make verify` proof; connector validation passed for the latest behavior checkpoint.
- ProductCatalog ListProducts query forwarding, success item mapping, and empty-list behavior are remote-visible through GitHub connector with focused `go test ./internal/modules/productcatalog/...` proof and aggregate `make verify` proof.
- ProductCatalog implementation slice 1 is closed after ListProductVersions behavior connector validation.
- ProductCatalog PostgreSQL persistence blueprint `docs/blueprints/0030_productcatalog_postgres_persistence_slice.md` is accepted. ProductCatalog PostgreSQL persistence blueprint includes performance and flexibility standards for CRUD, show, list, lookup, duplicate guard, and version-list query paths. ProductCatalog migration `0011_create_product_catalog_tables.up.sql` has local DB apply proof; ProductCatalog PostgreSQL repository skeletons are remote-visible with compile-time port assertions; ProductRepository Create, FindByID, Update, ProductReader GetByID, ProductReader List, ProductReader Lookup, ProductVersionRepository Append/ListByProductID, and ProductDuplicateChecker behavior are implemented with focused PostgreSQL integration proof and connector validation. EXPLAIN/query-plan proof passed for key ProductCatalog PostgreSQL read paths.
- ProductCatalog protected runtime/capability blueprint `docs/blueprints/0031_productcatalog_runtime_capability_slice.md` and API docs/error envelope blueprint `docs/blueprints/0032_api_docs_error_envelope_slice.md` are locally implemented with proof.
- ProductCatalog runtime smoke proof blueprint `docs/blueprints/0033_productcatalog_runtime_smoke_proof_slice.md` is locally proven by `docs/evidence/2026-06-14_productcatalog_runtime_smoke_proof.md`.



### Supplier implementation slice 1 closeout - 2026-06-14

Business Phase 1 progress updated: Supplier implementation slice 1 domain/ports/usecase is locally implemented with proof and remote connector validation checked.

Proof:

```bash
go test ./internal/modules/supplier/...
bash scripts/audit_hexagonal.sh
make verify
```

Remaining open gaps:

- Supplier PostgreSQL persistence.
- Supplier HTTP routes.
- Supplier capability seed.
- Faktur.
- Inventory/stock movement.
- Audit/outbox.
- Localization.
- Extended filters.

Auth/System ADR 0012 output contract centralization remains deferred by owner decision.

Next Valid Active Step: Supplier PostgreSQL persistence blueprint.

### Supplier PostgreSQL persistence migration-only checkpoint - 2026-06-14

Supplier PostgreSQL persistence blueprint 0039_supplier_postgres_persistence_slice.md is accepted locally.

Migration-only step is locally implemented and applied.

Files created:

```text
migrations/0014_create_suppliers_table.up.sql
migrations/0014_create_suppliers_table.down.sql
```

Local proof:

```bash
go test ./internal/modules/supplier/...
bash scripts/audit_hexagonal.sh
make verify
bash scripts/db_migrate.sh
```

Visible migration proof:

```text
[APPLY] 0014_create_suppliers_table.up.sql
CREATE TABLE
CREATE INDEX
CREATE INDEX
CREATE INDEX
CREATE INDEX
[PASS] db migrate completed
```

Current Supplier PostgreSQL persistence status:

- blueprint accepted;
- migration-only step complete;
- repository adapter skeletons locally implemented with compile-time port assertion;
- integration tests not started;
- repository behavior implementation not started;
- query-plan proof not collected;
- remote connector validation pending for final local changes.

Remaining open gaps:

- Supplier PostgreSQL repository behavior.
- Supplier repository integration tests.
- Supplier query-plan proof.
- Supplier HTTP routes.
- Supplier capability seed.
- Faktur.
- Inventory/stock movement.
- Audit/outbox.
- Localization.
- Extended filters.
- Laravel Supplier MySQL/source parity.

Auth/System ADR 0012 output contract centralization remains deferred by owner decision.

Next Valid Active Step: Supplier PostgreSQL repository Create and FindByID behavior.

### Supplier PostgreSQL repository adapter skeleton checkpoint - 2026-06-14

Supplier PostgreSQL repository adapter skeletons are locally implemented.

Files created:

```text
internal/platform/postgres/supplier_repository.go
internal/platform/postgres/supplier_repository_row.go
internal/platform/postgres/supplier_repository_write.go
internal/platform/postgres/supplier_repository_query.go
```

Local proof:

```bash
go test ./internal/modules/supplier/...
go test ./internal/platform/postgres/... -run Supplier
bash scripts/audit_hexagonal.sh
make verify
```

Current Supplier PostgreSQL persistence status:

- blueprint accepted;
- migration-only step complete;
- repository adapter skeletons implemented;
- compile-time port assertion exists;
- repository behavior is still pending;
- integration tests are not implemented;
- query-plan proof is not collected;
- remote connector validation pending for local changes.

Remaining open gaps:

- Supplier PostgreSQL repository Create and FindByID behavior.
- Supplier PostgreSQL repository Update behavior.
- Supplier PostgreSQL repository FindByNormalizedName and FindActiveByNormalizedName behavior.
- Supplier PostgreSQL repository SetActive behavior.
- Supplier PostgreSQL repository List and Lookup behavior.
- Supplier repository integration tests.
- Supplier query-plan proof.
- Supplier HTTP routes.
- Supplier capability seed.
- Faktur.
- Inventory/stock movement.
- Audit/outbox.
- Localization.
- Extended filters.
- Laravel Supplier MySQL/source parity.

Next Valid Active Step: Supplier PostgreSQL repository Create and FindByID behavior.

### Supplier PostgreSQL repository create and direct lookup checkpoint - 2026-06-14

Supplier PostgreSQL repository Create, FindByID, FindByNormalizedName, and FindActiveByNormalizedName behavior is locally implemented with compile and aggregate proof.

Files changed:

```text
internal/platform/postgres/supplier_repository_write.go
internal/platform/postgres/supplier_repository_query.go
internal/platform/postgres/supplier_repository_integration_helpers_test.go
internal/platform/postgres/supplier_repository_create_integration_test.go
internal/platform/postgres/supplier_repository_query_integration_test.go
docs/handoffs/2026-06-14-supplier-postgres-persistence-migration-only.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

Local proof:

```bash
go test ./internal/modules/supplier/...
go test ./internal/platform/postgres/... -run Supplier
go test -tags integration ./internal/platform/postgres/... -run Supplier -count=1 -v
bash scripts/audit_hexagonal.sh
make verify
```

Proof status:

- Supplier module proof passed.
- PostgreSQL package compile proof passed.
- Integration-tagged Supplier repository tests compile and the command passed, but every Supplier DB test skipped because `DATABASE_URL` is not set in this shell.
- Hexagonal import audit passed.
- Aggregate `make verify` passed.

Current Supplier PostgreSQL persistence status:

- blueprint accepted;
- migration-only step complete;
- repository adapter skeletons implemented;
- Create behavior implemented;
- FindByID behavior implemented;
- FindByNormalizedName behavior implemented;
- FindActiveByNormalizedName behavior implemented;
- integration tests for create and direct lookup behavior added;
- DB-backed integration execution proof pending because `DATABASE_URL` is not set in this shell;
- query-plan proof is not collected;
- remote connector validation pending for local changes.

Remaining open gaps:

- Supplier PostgreSQL repository Update behavior.
- Supplier PostgreSQL repository SetActive behavior.
- Supplier PostgreSQL repository List and Lookup behavior.
- Supplier DB-backed integration execution proof.
- Supplier query-plan proof.
- Supplier HTTP routes.
- Supplier capability seed.
- Faktur.
- Inventory/stock movement.
- Audit/outbox.
- Localization.
- Extended filters.
- Laravel Supplier MySQL/source parity.

Next Valid Active Step: Supplier PostgreSQL repository Update behavior.


## Next Valid Active Step

Supplier PostgreSQL repository Update behavior.

- Start from accepted blueprint `docs/blueprints/0039_supplier_postgres_persistence_slice.md`.
- Keep the next active step limited to Update behavior plus focused repository proof.
- Do not start SetActive, List, Lookup, HTTP routes, capability seed, Faktur, inventory/stock movement, audit/outbox, localization, extended filters, or architecture folder cleanup.
- Do not re-open ProductCatalog persistence, runtime/capability, API docs, error envelope, shared success envelope foundation, Capability envelope, Product API readiness, or runtime smoke proof work unless a bug is found.

## Handoff Requirement

Any Codex or web AI session that changes Laravel-to-Go transition docs, capability foundation, quality gates, or POS domain implementation must update this ledger or explicitly state why the ledger is unchanged.

The same session must create or update a handoff when durable work was done.

## Context Window Status

Current ledger update context status: updated after Supplier PostgreSQL repository create and direct lookup checkpoint. Auth/System output contract centralization is deferred by owner decision. The next valid step is Supplier PostgreSQL repository Update behavior.

## 2026-06-13 ProductCatalog runtime/capability closeout

FACT:
ProductCatalog protected HTTP runtime, presenter, route registration, permission seed, capability seed, route capability manifest, and disabled-capability proof are locally complete.

PROOF:
- `go test ./internal/modules/productcatalog/transport/http/... ./internal/presentation/http/id/productcatalog/...` passed.
- `go test ./internal/app/bootstrap/... ./internal/modules/productcatalog/transport/http/... ./internal/presentation/http/id/productcatalog/...` passed.
- `go test ./internal/transport/http/middleware/... -run TestProtectedRoutesRejectDisabledCapabilityBeforeHandler` passed.
- `bash scripts/audit_route_capabilities.sh` passed with 21 checked rows.
- `bash scripts/db_migrate.sh` applied `0013_seed_product_catalog_permissions_capabilities.up.sql`.
- `make verify` passed.

DECISION:
ProductCatalog runtime/capability slice is locally closed.

GAP:
Inventory stock mutation, stock adjustment create/reverse, ProductCatalog audit/outbox persistence, runtime smoke proof, and extended Laravel table filters remain out of scope.

PROGRESS:
- ProductCatalog persistence: 100%.
- ProductCatalog runtime/capability: 100% locally closed.
- Estimated ProductCatalog full transition: 78%.
- Estimated Business Phase 1: 56%.
- Estimated overall transition: 38%.

NEXT:
Select the next accepted slice before implementation.

## 2026-06-13 ProductCatalog API docs/error envelope closeout

FACT:

ProductCatalog developer-facing API documentation is locally implemented.

```text
docs/api/product_catalog.md
```

ProductCatalog API docs/error envelope blueprint is locally implemented.

```text
docs/blueprints/0032_api_docs_error_envelope_slice.md
```

Shared HTTP error envelope primitives and Echo HTTP error handler are locally implemented.

```text
internal/transport/http/response
```

Bootstrap wires the shared HTTP error handler.

ProductCatalog mapped errors now expose stable public error codes:

```text
product_not_found
product_code_already_exists
product_identity_already_exists
product_validation_failed
product_catalog_request_failed
invalid_request_body
invalid_query_parameter
```

Protected-route disabled capability proof now checks standard error envelope shape with:

```text
capability_disabled
```

ProductCatalog HTTP-level not-found proof now checks standard error envelope shape with:

```text
product_not_found
```

PROOF:

```text
go test ./internal/transport/http/response
PASS

go test ./internal/modules/productcatalog/transport/http/... ./internal/transport/http/response
PASS

go test ./internal/app/bootstrap/... ./internal/transport/http/response
PASS

go test ./internal/transport/http/middleware/... -run TestProtectedRoutesRejectDisabledCapabilityBeforeHandler
PASS

go test ./internal/app/bootstrap/... ./internal/modules/productcatalog/transport/http/... ./internal/presentation/http/id/productcatalog/... ./internal/transport/http/middleware/... ./internal/transport/http/response/...
PASS

bash scripts/audit_route_capabilities.sh
checked route capability rows: 21
[PASS] route capability audit passed

make verify
[PASS] aggregate audit passed
```

GAP:

Inventory stock mutation is not implemented.

Stock adjustment create/reverse is not implemented.

Broad audit sink is not implemented.

Runtime language switch/localization is not implemented.

Extended Laravel table filters remain unexposed:

```text
sort_by
sort_dir
merek
ukuran_min
ukuran_max
harga_min
harga_max
stok_saat_ini
```

Shared success envelope centralization is not implemented.

End-to-end runtime smoke proof with real HTTP server, auth token, and DB-backed ProductCatalog route is not proven in this slice.

ADR 0012 API output centralization remains partial because not every API surface has full response/error envelope coverage yet.

DECISION:

ProductCatalog backend API/runtime/capability/control-hex scope is locally closed.

ProductCatalog API docs and standardized error envelope slice is locally closed.

ProductCatalog full business transition remains open.

PROGRESS:

ProductCatalog implementation slice 1: 100%.

ProductCatalog PostgreSQL persistence: 100%.

ProductCatalog runtime/capability: 100% locally closed.

ProductCatalog API docs/error envelope: 100% locally closed.

Estimated ProductCatalog full transition: 82%.

Estimated Business Phase 1: 58%.

Estimated overall Laravel-to-Go transition: 40%.

NEXT:

Recommended next ProductCatalog slice:

```text
ProductCatalog runtime smoke proof with local auth token and DB-backed HTTP route.
```

Do not re-open ProductCatalog persistence, runtime/capability, API docs, or error envelope work unless a bug is found.

## 2026-06-14 ProductCatalog runtime smoke proof closeout

FACT:

ProductCatalog catalog API is locally proven through the real runtime path:

```text
local PostgreSQL -> migrations -> Echo server -> manual auth token -> protected ProductCatalog HTTP routes -> DB-backed responses
```

The smoke used blueprint:

```text
docs/blueprints/0033_productcatalog_runtime_smoke_proof_slice.md
```

Sanitized evidence is recorded at:

```text
docs/evidence/2026-06-14_productcatalog_runtime_smoke_proof.md
```

PROOF:

```text
make db-migrate
[PASS] db migrate completed

make db-status
[APPLIED] migrations through 0013_seed_product_catalog_permissions_capabilities.up.sql

GOCACHE=/tmp/gopos-go-build AUTH_DEBUG_ENABLED=true HTTP_PORT=18081 make run
server started against local PostgreSQL after unsandboxed DB access

health_status=200
health_status_field=ok

manual_login status=200 access_token_present=True refresh_token_present=True
unauth_products status=401 success=False error_code=authentication_required has_meta=True
list_products status=200 success=True has_meta=True
lookup_products status=200 success=True has_meta=True
create_product status=201 success=True id_present=True kode_barang_present=True has_meta=True
show_product status=200 success=True id_match=True kode_barang_present=True status_field=active has_meta=True
db_product_row table=products count=1
disable_capability key=product_catalog.list status=200 enabled=False
list_products_capability_disabled status=403 success=False error_code=capability_disabled has_meta=True
enable_capability key=product_catalog.list status=200 enabled=True
list_products_after_reenable status=200 success=True
[PASS] product catalog runtime smoke proof passed
```

The server process on port `18081` was manually stopped after the interrupted run, and the port was verified closed.

No bearer token, refresh token, raw auth payload, database credential, or secret is recorded.

GAP:

Product inventory/stock API is not implemented.

ProductCatalog audit/outbox persistence is not implemented.

Shared success envelope centralization is not implemented.

ADR `0012` API output centralization remains partial across all API surfaces.

ADR `0009` full debug auth lane remains partial; this smoke only proves manual login token issuance as part of the ProductCatalog runtime path.

Runtime language switch/localization is not implemented.

Extended Laravel table filters remain unexposed.

DECISION:

ProductCatalog runtime smoke proof is locally proven.

ProductCatalog catalog API is closed locally for pure API scope.

Product full backend transition remains partial.

PROGRESS:

ProductCatalog implementation slice 1: 100%.

ProductCatalog PostgreSQL persistence: 100%.

ProductCatalog runtime/capability: 100% locally closed.

ProductCatalog API docs/error envelope: 100% locally closed.

ProductCatalog runtime smoke proof: 100% locally proven.

Estimated ProductCatalog full transition: 84%.

Estimated Business Phase 1: 60%.

Estimated overall Laravel-to-Go transition: 41%.

NEXT:

Execution channel: Terminal Codex.

Recommended next slice:

```text
Shared success envelope and ADR 0012 output contract centralization.
```

Do not start inventory mutation, audit/outbox implementation, localization, extended filters, or router/server/bootstrap cleanup before the output-contract slice unless the owner changes scope.
