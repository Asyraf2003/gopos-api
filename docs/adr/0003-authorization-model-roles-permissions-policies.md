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

# ADR 0003: Authorization Model Roles Permissions Policies

## Status
accepted

## Context
Setelah auth minimum berhasil dari login Google sampai persistence session dan issue token, proyek membutuhkan model authorization yang:
- cukup kecil untuk fase awal
- cukup disiplin untuk berkembang
- tidak mengikat keputusan API ke nama role semata
- mendukung jalur live dan jalur debug

## Decision
Authorization model menggunakan pendekatan berikut:

### Identity vs authorization
- `account` adalah identitas principal
- `session` adalah status login
- `role` adalah grouping permission
- `permission` adalah capability yang diperiksa API
- `policy` adalah aturan keputusan pada level usecase/resource

### Default role
Setiap account baru otomatis memiliki role `base`.

### Initial roles
- `base`
- `cashier`
- `admin`

### Initial rule
API tidak boleh mengambil keputusan bisnis hanya dari nama role.
API harus memeriksa permission.
Role dipakai sebagai alat grouping permission.

### Initial authorization storage
Akan digunakan tabel:
- `roles`
- `permissions`
- `account_roles`
- `role_permissions`

### Initial permission strategy
Permission disimpan sebagai key string stabil, contoh:
- `auth.session.refresh`
- `auth.session.logout`
- `profile.self.read`
- `sale.order.create`
- `inventory.manage`
- `account.role.assign`

### Token strategy
Access token fase awal tidak akan menyimpan seluruh role/permission.
Token tetap ringan dan membawa identity/session claim minimum.
Role/permission akan di-resolve dari persistence layer pada request path yang memerlukan authorization.

### Live vs debug lane
- live lane memakai Google OIDC asli
- debug lane hanya aktif di local/debug mode
- debug lane harus tetap melewati pipeline auth utama, bukan bypass mentah

## Options considered
### Opsi A - role only
Kelebihan:
- sederhana

Kekurangan:
- cepat mentok
- logic API jadi bergantung ke nama role

### Opsi B - direct permission per account
Kelebihan:
- fleksibel

Kekurangan:
- sulit dikelola
- tidak ergonomis untuk operasi harian

### Opsi C - roles + permissions + policies
Kelebihan:
- seimbang
- mudah tumbuh
- keputusan API tetap presisi

Kekurangan:
- butuh tabel tambahan
- butuh resolver permission

## Consequences
### Positif
- authorization lebih stabil
- role tidak menjadi hardcoded business logic
- jalur admin/kasir/base bisa tumbuh tanpa refactor besar

### Negatif
- perlu migration tambahan
- perlu resolver principal + permission

## Follow-up
- buat migration authorization tables
- seed role dan permission minimum
- buat principal resolver
- buat permission guard
- lanjut ke refresh/logout/debug lane
