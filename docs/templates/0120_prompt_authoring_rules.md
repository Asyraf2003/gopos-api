# Prompt Authoring Rules

Use this file before writing prompts that will be copied into Codex, GPT web, or another AI surface.

## Goal

Prompts must be easy to copy, paste, audit, and continue later.

## Copy-Safe Rules

- Use ASCII characters only unless the source data requires non-ASCII.
- Use straight quotes: `'` and `"`.
- Avoid smart quotes, em dashes, decorative bullets, and hidden characters.
- Do not place a fenced code block inside another fenced code block.
- When the prompt itself needs to mention backticks, write `BACKTICK` instead of typing the character inside a copyable prompt.
- Keep placeholders obvious and uppercase, for example `REPLACE_WITH_SCOPE`.
- Prefer short sections with labels over long paragraphs.
- Put file paths on their own lines.
- Put commands on their own lines.
- Do not paste secrets, tokens, passwords, or private keys.
- Do not say tests passed unless command output is included.

## Required Prompt Sections

Use these sections for serious work:

```text
CONTEXT
ACTIVE SCOPE
FILES TO READ
FILES ALLOWED TO EDIT
FILES FORBIDDEN TO EDIT
TASK
RULES
EXPECTED OUTPUT
PROOF REQUIRED
OPEN GAPS
```

For a tiny question, a shorter prompt is fine.

## Web AI Source Rules

When a web AI has a GitHub connector available:

- Treat the GitHub connector as the source of truth for repository files, branches, pull requests, issues, and committed code.
- Do not ask the web AI to manage local git state unless the task is explicitly about git, GitHub, a branch, a commit, a PR, or CI.
- Do not ask the web AI to run destructive git actions.
- If local and GitHub are expected to be identical because the owner pushes frequently, say that assumption explicitly.
- Use local terminal/Codex context only for things the connector cannot see or cannot do.
- Examples of local-only context: `.env`, secrets, uncommitted sensitive files, generated local output, installed tools, `fd`, `rg`, local test output, local database state, and runtime logs.
- If GitHub connector data and local evidence disagree, mark it as GAP and ask for the smallest proof instead of choosing silently.

## Backtick Guidance For Web AI

Some web AI copy buttons can behave poorly when prompts contain nested Markdown fences.

Safe pattern:

```text
Use plain text blocks for paths and commands.
If you need to output Markdown that contains code fences, describe the fences as BACKTICK BACKTICK BACKTICK.
Do not wrap this whole prompt in Markdown fences.
```

Unsafe pattern:

```text
Do not paste a prompt that contains a Markdown fenced block that itself asks the AI to produce another fenced block.
```

## File Placement Rule

When asking an AI to create structured project knowledge:

- source proof goes to `docs/evidence/`;
- planned work goes to `docs/blueprints/`;
- accepted architecture decisions go to `docs/adr/`;
- session continuation notes go to `docs/handoffs/`;
- obsolete material goes to `docs/archive/`;
- reusable prompts go to `docs/templates/`.

Do not create loose notes at the repository root.
