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

# Cross-AI Prompt Templates Blueprint

## FACT
- `docs/` is the canonical rules and workflow package.
- `docs/templates/` currently contains only the domain scope packet template.
- Work may move between terminal Codex and web AI sessions.
- Web AI copy features can break or truncate prompts when prompt text contains confusing Markdown fences.
- The user needs English templates for recurring work: Codex sessions, web AI sessions, analysis, tests, data capture, long-pause resume, and session handoff.

## GAP
- The templates do not yet cover routine terminal Codex and web AI workflows.
- The templates do not yet specify safe copy rules for web AI prompts.
- The templates do not yet tell a web AI where generated data should be placed in this repository.

## DECISION
- Add reusable English prompt templates under `docs/templates/`.
- Keep templates copy-safe by using placeholders and avoiding nested Markdown code fences inside copyable prompt blocks.
- Separate Codex prompts from web AI prompts because their operating constraints differ.
- Add a data capture template so web AI output has a consistent target path under `docs/evidence/`, `docs/blueprints/`, `docs/handoffs/`, or `docs/archive/`.
- Keep this as documentation-only work. No Go runtime behavior changes.

## SCOPE-IN
- Template README update.
- Prompt authoring rules.
- Codex session prompts.
- Web AI session prompts.
- Analysis/review prompts.
- Testing/proof prompts.
- Data capture prompts.
- Resume-after-pause prompts.

## SCOPE-OUT
- Go source implementation.
- API behavior changes.
- Database migration changes.
- Capability implementation.

## PROOF PLAN
- Inspect new template files.
- Run `bash scripts/audit_ai_rules.sh`.
- Confirm `docs/templates/README.md` lists the new templates.

## STEP ORDER
1. Add this blueprint.
2. Add template documents.
3. Update `docs/templates/README.md`.
4. Run docs audit.

## NEXT ACTIVE STEP
Add the cross-AI prompt templates under `docs/templates/`.
