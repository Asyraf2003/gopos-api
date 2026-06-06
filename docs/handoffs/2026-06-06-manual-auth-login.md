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

## Decision

- Manual login is a debug/local auth lane, not production password login.
- Route: `POST /api/auth/manual/login`.
- Route is registered only when `AUTH_DEBUG_ENABLED=true`.
- Allowed emails:
  - `admin@example.com` -> `admin`
  - `kasir@example.com` -> `cashier`
- Both allowed emails require password `12345678`.
- The route reuses normal account, role, session, refresh token, access token, and auth middleware behavior.

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
```

## Blocked Or Partial Proof

`go test ./internal/app/bootstrap` with real tests could not run in the sandbox because it needs a PostgreSQL socket at `127.0.0.1:5432`, and the sandbox returned `socket: operation not permitted`.

## Open Gaps

- No runtime curl proof yet against a running local database.
- No production password login exists; this is intentionally debug/local only.
- No capability-control integration yet because capability foundation remains a separate blueprint.

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

Manual auth login foundation: 85%.

Remaining work is runtime DB proof and optional API envelope centralization if required by the next API contract pass.

## Context Window Status

Current session context is long but still usable. Handoff exists so the next model can resume safely from this file.
