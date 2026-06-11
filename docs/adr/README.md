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

# ADR Index

ADR dipakai untuk mengunci keputusan arsitektur yang penting.

## Kapan wajib membuat ADR
Buat ADR jika perubahan menyentuh salah satu hal berikut:
- boundary arsitektur
- auth flow
- storage strategy
- database contract
- security model
- dependency besar
- pola integrasi pihak ketiga

## Format penamaan
`0001-short-title.md`

## Status yang dipakai
- proposed
- accepted
- superseded
- deprecated

## Aturan
- satu ADR untuk satu keputusan utama
- alasan, alternatif, dan konsekuensi wajib ditulis
- jika accepted, ADR menjadi referensi aktif

## Proof Tracking

ADRs own decisions. Evidence owns proof.

Current ADR implementation proof status is tracked in:

```text
docs/evidence/0004_adr_implementation_proof_index.md
```

When an ADR is accepted, implemented, partially implemented, superseded, or proven by a new gate, update that proof index or state why it is unchanged.
