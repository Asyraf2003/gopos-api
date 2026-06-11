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

# Auth runtime local dev evidence

Date: 2026-06-06
Scope: Manual auth local runtime proof and local developer workflow cleanup for pos-go / gopos-api.

## FACT

This evidence file records only proof provided by the user in the active session and the prior handoff `docs/handoffs/2026-06-06-auth-runtime-local-dev.md`.

Local PostgreSQL runtime proof provided:

```text
PostgreSQL local is active on 127.0.0.1:5432.
App database connection works as posgo_app to posgo_app_db.
```

## GAP

This evidence file is incomplete. The note stopped after local PostgreSQL connectivity proof and did not record complete manual auth runtime proof.

Missing proof in this file:

- exact manual login HTTP command output;
- exact refresh/logout runtime output;
- database row/state proof after runtime auth actions;
- final command summary for the local dev workflow cleanup.

Use this file only for the DB connectivity facts above.

Use the active progress ledger and related handoff for current auth runtime gaps.

## REVIEW TRIGGER

Review and close or supersede this evidence by 2026-06-15 or before the next auth runtime change, whichever comes first.
