# Manual Auth Login Foundation Blueprint

## FACT
- Auth foundation already exists under `internal/modules/auth`.
- JWT issuer/verifier, session store, refresh flow, logout flow, request principal, role assignment, and permission middleware already exist.
- ADR 0009 accepts a debug auth lane for local/development verification.
- Existing auth routes are registered only when Google auth is configured.
- The owner needs manual login for `admin@example.com` and `kasir@example.com` so build, tests, and early API work can satisfy auth requirements.

## GAP
- There is no manual login endpoint yet.
- There is no password credential model yet.
- Capability-control foundation is still separate work.

## DECISION
- Add a debug/local manual login endpoint, not a production password login.
- Endpoint path: `POST /api/auth/manual/login`.
- Endpoint is registered only when `AUTH_DEBUG_ENABLED=true`.
- Allowed emails:
  - `admin@example.com` -> role `admin`
  - `kasir@example.com` -> role `cashier`
- Both allowed emails require debug password `12345678`.
- The endpoint must reuse normal account, role, session, token, refresh, authn, and principal mechanisms.
- Manual login creates or reuses an account by email, ensures the configured role, creates an `auth_sessions` row, and issues a normal access token plus refresh token.
- Manual login must not bypass middleware behavior for later protected requests.

## SCOPE-IN
- Auth use case for manual login.
- Auth port for resolving or creating an account by email.
- PostgreSQL adapter implementation.
- Echo handler for manual login.
- Bootstrap wiring under existing `/api/auth` group.
- Focused tests for use case and handler.

## SCOPE-OUT
- Production password login.
- Password hashing.
- User registration.
- Capability-control implementation.
- UI login page.
- New database schema unless a test proves it is required.

## SECURITY
- Debug endpoint is explicit opt-in through `AUTH_DEBUG_ENABLED=true`.
- It must not register by default.
- It must not accept arbitrary emails.
- It must not log tokens or request bodies.

## API CONTRACT

Request:

```json
{
  "email": "admin@example.com",
  "password": "12345678"
}
```

Success:

```json
{
  "access_token": "...",
  "access_exp": "...",
  "refresh_token": "...",
  "refresh_exp": "...",
  "trust_level": "aal1",
  "step_up_required": false
}
```

Errors:

- `400` for malformed request.
- `401` for unsupported email or wrong password.
- `404`-like role setup is treated as server setup error through adapter result if seed roles are missing.

## TEST PLAN
- Use case rejects unsupported email.
- Use case rejects wrong password.
- Use case creates session and issues token for admin.
- Use case assigns `admin` role for `admin@example.com`.
- Use case assigns `cashier` role for `kasir@example.com`.
- Handler returns `400` for bad JSON.
- Handler returns `401` for unsupported email.
- Bootstrap registers manual route only when debug auth is enabled.

## NEXT ACTIVE STEP
Implement manual auth login endpoint and focused tests.
