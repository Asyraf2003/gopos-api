##@ Format

.PHONY: fmt audit-format

fmt: ## Format all Go files with gofmt
	gofmt -w $$(fd -e go .)

audit-format: ## Check gofmt cleanliness
	bash scripts/audit_format.sh
