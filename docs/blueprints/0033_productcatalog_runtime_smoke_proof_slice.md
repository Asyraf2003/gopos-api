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

# ProductCatalog Runtime Smoke Proof Slice Blueprint

## Status

Accepted proof slice plan.

## Date

2026-06-13

## Active Scope

Prove the already-implemented ProductCatalog catalog API through the real local runtime path:

```text
local PostgreSQL -> migrations -> Echo server -> manual auth token -> protected ProductCatalog HTTP routes -> DB-backed responses
```

This is a proof slice, not a new ProductCatalog business-feature slice.

## FACT

ProductCatalog catalog API is locally closed at source/test level through:

- domain and usecase behavior;
- PostgreSQL persistence;
- protected runtime and capability route wiring;
- developer API docs;
- standardized ProductCatalog error envelope coverage.

The active transition ledger still marks ProductCatalog runtime smoke proof as not proven through a real local HTTP server, auth token, and DB-backed route.

The local debug auth route is available only when:

```text
AUTH_DEBUG_ENABLED=true
```

ProductCatalog routes are protected by auth, permission, and capability middleware.

## GAP

No current evidence proves this sequence end to end:

1. local PostgreSQL has required migrations applied;
2. Echo server starts against the local PostgreSQL database;
3. manual auth returns a valid local admin access token;
4. protected ProductCatalog routes work through HTTP with that bearer token;
5. at least one ProductCatalog write and one ProductCatalog read are DB-backed;
6. auth and capability guard failures still return the documented error envelope.

## DECISION

Run and record a narrow ProductCatalog runtime smoke proof.

Keep the smoke proof local and reversible. Do not implement inventory, stock adjustment, audit/outbox persistence, shared success envelope centralization, localization, extended filters, or route/server/presenter cleanup in this slice.

Use the existing manual auth lane with local admin credentials. Do not store access tokens in committed evidence.

Use a reversible capability-disable check for one ProductCatalog capability when feasible, then re-enable the capability before the slice ends.

## SCOPE-IN

- Apply or confirm PostgreSQL migrations.
- Start the Echo API server locally with debug auth enabled.
- Obtain an admin bearer token through `POST /api/auth/manual/login`.
- Prove unauthenticated ProductCatalog access fails with the standard error envelope.
- Prove authenticated DB-backed ProductCatalog routes:
  - `GET /api/products`
  - `GET /api/products/lookup`
  - `POST /api/products`
  - `GET /api/products/:id`
- Prove capability guard behavior by disabling and re-enabling a ProductCatalog capability when the admin capability API is available.
- Record sanitized proof output in evidence and the active ledger.

## SCOPE-OUT

- Inventory stock mutation.
- Stock adjustment create/reverse.
- ProductCatalog audit/outbox persistence.
- Shared success envelope centralization.
- ADR `0012` closeout.
- Runtime language switch or localization.
- Extended Laravel ProductCatalog filters.
- Large router/server/bootstrap/presenter folder renames.
- New production auth behavior.

## ROUTES UNDER PROOF

```text
POST /api/auth/manual/login
GET  /api/products
GET  /api/products/lookup
POST /api/products
GET  /api/products/:id
POST /api/admin/capabilities/:key/disable
POST /api/admin/capabilities/:key/enable
```

The admin capability routes are used only for reversible smoke proof of the existing capability guard.

## CAPABILITIES UNDER PROOF

```text
product_catalog.list
product_catalog.lookup
product_catalog.create
product_catalog.show
capability.manage
```

## PERMISSIONS UNDER PROOF

```text
product_catalog.read
product_catalog.manage
capability.manage
```

## RESPONSE CONTRACT

Success responses must keep the existing envelope shape:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

Error responses must keep the standardized error envelope shape:

```json
{
  "success": false,
  "error": {
    "code": "authentication_required",
    "message": "authentication required"
  },
  "meta": {}
}
```

The exact error code may differ by failure type, but the envelope family must be standard.

## PROOF PLAN

1. Run PostgreSQL migrations and status check.
2. Start the local Echo API server on a non-default smoke port with `AUTH_DEBUG_ENABLED=true`.
3. Login as local admin and keep the access token only in shell memory.
4. Run sanitized HTTP checks:
   - unauthenticated `GET /api/products` returns `success=false`;
   - authenticated `GET /api/products` returns `success=true`;
   - authenticated `GET /api/products/lookup` returns `success=true`;
   - authenticated `POST /api/products` returns `201` and a product ID;
   - authenticated `GET /api/products/:id` returns the same product ID;
   - direct SQL confirms the created product row exists;
   - disabling `product_catalog.list` makes authenticated `GET /api/products` fail with `403`;
   - re-enabling `product_catalog.list` restores authenticated `GET /api/products` to `200`.
5. Stop the local server.
6. Run final quality gates.

## PROOF REQUIRED

```bash
make db-migrate
make db-status
go test ./...
bash scripts/audit_hexagonal.sh
bash scripts/audit_route_capabilities.sh
make verify
```

Runtime proof must be recorded without committing raw bearer tokens or database credentials.

## NEXT

If this proof passes, ProductCatalog runtime smoke proof becomes locally proven.

The next valid slice should be shared success envelope and ADR `0012` output contract centralization, unless a smoke failure exposes a bug that must be fixed first.
