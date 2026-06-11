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

# Core

This folder defines AI work discipline and progress rules.

## Contents

- `0010_scope_and_facts.md`: scope control, fact gathering, and GAP handling.
- `0011_blueprint_first.md`: blueprint requirements and implementation gates.
- `0012_step_by_step_execution.md`: one-active-step execution discipline.
- `0013_proof_and_progress.md`: proof requirements and progress reporting rules.

## Use This Folder When

- starting a session;
- deciding whether implementation may begin;
- reporting progress;
- handling missing information without guessing.

Core rules apply to every technical task in this repository.
