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
-->

# Supplier Domain Contract Blueprint

## Status

Accepted locally for contract review.

This is a domain contract slice only. It must not implement Supplier code, migrations, HTTP routes, capability seed, Faktur, inventory mutation, stock movement, audit/outbox, localization, extended filters, or architecture folder cleanup.

## Active Scope

Define the Supplier domain contract as the next POS business-domain dependency after Product API readiness.

Supplier is needed before Faktur because purchase/sales documents need a stable external party reference.

## Domain

```text
supplier
```

## FACT

- Product API readiness for Supplier/Faktur planning is locally closed with proof.
- Product ID is accepted as the downstream Product reference key.
- ProductCatalog must not be accessed by downstream domains through PostgreSQL internals.
- ProductCatalog does not model stock-on-hand in the current Product API readiness slice.
- ServiceCatalog and ProductCatalog prior contracts use blueprint-first domain contract flow.
- Auth/System ADR `0012` output contract centralization is deferred by owner decision and must not block Supplier/Faktur progress.
- No Supplier domain contract is accepted yet.
- No Supplier implementation proof exists yet.

## GAP

- Supplier domain invariants are not accepted yet.
- Supplier duplicate policy is not accepted yet.
- Supplier lifecycle policy is not accepted yet.
- Supplier API contract is not accepted yet.
- Supplier PostgreSQL schema is not accepted yet.
- Supplier capability and permission keys are not accepted yet.
- Supplier relation to Faktur is not accepted yet.
- Supplier relation to ProductCatalog is not accepted yet.
- Laravel Supplier source parity is not proven.
- Supplier implementation, routes, migration, repository, and tests are not active in this slice.

## DECISION

Treat Supplier as master data for external business parties used by future Faktur flows.

Supplier contract decisions:

- Supplier is an independent master-data domain.
- Supplier must not own ProductCatalog data.
- Supplier must not mutate stock.
- Supplier must not create Faktur.
- Supplier may later be referenced by Faktur through `supplier_id`.
- Physical delete is forbidden for normal use.
- Use activate/deactivate lifecycle instead of delete.
- Use server-owned normalization for duplicate checks.
- Do not trust normalized fields from clients.
- Keep API JSON-only and Echo-only.
- Do not implement code in this step.
- Do not add Supplier migration before this contract is accepted.
- Do not start Faktur until Supplier contract is accepted unless owner explicitly changes the order.

## SCOPE-IN

- Supplier domain contract.
- Supplier lifecycle rules.
- Supplier identity and duplicate policy proposal.
- Supplier API route proposal.
- Supplier capability key proposal.
- Supplier authorization proposal.
- Supplier PostgreSQL schema proposal.
- Supplier request/response contract proposal.
- Supplier validation and invariant rules.
- Supplier test matrix proposal.
- Proof requirements for implementation readiness.

## SCOPE-OUT

- Go implementation.
- PostgreSQL migration implementation.
- Capability seed implementation.
- HTTP route registration.
- ProductCatalog implementation changes.
- Faktur implementation.
- Inventory behavior.
- Stock movement.
- Payment/accounting behavior.
- Audit/outbox persistence.
- Runtime localization.
- Extended filters.
- UI implementation.
- Git/GitHub mutation.

## DOMAIN CONTRACT

Domain name:

```text
supplier
```

Source table proposal:

```text
suppliers
```

Aggregate root:

```text
Supplier
```

Owned entity:

```text
Supplier
```

Read models:

```text
SupplierDetail
SupplierListRow
SupplierLookupRow
```

Future reference key for Faktur:

```text
supplier_id
```

Supplier must not depend on ProductCatalog tables directly.

Faktur may later reference both:

```text
supplier_id
product_id
```

Faktur relation rules are out of this Supplier contract slice.

## FIELDS

Supplier:

```text
id
name
name_normalized
phone nullable
email nullable
address nullable
notes nullable
is_active
created_at
updated_at
```

Field notes:

- `name` is the primary display identity.
- `name_normalized` is server-generated.
- `phone`, `email`, `address`, and `notes` are optional.
- `is_active` controls lookup visibility.
- No balance/accounting field belongs to Supplier in this slice.
- No Faktur summary field belongs to Supplier in this slice.

## LIFECYCLE

Statuses:

```text
active
inactive
```

Lifecycle rules:

- Newly created Supplier is active by default.
- Inactive Supplier remains stored for historical references.
- Inactive Supplier should not appear in default lookup results.
- List endpoint may include status filter.
- Physical delete is forbidden for normal use.
- Deactivation must not delete or rewrite future historical Faktur references.

Delete policy:

```text
forbidden
```

Reason:

Supplier may be referenced by future Faktur, purchase records, stock movement records, audit records, or historical documents.

## ALLOWED OPERATIONS

```text
create
update
activate
deactivate
show
list
lookup
```

## FORBIDDEN OPERATIONS

```text
physical_delete
supplier_balance_mutation
supplier_payment_mutation
faktur_create
stock_adjustment_create
stock_adjustment_reverse
inventory_movement_mutation
```

## INVARIANTS

- `id` is required.
- `name` is required after trim.
- `name_normalized` is required.
- `name_normalized` is generated by server normalizer.
- `name_normalized` must be unique among active suppliers.
- `phone` may be null or non-empty after trim.
- `email` may be null or non-empty after trim.
- `address` may be null or non-empty after trim.
- `notes` may be null or non-empty after trim.
- `is_active` must be explicit and default to true.
- Client input must not control `name_normalized`.
- Client input must not control created/updated timestamps.
- Client input must not control active status on create unless a later accepted blueprint allows it.
- Deactivate may accept a nullable reason only if audit/outbox support exists; until then reason may be ignored or deferred.

## NORMALIZATION

Name normalization must:

- Trim leading/trailing whitespace.
- Compact repeated internal whitespace.
- Lowercase consistently.
- Produce the value used by uniqueness checks.

Open proof needed:

```text
Confirm exact Laravel Supplier normalization behavior if Laravel Supplier source is later provided.
```

## DUPLICATE POLICY

Initial duplicate policy proposal:

```text
active supplier name must be unique by normalized name
```

Allowed:

- An inactive Supplier may keep its historical normalized name.
- A new active Supplier may reuse a name only if no active Supplier has that normalized name.

Forbidden:

```text
Two active suppliers with the same normalized name.
```

Owner decision needed before implementation:

```text
Should inactive Supplier names block reuse, or may active Supplier reuse an inactive Supplier name?
```

Default recommendation:

```text
Allow active reuse after deactivation, matching soft-delete style master-data behavior.
```

## POSTGRESQL SCHEMA PROPOSAL

Table:

```text
suppliers
```

Columns:

```sql
id text primary key,
name text not null,
name_normalized text not null,
phone text null,
email text null,
address text null,
notes text null,
is_active boolean not null default true,
created_at timestamptz not null default now(),
updated_at timestamptz not null default now()
```

Indexes:

```sql
create unique index suppliers_active_name_normalized_unique
on suppliers (name_normalized)
where is_active = true;

create index suppliers_active_name_idx
on suppliers (is_active, name_normalized, id);
```

Migration decision:

```text
Do not implement migration until owner accepts this domain contract and duplicate policy.
```

## API CONTRACT PROPOSAL

Base path:

```text
/api/suppliers
```

Routes:

```text
GET    /api/suppliers
POST   /api/suppliers
GET    /api/suppliers/lookup
GET    /api/suppliers/:id
PUT    /api/suppliers/:id
POST   /api/suppliers/:id/activate
POST   /api/suppliers/:id/deactivate
```

Route ordering note:

```text
Register fixed route /lookup before /:id.
```

## QUERY CONTRACT

List query:

```text
q
page
per_page
status = active|inactive|all
```

Defaults:

```text
page = 1
per_page = 10
status = active
```

Limits:

```text
page >= 1
1 <= per_page <= 50
```

Lookup query:

```text
q
limit
active_only
```

Defaults:

```text
limit = 20
active_only = true
```

Limits:

```text
1 <= limit <= 50
```

## COMMAND DTO PROPOSAL

Create command:

```text
name required
phone nullable
email nullable
address nullable
notes nullable
```

Update command:

```text
name required
phone nullable
email nullable
address nullable
notes nullable
```

Activate command:

```text
id required
```

Deactivate command:

```text
id required
reason nullable, deferred until audit/outbox support
```

## RESPONSE DTO PROPOSAL

Detail/list row:

```text
id
name
phone
email
address
notes
is_active
created_at
updated_at
available_operations
```

Lookup row:

```text
id
name
phone
email
```

Response envelope:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

## ERROR CONTRACT PROPOSAL

Use shared error envelope:

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

Suggested public error codes:

```text
invalid_supplier_name
supplier_not_found
duplicate_supplier_name
supplier_inactive
invalid_supplier_status
invalid_request_body
```

## AUTHORIZATION POLICY

Initial permission proposal:

```text
supplier.read
supplier.manage
```

Capability mapping proposal:

```text
supplier.list       -> supplier.read
supplier.lookup     -> supplier.read
supplier.show       -> supplier.read
supplier.create     -> supplier.manage
supplier.update     -> supplier.manage
supplier.activate   -> supplier.manage
supplier.deactivate -> supplier.manage
```

Owner decision needed before implementation:

```text
Use two permissions (supplier.read, supplier.manage) or operation-specific permissions.
```

Default recommendation:

```text
Use two permissions first, matching current ProductCatalog and ServiceCatalog style.
```

## RELATION TO PRODUCTCATALOG

Supplier must not own Product data.

Supplier must not require ProductCatalog at create/update time.

Future Faktur may combine:

```text
supplier_id
product_id
```

This Supplier contract must stay independent.

## RELATION TO FAKTUR

Supplier is a dependency for Faktur.

Faktur is out of scope for this contract.

Do not define invoice numbering, item lines, stock mutation, payment state, or accounting state here.

## TEST MATRIX PROPOSAL

Domain tests:

```text
create supplier trims and normalizes name
create supplier rejects blank name
create supplier accepts optional contact fields
create supplier defaults active
update supplier validates name
deactivate supplier marks inactive
activate supplier marks active
duplicate active normalized name is rejected
```

Usecase tests:

```text
create calls duplicate checker
update checks existence
update checks duplicate name
list forwards query filters
lookup defaults active-only
show returns not found for missing supplier
activate/deactivate return not found for missing supplier
```

PostgreSQL tests:

```text
create persists supplier
update persists supplier
duplicate active normalized name is rejected
inactive supplier name reuse follows accepted duplicate policy
list filters status
lookup filters active-only by default
show by ID returns supplier
```

HTTP tests:

```text
list returns canonical success envelope
lookup returns canonical success envelope
show returns canonical success envelope
create returns canonical success envelope
update returns canonical success envelope
activate returns canonical success envelope
deactivate returns canonical success envelope
invalid request returns shared error envelope
disabled capability rejects before handler
```

Audit/proof:

```bash
go test ./internal/modules/supplier/...
go test ./internal/app/bootstrap/...
bash scripts/audit_hexagonal.sh
bash scripts/audit_route_capabilities.sh
make verify
```

## ACCEPTANCE GATE

This contract is accepted when owner agrees on:

- Domain name.
- Table name.
- Field set.
- Lifecycle.
- Duplicate policy.
- Permission/capability model.
- API route proposal.
- Supplier before Faktur order.

## NEXT

After this contract is accepted, next valid active step:

```text
Supplier implementation slice 1: domain, ports, and usecase contracts only.
```

Do not implement PostgreSQL, HTTP routes, capability seed, or Faktur until the domain/usecase slice is proven.
