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

# Workflow

This folder defines documentation hygiene, handoff, and work sequencing rules.

## Contents

- `0070_docs_go_workflow.md`: document classes, blueprint rules, handoff rules, archive rules, naming, and cleanliness rules.
- `0071_handoff_protocol.md`: required handoff fields and scope packet behavior.
- `0072_transition_progress_ledger_protocol.md`: progress ledger rules for long-running transition scopes.

## Use This Folder When

- creating or updating blueprints, ADRs, handoffs, evidence, or archive material;
- moving work between sessions or tools;
- updating progress, proof, gaps, or next valid step for a long-running transition;
- deciding whether a document is active, historical, or proof-bearing.

Workflow rules keep standards, plans, proof, and history from overwriting each other.

For Laravel-to-Go continuation work, read `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md` before the latest handoff. The ledger states the valid next step; the handoff gives the latest proof and local context.
