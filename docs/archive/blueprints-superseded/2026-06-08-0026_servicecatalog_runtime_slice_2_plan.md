# Archived Blueprint: ServiceCatalog Runtime Slice 2 Blueprint

## Archive Status

Historical only.

This draft was archived on 2026-06-08 because it overlapped with the active `0026` blueprint:

```text
docs/blueprints/0026_servicecatalog_postgres_persistence_slice.md
```

Do not use this file as an active implementation plan.

The active next step is the narrower PostgreSQL persistence slice.

The later ServiceCatalog runtime/capability slice must get a new blueprint number and explicitly own HTTP transport, route registration, permission seeds, capability seeds, route manifest updates, and disabled-capability proof before protected routes are registered.

## Original Draft

## Status

Draft implementation slice plan.

## Date

2026-06-08

## Active Scope

Plan the next ServiceCatalog implementation slice after completed slice 1.

## Previous Completed Slice

```text
docs/blueprints/0025_servicecatalog_implementation_slice_1.md
```

## FACT

- ServiceCatalog slice 1 is implemented and proven.
- ServiceCatalog domain package exists.
- ServiceCatalog ports package exists.
- ServiceCatalog usecase package exists.
- ServiceCatalog domain and usecase unit tests pass.
- `make verify` passes after slice 1.
- HTTP transport is not implemented.
- PostgreSQL adapter is not implemented.
- PostgreSQL migrations are not implemented.
- Route registration is not implemented.
- Capability seeds are not implemented.
- ProductCatalog remains out of scope.
- Inventory remains out of scope.

## DECISION

Slice 2 should implement ServiceCatalog runtime integration in one narrow vertical slice.

The implementation should include:

- PostgreSQL migration for `service_catalog_items`.
- PostgreSQL repository adapter.
- HTTP request/response DTO mapping.
- Echo HTTP handlers.
- Route registration.
- Capability metadata seed entries.
- Route capability manifest update.
- Tests and proof.

## SCOPE-IN

- `migrations/*service_catalog_items*`
- `internal/platform/postgres/service_catalog_*`
- `internal/presentation/http/id/servicecatalog/*`
- `internal/modules/servicecatalog/transport/http/*`
- Route registration for ServiceCatalog endpoints
- Capability seed migration for ServiceCatalog endpoints
- `scripts/config/route_capabilities.tsv`
- Tests for repository, transport, bootstrap route wiring, and capability coverage
- Docs evidence and handoff updates

## SCOPE-OUT

- ProductCatalog
- Inventory
- Sales/transactions
- Payment
- Procurement
- Reporting
- UI
- Broad Laravel parity translation
- Audit sink implementation beyond explicit audit intent if required

## DOMAIN CONTRACT SOURCE

```text
docs/blueprints/0024_servicecatalog_domain_contract.md
```

## IMPLEMENTATION SOURCE

```text
docs/blueprints/0025_servicecatalog_implementation_slice_1.md
```

## DATABASE TARGET

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

## ROUTES

Base path:

```text
/api/service-catalog/items
```

Routes:

```text
GET    /api/service-catalog/items
POST   /api/service-catalog/items
GET    /api/service-catalog/items/lookup
GET    /api/service-catalog/items/:id
PUT    /api/service-catalog/items/:id
POST   /api/service-catalog/items/:id/activate
POST   /api/service-catalog/items/:id/deactivate
```

Route ordering rule:

Register `/lookup` before `/:id`.

## CAPABILITY KEYS

- `service_catalog.list`
- `service_catalog.create`
- `service_catalog.lookup`
- `service_catalog.show`
- `service_catalog.update`
- `service_catalog.activate`
- `service_catalog.deactivate`

## PERMISSIONS

- `service_catalog.read`
- `service_catalog.manage`

Mapping:

```text
service_catalog.list       -> service_catalog.read
service_catalog.lookup     -> service_catalog.read
service_catalog.show       -> service_catalog.read
service_catalog.create     -> service_catalog.manage
service_catalog.update     -> service_catalog.manage
service_catalog.activate   -> service_catalog.manage
service_catalog.deactivate -> service_catalog.manage
```

## REQUIRED RUNTIME ORDER

Every protected ServiceCatalog endpoint must pass:

```text
authn -> authz -> capability check -> request validation -> usecase execution
```

Capability disabled proof must show the protected endpoint does not reach handler/usecase execution.

## REQUEST CONTRACTS

Create request:

```json
{
  "name": "Potong Rambut",
  "default_price_rupiah": 10000
}
```

Update request:

```json
{
  "name": "Cuci Motor",
  "default_price_rupiah": 15000
}
```

List query:

- `q` nullable string
- `page` int default 1 min 1
- `per_page` int default 10 min 1 max 50
- `status` active|inactive|all default active
- `sort_by` name|default_price_rupiah|created_at|updated_at default name
- `sort_dir` asc|desc default asc

Lookup query:

- `q` nullable string
- `limit` int default 20 min 1 max 50
- `active_only` bool default true

## RESPONSE CONTRACTS

Success envelope:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

Error envelope:

```json
{
  "success": false,
  "error": {
    "code": "validation_failed",
    "message": "Validation failed",
    "fields": {}
  },
  "meta": {}
}
```

## TEST MATRIX

Migration/database tests:

- Migration creates `service_catalog_items`.
- Unique normalized name is enforced.
- Default price must be greater than zero.
- Active/name index exists or is proven by schema inspection.

Repository tests:

- Create stores item.
- Find by ID returns item.
- Find by normalized name returns item.
- Update changes item.
- List filters active/inactive/all.
- Lookup excludes inactive by default.
- Lookup enforces limit.
- Set active toggles lifecycle state.

HTTP tests:

- Create validates request body.
- Create returns 201.
- Duplicate create returns 409.
- List returns envelope.
- Lookup returns active items by default.
- Show missing item returns 404.
- Update validates request body.
- Activate returns active item.
- Deactivate returns inactive item.
- Route `/lookup` is not captured by `/:id`.

Capability tests:

- Each ServiceCatalog protected route has route capability manifest coverage.
- Read endpoints require `service_catalog.read`.
- Mutation endpoints require `service_catalog.manage`.
- Disabled capability returns 403.
- Disabled mutation endpoint does not reach usecase execution.

Quality proof:

- `go test ./internal/modules/servicecatalog/...`
- Repository/integration test command if database test convention exists
- Focused HTTP/bootstrap tests
- `make verify`

## ACCEPTANCE CRITERIA

Slice 2 can be accepted only after owner confirms:

- Table schema
- Capability key list
- Permission mapping
- Route list and ordering
- Whether audit sink remains deferred or audit intent must be recorded
- Proof commands

## NEXT ACTIVE STEP

Review and accept this blueprint before implementation.

Do not implement ServiceCatalog runtime slice 2 until this blueprint is accepted.

## PROGRESS IMPACT IF IMPLEMENTED

Expected after implementation proof:

- ServiceCatalog runtime integration: partial/complete depending repository + HTTP coverage.
- Business Phase 1 implementation: may increase from 15% to around 30%.
- Overall Laravel-to-Go transition: may increase from 22% to around 25%.
