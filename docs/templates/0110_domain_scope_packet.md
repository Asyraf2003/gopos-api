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

# Domain Scope Packet Template

Use this template when sending one domain/API scope to another AI or another session.

## Scope

Active domain/API:

```text
REPLACE_WITH_DOMAIN_OR_API
```

Goal:

```text
REPLACE_WITH_GOAL
```

## Included Files

Read-only context:

```text
REPLACE_WITH_FILES
```

Editable files:

```text
REPLACE_WITH_FILES
```

Forbidden files:

```text
REPLACE_WITH_FILES
```

## Required Rules To Read

- `docs/README.md`
- `docs/AGENTS.md`
- `docs/0002_decision_policy.md`
- `docs/architecture/0020_hexagonal_go_api.md`
- `docs/architecture/0022_api_capability_control.md`
- `docs/domain/0030_domain_contracts.md`
- `docs/api/0050_echo_http_contract.md`
- `docs/testing/0060_test_and_quality_gates.md`
- `docs/workflow/0071_handoff_protocol.md`

## Domain Contract

Tables:

```text
REPLACE_WITH_TABLES
```

Allowed operations:

```text
create:
update:
delete:
show:
list:
```

Forbidden operations:

```text
REPLACE_WITH_FORBIDDEN_OPERATIONS
```

Capability keys:

```text
REPLACE_WITH_CAPABILITY_KEYS
```

## Expected Output

The receiving AI must return:

- files changed;
- summary of implementation;
- proof commands run or commands the owner must run;
- gaps;
- handoff for next step;
- estimated progress percentage;
- estimated context-window status.

## Working Style

The receiving AI should:

- execute the largest safe slice that still fits one active step;
- keep progress updates short;
- keep the final report compact;
- ask for the smallest specific source batch if source data is missing;
- ask ADR-level decisions with 2-3 options, tradeoffs, and a recommended option when clear.

## Proof Required

Minimum proof:

```bash
make verify
```

Focused proof:

```bash
REPLACE_WITH_FOCUSED_COMMANDS
```

## Handoff

Next valid active step:

```text
REPLACE_WITH_NEXT_STEP
```

## Concrete Example: ServiceCatalog Slice 1

Use this as a shape example only. Replace values for the active scope.

Active domain/API:

```text
servicecatalog
```

Goal:

```text
Implement ServiceCatalog slice 1: domain, ports, usecase contracts, and unit tests only.
```

Read-only context:

```text
docs/README.md
docs/AGENTS.md
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
docs/blueprints/0024_servicecatalog_domain_contract.md
docs/blueprints/0025_servicecatalog_implementation_slice_1.md
```

Editable files:

```text
internal/modules/servicecatalog/domain/**
internal/modules/servicecatalog/ports/**
internal/modules/servicecatalog/usecase/**
```

Forbidden files:

```text
internal/modules/servicecatalog/transport/**
internal/platform/postgres/**
migrations/**
cmd/**
route registration
capability seed migrations
```

Capability keys:

```text
service_catalog.list
service_catalog.lookup
service_catalog.show
service_catalog.create
service_catalog.update
service_catalog.activate
service_catalog.deactivate
```

Focused proof:

```bash
go test ./internal/modules/servicecatalog/...
make verify
```
