# Domain Scope Packet Template

Use this template when sending one domain/API scope to another AI or another session.

## Scope

Active domain/API:

```text
REPLACE_WITH_DOMAIN_OR_API
```

Goal:

```text
REPLACE_WITH_GOAL
```

## Included Files

Read-only context:

```text
REPLACE_WITH_FILES
```

Editable files:

```text
REPLACE_WITH_FILES
```

Forbidden files:

```text
REPLACE_WITH_FILES
```

## Required Rules To Read

- `docs/README.md`
- `docs/AGENTS.md`
- `docs/0002_decision_policy.md`
- `docs/architecture/0020_hexagonal_go_api.md`
- `docs/architecture/0022_api_capability_control.md`
- `docs/domain/0030_domain_contracts.md`
- `docs/api/0050_echo_http_contract.md`
- `docs/testing/0060_test_and_quality_gates.md`
- `docs/workflow/0071_handoff_protocol.md`

## Domain Contract

Tables:

```text
REPLACE_WITH_TABLES
```

Allowed operations:

```text
create:
update:
delete:
show:
list:
```

Forbidden operations:

```text
REPLACE_WITH_FORBIDDEN_OPERATIONS
```

Capability keys:

```text
REPLACE_WITH_CAPABILITY_KEYS
```

## Expected Output

The receiving AI must return:

- files changed;
- summary of implementation;
- proof commands run or commands the owner must run;
- gaps;
- handoff for next step;
- estimated progress percentage;
- estimated context-window status.

## Working Style

The receiving AI should:

- execute the largest safe slice that still fits one active step;
- keep progress updates short;
- keep the final report compact;
- ask for the smallest specific source batch if source data is missing;
- ask ADR-level decisions with 2-3 options, tradeoffs, and a recommended option when clear.

## Proof Required

Minimum proof:

```bash
make verify
```

Focused proof:

```bash
REPLACE_WITH_FOCUSED_COMMANDS
```

## Handoff

Next valid active step:

```text
REPLACE_WITH_NEXT_STEP
```
