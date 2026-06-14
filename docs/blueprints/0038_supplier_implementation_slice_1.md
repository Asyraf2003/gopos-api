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

# Supplier Implementation Slice 1

## Status

Closed with local proof.

## Active Scope

Implement Supplier domain, ports, and usecase contracts only.

## FACT

- Supplier domain contract is accepted locally with proof.
- Supplier is the next dependency before Faktur.
- Accepted duplicate policy: active supplier name must be unique by normalized name; inactive supplier names do not block active supplier name reuse.
- Accepted permission model: `supplier.read` and `supplier.manage`.
- Product ID remains the downstream Product reference key.
- Supplier must not own ProductCatalog data.
- Supplier must not mutate stock.
- Supplier must not create Faktur.

## SCOPE-IN

- `internal/modules/supplier/domain`
- `internal/modules/supplier/ports`
- `internal/modules/supplier/usecase`
- Domain unit tests.
- Usecase unit tests with in-memory fake repository.

## SCOPE-OUT

- PostgreSQL persistence.
- HTTP transport.
- Presenter/output DTOs.
- Route registration.
- Capability seed migration.
- Supplier API docs.
- Faktur.
- Inventory mutation.
- Stock movement.
- Audit/outbox.
- Localization.
- Extended filters.

## PROOF REQUIRED

Run:

```bash
go test ./internal/modules/supplier/...
bash scripts/audit_hexagonal.sh
make verify
```

## CLOSEOUT

Supplier implementation slice 1 is closed with local proof.

Remote connector validation was checked on branch visible to the connector and confirms the Supplier domain, ports, and usecase implementation files are present for this slice.

Scope remains domain/ports/usecase only.

Local proof provided by owner:

```bash
go test ./internal/modules/supplier/...
bash scripts/audit_hexagonal.sh
make verify
```

Visible owner result:

- `go test ./internal/modules/supplier/...` passed for supplier domain/usecase, with ports having no test files.
- `bash scripts/audit_hexagonal.sh` passed.
- `make verify` passed aggregate audit, including `go test ./...`, vet audit, format audit, AI rules audit, license header audit, file size audit, hexagonal import audit, route capability audit, and gosec audit.

Remaining gaps stay out of this slice:

- Supplier PostgreSQL persistence.
- Supplier HTTP routes.
- Supplier capability seed.
- Faktur.
- Inventory/stock movement.
- Audit/outbox.
- Localization.
- Extended filters.


## NEXT

After proof, update ledger/handoff.

Next valid active step after this slice:

Supplier PostgreSQL persistence blueprint.
