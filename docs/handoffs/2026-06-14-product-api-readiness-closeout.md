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

# Product API Readiness Closeout

## Status

Closed with local proof.

## FACT

- Product API readiness for Supplier/Faktur planning is documented.
- Product ID is accepted as the downstream Product reference key.
- Product lookup/list/show endpoints are the intended Product selection path for future Supplier/Faktur flows.
- Supplier/Faktur must not depend on ProductCatalog PostgreSQL internals.
- Supplier/Faktur must not treat `kode_barang` as always present.
- Supplier/Faktur must not infer stock-on-hand from ProductCatalog threshold fields.
- Auth/System output contract centralization remains deferred by owner decision.

## PROOF

Owner local terminal proof passed:

```text
rg -n 'Product API Readiness|REQUIRED PRODUCT CONTRACT|READINESS CHECKLIST|NEXT DOMAIN PATH|product_id|Supplier domain contract blueprint' docs/blueprints/0036_product_api_readiness_for_supplier_faktur_slice.md
make verify

[PASS] blueprint marker check
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] license header audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed
```

## DECISION

Product API is sufficiently ready as a dependency boundary for Supplier and Faktur planning.

This does not mean inventory or stock mutation is implemented.

## GAP

- Supplier domain contract is not accepted yet.
- Supplier implementation is not started.
- Faktur domain contract is not accepted yet.
- Faktur implementation is not started.
- Product inventory/stock mutation is not implemented.
- Stock adjustment create/reverse is not implemented.
- ProductCatalog audit/outbox persistence is not implemented.
- Runtime localization/language switching is not implemented.
- Extended Laravel filters are not implemented.

## NEXT

Next valid active step:

```text
Supplier domain contract blueprint.
```

Scope guard:

```text
Do not implement Supplier before the Supplier domain contract is accepted.
Do not start Faktur before Supplier unless an explicit owner decision changes the order.
Do not start inventory mutation, stock movement, audit/outbox, localization, extended filters, or architecture folder rename work from this closeout.
```
