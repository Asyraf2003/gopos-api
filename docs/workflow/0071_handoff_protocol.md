# Handoff Protocol

## Purpose

Work may move between terminal Codex, GPT web, humans, and future AI assistants.

This protocol keeps scope, proof, and file ownership clear.

## Required Handoff Fields

Every handoff must include:

- date;
- active scope;
- current branch or source snapshot if available;
- files included;
- files changed;
- files forbidden to touch;
- blueprint referenced;
- ADR/rules referenced;
- decisions made;
- proof collected;
- tests or commands run;
- gaps still open;
- next valid active step.
- estimated scope progress percentage;
- estimated context-window status.

If the handoff recommends a next AI session, it must also include:

- target agent for the next session;
- template source for the next prompt;
- one next active step;
- files the next agent should read;
- files the next agent may edit, when the target is Terminal Codex;
- files the next agent must not edit;
- proof commands;
- progress percentage;
- context-window status.

## Handoff Target Types

A handoff that recommends further action must name one next execution channel:

- owner/local terminal;
- Terminal Codex;
- Web AI;
- explicit collaboration packet.

Use owner/local terminal when the owner can run a command plan and then push repository changes for later Web AI connector review.

Use Terminal Codex only when the owner asks for local implementation by Codex.

Use Web AI only when the owner asks for read-only analysis, planning, review, or paste-ready text through the connector-backed Web AI loop.

Use explicit collaboration packet only when the owner explicitly requests Web AI, Terminal Codex, owner/local terminal, and repository state to collaborate for a specific problem.

Every next section or handoff must identify exactly one next execution channel.

For normal Web AI analysis, the next execution channel should usually be owner/local terminal unless the owner named another channel.

## Automatic Handoff Triggers

Create or update a handoff without waiting for the owner to repeat the request when:

- the session is near context limit;
- a long-running scope has changed files, evidence, blueprint, or decisions;
- work moves from Codex to web AI or from web AI to Codex;
- a model switch is likely;
- the owner asks to pause;
- the active step ends with open GAP items that the next session must know.

If the session is small and no durable decision, file change, or evidence was created, a handoff is optional.

## Cascading Documentation Rule

When a docs update changes how future work should start, continue, verify, or hand off, update the impacted chain in the same active step when feasible:

- human entry point: `docs/README.md`;
- AI bootstrap: `docs/AGENTS.md`;
- canonical index: `docs/0001_index.md`;
- local folder README;
- related workflow/template docs;
- audit script when the new file or rule is mandatory.

## Scope Packet Rule

Before sending work to another AI, create a scope packet.

The packet must say:

- target agent;
- template source;
- what domain/API is active;
- which files are provided as context;
- which files may be edited;
- which files are read-only;
- which output format is expected;
- which command output counts as proof.

## Receiving AI Rule

The receiving AI must:

- read `docs/README.md`;
- read `docs/AGENTS.md`;
- read the scope packet;
- stay inside active scope;
- mark GAP instead of guessing missing files;
- return a handoff when done.

## Forbidden Handoff Behavior

- Do not say "continue from previous chat" without a written handoff.
- Do not omit the target agent from a next-session prompt.
- Do not write "Web AI or Codex next session" as a single prompt.
- Do not mix Web AI and Codex instructions unless the owner explicitly requested a collaboration packet.
- Do not write "send to Codex" when owner/local terminal can run the command plan and Web AI can validate after push.
- Do not make `HANDOFF TEXT FOR CODEX` mandatory Web AI output.
- Do not include Codex handoff sections in Web AI output by default.
- Do not treat `OPTIONAL HANDOFF TEXT FOR CODEX` as a harmless extra section.
- Do not include extra Codex handoff text unless the owner requested Codex.
- Do not leave a handoff or NEXT section without exactly one next execution channel: owner/local terminal, Web AI, Terminal Codex, or explicit collaboration packet.
- Do not use collaboration mode unless it says the owner explicitly requested it.
- Do not send only code files without rules and proof requirements.
- Do not let archive docs override active blueprint.
- Do not ask another AI to edit files outside the packet.
- Do not claim tests passed from intention.
- Do not leave a new mandatory workflow rule only in chat.

## Final Report Rule

For non-trivial work, the final response must include a compact status report:

- files or docs changed;
- proof run;
- estimated progress for the active scope;
- estimated context-window status;
- next valid active step.

Progress estimates are directional. They must not replace proof.
