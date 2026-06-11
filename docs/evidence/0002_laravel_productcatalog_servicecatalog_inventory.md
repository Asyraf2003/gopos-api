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

# Laravel ProductCatalog And ServiceCatalog Inventory

## Source

User-provided Laravel command output for:

- product soft delete/search/unique/threshold migrations;
- service catalog seed migration;
- `app/Core/ProductCatalog/**`;
- `app/Application/ProductCatalog/**`;
- `app/Ports/Out/ProductCatalog/**`;
- `app/Adapters/Out/ProductCatalog/**`;
- product controllers and requests;
- ProductCatalog feature tests;
- ServiceCatalog normalizer unit test.

This evidence is source inventory only. It does not prove Go implementation.

## Product Domain Facts

Product master fields:

```text
id
kode_barang nullable
nama_barang
nama_barang_normalized
merek
merek_normalized
ukuran nullable
harga_jual
reorder_point_qty nullable
critical_threshold_qty nullable
deleted_at nullable
deleted_by_actor_id nullable
delete_reason nullable
active_unique_marker legacy MySQL generated column
```

Product version fields:

```text
id
product_id
revision_no
event_name
changed_by_actor_id nullable
change_reason nullable
changed_at
snapshot_json
```

Important Laravel behavior:

- `kode_barang` is trimmed and blank becomes null.
- `nama_barang` and `merek` are trimmed.
- `nama_barang_normalized` and `merek_normalized` compact whitespace and lowercase.
- `harga_jual` must be greater than zero.
- `reorder_point_qty` and `critical_threshold_qty` must both be null or both filled.
- `reorder_point_qty` and `critical_threshold_qty` must be non-negative.
- `critical_threshold_qty` must not exceed `reorder_point_qty`.
- Soft delete keeps product rows and historical references.
- Restore clears deleted metadata.
- Product create/update records product version history and audit event in the versioned writer.
- Product lookup joins inventory and is bounded to max 50 rows.
- Product table uses pagination, filters, sorting, and inventory projection join.
- Product stock adjustment belongs to inventory behavior, even though product admin routes expose it.

## Product Duplicate Rules

Legacy MySQL used generated `active_unique_marker` to allow duplicate values after soft delete.

Go/PostgreSQL target should not copy the MySQL generated `TINYINT(1)` pattern literally.

PostgreSQL target should use partial unique indexes:

```sql
create unique index products_kode_barang_active_unique
on products (kode_barang)
where deleted_at is null and kode_barang is not null;

create unique index products_business_identity_active_unique
on products (nama_barang_normalized, merek_normalized, ukuran)
where deleted_at is null;
```

Important behavioral exception from Laravel:

- Two active products with the same `nama_barang`, `merek`, and `ukuran` may be allowed when both have distinct non-null `kode_barang`.
- Duplicate guard treats matching identity rows as allowed when candidate and existing `kode_barang` are both non-null and different.

PostgreSQL design impact:

- A strict unique index on `(nama_barang_normalized, merek_normalized, ukuran) where deleted_at is null` would be stricter than Laravel behavior.
- Product duplicate policy must be finalized before migration.
- If preserving Laravel behavior, DB uniqueness should enforce active unique `kode_barang`, while business identity duplicate rules stay in usecase/repository guard.

## Product Operations Inventory

| Operation | Laravel Source | Go Operation |
| --- | --- | --- |
| list table | `GetProductTableHandler`, `ProductTableReaderPort` | `products.list` |
| lookup | `ProductLookupReaderPort` | `products.lookup` |
| show detail | `GetProductDetailHandler`, `ProductDetailReaderPort` | `products.show` |
| create | `CreateProductHandler`, `ProductWriterPort` | `products.create` |
| update | `UpdateProductHandler`, duplicate guard | `products.update` |
| soft delete | `SoftDeleteProductHandler`, `ProductLifecyclePort` | `products.delete` |
| restore | `RestoreProductHandler`, `ProductLifecyclePort` | `products.restore` |
| stock adjustment | inventory use case through product admin route | `inventory.stock_adjustments.create` |
| stock adjustment reversal | inventory use case through product admin route | `inventory.stock_adjustments.reverse` |

## Product API Candidate Contract

Go should expose API-only routes, not page routes:

```text
GET    /api/products
POST   /api/products
GET    /api/products/:id
PUT    /api/products/:id
DELETE /api/products/:id
PATCH  /api/products/:id/restore
GET    /api/products/lookup
```

Stock adjustment should belong to inventory:

```text
POST  /api/inventory/products/:productId/stock-adjustments
PATCH /api/inventory/products/:productId/stock-adjustments/:adjustmentId/reverse
```

## Product Query Contract

List query fields from Laravel:

```text
q nullable string
page int default 1 min 1
per_page int default 10, currently only 10 in Laravel
sort_by one of nama_barang,merek,ukuran,harga_jual,stok_saat_ini
sort_dir one of asc,desc
status one of active,deleted,all default active
merek nullable string
ukuran_min nullable int min 0
ukuran_max nullable int min 0
harga_min nullable int min 0
harga_max nullable int min 0
```

Range validation:

- `ukuran_min <= ukuran_max` when both exist.
- `harga_min <= harga_max` when both exist.

Performance direction:

- Table list and lookup must be bounded and indexed.
- `stok_saat_ini` sorting joins `product_inventory`; verify query plan before claiming `<1s`.
- Search with leading wildcard may not satisfy `<1s` at scale; PostgreSQL trigram or normalized prefix strategy may be needed later.

## Product Audit And Versioning

Version events observed:

```text
Produk dibuat
Produk diperbarui
product_soft_deleted
product_restored
```

Audit event fields from Laravel writer:

```text
bounded_context = product_catalog
aggregate_type = product
aggregate_id = product id
event_name
occurred_at
actor_id
actor_role
reason
source_channel
metadata_json includes product snapshot and revision_no
```

Go target:

- product create/update/delete/restore must be transactional with product write, version row, and audit event.
- version and audit writes must happen through ports/usecase/adapters, not in HTTP handlers.

## Product Tests To Preserve

Minimum Go test matrix derived from Laravel:

- create product stores new product;
- create rejects duplicate identity when code exception does not apply;
- create allows same name/brand/size when distinct non-null code exception applies, if this behavior is accepted;
- create rejects zero price;
- create rejects one-sided threshold;
- create rejects critical threshold greater than reorder point;
- update existing product;
- update rejects duplicate active product code;
- update normalizes blank `kode_barang` to null;
- update rejects negative thresholds;
- update allows reuse of code from soft-deleted product;
- missing product returns not found behavior;
- list/table requires authz/capability;
- list/table returns inventory projection;
- list/table validates range filters;
- soft delete stores deleted metadata;
- soft delete keeps supplier payable, invoice line snapshot, inventory projection, costing, and movement history;
- restore deleted product;
- restore rejects missing or active product;
- stock adjustment reduces projection precisely;
- stock adjustment rejects negative stock;
- stock adjustment reversal restores projection;
- stock adjustment reversal cannot run twice.

## ServiceCatalog Facts

Table:

```text
service_catalog_items
id
name
normalized_name unique
default_price_rupiah
is_active default true
created_at
updated_at
```

Seed names:

```text
Sok Kopling (Besar) 120000
Sok Kopling (Kecil) 110000
Setting In (Kecil) 70000
Setting Ex (Kecil) 70000
Setting In (Besar) 85000
Setting Ex (Besar) 85000
Bosklep In (Kecil) 60000
Bosklep Ex (Kecil) 60000
Bosklep In (Besar) 75000
Bosklep Ex (Besar) 75000
Pasang Stang (Kecil) 50000
Pasang Stang (Besar) 60000
```

Normalizer behavior:

- Parentheses/plain variant should match.
- Spacing and punctuation are compacted.
- Example: `Setting--In   (Kecil)` -> `setting in kecil`.

Go target:

- implement normalizer as domain/service function.
- normalized name is unique.
- seed should be idempotent by `normalized_name`.

## PostgreSQL Design Notes

Product table target should include:

```text
id text primary key
kode_barang text null
nama_barang text not null
nama_barang_normalized text not null
merek text not null
merek_normalized text not null
ukuran integer null
harga_jual_rupiah bigint not null check > 0
reorder_point_qty integer null check >= 0
critical_threshold_qty integer null check >= 0
deleted_at timestamptz null
deleted_by_actor_id text null
delete_reason text null
created_at timestamptz not null default now()
updated_at timestamptz not null default now()
```

Additional product check:

```text
threshold pair must be both null or both non-null
critical_threshold_qty <= reorder_point_qty when both non-null
```

Product versions:

```text
id text primary key
product_id text not null references products(id) restrict
revision_no integer not null check > 0
event_name text not null
changed_by_actor_id text null
change_reason text null
changed_at timestamptz not null
snapshot_json jsonb not null
unique(product_id, revision_no)
```

Service catalog:

```text
id text primary key
name text not null
normalized_name text not null unique
default_price_rupiah bigint not null check > 0
is_active boolean not null default true
created_at timestamptz not null default now()
updated_at timestamptz not null default now()
```

## GAP

- Full `ServiceCatalog` source was not included, except the normalizer test and seed migration.
- Full inventory use cases for stock adjustment were not included.
- Full product presenters/view data are not needed for API implementation, but API output DTOs still need to be designed.
- Product duplicate policy conflict must be resolved before DB unique indexes are finalized.
- Current Go capability foundation blueprint is not implemented yet.

## Recommended Next Step

Create the first domain blueprint for catalog foundation:

```text
servicecatalog + productcatalog domain contract
PostgreSQL schema proposal
API contract proposal
capability keys
test matrix
implementation order
```
