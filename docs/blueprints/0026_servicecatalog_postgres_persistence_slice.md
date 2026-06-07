# ServiceCatalog PostgreSQL Persistence Slice Blueprint

## Status

Proposed implementation slice.

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
- ServiceCatalog PostgreSQL adapter is not implemented.
- ServiceCatalog PostgreSQL migration is not implemented.
- ServiceCatalog route registration is not implemented.
- ServiceCatalog capability seed migration is not implemented.

## DECISION

Slice 2 should implement only ServiceCatalog PostgreSQL persistence.

Do not implement HTTP transport in this slice.

Do not register routes in this slice.

Do not add ServiceCatalog capability seed rows in this slice.

Do not add ProductCatalog.

Do not add UI behavior.

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

## ACCEPTANCE GATE

This blueprint is accepted only when owner confirms:

- Implement ServiceCatalog PostgreSQL persistence slice only.
- No HTTP transport.
- No route registration.
- No capability seed migration.
- No ProductCatalog.

## NEXT ACTIVE STEP

After this blueprint is accepted:

Implement ServiceCatalog PostgreSQL persistence slice.

Do not implement HTTP transport, route registration, or capability seeds until a later accepted blueprint.
