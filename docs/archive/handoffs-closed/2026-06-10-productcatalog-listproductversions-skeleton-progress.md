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

# ProductCatalog ListProductVersions Skeleton Progress Handoff

## Date

2026-06-10

## Active Scope

ProductCatalog implementation slice 1.

Blueprint:

```text
docs/blueprints/0029_productcatalog_implementation_slice_1.md
```

Scope remains limited to:

internal/modules/productcatalog/domain
internal/modules/productcatalog/ports
internal/modules/productcatalog/usecase

## FACT

ListProductVersions contract and constructor/skeleton are locally implemented.

Implemented files:

internal/modules/productcatalog/usecase/list_product_versions_contract.go
internal/modules/productcatalog/usecase/list_product_versions.go

Implemented:

ListProductVersionsQuery
ListProductVersionsResult
ListProductVersionItem
ListProductVersions usecase type with ports.ProductVersionRepository dependency
NewListProductVersions constructor

Focused proof passed:

go test ./internal/modules/productcatalog/...

Line-count checkpoint passed:

wc -l internal/modules/productcatalog/usecase/list_product_versions*.go

## GAP

ListProductVersions behavior is not implemented yet.

ListProductVersions behavior tests are not added yet.

ProductCatalog PostgreSQL adapter, Echo HTTP transport, presenters, route registration, capability seed, inventory stock mutation, UI, and runtime HTTP slice are still not implemented and remain out of scope for ProductCatalog slice 1.

## DECISION

Stop ListProductVersions work after contract and constructor/skeleton checkpoint.

Do not start Execute behavior until this skeleton checkpoint is connector-validated.

Do not start PostgreSQL adapter, migrations, Echo HTTP transport, presenters, route registration, capability seed, inventory stock mutation, UI, or ProductCatalog runtime HTTP slice.

## PROOF

Focused ProductCatalog proof passed after ListProductVersions skeleton creation.

Progress ledger was updated after focused proof:

Business Phase 1: 41%
Overall Laravel-to-Go transition: 31%
ListProductVersions contract and constructor/skeleton are locally proven; connector validation pending.

## NEXT

Execution channel: owner/local terminal.

Next valid implementation step after connector validation:

Add ListProductVersions behavior.

Do not start ProductCatalog PostgreSQL, Echo/runtime, migrations, capability seed, inventory mutation, UI, or ProductCatalog runtime HTTP slice.

## PROGRESS

ProductCatalog domain: 100%.

ProductCatalog ports: 96%.

CreateProduct usecase behavior: 97%.

UpdateProduct usecase behavior: 100% locally proven.

SoftDeleteProduct usecase behavior: 100% locally proven and connector-validated.

RestoreProduct usecase behavior: 100% locally proven and connector-validated.

GetProductDetail usecase behavior: 100% locally proven and connector-validated.

ListProducts usecase behavior: 100% locally proven and connector-visible.

LookupProducts usecase behavior: 100% locally proven and connector-validated.

ListProductVersions contract/skeleton: 100% locally proven.

Business Phase 1: 41%.

Overall transition: 31%.

## CONTEXT WINDOW STATUS

Enough context remains to continue ProductCatalog slice 1 into ListProductVersions behavior only after connector validation of this skeleton checkpoint.

Forbidden scope remains out:

PostgreSQL adapter
migrations
Echo HTTP transport
presenters
route registration
capability seed
inventory stock mutation
UI
ProductCatalog runtime HTTP slice
