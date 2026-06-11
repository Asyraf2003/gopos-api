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

# Hexagonal Go API Baseline

## Purpose

The Go application must be a pure API with strict hexagonal architecture.

## Target Shape

Recommended top-level shape:

```text
cmd/api/
internal/core/
internal/application/
internal/ports/
internal/adapters/in/http/
internal/adapters/out/postgres/
internal/platform/
internal/modules/
migrations/
tests/
scripts/
docs/
```

The exact folder names may change by owner decision, but the dependency direction may not.

## Dependency Direction

- `domain/core` knows no transport, SQL, Echo, config, or HTTP.
- `application` knows domain and ports.
- `ports` define interfaces and DTO boundaries.
- `adapters/in/http` depends on application contracts.
- `adapters/out/postgres` depends on ports and PostgreSQL driver/query tooling.
- `platform` contains bootstrapping, config, logging, clock, id generation, and DB connection setup.

## Forbidden Imports

- Domain must not import Echo.
- Domain must not import PostgreSQL drivers.
- Domain must not import HTTP request/response types.
- Application must not import Echo handlers.
- HTTP handlers must not import PostgreSQL repositories directly.
- Repositories must not call Echo, middleware, or presentation code.

## Mutation Flow

All mutations must follow:

```text
Echo handler -> request DTO validation -> use case -> domain rules -> port -> postgres adapter -> transaction commit -> response presenter
```

No mutation may write directly from a handler.

## Read Flow

Reads may use query ports, but public response shape still belongs to API contract/presenter code, not raw SQL rows.

