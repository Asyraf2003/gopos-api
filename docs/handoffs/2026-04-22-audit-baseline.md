# Audit Handoff - 2026-04-22

## Status
Audit baseline repo untuk screening dan file hygiene sudah selesai.

## Completed
- audit docs/core, AI_RULES, dan ADR baseline
- audit scripts dan jalur screening repo
- audit auth middleware, permission middleware, principal resolver, dan session enforcement
- refactor file produksi oversize sampai lolos file size gate
- refactor file test oversize sampai lolos file size gate
- verifikasi aggregate screening lewat `make audit-all`

## Locked Facts
- `make audit-all` pass
- `go test ./...` pass
- `scripts/audit_ai_rules.sh` pass
- `scripts/audit_file_size.sh` pass
- `scripts/audit_security_gosec.sh` pass
- `gosec` summary menunjukkan `Issues : 0`
- file oversize yang tersisa hanya 2 file allowlist resmi:
  - `internal/app/bootstrap/app.go`
  - `internal/config/config.go`
- integration tests postgres yang dijalankan manual masih skip ketika `DATABASE_URL` tidak diset

## Files Changed
- `internal/platform/token/jwt/issuer.go`
- `internal/platform/token/jwt/issue_access_token.go`
- `internal/platform/token/jwt/token_codec.go`
- `internal/platform/token/jwt/verifier_test.go`
- `internal/platform/token/jwt/verifier_test_helpers.go`

- `internal/platform/postgres/auth_account_identity_repository.go`
- `internal/platform/postgres/auth_account_identity_resolve.go`
- `internal/platform/postgres/auth_account_identity_upsert.go`
- `internal/platform/postgres/principal_resolver.go`
- `internal/platform/postgres/principal_resolver_roles.go`
- `internal/platform/postgres/principal_resolver_permissions.go`

- `internal/platform/google/oidc_google.go`
- `internal/platform/google/oidc_google_auth_url.go`
- `internal/platform/google/oidc_google_exchange.go`

- `internal/modules/auth/usecase/google_callback.go`
- `internal/modules/auth/usecase/google_callback_input.go`
- `internal/modules/auth/usecase/google_callback_tx.go`
- `internal/modules/auth/usecase/google_callback_output.go`

- `internal/platform/postgres/principal_resolver_integration_test.go`
- `internal/platform/postgres/principal_resolver_integration_helpers_test.go`
- `internal/platform/postgres/account_role_remover_integration_test.go`
- `internal/platform/postgres/account_role_remover_integration_helpers_test.go`
- `internal/platform/postgres/auth_account_identity_repository_integration_test.go`
- `internal/platform/postgres/auth_account_identity_repository_integration_helpers_test.go`
- `internal/platform/postgres/refresh_session_repository_integration_test.go`
- `internal/platform/postgres/refresh_session_repository_integration_helpers_test.go`

- `internal/transport/http/middleware/authz_test.go`
- `internal/transport/http/middleware/authz_test_helpers.go`
- `internal/transport/http/middleware/authn_test.go`
- `internal/transport/http/middleware/authn_test_helpers.go`
- `internal/transport/http/middleware/authn_test_fakes.go`

## Verification Proof
Command set yang sudah dibuktikan hijau:
- `go test ./...`
- `bash scripts/audit_ai_rules.sh`
- `bash scripts/audit_file_size.sh`
- `bash scripts/audit_security_gosec.sh`
- `make audit-all`

## Pending
- review dan rapikan ADR `0011` terhadap kondisi script audit aktual
- putuskan apakah perlu implement `scripts/audit_error_policy.sh`, `scripts/audit_time_policy.sh`, dan `scripts/audit_structure_policy.sh`
- review gap ADR `0009` debug auth lane vs wiring bootstrap aktual
- jalankan integration tests postgres dengan `DATABASE_URL` aktif untuk proof DB live
- putuskan apakah dua file allowlist tetap dipertahankan atau dipecah lagi

## Recommended Next Step
Fokus berikutnya:
- rapikan governance docs dan ADR agar sesuai dengan kondisi repo saat ini
- setelah itu baru lanjut ke debug auth lane
