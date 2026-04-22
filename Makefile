SHELL := /usr/bin/env bash

APP_BIN := .bin/pos-go-api
HTTP_PORT ?= 8081

.PHONY: help test audit-ai-rules audit-file-size security-gosec audit-all screening build run auth-start db-migrate db-status db-adopt-existing git-status push

help:
	@printf '%s\n' \
	'Available targets:' \
	'  make test              - run go test ./...' \
	'  make audit-ai-rules    - run AI rules audit' \
	'  make audit-file-size   - run file size audit' \
	'  make security-gosec    - run gosec security audit' \
	'  make audit-all         - run test + all audit scripts' \
	'  make screening         - alias to audit-all' \
	'  make build             - build app binary' \
	'  make run               - run app on HTTP_PORT (default 8081)' \
	'  make auth-start        - start app, request Google auth URL, print it, and open browser' \
	'  make db-migrate        - apply pending *.up.sql migrations with tracking' \
	'  make db-status         - show migration status from schema_migrations' \
	'  make db-adopt-existing - mark existing *.up.sql as applied in schema_migrations' \
	'  make git-status        - show git status short' \
	'  make push MSG="..."   - git add, commit with MSG, and push current branch'

test:
	go test ./...

audit-ai-rules:
	bash scripts/audit_ai_rules.sh

audit-file-size:
	bash scripts/audit_file_size.sh

security-gosec:
	bash scripts/audit_security_gosec.sh

audit-all:
	bash scripts/audit_all.sh

screening: audit-all

build:
	mkdir -p .bin
	go build -o $(APP_BIN) ./cmd/api

run: build
	HTTP_PORT=$(HTTP_PORT) $(APP_BIN)

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
