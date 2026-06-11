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

# Handoffs

This folder contains current session continuation notes only.

Use the active transition ledger for current state first:

```text
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

Closed or historical handoffs live in:

```text
docs/archive/handoffs-closed/
```

Archive index:

```text
docs/archive/handoffs-closed/README.md
```

## Current

- `2026-06-08-servicecatalog-runtime-capability-implementation.md`: latest ServiceCatalog runtime/capability closeout context and next planning-only scope.
- `2026-06-08-servicecatalog-runtime-capability-blueprint-accepted.md`: accepted blueprint 0027 context retained until the next scoped session prompt is prepared.

## Pending / Needs Follow-Up

- `2026-06-06-auth-runtime-local-dev.md`: auth runtime proof remains incomplete; keep visible until ADR `0009` runtime proof is closed or superseded.

## Archive Policy

Move a handoff to `docs/archive/handoffs-closed/` when all conditions are true:

- the scope is closed with proof or clearly historical;
- the proof is summarized in the active ledger, evidence history, or archive index;
- no current next step depends on scanning the full handoff;
- the owner approves early archive or the handoff is old enough under project policy.

Archived handoffs are historical only and cannot override active ledgers, blueprints, ADRs, source code, or command output.

## Required Handoff Fields

Each active handoff must include:

- active scope;
- files changed;
- proof collected;
- tests run;
- open gaps;
- next valid active step;
- estimated scope progress percentage;
- estimated context-window status.
