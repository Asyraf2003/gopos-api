##@ Audit

.PHONY: audit-ai-rules audit-license-headers audit-file-size audit-hex arch audit-route-capabilities security-gosec security audit-all

audit-ai-rules: ## Run AI rules audit
	bash scripts/audit_ai_rules.sh

audit-license-headers: ## Run license header audit
	bash scripts/audit_license_headers.sh

audit-file-size: ## Run file size audit
	bash scripts/audit_file_size.sh

audit-hex: ## Run strict hexagonal import-boundary audit
	bash scripts/audit_hexagonal.sh

arch: audit-hex ## Alias to architecture audit

audit-route-capabilities: ## Run protected route capability coverage audit
	bash scripts/audit_route_capabilities.sh

security-gosec: ## Run gosec security audit
	bash scripts/audit_security_gosec.sh

security: security-gosec ## Alias to security audit

audit-all: ## Run the aggregate audit script
	bash scripts/audit_all.sh
