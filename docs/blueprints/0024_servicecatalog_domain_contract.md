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

# ServiceCatalog Domain Contract Blueprint

## Status

Accepted domain contract.

## Date

2026-06-08

## Current Implementation Note

This domain contract remains accepted as the ServiceCatalog source contract.

Post-contract implementation status now lives in:

```text
docs/blueprints/0025_servicecatalog_implementation_slice_1.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/archive/handoffs-closed/2026-06-08-servicecatalog-implementation-slice-1.md
```

ServiceCatalog slice 1 is implemented and closed with proof for:

```text
internal/modules/servicecatalog/domain
internal/modules/servicecatalog/ports
internal/modules/servicecatalog/usecase
```

This contract should still be used for lifecycle, API, PostgreSQL, authorization, capability, audit, and transaction decisions for later slices.

## Active Scope

Create the first POS business-domain blueprint/domain contract after capability-control foundation closeout.

## Domain

```text
servicecatalog
```

## FACT

- Capability-control foundation is closed.
- The first POS business-domain step is contract/blueprint work only.
- `servicecatalog` is safer to start before `productcatalog` because product duplicate policy still needs an owner decision.
- Current Laravel evidence for ServiceCatalog is limited to service catalog seed data and normalizer unit test.
- Full ServiceCatalog Laravel source has not been provided.
- Current Go repo has no `servicecatalog` business module implementation proof yet.
- Protected API endpoints must have capability metadata and runtime capability checks before validation/usecase execution.

## GAP

- Full ServiceCatalog Laravel source is not available yet.
- Existing Laravel route/controller/request coverage for ServiceCatalog has not been proven.
- Exact seed row contents for `service_catalog_items` need source proof before writing Go migrations/seeds.
- Audit module implementation is not available yet.
- No PostgreSQL migration has been accepted for `service_catalog_items`.
- ServiceCatalog HTTP handler, PostgreSQL repository, migration, route registration, and capability seed implementation are not active in this contract blueprint.

## DECISION

- Start with `servicecatalog` before `productcatalog`.
- Treat `servicecatalog` as master data.
- Keep physical delete forbidden for normal use.
- Use activate/deactivate lifecycle instead of delete.
- Use server-owned normalization for names.
- Do not trust `normalized_name` from clients.
- Keep API JSON-only and Echo-only.
- Do not implement code in this step.
- Do not add capability seed migration before this contract is accepted.
- Do not start product catalog until product duplicate policy is resolved or explicitly deferred by owner decision.

## SCOPE-IN

- ServiceCatalog domain contract.
- Lifecycle and operation rules.
- API route proposal.
- Capability key proposal.
- Authorization proposal.
- PostgreSQL schema proposal.
- Request/response contract proposal.
- Validation and invariant rules.
- Test matrix proposal.
- Proof requirements for implementation readiness.

## SCOPE-OUT

- Go implementation.
- PostgreSQL migration implementation.
- Capability seed implementation.
- ProductCatalog implementation.
- Inventory behavior.
- Payment, procurement, reporting, and transaction behavior.
- UI implementation.
- Git/GitHub mutation.

## DOMAIN CONTRACT

Domain name:

```text
servicecatalog
```

Source table:

```text
service_catalog_items
```

Aggregate root:

```text
ServiceCatalogItem
```

Owned entity:

```text
ServiceCatalogItem
```

Read models:

- `ServiceCatalogItemDetail`
- `ServiceCatalogItemListRow`
- `ServiceCatalogLookupRow`

## LIFECYCLE

Statuses:

- `active`
- `inactive`

Lifecycle rules:

- Newly created item is active by default.
- Inactive item remains stored for historical references.
- Inactive item should not appear in default lookup results.
- List endpoint may include status filter.
- Physical delete is forbidden for normal use.

Delete policy:

```text
forbidden
```

Reason:

Service catalog items may be referenced by future sales, service transactions, audit records, or historical documents.

## ALLOWED OPERATIONS

- `create`
- `update`
- `activate`
- `deactivate`
- `show`
- `list`
- `lookup`

## FORBIDDEN OPERATIONS

- `physical_delete`

## INVARIANTS

- `id` is required.
- `name` is required after trim.
- `normalized_name` is required.
- `normalized_name` is generated by server normalizer.
- `normalized_name` must be unique.
- `default_price_rupiah` must be greater than zero.
- `is_active` must be explicit and must default to true.
- Client input must not control `normalized_name`.
- Client input must not control created/updated timestamps.

## NORMALIZATION

Name normalization must:

- Trim leading/trailing whitespace.
- Compact repeated internal whitespace.
- Lowercase consistently.
- Produce the value used by uniqueness checks.

Open proof needed:

Confirm exact Laravel normalizer behavior from the ServiceCatalog normalizer unit test before implementation.

## POSTGRESQL SCHEMA PROPOSAL

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

Migration decision:

Do not implement migration until owner accepts this domain contract and seed policy is proven.

## API CONTRACT PROPOSAL

Base path:

```text
/api/service-catalog/items
```

Routes:

```text
GET    /api/service-catalog/items
POST   /api/service-catalog/items
GET    /api/service-catalog/items/:id
PUT    /api/service-catalog/items/:id
POST   /api/service-catalog/items/:id/activate
POST   /api/service-catalog/items/:id/deactivate
GET    /api/service-catalog/items/lookup
```

Route ordering note:

Register fixed route `/lookup` before `/:id` if Echo routing requires it.

## QUERY CONTRACT

List query:

- `q` nullable string
- `page` int default 1 min 1
- `per_page` int default 10 min 1 max 50
- `status` one of `active,inactive,all` default `active`
- `sort_by` one of `name,default_price_rupiah,created_at,updated_at`
- `sort_dir` one of `asc,desc` default `asc`

Lookup query:

- `q` nullable string
- `limit` int default 20 min 1 max 50
- `active_only` bool default true

## COMMAND DTO PROPOSAL

Create command:

- `name` string required
- `default_price_rupiah` int64 required > 0

Update command:

- `name` string required
- `default_price_rupiah` int64 required > 0

Activate command:

- `id` string required

Deactivate command:

- `id` string required
- `reason` nullable string

## RESPONSE DTO PROPOSAL

Detail/list row:

- `id`
- `name`
- `default_price_rupiah`
- `is_active`
- `created_at`
- `updated_at`
- `available_operations`

Lookup row:

- `id`
- `name`
- `default_price_rupiah`

Response envelope:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

## AUTHORIZATION POLICY

Initial permission proposal:

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

Owner decision needed before implementation:

Use two permissions (`read`, `manage`) or operation-specific permissions.

Recommended first implementation:

Use `service_catalog.read` and `service_catalog.manage` to avoid permission sprawl during initial domain migration.

## CAPABILITY KEYS

- `service_catalog.list`
- `service_catalog.lookup`
- `service_catalog.show`
- `service_catalog.create`
- `service_catalog.update`
- `service_catalog.activate`
- `service_catalog.deactivate`

Capability metadata must include:

- `key`
- `domain`
- `operation`
- `method`
- `path`
- `default_state`
- `required_permission`
- `risk_level`
- `audit_required`
- `idempotency_required`
- `owner_package`
- `test_proof`
- `disabled_reason`

## AUDIT POLICY

- `create` should be auditable.
- `update` should be auditable.
- `activate` should be auditable.
- `deactivate` should be auditable.
- First implementation may record audit intent if audit sink is not available yet.
- Do not hide audit behavior inside HTTP handlers.

Audit event proposal:

```text
bounded_context = service_catalog
aggregate_type = service_catalog_item
aggregate_id = item id
event_name = service_catalog_item_created | service_catalog_item_updated | service_catalog_item_activated | service_catalog_item_deactivated
occurred_at
actor_id
actor_role
reason
source_channel
metadata_json
```

## IDEMPOTENCY POLICY

Initial implementation:

```text
not required
```

Reason:

Normal admin master-data writes are not retry-heavy transaction endpoints.

Future decision:

Add idempotency key only if API clients need safe retry semantics.

## TRANSACTION BOUNDARY

Create/update/activate/deactivate must run inside one transaction when persistence is implemented.

Transaction must cover:

- `service_catalog_items` write
- Audit event write if audit sink exists

## PACKAGE PLAN

```text
internal/modules/servicecatalog/
  domain/
  ports/
  usecase/
  transport/http/

internal/platform/postgres/
  service_catalog_repository.go
  service_catalog_queries.go

internal/presentation/http/id/servicecatalog/
```

Rules:

- No Echo import in domain/usecase.
- No SQL in handlers/usecases.
- No UI authority inside domain.
- HTTP layer only parses request and calls usecase.

## TEST MATRIX

Domain tests:

- Normalize service name.
- Reject blank name.
- Reject zero default price.
- Reject negative default price.
- Preserve display name after trim.
- Produce stable normalized name.

Usecase tests:

- Create stores active item.
- Create rejects duplicate normalized name.
- Update changes name and price.
- Update rejects duplicate normalized name.
- Show missing item returns not found.
- Deactivate marks item inactive.
- Activate marks item active.
- Lookup excludes inactive by default.
- List filters active/inactive/all.
- Disabled capability prevents handler/usecase execution.

Repository/integration tests when `DATABASE_URL` is available:

- Unique `normalized_name`.
- Positive price check.
- Active lookup ordering.
- Pagination and limit behavior.
- Transaction rollback on write failure.

HTTP tests:

- List requires authn/authz/capability.
- Show requires authn/authz/capability.
- Create requires manage permission and capability.
- Update requires manage permission and capability.
- Activate/deactivate require manage permission and capability.
- Validation errors use public error envelope.
- Success responses use envelope.

Script/quality proof:

- `go test ./...`
- `make verify`

## IMPLEMENTATION READINESS GATES

Before implementation starts:

- Owner accepts this domain contract.
- ServiceCatalog Laravel normalizer source/test is provided or current evidence is accepted as enough.
- Seed row source is provided if Go seed migration is needed.
- Permission granularity decision is accepted.
- Capability seed list is accepted.
- PostgreSQL schema proposal is accepted.
- Audit behavior for first implementation is accepted.

## NEXT ACTIVE STEP

ServiceCatalog domain contract is accepted.

Next valid active step:

```text
Plan ServiceCatalog implementation slice 1 without coding yet.
```

Implementation remains blocked until an implementation blueprint/patch plan is accepted.

## ACCEPTANCE

ServiceCatalog domain contract accepted on 2026-06-08.

Accepted decisions:

```text
domain = servicecatalog
delete policy = forbidden physical delete
lifecycle = active/inactive
initial permissions = service_catalog.read, service_catalog.manage
capability keys = service_catalog.list, service_catalog.lookup, service_catalog.show, service_catalog.create, service_catalog.update, service_catalog.activate, service_catalog.deactivate
idempotency = not required for first implementation
audit = required by decision, first implementation may record audit intent if audit sink is not available
```

Implementation remains blocked until the implementation slice is planned and accepted.
