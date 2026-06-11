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

# Database

This folder defines PostgreSQL and persistence rules.

## Contents

- `0040_postgresql_policy.md`: schema, migration, transaction, repository, and data-integrity policy.

## Use This Folder When

- adding or changing migrations;
- designing repositories or query adapters;
- deciding transaction boundaries;
- defining database-owned invariants and constraints.

Database-backed mutations must be covered by domain contracts, capability rules, transactions, and test proof.
