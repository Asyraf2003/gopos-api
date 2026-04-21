# ADR 0004: Authorization Minimum Schema

## Status
accepted

## Context
Setelah model authorization roles + permissions + policies dipilih, proyek membutuhkan schema minimum yang:
- cukup kecil untuk fase awal
- cukup kuat untuk POS yang membutuhkan pembatasan aksi
- mendukung role default untuk semua account
- dapat berkembang ke policy dan audit trail tanpa refactor besar

## Decision
Schema authorization minimum menggunakan tabel:

### 1. roles
Tujuan:
- daftar role stabil untuk grouping permission

Field minimum:
- id
- key
- name
- created_at

Constraint minimum:
- unique (key)

### 2. permissions
Tujuan:
- daftar capability stabil yang diperiksa API

Field minimum:
- id
- key
- name
- created_at

Constraint minimum:
- unique (key)

### 3. account_roles
Tujuan:
- relasi many-to-many account ke role

Field minimum:
- account_id
- role_id
- created_at

Constraint minimum:
- unique (account_id, role_id)

### 4. role_permissions
Tujuan:
- relasi many-to-many role ke permission

Field minimum:
- role_id
- permission_id
- created_at

Constraint minimum:
- unique (role_id, permission_id)

## Default rule
- setiap account baru harus memiliki role `base`

## Initial roles
- `base`
- `cashier`
- `admin`

## Initial permission examples
- `auth.session.refresh`
- `auth.session.logout`
- `profile.self.read`
- `sale.order.create`
- `sale.order.read`
- `payment.create`
- `inventory.manage`
- `report.read`
- `account.role.assign`

## Options considered
### Opsi A - role only
Kelebihan:
- sederhana

Kekurangan:
- keputusan API jadi bergantung langsung ke nama role
- cepat mentok

### Opsi B - direct permission per account
Kelebihan:
- fleksibel

Kekurangan:
- sulit dikelola
- tidak ergonomis untuk operasi harian

### Opsi C - roles + permissions join tables
Kelebihan:
- seimbang
- stabil
- mudah berkembang ke policy

Kekurangan:
- perlu tabel tambahan
- perlu resolver permission

## Consequences
### Positif
- role default bisa diterapkan konsisten
- permission lebih presisi daripada role name
- policy layer bisa ditambahkan tanpa mengubah schema inti

### Negatif
- perlu seed data minimum
- perlu resolver principal dan permission guard

## Follow-up
- buat migration authorization tables
- buat seed minimum role dan permission
- buat principal resolver
- buat permission guard
