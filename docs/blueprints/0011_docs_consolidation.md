# Docs Consolidation Blueprint

## FACT
- `docsgo/` contains the newer canonical AI and engineering rules package.
- `docs/` already contains active product documentation: ADRs, blueprints, evidence, handoffs, and archive.
- The latest user scope asks to merge old `docs/` and new `docsgo/` under `docs/`, using `docsgo` as the standard.
- Root `AGENTS.md`, docs references, and `scripts/audit_ai_rules.sh` still point to the split structure.

## GAP
- No source-code behavior is in scope for this consolidation.
- No product blueprint, ADR, handoff, or evidence file should be deleted as part of this scope.

## DECISION
- `docs/` becomes the single canonical documentation root.
- Former `docsgo` standards move into `docs/` with path references rewritten from `docsgo/` to `docs/`.
- Legacy `docs/AI_RULES` and `docs/core` are archived, not mixed into active standards.
- Useful legacy layout guidance is preserved as an active `docs/architecture` standard.
- `docsgo/` is removed after successful consolidation.

## SCOPE-IN
- Move the `docsgo` standards package under `docs/`.
- Preserve existing `docs/adr`, `docs/blueprints`, `docs/evidence`, `docs/handoffs`, and `docs/archive`.
- Archive legacy `docs/AI_RULES` and `docs/core`.
- Update root and docs-local `AGENTS.md`.
- Update path references and AI rules audit script.

## SCOPE-OUT
- Product capability-control implementation.
- Go source changes.
- ADR rewrites.
- Handoff rewrites beyond path reference cleanup.

## PUBLIC CONTRACT IMPACT
- Documentation paths change from `docsgo/...` to `docs/...`.
- No API contract changes.

## DOMAIN/DB/CAPABILITY IMPACT
- No domain logic, DB schema, or runtime capability behavior changes.

## TEST/PROOF PLAN
- Inspect file tree with `fd`.
- Search for stale `docsgo` references with `rg`.
- Run `bash scripts/audit_ai_rules.sh`.
- Inspect git diff summary.

## STEP ORDER
1. Add this blueprint.
2. Archive legacy `docs` rule folders.
3. Move `docsgo` standards into `docs`.
4. Rewrite stale references and update audit script.
5. Remove `docsgo`.
6. Verify with file tree, reference search, and audit script.

## NEXT ACTIVE STEP
Archive legacy docs rule folders and move `docsgo` standards into `docs`.
