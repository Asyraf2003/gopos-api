# Handoff: CLI Command Formatter Rules

## Date

2026-06-09

## Active Scope

Add repository-visible CLI command formatter rules so future AI sessions can consistently transform Asyraf's raw patch text, heredocs, shell fragments, proof commands, and mixed command/narrative input into paste-ready terminal commands.

## Files Changed

```text
docs/templates/0123_cli_command_formatter_rules.md
docs/AGENTS.md
docs/0003_session_start_protocol.md
docs/templates/0122_web_ai_session_prompts.md
scripts/audit_ai_rules.sh
docs/handoffs/2026-06-09-cli-command-formatter-rules.md
```

## FACT

- Existing repository AI rules cover scope control, proof, Web AI read-only behavior, command plans, and progress gates.
- Existing rules did not explicitly encode Asyraf's CLI formatter preferences.
- This handoff records the addition of a dedicated CLI command formatter template.
- The formatter template captures paste-ready terminal command output.
- The formatter template forbids default `cd /home/asyraf/Code/go/pos-go`.
- The formatter template forbids `set -euo pipefail`.
- The formatter template prefers `rg` over `grep`.
- The formatter template prefers `fd` over `find`.
- The formatter template requires safe heredoc handling.
- The formatter template requires outer markdown fences longer than inner markdown fences.
- The formatter template forbids Git mutation commands unless explicitly requested.

## PROOF

Run:

```bash
bash scripts/audit_ai_rules.sh
make verify
```

Expected result:

```text
[PASS] AI rules audit passed
[PASS] aggregate audit passed
```

## GAP

- This only makes the rule repository-visible and audit-anchored.
- Future AI sessions still need to be told to read repository rules or this formatter template.
- It does not change Laravel-to-Go implementation progress.

## PROGRESS

Workflow consistency improves.

Laravel-to-Go implementation progress is unchanged.

## NEXT

Execution channel:

```text
owner/local terminal
```

Run audit proof locally and publish through the normal repository flow if proof passes.
