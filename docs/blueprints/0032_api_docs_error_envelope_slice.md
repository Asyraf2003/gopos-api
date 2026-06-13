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

# API Docs Error Envelope Slice Blueprint

## Status

Draft implementation slice plan.

## Date

2026-06-13

## Active Scope

Document the existing ProductCatalog HTTP API contract and standardize public HTTP error responses using the existing Echo HTTP contract.

## FACT

ProductCatalog protected runtime is complete for eight backend API routes.

ProductCatalog success responses use:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

The generic Echo HTTP contract already requires a consistent envelope family for success and error responses.

ProductCatalog mapped errors currently return Echo HTTPError values and therefore do not yet guarantee the standard public error envelope shape.

ProductCatalog-specific developer API documentation is not yet present.

## GAP

No ProductCatalog-specific API doc exists yet at:

```text
docs/api/product_catalog.md
```

No accepted blueprint exists yet for the API docs and error envelope implementation slice.

No shared HTTP error handler has been proven for converting Echo HTTPError responses into:

```json
{
  "success": false,
  "error": {
    "code": "product_not_found",
    "message": "product not found"
  },
  "meta": {}
}
```

## DECISION

Implement a small API docs and error envelope slice.

Preserve the existing success envelope.

Standardize error responses through the HTTP stack instead of hand-building inconsistent error JSON inside individual handlers.

Use stable machine-readable error codes.

Keep ProductCatalog domain, PostgreSQL persistence, inventory mutation, UI, and audit sink behavior unchanged.

## SCOPE-IN

Create ProductCatalog developer API documentation.

Document request and response examples for all eight ProductCatalog routes.

Define ProductCatalog error codes and mapped HTTP statuses.

Add or wire a shared Echo HTTP error handler if the current Echo default handler cannot produce the standard envelope.

Update ProductCatalog error mapping so known errors carry stable public codes.

Add focused ProductCatalog tests proving the error envelope shape.

Add at least one protected-route test proving middleware errors use the same envelope family.

## SCOPE-OUT

Inventory stock mutation.

Stock adjustment create/reverse.

ProductCatalog UI.

Runtime language switch.

OpenAPI generator.

Broad audit sink implementation.

ProductCatalog persistence changes.

Route/capability seed changes unless a bug is found.

## ERROR CONTRACT

Standard error response:

```json
{
  "success": false,
  "error": {
    "code": "product_not_found",
    "message": "product not found"
  },
  "meta": {}
}
```

Validation-style errors may include fields later, but this slice does not need to implement field-level validation details unless already available.

## INITIAL PRODUCTCATALOG ERROR CODE MAP

```text
404 product_not_found                  product not found
409 product_code_already_exists        product code already exists
409 product_identity_already_exists    product identity already exists
400 product_validation_failed          domain validation message
400 invalid_request_body               invalid request body
400 invalid_query_parameter            invalid query parameter message
500 product_catalog_request_failed     product catalog request failed
```

## TARGET FILES

Expected new or changed files:

```text
docs/blueprints/0032_api_docs_error_envelope_slice.md
docs/api/product_catalog.md
internal/transport/http/response/error.go
internal/transport/http/response/error_handler.go
internal/transport/http/response/error_handler_test.go
internal/modules/productcatalog/transport/http/product_catalog_handler_error.go
internal/modules/productcatalog/transport/http/product_catalog_handler_request.go
internal/modules/productcatalog/transport/http/product_catalog_handler_*_test.go
internal/transport/http/middleware/*_test.go
internal/app/bootstrap/app.go
```

Exact file names may be adjusted to match existing package layout and file-size audit rules.

## TEST MATRIX

ProductCatalog error envelope tests:

```text
not found returns 404 with success=false, error.code=product_not_found, meta={}
duplicate code returns 409 with success=false, error.code=product_code_already_exists
duplicate identity returns 409 with success=false, error.code=product_identity_already_exists
invalid body returns 400 with success=false, error.code=invalid_request_body
invalid list status returns 400 with success=false, error.code=invalid_query_parameter
```

Protected-route envelope proof:

```text
disabled capability returns 403 with success=false, error.code=capability_disabled
```

## PROOF REQUIRED

```bash
go test ./internal/transport/http/...
go test ./internal/modules/productcatalog/transport/http/... ./internal/presentation/http/id/productcatalog/...
go test ./internal/app/bootstrap/... ./internal/transport/http/middleware/...
bash scripts/audit_route_capabilities.sh
make verify
```

## NEXT

After this blueprint is accepted, implement the shared HTTP error envelope first, then update ProductCatalog mappings and docs.
