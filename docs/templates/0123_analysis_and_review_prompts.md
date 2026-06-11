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

# Analysis And Review Prompts

Use these prompts when asking an AI to analyze source, schema, routes, logs, or a design decision.

## Analyze A New Source Batch

```text
Analyze this source batch for the Go API migration.

ACTIVE SCOPE
REPLACE_WITH_SCOPE

SOURCE BATCH
REPLACE_WITH_SOURCE_BATCH

OUTPUT TARGET
docs/evidence/REPLACE_WITH_EVIDENCE_FILE.md

RULES
- Extract facts only from the provided source.
- Separate FACT from GAP.
- Identify tables, fields, invariants, operations, routes, capability candidates, tests to preserve, and unresolved decisions.
- Do not design implementation until the facts are extracted.
- Do not include unrelated source outside the batch.

EXPECTED OUTPUT
- Evidence document content.
- Follow-up blueprint changes needed.
- Smallest missing data request.
- If a decision is blocked, provide 2-3 owner decision options with tradeoffs and a recommended option when clear.
```

## Analyze A Domain For Migration

```text
Analyze this domain for migration from Laravel to Go.

DOMAIN
REPLACE_WITH_DOMAIN

KNOWN EVIDENCE
REPLACE_WITH_EVIDENCE_FILES_OR_TEXT

REQUIRED OUTPUT
- domain aggregate candidates;
- allowed operations;
- forbidden operations;
- database tables;
- invariants;
- transaction boundaries;
- audit requirements;
- capability keys;
- API route candidates;
- tests to preserve;
- gaps.
- smallest missing source batch request;
- ADR or owner decision question with 2-3 options and tradeoffs when needed.

Do not write Go code. Produce a blueprint-ready analysis.
```

## Review A Blueprint

```text
Review this blueprint for implementation readiness.

BLUEPRINT
REPLACE_WITH_BLUEPRINT_TEXT

CHECK
- Does it have facts?
- Does it mark gaps?
- Does it define scope in and scope out?
- Does it protect API contracts?
- Does it define domain invariants?
- Does it define DB constraints?
- Does it define capability keys?
- Does it define tests and proof?
- Does it have exactly one next active step?

Return blockers first.
```

## Review Generated Code

```text
Review the code below against the project rules.

CODE OR DIFF
REPLACE_WITH_CODE_OR_DIFF

FOCUS
- hexagonal boundary violations;
- domain logic in handlers;
- SQL outside adapters;
- missing capability check;
- missing transaction;
- missing audit decision;
- weak validation;
- missing tests.

Return findings first with severity.
```
