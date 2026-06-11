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

Active scope

Manual auth local runtime proof and local developer workflow cleanup for pos-go / gopos-api.

Blueprint referenced

docs/archive/handoffs-closed/2026-06-06-manual-auth-login.md

Files that Codex should read

AGENTS.md
docs/AGENTS.md
docs/archive/handoffs-closed/2026-06-06-manual-auth-login.md
Makefile
README.md
.env.example
.gitignore
scripts/db_dev_setup.sh
scripts/dev_smoke.sh
scripts/dev_manual_login.sh
scripts/db_migrate.sh
scripts/db_status.sh
scripts/db_adopt_existing.sh
internal/app/bootstrap/app.go
internal/config/config.go
internal/modules/auth/transport/http/manual_login_handler.go
internal/transport/http/middleware/authn.go
migrations/0001_auth_minimum.up.sql
migrations/0002_authorization_minimum.up.sql
migrations/0003_authorization_seed_minimum.up.sql
migrations/0004_authorization_role_permissions_seed.up.sql
migrations/0005_authorization_assign_base_role_to_existing_accounts.up.sql

Files Codex may edit

Makefile
README.md
scripts/db_dev_setup.sh
scripts/dev_smoke.sh
scripts/dev_manual_login.sh
docs/handoffs/2026-06-06-auth-runtime-local-dev.md
docs/evidence/2026-06-06-auth-runtime-local-dev.md

Files Codex must not edit

.env
.git/
go.sum, unless dependency changes are explicitly required and proven
migrations/*.sql, unless the active scope changes to schema work
internal/modules/* domain or use case files, unless a failing proof shows code-level auth defects
internal/transport/http/middleware/authn.go, unless a failing proof shows middleware defects

Decisions made

.env remains local-only and ignored by git.

.env.example remains tracked and may contain local development example credentials only.

README.md must document the local PostgreSQL setup, common local flow, auth debug mode, and developer commands.

Local PostgreSQL setup should be automated through Makefile instead of requiring repeated manual SQL.

The local developer flow should support these commands:

make db-dev-setup
make db-migrate
make run
make smoke
make auth-login-admin
make auth-login-cashier

Manual auth remains debug/local only and requires AUTH_DEBUG_ENABLED=true.

The health endpoint path used for local smoke testing is /api/health, not /health.

The default Makefile runtime port is 8081.

The smoke and manual-login helper scripts must preserve command-line or Makefile-provided HTTP_PORT/API_BASE_URL over .env values.

Facts proven from provided data

Repository local path:

/home/asyraf/Code/go/pos-go

Local repository has:

AGENTS.md
docs/
cmd/
internal/
migrations/
scripts/
Makefile
go.mod
go.sum
.env
.env.example
.gitignore

go.mod exists.

cmd/api/main.go exists.

Makefile help originally contained:

make run
make db-migrate
make db-status
make db-adopt-existing

.env is ignored by .gitignore.

.env is not tracked by git.

.env.example is tracked by git.

Local .env contains APP_ENV=local.

Local .env contains HTTP_PORT=8080.

Local .env contains DATABASE_URL for user posgo_app at 127.0.0.1:5432 database posgo_app_db.

Local .env contains AUTH_DEBUG_ENABLED=true after local update.

Local PostgreSQL status was proven:

systemctl is-active postgresql returned active.
127.0.0.1:5432 was listening.
psql version returned PostgreSQL 18.4.
Admin connection as postgres to postgres succeeded.

Local app database connection was proven:

current_user = posgo_app
current_database = posgo_app_db

scripts/db_dev_setup.sh was added locally and proven executable.

scripts/dev_smoke.sh was added locally and proven executable.

scripts/dev_manual_login.sh was added locally and proven executable.

Makefile was locally updated to include:

make dev
make smoke
make auth-login-admin
make auth-login-cashier
make db-dev-setup

README.md was locally created or updated with local PostgreSQL setup, common local flow, auth debug mode, and developer commands.

make db-dev-setup proof output showed:

ALTER ROLE
GRANT
current_user = posgo_app
current_database = posgo_app_db
[PASS] db dev setup completed

make db-migrate proof output showed:

[SKIP] already applied: 0001_auth_minimum.up.sql
[SKIP] already applied: 0002_authorization_minimum.up.sql
[SKIP] already applied: 0003_authorization_seed_minimum.up.sql
[SKIP] already applied: 0004_authorization_role_permissions_seed.up.sql
[SKIP] already applied: 0005_authorization_assign_base_role_to_existing_accounts.up.sql
[PASS] db migrate completed

make db-status proof output showed all five migrations APPLIED.

make run built .bin/pos-go-api and started the app with HTTP_PORT=8081.

Initial /health request returned 404.

Correct /api/health request through make smoke returned 200 OK with:

{"database":"up","status":"ok"}

GET /api/me without token returned 401 Unauthorized with:

{"message":"missing bearer token"}

Initial manual login failed when server was not restarted with AUTH_DEBUG_ENABLED=true or when helper scripts targeted port 8080 from .env.

After setting AUTH_DEBUG_ENABLED=true and restarting with explicit debug environment, make auth-login-admin succeeded.

POST /api/auth/manual/login for admin@example.com returned HTTP_STATUS=200 and an access_token.

The generated admin access token was then used against GET /api/me.

GET /api/me with the admin bearer token returned 200 OK and included:

roles: ["admin"]

permissions included:

account.role.assign
auth.session.logout
auth.session.refresh
inventory.manage
payment.create
profile.self.read
report.read
sale.order.create
sale.order.read

trust_level was aal1.

Gaps

No proof has been provided yet for make auth-login-cashier.

No proof has been provided yet for AUTH_DEBUG_ENABLED=false disabling POST /api/auth/manual/login.

No proof has been provided yet for make check after local Makefile, README, and script changes.

No proof has been provided yet for make security-gosec after local changes.

No proof has been provided yet for git status after local changes.

No proof has been provided yet that scripts pass shellcheck.

No proof has been provided yet that README.md existed before this session. It may have been created locally during this work.

No proof has been provided yet that local changes have been committed or pushed.

Do not claim these changes are in the remote repository unless git proof is provided.

Recommended next active step

Create evidence file for the local auth runtime proof and local developer workflow changes.

Suggested evidence file:

docs/evidence/2026-06-06-auth-runtime-local-dev.md

After evidence is written, run verification commands.

If verification passes, the next scope can move toward Laravel legacy API migration inventory.

Do not start Laravel legacy migration implementation until auth runtime proof and required quality gates are recorded.

Proof commands Codex should run

Print local status and changed files:

git status --short

Check tracked env safety:

git ls-files -- .env .env.example .gitignore

Check secret exposure in tracked diff, excluding .env:

git diff -- . ':!.env' | grep -Ei 'password|secret|token|DATABASE_URL|PRIVATE|KEY' || true

Check Makefile targets:

make help | rg -e 'make dev' -e 'make smoke' -e 'auth-login' -e 'db-dev-setup' -e 'db-migrate' -e 'make run'

Check local PostgreSQL setup:

unset DATABASE_URL
make db-dev-setup

Check migrations:

unset DATABASE_URL
make db-migrate
make db-status

Start server in a dedicated terminal:

AUTH_DEBUG_ENABLED=true HTTP_PORT=8081 make run

From a second terminal, run smoke proof:

make smoke

Run admin manual login proof:

make auth-login-admin

Run cashier manual login proof:

make auth-login-cashier

Run debug disabled proof:

Stop the running server.

AUTH_DEBUG_ENABLED=false HTTP_PORT=8081 make run

From a second terminal:

curl -i -X POST 'http://127.0.0.1:8081/api/auth/manual/login' -H 'Content-Type: application/json' -d '{"email":"admin@example.com","password":"12345678"}'

Expected result for debug disabled proof:

POST /api/auth/manual/login must not return HTTP 200.

Run quality gates:

make check
make security-gosec

If make security-gosec fails because of local toolchain or gosec installation, capture the exact command output and classify it as BLOCKED, not PASS.

Suggested evidence content structure

FACT

Record exact command outputs only.

DECISION

Record that .env remains local-only, .env.example remains tracked, README.md documents the local developer workflow, and local DB bootstrap is handled by make db-dev-setup.

PROOF

Paste command and output pairs for:

git ls-files -- .env .env.example .gitignore
make db-dev-setup
make db-migrate
make db-status
make smoke
make auth-login-admin
make auth-login-cashier
AUTH_DEBUG_ENABLED=false manual login check
make check
make security-gosec

GAP

Record any command not run or any command that failed.

NEXT

If all required proofs pass, next active scope can move from manual auth runtime proof to Laravel legacy API migration inventory.
