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

# PostgreSQL Policy

## Purpose

PostgreSQL is the database target and must be treated as a first-class runtime, not a later compatibility detail.

## Schema Rules

- Use explicit primary keys.
- Use explicit foreign keys where referential integrity belongs to the database.
- Use `check` constraints for non-negative money, quantity, counters, and lifecycle-safe fields.
- Use integer or bigint for money in rupiah or smallest currency unit.
- Use `timestamptz` for API-facing event timestamps unless a domain reason says otherwise.
- Use JSONB only for structured payloads that need JSON behavior.
- Use text for opaque historical snapshots only by decision.

## Migration Rules

- Every migration must have forward behavior and rollback/compensation notes.
- Migrations must be small and reviewable.
- Do not mix unrelated domains in one migration.
- Do not change historical data without a data migration plan and proof query.

## Repository Rules

- SQL lives only in PostgreSQL adapters or query packages approved by the architecture.
- Repositories return domain/application DTOs, not raw driver rows.
- Row mapping must be explicit.
- Transactional methods must accept an explicit transaction context or unit-of-work boundary.

## Transaction Rules

Mutations that affect finance, stock, payment, refund, audit, or lifecycle state must define:

- transaction boundary;
- lock strategy;
- retry/idempotency behavior;
- audit write behavior;
- consistency checks after commit.

## Forbidden Behavior

- Do not rely on MySQL unsigned semantics.
- Do not use floating point for money.
- Do not store machine-readable status only in free text.
- Do not let handlers build SQL.
- Do not let SQL decide business lifecycle without domain/application approval.

