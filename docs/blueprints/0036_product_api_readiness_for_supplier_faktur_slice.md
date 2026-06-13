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

# Product API Readiness For Supplier And Faktur Path

## Status

Accepted locally for readiness analysis.

This is a readiness and contract-stabilization slice. It must not implement Supplier, Faktur, inventory mutation, stock movement, audit/outbox, localization, extended filters, or architecture folder cleanup.

## Active Scope

Review and stabilize the existing ProductCatalog API as the dependency path toward Supplier and Faktur.

The goal is to define whether Product API contracts are sufficient for the next business-domain slices.

## FACT

- ProductCatalog already has protected HTTP routes for list, create, lookup, versions, restore, show, update, and delete.
- ProductCatalog routes are wired through auth, permission, and capability middleware.
- ProductCatalog public fields include `kode_barang`, `nama_barang`, `merek`, `ukuran`, `harga_jual`, `reorder_point_qty`, and `critical_threshold_qty`.
- ProductCatalog PostgreSQL schema stores product identity, normalized lookup fields, sale price, reorder threshold, critical threshold, soft delete metadata, timestamps, and product versions.
- ProductCatalog does not currently model stock-on-hand as a product field.
- ProductCatalog runtime smoke proof already proved local HTTP/auth/DB behavior for selected protected ProductCatalog routes.
- Shared success envelope coverage for ProductCatalog, ServiceCatalog, and Capability is closed locally.
- Auth/System output contract centralization is deferred by owner decision and must not block Product/Supplier/Faktur progress.

## GAP

- There is no explicit Product API readiness checklist for Supplier/Faktur dependency use.
- Supplier and Faktur domain contracts are not accepted yet.
- Product inventory/stock mutation is not implemented.
- Stock adjustment create/reverse is not implemented.
- ProductCatalog audit/outbox persistence is not implemented.
- Runtime localization/language switching is not implemented.
- Extended Laravel filters are not implemented.
- There is no accepted rule yet for whether Faktur will directly mutate stock or only record purchase/sales documents in the first slice.

## DECISION

Treat Product API as the next dependency boundary.

Before starting Supplier or Faktur implementation:

- Freeze the current Product API identity fields for dependency use.
- Confirm lookup behavior is sufficient for selecting products from Supplier/Faktur flows.
- Confirm product ID is the stable reference key for downstream domains.
- Confirm `kode_barang` remains optional but unique for active products when provided.
- Confirm product name, brand, and size remain the human-readable identity tuple.
- Confirm price and threshold fields are ProductCatalog-owned.
- Do not add stock-on-hand to ProductCatalog in this readiness slice.
- Do not implement Supplier in this readiness slice.
- Do not implement Faktur in this readiness slice.
- Do not introduce invoice-numbering, stock movement, payment, or accounting rules in this readiness slice.

## READINESS CHECKLIST

Product API is ready for Supplier/Faktur planning when all of these are true:

- List endpoint is documented and protected.
- Lookup endpoint is documented and protected.
- Show endpoint is documented and protected.
- Create/update endpoint request fields are documented.
- Soft delete/restore behavior is documented.
- Versions endpoint is documented as ProductCatalog history, not audit/outbox.
- Response envelope is canonical.
- Error envelope is canonical for ProductCatalog errors.
- Product ID is documented as the downstream reference key.
- Product display identity is documented for UI/search use.
- Optional `kode_barang` behavior is documented.
- Duplicate behavior is documented.
- ProductCatalog scope explicitly excludes stock-on-hand mutation.
- ProductCatalog scope explicitly excludes Supplier/Faktur ownership.

## SCOPE-IN

- Product API readiness blueprint.
- Product API dependency checklist for Supplier/Faktur.
- Documentation-only clarification if gaps are found.
- Focused connector/file inspection.
- Local `make verify` proof after docs changes.

## SCOPE-OUT

- New migrations.
- Product inventory mutation.
- Stock adjustment create/reverse.
- Supplier implementation.
- Faktur implementation.
- Payment/accounting behavior.
- Audit/outbox persistence.
- Runtime localization.
- Extended filters.
- Auth/System output contract centralization.
- Router/bootstrap cleanup.
- ProductCatalog persistence rewrite.

## REQUIRED PRODUCT CONTRACT FOR NEXT DOMAINS

Downstream domains must reference products by:

```text
product_id
```

Human-facing selection should use existing ProductCatalog read models:

```text
GET /api/products/lookup
GET /api/products
GET /api/products/:id
```

Supplier/Faktur must not depend on ProductCatalog internal PostgreSQL columns directly.

Supplier/Faktur must not treat `kode_barang` as always present.

Supplier/Faktur must not infer stock-on-hand from ProductCatalog threshold fields.

## NEXT DOMAIN PATH

Recommended order after this readiness slice:

```text
1. Supplier domain contract
2. Supplier implementation slice 1
3. Faktur domain contract
4. Faktur implementation slice 1
5. Stock movement design, if Faktur must affect inventory
```

Do not start Faktur before Supplier unless an owner decision explicitly changes the business order.

## PROOF REQUIRED

Run:

```bash
rg -n 'Product API Readiness|REQUIRED PRODUCT CONTRACT|READINESS CHECKLIST|NEXT DOMAIN PATH|product_id' docs/blueprints/0036_product_api_readiness_for_supplier_faktur_slice.md
make verify
```

## NEXT

After this blueprint is accepted and proven, the next valid active step should be:

```text
Supplier domain contract blueprint.
```
