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

# Blueprint First

## Purpose

Implementation must start from a blueprint so decisions are auditable.

## Minimum Blueprint

Every blueprint must state:

- target;
- current state;
- facts already known;
- gaps still open;
- scope in;
- scope out;
- affected public API contracts;
- affected domain invariants;
- affected DB tables;
- affected capability-control rules;
- dependencies;
- risks;
- test/proof plan;
- step order.

## Implementation Gate

Code may start only when:

- active scope is clear;
- public contract impact is known;
- domain owner is known;
- package boundary is known;
- DB transaction boundary is known for mutations;
- capability rule is known for any endpoint;
- test path is known.

## Forbidden Behavior

- Do not implement a mutation from a vague request.
- Do not create packages before the package role is defined.
- Do not add CRUD endpoints before domain lifecycle rules are defined.
- Do not replace a blueprint with comments in code.

