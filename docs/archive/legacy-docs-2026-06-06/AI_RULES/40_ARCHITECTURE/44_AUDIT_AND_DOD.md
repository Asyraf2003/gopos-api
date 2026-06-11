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

# P1 - Audit and DoD

## Tujuan
Menjadikan auditability dan Definition of Done sebagai bagian wajib dari delivery.

## Mandatory Rule
- Perubahan penting harus dapat diaudit.
- Klaim "selesai" harus ditopang oleh verification yang relevan terhadap scope step.
- DoD mengikuti konteks perubahan, tetapi tidak boleh kosong.

## Typical DoD Components
Tergantung konteks, DoD dapat mencakup:
- format/lint
- test
- audit
- sanity check
- inspection file/output

## Proof Rule
Jika menyebut verifikasi:
- sertakan command atau artefak
- sertakan hasil
- sertakan arti hasil terhadap step aktif

## Forbidden Behavior
- Jangan menulis DoD seolah selesai jika baru rencana.
- Jangan menulis verifikasi abstrak tanpa bukti konkret.
