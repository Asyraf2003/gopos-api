# AGENTS.md

## Repository instruction source

This Go API workspace uses `docs/` as the canonical AI_RULES package.

Before giving technical guidance, planning implementation, editing files, or proposing commands, read and follow:

1. `docs/0001_index.md`
2. `docs/0002_decision_policy.md`
3. `docs/0003_session_start_protocol.md`
4. `docs/core/0010_scope_and_facts.md`
5. `docs/core/0011_blueprint_first.md`
6. `docs/core/0012_step_by_step_execution.md`
7. `docs/core/0013_proof_and_progress.md`
8. `docs/architecture/0020_hexagonal_go_api.md`
9. `docs/architecture/0021_package_boundaries.md`
10. `docs/architecture/0022_api_capability_control.md`
11. `docs/architecture/0023_public_contracts.md`
12. `docs/architecture/0024_current_repo_layout.md`
13. `docs/domain/0030_domain_contracts.md`
14. `docs/db/0040_postgresql_policy.md`
15. `docs/api/0050_echo_http_contract.md`
16. `docs/testing/0060_test_and_quality_gates.md`
17. `docs/workflow/0070_docs_go_workflow.md`
18. `docs/workflow/0071_handoff_protocol.md`
19. `docs/workflow/0072_transition_progress_ledger_protocol.md`
20. `docs/security/0080_security_baseline.md`
21. `docs/scripts/0090_makefile_and_scripts.md`
22. `docs/style/0100_go_style.md`
23. `docs/templates/0110_domain_scope_packet.md`
24. `docs/templates/0120_prompt_authoring_rules.md`
25. `docs/templates/0121_codex_session_prompts.md`
26. `docs/templates/0122_web_ai_session_prompts.md`
27. `docs/templates/0123_cli_command_formatter_rules.md`

If the user names a blueprint, ADR, handoff, error log, branch, commit, command output, API, domain, table, or module, that reference defines the active scope until the user changes it.

## Mandatory working behavior

- Do not invent facts, repo state, file contents, test results, or completion status.
- Separate FACT, GAP, DECISION, BLUEPRINT, ACTIVE STEP, PROOF, NEXT, and PROGRESS for technical work.
- Start from a blueprint before implementation.
- Use one active step per response.
- Do not continue to the next step without proof and user feedback.
- Progress may increase only when there is real proof.
- User command output is the primary proof.
- The Go API must be pure API, Echo-based, PostgreSQL-backed, and hexagonal.
- UI must consume API contracts dynamically; UI rules must not leak into domain or persistence.
- API capability control must exist as a first-class admin surface and backend policy.
- Every domain database resource must have explicit create, edit/update, delete, show, list, and capability rules unless the domain contract marks an operation forbidden.
- Tests and scripts are required gates, not optional polish.
- Cross-AI work must use a scope packet and handoff.
- Before drafting a next-session prompt, select exactly one target agent and one matching template source.
- Do not write hybrid Web AI/Codex prompts unless the owner explicitly requests a collaboration packet.
- If the target agent is missing, ask one concise clarification question instead of drafting a mixed prompt.
- Web AI is read-only analysis/planning by default and should normally produce command plans for owner/local terminal execution.
- Normal Web AI analysis must omit Codex handoff text unless the owner explicitly requests Codex.
- The default terminal execution channel for Web AI command plans is owner/local terminal.
- Every NEXT or handoff section must name exactly one next execution channel.
- Terminal Codex is the local CLI implementation agent and must not assume Web AI collaboration unless the owner explicitly provided a collaboration packet.
- Web AI with a GitHub connector must use the connector as read-only repository source of truth. Only fetch, search, list, and get actions are allowed by default.
- Web AI must not create, update, or delete files; create branches or commits; update refs; open or merge PRs; create or update issues; comment; label; request reviewers; rerun remote CI; or otherwise mutate GitHub unless the prompt gives exact mutation permission naming the action, target repository, branch, path or issue/PR, and intended content.
- Web AI prompts that say "write docs/...", "update docs/...", "create evidence", "prepare handoff", or "close scope" mean draft paste-ready response content, not repository mutation.
- Docs workflow changes must cascade to impacted README/index/AGENTS/template/audit files when feasible.
- When Laravel source data is missing, ask for the smallest specific source batch by folder, file, route, migration, seeder, test, or command output.
- When an ADR or owner decision is needed, ask a concise question with 2-3 viable options, include tradeoffs, and put the recommended option first when clear.
- Non-trivial final reports must include proof, estimated active-scope progress, context-window status, and next valid step.
- Create or update a handoff before context runs low or when durable work needs continuation.
- For long-running transition scopes, update the active progress ledger when progress, proof, gaps, or next valid step changes.
- Web AI must not provide Git mutation instructions unless the owner explicitly requested Git operations.
- Local proof and remote connector proof must not be conflated; use locally implemented with proof; connector validation pending when remote validation is missing.
- If durable proof changes progress, update, cite, draft, or command-plan the active ledger and relevant handoff before naming the next implementation step.
- Before NEXT, apply the Progress Write Gate when new proof may change project progress.
- When the owner asks for command formatting, follow `docs/templates/0123_cli_command_formatter_rules.md`.
- Makefile/script contracts must stay stable so terminal Codex and GPT web workflows use the same proof language.
