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

## Prompt Target Selection Rule

Before drafting any next-session prompt, identify exactly one target agent:

- Terminal Codex;
- Web AI.

Then select exactly one matching template:

- Terminal Codex uses `docs/templates/0121_codex_session_prompts.md`;
- Web AI uses `docs/templates/0122_web_ai_session_prompts.md`.

Do not combine Terminal Codex and Web AI instructions in one prompt.

If the owner does not specify the target agent, ask one concise clarification question instead of drafting a prompt.

Create a collaboration packet only when the owner explicitly requests Web AI/Codex collaboration.

Every generated next-session prompt must include:

- target agent;
- template source;
- active scope;
- one active step;
- files to read first;
- allowed files for Terminal Codex implementation prompts;
- forbidden files;
- proof requirements;
- expected output;
- handoff, progress, and context-window requirement.

Final self-check before sharing a generated prompt:

- exactly one target agent is named;
- exactly one template source is named;
- no hybrid Terminal Codex/Web AI instructions are present;
- proof commands are grouped in the correct proof section for that target agent;
- no git commands are included unless the owner explicitly asked for them.

## AI Execution Channel Rule

Web AI is a read-only analysis and planning agent by default.

Web AI normally returns:

- analysis;
- patch plans;
- command plans;
- docs text;
- evidence text;
- validation notes.

Web AI may provide commands for the owner/local terminal to run.

Web AI must not assume Terminal Codex is the executor.

Terminal Codex is a local CLI implementation agent. Terminal Codex works through local CLI execution and owner feedback.

Owner/local terminal may execute command plans, collect proof, and push changes.

Collaboration packets involving Web AI, Terminal Codex, owner/local terminal, and repository state are special-case only. They require explicit owner instruction for a specific problem.

If the next executor is unclear, ask one concise clarification question.

Do not convert a Web AI analysis task into a Terminal Codex handoff unless the owner asked for Codex.

Web AI analysis output must not include Codex handoff text unless the owner explicitly requested Codex.

If a Web AI output includes a Codex handoff without explicit owner request, treat it as a template-compliance failure.

Default terminal execution channel for Web AI command plans is owner/local terminal, not Terminal Codex.

The next execution channel must be named explicitly.

Normal Web AI loop:

```text
owner -> Web AI -> GitHub connector read-only -> Web AI -> owner -> owner/local terminal -> owner pushes repo -> Web AI reads repo via connector -> loop
```

Normal Terminal Codex loop:

```text
owner -> Terminal Codex -> local CLI -> owner pushes repo -> Terminal Codex continues through local CLI/user feedback -> loop
```

Explicit collaboration loop:

```text
Web AI + Terminal Codex + owner/local terminal + repo may collaborate only when the owner explicitly requests that mode.
```

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

## Working Style Rule

For coding or migration work, prefer one focused work batch over many tiny back-and-forth prompts.

The prompt should ask the AI to:

- execute the largest safe slice that still fits one active step;
- send short progress updates while working;
- avoid long speculative discussion before inspecting files;
- stop and ask only when missing data blocks the decision;
- ask for the smallest specific Laravel source batch when source data is missing;
- ask ADR-level questions with 2-3 options, tradeoffs, and a recommended option when clear;
- finish with a compact report: files changed, proof, progress, context-window status, and next step.

This is usually cheaper than many small chats because repeated prompts often resend or restate context.

## Web AI Source Rules

When a web AI has a GitHub connector available:

- Treat the GitHub connector as the read-only source of truth for repository files, branches, pull requests, issues, and committed code.
- GitHub connector access is read-only by default.
- Allow only read-only fetch, search, list, and get connector actions by default.
- Forbid connector mutation actions by default, including create_file, update_file, delete_file, create_branch, update_ref, create_commit, create_pull_request, merge_pull_request, create_issue, update_issue, issue comments, PR comments, labels, reviewer requests, and remote CI reruns.
- Require exact mutation permission before any GitHub mutation action. The permission must name the action, target repository, branch, path or issue/PR, and intended content.
- Treat "write docs/...", "update docs/...", "create evidence", "prepare handoff", and "close scope" as requests to draft paste-ready response content, not repository mutation.
- Do not ask the web AI to manage local git state unless the task gives exact permission for a specific git or GitHub action.
- Do not ask the web AI to run destructive git actions.
- If local and GitHub are expected to be identical because the owner pushes frequently, say that assumption explicitly.
- Use local terminal/Codex context only for things the connector cannot see or cannot do.
- Examples of local-only context: `.env`, secrets, uncommitted sensitive files, generated local output, installed tools, `fd`, `rg`, local test output, local database state, and runtime logs.
- If GitHub connector data and local evidence disagree, mark it as GAP and ask for the smallest proof instead of choosing silently.
- Do not claim local commands, tests, database checks, runtime checks, or git-status checks were run unless exact output is provided.
- Put local proof that Web AI cannot run under `PROOF THE TERMINAL AGENT MUST RUN`.

Web AI execution model:

- Web AI drafts patch plans, exact shell commands, docs text, evidence text, and handoff text.
- Owner/local terminal normally runs command plans and applies repository changes unless exact GitHub mutation permission is provided.
- Terminal Codex runs command plans only when the owner has targeted Terminal Codex or requested a Codex handoff.
- Web AI must separate proposed commands from commands that actually ran.
- Web AI should produce copyable CLI commands for owner/local terminal when terminal execution is required.

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
