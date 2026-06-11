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

# Handoff: ServiceCatalog Domain Contract Blueprint

## Date

2026-06-08

## Active Scope

Create the first POS business-domain blueprint/domain contract after capability-control foundation closeout.

## Domain

```text
servicecatalog
```

## Files Included

- `docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md`
- `docs/blueprints/0020_catalog_foundation_migration.md`
- `docs/domain/0030_domain_contracts.md`
- `docs/architecture/0022_api_capability_control.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`
- `docs/blueprints/0024_servicecatalog_domain_contract.md`

## Files Changed

- `docs/blueprints/0024_servicecatalog_domain_contract.md`
- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`
- `docs/handoffs/2026-06-08-servicecatalog-domain-contract-blueprint.md`

## Files Forbidden To Touch

- Go implementation files
- PostgreSQL migrations
- Capability seed migrations
- ProductCatalog implementation
- Inventory implementation
- Auth behavior
- Admin capability HTTP behavior
- Production secrets
- GitHub refs, branches, commits, pull requests, issues, labels, reviewers, merges, or CI by Web AI

## Blueprint Referenced

- `docs/blueprints/0024_servicecatalog_domain_contract.md`
- `docs/blueprints/0020_catalog_foundation_migration.md`

## Decisions Made

- Start POS business-domain blueprint work with `servicecatalog`.
- Do not start `productcatalog` first because product duplicate policy remains unresolved.
- Treat ServiceCatalog as master data.
- Physical delete is forbidden for normal use.
- Activate/deactivate lifecycle is preferred.
- Do not implement Go code in this step.
- Do not add capability seeds before the contract is accepted.

## Proof Collected

Owner provided local source excerpts from:

```text
sed -n '1,220p' docs/evidence/0002_laravel_productcatalog_servicecatalog_inventory.md
sed -n '1,220p' docs/blueprints/0020_catalog_foundation_migration.md
sed -n '1,180p' docs/domain/0030_domain_contracts.md
sed -n '1,180p' docs/architecture/0022_api_capability_control.md
```

The excerpts show:

- ServiceCatalog source evidence is limited to seed and normalizer test.
- ServiceCatalog candidate domain contract already exists in catalog foundation blueprint.
- Domain contracts are required before implementation.
- Protected endpoints require authn, authz, capability check, validation, then usecase execution.

## Tests Or Commands Run

Pending after this docs update:

- `make verify`

## Gaps Still Open

- Full ServiceCatalog Laravel source is not available yet.
- Exact seed row source needs proof before Go seed migration.
- ServiceCatalog domain contract has not been owner-accepted yet.
- No Go ServiceCatalog implementation proof exists.
- No ServiceCatalog PostgreSQL migration proof exists.
- No ServiceCatalog capability seed proof exists.

## Next Valid Active Step

Review and accept or adjust `docs/blueprints/0024_servicecatalog_domain_contract.md`.

Do not start implementation until the contract is accepted.

## Estimated Scope Progress Percentage

ServiceCatalog domain contract blueprint: 70%.

Business Phase 1 implementation: 0%.

Overall Laravel-to-Go transition: 20%.

## Estimated Context-Window Status

Enough context remains for one focused ServiceCatalog contract review.
