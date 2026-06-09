##@ Test

.PHONY: test test-unit test-api test-db vet lint check verify ci screening

test: ## Run all Go tests
	$(GO_TEST) ./...

test-unit: ## Run module-focused tests
	$(GO_TEST) ./internal/modules/... ./internal/platform/state/... ./internal/platform/token/... ./internal/config

test-api: ## Run HTTP transport and presentation tests
	$(GO_TEST) ./internal/modules/*/transport/http ./internal/transport/http/... ./internal/presentation/http/...

test-db: ## Run PostgreSQL adapter tests
	$(GO_TEST) ./internal/platform/postgres/...

vet: ## Run go vet audit
	bash scripts/audit_go_vet.sh

lint: vet ## Run static analysis currently wired as go vet

check: audit-format vet audit-file-size audit-hex audit-ai-rules ## Run local doc and structure checks

verify: audit-all ## Run the aggregate local quality gate

ci: verify ## Alias to verify

screening: audit-all ## Alias to aggregate audit
