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

# Public Contracts

## Purpose

API contracts are public contracts. They must be stable, tested, and versioned.

## Contract Scope

A public contract includes:

- method and path;
- request body;
- query parameters;
- path parameters;
- response envelope;
- data DTO;
- error codes;
- validation errors;
- auth requirement;
- permission requirement;
- capability key;
- idempotency behavior;
- pagination/filter/sort behavior;
- timestamp format;
- money format.

## Versioning Rule

Breaking changes require one of:

- new API version;
- explicit migration period;
- owner decision and contract test update.

## Forbidden Behavior

- Do not return raw database rows.
- Do not expose internal enum names unless they are public terms.
- Do not change JSON field names casually.
- Do not return mixed envelope formats.
- Do not make UI parse human messages as machine state.

