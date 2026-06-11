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

# ADR 0010: Authorization Admin API Minimum

## Status
accepted

## Context
Setelah auth core live lane selesai, proyek membutuhkan authorization admin API minimum agar role account dapat dikelola secara operasional tanpa SQL manual.

Fondasi yang sudah ada:
- principal resolver sudah dapat membaca roles dan permissions account
- role assignment minimum sudah ada melalui adapter application-level
- middleware auth dan permission enforcement sudah aktif

Kebutuhan minimum saat ini:
- user dapat melihat authz dirinya sendiri
- admin dapat menambahkan role ke account
- admin dapat menghapus role dari account

## Decision
Dipilih authorization admin API minimum dengan tiga endpoint:

1. `GET /api/authz/me`
   - protected
   - menampilkan principal request saat ini
   - response minimum:
     - `account_id`
     - `session_id`
     - `roles`
     - `permissions`
     - `trust_level`

2. `POST /api/admin/accounts/:account_id/roles`
   - protected
   - butuh permission `account.role.assign`
   - request body minimum:
     - `role_key`
   - behavior:
     - assign role ke account secara idempotent

3. `DELETE /api/admin/accounts/:account_id/roles/:role_key`
   - protected
   - butuh permission `account.role.assign`
   - behavior:
     - hapus role dari account
     - minimum guard:
       - role `base` tidak boleh dihapus

## Rules
- endpoint admin wajib memakai auth middleware normal
- endpoint admin wajib memakai permission enforcement normal
- `base` adalah role default minimum dan tidak boleh diremove lewat API
- assign role bersifat idempotent
- remove role untuk role yang tidak ada bersifat idempotent

## Options considered
### Opsi A - langsung pakai SQL/manual
Kelebihan:
- cepat

Kekurangan:
- tidak operasional
- tidak cocok untuk API platform

### Opsi B - admin API minimum
Kelebihan:
- cukup kecil
- operasional
- cocok sebagai pondasi API bisnis berikutnya

Kekurangan:
- masih perlu slice tambahan sebelum API bisnis

### Opsi C - tunggu policy layer penuh dulu
Kelebihan:
- desain lebih lengkap

Kekurangan:
- menunda kebutuhan operasional dasar
- scope membesar

## Consequences
### Positif
- admin/cashier/base bisa dikelola tanpa SQL manual
- API bisnis berikutnya bisa langsung memakai permission check
- scope tetap kecil

### Negatif
- perlu adapter remove role
- perlu handler admin authz baru
- audit trail detail belum masuk di slice ini

## Follow-up
- buat port account role remover
- buat adapter PostgreSQL remove role
- buat usecase dan handler:
  - authz me
  - assign role
  - remove role
- wiring bootstrap
- tambah test dan runtime proof
