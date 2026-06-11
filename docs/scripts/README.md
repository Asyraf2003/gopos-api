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

# Scripts

This folder defines command and automation contracts.

## Contents

- `0090_makefile_and_scripts.md`: stable Makefile targets, script behavior, seed profiles, and architecture audit expectations.

## Use This Folder When

- adding or changing a Makefile target;
- writing an audit, test, migration, or seed script;
- deciding what command output counts as proof;
- maintaining repeatable checks for terminal, CI, and AI handoffs.

Scripts must be deterministic, print clear PASS/FAIL output, and exit non-zero on failure.
