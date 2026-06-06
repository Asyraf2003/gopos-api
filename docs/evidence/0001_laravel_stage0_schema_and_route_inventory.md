# Laravel Stage 0 Schema And Route Inventory

## Source

User-provided Laravel command output from the Laravel root:

```bash
for f in $(fd . database/migrations -e php | sort); do
    printf '\n===== %s =====\n' "$f"
    sed -n '/Schema::create(/,/^[[:space:]]*});/p' "$f"
done

cat routes/web.php
for f in $(fd . routes/web -e php | sort); do
    cat "$f"
done
```

`routes/api.php` was intentionally skipped by the user because it was experimental and not the source of truth for the future Go API.

## What This Evidence Proves

- Laravel production behavior is mostly web-route driven, but the Go target must expose API-only contracts.
- The database contains enough mature POS schema information to start a PostgreSQL parity matrix.
- PostgreSQL work cannot be finalized from this output alone because many `alter`, foreign key, index, timestamp, and seed migrations were not captured by the `Schema::create` extraction.
- The first Go API migration work should be inventory and schema design, not direct implementation.

## High-Level Bounded Contexts

| Context | Laravel Evidence | Target Go Module |
| --- | --- | --- |
| Auth/Account | `users`, sessions, web auth routes | existing `auth`, `account` |
| Identity Access/Capability | `actor_accesses`, admin capability state routes | `capability`, existing authz |
| Audit | `audit_logs`, `audit_events`, `audit_event_snapshots`, `audit_outbox` | `audit` |
| Product Catalog | `products`, `product_versions`, product routes | `productcatalog` |
| Supplier/Procurement | suppliers, invoices, receipts, payments, payment proof routes | `procurement` |
| Inventory | movements, product inventory, costing, cost adjustments | `inventory` |
| Note/Transaction | notes, work items, revisions, workspace drafts, correction routes | `note` |
| Payment/Refund | customer payments, allocations, refunds, component allocations | `payment` |
| Employee Finance | employees, debts, payroll, reversals, adjustments | `employeefinance` |
| Expense | categories, operational expenses | `expense` |
| Reporting | report routes and projection tables | `reporting` |
| Notification | push subscriptions | `notification` |
| Mobile API | mobile tokens table only in this batch | `mobileapi` |
| Service Catalog | service catalog items | `servicecatalog` |

## Table Inventory By Module

### Existing/Platform-Like Tables

- `users`
- `password_reset_tokens`
- `sessions`
- `cache`
- `cache_locks`
- `jobs`
- `job_batches`
- `failed_jobs`

Go decision:
- Do not automatically port Laravel cache/jobs/session tables.
- `users` is relevant only as auth/account reference.
- Go should keep its own auth/session/token tables already started in current migrations unless a separate auth blueprint changes that.

### Identity Access And Capability

- `actor_accesses`
- `admin_transaction_capability_states`
- `admin_cashier_area_access_states`

Go decision:
- Treat legacy admin/cashier capability state as source behavior evidence.
- Replace with a first-class API capability registry/control surface in `capability`.
- Do not expose protected POS operations before capability keys exist.

### Audit

- `audit_logs`
- `note_mutation_events`
- `note_mutation_snapshots`
- `audit_events`
- `audit_event_snapshots`
- `audit_outbox`

PostgreSQL notes:
- JSON columns should become `jsonb` where machine-readable.
- `occurred_at`, `created_at`, `updated_at`, `available_at`, `locked_at`, `processed_at` should become `timestamptz`.
- `audit_outbox(status, available_at)` is a queue hot path and needs index proof.

### Product Catalog

- `products`
- `product_versions`

Observed fields:
- product id, code, name, brand, size, selling price;
- soft delete and search/duplicate hardening exist in later alter migrations but are not captured here;
- product versions preserve revision snapshots.

PostgreSQL notes:
- money fields should be `bigint` or constrained integer rupiah.
- active uniqueness and soft delete behavior require full alter migration files before final schema.
- version snapshots should use `jsonb` only if queried; otherwise store as controlled historical payload by decision.

### Supplier And Procurement

- `suppliers`
- `supplier_versions`
- `supplier_invoices`
- `supplier_invoice_lines`
- `supplier_receipts`
- `supplier_receipt_lines`
- `supplier_payments`
- `supplier_payment_proof_attachments`
- `supplier_invoice_versions`
- `supplier_receipt_reversals`
- `supplier_payment_reversals`
- `supplier_invoice_list_projection`
- `supplier_list_projection`

Observed behavior:
- invoice create/update/revise/void;
- receive invoice;
- record/reverse supplier payment;
- attach/serve payment proof;
- supplier lookup and table queries;
- list projections for invoice and supplier read performance.

PostgreSQL notes:
- add FKs from invoice -> supplier, lines -> invoice/product, receipts -> invoice, receipt lines -> receipt/invoice line, payments -> invoice, proof -> payment.
- preserve reversal uniqueness constraints.
- projection tables are valid for `<1s` list/report goals but need rebuild/sync proof.

### Inventory

- `inventory_movements`
- `product_inventory`
- `product_inventory_costing`
- `inventory_cost_adjustments`

Observed behavior:
- stock movement writer;
- projection of current quantity;
- costing projection;
- stock adjustment and reversal routes via product admin.

PostgreSQL notes:
- use `check` constraints for non-negative projected quantity where policy requires.
- movement `qty_delta` can be signed.
- movement source lookup needs `(source_type, source_id)` index.
- unique reversal source key exists in an omitted alter migration and must be inspected.

### Note And Transaction

- `notes`
- `work_items`
- `work_item_service_details`
- `work_item_external_purchase_lines`
- `work_item_store_stock_lines`
- `transaction_workspace_drafts`
- `note_history_projection`
- `note_revisions`
- `note_revision_lines`
- `note_revision_settlements`
- `note_revision_surplus_dispositions`
- `note_revision_surplus_refund_payments`

Observed behavior:
- create note;
- create/update transaction workspace;
- workspace drafts;
- add rows;
- payments/refunds;
- corrections;
- admin and cashier note histories;
- revisions, settlements, surplus refund due, refund paid;
- reopen closed notes.

PostgreSQL notes:
- this is high-risk and must not be first implementation domain.
- use strict transaction boundaries and idempotency for create/update/payment/refund.
- projection table `note_history_projection` is important for `<1s` list/history behavior.
- current revision pointer and line uniqueness were added by omitted alter migrations and must be inspected.

### Payment And Refund

- `customer_payments`
- `customer_payment_cash_details`
- `payment_allocations`
- `customer_refunds`
- `payment_component_allocations`
- `refund_component_allocations`

Observed behavior:
- note payments;
- closed note refunds;
- component allocation;
- refund allocation;
- cash details and change calculation.

PostgreSQL notes:
- all money must be integer rupiah with non-negative checks where appropriate.
- over-allocation prevention needs DB constraints plus transactional locks, not only application code.
- component allocation uniqueness exists and must be preserved.

### Employee Finance

- `employees`
- `employee_versions`
- `employee_debts`
- `employee_debt_payments`
- `employee_debt_payment_reversals`
- `employee_debt_adjustments`
- `payroll_disbursements`
- `payroll_disbursement_reversals`

Observed behavior:
- employee create/update/show/list;
- payroll list/create/reverse;
- employee debt create/pay/adjust/reverse.

PostgreSQL notes:
- use `bigint` money fields with non-negative checks.
- reversal tables use one-to-one uniqueness and restrict delete.
- versioning exists for employee master.

### Expense

- `expense_categories`
- `operational_expenses`

Observed behavior:
- category create/update/activate/deactivate/list;
- operational expense create/list/soft delete.

PostgreSQL notes:
- `operational_expenses` uses soft delete.
- category delete appears replaced by activation/deactivation.
- operational expense delete should be soft-delete only unless a domain contract says otherwise.

### Notification And Mobile

- `push_subscriptions`
- `mobile_api_tokens`

Go decision:
- notification must be modular and consume events/read models.
- mobile API contracts should come after stable core modules.
- token/auth details remain in auth/mobileapi blueprints, not this schema evidence.

### Service Catalog

- `service_catalog_items`

Observed behavior:
- service lookup and create route appears under note workspace routes.

Go decision:
- good candidate for the first POS business domain because it is small and useful for transaction workspace.

## Route Inventory By Target API Module

Laravel web routes should become API contracts only where behavior is needed. Page routes become read/show/list API candidates or UI-only references.

### Health/System

- `GET /health`

Go target:
- existing or future `GET /api/health`.
- public infrastructure; capability may be exempt.

### Auth

- `GET /login`
- `POST /login`
- `POST /logout`

Go target:
- auth remains its own module.
- do not use Laravel web auth routes as future API source of truth.

### Capability/Identity Access

- `POST /identity-access/admin-transaction-capability/enable`
- `POST /identity-access/admin-transaction-capability/disable`

Go target:
- replace with admin capability API:
  - `GET /api/admin/capabilities`
  - `GET /api/admin/capabilities/:key`
  - `POST /api/admin/capabilities/:key/enable`
  - `POST /api/admin/capabilities/:key/disable`

### Product Catalog

Laravel routes:
- `GET /admin/products/table`
- `POST /admin/products/{productId}/stock-adjustments`
- `PATCH /admin/products/{productId}/stock-adjustments/{adjustmentId}/reverse`
- `GET /admin/products`
- `GET /admin/products/create`
- `POST /admin/products`
- `GET /admin/products/{productId}`
- `GET /admin/products/{productId}/edit`
- `GET /admin/products/{productId}/stock`
- `PUT /admin/products/{productId}`
- `PATCH /admin/products/{productId}/restore`
- `DELETE /admin/products/{productId}`
- `POST /product-catalog/products/create`
- `POST /product-catalog/products/{productId}/update`

Go API candidates:
- `GET /api/products`
- `POST /api/products`
- `GET /api/products/:id`
- `PUT /api/products/:id`
- `DELETE /api/products/:id`
- `PATCH /api/products/:id/restore`
- `POST /api/products/:id/stock-adjustments`
- `PATCH /api/products/:id/stock-adjustments/:adjustmentId/reverse`
- `GET /api/products/lookup`

### Procurement And Supplier

Laravel routes:
- supplier list/table/edit/update;
- procurement invoice table/list/create/show/edit/revise/update;
- receive invoice;
- record payment;
- reverse receipt;
- reverse payment;
- void invoice;
- attach/serve payment proof;
- product and supplier lookup;
- non-admin transaction-entry routes for create/receive.

Go API candidates:
- `GET /api/suppliers`
- `PUT /api/suppliers/:id`
- `GET /api/suppliers/lookup`
- `GET /api/procurement/supplier-invoices`
- `POST /api/procurement/supplier-invoices`
- `GET /api/procurement/supplier-invoices/:id`
- `PUT /api/procurement/supplier-invoices/:id`
- `POST /api/procurement/supplier-invoices/:id/receive`
- `POST /api/procurement/supplier-invoices/:id/payments`
- `POST /api/procurement/supplier-invoices/:id/void`
- `POST /api/procurement/supplier-receipts/:id/reverse`
- `POST /api/procurement/supplier-payments/:id/reverse`
- `POST /api/procurement/supplier-payments/:id/proof`
- `GET /api/procurement/supplier-payment-proof-attachments/:id`

### Note, Transaction, Payment, Refund

Laravel routes:
- create note;
- store transaction workspace;
- admin/cashier note history;
- table/list;
- product/service lookup;
- service catalog create;
- workspace draft show/save;
- workspace create/edit/update;
- row add;
- note payment;
- closed note refund;
- corrections;
- surplus refund due;
- surplus refund paid;
- reopen closed note.

Go API candidates:
- `GET /api/notes`
- `POST /api/notes`
- `GET /api/notes/:id`
- `POST /api/notes/workspace`
- `GET /api/notes/workspace/draft`
- `POST /api/notes/workspace/draft`
- `PATCH /api/notes/:id/workspace`
- `POST /api/notes/:id/rows`
- `POST /api/notes/:id/payments`
- `POST /api/notes/:id/refunds`
- `POST /api/notes/:id/corrections/status`
- `POST /api/notes/:id/corrections/service-only`
- `POST /api/notes/:id/reopen`
- `POST /api/note-revision-settlements/:id/refund-due`
- `POST /api/note-revision-surplus-dispositions/:id/refund-paid`

### Employee Finance

Go API candidates from routes:
- `GET /api/employees`
- `POST /api/employees`
- `GET /api/employees/:id`
- `PUT /api/employees/:id`
- `GET /api/employees/:id/payrolls`
- `GET /api/employee-debts`
- `POST /api/employee-debts`
- `GET /api/employee-debts/:id`
- `POST /api/employee-debts/:id/payments`
- `POST /api/employee-debts/:id/adjustments`
- `POST /api/employee-debt-payments/:id/reverse`
- `GET /api/payrolls`
- `POST /api/payrolls`
- `POST /api/payrolls/:id/reverse`

### Expense

Go API candidates:
- `GET /api/expenses`
- `POST /api/expenses`
- `DELETE /api/expenses/:id`
- `GET /api/expense-categories`
- `POST /api/expense-categories`
- `PUT /api/expense-categories/:id`
- `PATCH /api/expense-categories/:id/activate`
- `PATCH /api/expense-categories/:id/deactivate`

### Reporting

Laravel reports:
- transaction cash ledger;
- payroll;
- employee debt;
- operational profit;
- operational expense;
- supplier payable;
- inventory stock value;
- transaction summary;
- Excel and PDF exports.

Go API direction:
- reporting should be read-model APIs first.
- export endpoints should be designed later as API-owned file/export contracts, not copied from Blade/PDF routes.

### Notification

Laravel routes:
- `POST /push-notifications/subscriptions`
- `DELETE /push-notifications/subscriptions`

Go API candidates:
- `POST /api/notifications/push/subscriptions`
- `DELETE /api/notifications/push/subscriptions`

## Initial Capability Key Families

These are candidates only. Final keys require domain contracts.

```text
products.list
products.show
products.create
products.update
products.delete
products.restore
products.stock_adjust
products.stock_adjust.reverse

suppliers.list
suppliers.update
suppliers.lookup

procurement.supplier_invoices.list
procurement.supplier_invoices.show
procurement.supplier_invoices.create
procurement.supplier_invoices.update
procurement.supplier_invoices.receive
procurement.supplier_invoices.void
procurement.supplier_payments.create
procurement.supplier_payments.reverse
procurement.supplier_payment_proofs.attach
procurement.supplier_payment_proofs.show

notes.list
notes.show
notes.create
notes.workspace.save
notes.rows.add
notes.payments.create
notes.refunds.create
notes.corrections.status
notes.corrections.service_only
notes.reopen
notes.revisions.update
notes.surplus_refund_due.create
notes.surplus_refund_paid.create

employees.list
employees.show
employees.create
employees.update

employee_debts.list
employee_debts.show
employee_debts.create
employee_debts.payments.create
employee_debts.adjustments.create
employee_debt_payments.reverse

payrolls.list
payrolls.create
payrolls.reverse

expenses.list
expenses.create
expenses.delete
expense_categories.list
expense_categories.create
expense_categories.update
expense_categories.activate
expense_categories.deactivate

reports.transaction_cash_ledger.view
reports.transaction_summary.view
reports.inventory_stock_value.view
reports.operational_profit.view
reports.operational_expense.view
reports.supplier_payable.view
reports.employee_debt.view
reports.payroll.view

notifications.push_subscriptions.create
notifications.push_subscriptions.delete
```

## Performance Direction

The user target is all normal API processes under one second.

This is accepted as a design target, not as proof yet.

Required implementation rules:
- list/search endpoints must use pagination and allow-listed sort/filter fields;
- hot list/report paths should use projection/read-model tables when direct joins are too expensive;
- indexes must be tied to actual query shapes;
- any `<1s` claim must have benchmark, query plan, or integration proof;
- no unbounded lookup endpoint.

## PostgreSQL First Pass Decisions

- Use PostgreSQL as the only target DB for new Go POS domains.
- Convert Laravel `json` payloads to PostgreSQL `jsonb` only when machine-readable.
- Use `bigint` for money if values may exceed 32-bit range or for consistency across finance modules.
- Use `timestamptz` for API-facing event timestamps.
- Preserve append-only/reversal patterns for financial and stock lifecycle records.
- Prefer soft delete, void, cancel, or reverse over physical delete for financial/stock/transaction records.
- Master data may support soft delete/restore if domain contract allows it.

## Data Still Needed

To turn this inventory into PostgreSQL migrations and Go domain contracts, request these next:

1. Full contents of omitted migration files that only had `alter`, `index`, `foreign`, `softDeletes`, or timestamp changes.
2. Full contents of `routes/web/admin_products.php`, `admin_procurement.php`, and `note.php` are already available in this batch, but controller/usecase files are still needed before behavior can be rewritten.
3. For the first implementation domain, send source files for `ServiceCatalog` or `ProductCatalog`:
   - `app/Core/ProductCatalog/**`
   - `app/Application/ProductCatalog/**`
   - `app/Ports/Out/ProductCatalog/**`
   - `app/Adapters/Out/ProductCatalog/**`
   - relevant request/controller files;
   - relevant tests.
4. Seeder files for `CreateOnly` profiles are needed before Go seed profiles can be designed.

## Recommended Next Batch

Because `servicecatalog` and `productcatalog` are safest first domains, request this next:

```text
database/migrations/2026_04_06_230200_add_soft_delete_foundation_to_products_and_suppliers.php
database/migrations/2026_04_06_230400_add_product_search_normalization_and_duplicate_hardening.php
database/migrations/2026_04_07_160100_fix_products_unique_constraints_for_soft_delete.php
database/migrations/2026_04_07_160200_rename_product_active_unique_indexes_to_legacy_names.php
database/migrations/2026_04_17_013500_add_stock_threshold_columns_to_products_table.php
database/migrations/2026_06_04_000100_create_service_catalog_items_table.php
database/migrations/2026_06_04_000200_seed_default_service_catalog_items.php
app/Core/ProductCatalog/**
app/Application/ProductCatalog/**
app/Ports/Out/ProductCatalog/**
app/Adapters/Out/ProductCatalog/**
app/Adapters/In/Http/Controllers/Admin/Product/**
app/Adapters/In/Http/Controllers/ProductCatalog/**
app/Adapters/In/Http/Requests/ProductCatalog/**
tests/Feature/ProductCatalog/**
tests/Unit/Core/ServiceCatalog/**
database/seeders/CreateOnly/CreateMasterBasicSeeder.php
database/seeders/CreateOnly/CreateInventorySeeder.php
```

## Current Stage 0 Status

- Schema inventory: partial, from `Schema::create` output only.
- Route inventory: partial but useful, from web routes.
- API inventory: intentionally not sourced from Laravel `routes/api.php`.
- PostgreSQL schema: not yet ready for implementation.
- First safe domain candidate: `servicecatalog`, then `productcatalog`.
