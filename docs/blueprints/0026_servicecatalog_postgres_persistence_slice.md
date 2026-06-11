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

# ServiceCatalog PostgreSQL Persistence Slice Blueprint

## Status

Closed / implemented with proof.

## Date

2026-06-08

## Active Scope

Plan ServiceCatalog PostgreSQL persistence after ServiceCatalog slice 1 domain/usecase proof.

## Source Contract

```text
docs/blueprints/0024_servicecatalog_domain_contract.md
```

## Prior Slice

```text
docs/blueprints/0025_servicecatalog_implementation_slice_1.md
```

## FACT

- ServiceCatalog domain contract is accepted.
- ServiceCatalog slice 1 is implemented and closed with proof.
- ServiceCatalog domain, ports, usecase contracts, and unit tests exist.
- ServiceCatalog repository port already defines create, update, find by ID, find by normalized name, list, lookup, and set active behavior.
- ServiceCatalog HTTP transport is not implemented.
- ServiceCatalog PostgreSQL adapter is implemented and closed with proof.
- ServiceCatalog PostgreSQL migration is implemented and closed with proof.
- ServiceCatalog route registration is not implemented.
- ServiceCatalog capability seed migration is not implemented.
- Superseded broader runtime draft `docs/blueprints/0026_servicecatalog_runtime_slice_2_plan.md` is archived as historical-only context.

## DECISION

Slice 2 implemented only ServiceCatalog PostgreSQL persistence.

Do not implement HTTP transport in this slice.

Do not register routes in this slice.

Do not add ServiceCatalog capability seed rows in this slice.

Do not add ProductCatalog.

Do not add UI behavior.

## IMPLEMENTED SCOPE

- PostgreSQL migration for `service_catalog_items`.
- PostgreSQL repository adapter for the existing ServiceCatalog repository port.
- Repository integration tests for create, find, update, lifecycle, list, and lookup behavior.
- Transaction context support through the existing PostgreSQL transaction context pattern.

## SCOPE-IN

- PostgreSQL migration for `service_catalog_items`.
- PostgreSQL repository adapter for the existing ServiceCatalog repository port.
- Repository/integration tests when `DATABASE_URL` is available.
- Transaction boundary support for write operations if required by existing platform transaction pattern.
- Persistence proof for uniqueness, positive price, lifecycle state, list, lookup, and pagination/limit behavior.

## SCOPE-OUT

- Echo HTTP handlers.
- Request/response DTO presenters.
- Route registration.
- Capability seed migrations.
- Authorization middleware wiring.
- Audit sink implementation.
- ProductCatalog.
- Inventory.
- UI.

## IMPLEMENTED FILES

```text
migrations/0009_create_service_catalog_items.up.sql
migrations/0009_create_service_catalog_items.down.sql
internal/modules/servicecatalog/domain/restore_service_catalog_item.go
internal/platform/postgres/service_catalog_repository.go
internal/platform/postgres/service_catalog_repository_query.go
internal/platform/postgres/service_catalog_repository_row.go
internal/platform/postgres/service_catalog_repository_string.go
internal/platform/postgres/service_catalog_repository_read.go
internal/platform/postgres/service_catalog_repository_write.go
internal/platform/postgres/service_catalog_repository_list.go
internal/platform/postgres/service_catalog_repository_lookup.go
internal/platform/postgres/time.go
internal/platform/postgres/service_catalog_repository_create_integration_test.go
internal/platform/postgres/service_catalog_repository_integration_helpers_test.go
internal/platform/postgres/service_catalog_repository_query_integration_test.go
internal/platform/postgres/service_catalog_repository_update_integration_test.go
```

## TARGET FILES

Expected new or changed files:

```text
migrations/0009_create_service_catalog_items.up.sql
migrations/0009_create_service_catalog_items.down.sql
internal/platform/postgres/service_catalog_repository.go
internal/platform/postgres/service_catalog_repository_query.go
internal/platform/postgres/service_catalog_repository_row.go
internal/platform/postgres/service_catalog_repository_test.go
```

If file-size audit requires smaller files, split repository tests/helpers without changing package boundaries.

## POSTGRESQL SCHEMA

Table:

```text
service_catalog_items
```

Columns:

```text
id text primary key,
name text not null,
normalized_name text not null,
default_price_rupiah bigint not null check (default_price_rupiah > 0),
is_active boolean not null default true,
created_at timestamptz not null default now(),
updated_at timestamptz not null default now()
```

Indexes:

```sql
create unique index service_catalog_items_normalized_name_unique
on service_catalog_items (normalized_name);

create index service_catalog_items_active_name_idx
on service_catalog_items (is_active, normalized_name);
```

## REPOSITORY BEHAVIOR

The PostgreSQL adapter must implement:

- `Create`
- `Update`
- `FindByID`
- `FindByNormalizedName`
- `List`
- `Lookup`
- `SetActive`

Rules:

- Repository accepts domain objects and returns domain objects.
- Repository must not perform HTTP parsing.
- Repository must not own business validation that already belongs in domain/usecase.
- Repository must preserve server-generated normalized name from the domain object.
- Repository must return not-found as `(zero, false, nil)` where the port uses a boolean found result.
- Duplicate normalized name must surface as an error that usecase can map to duplicate behavior.
- List defaults remain owned by usecase/port input, not SQL magic.
- Lookup should support active-only filtering.

## TRANSACTION POLICY

Create, update, activate, and deactivate persistence operations should be usable inside a transaction when HTTP/write orchestration is added.

If the existing PostgreSQL transactor can be reused cleanly, this slice may wire repository methods to use the transaction context pattern already used by auth/capability.

If no clean transaction pattern exists for this repository yet, document the gap and keep transaction wiring deferred to the HTTP/write slice.

## TEST MATRIX

Repository/integration tests:

- Create stores a service catalog item.
- Create rejects duplicate `normalized_name`.
- Create rejects non-positive `default_price_rupiah` through DB constraint.
- Find by ID returns found item.
- Find by normalized name returns found item.
- Update changes name, normalized name, price, and updated timestamp.
- `SetActive false` marks item inactive.
- `SetActive true` marks item active.
- List filters active, inactive, and all.
- Lookup excludes inactive by default.
- Lookup respects limit.
- Lookup orders by normalized name or stable deterministic ordering.
- Down migration drops table and indexes created by this slice.

## PROOF REQUIRED

Focused proof:

```text
go test ./internal/platform/postgres/... -run ServiceCatalog
```

Full proof:

```text
make verify
```

Migration proof if local DB is available:

```text
make db-up
make db-status
```

Expected migration status must include:

```text
0009_create_service_catalog_items.up.sql applied
```

## PROOF COLLECTED

```text
Focused compile/repository package proof:

go test ./internal/modules/servicecatalog/... ./internal/platform/postgres/...
ok  	pos-go/internal/modules/servicecatalog/domain
?   	pos-go/internal/modules/servicecatalog/ports	[no test files]
ok  	pos-go/internal/modules/servicecatalog/usecase
?   	pos-go/internal/platform/postgres	[no test files]

Focused integration proof:

go test -tags=integration ./internal/platform/postgres/... -run ServiceCatalog -count=1
ok  	pos-go/internal/platform/postgres	0.006s

Full proof:

make verify
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed

Security proof:

Gosec  : dev
Files  : 122
Lines  : 5081
Nosec  : 0
Issues : 0
```

## ACCEPTANCE GATE

This blueprint is accepted with owner confirmation:

- Implement ServiceCatalog PostgreSQL persistence slice only.
- No HTTP transport.
- No route registration.
- No capability seed migration.
- No ProductCatalog.

## NEXT ACTIVE STEP

ServiceCatalog PostgreSQL persistence slice is closed.

Next valid active step:

Plan ServiceCatalog HTTP transport, route registration, and capability seed slice.

Do not implement HTTP transport, route registration, or capability seeds until a later accepted blueprint defines exact files, route/capability mapping, authorization behavior, response envelope behavior, and proof commands.

The later runtime/capability blueprint must explicitly own ServiceCatalog permission seed rows, capability seed rows, route capability manifest updates, and disabled-capability proof before any protected ServiceCatalog HTTP route is registered.
