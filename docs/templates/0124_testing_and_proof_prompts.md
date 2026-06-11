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

# Testing And Proof Prompts

Use these prompts when asking an AI to design, run, or interpret proof.

## Ask For A Test Plan

```text
Create a test plan for this active step.

ACTIVE STEP
REPLACE_WITH_STEP

BLUEPRINT
REPLACE_WITH_BLUEPRINT_OR_SUMMARY

CODE OR DESIGN
REPLACE_WITH_CONTEXT

REQUIRED TEST TYPES
- domain unit tests;
- usecase tests;
- PostgreSQL adapter tests;
- API contract tests;
- capability-disabled tests;
- migration/constraint tests when DB changes exist.

OUTPUT
- test cases;
- test files;
- proof commands;
- residual risk if only focused tests are run.
```

## Ask Codex To Run Proof

```text
Run proof for the current active step.

COMMANDS
REPLACE_WITH_FOCUSED_COMMANDS
make verify

RULES
- Run focused proof first.
- Then run broader proof if the focused proof passes.
- If a command fails, report the failing command and key output.
- Do not edit files unless the failure is clearly in scope.
```

## Ask Web AI To Interpret Test Output

```text
Interpret this test output.

ACTIVE SCOPE
REPLACE_WITH_SCOPE

COMMAND
REPLACE_WITH_COMMAND

OUTPUT
REPLACE_WITH_OUTPUT

TASK
- Identify the actual failure.
- Separate test setup failures from product behavior failures.
- Suggest the smallest next debug step.
- Do not claim the fix without code proof.
```

## Proof Summary Template

```text
PROOF SUMMARY

Active scope:
REPLACE_WITH_SCOPE

Commands run:
REPLACE_WITH_COMMANDS

Passed:
REPLACE_WITH_PASSED

Failed:
REPLACE_WITH_FAILED

Key output:
REPLACE_WITH_KEY_OUTPUT

Conclusion:
REPLACE_WITH_CONCLUSION

Residual risk:
REPLACE_WITH_RISK
```
