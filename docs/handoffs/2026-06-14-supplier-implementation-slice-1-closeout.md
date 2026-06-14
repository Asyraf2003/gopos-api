<!--
Copyright (C) 2026 Asyraf Mubarak

This file is part of gopos-api.

gopos-api is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, version 3 only.

gopos-api is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with gopos-api. If not, see <https://www.gnu.org/licenses/>.
-->

# Supplier Implementation Slice 1 Closeout

## Status

Closed with local proof.

Remote connector validation confirms the implementation files are visible on the repository branch available to the connector.

## FACT

- Supplier implementation slice 1 is limited to domain, ports, and usecase contracts.
- Supplier domain model exists.
- Supplier status supports active and inactive.
- Supplier contact fields include phone, email, address, and notes.
- Supplier normalization and validation exist.
- Supplier lifecycle behavior exists for create, update, activate, and deactivate.
- Supplier repository port exists.
- Supplier repository duplicate lookup exists for normalized supplier names.
- Supplier usecases exist for create, update, show, list, lookup, activate, and deactivate.
- Duplicate policy is enforced:
  - active Supplier name must be unique by normalized name;
  - inactive Supplier names do not block active Supplier name reuse;
  - reactivating inactive Supplier rejects if another active Supplier already owns the same normalized name.

## Proof

Owner local proof:

```bash
go test ./internal/modules/supplier/...
bash scripts/audit_hexagonal.sh
make verify
```

Visible owner result:

- `go test ./internal/modules/supplier/...` passed.
- `bash scripts/audit_hexagonal.sh` passed.
- `make verify` passed aggregate audit.

## Remote Connector Validation Status

Remote connector validation confirms the Supplier implementation slice 1 files are visible and match the expected domain/ports/usecase-only scope.

## Gap

Still incomplete and intentionally not started in this slice:

- Supplier PostgreSQL persistence.
- Supplier HTTP routes.
- Supplier capability seed.
- Faktur.
- Inventory/stock movement.
- Audit/outbox.
- Localization.
- Extended filters.

Auth/System ADR 0012 output contract centralization remains deferred by owner decision.

## Next

Supplier PostgreSQL persistence blueprint.

## Scope Guard

Do not start HTTP transport, capability seed, Faktur, stock movement, audit/outbox, localization, extended filters, or architecture cleanup from this closeout.
