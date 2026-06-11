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

# P1 - Go Rules

## Tujuan
Menjaga hygiene implementasi Go tetap konsisten.

## Aturan
- Satu folder = satu package.
- Jaga ukuran file tetap terkontrol; bila melewati batas internal project, harus ada alasan jelas.
- Patuhi boundary dan import discipline.
- Jangan campur domain, transport, dan persistence tanpa jalur yang sah.
- `gofmt` wajib untuk file yang berubah.
- `go test ./...` wajib lulus untuk perubahan yang menyentuh kode Go.
- `go vet ./...` dijalankan bila perubahan sudah menyentuh fondasi runtime atau contract penting.
