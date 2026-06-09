##@ Build

.PHONY: build run dev

build: ## Build the API binary
	mkdir -p .bin
	go build -o $(APP_BIN) ./cmd/api

run: build ## Run the API on HTTP_PORT
	HTTP_PORT=$(HTTP_PORT) $(APP_BIN)

dev: db-dev-setup db-migrate run ## Setup local DB, migrate, then run the API
