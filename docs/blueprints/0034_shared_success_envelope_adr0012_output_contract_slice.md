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

# Shared Success Envelope And ADR 0012 Output Contract Slice

## Status

Proposed.

Analysis-only blueprint draft. Not implemented yet.

## Active Scope

Add shared success envelope primitives for the HTTP transport response package and migrate only success responses that already match the canonical public JSON shape.

This slice advances ADR `0012` output contract centralization without claiming full ADR closeout.

## FACT

- ADR `0012` is accepted and requires centralized API output contracts through presenter/output packages.
- `docs/api/0050_echo_http_contract.md` defines the canonical success envelope with `success`, `data`, and `meta`.
- `docs/api/0050_echo_http_contract.md` defines the canonical error envelope with `success`, `error`, and `meta`.
- Shared error envelope primitives already exist in `internal/transport/http/response`.
- Bootstrap already wires `httpresponse.HTTPErrorHandler` as the global Echo error handler.
- ProductCatalog transport currently has a local success envelope helper with `success`, `data`, and `meta`.
- ServiceCatalog transport currently has a local success envelope helper with `success`, `data`, and `meta`.
- Capability transport currently has a local success envelope with `success` and `data`, but no `meta`.
- Auth transport currently returns raw success DTOs for login, refresh, and Google auth flows, and `204 No Content` for logout/account role mutation success.
- System transport currently returns raw DTO/map success responses for `/api/me` and `/api/health`.

## GAP

- Shared success envelope primitives are not implemented in `internal/transport/http/response`.
- ProductCatalog and ServiceCatalog duplicate equivalent local success envelope helpers.
- Capability success responses do not currently include `meta`.
- Auth and System success responses are not centrally enveloped.
- Full ADR `0012` response/error coverage is not proven for every API surface.

## DECISION

Implement the smallest safe shared success envelope foundation now:

- Add shared success envelope primitives to `internal/transport/http/response`.
- Migrate ProductCatalog and ServiceCatalog success responses to use the shared helper because their current public JSON shape already matches the canonical shape.
- Preserve existing shared error envelope behavior.
- Preserve ProductCatalog business behavior.
- Preserve capability middleware behavior.
- Preserve public JSON field names.
- Do not change Capability/Auth/System success response shapes in this first patch unless a later accepted blueprint or owner decision explicitly allows those public contract changes.
- Do not claim ADR `0012` fully closed in this slice.

## SCOPE-IN

- `internal/transport/http/response/success.go`
- `internal/transport/http/response/success_test.go`
- ProductCatalog HTTP transport success response wrapping only.
- ServiceCatalog HTTP transport success response wrapping only.
- Focused tests proving canonical shared success envelope behavior.
- Existing transport package tests proving migrated handlers still return expected success envelopes.
- Existing docs remain source-of-truth unless implementation proof later requires ledger/handoff update.

## SCOPE-OUT

- Error envelope behavior changes.
- Capability middleware behavior changes.
- Capability response shape migration unless explicitly accepted later.
- Auth response shape migration.
- System `/api/me` response shape migration.
- Health response shape migration.
- ProductCatalog domain/usecase/postgres behavior changes.
- Product inventory/stock API.
- Stock adjustment create/reverse.
- ProductCatalog audit/outbox persistence.
- Runtime localization/language switching.
- Extended ProductCatalog filters.
- Router/server/bootstrap folder rename or cleanup.
- Migrations.
- Route/capability key changes.

## RESPONSE CONTRACT

Canonical success response:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

Canonical error response remains unchanged:

```json
{
  "success": false,
  "error": {
    "code": "example_code",
    "message": "example message"
  },
  "meta": {}
}
```

Recommended shared helper design:

```go
type SuccessEnvelope struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Meta    any  `json:"meta"`
}

func Success(data any) SuccessEnvelope {
	return SuccessWithMeta(data, EmptyMeta())
}

func SuccessWithMeta(data any, meta any) SuccessEnvelope {
	if meta == nil {
		meta = EmptyMeta()
	}

	return SuccessEnvelope{
		Success: true,
		Data:    data,
		Meta:    meta,
	}
}

func EmptyMeta() map[string]any {
	return map[string]any{}
}
```

Design notes:

- `Success(data)` must always include `meta`.
- `SuccessWithMeta(data, nil)` must normalize nil meta to an empty object.
- `EmptyMeta()` must return an empty `map[string]any`.
- Do not add map-based ad-hoc response construction in handlers.
- Do not add a shared no-meta success envelope in this slice because it would preserve inconsistency instead of centralizing the canonical contract.

## TARGET FILES

Implementation target files:

```text
internal/transport/http/response/success.go
internal/transport/http/response/success_test.go
internal/modules/productcatalog/transport/http/product_catalog_handler_response.go
internal/modules/productcatalog/transport/http/product_catalog_handler_read.go
internal/modules/productcatalog/transport/http/product_catalog_handler_write.go
internal/modules/productcatalog/transport/http/product_catalog_handler_lifecycle.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_response.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_read.go
internal/modules/servicecatalog/transport/http/service_catalog_handler_write.go
```

Read-only/reference files for this slice:

```text
docs/adr/0012-api-output-contract-centralization.md
docs/api/0050_echo_http_contract.md
docs/api/product_catalog.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/0004_adr_implementation_proof_index.md
docs/evidence/2026-06-14_productcatalog_runtime_smoke_proof.md
internal/modules/capability/transport/http
internal/modules/auth/transport/http
internal/modules/system/transport/http
internal/app/bootstrap/app.go
```

## TEST PLAN

Add or update focused tests to prove:

- `response.Success(data)` returns `success=true`, preserves data, and emits `meta` as an empty object.
- `response.SuccessWithMeta(data, meta)` preserves supplied metadata.
- `response.SuccessWithMeta(data, nil)` emits empty metadata instead of null.
- ProductCatalog success handler tests still pass after using the shared helper.
- ServiceCatalog success handler tests still pass after using the shared helper.
- Existing ProductCatalog error envelope tests still pass.
- Existing shared error envelope tests still pass.
- Capability/Auth/System tests still pass, proving this slice did not accidentally alter their current response behavior.

## PROOF REQUIRED

Run:

```bash
go test ./internal/transport/http/response/...
go test ./internal/modules/productcatalog/transport/http/... ./internal/modules/servicecatalog/transport/http/... ./internal/modules/capability/transport/http/... ./internal/modules/system/transport/http/... ./internal/modules/auth/transport/http/...
bash scripts/audit_hexagonal.sh
bash scripts/audit_route_capabilities.sh
make verify
```

The implementation cannot be called complete until visible command output proves the focused tests, architecture audits, route capability audit, and aggregate verification pass.

If implementation proof changes durable progress, update or draft updates for:

```text
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/0004_adr_implementation_proof_index.md
docs/handoffs/
```

## NEXT

Execution channel after blueprint acceptance:

```text
owner/local terminal
```

One next active step:

```text
Create shared success envelope primitives and migrate ProductCatalog plus ServiceCatalog success wrapping only.
```

## PATCH PLAN

1. Add `internal/transport/http/response/success.go` with `SuccessEnvelope`, `Success`, `SuccessWithMeta`, and `EmptyMeta`.
2. Add `internal/transport/http/response/success_test.go` proving `meta` is emitted as an empty object, supplied meta is preserved, and nil meta is normalized.
3. Replace ProductCatalog local response helper usage with `httpresponse.Success(...)`.
4. Remove the local `responseEnvelope` and `successEnvelope` from `product_catalog_handler_response.go` if no longer used.
5. Replace ServiceCatalog local response helper usage with `httpresponse.Success(...)`.
6. Remove the local `responseEnvelope` and `successEnvelope` from `service_catalog_handler_response.go` if no longer used.
7. Leave Capability response shape unchanged in this patch. Record it as an ADR `0012` remaining gap because it lacks `meta`.
8. Leave Auth and System success responses unchanged in this patch. Record them as ADR `0012` remaining gaps because enveloping them changes public response shape.
9. Preserve shared error behavior. Do not alter `error.go`, `error_response.go`, `error_handler.go`, or `error_envelope.go` except if tests need imports adjusted, which they should not.
10. Update only the blueprint/docs if implementation proof later changes progress. Do not update ledger/handoff until there is visible implementation/test proof.
