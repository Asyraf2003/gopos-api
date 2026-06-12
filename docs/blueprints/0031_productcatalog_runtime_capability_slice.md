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

# ProductCatalog Runtime Capability Slice Blueprint

## Status

Accepted implementation slice plan.

## Date

2026-06-12

## Active Scope

Plan ProductCatalog HTTP runtime, route registration, response presenters, permission seeds, capability seeds, route capability manifest coverage, and disabled-capability proof after ProductCatalog PostgreSQL persistence closeout.

## Source Contracts

```text
docs/blueprints/0028_productcatalog_domain_contract.md
docs/blueprints/0029_productcatalog_implementation_slice_1.md
docs/blueprints/0030_productcatalog_postgres_persistence_slice.md
docs/handoffs/2026-06-12-productcatalog-postgres-persistence-closeout.md
```

## FACT

ProductCatalog domain contract is accepted.

ProductCatalog domain, ports, and usecases are implemented with proof.

ProductCatalog PostgreSQL migration and repository adapter are implemented with proof.

ProductCatalog PostgreSQL persistence closeout has query-plan proof.

ProductCatalog HTTP transport is not implemented.

ProductCatalog request/response presenters are not implemented.

ProductCatalog route registration is not implemented.

ProductCatalog permission seeds are not implemented.

ProductCatalog capability seeds are not implemented.

Route capability manifest does not yet include ProductCatalog protected routes.

## DECISION

This slice implements protected ProductCatalog runtime HTTP integration.

This slice must include route capability enforcement and proof.

Do not implement inventory mutation.

Do not implement UI.

Do not implement broad audit sink behavior beyond capability metadata and proof.

Do not change ProductCatalog persistence behavior unless required to wire existing usecases.

## Laravel Parity Stance

Laravel is behavioral evidence, not a file-by-file implementation template.

The Go target keeps the existing separation:

```text
domain
ports
usecase
transport/http
presentation/http/id
platform/postgres
app/bootstrap
```

This slice must not copy Laravel page controllers, Blade view behavior, or UI table endpoints.

ProductCatalog runtime routes are Go API contracts.

Laravel runtime evidence confirms:

```text
ProductCatalog create/update web actions use:
POST /product-catalog/products/create
POST /product-catalog/products/{productId}/update

Admin lifecycle actions use:
DELETE /admin/products/{productId}
PATCH  /admin/products/{productId}/restore
```

The Go API route shape remains API-only, while request and response field names preserve the Laravel/ProductCatalog public vocabulary where practical.

Laravel runtime evidence confirms:

```text
ProductCatalog create/update web actions use:
POST /product-catalog/products/create
POST /product-catalog/products/{productId}/update

Admin lifecycle actions use:
DELETE /admin/products/{productId}
PATCH  /admin/products/{productId}/restore
```

The Go API route shape remains API-only, while request and response field names preserve the Laravel/ProductCatalog public vocabulary where practical.

## SCOPE-IN

ProductCatalog HTTP transport.

Request parsing for list, lookup, show, create, update, soft delete, restore, and version timeline.

Response presenter for Indonesian/public HTTP JSON contracts.

Bootstrap route registration.

Permission seed migration for:

```text
product_catalog.read
product_catalog.manage
```

Role permission assignment:

```text
admin gets product_catalog.read
admin gets product_catalog.manage
cashier gets product_catalog.read
base gets no ProductCatalog permission in this slice.
```

Capability seed migration for all protected ProductCatalog routes.

Route capability manifest update.

Route capability audit script update if the audit source list requires explicit ProductCatalog files.

HTTP handler tests for request parsing, validation, envelope shape, not-found mapping, duplicate mapping, and handler/usecase boundary.

Runtime disabled-capability proof through route-level test or bootstrap test.

## SCOPE-OUT

Inventory stock mutation.

Inventory stock adjustment create/reverse.

UI.

Public unauthenticated ProductCatalog routes.

Physical delete.

Audit event sink implementation.

Idempotency-key implementation.

Pagination total-count metadata unless explicitly required by existing API contract.

Broad ADR 0012 closeout beyond proof for this slice.

New ProductCatalog persistence behavior.

## ROUTE CONTRACT

Base path:

```text
/api/products
```

Routes:

```text
GET    /api/products
POST   /api/products
GET    /api/products/lookup
GET    /api/products/:id
PUT    /api/products/:id
DELETE /api/products/:id
PATCH  /api/products/:id/restore
GET    /api/products/:id/versions
```

Route ordering rule:

Register `/lookup` before `/:id`.

Register `/:id/restore` and `/:id/versions` before generic `/:id` if the router requires explicit ordering for nested paths.

Restore uses `PATCH` to match the accepted Laravel ProductCatalog lifecycle evidence.

## ROUTE AND CAPABILITY MAPPING

```text
GET    /api/products              product_catalog.list     product_catalog.read
POST   /api/products              product_catalog.create   product_catalog.manage
GET    /api/products/lookup       product_catalog.lookup   product_catalog.read
GET    /api/products/:id          product_catalog.show     product_catalog.read
PUT    /api/products/:id          product_catalog.update   product_catalog.manage
DELETE /api/products/:id          product_catalog.delete   product_catalog.manage
PATCH  /api/products/:id/restore  product_catalog.restore  product_catalog.manage
GET    /api/products/:id/versions product_catalog.versions product_catalog.read
```

## AUTHORIZATION POLICY

Initial permission keys:

```text
product_catalog.read
product_catalog.manage
```

Initial role assignment:

```text
base    -> no ProductCatalog permission
cashier -> product_catalog.read
admin   -> product_catalog.read, product_catalog.manage
```

Reason:

Cashier may need read/lookup access for future sales flow.

Admin owns master-data management.

Base role should remain minimal.

## CAPABILITY SEED POLICY

Seed all ProductCatalog route capabilities into api_capabilities.

Default enabled:

```text
true
```

Risk levels:

```text
product_catalog.list     low
product_catalog.lookup   low
product_catalog.show     low
product_catalog.versions low
product_catalog.create   medium
product_catalog.update   medium
product_catalog.delete   medium
product_catalog.restore  medium
```

Audit required:

```text
read routes: false
write/lifecycle routes: true
```

Idempotency required:

```text
false
```

Owner package:

```text
internal/modules/productcatalog/transport/http
```

## REQUEST CONTRACT

List query:

```text
q
page
per_page
status = active|deleted|all

Laravel inventory also records richer list filters and sorting:

sort_by = nama_barang|merek|ukuran|harga_jual|stok_saat_ini
sort_dir = asc|desc
merek
ukuran_min
ukuran_max
harga_min
harga_max

This runtime slice only exposes the currently proven Go ProductReader subset unless a later persistence/query slice adds the richer filters.

Do not expose stok_saat_ini sorting in this slice because inventory projection belongs to inventory behavior.
```

Lookup query:

```text
q
limit
include_deleted
```

Create JSON body:

```json
{
  "kode_barang": "SKU-001",
  "nama_barang": "Kampas Rem",
  "merek": "Honda",
  "ukuran": 14,
  "harga_jual": 40000,
  "reorder_point_qty": 5,
  "critical_threshold_qty": 2,
  "reason": "optional reason"
}
```

Update JSON body:

```json
{
  "kode_barang": "SKU-001",
  "nama_barang": "Kampas Rem",
  "merek": "Honda",
  "ukuran": 14,
  "harga_jual": 40000,
  "reorder_point_qty": 5,
  "critical_threshold_qty": 2,
  "reason": "optional reason"
}
```

Soft delete JSON body:

```json
{
  "reason": "optional reason"
}
```

Laravel admin soft delete does not submit a reason; it passes actor identity only. Reason remains optional in Go runtime because the Go usecase command already supports it.

Restore JSON body:

```json
{
  "reason": "optional reason"
}
```

Laravel admin restore does not submit a reason; it passes actor identity only. Reason remains optional in Go runtime because the Go usecase command already supports it.

## RESPONSE CONTRACT

Success response envelope:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

List response data item:

```json
{
  "id": "product-id",
  "kode_barang": "SKU-001",
  "nama_barang": "Kampas Rem",
  "merek": "Honda",
  "ukuran": 14,
  "harga_jual": 40000,
  "status": "active"
}
```

Detail response data:

```json
{
  "id": "product-id",
  "kode_barang": "SKU-001",
  "nama_barang": "Kampas Rem",
  "nama_barang_normalized": "kampas rem",
  "merek": "Honda",
  "merek_normalized": "honda",
  "ukuran": 14,
  "harga_jual": 40000,
  "reorder_point_qty": 5,
  "critical_threshold_qty": 2,
  "status": "active"
}
```

Write response data should use the public ProductCatalog field names above and include revision number where available.

Lookup response data item:

```json
{
  "id": "product-id",
  "kode_barang": "SKU-001",
  "nama_barang": "Kampas Rem",
  "merek": "Honda",
  "ukuran": 14,
  "harga_jual": 40000,
  "status": "active"
}
```

Version timeline response data item:

```json
{
  "product_id": "product-id",
  "revision_no": 1,
  "event_name": "product.created",
  "changed_by_actor_id": "actor-id",
  "change_reason": "created from API",
  "changed_at": "RFC3339"
}
```

Error responses must use the existing public error envelope behavior available in the HTTP stack.

## TARGET FILES

Expected new or changed files:

```text
internal/modules/productcatalog/transport/http/product_catalog_handler.go
internal/modules/productcatalog/transport/http/product_catalog_handler_read.go
internal/modules/productcatalog/transport/http/product_catalog_handler_write.go
internal/modules/productcatalog/transport/http/product_catalog_handler_request.go
internal/modules/productcatalog/transport/http/product_catalog_handler_response.go
internal/modules/productcatalog/transport/http/product_catalog_handler_error.go
internal/modules/productcatalog/transport/http/product_catalog_handler_test.go
internal/presentation/http/id/productcatalog/product.go
internal/presentation/http/id/productcatalog/product_lookup.go
internal/presentation/http/id/productcatalog/product_version.go
internal/app/bootstrap/app.go
migrations/0013_seed_product_catalog_permissions_capabilities.up.sql
migrations/0013_seed_product_catalog_permissions_capabilities.down.sql
scripts/config/route_capabilities.tsv
scripts/audit_route_capabilities.sh
```

If file-size audit requires smaller files, split handler tests/helpers without changing package boundaries.

## BOOTSTRAP WIRING PLAN

Bootstrap should wire:

```text
postgres.NewProductRepository(pool)
ProductCatalog usecases
ProductCatalog HTTP handler
```

Protected groups should preserve order:

```text
authn -> authz permission -> runtime capability check -> handler
```

Because different ProductCatalog routes use different permissions/capabilities, bootstrap may use separate groups or per-route registration helpers to keep exact capability mapping obvious.

## ROUTE CAPABILITY MANIFEST PLAN

Add rows to `scripts/config/route_capabilities.tsv`:

```text
GET	/api/products	product_catalog.list	product_catalog.read	exact	group.GET("/products"
POST	/api/products	product_catalog.create	product_catalog.manage	exact	group.POST("/products"
GET	/api/products/lookup	product_catalog.lookup	product_catalog.read	exact	group.GET("/products/lookup"
GET	/api/products/:id	product_catalog.show	product_catalog.read	exact	group.GET("/products/:id"
PUT	/api/products/:id	product_catalog.update	product_catalog.manage	exact	group.PUT("/products/:id"
DELETE	/api/products/:id	product_catalog.delete	product_catalog.manage	exact	group.DELETE("/products/:id"
PATCH	/api/products/:id/restore	product_catalog.restore	product_catalog.manage	exact	group.PATCH("/products/:id/restore"
GET	/api/products/:id/versions	product_catalog.versions	product_catalog.read	exact	group.GET("/products/:id/versions"
```

Update `scripts/audit_route_capabilities.sh` if the audit script requires explicit seed/source file lists.

## TEST MATRIX

Handler tests:

```text
List parses query and calls usecase.
Lookup parses query and calls usecase.
Show requires ID and calls usecase.
Create rejects invalid body.
Create maps command fields correctly.
Update requires ID and valid body.
Soft delete requires ID and maps reason.
Restore requires ID and maps reason.
Versions requires product ID and calls usecase.
Not-found usecase error maps to 404.
Duplicate code/identity usecase errors map to conflict.
Success responses use envelope.
Handler package must not import PostgreSQL adapter.
```

Bootstrap/route tests:

```text
ProductCatalog routes are registered only behind authn/authz/capability checks.
Disabled ProductCatalog capability returns 403 before handler/usecase execution.
Route capability manifest audit passes with ProductCatalog rows.
```

Migration/seed tests or proof:

```text
product_catalog.read permission exists.
product_catalog.manage permission exists.
cashier has product_catalog.read.
admin has product_catalog.read.
admin has product_catalog.manage.
api_capabilities contains all eight ProductCatalog capability keys.
```

## PROOF REQUIRED

Focused HTTP proof:

```bash
go test ./internal/modules/productcatalog/transport/http/...
```

Focused bootstrap/capability proof:

```bash
go test ./internal/app/bootstrap/... ./internal/transport/http/middleware/...
```

Route capability proof:

```bash
bash scripts/audit_route_capabilities.sh
```

Full proof:

```bash
make verify
```

Runtime DB proof if local DB is available:

```bash
set -a
. ./.env
set +a
bash scripts/db_migrate.sh
```

Expected applied migration:

```text
0013_seed_product_catalog_permissions_capabilities.up.sql
```

Optional SQL proof:

```sql
select key from permissions where key in ('product_catalog.read', 'product_catalog.manage');

select capability_key from api_capabilities
where capability_key like 'product_catalog.%'
order by capability_key;
```

## PROGRESS UPDATE REQUIREMENT

After implementation proof passes:

Update `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`.

Create or update a ProductCatalog runtime/capability handoff.

Record focused HTTP proof, bootstrap/capability proof, route capability proof, DB migration proof if available, and aggregate make verify proof.

## NEXT

After this blueprint is accepted, implement ProductCatalog HTTP transport and presenter first.

Do not start route registration or seed migration until handler/presenter proof is stable.
