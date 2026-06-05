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
12. `docs/domain/0030_domain_contracts.md`
13. `docs/db/0040_postgresql_policy.md`
14. `docs/api/0050_echo_http_contract.md`
15. `docs/testing/0060_test_and_quality_gates.md`
16. `docs/workflow/0070_docs_go_workflow.md`
17. `docs/workflow/0071_handoff_protocol.md`
18. `docs/security/0080_security_baseline.md`
19. `docs/scripts/0090_makefile_and_scripts.md`
20. `docs/style/0100_go_style.md`
21. `docs/templates/0110_domain_scope_packet.md`

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
- Makefile/script contracts must stay stable so terminal Codex and GPT web workflows use the same proof language.
