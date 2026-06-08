# ServiceCatalog Runtime Capability Slice Blueprint

## Status

Proposed implementation slice.

## Date

2026-06-08

## Active Scope

Plan ServiceCatalog HTTP runtime, route registration, response presenters, permission seeds, capability seeds, route capability manifest coverage, and disabled-capability proof after ServiceCatalog PostgreSQL persistence closeout.

## Source Contracts

```text
docs/blueprints/0024_servicecatalog_domain_contract.md
docs/blueprints/0025_servicecatalog_implementation_slice_1.md
docs/blueprints/0026_servicecatalog_postgres_persistence_slice.md
```

## FACT

- ServiceCatalog domain contract is accepted.
- ServiceCatalog domain, ports, and usecases are implemented with proof.
- ServiceCatalog PostgreSQL migration and repository adapter are implemented with proof.
- ServiceCatalog HTTP transport is not implemented.
- ServiceCatalog request/response presenters are not implemented.
- ServiceCatalog route registration is not implemented.
- ServiceCatalog permission seeds are not implemented.
- ServiceCatalog capability seeds are not implemented.
- Route capability manifest does not yet include ServiceCatalog protected routes.
- Route capability audit script currently has explicit seed/source file lists that must be updated when new protected route seed/source files are added.

## DECISION

This slice implements ServiceCatalog protected runtime HTTP integration.

This slice must include route capability enforcement and proof.

Do not implement ProductCatalog.

Do not implement UI.

Do not implement inventory behavior.

Do not implement broad audit sink behavior beyond capability metadata and proof.

Do not change ServiceCatalog persistence behavior unless required to wire existing usecases.

## SCOPE-IN

- ServiceCatalog HTTP transport.
- Request parsing for list, lookup, show, create, update, activate, and deactivate.
- Response presenter for Indonesian/public HTTP JSON contracts.
- Bootstrap route registration.
- Permission seed migration for:
  - `service_catalog.read`
  - `service_catalog.manage`
- Role permission assignment:
  - admin gets `service_catalog.read`
  - admin gets `service_catalog.manage`
  - cashier gets `service_catalog.read`
  - base gets no ServiceCatalog permission in this slice.
- Capability seed migration for all protected ServiceCatalog routes.
- Route capability manifest update.
- Route capability audit script update so the new seed migration and ServiceCatalog handler source are checked.
- HTTP handler tests for request parsing, validation, envelope shape, not-found mapping, and handler/usecase boundary.
- Runtime disabled-capability proof through route-level test or bootstrap test.

## SCOPE-OUT

- ProductCatalog.
- Inventory.
- UI.
- Public unauthenticated ServiceCatalog routes.
- Physical delete.
- Audit event sink implementation.
- Idempotency-key implementation.
- Pagination total-count metadata unless explicitly required by existing API contract.
- Broad ADR 0012 closeout beyond proof for this slice.

## ROUTE CONTRACT

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

## ROUTE AND CAPABILITY MAPPING

```text
GET    /api/service-catalog/items                service_catalog.list       service_catalog.read
POST   /api/service-catalog/items                service_catalog.create     service_catalog.manage
GET    /api/service-catalog/items/lookup         service_catalog.lookup     service_catalog.read
GET    /api/service-catalog/items/:id            service_catalog.show       service_catalog.read
PUT    /api/service-catalog/items/:id            service_catalog.update     service_catalog.manage
POST   /api/service-catalog/items/:id/activate   service_catalog.activate   service_catalog.manage
POST   /api/service-catalog/items/:id/deactivate service_catalog.deactivate service_catalog.manage
```

## AUTHORIZATION POLICY

Initial permission keys:

```text
service_catalog.read
service_catalog.manage
```

Initial role assignment:

```text
base    -> no ServiceCatalog permission
cashier -> service_catalog.read
admin   -> service_catalog.read, service_catalog.manage
```

Reason:

Cashier may need read/lookup access for future sales flow.
Admin owns master-data management.
Base role should remain minimal.

## CAPABILITY SEED POLICY

Seed all ServiceCatalog route capabilities into `api_capabilities`.

Default enabled:

```text
true
```

Risk levels:

```text
service_catalog.list       low
service_catalog.lookup     low
service_catalog.show       low
service_catalog.create     medium
service_catalog.update     medium
service_catalog.activate   medium
service_catalog.deactivate medium
```

Audit required:

```text
read routes: false
write/lifecycle routes: true
```

Idempotency required:

```text
false
```

Owner package:

```text
internal/modules/servicecatalog/transport/http
```

## REQUEST CONTRACT

List query:

```text
q
page
per_page
status = active|inactive|all
```

Lookup query:

```text
q
limit
active_only
```

Create JSON body:

```json
{
  "name": "Express Wash",
  "default_price_rupiah": 15000
}
```

Update JSON body:

```json
{
  "name": "Express Wash",
  "default_price_rupiah": 15000
}
```

Deactivate JSON body:

```json
{
  "reason": "optional reason"
}
```

Activate body:

```text
empty body
```

## RESPONSE CONTRACT

Success response envelope:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

List response:

```json
{
  "success": true,
  "data": [
    {
      "id": "item-id",
      "name": "Express Wash",
      "normalized_name": "express wash",
      "default_price_rupiah": 15000,
      "is_active": true,
      "status": "active",
      "created_at": "RFC3339",
      "updated_at": "RFC3339",
      "available_operations": []
    }
  ],
  "meta": {}
}
```

Lookup response:

```json
{
  "success": true,
  "data": [
    {
      "id": "item-id",
      "name": "Express Wash",
      "default_price_rupiah": 15000
    }
  ],
  "meta": {}
}
```

Error responses must use the existing public error envelope behavior available in the HTTP stack.

## TARGET FILES

Expected new or changed files:

```text
internal/modules/servicecatalog/transport/http/service_catalog_handler.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_read.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_write.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_request.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_response.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_error.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_test.go
internal/presentation/http/id/servicecatalog/service_catalog_item.go
internal/presentation/http/id/servicecatalog/service_catalog_lookup.go
internal/app/bootstrap/app.go
migrations/0010_seed_service_catalog_permissions_capabilities.up.sql
migrations/0010_seed_service_catalog_permissions_capabilities.down.sql
scripts/config/route_capabilities.tsv
scripts/audit_route_capabilities.sh
```

If file-size audit requires smaller files, split handler tests/helpers without changing package boundaries.

## BOOTSTRAP WIRING PLAN

Bootstrap should wire:

```text
postgres.NewServiceCatalogRepository(pool)
servicecatalog usecases
servicecatalog HTTP handler
```

Protected groups should preserve order:

```text
authn -> authz permission -> runtime capability check -> handler
```

Because different ServiceCatalog routes use different permissions/capabilities, bootstrap may use separate groups or per-route registration helpers to keep exact capability mapping obvious.

## ROUTE CAPABILITY MANIFEST PLAN

Add rows to `scripts/config/route_capabilities.tsv`:

```tsv
GET	/api/service-catalog/items	service_catalog.list	service_catalog.read	exact	group.GET("/items"
POST	/api/service-catalog/items	service_catalog.create	service_catalog.manage	exact	group.POST("/items"
GET	/api/service-catalog/items/lookup	service_catalog.lookup	service_catalog.read	exact	group.GET("/items/lookup"
GET	/api/service-catalog/items/:id	service_catalog.show	service_catalog.read	exact	group.GET("/items/:id"
PUT	/api/service-catalog/items/:id	service_catalog.update	service_catalog.manage	exact	group.PUT("/items/:id"
POST	/api/service-catalog/items/:id/activate	service_catalog.activate	service_catalog.manage	exact	group.POST("/items/:id/activate"
POST	/api/service-catalog/items/:id/deactivate	service_catalog.deactivate	service_catalog.manage	exact	group.POST("/items/:id/deactivate"
```

Update `scripts/audit_route_capabilities.sh` so it checks:

```text
migrations/0010_seed_service_catalog_permissions_capabilities.up.sql
internal/modules/servicecatalog/transport/http/service_catalog_handler.go
```

and any split handler files that contain route source patterns.

## TEST MATRIX

Handler tests:

- List parses query and calls usecase.
- Lookup parses query and calls usecase.
- Show requires ID and calls usecase.
- Create rejects invalid body.
- Create maps command fields correctly.
- Update requires ID and valid body.
- Activate requires ID.
- Deactivate requires ID and optional reason.
- Not-found usecase error maps to 404.
- Duplicate normalized name usecase error maps to conflict or accepted project error mapping.
- Success responses use envelope.
- Handler package must not import PostgreSQL adapter.

Bootstrap/route tests:

- ServiceCatalog routes are registered only behind authn/authz/capability checks.
- Disabled ServiceCatalog capability returns 403 before handler/usecase execution.
- Route capability manifest audit passes with ServiceCatalog rows.

Migration/seed tests or proof:

- `service_catalog.read` permission exists.
- `service_catalog.manage` permission exists.
- cashier has `service_catalog.read`.
- admin has `service_catalog.read`.
- admin has `service_catalog.manage`.
- `api_capabilities` contains all seven ServiceCatalog capability keys.

## PROOF REQUIRED

Focused HTTP proof:

```bash
go test ./internal/modules/servicecatalog/transport/http/...
```

Focused bootstrap/capability proof:

```bash
go test ./internal/app/bootstrap/... ./internal/transport/http/middleware/...
```

Route capability proof:

```bash
bash scripts/audit_route_capabilities.sh
```

Full proof:

```bash
make verify
```

Runtime DB proof if local DB is available:

```bash
make db-up
make db-status
```

Expected applied migration:

```text
0010_seed_service_catalog_permissions_capabilities.up.sql
```

Optional SQL proof:

```sql
select key from permissions where key in ('service_catalog.read', 'service_catalog.manage');
select key from api_capabilities where key like 'service_catalog.%' order by key;
```

## ACCEPTANCE GATE

This blueprint is accepted only when owner confirms:

- Implement ServiceCatalog runtime/capability slice only.
- No ProductCatalog.
- No UI.
- No physical delete.
- No broad audit sink.

## NEXT ACTIVE STEP

After this blueprint is accepted:

Implement ServiceCatalog runtime/capability slice.

Do not implement ProductCatalog or UI in this slice.
