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

# API Capability Control

## Purpose

The system must have one administrative control surface for API capability exposure.

This exists so features can be enabled, disabled, focused, audited, and discovered dynamically by UI and operators.

## Capability Concept

An API capability is a named operation such as:

- `products.create`
- `products.update`
- `products.delete`
- `products.show`
- `products.list`
- `transactions.close`
- `payments.refund`

## Required Fields

Every capability must declare:

- key;
- domain;
- operation;
- endpoint method/path;
- default state;
- required permission;
- risk level;
- audit requirement;
- idempotency requirement;
- owner package;
- test proof reference.

## Runtime Rule

Every protected endpoint must pass:

```text
authn -> authz -> capability check -> request validation -> use case
```

Capability checks must happen before business mutation.

## Admin Control Surface

The API must provide an admin-only surface to:

- list capabilities;
- show capability detail;
- enable capability;
- disable capability;
- explain why a capability is disabled;
- expose capability metadata for dynamic UI behavior.

## Forbidden Behavior

- Do not hide feature flags in handlers.
- Do not use environment variables as the only capability control.
- Do not let UI decide whether a forbidden API is actually allowed.
- Do not allow disabled mutation endpoints to reach use cases.
- Do not create endpoint routes without capability registry entries, unless the endpoint is explicitly public infrastructure such as health.

