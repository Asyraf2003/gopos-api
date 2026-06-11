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

# Testing

This folder defines test and quality-gate rules.

## Contents

- `0060_test_and_quality_gates.md`: unit, integration, API, contract, architecture, migration, and script proof expectations.

## Use This Folder When

- planning proof for a blueprint step;
- adding or changing use cases, repositories, handlers, migrations, or contracts;
- deciding whether `go test`, API tests, DB tests, or architecture checks are required.

Tests are delivery gates. Do not claim completion without visible proof appropriate to the changed surface.
