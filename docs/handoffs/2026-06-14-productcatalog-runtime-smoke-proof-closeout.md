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

# ProductCatalog Runtime Smoke Proof Closeout Handoff

Date: 2026-06-14

Target agent for next session: Terminal Codex.

Template source for next prompt: `docs/templates/0121_codex_session_prompts.md`.

## Active Scope

Close out ProductCatalog runtime smoke proof after the interrupted Terminal Codex run.

## Current Snapshot

Source snapshot before closeout edits:

```text
git status --short
<clean>

git log -1 --oneline
b3a5d07 commit 291
```

## Files Included

```text
docs/blueprints/0033_productcatalog_runtime_smoke_proof_slice.md
docs/evidence/2026-06-14_productcatalog_runtime_smoke_proof.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/README.md
docs/evidence/2026-06-13_api_architecture_product_status_review.md
docs/handoffs/README.md
```

## Files Changed

```text
docs/evidence/2026-06-14_productcatalog_runtime_smoke_proof.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/README.md
docs/evidence/2026-06-13_api_architecture_product_status_review.md
docs/handoffs/2026-06-14-productcatalog-runtime-smoke-proof-closeout.md
docs/handoffs/README.md
```

## Files Forbidden To Touch

```text
internal/
cmd/
migrations/
scripts/
make/
```

No application code, migration, script, or Makefile change is required for this closeout.

## Blueprint Referenced

```text
docs/blueprints/0033_productcatalog_runtime_smoke_proof_slice.md
```

## ADR / Rules Referenced

```text
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/workflow/0071_handoff_protocol.md
docs/workflow/0072_transition_progress_ledger_protocol.md
docs/api/0050_echo_http_contract.md
docs/adr/0012-api-output-contract-centralization.md
```

## FACT

ProductCatalog runtime smoke proof is locally proven.

The smoke proof covered:

- local PostgreSQL migrations through `0013_seed_product_catalog_permissions_capabilities.up.sql`;
- Echo server started with `AUTH_DEBUG_ENABLED=true` on port `18081`;
- health endpoint returned `200` with `status=ok`;
- manual login returned `200` and access/refresh tokens were present;
- unauthenticated `GET /api/products` returned `401` with `authentication_required`;
- authenticated ProductCatalog `list`, `lookup`, `create`, and `show` routes succeeded;
- direct `products` table row check returned count `1`;
- reversible disable of `product_catalog.list` produced `403 capability_disabled`;
- capability was re-enabled and final list returned `200`;
- server process was manually stopped and port `18081` was verified closed.

Raw tokens, auth payloads, database credentials, and secrets were not written to docs.

## PROOF

Runtime proof summary:

```text
make db-migrate
[PASS] db migrate completed

make db-status
[APPLIED] migrations through 0013_seed_product_catalog_permissions_capabilities.up.sql

health_status=200
health_status_field=ok

manual_login status=200 access_token_present=True refresh_token_present=True
unauth_products status=401 success=False error_code=authentication_required has_meta=True
list_products status=200 success=True has_meta=True
lookup_products status=200 success=True has_meta=True
create_product status=201 success=True id_present=True kode_barang_present=True has_meta=True
show_product status=200 success=True id_match=True kode_barang_present=True status_field=active has_meta=True
db_product_row table=products count=1
disable_capability key=product_catalog.list status=200 enabled=False
list_products_capability_disabled status=403 success=False error_code=capability_disabled has_meta=True
enable_capability key=product_catalog.list status=200 enabled=True
list_products_after_reenable status=200 success=True
[PASS] product catalog runtime smoke proof passed
```

Docs closeout verification:

```text
GOCACHE=/tmp/gopos-go-build make verify
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] license header audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed
```

## GAP

Product inventory/stock API is not implemented.

ProductCatalog audit/outbox persistence is not implemented.

Shared success envelope centralization is not implemented.

ADR `0012` output contract centralization remains partial across all API surfaces.

ADR `0009` full debug auth runtime closeout remains partial; this smoke only proves manual login token issuance as part of the ProductCatalog runtime path.

Runtime language/output policy remains open.

Extended Laravel ProductCatalog filters remain unexposed.

Router/server/bootstrap cleanup remains a future blueprint candidate, not a blocker for the next slice.

## Decisions Made

ProductCatalog runtime smoke proof is marked locally proven in the active ledger.

ProductCatalog catalog API remains closed locally for pure API scope.

Product/inventory area remains partial and must not be called fully closed.

The next valid slice is shared success envelope and ADR `0012` output contract centralization.

## PROGRESS

ProductCatalog runtime smoke proof: 100% locally proven.

Estimated ProductCatalog full transition: 84%.

Estimated Business Phase 1: 60%.

Estimated overall Laravel-to-Go transition: 41%.

## NEXT

Execution channel: Terminal Codex.

One next active step:

```text
Shared success envelope and ADR 0012 output contract centralization.
```

Files the next agent should read:

```text
docs/README.md
docs/AGENTS.md
docs/0001_index.md
docs/0002_decision_policy.md
docs/0003_session_start_protocol.md
docs/adr/0012-api-output-contract-centralization.md
docs/api/0050_echo_http_contract.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/2026-06-14_productcatalog_runtime_smoke_proof.md
internal/transport/http/response
internal/modules/productcatalog/transport/http
internal/modules/servicecatalog/transport/http
internal/modules/capability/transport/http
internal/modules/system/transport/http
internal/modules/auth/transport/http
```

Files the next agent may edit:

```text
docs/blueprints/
docs/api/
docs/evidence/
docs/handoffs/
internal/transport/http/response
internal/modules/*/transport/http
internal/app/bootstrap
```

Files the next agent must not edit unless the owner changes scope:

```text
migrations/
internal/modules/productcatalog/domain
internal/modules/productcatalog/usecase
internal/platform/postgres
```

Proof commands for the next slice:

```text
go test ./internal/transport/http/response/...
go test ./internal/modules/productcatalog/transport/http/... ./internal/modules/servicecatalog/transport/http/... ./internal/modules/capability/transport/http/... ./internal/modules/system/transport/http/... ./internal/modules/auth/transport/http/...
bash scripts/audit_hexagonal.sh
bash scripts/audit_route_capabilities.sh
make verify
```

## CONTEXT WINDOW STATUS

Current closeout context status: fresh after interrupted runtime smoke proof was resumed. Runtime proof is recorded; next session can start from ADR `0012` without rerunning the ProductCatalog smoke unless a concrete gap is found.
