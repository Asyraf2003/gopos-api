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

# Project Documentation Index

Dokumen di folder ini adalah pengendali utama arah proyek.

## Kelompok dokumen
- `docs/core/*` = ringkasan arah, struktur, workflow, dan quality gate proyek
- `docs/AI_RULES/*` = konstitusi operasional detail untuk AI assistant
- `docs/adr/*` = keputusan arsitektur yang sudah dikunci

## Urutan rujukan
1. `docs/adr/*` dengan status accepted
2. `docs/core/BLUEPRINT.md`
3. `docs/core/STRUCTURE.md`
4. `docs/core/AI_RULES.md`
5. `docs/AI_RULES/00_INDEX.md`
6. `docs/AI_RULES/01_DECISION_POLICY.md`
7. `docs/AI_RULES/10_CORE/11_BLUEPRINT_FIRST.md`
8. `docs/AI_RULES/10_CORE/12_STEP_BY_STEP_EXECUTION.md`
9. `docs/AI_RULES/10_CORE/13_PROOF_AND_PROGRESS.md`
10. `docs/AI_RULES/40_ARCHITECTURE/44_AUDIT_AND_DOD.md`
11. `docs/AI_RULES/60_STACK/61_GO_RULES.md`
12. `docs/core/WORKFLOW.md`
13. `docs/core/DOD.md`

## Aturan kerja
- Jangan menambah fitur atau struktur yang melanggar blueprint aktif.
- Workflow boleh berubah hanya bila ada bukti kebutuhan nyata, konflik teknis, atau keputusan ADR baru.
- Definition of Done adalah pagar minimum. Sesuatu belum dianggap selesai jika DoD belum terpenuhi.
- AI operational detail harus mengikuti `docs/AI_RULES/*`, bukan menebak dari kebiasaan.
