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

# ProductCatalog PostgreSQL Migration Checkpoint Handoff

## Date

2026-06-10

## Active Scope

ProductCatalog PostgreSQL persistence migration-only implementation.

## FACT

ProductCatalog PostgreSQL persistence blueprint is accepted:

docs/blueprints/0030_productcatalog_postgres_persistence_slice.md

Migration-only implementation exists:

migrations/0011_create_product_catalog_tables.up.sql
migrations/0011_create_product_catalog_tables.down.sql

Local database migration proof passed.

Applied migration:

0011_create_product_catalog_tables.up.sql

Created tables:

products
product_versions

Created ProductCatalog performance-supporting indexes:

products_kode_barang_unique
products_active_list_idx
products_active_identity_lookup_idx
product_versions_product_revision_unique
product_versions_product_changed_at_idx
product_versions_event_name_idx

The migration preserves the accepted PostgreSQL direction:

- no MySQL active_unique_marker generated column;
- active kode_barang uniqueness uses a PostgreSQL partial unique index;
- strict business identity unique index is not created;
- threshold constraints are enforced in PostgreSQL;
- product_versions.product_id uses ON DELETE RESTRICT.

## GAP

No ProductCatalog PostgreSQL repository adapter has been implemented yet.

No ProductCatalog adapter integration tests have been implemented yet.

No query-plan EXPLAIN proof exists yet because adapter queries are not implemented.

No ProductCatalog HTTP/runtime/capability/UI work has started.

## DECISION

Migration-only checkpoint is locally proven.

Do not start Echo HTTP transport, presenters, route registration, capability seed, inventory mutation, UI, or ProductCatalog runtime HTTP slice.

Next implementation step after connector validation is repository adapter skeletons.

## PROOF

Local migration proof:

make db-migrate
make db-status

Observed status:

[APPLIED] 0011_create_product_catalog_tables.up.sql

Aggregate proof before DB apply also passed:

make verify

## NEXT

Validate this checkpoint through GitHub connector.

After connector validation, start ProductCatalog PostgreSQL repository adapter skeletons only.

Do not implement full adapter behavior until skeleton shape is proven against existing ports.

## PROGRESS

ProductCatalog implementation slice 1: 100% closed.

ProductCatalog PostgreSQL persistence slice: migration-only checkpoint locally proven.

Estimated ProductCatalog full transition: 58%.

Business Phase 1: 43%.

Overall transition: 32%.

## CONTEXT WINDOW STATUS

Enough context remains to validate migration checkpoint and start repository adapter skeleton planning.

Forbidden until connector validation:

ProductCatalog PostgreSQL repository adapter behavior
Echo HTTP transport
presenters
route registration
capability seed
inventory stock mutation
UI
ProductCatalog runtime HTTP slice
