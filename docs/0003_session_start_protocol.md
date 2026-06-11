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

# Session Start Protocol

## Purpose

Every session must start with scope control before coding, planning, or commands.

## Mandatory Opening Flow

1. Read `docs/0001_index.md`.
2. Read `docs/0002_decision_policy.md`.
3. Identify the active scope from the latest user prompt.
4. Identify named files, commands, branches, blueprints, ADRs, handoffs, APIs, domains, or tables.
5. Classify those references as ACTIVE, CONSTRAINT, REFERENCE, or DEFERRED.
6. Inspect current files before making repo-state claims.
7. Build a short blueprint for the active scope.
8. State exactly one active step.
9. Define expected proof.
10. Before naming NEXT, apply the Progress Write Gate when new proof may change project progress.
11. If the active task is CLI command formatting, apply `docs/templates/0123_cli_command_formatter_rules.md`.

## Required Work Sections

For technical work, responses must separate:

- FACT
- GAP
- DECISION
- BLUEPRINT
- ACTIVE STEP
- PROOF
- NEXT
- PROGRESS

Use concise sections for small tasks.

## Active Scope Rule

The latest user prompt controls the active scope.

Do not switch to another domain, package, or cleanup just because it is nearby.

## Stop Conditions

Stop and mark GAP if:

- the requested source file cannot be found;
- public contract impact is unknown;
- DB ownership is unknown;
- capability-control behavior is unknown;
- tests cannot be identified for a mutation;
- the change would cross a forbidden boundary;
- durable proof changes progress but the active ledger and relevant handoff have not been updated, cited, drafted, or placed in an owner/local command plan;

