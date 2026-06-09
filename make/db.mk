##@ Database

.PHONY: db-dev-setup db-migrate db-status db-adopt-existing migrate-up migrate-down seed

db-dev-setup: ## Create or update the local PostgreSQL role and database from DATABASE_URL
	bash scripts/db_dev_setup.sh

db-migrate: ## Apply pending up migrations with tracking
	bash scripts/db_migrate.sh

db-status: ## Show migration status from schema_migrations
	bash scripts/db_status.sh

db-adopt-existing: ## Mark existing up migrations as applied
	bash scripts/db_adopt_existing.sh

migrate-up: db-migrate ## Alias to apply pending migrations

migrate-down: ## Roll back migrations one safe step when support is added
	@printf '%s\n' 'migrate-down is not implemented in this repository yet.' >&2
	@exit 1

seed: ## Load deterministic seed data when support is added
	@printf '%s\n' 'seed is not implemented in this repository yet.' >&2
	@exit 1
