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

# ProductCatalog PostgreSQL Persistence Slice Blueprint

## Status

Accepted.

## Date

2026-06-10

## Active Scope

Plan ProductCatalog PostgreSQL persistence after ProductCatalog implementation slice 1 was closed.

Source slice:

```text
docs/blueprints/0029_productcatalog_implementation_slice_1.md
```

Closeout handoff:

```text
docs/handoffs/2026-06-10-productcatalog-implementation-slice-1-closeout.md
```

## FACT

ProductCatalog implementation slice 1 is closed with local proof and GitHub connector validation.

Closed Go package scope:

```text
internal/modules/productcatalog/domain
internal/modules/productcatalog/ports
internal/modules/productcatalog/usecase
```

Laravel migration evidence shows ProductCatalog persistence requires:

```text
products
product_versions
```

Laravel product table fields observed:

```text
id
kode_barang nullable
nama_barang
nama_barang_normalized nullable
merek
merek_normalized nullable
ukuran nullable
harga_jual
reorder_point_qty nullable
critical_threshold_qty nullable
deleted_at nullable
deleted_by_actor_id nullable
delete_reason nullable
active_unique_marker legacy MySQL generated column
```

Laravel product_versions fields observed:

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

Laravel indexes observed:

```text
products_merek_idx
products_ukuran_idx
products_harga_jual_idx
products_duplicate_lookup_idx
products_deleted_at_idx
products_nama_barang_normalized_idx
products_merek_normalized_idx
products_kode_barang_unique legacy active unique
products_business_identity_unique legacy active unique
product_versions_product_revision_unique
product_versions_product_changed_at_idx
product_versions_event_name_idx
fk_product_versions_product
```

## GAP

No ProductCatalog PostgreSQL migration exists in Go yet.

No ProductCatalog PostgreSQL repository adapter exists in Go yet.

No ProductCatalog PostgreSQL integration tests exist yet.

No ProductCatalog runtime HTTP surface exists yet.

No ProductCatalog route registration, presenter, capability seed, inventory mutation, or UI exists yet.

## DECISION

This slice should implement ProductCatalog PostgreSQL persistence only.

Do not implement Echo HTTP transport.

Do not register routes.

Do not add capability seed migrations.

Do not implement inventory stock mutation.

Do not implement UI.

Do not start ProductCatalog runtime HTTP slice.

## SCOPE-IN

Expected new files:

```text
migrations/0011_create_product_catalog_tables.up.sql
migrations/0011_create_product_catalog_tables.down.sql
internal/platform/postgres/product_repository.go
internal/platform/postgres/product_repository_query.go
internal/platform/postgres/product_repository_row.go
internal/platform/postgres/product_repository_read.go
internal/platform/postgres/product_repository_write.go
internal/platform/postgres/product_repository_list.go
internal/platform/postgres/product_repository_lookup.go
internal/platform/postgres/product_version_repository.go
internal/platform/postgres/product_repository_integration_helpers_test.go
internal/platform/postgres/product_repository_create_integration_test.go
internal/platform/postgres/product_repository_update_integration_test.go
internal/platform/postgres/product_repository_lifecycle_integration_test.go
internal/platform/postgres/product_repository_query_integration_test.go
internal/platform/postgres/product_version_repository_integration_test.go
```

If file-size audit requires smaller files, split files without changing package boundaries.

## SCOPE-OUT

```text
internal/modules/productcatalog/transport/http
internal/presentation/http/id/productcatalog
cmd/server route wiring
capability seed migration
inventory stock mutation
inventory stock reversal
UI
```

## POSTGRESQL SCHEMA

Table:

```text
products
```

Columns:

```sql
id text primary key,
kode_barang text null,
nama_barang text not null,
nama_barang_normalized text not null,
merek text not null,
merek_normalized text not null,
ukuran integer null,
harga_jual bigint not null check (harga_jual > 0),
reorder_point_qty integer null check (reorder_point_qty >= 0),
critical_threshold_qty integer null check (critical_threshold_qty >= 0),
deleted_at timestamptz null,
deleted_by_actor_id text null,
delete_reason text null,
created_at timestamptz not null default now(),
updated_at timestamptz not null default now(),
constraint products_threshold_pair_check check (
  (reorder_point_qty is null and critical_threshold_qty is null)
  or
  (reorder_point_qty is not null and critical_threshold_qty is not null)
),
constraint products_threshold_order_check check (
  critical_threshold_qty is null
  or reorder_point_qty is null
  or critical_threshold_qty <= reorder_point_qty
)
```

Table:

```text
product_versions
```

Columns:

```sql
id text primary key,
product_id text not null references products(id) on delete restrict,
revision_no integer not null,
event_name text not null,
changed_by_actor_id text null,
change_reason text null,
changed_at timestamptz not null,
snapshot_json jsonb not null,
created_at timestamptz not null default now()
```

## POSTGRESQL INDEXES

Product read/search indexes:

```sql
create index products_merek_idx
on products (merek);

create index products_ukuran_idx
on products (ukuran);

create index products_harga_jual_idx
on products (harga_jual);

create index products_duplicate_lookup_idx
on products (nama_barang, merek, ukuran);

create index products_deleted_at_idx
on products (deleted_at);

create index products_nama_barang_normalized_idx
on products (nama_barang_normalized);

create index products_merek_normalized_idx
on products (merek_normalized);
```

Product active code uniqueness:

```sql
create unique index products_kode_barang_unique
on products (kode_barang)
where deleted_at is null and kode_barang is not null;
```

Important duplicate-policy decision:

Do not create a strict active unique index on:

```text
nama_barang_normalized, merek_normalized, ukuran
```

Reason:

Laravel legacy MySQL used active_unique_marker for soft-delete uniqueness, but accepted Go duplicate policy Option A preserves the business exception that allows same normalized identity when both products have distinct non-null kode_barang.

Therefore, business identity duplicate enforcement remains in repository/usecase duplicate guard, not a strict PostgreSQL unique index.

Product version indexes:

```sql
create unique index product_versions_product_revision_unique
on product_versions (product_id, revision_no);

create index product_versions_product_changed_at_idx
on product_versions (product_id, changed_at);

create index product_versions_event_name_idx
on product_versions (event_name);
```

## REPOSITORY BEHAVIOR

ProductRepository adapter must implement:

```text
Create
Update
FindByID
```

ProductReader adapter must implement:

```text
GetByID
List
Lookup
```

ProductVersionRepository adapter must implement:

```text
Append
ListByProductID
```

ProductDuplicateChecker adapter must implement:

```text
CheckCreateDuplicate
CheckUpdateDuplicate
```

Rules:

Repository accepts and returns domain/usecase port objects.

Repository must not import Echo or HTTP.

Repository must not own capability behavior.

Repository must not mutate inventory.

Repository must preserve domain-generated normalized values.

Repository must translate missing rows to ports.ErrProductNotFound where the port expects not found.

Repository must support transaction context pattern already used by platform/postgres.

Product write, product version append, and audit write must be usable inside the same transaction in later orchestration.

## QUERY BEHAVIOR

List should support existing ProductListQuery fields:

```text
Search
Status
Page
PerPage
```

Status behavior:

```text
active -> deleted_at is null
deleted -> deleted_at is not null
all -> no deleted_at filter
```

Lookup should support existing ProductLookupQuery fields:

```text
Query
Limit
IncludeDeleted
```

Lookup rules:

default active-only unless IncludeDeleted is true;

bounded max limit must remain usecase/adapter safe;

deterministic order by normalized name, brand, size, and id unless a later blueprint changes it.

## PERFORMANCE AND FLEXIBILITY STANDARD

This slice must preserve fast admin CRUD, list, lookup, and show/read behavior without hardcoding runtime HTTP concerns.

Performance rules:

```text
Create -> primary write plus active code uniqueness check must rely on indexed/unique paths.
Update -> primary key lookup/write path must use products primary key.
FindByID/GetByID/show -> primary key lookup must use products primary key.
List active/default -> must be bounded by Page and PerPage and must filter deleted_at predictably.
Lookup -> must be bounded by Limit and must not scan unbounded rows.
Version list -> must use product_versions_product_changed_at_idx or product_versions_product_revision_unique-compatible ordering.
Duplicate guard -> active kode_barang check must use products_kode_barang_unique.
```

Query-plan proof rule:

Integration proof should include EXPLAIN or EXPLAIN ANALYZE notes for list, lookup, show/get-by-id, and duplicate guard queries when a local database is available.

No fake SLA rule:

Do not claim sub-second or millisecond performance without local database proof.

Flexibility rules:

Keep ProductListQuery and ProductLookupQuery adapter translation centralized.
Do not spread SQL filter construction across usecase or HTTP layers.
Keep future sort/filter additions localized to product_repository_query.go.
Do not add inventory projection joins in this slice.
Do not over-index speculative filters until a query exists and proof shows the need.

Baseline acceptable behavior:

CRUD and show must be index-backed.
List and lookup must be bounded.
Search behavior may start with normalized LIKE/ILIKE-compatible filtering, but trigram or full-text search must be deferred unless proof shows the current strategy is too slow.


## TRANSACTION POLICY

This slice may use the existing platform/postgres transaction context pattern.

Write-side operations must be compatible with later usecase orchestration where product write, version append, duplicate check, and audit record happen inside one transaction.

If transaction reuse is blocked by current platform shape, document the gap in the handoff and keep the behavior isolated.

## TEST MATRIX

Migration proof:

```text
Up migration creates products and product_versions.
Down migration drops product_versions before products.
product_versions.product_id restricts product delete.
products.harga_jual rejects non-positive values.
threshold pair check rejects one-sided threshold.
threshold order check rejects critical threshold above reorder point.
active non-null kode_barang is unique.
kode_barang can be reused after soft delete.
strict business identity unique index is not present.
```

Repository integration tests:

```text
Create stores product fields.
Create stores normalized fields from domain.
Create rejects duplicate active code.
Create allows code reuse from soft-deleted product.
Update changes fields.
FindByID returns product.
FindByID returns ErrProductNotFound for missing product.
GetByID returns product detail.
List filters active, deleted, and all.
List supports pagination.
Lookup excludes deleted by default.
Lookup includes deleted when requested.
Lookup respects limit.
ProductVersionRepository.Append stores version.
ProductVersionRepository.ListByProductID returns versions ordered deterministically by revision or changed_at.
ProductVersionRepository.ListByProductID returns empty list for products with no versions.
```

## PROOF REQUIRED

Focused PostgreSQL proof:

```bash
go test ./internal/platform/postgres/... -run Product
```

Focused ProductCatalog proof:

```bash
go test ./internal/modules/productcatalog/...
```

Aggregate proof:

```bash
make verify
```

Migration and query-plan proof if local DB is available:

```bash
make db-up
make db-status
```

Additional local DB proof should capture EXPLAIN or EXPLAIN ANALYZE for:

```text
products primary-key show/get-by-id
products active list first page
products lookup bounded search
products active kode_barang duplicate guard
product_versions list by product_id
```

Expected migration status must include:

```text
0011_create_product_catalog_tables.up.sql applied
```

## API CONTRACT IMPACT

No HTTP API contract change in this slice.

ProductCatalog runtime API remains out of scope.

## CAPABILITY IMPACT

No capability key or seed change in this slice.

Capability seed belongs to a later ProductCatalog runtime/capability slice.

## ARCHITECTURE IMPACT

This slice adds ProductCatalog persistence under:

```text
internal/platform/postgres
migrations
```

It must not add imports from domain/usecase into transport or HTTP layers.

Hexagonal boundaries must remain enforced by make verify.

## RISKS

Copying Laravel active_unique_marker literally would preserve MySQL shape but not fit PostgreSQL cleanly.

Creating a strict identity unique index would break accepted Option A duplicate behavior.

Search with leading wildcard may need trigram or prefix strategy later.

Inventory projection joins are out of scope and must not leak into product persistence writes.

## STEP ORDER

1. Accept or revise this blueprint.
2. Add migration only.
3. Add repository adapter skeletons.
4. Add repository integration helpers.
5. Implement ProductRepository create/find/update behavior.
6. Implement ProductReader get/list/lookup behavior.
7. Implement ProductVersionRepository append/list behavior.
8. Implement ProductDuplicateChecker behavior.
9. Run focused PostgreSQL proof.
10. Run aggregate make verify.
11. Update ledger and handoff.
12. Close this slice only after local proof and connector validation.

## ACCEPTANCE RULE

This blueprint is accepted. Implementation must follow the documented step order and start with migration-only work.
