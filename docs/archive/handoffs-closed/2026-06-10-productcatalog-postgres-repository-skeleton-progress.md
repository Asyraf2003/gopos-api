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

# ProductCatalog PostgreSQL Repository Skeleton Progress Handoff

## Date

2026-06-10

## Active Scope

ProductCatalog PostgreSQL repository adapter skeletons.

## FACT

ProductCatalog PostgreSQL migration-only checkpoint is locally proven and remote-visible.

Repository adapter skeletons were added under:

internal/platform/postgres

Skeleton files:

internal/platform/postgres/product_repository.go
internal/platform/postgres/product_repository_query.go
internal/platform/postgres/product_repository_write.go
internal/platform/postgres/product_repository_read.go
internal/platform/postgres/product_repository_list.go
internal/platform/postgres/product_repository_lookup.go
internal/platform/postgres/product_repository_duplicate.go
internal/platform/postgres/product_version_repository.go

The skeleton provides:

- NewProductRepository constructor.
- pgxpool-backed ProductRepository struct.
- TxFromContext-compatible query helpers.
- compile-time assertions for ProductRepository, ProductReader, ProductVersionRepository, and ProductDuplicateChecker ports.
- method stubs that return errProductRepositoryNotImplemented.

## GAP

No ProductCatalog PostgreSQL repository behavior has been implemented yet.

No ProductCatalog repository integration tests have been implemented yet.

No query-plan EXPLAIN proof exists yet because SQL behavior is still not implemented.

No ProductCatalog HTTP/runtime/capability/UI work has started.

## DECISION

Stop at repository adapter skeletons.

Do not implement CRUD SQL, list, lookup, duplicate guard, version append/list, or query-plan proof in this step.

Next step after connector validation is the first behavior slice: ProductRepository create/find/update behavior with focused PostgreSQL tests.

## PROOF

Local proof required:

go test ./internal/platform/postgres/...
go test ./internal/modules/productcatalog/...
make verify

## NEXT

Validate repository skeletons through GitHub connector.

After connector validation, implement ProductRepository create/find/update behavior only.

## PROGRESS

ProductCatalog implementation slice 1: 100% closed.

ProductCatalog PostgreSQL persistence slice: migration-only checkpoint locally proven; repository skeletons locally added.

Estimated ProductCatalog full transition: 59%.

Business Phase 1: 44%.

Overall transition: 32%.

## CONTEXT WINDOW STATUS

Enough context remains to validate repository skeletons and start first repository behavior slice.

Forbidden until connector validation:

ProductCatalog PostgreSQL repository behavior
Echo HTTP transport
presenters
route registration
capability seed
inventory stock mutation
UI
ProductCatalog runtime HTTP slice
