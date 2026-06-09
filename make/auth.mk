##@ Auth And Smoke

.PHONY: smoke auth-login-admin auth-login-cashier auth-start

smoke: ## Call /api/health and unauthenticated /api/me
	HTTP_PORT=$(HTTP_PORT) bash scripts/dev_smoke.sh

auth-login-admin: ## Manual login as admin and print bearer token
	HTTP_PORT=$(HTTP_PORT) bash scripts/dev_manual_login.sh admin

auth-login-cashier: ## Manual login as cashier and print bearer token
	HTTP_PORT=$(HTTP_PORT) bash scripts/dev_manual_login.sh cashier

auth-start: build ## Start app, print Google auth URL, and open browser
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
