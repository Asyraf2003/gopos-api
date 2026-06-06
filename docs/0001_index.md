# Go API AI_RULES Index

## Status

This document is the mandatory entry point for AI work in the future Go Echo API workspace.

`docs/` is the canonical documentation root. It combines active project documentation with the clean rules package for a pure Go API + PostgreSQL project.

## Purpose

The package exists so future work becomes strict, repeatable, and easy to audit.

It prevents:

- mixed architecture boundaries;
- undocumented API behavior;
- controller-driven business logic;
- database access outside persistence adapters;
- UI assumptions inside core domain;
- uncontrolled API capability exposure;
- handoff, blueprint, archive, and implementation docs overwriting each other;
- claims of completion without test or command proof.

## Mandatory Read Order

1. `0002_decision_policy.md`
2. `0003_session_start_protocol.md`
3. `core/0010_scope_and_facts.md`
4. `core/0011_blueprint_first.md`
5. `core/0012_step_by_step_execution.md`
6. `core/0013_proof_and_progress.md`
7. `architecture/0020_hexagonal_go_api.md`
8. `architecture/0021_package_boundaries.md`
9. `architecture/0022_api_capability_control.md`
10. `architecture/0023_public_contracts.md`
11. `architecture/0024_current_repo_layout.md`
12. `domain/0030_domain_contracts.md`
13. `db/0040_postgresql_policy.md`
14. `api/0050_echo_http_contract.md`
15. `testing/0060_test_and_quality_gates.md`
16. `workflow/0070_docs_go_workflow.md`
17. `workflow/0071_handoff_protocol.md`
18. `security/0080_security_baseline.md`
19. `scripts/0090_makefile_and_scripts.md`
20. `style/0100_go_style.md`
21. `templates/0110_domain_scope_packet.md`
22. `templates/0120_prompt_authoring_rules.md`
23. `templates/0121_codex_session_prompts.md`
24. `templates/0122_web_ai_session_prompts.md`

## Constitution Summary

- Blueprint first.
- One active step.
- Proof before progress.
- Hexagonal architecture is mandatory.
- PostgreSQL is the database target.
- Echo is the HTTP transport.
- API contracts are public contracts.
- UI must be dynamic over API contracts and capability metadata.
- Capability control is a core system feature.
- Every domain operation must declare authorization, capability, transaction, audit, and test requirements.
- A file or package may only own one kind of responsibility.

## Canonical Package Map

- `core/`: AI behavior, scope, proof, and step discipline.
- `architecture/`: Go hexagonal architecture, package boundaries, public API contracts, and capability control.
- `domain/`: domain CRUD and mutation contract rules.
- `db/`: PostgreSQL schema, repository, migration, and transaction rules.
- `api/`: Echo HTTP contract and response envelope rules.
- `testing/`: unit, integration, contract, API, architecture, and script gates.
- `workflow/`: docs hygiene, blueprint/handoff/archive separation, and work sequencing.
- `security/`: authentication, authorization, capability, input, output, secret, audit, and abuse-control rules.
- `scripts/`: Makefile and command contracts for repeatable proof.
- `style/`: Go writing style and forbidden patterns.
- `templates/`: copyable work packet, Codex prompt, web AI prompt, proof, evidence, and resume templates.
- `adr/`, `blueprints/`, `evidence/`, `handoffs/`, `archive/`: active project decisions, plans, proof, continuation notes, and historical material.

## Non-Negotiable Behavior

- Do not implement from memory when a source file can be inspected.
- Do not make a domain decision in transport, SQL, or UI.
- Do not add an endpoint without a use case contract.
- Do not add a repository method without a port or use case need.
- Do not add a table without migration, constraints, repository tests, and rollback/forward policy.
- Do not expose an API without capability metadata.
- Do not claim tests passed without visible output.
- Do not hand work to another AI without a scope packet and handoff target.
- For web AI with a GitHub connector, treat the connector as read-only repository source of truth. Allow only fetch, search, list, and get connector actions by default.
- Web AI must not mutate files, branches, commits, refs, pull requests, issues, labels, reviewers, merges, or remote CI unless the prompt gives exact mutation permission naming the action, target repository, branch, path or issue/PR, and intended content.
- Web AI prompts that say "write docs/...", "update docs/...", "create evidence", "prepare handoff", or "close scope" mean draft paste-ready response content, not repository mutation.
- Do not leave mandatory workflow updates only in chat; cascade impacted docs and audit rules.
- For non-trivial work, final status must include proof, estimated active-scope progress, context-window status, and next valid step.
