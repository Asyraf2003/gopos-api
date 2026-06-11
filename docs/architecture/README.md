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

# Architecture

This folder defines the structural rules for the Go API.

## Contents

- `0020_hexagonal_go_api.md`: required hexagonal dependency direction and mutation/read flow.
- `0021_package_boundaries.md`: package ownership, file-size policy, and forbidden responsibility mixing.
- `0022_api_capability_control.md`: capability metadata and runtime capability-control requirements.
- `0023_public_contracts.md`: public contract protection rules.
- `0024_current_repo_layout.md`: current repository layout, protected paths, and layout-change rules.

## Use This Folder When

- creating or moving packages;
- adding modules, adapters, handlers, repositories, or use cases;
- changing route exposure or capability-control behavior;
- deciding whether a change needs an ADR.

Architecture rules override convenience. If implementation pressure conflicts with these documents, mark GAP and update the blueprint or ADR first.
