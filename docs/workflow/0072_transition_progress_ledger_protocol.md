# Transition Progress Ledger Protocol

## Purpose

Long migrations need one current progress ledger in addition to blueprints, evidence, and handoffs.

The ledger keeps the owner, terminal Codex, web AI, and future sessions aligned on:

- accepted stage order;
- current stage status;
- proof behind progress percentages;
- open gaps;
- next valid active step;
- handoff and README cascade requirements.

## Applies To

Use this protocol for long-running transition scopes, especially:

- Laravel-to-Go API transition;
- multi-stage domain migration;
- capability-control foundation;
- quality/security gate rollout;
- cross-AI work that spans more than one session.

## Ledger Location

Current transition progress ledgers live under `docs/evidence/` because they are proof-linked status documents.

The active Laravel-to-Go ledger is:

```text
docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

## Ledger Rules

Every progress ledger must include:

- date updated;
- active blueprint;
- linked handoffs;
- linked evidence;
- stage table;
- per-stage status;
- proof reference for every non-zero progress estimate;
- open gaps;
- next valid active step;
- context-window status when updated by an AI session.

## Progress Rules

Progress may increase only from proof.

Valid proof includes:

- inspected source files;
- inspected docs;
- command output;
- test output;
- user-provided runtime output;
- accepted ADR or blueprint decision.

Plans alone do not increase progress.

When a stage has partial implementation but missing runtime or test proof, mark it partial and list the missing proof.

## Progress Update Gate

Any session that receives new proof for durable implementation work must update or draft an update for the active progress ledger before naming the next implementation step.

Durable proof includes:

- passing aggregate audit;
- passing focused tests;
- applied database migrations;
- accepted blueprint changes;
- route/capability audit proof;
- runtime proof;
- security audit proof.

If the session cannot mutate files, it must provide paste-ready ledger and handoff text or an owner/local terminal command plan.

A session must not advance to a new domain, new implementation slice, or closeout until the current proof is reflected in:

- the active progress ledger;
- the relevant handoff;
- the `NEXT` field.

Local proof must be marked as local proof until connector or remote repository validation exists.

Remote validation must not be claimed from local terminal output.

When local proof exists but connector validation is missing, use this status wording:

```text
locally implemented with proof; connector validation pending
```

## Cascade Rules

When the ledger is created or materially updated, update the impacted chain when feasible:

- `docs/README.md`;
- `docs/AGENTS.md`;
- `docs/0001_index.md`;
- `docs/workflow/README.md`;
- `docs/evidence/README.md`;
- `docs/handoffs/README.md`;
- `scripts/audit_ai_rules.sh` if the ledger becomes mandatory.

## Handoff Rules

After durable progress on a long-running transition scope, create or update a handoff under `docs/handoffs/`.

The handoff must point to the current progress ledger and state:

- what changed;
- proof run;
- gaps still open;
- estimated active-scope progress;
- estimated context-window status;
- next valid active step.

## Web AI Rules

Web AI with a GitHub connector stays read-only by default.

If web AI updates progress, it must draft paste-ready ledger or handoff text in the chat response unless the owner gives exact mutation permission.

Connector facts and local terminal facts can differ. If they do, mark `GAP` and request the smallest proof.

## Stop Condition

Do not start a new Laravel-to-Go implementation stage when the ledger says the required foundation proof is missing.

For protected POS API implementation, capability-control proof is required before exposing endpoints.
