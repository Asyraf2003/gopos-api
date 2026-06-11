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

# ADR 0012: API Output Contract Centralization

## Status
accepted

## Context
Proyek membutuhkan kontrol terpusat untuk seluruh output API user-facing agar:
- contract response mudah diaudit
- bentuk output tidak tersebar liar di handler
- perbedaan bahasa dapat dikelola tanpa membongkar logic aplikasi
- developer dapat melihat seluruh permukaan output API dengan cepat

Kondisi saat ini:
- beberapa handler masih membentuk payload JSON langsung
- beberapa error response masih ditulis langsung di middleware/handler
- JSONB internal persistence sudah terpisah dan tidak boleh dicampur dengan contract output API

## Decision
Dipilih pendekatan centralized API output contracts melalui folder presenter khusus.

Struktur minimum:
- `internal/presentation/http/id/...` untuk output API bahasa Indonesia
- `internal/presentation/http/en/...` disiapkan untuk versi English di masa depan

Rule minimum:
1. handler tidak boleh membentuk payload user-facing secara liar
2. response DTO/presenter harus berada di folder presentation
3. JSON output API dan JSONB persistence harus tetap dipisah
4. perubahan bahasa output tidak boleh memaksa perubahan business logic
5. handler cukup:
   - ambil input
   - panggil usecase
   - panggil presenter/output contract
   - kirim response

## Scope awal
Refactor dilakukan bertahap, dimulai dari endpoint yang sudah ada:
- `/api/me`
- health response
- auth response berikutnya

## Options considered

### Opsi A - biarkan DTO dekat di handler
Kelebihan:
- cepat

Kekurangan:
- tetap tersebar
- sulit diaudit
- sulit disiapkan untuk i18n

### Opsi B - centralized presenter folder per bahasa
Kelebihan:
- mudah diaudit
- siap i18n
- handler lebih tipis
- contract output terlihat jelas

Kekurangan:
- perlu refactor bertahap

### Opsi C - satu file besar semua output
Kelebihan:
- terlihat terpusat

Kekurangan:
- cepat menjadi file raksasa
- bertentangan dengan disiplin file kecil

## Consequences

### Positif
- output API mudah dipetakan
- i18n lebih murah
- repo police dapat mengaudit contract output lebih tegas

### Negatif
- perlu refactor endpoint yang sudah ada
- perlu aturan audit tambahan agar handler tidak kembali liar

## Follow-up
- buat folder `internal/presentation/http/id`
- pindahkan contract output endpoint awal ke presenter
- refactor handler `/api/me` sebagai baseline
- buat audit rule untuk melarang payload user-facing liar di handler
