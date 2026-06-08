# Handoffs

This folder contains session continuation notes.

Each handoff must include:
- active scope;
- files changed;
- proof collected;
- tests run;
- open gaps;
- next valid active step.

For long-running transition scopes, each handoff must also point to the active progress ledger and state whether the ledger was updated.

Active transition ledger:

- `docs/evidence/0003_laravel_to_go_transition_progress_ledger.md`

## Current Continuation Index

Use the active transition ledger for current state first.

Current handoffs:

- `2026-06-08-servicecatalog-postgres-persistence-blueprint.md`: current ServiceCatalog persistence blueprint proposal context.
- `2026-06-08-docs-quality-feedback-crosscheck.md`: current docs-quality cleanup and proof context.

Historical handoffs:

- Capability-control handoffs before `2026-06-08-capability-control-closeout.md` are historical unless the ledger names them as proof references.
- Prompt/template hardening handoffs from 2026-06-07 are historical unless the active task is AI workflow maintenance.
- ServiceCatalog slice 1 handoffs are historical after `2026-06-08-servicecatalog-implementation-slice-1.md` and the active ledger summary.

## Archiving Policy

Closed handoffs may be moved to `docs/archive/handoffs-closed/` when all conditions are true:

- the scope is closed with proof;
- the proof is referenced in the active ledger or a historical evidence snapshot;
- no current next step depends on the handoff details;
- the handoff is at least 30 days old, unless the owner explicitly approves earlier archiving.

Archived handoffs are historical only and cannot override active ledgers, blueprints, ADRs, source code, or command output.
