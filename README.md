# pos-go API

Go API service for POS backend development.

## Requirements

- Go
- PostgreSQL
- `psql`
- Bash
- Make

## Local environment

Create local env from the tracked example:

```bash
cp .env.example .env
```

`.env` is local-only and must not be committed.

Default local database values from `.env.example`:

```env
DATABASE_URL=postgres://posgo_app:posgo_local_dev_123@127.0.0.1:5432/posgo_app_db?sslmode=disable
```

## Local PostgreSQL setup

Start PostgreSQL first.

On systems using systemd:

```bash
systemctl status postgresql
```

Then create/update the local app role and database from `DATABASE_URL`:

```bash
make db-dev-setup
```

If your PostgreSQL admin user is not `postgres`, override it:

```bash
POSTGRES_ADMIN_USER=<admin_user> make db-dev-setup
```

## Migrations

Apply pending migrations:

```bash
make db-migrate
```

Show migration status:

```bash
make db-status
```

## Run API

```bash
make run
```

The Makefile default port is `8081`.

To override:

```bash
HTTP_PORT=8080 make run
```

## Common local flow

```bash
cp .env.example .env
make db-dev-setup
make db-migrate
make run
```

## Auth debug mode

Manual auth login is intended for local/build/testing only.

Enable it in `.env` when needed:

```env
AUTH_DEBUG_ENABLED=true
```

Then run:

```bash
make run
```

## Developer commands

Run the full local development flow:

```bash
make dev
```

This runs:

```bash
make db-dev-setup
make db-migrate
make run
```

Smoke-test a running local API:

```bash
make smoke
```

Manual auth login helpers:

```bash
make auth-login-admin
make auth-login-cashier
```

These commands call `POST /api/auth/manual/login`, print the HTTP response, extract the bearer token, then call `GET /api/me` with that token.

Manual auth requires:

```env
AUTH_DEBUG_ENABLED=true
```

Health check path:

```text
/api/health
```
