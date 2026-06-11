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

# Step By Step Execution

## Purpose

Keep work controlled and reviewable.

## Rule

One response may contain only one active step.

An active step must have:

- one target;
- one bounded scope;
- one expected output;
- one proof method.

## Valid Active Step Examples

- Add a domain contract doc for one module.
- Add one use case and its unit test.
- Add one repository method and integration test.
- Add one Echo handler for an existing use case.
- Add one migration and database test.
- Add one capability-control registry entry and API test.

## Invalid Active Step Examples

- "Build the whole finance module."
- "Make all CRUD."
- "Clean architecture and tests."
- "Fix docs and implementation."

## Completion Rule

After the active step has proof, stop and wait for user feedback before the next step.

