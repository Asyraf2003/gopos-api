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

# Test And Quality Gates

## Purpose

Tests and scripts discipline the project so architecture does not drift.

## Required Gates

The project should provide a single verification entry point:

```bash
make verify
```

`make verify` should run:

- formatting check;
- lint/static analysis;
- unit tests;
- integration tests where configured;
- API contract tests;
- architecture boundary checks;
- migration checks where configured.

## Test Layers

- Domain tests: invariants and lifecycle behavior.
- Application tests: use case orchestration, ports, errors, idempotency.
- Repository tests: PostgreSQL SQL behavior, constraints, row mapping.
- HTTP tests: Echo routing, auth, capability, validation, response envelope.
- Contract tests: public API request/response stability.
- Architecture tests: forbidden imports and package boundaries.

## Mutation Test Minimum

Every mutation needs:

- use case unit test;
- repository/integration test if persistence changes;
- HTTP/API test if endpoint exists;
- capability-disabled test;
- authorization test;
- validation error test;
- audit/idempotency test when required by domain contract.

## Scripts

Scripts must be deterministic and documented.

Recommended scripts:

- forbidden import scanner;
- max file line checker;
- route capability registry checker;
- API envelope checker;
- migration naming checker;
- PostgreSQL constraint checker.

## Forbidden Completion Claims

- Do not say "done" without command output or inspected proof.
- Do not call a mutation safe without tests.
- Do not skip capability-disabled tests for protected endpoints.

