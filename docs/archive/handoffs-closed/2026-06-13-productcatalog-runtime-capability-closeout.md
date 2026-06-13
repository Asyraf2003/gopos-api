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
-->

# ProductCatalog Runtime Capability Closeout

## Date

2026-06-13

## Active Scope

ProductCatalog protected HTTP runtime, presenter, route registration, permission seed, capability seed, route capability manifest, and disabled-capability proof.

## FACT

ProductCatalog PostgreSQL persistence was already closed before this slice.

ProductCatalog runtime handler and presenter are implemented.

ProductCatalog API routes are registered under `/api/products`.

ProductCatalog public JSON field names preserve Laravel/ProductCatalog vocabulary where practical:

```text
kode_barang
nama_barang
merek
ukuran
harga_jual
reorder_point_qty
critical_threshold_qty
```

ProductCatalog write handlers map public request fields to Go usecase commands.

ProductCatalog write handlers use authenticated principal account ID as actor ID when available.

ProductCatalog create, update, soft delete, and restore use a no-op ProductAuditRecorder in bootstrap because broad audit sink implementation is outside this slice.

ProductCatalog permission seed migration is implemented:

```text
product_catalog.read
product_catalog.manage
```

Initial role assignment is implemented:

```text
cashier -> product_catalog.read
admin   -> product_catalog.read
admin   -> product_catalog.manage
base    -> no ProductCatalog permission
```

ProductCatalog capability seed migration is implemented for:

```text
product_catalog.list
product_catalog.create
product_catalog.lookup
product_catalog.show
product_catalog.update
product_catalog.delete
product_catalog.restore
product_catalog.versions
```

Route capability manifest includes all protected ProductCatalog routes.

Disabled-capability middleware proof includes all protected ProductCatalog routes.

## PROOF

Focused ProductCatalog HTTP proof:

```text
go test ./internal/modules/productcatalog/transport/http/... ./internal/presentation/http/id/productcatalog/...
ok      pos-go/internal/modules/productcatalog/transport/http    (cached)
?       pos-go/internal/presentation/http/id/productcatalog      [no test files]
```

Focused bootstrap/runtime proof:

```text
go test ./internal/app/bootstrap ./internal/modules/productcatalog/transport/http/... ./internal/presentation/http/id/productcatalog/...
ok      pos-go/internal/app/bootstrap                            0.187s
ok      pos-go/internal/modules/productcatalog/transport/http    (cached)
?       pos-go/internal/presentation/http/id/productcatalog      [no test files]
```

Disabled-capability proof:

```text
go test ./internal/transport/http/middleware/... -run TestProtectedRoutesRejectDisabledCapabilityBeforeHandler
ok      pos-go/internal/transport/http/middleware                0.007s
```

Route capability proof:

```text
bash scripts/audit_route_capabilities.sh

checked route capability rows: 21
[PASS] route capability audit passed
```

Migration proof:

```text
bash scripts/db_migrate.sh

[APPLY] 0013_seed_product_catalog_permissions_capabilities.up.sql
BEGIN
INSERT 0 2
INSERT 0 2
INSERT 0 1
INSERT 0 8
INSERT 0 1
COMMIT

[PASS] db migrate completed
```

Aggregate proof:

```text
make verify

[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] license header audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed
```

## GAP

ProductCatalog UI is not implemented.

Inventory stock mutation is not implemented.

Inventory stock adjustment create/reverse is not implemented.

Broad audit sink implementation is not implemented.

Runtime ProductCatalog uses no-op ProductAuditRecorder until a future audit/outbox slice exists.

Extended Laravel table filters are not implemented in this runtime slice:

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

Only the currently proven ProductReader subset is exposed:

```text
q
page
per_page
status
```

## DECISION

ProductCatalog runtime/capability slice is closed locally with focused handler proof, bootstrap proof, disabled-capability proof, route capability audit proof, migration proof, and aggregate proof.

Do not start UI or inventory mutation from this closeout.

Next valid ProductCatalog work should be selected as a new accepted slice.

## PROGRESS

ProductCatalog PostgreSQL persistence slice: 100%.

ProductCatalog runtime/capability slice: 100% locally closed.

Estimated ProductCatalog full transition: 78%.

Estimated Business Phase 1: 56%.

Estimated overall Laravel-to-Go transition: 38%.

## NEXT

Pick the next accepted slice.

Recommended next candidates:

```text
1. ProductCatalog runtime smoke/integration proof if local auth token workflow is ready.
2. Inventory stock projection/mutation blueprint.
3. Employee/master-data domain contract if ProductCatalog should pause.
4. Audit/outbox implementation slice to replace ProductCatalog no-op audit recorder.
```

## CONTEXT WINDOW STATUS

The next session should treat ProductCatalog persistence and runtime/capability as locally proven.

Do not re-open ProductCatalog persistence unless a bug is found.

Do not implement UI or inventory behavior without a new accepted blueprint.
