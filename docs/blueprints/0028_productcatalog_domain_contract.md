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

# ProductCatalog Domain Contract Blueprint

## Status

Accepted.

## Date

2026-06-08

## Active Scope

ProductCatalog domain contract and duplicate-policy decision planning only.

This blueprint does not implement ProductCatalog.

## FACT

- ProductCatalog source evidence exists in `docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md`.
- ProductCatalog is part of Business Phase 1 together with ServiceCatalog.
- ServiceCatalog has implementation proof, but ProductCatalog does not.
- ProductCatalog domain contract has not been accepted yet.
- Product duplicate policy must be accepted before PostgreSQL unique indexes are finalized.
- Product stock adjustment belongs to future Inventory scope, not ProductCatalog.

## GAP

- No accepted standalone ProductCatalog domain contract exists yet.
- No ProductCatalog Go business module implementation proof exists.
- No ProductCatalog PostgreSQL migration is accepted yet.
- No ProductCatalog API route/capability seed implementation proof exists.
- Full Inventory stock adjustment source is not available and must stay deferred.
- Product duplicate policy needs owner acceptance before implementation.

## DECISION

Recommended duplicate policy:

Preserve Laravel-compatible duplicate behavior.

Rules:

1. Active non-null `kode_barang` must be unique.
2. `kode_barang` is trimmed.
3. Blank `kode_barang` becomes null.
4. Same active normalized product identity is rejected when the code exception does not apply.
5. Same normalized `nama_barang + merek + ukuran` is allowed only when both existing and candidate products have distinct non-null `kode_barang`.
6. Do not create a strict active unique PostgreSQL index on `(nama_barang_normalized, merek_normalized, ukuran)`.
7. Enforce nuanced product identity duplicate behavior in ProductCatalog usecase/repository guard inside a transaction.
8. Use PostgreSQL partial unique index only for active non-null `kode_barang`.

## DUPLICATE POLICY OPTIONS

### Option A - Recommended: Laravel-compatible identity exception

Database:

- active non-null `kode_barang` is unique.

Usecase/repository guard:

- same normalized name/brand/size is rejected unless both products have distinct non-null `kode_barang`.

Pros:

- Preserves Laravel behavior.
- Avoids data tightening during migration.
- Supports multiple coded products with same display identity.

Cons:

- Requires careful transactional duplicate guard.
- Cannot rely only on one simple database unique index.

### Option B - Strict business identity uniqueness

Database:

- active unique `(nama_barang_normalized, merek_normalized, ukuran)`.

Pros:

- Simple database integrity.
- Easier concurrency control.

Cons:

- Stricter than Laravel.
- May reject product data Laravel currently allows.

### Option C - Code-only uniqueness

Database:

- active non-null `kode_barang` unique only.

Pros:

- Simple migration.

Cons:

- Too loose versus Laravel duplicate guard.
- Allows accidental duplicate identity when code is null.

## DOMAIN CONTRACT

Domain:

```text
productcatalog
```

Aggregate root:

```text
Product
```

Owned tables:

```text
products
product_versions
```

Read-only dependency:

```text
product_inventory
```

`product_inventory` may be used only for list/detail/lookup projection.

ProductCatalog must not mutate stock.

## FIELDS

Product:

```text
id
kode_barang nullable
nama_barang
nama_barang_normalized
merek
merek_normalized
ukuran nullable
harga_jual_rupiah
reorder_point_qty nullable
critical_threshold_qty nullable
deleted_at nullable
deleted_by_actor_id nullable
delete_reason nullable
created_at
updated_at
```

ProductVersion:

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

## LIFECYCLE

Statuses:

```text
active  = deleted_at is null
deleted = deleted_at is not null
```

Rules:

- New product is active.
- Soft delete sets deleted metadata.
- Restore clears deleted metadata.
- Physical delete is forbidden.
- Deleted products remain for historical references.
- Code from soft-deleted product may be reused by active product.

## ALLOWED OPERATIONS

```text
create
update
soft delete
restore
show
list
lookup
version timeline
```

## FORBIDDEN OPERATIONS

```text
physical delete
stock adjustment create
stock adjustment reverse
inventory costing mutation
inventory movement mutation
```

Deferred inventory operations:

```text
inventory.stock_adjustments.create
inventory.stock_adjustments.reverse
```

## INVARIANTS

- `id` is required.
- `nama_barang` is required after trim.
- `merek` is required after trim.
- `kode_barang` blank normalizes to null.
- `nama_barang_normalized` is generated by server normalizer.
- `merek_normalized` is generated by server normalizer.
- `harga_jual_rupiah > 0`.
- `reorder_point_qty` and `critical_threshold_qty` must both be null or both filled.
- Threshold values must be non-negative.
- `critical_threshold_qty <= reorder_point_qty`.
- Duplicate behavior follows accepted duplicate policy.
- Create/update/delete/restore must write product row, version row, and audit event or audit intent transactionally.

## POSTGRESQL SCHEMA PROPOSAL

Table:

```text
products
```

Columns:

```text
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
);

check (
  reorder_point_qty is null
  or critical_threshold_qty <= reorder_point_qty
);
```

Indexes:

```sql
create unique index products_kode_barang_active_unique
on products (kode_barang)
where deleted_at is null and kode_barang is not null;

create index products_deleted_at_idx
on products (deleted_at);

create index products_name_norm_idx
on products (nama_barang_normalized);

create index products_brand_norm_idx
on products (merek_normalized);

create index products_identity_guard_idx
on products (nama_barang_normalized, merek_normalized, ukuran)
where deleted_at is null;
```

Important:

Do not create this strict index under recommended policy:

```sql
-- forbidden under Option A
create unique index products_business_identity_active_unique
on products (nama_barang_normalized, merek_normalized, ukuran)
where deleted_at is null;
```

Table:

```text
product_versions
```

Columns:

```text
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
create index product_versions_product_changed_at_idx
on product_versions(product_id, changed_at);

create index product_versions_event_name_idx
on product_versions(event_name);
```

## API ROUTE PROPOSAL

ProductCatalog routes:

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

Route ordering:

Register `/api/products/lookup` before `/api/products/:id`.

Inventory routes are deferred:

```text
POST  /api/inventory/products/:productId/stock-adjustments
PATCH /api/inventory/products/:productId/stock-adjustments/:adjustmentId/reverse
```

## QUERY CONTRACT

List query:

```text
q nullable string
page int default 1 min 1
per_page int default 10 min 1 max 50
sort_by one of nama_barang,merek,ukuran,harga_jual_rupiah,stok_saat_ini
sort_dir one of asc,desc
status one of active,deleted,all default active
merek nullable string
ukuran_min nullable int min 0
ukuran_max nullable int min 0
harga_min nullable int min 0
harga_max nullable int min 0
```

Rules:

- Reject unknown filters.
- Reject unknown sort fields.
- Reject invalid ranges.
- `ukuran_min <= ukuran_max` when both exist.
- `harga_min <= harga_max` when both exist.
- Lookup must be bounded to max 50 rows.

## CAPABILITY KEYS

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

Deferred inventory capability keys:

```text
inventory.stock_adjustments.create
inventory.stock_adjustments.reverse
```

## AUTHORIZATION PROPOSAL

Initial permissions:

```text
products.read
products.manage
```

Mapping:

```text
products.list          -> products.read
products.lookup        -> products.read
products.show          -> products.read
products.versions.list -> products.read
products.create        -> products.manage
products.update        -> products.manage
products.delete        -> products.manage
products.restore       -> products.manage
```

Capability middleware must reject disabled protected routes before validation/usecase execution.

## AUDIT AND VERSIONING

Auditable operations:

```text
create
update
soft delete
restore
```

Audit event fields:

```text
bounded_context = product_catalog
aggregate_type = product
aggregate_id
event_name
occurred_at
actor_id
actor_role
reason
source_channel
metadata_json
```

Version events:

```text
product_created
product_updated
product_soft_deleted
product_restored
```

Each version snapshot must include the product state and revision number.

## AUDIT PERFORMANCE AND PAYLOAD CONTRACT

Audit policy:

```text
transactional audit intent/outbox in the hot path; full audit telemetry deferred
```

ProductCatalog create/update/soft-delete/restore must not perform full telemetry, reporting projection, notification, external logging, or heavy audit enrichment inside the request hot path.

The request hot path must only write minimal audit intent/outbox data inside the same transaction as the product row and product version row.

Required hot-path transaction:

```text
begin transaction
write products
write product_versions
write audit_outbox or audit_intents minimal record
commit transaction
return API response
```

Forbidden hot-path audit behavior:

```text
external network logging
full telemetry export
dashboard/report generation
notification dispatch
untracked goroutine audit writes
slow audit enrichment
```

Full audit materialization may run later through a background worker that reads audit outbox/intent records.

Audit intent/outbox minimal fields:

```text
id
bounded_context = product_catalog
aggregate_type = product
aggregate_id
event_name
operation
actor_id
actor_role
occurred_at
reason nullable
source_channel
revision_no
before_snapshot nullable
after_snapshot nullable
metadata_json
status = pending
created_at
```

Audit meaning:

- `actor_id` answers who performed the action.
- `actor_role` answers which role/authority performed the action.
- `occurred_at` answers when the action happened and must be server-owned.
- `event_name` / `operation` answers what action happened.
- `before_snapshot` records relevant data before the action when available.
- `after_snapshot` records relevant data after the action when available.
- `reason` records why the action happened when supplied.
- `metadata_json` may include request id, correlation id, product version id, and compact domain-specific context.

Snapshot policy:

- create may have `before_snapshot = null` and `after_snapshot = created product snapshot`.
- update should include both before and after product snapshots or references to before/after versions.
- soft delete should include before active snapshot and after deleted snapshot.
- restore should include before deleted snapshot and after restored active snapshot.

Performance rule:

The audit outbox/intent write must be bounded and local to PostgreSQL. The API response must not wait for full audit enrichment. This preserves the <1 second API mission while keeping audit evidence durable.

## TEST MATRIX

Domain tests:

- trims product name and brand.
- normalizes product name and brand.
- blank `kode_barang` becomes null.
- rejects blank name.
- rejects blank brand.
- rejects price <= 0.
- rejects one-sided threshold.
- rejects negative threshold.
- rejects critical threshold greater than reorder point.

Usecase tests:

- create product success.
- update product success.
- show missing product returns not found.
- soft delete active product.
- soft delete missing product returns not found.
- soft delete deleted product is rejected.
- restore deleted product.
- restore active product is rejected.
- restore missing product returns not found.
- duplicate active code is rejected.
- code from soft-deleted product can be reused.
- duplicate identity is rejected when code exception does not apply.
- same identity is allowed when both products have distinct non-null `kode_barang`.

PostgreSQL tests:

- active non-null code uniqueness.
- deleted product does not block code reuse.
- threshold checks are enforced.
- positive price check is enforced.
- version revision uniqueness is enforced.
- product write + version write + audit write rollback together.
- duplicate guard has transaction/concurrency proof.

HTTP tests:

- list requires authn/authz/capability.
- lookup requires authn/authz/capability.
- show requires authn/authz/capability.
- create requires manage permission and capability.
- update requires manage permission and capability.
- delete requires manage permission and capability.
- restore requires manage permission and capability.
- versions list requires read permission and capability.
- disabled capability returns 403 before usecase execution.
- validation errors use public error envelope.
- success responses use public success envelope.

Deferred Inventory tests:

- stock adjustment reduces projection.
- stock adjustment rejects negative stock.
- stock adjustment reversal restores projection.
- stock adjustment reversal cannot run twice.

These deferred tests must not be implemented in ProductCatalog scope.

## IMPLEMENTATION SLICE ORDER

1. Accept this ProductCatalog domain contract and duplicate policy.
2. Add ProductCatalog domain/value tests.
3. Add ProductCatalog ports and usecase contract tests.
4. Add ProductCatalog PostgreSQL schema and adapter tests.
5. Add ProductCatalog usecases.
6. Add ProductCatalog response DTOs/presenters.
7. Add Echo handlers/routes after capability metadata is accepted.
8. Add capability seed migration after route/capability list is accepted.
9. Add route capability audit coverage and disabled-capability tests.
10. Add read-only inventory projection only if inventory projection table exists.
11. Keep inventory stock adjustment for future Inventory scope.

## IMPLEMENTATION BLOCKERS

Do not start implementation until:

- Owner accepts duplicate policy.
- Owner accepts this domain contract.
- PostgreSQL schema proposal is accepted.
- Capability keys are accepted.
- Permission mapping is accepted.
- Audit/versioning behavior is accepted.
- ProductCatalog implementation slice blueprint is created and accepted.

## ACCEPTANCE

Accepted on: 2026-06-09

Owner decision:

```text
ACCEPT OPTION A
```

Accepted duplicate policy:

```text
Preserve Laravel-compatible duplicate behavior.
```

Accepted rules:

- Active non-null `kode_barang` must be unique.
- Blank `kode_barang` becomes null.
- Same normalized `nama_barang + merek + ukuran` is rejected when the code exception does not apply.
- Same normalized identity is allowed only when both existing and candidate products have distinct non-null `kode_barang`.
- PostgreSQL must not use a strict active unique index on `(nama_barang_normalized, merek_normalized, ukuran)`.
- PostgreSQL should use a partial unique index for active non-null `kode_barang`.
- The nuanced identity duplicate rule belongs in ProductCatalog usecase/repository guard inside a transaction.
- Inventory stock adjustment remains deferred to future Inventory scope.

## NEXT ACTIVE STEP

Owner decision accepted.

Next valid step:

Plan ProductCatalog implementation slice 1 without coding yet.
