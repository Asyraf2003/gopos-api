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

audit:allow-oversize reason=bootstrap-wiring
-->

# ADR Implementation Proof Index

## Status

Date updated: 2026-06-13

This evidence file maps accepted ADRs to current implementation proof.

It does not replace ADRs. ADRs own decisions. This file owns proof status.

## Proof Sources

Current proof sources inspected:

```text
docs/adr/*.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/2026-06-06-auth-runtime-local-dev.md
migrations/
internal/
scripts/
make verify output from 2026-06-13
docs/api/product_catalog.md
internal/transport/http/response
```

Current full gate proof:

```text
make verify
```

Result from 2026-06-13:

```text
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed
```

Gosec summary:

```text
Files: 112
Lines: 4659
Nosec: 0
Issues: 0
```

## Status Key

- `Proven`: implemented and backed by local proof or gate proof.
- `Partial`: some implementation or tests exist, but a known proof gap remains.
- `Decision only`: ADR exists, but implementation proof is not present yet.

## ADR Proof Matrix

| ADR | Decision Area | Proof Status | Current Proof | Remaining Gap |
| --- | --- | --- | --- | --- |
| `0001` | Raw Go, Echo, PostgreSQL, hexagonal foundation | Proven | Go module layout, Echo bootstrap, PostgreSQL adapters, hexagonal import audit, `make verify` pass | None for current foundation scope |
| `0002` | Auth minimum data model | Proven | `migrations/0001_auth_minimum.up.sql`, auth domain/ports/usecases/adapters, auth tests, `make verify` pass | Runtime manual auth proof remains tracked separately |
| `0003` | Roles, permissions, policies model | Proven | `migrations/0002_authorization_minimum.up.sql`, permission principal model, authz middleware, principal resolver, tests, `make verify` pass | None for current authorization foundation |
| `0004` | Authorization minimum schema | Proven | authorization migrations and PostgreSQL role/permission adapters/tests, `make verify` pass | None for current authorization foundation |
| `0005` | Default base role assignment strategy | Proven | `migrations/0005_authorization_assign_base_role_to_existing_accounts.up.sql`, manual login role tests, ledger proof | Runtime manual login proof remains incomplete |
| `0006` | Request principal model | Proven | `internal/modules/auth/domain/principal.go`, authn/authz middleware tests, principal resolver tests, `make verify` pass | None for current principal foundation |
| `0007` | Session revocation enforcement | Proven | logout usecase/handler, session revoker/status checker adapters and integration tests, `make verify` pass | Runtime end-to-end logout proof is not recorded in auth runtime evidence |
| `0008` | Refresh token rotation strategy | Proven | refresh usecase/handler tests, refresh session repository integration tests, `make verify` pass | Runtime end-to-end refresh proof is not recorded in auth runtime evidence |
| `0009` | Debug auth lane strategy | Partial | manual login usecase/handler tests, manual account repository, debug accounts documented in ledger | `docs/evidence/2026-06-06-auth-runtime-local-dev.md` is incomplete for manual auth runtime proof |
| `0010` | Authorization admin API minimum | Proven | account role assign/remove handlers/usecases/adapters/tests, protected admin routes, capability seed coverage, `make verify` pass | None for current admin authorization scope |
| `0011` | Code discipline audit gates | Proven | audit scripts, file-size policy, hexagonal audit, route capability audit, gosec, `make verify` pass | Allowlisted oversized bootstrap/config files remain documented warnings |
| `0012` | API output contract centralization | Partial | presentation DTO packages exist for system/capability/ProductCatalog; shared HTTP error envelope primitives and Echo error handler exist; shared success envelope primitives exist in `internal/transport/http/response`; ProductCatalog and ServiceCatalog success responses use the shared success helper; ProductCatalog mapped errors have stable public codes; ProductCatalog and disabled-capability error envelope tests pass; focused response/transport tests and `make verify` pass | Capability/Auth/System success response shape centralization remains incomplete; full response/error envelope coverage is not proven for every API surface |

## Current ADR Quality Assessment

ADR decision quality is strong: accepted ADRs are scoped, named, and aligned with current architecture.

ADR proof tracking is now explicit through this file.

Current ADR quality rating from local evidence:

```text
8.5/10
```

Reason:

- most foundation ADRs have implementation and `make verify` proof;
- known partial ADRs are labeled instead of hidden;
- full runtime auth proof, shared success envelope centralization, and full output centralization proof still need dedicated closeout evidence.

## Next Proof Improvements

- Complete manual auth runtime evidence for ADR `0009`.
- Complete remaining Capability/Auth/System response/error envelope coverage proof for ADR `0012` without changing public response shapes outside an accepted blueprint.
- Keep this index updated when future ADRs are accepted or implemented.
