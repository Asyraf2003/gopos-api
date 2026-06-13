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

# ProductCatalog Domain Slice 1 Handoff

## Date

2026-06-09

## Active Scope

ProductCatalog implementation slice 1.

Current active slice:

```text
domain, ports, usecase contract, and unit tests only
```

Current completed local work in this slice:

```text
ProductCatalog domain tests and minimal domain implementation
```

## Current Branch Or Source Snapshot

Owner/local terminal snapshot.

GitHub connector validation for ProductCatalog domain package is still pending.

## Files Included

```text
docs/blueprints/0028_productcatalog_domain_contract.md
docs/blueprints/0029_productcatalog_implementation_slice_1.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
internal/modules/productcatalog/domain/
```

## Files Changed

ProductCatalog domain package was implemented locally.

Expected local files include:

```text
internal/modules/productcatalog/domain/product_errors.go
internal/modules/productcatalog/domain/product_types.go
internal/modules/productcatalog/domain/product_constructor.go
internal/modules/productcatalog/domain/product_accessors.go
internal/modules/productcatalog/domain/product_lifecycle.go
internal/modules/productcatalog/domain/product_normalization.go
internal/modules/productcatalog/domain/product_constructor_test.go
internal/modules/productcatalog/domain/product_validation_test.go
internal/modules/productcatalog/domain/product_lifecycle_test.go
```

Ledger updated locally:

```text
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

Blueprints accepted locally:

```text
docs/blueprints/0028_productcatalog_domain_contract.md
docs/blueprints/0029_productcatalog_implementation_slice_1.md
```

## Files Forbidden To Touch In The Next Step

```text
internal/platform/postgres/
migrations/
cmd/server route wiring
internal/modules/productcatalog/transport/http/
internal/presentation/http/id/productcatalog/
capability seed migrations
inventory stock adjustment code
inventory stock reversal code
```

## Blueprint Referenced

```text
docs/blueprints/0028_productcatalog_domain_contract.md
docs/blueprints/0029_productcatalog_implementation_slice_1.md
```

## Decisions Made

ProductCatalog duplicate policy Option A is accepted:

```text
Preserve Laravel-compatible duplicate behavior.
```

Accepted duplicate behavior:

- Active non-null `kode_barang` must be unique.
- Blank `kode_barang` becomes null.
- Same normalized `nama_barang + merek + ukuran` is rejected when code exception does not apply.
- Same normalized identity is allowed only when both existing and candidate products have distinct non-null `kode_barang`.
- PostgreSQL must not use a strict active unique index on `(nama_barang_normalized, merek_normalized, ukuran)`.
- Nuanced duplicate behavior belongs in ProductCatalog usecase/repository guard inside a transaction.

## FACT

- ProductCatalog domain contract blueprint 0028 is accepted locally.
- ProductCatalog implementation slice 1 blueprint 0029 is accepted locally.
- ProductCatalog domain package is locally implemented with focused test proof.
- Aggregate `make verify` passed after domain implementation and file split.
- File size audit passed after splitting oversized ProductCatalog domain files.
- GitHub connector validation for ProductCatalog domain package is still pending.
- ProductCatalog PostgreSQL persistence has not started.
- ProductCatalog Echo/API/capability seed work has not started.
- Inventory stock adjustment remains deferred to future Inventory scope.

## PROOF

Focused proof:

```text
go test ./internal/modules/productcatalog/domain
ok  	pos-go/internal/modules/productcatalog/domain
```

Aggregate proof:

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

Working tree proof provided by owner:

```text
git status --short

Output was empty.
```

## GAP

- Connector validation for `internal/modules/productcatalog/domain` is pending.
- ProductCatalog ports are not implemented yet.
- ProductCatalog usecase contracts are not implemented yet.
- ProductCatalog duplicate-policy usecase tests are not implemented yet.
- ProductCatalog PostgreSQL persistence is not implemented.
- ProductCatalog HTTP/runtime/capability slice is not implemented.
- Inventory stock adjustment is not implemented and must remain deferred.

## PROGRESS

Business Phase 1 remains partial.

ProductCatalog progress is now:

```text
domain package locally implemented with proof; connector validation pending
```

Do not claim ProductCatalog slice 1 closed yet.

## CONTEXT WINDOW STATUS

Enough context remains to continue ProductCatalog slice 1 into ports contract planning.

## NEXT

Execution channel:

```text
owner/local terminal
```

Next valid active step:

Validate ProductCatalog domain package connector visibility, then continue ProductCatalog implementation slice 1 with ports contract planning.

Do not start PostgreSQL, Echo routes, capability seed migration, or inventory stock adjustment.
