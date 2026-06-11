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

# Scope And Facts

## Purpose

Keep every work step grounded in proof.

## FACT

FACT is only:

- inspected source code;
- inspected docs;
- user-provided command output;
- visible test output;
- explicit owner decision.

## GAP

GAP is anything important that is not proven yet.

Examples:

- unknown table owner;
- unknown API contract;
- unknown capability rule;
- unknown transaction boundary;
- unknown authorization behavior;
- unknown test coverage.

## SCOPE-IN

Only the package, API, domain, table, migration, or document explicitly selected for the active step.

## SCOPE-OUT

Related cleanup, refactors, route redesign, schema redesign, UI redesign, and unrelated tests stay out unless explicitly selected.

## Forbidden Behavior

- Do not infer file contents.
- Do not claim a package is clean without inspection.
- Do not claim an endpoint works without request/test proof.
- Do not call a design "done" when only a plan exists.

