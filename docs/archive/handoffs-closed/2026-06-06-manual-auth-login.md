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

# Manual Auth Login Handoff

## Date

2026-06-06

## Active Scope

Manual auth login foundation for local/build/testing auth requirements.

## Blueprint

- `docs/blueprints/0022_manual_auth_login_foundation.md`

## Files Changed

- `internal/app/bootstrap/app.go`
- `internal/modules/auth/ports/manual_account_repository.go`
- `internal/modules/auth/usecase/manual_login.go`
- `internal/modules/auth/usecase/manual_login_config.go`
- `internal/modules/auth/usecase/manual_login_seed.go`
- `internal/modules/auth/usecase/manual_login_types.go`
- `internal/modules/auth/usecase/manual_login_test.go`
- `internal/modules/auth/usecase/manual_login_role_test.go`
- `internal/modules/auth/usecase/manual_login_test_helpers_test.go`
- `internal/modules/auth/transport/http/manual_login_handler.go`
- `internal/modules/auth/transport/http/manual_login_handler_test.go`
- `internal/modules/auth/transport/http/manual_login_handler_error_test.go`
- `internal/modules/auth/transport/http/manual_login_handler_test_helpers_test.go`
- `internal/platform/postgres/manual_account_repository.go`
- `docs/blueprints/0023_quality_security_hex_audit_gates.md`
- `internal/modules/system/ports/health_checker.go`
- `internal/platform/postgres/health_checker.go`
- `scripts/audit_format.sh`
- `scripts/audit_go_vet.sh`
- `scripts/audit_hexagonal.sh`
- `scripts/audit_all.sh`
- `Makefile`
- `docs/scripts/0090_makefile_and_scripts.md`
- `scripts/audit_ai_rules.sh`

## Decision

- Manual login is a debug/local auth lane, not production password login.
- Route: `POST /api/auth/manual/login`.
- Route is registered only when `AUTH_DEBUG_ENABLED=true`.
- Allowed emails:
  - `admin@example.com` -> `admin`
  - `kasir@example.com` -> `cashier`
- Both allowed emails require password `12345678`.
- The route reuses normal account, role, session, refresh token, access token, and auth middleware behavior.
- Quality gate now includes format, vet, file size, hexagonal import-boundary, docs audit, and security gosec target.
- `make check` runs format, vet, file size, hexagonal, and docs audits.
- `make verify` delegates to aggregate audit.
- Health transport no longer imports `pgxpool`; it uses a `system/ports.HealthChecker` implemented by PostgreSQL platform code.

## Proof Collected

```text
GOCACHE=/tmp/go-build-cache go test ./internal/modules/auth/usecase
PASS

GOCACHE=/tmp/go-build-cache go test ./internal/modules/auth/transport/http
PASS

GOCACHE=/tmp/go-build-cache go test ./internal/modules/auth/...
PASS

GOCACHE=/tmp/go-build-cache go test ./internal/app/bootstrap -run '^$'
PASS compile-only, no tests run

GOCACHE=/tmp/go-build-cache go test ./internal/platform/postgres -run '^$'
PASS compile-only, no tests

bash scripts/audit_file_size.sh
PASS

bash scripts/audit_ai_rules.sh
PASS

make check
PASS

bash scripts/audit_format.sh
PASS

bash scripts/audit_hexagonal.sh
PASS

bash scripts/audit_go_vet.sh
PASS

GOCACHE=/tmp/go-build-cache go test ./internal/modules/system/...
PASS
```

## Blocked Or Partial Proof

`go test ./internal/app/bootstrap` with real tests could not run in the sandbox because it needs a PostgreSQL socket at `127.0.0.1:5432`, and the sandbox returned `socket: operation not permitted`.

`bash scripts/audit_security_gosec.sh` did not complete because the installed gosec binary reports `Version: dev` and fails to load Go stdlib package `unsafe` from `/usr/lib/go/src/unsafe`. The script is wired as a security gate, but this environment needs gosec/toolchain repair before claiming security scan proof.

## Open Gaps

- No runtime curl proof yet against a running local database.
- No production password login exists; this is intentionally debug/local only.
- No capability-control integration yet because capability foundation remains a separate blueprint.
- No passing gosec proof yet due local gosec/toolchain issue.

## Next Valid Active Step

Run local DB migration and runtime proof:

```text
AUTH_DEBUG_ENABLED=true
AUTH_JWT_SECRET=<local secret>
POST /api/auth/manual/login {"email":"admin@example.com","password":"12345678"}
GET /api/me with returned bearer token
POST /api/auth/manual/login {"email":"kasir@example.com","password":"12345678"}
GET /api/me with returned bearer token
```

## Estimated Progress

Manual auth login foundation: 88%.

Quality/security/hex audit gates: 75%.

Remaining work is runtime DB proof, gosec/toolchain repair, optional API envelope centralization if required by the next API contract pass, and future capability-control audit once capability foundation exists.

## Context Window Status

Current session context is long but still usable. Estimated context capacity used: 78%. Handoff exists so the next model can resume safely from this file.
