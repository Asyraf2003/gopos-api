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

audit:allow-oversize reason=bootstrap-wiring
-->

# Closed Handoffs Archive

This folder stores closed or historical handoffs that should not be scanned as active continuation context.

Backlinks:

```text
docs/handoffs/README.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/evidence/0005_laravel_to_go_transition_history_2026_06_08.md
```

## Current Active Handoffs

Use `docs/handoffs/` for current and pending continuation notes.

## Archived Groups

Early automation and audit:

- `2026-04-21-automation-stage1-handoff.md`
- `2026-04-22-audit-baseline.md`

Auth foundation history:

- `2026-06-06-manual-auth-login.md`

Capability-control history:

- `2026-06-07-capability-contracts.md`
- `2026-06-07-capability-postgres-state.md`
- `2026-06-07-capability-runtime-middleware.md`
- `2026-06-07-capability-route-seeds.md`
- `2026-06-07-capability-admin-http-surface.md`
- `2026-06-08-capability-route-audit-script.md`
- `2026-06-08-capability-route-disabled-proof.md`
- `2026-06-08-capability-control-closeout-prep.md`
- `2026-06-08-capability-control-closeout.md`

AI workflow and docs guardrail history:

- `2026-06-07-ai-execution-channel-boundaries.md`
- `2026-06-07-ai-workstyle-template-update.md`
- `2026-06-07-decision-data-request-protocol.md`
- `2026-06-07-prompt-template-selection-rule.md`
- `2026-06-07-transition-progress-ledger.md`
- `2026-06-07-web-ai-owner-terminal-output-test.md`
- `2026-06-08-docs-quality-feedback-crosscheck.md`
- `2026-06-08-docs-scalability-blueprint-cleanup.md`
- `2026-06-08-progress-write-gate-entrypoint-hardening.md`
- `2026-06-08-progress-write-gate-hardening.md`

ServiceCatalog completed slice history:

- `2026-06-08-servicecatalog-domain-contract-blueprint.md`
- `2026-06-08-servicecatalog-domain-contract-accepted.md`
- `2026-06-08-servicecatalog-implementation-slice-1-plan.md`
- `2026-06-08-servicecatalog-implementation-slice-1-accepted.md`
- `2026-06-08-servicecatalog-implementation-slice-1.md`
- `2026-06-08-servicecatalog-postgres-persistence-blueprint.md`

ProductCatalog completed slice history:

- `2026-06-09-productcatalog-domain-slice-1.md`
- `2026-06-09-productcatalog-usecase-createproduct-progress.md`
- `2026-06-10-productcatalog-updateproduct-progress.md`
- `2026-06-10-productcatalog-softdeleteproduct-progress.md`
- `2026-06-10-productcatalog-restoreproduct-progress.md`
- `2026-06-10-productcatalog-getproductdetail-progress.md`
- `2026-06-10-productcatalog-listproducts-skeleton-progress.md`
- `2026-06-10-productcatalog-lookupproducts-skeleton-progress.md`
- `2026-06-10-productcatalog-listproductversions-skeleton-progress.md`
- `2026-06-10-productcatalog-listproductversions-behavior-progress.md`
- `2026-06-10-productcatalog-implementation-slice-1-closeout.md`
- `2026-06-10-productcatalog-postgres-persistence-blueprint-progress.md`
- `2026-06-10-productcatalog-postgres-persistence-blueprint-accepted.md`
- `2026-06-10-productcatalog-postgres-migration-checkpoint.md`
- `2026-06-10-productcatalog-postgres-repository-skeleton-progress.md`
- `2026-06-11-productcatalog-postgres-create-find-update-progress.md`
- `2026-06-11-productcatalog-postgres-reader-getbyid-progress.md`
- `2026-06-11-productcatalog-postgres-reader-list-progress.md`
- `2026-06-12-productcatalog-postgres-reader-lookup-progress.md`
- `2026-06-12-productcatalog-postgres-version-repository-progress.md`
- `2026-06-12-productcatalog-postgres-duplicate-checker-progress.md`
- `2026-06-12-productcatalog-postgres-persistence-closeout.md`
- `2026-06-13-productcatalog-runtime-capability-closeout.md`
- `2026-06-13-api-docs-error-envelope-closeout.md`

## Rule

Archived handoffs are proof history only. They do not define the next active step unless the active ledger or current handoff explicitly revives them.
