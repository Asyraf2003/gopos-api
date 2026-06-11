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

# DEFINITION OF DONE

## Done minimum untuk perubahan kode
Suatu pekerjaan baru dianggap selesai bila semua yang relevan di bawah ini terpenuhi:

- scope perubahan jelas
- file yang diubah benar-benar tertulis di repo
- `gofmt -w` sudah dijalankan pada file Go yang diubah
- `go test ./...` lulus
- perubahan tidak melanggar blueprint aktif
- jika menyentuh boundary besar, auth flow, atau storage contract, ADR/blueprint sudah diperbarui bila diperlukan
- jika menambah config baru, dokumentasi config ikut diperbarui
- jika menambah endpoint, ada bukti route/handler shape yang jelas
- jika menambah persistence, kontrak port dan adapter sama-sama jelas

## Done minimum untuk perubahan docs
- dokumen baru berada di lokasi yang konsisten
- dokumen tidak bertentangan dengan blueprint aktif
- bila dokumen mengubah keputusan arsitektur, perubahan itu juga tercermin di ADR atau blueprint

## Done minimum untuk pekerjaan fondasi
- bisa di-compile
- bisa diuji
- punya proof output
- tidak menyisakan keputusan kritis yang masih implisit
