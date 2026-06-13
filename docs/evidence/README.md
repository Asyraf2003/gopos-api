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

# Evidence

This folder stores proof references and command-output summaries.

Do not claim completion from plans. Completion requires proof.

Progress ledgers for long-running transitions may also live here when each estimate is tied to proof.

Active ledgers:

- `0003_laravel_to_go_transition_progress_ledger.md`: current Laravel-to-Go API transition status.

## Evidence Status Index

Use this index before treating an evidence file as complete proof.

| File | Status | Use Guidance |
| --- | --- | --- |
| `0001_laravel_stage0_schema_and_route_inventory.md` | Partial source inventory | Use for captured Laravel schema and route facts only. Do not treat omitted alter, foreign key, index, timestamp, or seed migrations as proven. |
| `0002_laravel_productcatalog_servicecatalog_inventory.md` | Partial source inventory | Use for ProductCatalog and ServiceCatalog facts captured from provided Laravel source. Product duplicate policy still needs owner decision. |
| `0003_laravel_to_go_transition_progress_ledger.md` | Active ledger | Use as the current transition status, progress, gaps, and next valid active step source. |
| `0004_adr_implementation_proof_index.md` | Active proof index | Use as the current map from accepted ADRs to implementation proof status and remaining ADR proof gaps. |
| `0005_laravel_to_go_transition_history_2026_06_08.md` | Historical proof snapshot | Use for completed-work history that was moved out of the active ledger to reduce context load. |
| `2026-06-13_api_architecture_product_status_review.md` | Active status review | Use for the current ProductCatalog catalog API, Product/inventory, architecture, docs/handoff cleanliness, and next-slice assessment. |
| `2026-06-06-auth-runtime-local-dev.md` | Incomplete local runtime evidence | Use only for the DB connectivity facts explicitly recorded in the file. Do not use it as complete manual auth runtime proof. |

Partial or incomplete evidence can support only the facts written inside that file.

Do not close a gap from a partial or incomplete evidence file unless the missing proof is added or referenced by the active ledger.

Incomplete evidence review trigger:

- review incomplete evidence at least before the next active scope in the same domain;
- review `2026-06-06-auth-runtime-local-dev.md` by 2026-06-15 or before the next auth runtime change, whichever comes first.
