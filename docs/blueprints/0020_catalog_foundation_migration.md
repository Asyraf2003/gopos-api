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

# Catalog Foundation Migration Blueprint

## FACT
- Master transition blueprint is `docs/blueprints/0012_laravel_to_go_api_transition_master_plan.md`.
- Stage 0 route/schema evidence is `docs/evidence/0001_laravel_stage0_schema_and_route_inventory.md`.
- ProductCatalog/ServiceCatalog evidence is `docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md`.
- Current Go repo has no `servicecatalog`, `productcatalog`, or `inventory` business modules yet.
- Current Go repo has PostgreSQL/auth foundation, but POS domain migrations are not implemented.
- Product stock adjustment is exposed near product UI in Laravel, but behavior belongs to `inventory`.

## GAP
- Capability foundation from `0010_capability_control_foundation.md` is not implemented yet.
- Full `ServiceCatalog` Laravel source was not provided, only seed and normalizer test.
- Full inventory stock adjustment source was not provided.
- Product duplicate behavior has an important policy decision before PostgreSQL uniqueness is finalized.

## DECISION
- Implement catalog foundation before procurement, note, payment, and reporting.
- Keep `servicecatalog` and `productcatalog` as separate Go modules.
- Do not implement inventory stock adjustment in this blueprint; only expose product inventory read data where needed for product list/lookup if the dependency is already available.
- Use PostgreSQL-native partial unique indexes instead of MySQL generated `active_unique_marker`.
- Keep ProductCatalog write behavior transactional: product row + product version + audit event.
- API is Echo-only and JSON-only. Laravel page routes are not ported.
- User-facing Indonesian messages belong in presentation/output layer, not domain.

## SCOPE-IN
- `servicecatalog` domain contract, schema, seed, lookup/list API.
- `productcatalog` domain contract, schema, create/update/soft-delete/restore/show/list/lookup API.
- Product version timeline.
- Product audit event write decision for create/update/delete/restore.
- Product table/list query contract with pagination/filter/sort.
- Product duplicate and threshold validation.
- PostgreSQL migration blueprint and test plan.
- Capability key proposal.

## SCOPE-OUT
- Inventory stock adjustment create/reverse.
- Procurement invoice behavior.
- Note transaction workspace.
- Customer payment/refund.
- Reporting dashboards.
- Blade or page output.
- Mobile API specialization.

## DOMAIN CONTRACT - SERVICE CATALOG

Domain: `servicecatalog`

Source table:

- `service_catalog_items`

Aggregate root:

- `ServiceCatalogItem`

Allowed operations:

- create
- update default price/name
- activate
- deactivate
- show
- list/lookup

Delete policy:

- physical delete forbidden for normal use.
- deactivate instead.

Invariants:

- `name` required.
- `normalized_name` required and unique.
- `default_price_rupiah > 0`.
- `normalized_name` produced by normalizer, not trusted from clients.

Audit policy:

- create/update/activate/deactivate should emit audit event once audit module exists.

Capability keys:

```text
service_catalog.list
service_catalog.show
service_catalog.create
service_catalog.update
service_catalog.activate
service_catalog.deactivate
```

## DOMAIN CONTRACT - PRODUCT CATALOG

Domain: `productcatalog`

Source tables:

- `products`
- `product_versions`
- `product_inventory` for read-only stock projection join

Aggregate root:

- `Product`

Allowed operations:

- create
- update
- soft delete
- restore
- show
- list
- lookup
- version timeline

Forbidden operations:

- physical delete.
- stock adjustment mutation in ProductCatalog.

Delete policy:

- soft delete with `deleted_at`, `deleted_by_actor_id`, and optional `delete_reason`.
- restore clears deleted metadata.

Lifecycle statuses:

- active: `deleted_at is null`.
- deleted: `deleted_at is not null`.

Invariants:

- id required.
- `nama_barang` required after trim.
- `merek` required after trim.
- `harga_jual_rupiah > 0`.
- `kode_barang` blank normalizes to null.
- `nama_barang_normalized` and `merek_normalized` come from server normalizer.
- threshold fields must be both null or both non-null.
- threshold fields must be non-negative.
- `critical_threshold_qty <= reorder_point_qty`.

Duplicate policy pending decision:

- Laravel allows same name/brand/size when both existing and candidate have distinct non-null `kode_barang`.
- PostgreSQL strict identity unique index would reject that.
- Recommended Go decision: enforce active unique `kode_barang` in DB and enforce business duplicate behavior in usecase/query guard.

Audit policy:

- create/update/soft-delete/restore are auditable.
- audit metadata should include snapshot and revision number.

Transaction boundary:

- create/update/soft-delete/restore must write product row, product version, and audit event in one transaction.

Idempotency:

- not required for normal admin product create/update in first implementation.
- may be introduced for API clients later if duplicate-submit risk is material.

Capability keys:

```text
products.list
products.lookup
products.show
products.create
products.update
products.delete
products.restore
products.versions.list
```

## TARGET GO PACKAGE PLAN

```text
internal/modules/servicecatalog/
  domain/
  ports/
  usecase/
  transport/http/

internal/modules/productcatalog/
  domain/
  ports/
  usecase/
  transport/http/

internal/platform/postgres/
  service_catalog_*.go
  product_catalog_*.go

internal/presentation/http/id/servicecatalog/
internal/presentation/http/id/productcatalog/
```

No Echo import in domain/usecase.

No SQL in handlers/usecases.

No product module import of inventory mutation logic.

## POSTGRESQL SCHEMA PLAN

### `service_catalog_items`

```sql
id text primary key
name text not null
normalized_name text not null unique
default_price_rupiah bigint not null check (default_price_rupiah > 0)
is_active boolean not null default true
created_at timestamptz not null default now()
updated_at timestamptz not null default now()
```

### `products`

```sql
id text primary key
kode_barang text null
nama_barang text not null
nama_barang_normalized text not null
merek text not null
merek_normalized text not null
ukuran integer null
harga_jual_rupiah bigint not null check (harga_jual_rupiah > 0)
reorder_point_qty integer null check (reorder_point_qty >= 0)
critical_threshold_qty integer null check (critical_threshold_qty >= 0)
deleted_at timestamptz null
deleted_by_actor_id text null
delete_reason text null
created_at timestamptz not null default now()
updated_at timestamptz not null default now()
```

Checks:

```sql
check (
  (reorder_point_qty is null and critical_threshold_qty is null)
  or
  (reorder_point_qty is not null and critical_threshold_qty is not null)
)

check (
  reorder_point_qty is null
  or critical_threshold_qty <= reorder_point_qty
)
```

Indexes:

```sql
create index products_deleted_at_idx on products(deleted_at);
create index products_name_norm_idx on products(nama_barang_normalized);
create index products_brand_norm_idx on products(merek_normalized);
create unique index products_kode_barang_active_unique
  on products(kode_barang)
  where deleted_at is null and kode_barang is not null;
```

Business identity unique index is deferred until duplicate policy is accepted.

### `product_versions`

```sql
id text primary key
product_id text not null references products(id) on delete restrict
revision_no integer not null check (revision_no > 0)
event_name text not null
changed_by_actor_id text null
change_reason text null
changed_at timestamptz not null
snapshot_json jsonb not null
unique(product_id, revision_no)
```

Indexes:

```sql
create index product_versions_product_changed_at_idx on product_versions(product_id, changed_at);
create index product_versions_event_name_idx on product_versions(event_name);
```

## API CONTRACT PLAN

Service catalog:

```text
GET    /api/service-catalog/items
POST   /api/service-catalog/items
GET    /api/service-catalog/items/:id
PUT    /api/service-catalog/items/:id
PATCH  /api/service-catalog/items/:id/activate
PATCH  /api/service-catalog/items/:id/deactivate
GET    /api/service-catalog/items/lookup
```

Product catalog:

```text
GET    /api/products
POST   /api/products
GET    /api/products/:id
PUT    /api/products/:id
DELETE /api/products/:id
PATCH  /api/products/:id/restore
GET    /api/products/:id/versions
GET    /api/products/lookup
```

Envelope:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

Errors:

```json
{
  "success": false,
  "error": {
    "code": "validation_failed",
    "message": "Validasi gagal",
    "fields": {}
  },
  "meta": {}
}
```

## PRODUCT LIST QUERY CONTRACT

Allowed query:

```text
q
page
per_page
sort_by: nama_barang,merek,ukuran,harga_jual_rupiah,stok_saat_ini
sort_dir: asc,desc
status: active,deleted,all
merek
ukuran_min
ukuran_max
harga_min
harga_max
```

Rules:

- default status: `active`.
- default page: `1`.
- default per_page: `10`.
- max per_page should be explicitly chosen before implementation.
- reject invalid ranges.
- reject unknown sort/filter fields.

## TEST PLAN

Domain tests:

- product create trims/normalizes valid input.
- product rejects empty name/brand.
- product rejects price <= 0.
- product rejects one-sided threshold.
- product rejects negative threshold.
- product rejects critical threshold > reorder point.
- service normalizer matches punctuation/spacing variants.

Usecase tests:

- create product success.
- create duplicate behavior.
- update product success.
- update missing product.
- soft delete active product.
- soft delete missing/deleted product.
- restore deleted product.
- restore active/missing product.

PostgreSQL tests:

- product constraints.
- active code uniqueness ignores deleted products.
- version revision uniqueness.
- service normalized name uniqueness.
- seed idempotency by normalized name.

HTTP/API tests:

- capability/authz required for protected routes.
- create/update validation errors.
- list filter/sort/pagination contract.
- show includes detail and version timeline.
- lookup is bounded.
- disabled capability returns `403` before usecase.

Performance proof:

- product list with inventory projection has query plan/index proof before `<1s` claim.
- lookup has bounded limit and no unbounded scan claim.

## IMPLEMENTATION ORDER

1. Resolve product duplicate policy.
2. Add service catalog domain/usecase contracts and tests.
3. Add service catalog migration and PostgreSQL adapter.
4. Add service catalog seed profile.
5. Add product domain tests and value rules.
6. Add product migration and PostgreSQL adapter tests.
7. Add product usecases for create/update/show/list/lookup/delete/restore/version timeline.
8. Add presentation DTOs.
9. Add Echo handlers/routes after capability foundation is ready.
10. Add API and capability-disabled tests.

## NEXT ACTIVE STEP

Resolve product duplicate policy before writing PostgreSQL migration:

```text
Choose whether Go preserves Laravel's "same name/brand/size allowed when distinct non-null kode_barang" behavior,
or tightens product identity to forbid same normalized name/brand/size for active products.
```

## DOD

- Blueprint exists before implementation.
- Product and service catalog domain contracts are declared.
- PostgreSQL schema is proposed but does not claim migration completion.
- API routes and capability keys are proposed.
- Test plan maps Laravel behavior into Go tests.
- The only open policy blocker is explicit and small.
