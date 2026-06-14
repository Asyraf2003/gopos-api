<!--
Copyright (C) 2026 Asyraf Mubarak

This file is part of gopos-api.

gopos-api is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, version 3 only.

gopos-api is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with gopos-api. If not, see <https://www.gnu.org/licenses/>.
-->

# Supplier PostgreSQL Persistence Handoff

## Active Scope

Supplier PostgreSQL persistence slice, migration-only through repository SetActive checkpoint.

## Status

In progress.

Migration-only step is locally implemented and applied.

Repository adapter skeletons are locally implemented with compile-time port assertion.

Repository Create, FindByID, FindByNormalizedName, FindActiveByNormalizedName, Update, and SetActive behavior are locally implemented with compile, targeted DB-backed integration, and aggregate proof.

List and Lookup remain explicit placeholder behavior.

## Files Changed

```text
docs/blueprints/0039_supplier_postgres_persistence_slice.md
migrations/0014_create_suppliers_table.up.sql
migrations/0014_create_suppliers_table.down.sql
internal/platform/postgres/supplier_repository.go
internal/platform/postgres/supplier_repository_row.go
internal/platform/postgres/supplier_repository_write.go
internal/platform/postgres/supplier_repository_query.go
internal/platform/postgres/supplier_repository_integration_helpers_test.go
internal/platform/postgres/supplier_repository_create_integration_test.go
internal/platform/postgres/supplier_repository_query_integration_test.go
internal/platform/postgres/supplier_repository_update_integration_test.go
internal/platform/postgres/supplier_repository_lifecycle_integration_test.go
docs/handoffs/2026-06-14-supplier-postgres-persistence-migration-only.md
docs/handoffs/README.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

## FACT

- Supplier implementation slice 1 was restored and closed locally with proof after fixing duplicate/split helper issues and aligning ActivateSupplier with FindActiveByNormalizedName.
- Supplier PostgreSQL persistence blueprint `0039_supplier_postgres_persistence_slice.md` is accepted by owner decision.
- The accepted step order says implementation must start with migration-only work.
- Migration slot `0014` was confirmed empty before creating Supplier migration files.
- Migration files created:
  - `migrations/0014_create_suppliers_table.up.sql`
  - `migrations/0014_create_suppliers_table.down.sql`
- Migration creates `suppliers`.
- Supplier table includes:
  - `id`
  - `name`
  - `name_normalized`
  - `phone`
  - `email`
  - `address`
  - `notes`
  - `is_active`
  - `created_at`
  - `updated_at`
- Supplier indexes include:
  - `suppliers_active_name_normalized_unique`
  - `suppliers_active_name_idx`
  - `suppliers_name_normalized_idx`
  - `suppliers_updated_at_idx`
- Repository adapter skeleton files created:
  - `internal/platform/postgres/supplier_repository.go`
  - `internal/platform/postgres/supplier_repository_row.go`
  - `internal/platform/postgres/supplier_repository_write.go`
  - `internal/platform/postgres/supplier_repository_query.go`
- `SupplierRepository` stores the existing PostgreSQL pool dependency pattern.
- `NewSupplierRepository` is the constructor.
- Compile-time assertion verifies the adapter satisfies `ports.SupplierRepository`.
- Query helpers follow the existing transaction-aware `TxFromContext` pattern.
- Row mapping structure targets the `suppliers` table columns from migration `0014`.
- Create inserts all Supplier fields into `suppliers`.
- FindByID returns `(supplier, true, nil)` when found and `(domain.Supplier{}, false, nil)` when missing.
- FindByNormalizedName returns a matching supplier by `name_normalized`, preferring active rows to keep active duplicate guards safe when inactive duplicates exist.
- FindActiveByNormalizedName filters by `name_normalized` and `is_active = true`.
- Update persists Supplier `name`, `name_normalized`, contact fields, `is_active`, and `updated_at` by primary key.
- Update follows ProductCatalog and ServiceCatalog local convention for missing ids: no explicit not-found error is returned when zero rows are affected.
- SetActive deactivates an existing Supplier and returns the stored inactive Supplier row.
- SetActive activates an existing Supplier and returns the stored active Supplier row.
- SetActive returns `(domain.Supplier{}, false, nil)` for missing Supplier ids.
- SetActive updates `updated_at` when the active state changes.
- Activating an inactive Supplier is rejected by the PostgreSQL partial unique index when another active Supplier already owns the same normalized name.
- List and Lookup remain explicit placeholder behavior.
- Supplier repository integration test files were added for Create and direct find behavior.
- Supplier repository integration tests were added for Update behavior.
- Supplier repository integration tests were added for SetActive behavior.

The active normalized-name uniqueness rule uses a PostgreSQL partial unique index:

```sql
create unique index suppliers_active_name_normalized_unique
on suppliers (name_normalized)
where is_active = true;
```

This preserves the accepted duplicate policy:

- active Supplier name must be unique by normalized name;
- inactive Supplier names do not block active Supplier name reuse;
- reactivating inactive Supplier rejects if another active Supplier already owns the same normalized name.

## Proof Collected

Supplier module proof:

```bash
go test ./internal/modules/supplier/...
```

Visible result:

```text
ok   pos-go/internal/modules/supplier/domain
?    pos-go/internal/modules/supplier/ports [no test files]
ok   pos-go/internal/modules/supplier/usecase
```

Hexagonal proof:

```bash
bash scripts/audit_hexagonal.sh
```

Visible result:

```text
[PASS] hexagonal import audit passed
```

Aggregate proof:

```bash
make verify
```

Visible result:

```text
[PASS] aggregate audit passed
```

Supplier PostgreSQL skeleton compile proof:

```bash
go test ./internal/platform/postgres/... -run Supplier
```

Visible result:

```text
?    pos-go/internal/platform/postgres [no test files]
```

Supplier PostgreSQL integration-tagged proof:

```bash
set -a
source .env
set +a
go test -tags integration ./internal/platform/postgres/... -run Supplier -count=1 -v
```

Visible result:

```text
--- PASS: TestSupplierRepository_CreateStoresFields
--- PASS: TestSupplierRepository_CreateRejectsDuplicateActiveNormalizedName
--- PASS: TestSupplierRepository_CreateAllowsInactiveNameReuse
--- PASS: TestSupplierRepository_SetActiveDeactivatesSupplier
--- PASS: TestSupplierRepository_SetActiveActivatesSupplier
--- PASS: TestSupplierRepository_SetActiveMissing
--- PASS: TestSupplierRepository_SetActiveRejectsDuplicateActivation
--- PASS: TestSupplierRepository_FindByIDMissing
--- PASS: TestSupplierRepository_FindByNormalizedName
--- PASS: TestSupplierRepository_FindActiveByNormalizedName
--- PASS: TestSupplierRepository_FindActiveByNormalizedNameIgnoresInactive
--- PASS: TestSupplierRepository_UpdateChangesFields
--- PASS: TestSupplierRepository_UpdateStoresNormalizedNameFromDomain
--- PASS: TestSupplierRepository_UpdateMissingSupplierUsesLocalConvention
--- PASS: TestSupplierRepository_UpdateRejectsDuplicateActiveNormalizedName
PASS
ok   pos-go/internal/platform/postgres
```

Migration proof:

```bash
bash scripts/db_migrate.sh
```

Visible result:

```text
[APPLY] 0014_create_suppliers_table.up.sql
BEGIN
CREATE TABLE
CREATE INDEX
CREATE INDEX
CREATE INDEX
CREATE INDEX
INSERT 0 1
COMMIT

[PASS] db migrate completed
```

## Tests Run

```bash
go test ./internal/modules/supplier/...
go test ./internal/platform/postgres/... -run Supplier
set -a
source .env
set +a
go test -tags integration ./internal/platform/postgres/... -run Supplier -count=1 -v
bash scripts/audit_hexagonal.sh
make verify
bash scripts/db_migrate.sh
```

## Open Gaps

- Supplier PostgreSQL repository List and Lookup behavior is not implemented.
- Supplier query-plan proof is not collected.
- Supplier HTTP runtime is not implemented.
- Supplier capability seed is not implemented.
- Supplier route capability manifest rows are not implemented.
- Faktur is not implemented.
- Inventory/stock movement is not implemented.
- Audit/outbox persistence is not implemented.
- Localization is not implemented.
- Extended filters are not implemented.
- Laravel Supplier MySQL/source parity remains unproven.
- Remote connector validation for final local Supplier PostgreSQL changes remains pending unless the files become visible through connector.

## Next Valid Active Step

Supplier PostgreSQL repository List and Lookup behavior.

Keep the next step limited to List and Lookup behavior plus the focused repository proof needed for that behavior.

## Scope Guard

Do not start Supplier HTTP transport.

Do not add Supplier capability seed.

Do not start Faktur.

Do not start inventory mutation.

Do not start stock movement.

Do not start audit/outbox.

Do not start localization.

Do not start extended filters.

Do not start architecture folder cleanup.

Auth/System ADR 0012 output contract centralization remains deferred by owner decision and must not block Supplier/Faktur progress.

## Estimated Scope Progress Percentage

Supplier PostgreSQL persistence slice: 65%.

Reason:

- blueprint accepted;
- migration files created;
- migration applied locally;
- repository adapter skeleton files created;
- compile-time port assertion exists;
- Create, FindByID, FindByNormalizedName, FindActiveByNormalizedName, Update, and SetActive behavior implemented;
- focused Supplier and PostgreSQL compile proof passed;
- targeted Supplier DB-backed integration proof passed with `.env` loaded;
- aggregate `make verify` proof passed;
- integration tests for Create, direct lookup, Update, and SetActive behavior were added;
- List and Lookup remain pending;
- query-plan proof not collected.

## Estimated Context-Window Status

Current context is sufficient to start Supplier PostgreSQL repository List and Lookup behavior in the next session.

Recommended next session target:

```text
Terminal Codex
```

Recommended template source:

```text
docs/templates/0121_codex_session_prompts.md
```
