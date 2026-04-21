# AI Rules Summary

## Tujuan
Dokumen ini adalah ringkasan manusiawi dari aturan kerja AI di repo `pos-go`.

## Posisi dokumen ini
- `docs/core/AI_RULES.md` adalah ringkasan singkat
- `docs/AI_RULES/*` adalah konstitusi operasional yang lebih detail
- Jika ada konflik, ikuti urutan referensi di `docs/README.md`

## Ringkasan aturan inti
- jangan berasumsi
- mulai dari blueprint
- kerjakan step-by-step
- satu respons kerja hanya boleh punya satu active step
- jangan klaim selesai tanpa proof
- jangan naikkan progress tanpa bukti nyata
- jaga boundary hexagonal
- jangan melanggar kontrak repo demi convenience

## Mandatory operational reference
Sebelum perubahan besar atau arahan implementasi, baca:
1. `docs/AI_RULES/00_INDEX.md`
2. `docs/AI_RULES/01_DECISION_POLICY.md`
3. `docs/AI_RULES/10_CORE/11_BLUEPRINT_FIRST.md`
4. `docs/AI_RULES/10_CORE/12_STEP_BY_STEP_EXECUTION.md`
5. `docs/AI_RULES/10_CORE/13_PROOF_AND_PROGRESS.md`
6. `docs/AI_RULES/40_ARCHITECTURE/44_AUDIT_AND_DOD.md`
7. `docs/AI_RULES/60_STACK/61_GO_RULES.md`

## Catatan
Jika nanti modul AI_RULES bertambah, dokumen ini cukup diperbarui sebagai ringkasan, bukan tempat semua aturan detail ditumpuk.
