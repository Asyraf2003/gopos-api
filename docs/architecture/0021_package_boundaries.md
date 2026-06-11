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

# Package Boundaries

## Purpose

Prevent packages from absorbing unrelated responsibilities.

## One Folder One Package

Each folder contains one Go package.

Package names must reflect role, not technology fashion.

## File Size Policy

Default maximum:

- 100 lines for Echo handlers.
- 120 lines for request/response DTO files.
- 150 lines for use case files.
- 180 lines for repository files.
- 120 lines for domain entity/value object files.
- 200 lines for tests, unless table-driven cases need more.

If a file exceeds the limit, the reason must be documented in the blueprint or review note.

## Package Responsibility Rules

- `domain`: entities, value objects, domain errors, domain services, invariants.
- `application`: use cases, commands, queries, orchestration, transaction boundary request.
- `ports`: interfaces needed by application/domain boundaries.
- `adapters/in/http`: Echo handlers, route registration, request parsing, response mapping.
- `adapters/out/postgres`: SQL, row mapping, repository implementation.
- `platform`: config, logger, clock, id generator, DB pool, middleware primitives.
- `tests`: external behavior and contract proof.
- `scripts`: discipline automation.

## Forbidden Mixing

- No SQL in use cases.
- No business branching in handlers.
- No HTTP status selection in domain.
- No raw `map[string]any` as stable public API DTO.
- No global mutable capability state outside the capability registry/storage boundary.

