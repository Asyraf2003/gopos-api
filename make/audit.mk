# Copyright (C) 2026 Asyraf Mubarak
#
# This file is part of gopos-api.
#
# gopos-api is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, version 3 only.
#
# gopos-api is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

##@ Audit

.PHONY: audit-ai-rules license-headers audit-license-headers audit-file-size audit-hex arch audit-route-capabilities security-gosec security audit-all

audit-ai-rules: ## Run AI rules audit
	bash scripts/audit_ai_rules.sh

license-headers: ## Apply license headers to supported source files
	bash scripts/apply_license_headers.sh
	fd -e go . -x gofmt -w

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
