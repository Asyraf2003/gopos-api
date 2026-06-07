SHELL := /usr/bin/env bash

APP_BIN := .bin/pos-go-api
HTTP_PORT ?= 8081

.PHONY: help fmt test vet audit-format audit-ai-rules audit-file-size audit-hex audit-route-capabilities security-gosec audit-all screening check verify ci build run dev smoke auth-login-admin auth-login-cashier auth-start db-dev-setup db-migrate db-status db-adopt-existing git-status push

help:
	@printf '%s\n' \
	'Available targets:' \
	'  make fmt               - run gofmt on all Go files' \
	'  make test              - run go test ./...' \
	'  make vet               - run go vet ./...' \
	'  make audit-format      - check gofmt cleanliness' \
	'  make audit-ai-rules    - run AI rules audit' \
	'  make audit-file-size   - run file size audit' \
	'  make audit-hex         - run strict hexagonal import-boundary audit' \
		'  make audit-route-capabilities - run protected route capability coverage audit' \
	'  make security-gosec    - run gosec security audit' \
	'  make audit-all         - run test + all audit scripts' \
	'  make screening         - alias to audit-all' \
	'  make check             - run format + vet + file size + hex + docs audits' \
	'  make verify            - run aggregate audit gate' \
	'  make ci                - alias to verify' \
	'  make build             - build app binary' \
	'  make run               - run app on HTTP_PORT (default 8081)' \
	'  make dev               - setup local DB, migrate, then run API' \
	'  make smoke             - call /api/health and unauthenticated /api/me' \
	'  make auth-login-admin  - manual login as admin and print bearer token' \
	'  make auth-login-cashier - manual login as cashier and print bearer token' \
	'  make auth-start        - start app, request Google auth URL, print it, and open browser' \
	'  make db-dev-setup      - create/update local PostgreSQL role and database from DATABASE_URL' \
	'  make db-migrate        - apply pending *.up.sql migrations with tracking' \
	'  make db-status         - show migration status from schema_migrations' \
	'  make db-adopt-existing - mark existing *.up.sql as applied in schema_migrations' \
	'  make git-status        - show git status short' \
	'  make push MSG="..."   - git add, commit with MSG, and push current branch'

fmt:
	gofmt -w $$(fd -e go .)

test:
	GOCACHE=$${GOCACHE:-/tmp/go-build-cache} go test ./...

vet:
	bash scripts/audit_go_vet.sh

audit-format:
	bash scripts/audit_format.sh

audit-ai-rules:
	bash scripts/audit_ai_rules.sh

audit-file-size:
	bash scripts/audit_file_size.sh

audit-hex:
	bash scripts/audit_hexagonal.sh

audit-route-capabilities:
	bash scripts/audit_route_capabilities.sh

security-gosec:
	bash scripts/audit_security_gosec.sh

audit-all:
	bash scripts/audit_all.sh

screening: audit-all

check: audit-format vet audit-file-size audit-hex audit-ai-rules

verify: audit-all

ci: verify

build:
	mkdir -p .bin
	go build -o $(APP_BIN) ./cmd/api

run: build
	HTTP_PORT=$(HTTP_PORT) $(APP_BIN)

dev: db-dev-setup db-migrate run

smoke:
	HTTP_PORT=$(HTTP_PORT) bash scripts/dev_smoke.sh

auth-login-admin:
	HTTP_PORT=$(HTTP_PORT) bash scripts/dev_manual_login.sh admin

auth-login-cashier:
	HTTP_PORT=$(HTTP_PORT) bash scripts/dev_manual_login.sh cashier

auth-start: build
	PORT_PIDS="$$(ss -ltnp | sed -n 's/.*:$(HTTP_PORT) .*pid=\([0-9]\+\).*/\1/p' | sort -u)"; \
	printf 'PORT_PIDS_BEFORE=%s\n' "$$PORT_PIDS"; \
	if [ -n "$$PORT_PIDS" ]; then kill -9 $$PORT_PIDS; fi; \
	sleep 1; \
	HTTP_PORT=$(HTTP_PORT) $(APP_BIN) > /tmp/pos-go-api-$(HTTP_PORT).log 2>&1 & \
	API_PID=$$!; \
	sleep 3; \
	AUTH_URL="$$(curl -s 'http://127.0.0.1:$(HTTP_PORT)/api/auth/google/start?purpose=login&redirect_url=http://127.0.0.1:$(HTTP_PORT)/api/auth/google/callback' | python3 -c 'import sys, json; print(json.load(sys.stdin)["redirect_to"])')"; \
	printf 'API_PID=%s\n' "$$API_PID"; \
	printf 'AUTH_URL=%s\n' "$$AUTH_URL"; \
	xdg-open "$$AUTH_URL" || true

db-dev-setup:
	bash scripts/db_dev_setup.sh

db-migrate:
	bash scripts/db_migrate.sh

db-status:
	bash scripts/db_status.sh


db-adopt-existing:
	bash scripts/db_adopt_existing.sh


git-status:
	git status --short

push:
	bash scripts/git_push.sh
