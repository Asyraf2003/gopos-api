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

# Proof And Progress

## Purpose

Progress must be tied to evidence.

## Accepted Proof

- `go test` output.
- `make verify` output.
- migration up/down output.
- API test output.
- contract test output.
- architecture/lint script output.
- inspected diff or file contents.
- explicit owner approval.

## Progress Rule

- Plans do not increase progress.
- Created files increase progress only if the active step was file creation.
- Tests passing increase progress only when output is visible.
- A mutation is not complete without unit, adapter, and API proof where relevant.
- When proof changes durable project status, update or draft the active ledger and handoff before naming `NEXT`.
- Local proof and remote connector proof must be separated.
- Remote validation must not be claimed from local terminal output.
- Do not use "closed", "complete", "done", or "ready" until every acceptance gate and repository fact is proven.
- Web AI file update requests require `COMMAND PLAN FOR OWNER / LOCAL TERMINAL` when exact mutation permission is absent.
- Paste-ready text must not replace the command plan unless the owner explicitly asks for draft-only text.

If local proof exists but connector validation is missing, use this status wording:

```text
locally implemented with proof; connector validation pending
```

## Required Proof Statement

Every completion claim must state:

- command or artifact;
- visible result;
- meaning for the active step.
