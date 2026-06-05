# Current Repo Layout

## Purpose

Preserve the active repository layout contract while the main architecture rules stay framework-agnostic enough for future Go API work.

## Active Baseline

- Main language: Go.
- HTTP adapter: Echo.
- Database: PostgreSQL.
- Architecture: hexagonal / ports and adapters.
- Product mode: API-first.
- Auth direction: token-based API flow with current auth and authorization contracts preserved by ADRs.

## Current Layout

```text
cmd/
  api/
internal/
  app/bootstrap/
  config/
  modules/<module>/
    domain/
    ports/
    usecase/
    transport/http/
  platform/
    postgres/
  transport/http/
    middleware/
migrations/
scripts/
docs/
```

## Layer Contracts

- `cmd/api`: runtime entrypoint only.
- `internal/config`: load and validate runtime configuration.
- `internal/app/bootstrap`: dependency assembly and route wiring.
- `internal/modules/<module>/domain`: entities, value objects, invariants, semantic errors.
- `internal/modules/<module>/ports`: module interfaces and boundary contracts.
- `internal/modules/<module>/usecase`: orchestration and transaction intent.
- `internal/modules/<module>/transport/http`: request/response mapping and use case calls.
- `internal/transport/http/middleware`: cross-module HTTP middleware.
- `internal/platform`: technical adapters such as PostgreSQL, tokens, state, clocks, IDs, providers, and crypto.
- `migrations`: PostgreSQL schema changes.
- `scripts`: repeatable proof and audit commands.
- `docs`: standards, blueprints, ADRs, evidence, handoffs, and archive.

## Protected Contracts

These paths are protected until an ADR or active blueprint changes them:

- `cmd/api/main.go` as the API runtime entrypoint.
- `internal/config/*` as the active config source.
- `internal/app/bootstrap/*` as runtime wiring.
- `internal/modules/*/ports/*` as module boundary contracts.
- Public API response envelopes defined by active ADRs and `docs/api/0050_echo_http_contract.md`.

## Change Rule

If a change modifies layout, dependency direction, module ownership, or a protected contract:

- update the active blueprint;
- update this document when the layout contract changes;
- create or update an ADR when the change affects an architecture decision;
- show proof through diff inspection and the relevant test or audit command.
