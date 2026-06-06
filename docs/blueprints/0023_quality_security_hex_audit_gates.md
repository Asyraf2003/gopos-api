# Quality Security Hex Audit Gates Blueprint

## FACT
- Current Makefile has basic targets for `test`, `audit-ai-rules`, `audit-file-size`, `security-gosec`, and `audit-all`.
- Current scripts do not yet include strict hexagonal import-boundary audit.
- Current scripts do not yet include a dedicated formatting audit or `go vet` audit.
- `internal/modules/system/transport/http/health_handler.go` currently depends directly on `pgxpool`, which weakens the transport boundary.
- The owner wants strict testing, security, code uniformity, and hexagonal folder responsibility gates.

## GAP
- Capability-control audit is not implemented yet because capability foundation remains separate.
- Runtime DB proof may require local PostgreSQL access outside sandbox.

## DECISION
- Add strict production-code import-boundary audit script.
- Add formatting audit script based on `gofmt -l`.
- Add Go vet audit script.
- Keep security audit through existing `gosec` script.
- Update aggregate audit to run tests and all audit scripts.
- Add Makefile targets with stable names: `fmt`, `vet`, `audit-format`, `audit-hex`, `check`, `verify`, `ci`.
- Refactor health handler so HTTP transport depends on a health-check port instead of `pgxpool`.

## SCOPE-IN
- Makefile quality targets.
- Audit scripts.
- Health checker port and PostgreSQL adapter.
- Docs handoff update.

## SCOPE-OUT
- Full capability registry audit.
- Seed profiles.
- Production password auth.
- Runtime DB migration proof.

## TEST/PROOF PLAN
- `GOCACHE=/tmp/go-build-cache go test ./internal/modules/auth/...`
- `GOCACHE=/tmp/go-build-cache go test ./internal/modules/system/...`
- `GOCACHE=/tmp/go-build-cache go test ./internal/app/bootstrap -run '^$'`
- `GOCACHE=/tmp/go-build-cache go vet ./...`
- `bash scripts/audit_format.sh`
- `bash scripts/audit_hexagonal.sh`
- `bash scripts/audit_file_size.sh`
- `bash scripts/audit_ai_rules.sh`

## NEXT ACTIVE STEP
Implement the quality/security/hex audit gates and focused proof.
