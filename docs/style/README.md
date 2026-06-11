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

# Style

This folder defines Go code style and forbidden implementation patterns.

## Contents

- `0100_go_style.md`: naming, dependency, context, error, DTO, and forbidden-pattern rules.

## Use This Folder When

- naming use cases, repositories, commands, queries, DTOs, or handlers;
- deciding where dependencies should be passed;
- reviewing Go code for framework leakage or hidden globals;
- checking whether a helper or package is becoming too broad.

Style rules do not override architecture or security, but they keep implementation predictable.
