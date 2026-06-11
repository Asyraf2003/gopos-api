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

# CLI Command Formatter Rules

## Status

Active support template.

## Purpose

This document defines the paste-ready CLI command formatting behavior expected when Asyraf sends raw text, patch fragments, heredocs, shell snippets, Python scripts, proof commands, Git commands, Make commands, or mixed markdown/command content for the POS Go repository.

Goal:

```text
Asyraf sends raw material.
AI returns valid terminal commands that are safe to copy-paste.
```

## Repository Context

Default repository:

```text
/home/asyraf/Code/go/pos-go
```

Asyraf is usually already inside the repository.

Do not prepend this unless explicitly needed:

```bash
cd /home/asyraf/Code/go/pos-go
```

## Primary Behavior

When the active task is CLI formatting, output the cleaned command first.

The response must prioritize:

```text
valid shell syntax
valid heredoc syntax
closed markdown fences inside files
safe copy-paste shape
minimal explanation
no new implementation scope
```

## Hard Rules

1. Output terminal commands as the primary answer.
2. Do not invent implementation details.
3. Do not expand the requested scope.
4. Do not jump to CRUD, ProductCatalog, Inventory, HTTP, PostgreSQL, route registration, capability seed migration, Git operations, or adjacent work unless the owner explicitly requested it.
5. If user input contains markdown code fences inside file content, wrap the final assistant command response with a larger markdown fence than the inner file fences.
6. Never place free text after an `EOF`, `PY`, `SQL`, or other heredoc terminator inside the same shell block.
7. Put important narrative inside the target heredoc file when the task is creating docs, blueprints, ADRs, evidence, or handoffs.
8. If raw input contains corrupt command syntax, broken Python indent, unclosed markdown fences, typo-prone shell, or mixed command/narrative sections, silently repair it while preserving the user's technical intent.
9. Do not use `set -euo pipefail`.
10. Use `set +e +u +o pipefail` only when it is useful to neutralize strict shell mode.
11. Convert `grep` commands to `rg`.
12. Convert `find` commands to `fd` unless `/usr/bin/find` is explicitly required.
13. Use `python3 - <<'PY'` for Python heredocs.
14. Use single-quoted heredoc delimiters for literal file writes, for example `<<'EOF'`.
15. Keep proof commands included when the user provides them.
16. If Git commit or push commands are requested, place them only after proof or validation commands such as `make verify`, `git diff`, or `git status --short`.
17. Do not include Git mutation commands unless the owner explicitly requested Git operations.
18. If the user asks "kasi commandnya aja", return only the command block.
19. If the user asks whether a command is valid, check syntax and available repo facts before giving confidence.
20. If repo facts are missing, provide the smallest proof command needed instead of guessing.

## Command Formatting Preferences

Use:

```bash
rg
fd
python3
make verify
git status --short
git diff --
```

Avoid by default:

```bash
grep
find
python
set -euo pipefail
cd /home/asyraf/Code/go/pos-go
```

Use `/usr/bin/find` only when exact POSIX `find` behavior is required or when `fd` is not suitable.

## Mixed Input Handling

If the user sends several copied sections, classify them as:

```text
terminal command
file content
proof command
commit command
prompt for another session
plain instruction
duplicate copied text
```

Then return only the valid usable parts in a clean shape.

Duplicate command or file sections may be deduplicated when they are exact or clearly accidental duplicates.

Do not delete similar-looking sections if they carry different technical meaning.

## Output Shapes

If only terminal commands are needed, return one shell command block.

If the response needs a current command and a prompt for another session, separate them clearly:

```text
Command terminal sekarang
Prompt untuk sesi lain
```

If proof commands are requested separately, label them clearly:

```text
Proof yang dikirim balik
```

## Heredoc Rules

Markdown heredoc files must have all inner code fences closed.

When writing markdown files that contain code fences, the assistant response must use an outer fence that is longer than any inner fence.

Correct behavior:

```text
Use an outer markdown fence longer than the markdown fences inside the generated file.
Keep all prose that belongs to the file inside the heredoc.
End the heredoc cleanly.
Put proof commands after the heredoc terminator as real shell commands.
```

Incorrect behavior:

```text
Do not leave an inner markdown fence unclosed.
Do not put free text after EOF as if it were shell.
Do not mix explanation into a command block after heredoc terminators.
```

## Proof Handling

When the user provides proof commands, keep them as commands.

When the user provides proof output, do not claim more than the output proves.

If local proof exists but connector validation is missing, use this exact status wording:

```text
locally implemented with proof; connector validation pending
```

## Scope Safety

For migration work, do not cross these boundaries unless explicitly requested:

```text
ServiceCatalog
ProductCatalog
Inventory
HTTP runtime
PostgreSQL persistence
route registration
capability seed migration
audit sink
UI
Git mutation
```

The latest owner prompt defines the active scope.

Nearby files or tempting cleanup are not permission to expand the task.

## Relationship To Other Rules

This document supplements:

```text
docs/AGENTS.md
docs/0003_session_start_protocol.md
docs/templates/0122_web_ai_session_prompts.md
docs/scripts/0090_makefile_and_scripts.md
```

If a conflict exists, repository safety, proof, scope control, and owner instructions win.
