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

# ADR 0006: Request Principal Model

## Status
accepted

## Context
Setelah identity, session, role, dan permission foundation selesai, API membutuhkan model principal yang konsisten untuk request yang sudah terautentikasi.

Model ini harus:
- cukup kecil untuk fase awal
- cukup kaya untuk permission enforcement
- dapat dipakai bersama oleh middleware, usecase, debug lane, dan audit trail
- tidak memaksa API mengambil keputusan langsung dari nama role

## Decision
Digunakan model request principal minimum berikut:

### Principal fields
- account_id
- session_id
- roles
- permissions
- trust_level

### Principal behavior
Principal menyediakan helper minimum:
- `HasPermission(permissionKey string) bool`

### Source of truth
- `account_id` dan `session_id` berasal dari access token claim
- `roles` dan `permissions` di-resolve dari persistence layer
- `trust_level` berasal dari access token claim

### Request flow
1. middleware memverifikasi access token
2. middleware mengekstrak identity/session claim minimum
3. resolver memuat roles dan permissions dari persistence layer
4. principal ditempel ke context request
5. handler/usecase/guard memakai principal untuk keputusan authorization

### Rule
API tidak boleh mengambil keputusan hanya dari nama role bila permission check sudah tersedia.
Role adalah grouping.
Permission adalah capability yang diperiksa.

## Options considered
### Opsi A - token-only principal
Kelebihan:
- cepat

Kekurangan:
- role/permission di token cepat stale
- perubahan authorization tidak langsung berlaku

### Opsi B - minimal claim + DB-backed principal
Kelebihan:
- stabil
- cocok untuk fase awal
- perubahan role/permission langsung efektif

Kekurangan:
- butuh resolver tambahan

### Opsi C - DB lookup penuh tanpa claim session
Kelebihan:
- sederhana di token

Kekurangan:
- kehilangan ikatan session dari token
- kurang baik untuk revoke/debug/audit

## Consequences
### Positif
- satu model principal dipakai lintas layer
- permission enforcement lebih presisi
- debug lane dan audit trail bisa memakai bentuk yang sama

### Negatif
- perlu middleware verifier
- perlu resolver principal

## Follow-up
- buat domain principal minimum
- buat port principal resolver
- buat JWT verifier minimum
- buat middleware auth context
- buat permission guard
