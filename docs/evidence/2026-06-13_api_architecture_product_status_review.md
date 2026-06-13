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

audit:allow-oversize reason=evidence-report
-->

# API Architecture Product Status Review

Date: 2026-06-13

## FACT

- The repository is a pure Go Echo and PostgreSQL API backend. No UI scope is active.
- ProductCatalog catalog API routes are present in source, docs, capability seed migration, and route capability manifest:
  - `GET /api/products`
  - `POST /api/products`
  - `GET /api/products/lookup`
  - `GET /api/products/:id`
  - `PUT /api/products/:id`
  - `DELETE /api/products/:id`
  - `PATCH /api/products/:id/restore`
  - `GET /api/products/:id/versions`
- ProductCatalog capability keys are present in `scripts/config/route_capabilities.tsv` and `migrations/0013_seed_product_catalog_permissions_capabilities.up.sql`.
- ProductCatalog domain imports only standard library packages.
- ProductCatalog ports import `context`, `errors`, `time`, and ProductCatalog domain only.
- ProductCatalog usecases import ProductCatalog domain and ports, not Echo or PostgreSQL.
- ProductCatalog HTTP transport imports Echo, usecases, `internal/presentation/http/id/productcatalog`, and `internal/transport/http/response`, not PostgreSQL.
- PostgreSQL ProductCatalog adapter code lives in `internal/platform/postgres` and satisfies `ProductRepository`, `ProductReader`, `ProductVersionRepository`, and `ProductDuplicateChecker`.
- `internal/app/bootstrap/app.go` wires PostgreSQL adapters into usecases and registers ProductCatalog routes behind auth, permission, and capability middleware.
- Bootstrap currently injects `productCatalogNoopAuditRecorder`, so ProductCatalog audit/outbox persistence is not implemented.
- ProductCatalog success responses still use a local `successEnvelope` in the handler package. Shared success envelope centralization is not implemented.
- Shared HTTP error envelope primitives exist in `internal/transport/http/response`, and bootstrap wires `httpresponse.HTTPErrorHandler`.
- Other API surfaces still use local envelopes and plain Echo HTTP errors, so ADR 0012 remains partial.
- Active ProductCatalog handoffs under `docs/handoffs/` are empty. ProductCatalog handoff history is staged/recorded under `docs/archive/handoffs-closed/`.

## GAP

- Product inventory/stock API is not closed.
- Product audit/outbox persistence is not closed.
- Product runtime smoke proof through a real local HTTP server, auth token, and DB-backed ProductCatalog route is not proven.
- Runtime localization/language switching is not implemented.
- Extended Laravel ProductCatalog filters remain unexposed:
  - `sort_by`
  - `sort_dir`
  - `merek`
  - `ukuran_min`
  - `ukuran_max`
  - `harga_min`
  - `harga_max`
  - `stok_saat_ini`
- Shared success envelope centralization is not implemented.
- ADR 0012 is not fully closed across all API surfaces.

## DECISION

ProductCatalog catalog API: CLOSED locally.

Product inventory/stock API: NOT CLOSED.

Product audit/outbox: NOT CLOSED.

Product runtime smoke proof: NOT PROVEN.

Product full backend transition: PARTIAL.

ProductCatalog catalog API should be treated separately from the wider Product/inventory area. Do not count UI as a ProductCatalog backend gap.

Architecture cleanup is useful but not required before the next feature slice. The current repo is hexagonal enough for the ProductCatalog catalog API based on imports, package roles, route capability audit, and bootstrap wiring. The next improvement should prove the runtime path before larger refactors.

Workflow/docs needed repair. Stale active ledger text still named ProductCatalog UI as a gap and said HTTP/runtime/capability work was not started. This review updates those active docs.

## ARCHITECTURE ASSESSMENT

| Area | Current State | Good Enough? | Gap | Recommendation |
| --- | --- | --- | --- | --- |
| Domain | `internal/modules/productcatalog/domain` imports standard library only and owns Product invariants/lifecycle. | Yes | None for current catalog API. | Keep as is. |
| Ports | `product_ports.go` defines repository, reader, version, duplicate checker, and audit recorder ports. | Yes | Audit port is not backed by durable outbox; one PostgreSQL adapter implements several ports. | Keep current ports for catalog API; split only when audit/outbox or inventory pressure creates real complexity. |
| Usecase | Usecases depend on ports/domain and record version/audit side effects through ports. | Yes | Mutations are not wrapped in an explicit ProductCatalog unit-of-work boundary. | Address with audit/outbox slice, not incidental cleanup. |
| PostgreSQL adapter | Product SQL and pgx usage live in `internal/platform/postgres`; bootstrap wires adapters. | Yes | `internal/platform/postgres` is growing as a shared adapter package. | Defer `internal/modules/*/store/postgres` until an accepted architecture blueprint proves the current package is hurting ownership. |
| HTTP transport | ProductCatalog handlers parse input, call one usecase, map errors, and call presenters. | Yes | Handlers still own success envelope wrapping. | Centralize success envelope in a small shared response slice later. |
| Presenter/output | ProductCatalog DTO mapping is in `internal/presentation/http/id/productcatalog`. | Yes | Package path is ADR-backed but not as self-describing as `internal/transport/http/presenter`; English/localization path is not implemented. | Keep `internal/presentation/http/id` for now; consider a localization/output blueprint before renaming. |
| Error envelope | Shared error primitives and Echo handler exist in `internal/transport/http/response`; ProductCatalog known errors use stable codes. | Partial | Other API surfaces still return plain Echo HTTP errors. | Complete ADR 0012 across all API surfaces after runtime smoke proof or as its own slice. |
| Success envelope | ProductCatalog, ServiceCatalog, and Capability own local response envelope structs. | No | No shared success envelope helper. | Create a focused shared success envelope centralization slice. |
| Auth middleware | Authn/authz middleware is centralized in `internal/transport/http/middleware`. | Yes | Middleware errors are normalized by shared error handler but still originate as plain Echo errors. | Leave until ADR 0012 completion. |
| Capability control | ProductCatalog routes have manifest rows, seed rows, permission keys, and disabled-capability tests. | Yes | None for current catalog API. | Keep route capability audit mandatory. |
| Router/server bootstrap | `internal/app/bootstrap/app.go` wires routes directly; ProductCatalog has a small helper for per-route capability groups. | Partial | Bootstrap is large and owns route grouping detail. | Draft router/server cleanup blueprint later; do not rename large folders without ADR/blueprint. |
| Reporting/docs workflow | Ledger, archive README, API docs, and handoff README now reflect catalog API closeout and no UI scope. | Yes | Runtime smoke proof and ADR 0012 remain open. | Keep this report as current evidence and run final proof gates. |

## PRODUCT STATUS

ProductCatalog catalog API: CLOSED locally.

Product inventory/stock API: NOT CLOSED.

Product audit/outbox: NOT CLOSED.

Product runtime smoke proof: NOT PROVEN.

Product full backend transition: PARTIAL.

## RECOMMENDED NEXT SLICE

Choose exactly one next slice:

```text
ProductCatalog runtime smoke proof with local auth token and DB-backed HTTP route.
```

Why this one now:

- It validates the already-closed catalog API through the real runtime path.
- It is smaller and less invasive than shared envelope or architecture cleanup.
- It gives stronger proof before inventory, audit/outbox, or localization work builds on ProductCatalog.

Candidates not chosen:

- Shared success envelope centralization: useful, but should not precede runtime proof that the current API actually works end to end.
- Complete ADR 0012 across all API surfaces: important, but broader than ProductCatalog status validation.
- Audit/outbox implementation: required for ProductCatalog full transition, but should follow runtime smoke proof or an accepted audit/outbox blueprint.
- Product inventory/stock mutation API blueprint: valid future product-area work, but catalog runtime proof should come first.
- Runtime localization/language/output contract blueprint: useful later; current API status clarity does not require it first.
- Architecture/router/server/presenter cleanup blueprint: useful when bootstrap or output package friction becomes the active problem; not required before runtime smoke proof.

## FILE CHANGES MADE

- `docs/evidence/2026-06-13_api_architecture_product_status_review.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`
- `docs/evidence/0004_adr_implementation_proof_index.md`
- `docs/evidence/README.md`
- `docs/architecture/0024_current_repo_layout.md`
- `docs/api/README.md`
- `docs/handoffs/README.md`
- `docs/handoffs/2026-06-08-servicecatalog-runtime-capability-implementation.md`

Pre-existing worktree changes observed but not created by this review:

- ProductCatalog handoffs staged as renames from `docs/handoffs/` to `docs/archive/handoffs-closed/`.
- `docs/archive/handoffs-closed/README.md` already modified to list ProductCatalog completed slice history.

## PROOF

Review proof collected before this file was created:

```text
git status --short
R  ProductCatalog handoffs from docs/handoffs/ to docs/archive/handoffs-closed/
M  docs/archive/handoffs-closed/README.md
```

```text
fd -HI -t f . docs/handoffs --max-depth 1
docs/handoffs/2026-06-06-auth-runtime-local-dev.md
docs/handoffs/2026-06-08-servicecatalog-runtime-capability-blueprint-accepted.md
docs/handoffs/2026-06-08-servicecatalog-runtime-capability-implementation.md
docs/handoffs/2026-06-09-cli-command-formatter-rules.md
docs/handoffs/README.md
```

```text
fd -HI -t f '(?i)productcatalog' docs/handoffs --max-depth 1
<no output>
```

```text
GOCACHE=/tmp/gopos-go-build go list ./...
31 packages listed successfully.
```

```text
bash scripts/audit_hexagonal.sh
[PASS] hexagonal import audit passed
```

```text
bash scripts/audit_route_capabilities.sh
checked route capability rows: 21
[PASS] route capability audit passed
```

Final proof after docs changes must include:

```text
GOCACHE=/tmp/gopos-go-build go test ./...
bash scripts/audit_hexagonal.sh
bash scripts/audit_route_capabilities.sh
make verify
```

## NEXT

Execution channel: Terminal Codex.

Create or accept a small ProductCatalog runtime smoke proof blueprint, then prove a protected DB-backed ProductCatalog HTTP route using a local auth token.
