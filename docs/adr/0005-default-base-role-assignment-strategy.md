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

# ADR 0005: Default Base Role Assignment Strategy

## Status
accepted

## Context
Setelah schema authorization, seed role/permission, dan backfill account existing selesai, proyek membutuhkan strategi agar setiap account baru memiliki role default `base`.

Strategi ini harus:
- eksplisit
- mudah diuji
- tidak menyembunyikan logika penting di database
- tidak mencampur repository identity terlalu jauh dengan provisioning authorization

## Decision
Dipilih strategi application-level idempotent role ensure.

Flow minimum:
- setelah account berhasil di-resolve atau dibuat pada auth flow
- sistem harus memastikan account memiliki role `base`
- operasi ini harus idempotent
- implementasi dilakukan melalui port dan adapter PostgreSQL

## Options considered
### Opsi A - database trigger
Kelebihan:
- otomatis di level DB

Kekurangan:
- logika tersembunyi
- sulit diuji
- sulit diaudit dalam flow aplikasi

### Opsi B - inject langsung di repository create account
Kelebihan:
- cepat

Kekurangan:
- repository identity menjadi menanggung logika authorization
- coupling meningkat

### Opsi C - application-level ensure role
Kelebihan:
- eksplisit
- mudah diuji
- idempotent
- cocok dengan boundary hexagonal

Kekurangan:
- perlu port dan adapter tambahan

## Consequences
### Positif
- account baru dan lama bisa diperlakukan konsisten
- jalur live dan debug bisa memakai mekanisme yang sama
- default role tidak tergantung trigger DB

### Negatif
- perlu tambahan adapter authorization kecil
- usecase auth perlu satu langkah tambahan

## Follow-up
- buat port role assignment minimum
- buat adapter PostgreSQL idempotent
- panggil ensure base role dari auth flow
- tambah test
