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
