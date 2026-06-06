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
19. `docs/security/0080_security_baseline.md`
20. `docs/scripts/0090_makefile_and_scripts.md`
21. `docs/style/0100_go_style.md`
22. `docs/templates/0110_domain_scope_packet.md`
23. `docs/templates/0120_prompt_authoring_rules.md`
24. `docs/templates/0121_codex_session_prompts.md`
25. `docs/templates/0122_web_ai_session_prompts.md`

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
- Web AI with a GitHub connector must use the connector as repository source of truth and must not manage git unless the active task explicitly asks for GitHub, git, CI, branch, commit, or PR work.
- Docs workflow changes must cascade to impacted README/index/AGENTS/template/audit files when feasible.
- Non-trivial final reports must include proof, estimated active-scope progress, context-window status, and next valid step.
- Create or update a handoff before context runs low or when durable work needs continuation.
- Makefile/script contracts must stay stable so terminal Codex and GPT web workflows use the same proof language.
