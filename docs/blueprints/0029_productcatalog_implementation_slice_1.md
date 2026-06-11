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

# ProductCatalog Implementation Slice 1 Blueprint

## Status

Accepted.

## Date

2026-06-09

## Active Scope

ProductCatalog implementation slice 1 planning only.

This slice covers:

```text
internal/modules/productcatalog/domain
internal/modules/productcatalog/ports
internal/modules/productcatalog/usecase
```

This slice does not implement PostgreSQL persistence, Echo routes, capability seed migration, or inventory stock adjustment.

## FACT

- ProductCatalog domain contract is accepted in `docs/blueprints/0028_productcatalog_domain_contract.md`.
- Duplicate policy Option A is accepted.
- Option A preserves Laravel-compatible duplicate behavior.
- ProductCatalog contract is connector-visible.
- Current ledger records ProductCatalog contract accepted with local proof and connector-visible docs.
- ProductCatalog Go business module implementation proof does not exist yet.
- Product stock adjustment belongs to future Inventory scope.

## GAP

- No `internal/modules/productcatalog` implementation exists yet.
- No ProductCatalog domain tests exist yet.
- No ProductCatalog usecase tests exist yet.
- No ProductCatalog PostgreSQL adapter exists yet.
- No ProductCatalog HTTP handler/routes exist yet.
- No ProductCatalog capability seed migration exists yet.
- No ProductCatalog runtime proof exists yet.

## DECISION

Slice 1 must implement only ProductCatalog domain, ports, and usecase contract.

Slice 1 must not create:

```text
database migrations
PostgreSQL adapters
Echo handlers
route registration
capability seed migration
inventory stock adjustment
inventory stock reversal
```

## ACCEPTED DUPLICATE POLICY

Option A:

```text
Preserve Laravel-compatible duplicate behavior.
```

Rules:

- Active non-null `kode_barang` must be unique.
- Blank `kode_barang` becomes null.
- Same normalized `nama_barang + merek + ukuran` is rejected when code exception does not apply.
- Same normalized identity is allowed only when both existing and candidate products have distinct non-null `kode_barang`.
- PostgreSQL must not use a strict active unique index on `(nama_barang_normalized, merek_normalized, ukuran)`.
- Usecase/repository guard must enforce nuanced identity duplicate behavior transactionally.

## SLICE 1 PACKAGE PLAN

Allowed package paths:

```text
internal/modules/productcatalog/domain
internal/modules/productcatalog/ports
internal/modules/productcatalog/usecase
```

Forbidden package paths in this slice:

```text
internal/platform/postgres
internal/transport/http
internal/modules/productcatalog/transport/http
internal/presentation/http/id/productcatalog
migrations
seed migrations
cmd/server route wiring
```

## DOMAIN OBJECTS

Candidate domain types:

```text
Product
ProductID
ProductCode
ProductName
ProductBrand
ProductSize
MoneyRupiah
ThresholdPolicy
ProductStatus
ProductSnapshot
ProductVersion
DuplicatePolicy
```

## DOMAIN INVARIANTS

- Product ID is required.
- Product name is required after trim.
- Product brand is required after trim.
- Product code is trimmed.
- Blank product code becomes null.
- Product name normalized value is server-generated.
- Product brand normalized value is server-generated.
- Sale price must be greater than zero.
- Reorder point and critical threshold must both be null or both filled.
- Threshold values must be non-negative.
- Critical threshold must not exceed reorder point.
- Deleted product has `deleted_at`.
- Active product has no `deleted_at`.
- Restore clears deleted metadata.
- Physical delete is forbidden.

## USECASE CONTRACTS

Slice 1 usecase commands:

```text
CreateProduct
UpdateProduct
SoftDeleteProduct
RestoreProduct
GetProductDetail
ListProducts
LookupProducts
ListProductVersions
```

Slice 1 may define contracts for all operations, but tests should focus first on command behavior and duplicate policy.

## PORTS

Repository/read ports:

```text
ProductRepository
ProductReader
ProductVersionRepository
ProductDuplicateChecker
```

Transaction/audit/time/id ports:

```text
TransactionRunner
AuditRecorder
Clock
IDGenerator
```

Inventory projection port:

```text
ProductInventoryProjectionReader
```

Important:

`ProductInventoryProjectionReader` is read-only and may be defined only as a future dependency for list/detail/lookup. It must not expose stock mutation.

## FORBIDDEN USECASE BEHAVIOR

ProductCatalog usecases must not:

- adjust stock;
- reverse stock adjustment;
- write inventory movement;
- write costing records;
- import Echo;
- import SQL driver packages;
- know PostgreSQL details;
- own capability middleware behavior.

Capability and authz are enforced by HTTP/runtime layer in later slices.

## TEST MATRIX

Domain tests:

- trims product name and brand.
- normalizes product name and brand.
- blank product code becomes null.
- rejects blank product name.
- rejects blank product brand.
- rejects sale price less than or equal to zero.
- rejects one-sided threshold.
- rejects negative threshold.
- rejects critical threshold greater than reorder point.
- soft delete changes lifecycle state.
- restore clears deleted metadata.
- physical delete is not modeled as an allowed operation.

Usecase tests:

- create product success.
- create product records version.
- create product records audit intent/event.
- create rejects duplicate active code.
- create rejects duplicate identity when code exception does not apply.
- create allows same identity when both products have distinct non-null code.
- update product success.
- update rejects duplicate active code.
- update rejects duplicate identity when code exception does not apply.
- soft delete active product success.
- soft delete missing product returns not found.
- soft delete already deleted product is rejected.
- restore deleted product success.
- restore active product is rejected.
- restore missing product returns not found.
- list product uses query contract object.
- lookup product is bounded by limit contract.
- versions list uses product id.

Architecture tests:

- no Echo import in domain/usecase.
- no SQL import in domain/usecase.
- no inventory mutation import in ProductCatalog.

## PROOF REQUIRED FOR SLICE 1 IMPLEMENTATION

Implementation is not accepted until these pass locally:

```bash
go test ./internal/modules/productcatalog/...
make verify
```

Expected aggregate proof:

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

## IMPLEMENTATION ORDER

1. Create ProductCatalog domain package.
2. Add normalizer/value object tests.
3. Add product invariant tests.
4. Add lifecycle tests.
5. Define ports.
6. Add usecase command/query contracts.
7. Add fake repository/test doubles.
8. Add create/update duplicate-policy tests.
9. Add soft-delete/restore tests.
10. Run focused tests.
11. Run `make verify`.

## OUT OF SCOPE

- PostgreSQL schema.
- PostgreSQL repository.
- Database migrations.
- Product version table migration.
- Audit table migration.
- Echo handlers.
- Route registration.
- Capability seed migration.
- Runtime capability-disabled proof.
- Inventory stock adjustment.
- Inventory stock reversal.

## ACCEPTANCE

Accepted on: 2026-06-09

Accepted scope:

```text
ProductCatalog implementation slice 1:
domain, ports, usecase contract, and unit tests only.
```

Accepted constraints:

```text
Do not implement PostgreSQL persistence in this slice.
Do not create database migrations in this slice.
Do not implement Echo routes in this slice.
Do not create capability seed migrations in this slice.
Do not implement inventory stock adjustment in this slice.
Preserve duplicate policy Option A from docs/blueprints/0028_productcatalog_domain_contract.md.
```

## NEXT ACTIVE STEP

ProductCatalog slice 1 blueprint is accepted.

First implementation step after acceptance:

Create ProductCatalog domain tests for normalization, product code nulling, price, threshold, and lifecycle invariants.
