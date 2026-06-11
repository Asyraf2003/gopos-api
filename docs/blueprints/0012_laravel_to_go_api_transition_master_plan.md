// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

// audit:allow-oversize reason=bootstrap-wiring

# Laravel To Go API Transition Master Plan

## FACT
- Current target stack is Go, Echo, PostgreSQL, pure API, and hexagonal ports/adapters.
- Current Go repo already has foundation auth/account/system code, PostgreSQL auth migrations, JWT/token platform code, request middleware, and centralized presentation direction.
- Current Go repo does not yet have POS business domains such as product catalog, inventory, procurement, note/transaction, payment, expense, employee finance, reporting, push notification, or mobile API.
- User-provided Laravel tree shows mature production modules:
  - `ProductCatalog`
  - `Procurement`
  - `Inventory`
  - `Note`
  - `Payment`
  - `EmployeeFinance`
  - `Expense`
  - `Reporting`
  - `Audit`
  - `IdentityAccess`
  - `MobileApi`
  - `PushNotification`
  - `ServiceCatalog`
- User direction: auth has its own modular field; this blueprint must focus on moving production Laravel behavior into the Go API.
- User direction: Go implementation should be more modular and mature than the Laravel production app, with complete Makefile, testing, language/output control, notification module, and full custom control.
- Stage 0 batch 1 evidence exists at `docs/evidence/0001_laravel_stage0_schema_and_route_inventory.md`.
- Stage 0 batch 2 catalog evidence exists at `docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md`.
- First catalog migration blueprint exists at `docs/blueprints/0020_catalog_foundation_migration.md`.

## GAP
- The Laravel source tree and initial migration/route batch were provided as command output, but the Go repo does not currently contain the Laravel source files for direct inspection.
- Exact Laravel route definitions, database columns, constraints, indexes, seed profiles, and production data edge cases must be inspected before implementation.
- Several Laravel alter/index/FK/timestamp migrations were not captured by the Stage 0 batch 1 `Schema::create` extraction.
- Product duplicate behavior needs owner decision before PostgreSQL unique indexes are finalized.
- Exact public API contract for the future UI/mobile clients is not accepted yet.
- PostgreSQL target schema for POS domains is not accepted yet.
- Capability keys for POS business operations are not defined yet.

## DECISION
- Do not port Laravel files mechanically.
- Migrate behavior by bounded context and domain contract, not by controller file order.
- Start with migration inventory, PostgreSQL schema parity design, and platform proof before business CRUD.
- Keep auth as a separate foundation module; do not block the migration plan on rebuilding auth again.
- Capability control remains a mandatory foundation before protected POS endpoints are exposed.
- The first POS domains should be low-risk master/read surfaces before high-risk transactional flows.
- High-risk financial, stock, note revision, payment, refund, procurement, payroll, and reporting behavior must move only after characterization tests and schema contracts exist.
- Blade/server-rendered UI does not move into this repo. Go owns API contracts only.

## SCOPE-IN
- Define Laravel-to-Go migration sequence.
- Define target Go module map.
- Define PostgreSQL-first readiness work.
- Define Makefile/script/test gates needed before domain migration.
- Define language/output contract direction.
- Define notification module boundary.
- Define per-domain blueprint order.
- Define proof required before moving from one stage to the next.

## SCOPE-OUT
- Implementing POS business code.
- Rewriting auth in this blueprint.
- Building Blade/UI screens in Go.
- Copying Laravel service providers, controllers, or Eloquent models directly.
- Migrating production data without a data migration blueprint.
- Finalizing every endpoint path.

## SOURCE LARAVEL MAP
Laravel source areas from the provided tree should be interpreted as follows:

```text
app/Core/*                 -> Go module domain packages
app/Application/*          -> Go module usecase packages
app/Ports/*                -> Go module ports packages
app/Adapters/In/Http/*     -> Echo HTTP adapters and API presenters
app/Adapters/Out/*         -> PostgreSQL/platform adapters
database/migrations/*      -> PostgreSQL target schema design inputs
database/seeders/CreateOnly -> Go seed profiles and fixture generators
routes/api.php             -> API contract candidates
routes/web/*               -> behavior inventory only, not Go web routes
resources/views/*          -> API output and UI-consumption hints only
tests/Feature/*            -> characterization and contract test inputs
tests/Unit/*               -> domain/application test inputs
mk/* and scripts/*         -> Makefile and audit command inputs
lang/id/*                  -> API output language catalog inputs
```

## TARGET GO MODULE MAP
Use `internal/modules/<module>/` with `domain`, `ports`, `usecase`, and `transport/http` where applicable.

Initial module map:

```text
account              existing auth/account administration boundary
auth                 existing authentication/session boundary
capability           API capability registry and runtime policy
audit                audit event/outbox/read model boundary
language             API message/catalog/output language boundary
notification         push notification and future channel adapters
servicecatalog       service catalog master data
productcatalog       product master, lifecycle, lookup, soft delete, versioning
inventory            stock movements, projection, costing
procurement          supplier, supplier invoice, receipt, payable, payment proof
payment              customer payment, allocation, refund, settlement
note                 transaction workspace, note, work item, revision, correction
expense              expense category and operational expense
employeefinance      employee, debt, payroll, reversal
reporting            read models, dashboards, exports-as-data contracts
mobileapi            mobile-facing API contracts after core domains exist
system               health and infrastructure endpoints
```

## TARGET PLATFORM MAP
Platform and cross-cutting packages should stay outside business modules:

```text
internal/platform/postgres       PostgreSQL pool, tx, adapters, migration helpers
internal/platform/token          JWT/token primitives
internal/platform/notification   web push, future Telegram/WhatsApp/email adapters
internal/platform/i18n           catalog loading if runtime localization is needed
internal/transport/http/middleware authn/authz/capability/rate-limit/request-id/recovery
internal/presentation/http/id    Indonesian API output contracts
internal/presentation/http/en    future English API output contracts
scripts/                         audit, migration, seed, architecture, contract checks
```

## FOUNDATION ORDER
The migration should start with foundations, not with PostgreSQL alone and not with business CRUD alone.

### Stage 0 - Source Inventory And Parity Matrix
Goal: know exactly what is moving.

Deliverables:
- route inventory from Laravel `routes/api.php` and `routes/web/*`;
- domain inventory from Laravel `app/Core`, `app/Application`, `app/Ports`, and `app/Adapters/Out`;
- table/index/constraint inventory from Laravel `database/migrations`;
- seed profile inventory from `database/seeders/CreateOnly`;
- test inventory from Laravel `tests/Feature` and `tests/Unit`;
- risk ranking per domain.

Proof:
- generated inventory docs under `docs/evidence/`;
- current partial proof: `docs/evidence/0001_laravel_stage0_schema_and_route_inventory.md`;
- current catalog proof: `docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md`;
- no implementation yet.

### Stage 1 - Go Quality Foundation
Goal: make the Go repo able to prove work consistently.

Deliverables:
- complete Makefile aliases aligned with `docs/scripts/0090_makefile_and_scripts.md`;
- `make verify`, `make test`, `make test-unit`, `make test-api`, `make test-db`, `make lint`, `make fmt`, `make arch`, `make security`, `make migrate-up`, `make migrate-down`, `make seed`;
- architecture audit for forbidden imports and package boundaries;
- route-to-capability audit placeholder;
- API envelope audit placeholder;
- migration naming/status checks.

Proof:
- `make verify` runs the available gates;
- missing optional external services are reported clearly, not silently skipped.

### Stage 2 - PostgreSQL Target Baseline
Goal: design PostgreSQL correctly before porting behavior.

Deliverables:
- accepted PostgreSQL schema blueprint for POS domains;
- MySQL-to-PostgreSQL type mapping notes;
- integer rupiah money policy across all finance tables;
- `timestamptz` operational/system timestamp policy;
- FK/check/index policy per table;
- migration rollback/compensation notes;
- seed profile design: `minimal`, `dev`, `stress`, `security`.

Proof:
- migration dry run against local PostgreSQL;
- schema inspection queries;
- DB tests for constraints where implemented.

### Stage 3 - API Foundation And Capability Control
Goal: expose protected APIs only through stable contracts.

Deliverables:
- capability registry and middleware from `0010_capability_control_foundation.md`;
- centralized response/error presenters per language;
- request validation patterns;
- public DTO naming and envelope rules;
- route registration convention.

Proof:
- disabled capability returns `403` before validation/usecase;
- API envelope tests;
- route capability audit.

### Stage 4 - Cross-Cutting Modules
Goal: build shared behavior once, not inside every domain.

Deliverables:
- `audit` module: event writer, outbox decision, read model boundary;
- `language` module: stable message keys and `internal/presentation/http/id` ownership;
- `notification` module: push subscription, notification sender ports, future channel adapters;
- `idempotency` support: store, conflict semantics, retry policy;
- `transaction` support: PostgreSQL transaction manager;
- `servicecatalog` module if needed by note/workspace flows.

Proof:
- focused unit tests;
- PostgreSQL adapter tests where storage exists;
- no domain module imports platform directly.

## BUSINESS MIGRATION ORDER
Recommended order moves from low-risk foundations to high-risk transaction flows.

### Phase 1 - Master Data And Lookups
Order:
1. `servicecatalog`
2. `productcatalog`
3. `supplier` inside `procurement`
4. `expense` category

Reason:
- relatively isolated;
- creates lookup surfaces needed by procurement and note;
- validates CRUD pattern, soft delete/versioning, table queries, capability keys, and API envelopes.

Required per domain:
- domain contract;
- PostgreSQL migration;
- ports/usecase/adapters;
- list/show/create/update/delete or explicit forbidden operation;
- capability keys;
- API tests and DB tests;
- seed fixture.

### Phase 2 - Inventory Foundation
Order:
1. inventory product projection;
2. movement writer;
3. costing projection;
4. stock adjustment and reversal.

Reason:
- procurement and note flows depend on stock correctness;
- stock mutation needs lock, audit, and negative stock policy before transaction flows.

Required proof:
- movement invariants;
- projection rebuild tests;
- negative stock policy tests;
- idempotent reversal tests.

### Phase 3 - Procurement
Order:
1. supplier invoice create/list/show;
2. receive invoice and stock impact;
3. supplier payment;
4. payment proof attachment metadata;
5. reversal/void/revision.

Reason:
- procurement exercises product, supplier, inventory, payment proof, audit, and reporting source data before customer transaction complexity.

Required proof:
- duplicate invoice guards;
- product-per-revision constraints;
- receipt/payment/reversal transaction tests;
- attachment output security tests;
- capability-disabled tests.

### Phase 4 - Customer Transaction Core
Order:
1. note and work item create;
2. transaction workspace draft;
3. service/product/external/store-stock line behavior;
4. payment allocation;
5. note close/open state;
6. note history projection.

Reason:
- this is the central POS workflow and should start only after master data, inventory, audit, idempotency, and API foundation exist.

Required proof:
- create transaction idempotency;
- inventory issue/reversal;
- payment allocation correctness;
- rollback tests;
- current revision/history projection tests;
- API contract tests.

### Phase 5 - Corrections, Refunds, Revisions
Order:
1. paid work item correction;
2. note revision settlement;
3. surplus/refund due;
4. refund paid execution;
5. reopen/close correction policies.

Reason:
- Laravel tree shows many hardening tests and historical fixes here; this is high risk and must follow characterization.

Required proof:
- concurrency tests;
- over-allocation protection;
- current-only refund boundary;
- audit timeline proof;
- reporting impact proof.

### Phase 6 - Employee Finance And Expense
Order:
1. employee master;
2. employee debt;
3. debt payment/reversal;
4. payroll disbursement/reversal;
5. operational expense.

Reason:
- financially important, but less coupled to POS transaction engine than note/payment/inventory.

Required proof:
- non-negative balances;
- reversal invariants;
- audit events;
- report source consistency.

### Phase 7 - Reporting And Dashboard
Order:
1. read model source contracts;
2. transaction summary;
3. cash ledger;
4. inventory stock value;
5. operational profit;
6. supplier payable;
7. payroll and employee debt;
8. dashboard datasets.

Reason:
- reporting should consume accepted write models, not define business truth.

Required proof:
- reconciliation tests against source tables;
- contract tests for API outputs;
- export data contracts if exports remain API-owned.

### Phase 8 - Mobile API
Order:
1. mobile read contracts;
2. product search;
3. procurement invoice read;
4. upload/payment proof APIs;
5. mobile-specific auth/session integration.

Reason:
- mobile should consume stable core APIs and not force core domain shortcuts.

Required proof:
- mobile API envelope tests;
- auth/capability tests;
- upload safety tests.

### Phase 9 - Notification
Order:
1. push subscription API;
2. due note reminder notification;
3. supplier payable reminder notification;
4. future Telegram/WhatsApp/email adapters.

Reason:
- notification is modular and should consume domain events/read models, not own transaction truth.

Required proof:
- sender port tests;
- payload contract tests;
- no secret/log leakage;
- retry/failure behavior.

## API SURFACE POLICY
- Laravel `web` routes are behavior inventory only.
- Go exposes `/api/...` only.
- Any future UI consumes:
  - API response DTOs;
  - capability metadata;
  - language/output contracts;
  - pagination/filter metadata.
- No server-rendered page controller is ported to Go.
- Table-data endpoints should become list/query APIs with stable filters, sorting, pagination, and capability keys.

## LANGUAGE AND OUTPUT POLICY
- Indonesian output contracts live in `internal/presentation/http/id`.
- English output contracts may later live in `internal/presentation/http/en`.
- Message keys and machine states must be stable and separate from human text.
- Domain/usecase packages return typed errors/states, not localized strings.
- Handlers call presenters/output mappers for user-facing response shapes.

## NOTIFICATION POLICY
- Notification is its own module.
- Domain modules may emit events or expose read models.
- Notification module decides subscription, channel, template, and send policy.
- External providers are adapters under `internal/platform/notification`.
- No domain package imports web push, Telegram, WhatsApp, email, or provider SDKs.

## MAKEFILE AND SCRIPT POLICY
Before POS domain work accelerates, Makefile should converge to:

```bash
make verify
make test
make test-unit
make test-api
make test-db
make lint
make fmt
make arch
make security
make migrate-up
make migrate-down
make seed
make seed-minimal
make seed-dev
make seed-stress
make seed-security
make routes-audit
make api-contract-audit
make migration-audit
```

Existing `make audit-all` may stay as an alias, but `make verify` should become the canonical full gate.

## BLUEPRINT POLICY PER DOMAIN
Each domain migration must get its own blueprint before code:

```text
docs/blueprints/0020_servicecatalog_migration.md
docs/blueprints/0021_productcatalog_migration.md
docs/blueprints/0022_inventory_foundation.md
docs/blueprints/0023_procurement_migration.md
docs/blueprints/0024_note_transaction_foundation.md
docs/blueprints/0025_payment_allocation_refund.md
docs/blueprints/0026_employee_finance_migration.md
docs/blueprints/0027_reporting_read_models.md
docs/blueprints/0028_notification_module.md
docs/blueprints/0029_mobile_api_contracts.md
```

Each domain blueprint must include:
- Laravel source files to inspect;
- tables and migrations involved;
- accepted API routes;
- capability keys;
- domain lifecycle operations;
- transaction/lock/idempotency/audit policy;
- test mapping from Laravel tests to Go tests;
- migration proof and rollback proof.

## FIRST ACTIVE IMPLEMENTATION STEP
Do not start with business CRUD.

Next valid active implementation step:

```text
Resolve the ProductCatalog duplicate policy in `docs/blueprints/0020_catalog_foundation_migration.md`,
then implement catalog foundation in the blueprint order.
```

## DOD
- This master blueprint exists.
- It defines whether PostgreSQL or business code comes first.
- It separates auth from POS migration while preserving auth/capability as foundation.
- It defines target Go modules.
- It defines Makefile, testing, language, notification, API, and DB directions.
- It defines per-domain migration order.
- It leaves exactly one next active implementation step.
