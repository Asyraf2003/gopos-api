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

# ProductCatalog API

## Status

Developer-facing contract for the currently implemented ProductCatalog backend API runtime.

## Base Path

```text
/api/products
```

## Authentication and Authorization

All ProductCatalog routes are protected.

Read routes require:

```text
permission: product_catalog.read
```

Write and lifecycle routes require:

```text
permission: product_catalog.manage
```

Each route is also guarded by its route capability key.

## Response Envelope

Success responses use:

```json
{
  "success": true,
  "data": {},
  "meta": {}
}
```

Error responses use:

```json
{
  "success": false,
  "error": {
    "code": "product_not_found",
    "message": "product not found"
  },
  "meta": {}
}
```

## Routes

| Method | Path | Permission | Capability |
| --- | --- | --- | --- |
| GET | `/api/products` | `product_catalog.read` | `product_catalog.list` |
| POST | `/api/products` | `product_catalog.manage` | `product_catalog.create` |
| GET | `/api/products/lookup` | `product_catalog.read` | `product_catalog.lookup` |
| GET | `/api/products/:id` | `product_catalog.read` | `product_catalog.show` |
| PUT | `/api/products/:id` | `product_catalog.manage` | `product_catalog.update` |
| DELETE | `/api/products/:id` | `product_catalog.manage` | `product_catalog.delete` |
| PATCH | `/api/products/:id/restore` | `product_catalog.manage` | `product_catalog.restore` |
| GET | `/api/products/:id/versions` | `product_catalog.read` | `product_catalog.versions` |

## Public Product Fields

```text
kode_barang
nama_barang
merek
ukuran
harga_jual
reorder_point_qty
critical_threshold_qty
```

## GET /api/products

Query parameters:

```text
q
page
per_page
status = active|deleted|all
```

Example response:

```json
{
  "success": true,
  "data": [
    {
      "id": "product-id",
      "kode_barang": "SKU-001",
      "nama_barang": "Kampas Rem",
      "merek": "Honda",
      "ukuran": 14,
      "harga_jual": 40000,
      "status": "active"
    }
  ],
  "meta": {}
}
```

## POST /api/products

Example request:

```json
{
  "kode_barang": "SKU-001",
  "nama_barang": "Kampas Rem",
  "merek": "Honda",
  "ukuran": 14,
  "harga_jual": 40000,
  "reorder_point_qty": 5,
  "critical_threshold_qty": 2,
  "reason": "created from API"
}
```

Example response:

```json
{
  "success": true,
  "data": {
    "id": "product-id",
    "kode_barang": "SKU-001",
    "nama_barang": "Kampas Rem",
    "nama_barang_normalized": "kampas rem",
    "merek": "Honda",
    "merek_normalized": "honda",
    "ukuran": 14,
    "harga_jual": 40000,
    "reorder_point_qty": 5,
    "critical_threshold_qty": 2,
    "status": "active",
    "created_at": "2026-06-13T00:00:00Z",
    "updated_at": "2026-06-13T00:00:00Z"
  },
  "meta": {}
}
```

## GET /api/products/lookup

Query parameters:

```text
q
limit
include_deleted
```

Example response:

```json
{
  "success": true,
  "data": [
    {
      "id": "product-id",
      "kode_barang": "SKU-001",
      "nama_barang": "Kampas Rem",
      "merek": "Honda",
      "ukuran": 14,
      "harga_jual": 40000,
      "status": "active"
    }
  ],
  "meta": {}
}
```

## GET /api/products/:id

Example response:

```json
{
  "success": true,
  "data": {
    "id": "product-id",
    "kode_barang": "SKU-001",
    "nama_barang": "Kampas Rem",
    "nama_barang_normalized": "kampas rem",
    "merek": "Honda",
    "merek_normalized": "honda",
    "ukuran": 14,
    "harga_jual": 40000,
    "reorder_point_qty": 5,
    "critical_threshold_qty": 2,
    "status": "active"
  },
  "meta": {}
}
```

## PUT /api/products/:id

Example request:

```json
{
  "kode_barang": "SKU-001",
  "nama_barang": "Kampas Rem Updated",
  "merek": "Honda",
  "ukuran": 14,
  "harga_jual": 45000,
  "reorder_point_qty": 6,
  "critical_threshold_qty": 3,
  "reason": "price update"
}
```

Example response:

```json
{
  "success": true,
  "data": {
    "id": "product-id",
    "kode_barang": "SKU-001",
    "nama_barang": "Kampas Rem Updated",
    "nama_barang_normalized": "kampas rem updated",
    "merek": "Honda",
    "merek_normalized": "honda",
    "ukuran": 14,
    "harga_jual": 45000,
    "reorder_point_qty": 6,
    "critical_threshold_qty": 3,
    "status": "active",
    "updated_at": "2026-06-13T00:00:00Z",
    "revision_no": 2
  },
  "meta": {}
}
```

## DELETE /api/products/:id

Example request:

```json
{
  "reason": "discontinued"
}
```

Example response:

```json
{
  "success": true,
  "data": {
    "id": "product-id",
    "status": "deleted",
    "deleted_at": "2026-06-13T00:00:00Z",
    "revision_no": 3
  },
  "meta": {}
}
```

## PATCH /api/products/:id/restore

Example request:

```json
{
  "reason": "restored by admin"
}
```

Example response:

```json
{
  "success": true,
  "data": {
    "id": "product-id",
    "status": "active",
    "restored_at": "2026-06-13T00:00:00Z",
    "revision_no": 4
  },
  "meta": {}
}
```

## GET /api/products/:id/versions

Example response:

```json
{
  "success": true,
  "data": [
    {
      "product_id": "product-id",
      "revision_no": 1,
      "event_name": "product.created",
      "changed_by_actor_id": "account-id",
      "change_reason": "created from API",
      "changed_at": "2026-06-13T00:00:00Z"
    }
  ],
  "meta": {}
}
```

## Error Codes

| HTTP Status | Code | Message |
| --- | --- | --- |
| 400 | `invalid_request_body` | invalid request body |
| 400 | `invalid_query_parameter` | Query-specific validation message |
| 400 | `product_validation_failed` | Domain validation message |
| 401 | `authentication_required` | Authentication error message |
| 403 | `forbidden` | forbidden |
| 403 | `capability_disabled` | capability disabled |
| 404 | `product_not_found` | product not found |
| 409 | `product_code_already_exists` | product code already exists |
| 409 | `product_identity_already_exists` | product identity already exists |
| 500 | `product_catalog_request_failed` | product catalog request failed |

## Scope Notes

This API does not expose ProductCatalog UI behavior.

This API does not mutate inventory stock.

This API does not create or reverse stock adjustments.

This API does not expose the extended Laravel table filters yet:

```text
sort_by
sort_dir
merek
ukuran_min
ukuran_max
harga_min
harga_max
stok_saat_ini
```
