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

# P0 - Blueprint First

## Tujuan
Menetapkan bahwa pekerjaan harus dimulai dari blueprint sebelum workflow detail atau implementasi.

## Mandatory Rule
Sebelum implementasi, GPT wajib menyusun blueprint yang menjelaskan:
- target
- kondisi saat ini
- constraints
- scope in
- scope out
- dependensi
- risiko
- outcome yang diinginkan

## Minimum Blueprint Format
- masalah yang sedang diselesaikan
- fakta yang sudah diketahui
- gap yang masih terbuka
- rule yang mengikat
- opsi pendekatan bila ada lebih dari satu jalan
- rekomendasi pendekatan
- urutan step setelah blueprint

## Implementation Gate
Implementasi hanya boleh dimulai jika:
- blueprint sudah cukup jelas
- scope step aktif jelas
- rule P0 relevan sudah dicek
- tidak ada GAP kritis yang membuat implementasi spekulatif

## Forbidden Behavior
- Jangan langsung coding jika blueprint belum jelas.
- Jangan membuka area baru di luar blueprint tanpa menandai scope expansion.
- Jangan menggunakan output implementasi untuk menggantikan proses berpikir blueprint.
