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

# P1 - Proof and Progress

## Tujuan
Memastikan progress selalu terkait langsung dengan bukti nyata, bukan keyakinan atau proposal.

## Mandatory Rule
- Progress tidak boleh naik tanpa proof.
- Setiap klaim selesai harus menunjuk ke bukti nyata.
- Setelah satu step workflow selesai, tampilkan progress dalam persen.

## Accepted Proof
Proof yang valid dapat berupa:
- output command
- isi file
- diff yang terverifikasi
- hasil test
- hasil verifikasi manual
- snapshot atau ADR yang eksplisit

## Mandatory Proof Structure
Setiap proof minimal harus menjelaskan:
- command atau artefak
- hasil yang terlihat
- arti hasil terhadap step aktif

## Progress Rule
- Progress merepresentasikan status workflow, bukan sekadar banyaknya teks atau ide.
- Proposal tanpa eksekusi tidak menaikkan progress.
- Struktur file yang baru dibuat boleh menaikkan progress hanya jika memang target step adalah pembentukan struktur itu.
- Revisi rule hanya menaikkan progress jika file benar-benar sudah berubah dan diverifikasi.

## Forbidden Behavior
- Jangan mengklaim hijau tanpa output.
- Jangan mengklaim selesai jika baru menulis rencana.
- Jangan memanipulasi progress untuk terlihat maju.
