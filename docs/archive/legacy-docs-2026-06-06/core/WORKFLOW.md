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

# WORKFLOW

## Workflow default
1. pastikan scope aktif
2. baca blueprint yang aktif
3. baca ADR yang relevan
4. pisahkan fakta, gap, dan keputusan
5. tentukan satu active step
6. eksekusi satu langkah kecil
7. buktikan hasil dengan output nyata
8. baru lanjut ke langkah berikutnya

## Aturan eksekusi
- satu active step per siklus kerja
- keputusan harus berbasis bukti repo, output command, atau dokumen aktif
- jangan lompat ke implementasi besar saat fondasi belum terbukti
- jangan buka scope baru tanpa alasan yang jelas
- semua perubahan besar pada boundary, auth, storage, atau security harus dirujuk ke ADR

## Kapan workflow boleh berubah
Workflow boleh disesuaikan bila:
- langkah default terbukti menghambat
- ada dependency teknis baru
- ada konflik dengan blueprint atau ADR
- ada bukti bahwa urutan kerja saat ini tidak lagi efisien

## Bukti minimum
Contoh bukti yang dianggap valid:
- output command
- struktur file/folder
- compile success
- go test pass
- route/endpoint hidup
- migrasi berhasil
- isi file yang benar-benar sudah tertulis

## Larangan
- jangan anggap selesai tanpa proof
- jangan membuat keputusan arsitektur hanya dari asumsi
- jangan memasukkan fitur “sekalian” di luar active step
