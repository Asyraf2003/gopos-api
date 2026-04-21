# AI_RULES Index

## Status
Dokumen ini adalah entrypoint wajib untuk setiap GPT/AI assistant yang bekerja di repo `pos-go`.

## Tujuan
AI_RULES mengunci cara kerja AI agar:
- tidak berasumsi
- tidak keluar dari blueprint aktif
- tidak melompati step aktif
- tidak mengarang fakta, status repo, hasil test, atau keputusan
- tetap patuh pada boundary dan contract project

## Mandatory Read Order
Setiap GPT wajib membaca urutan ini sebelum memberi arahan kerja:

1. `01_DECISION_POLICY.md`
2. `10_CORE/11_BLUEPRINT_FIRST.md`
3. `10_CORE/12_STEP_BY_STEP_EXECUTION.md`
4. `10_CORE/13_PROOF_AND_PROGRESS.md`
5. `40_ARCHITECTURE/44_AUDIT_AND_DOD.md`
6. `60_STACK/61_GO_RULES.md`

## Constitution Summary
- Jangan berasumsi.
- Semua arahan harus berbasis fakta, kondisi saat ini, tujuan step, dan bukti.
- Mulai dari blueprint.
- Setelah blueprint, susun workflow step-by-step.
- Satu respons kerja hanya boleh punya satu step aktif.
- Setelah satu step aktif selesai, tunggu feedback user.
- Progress hanya boleh naik jika ada proof nyata.
- Jangan membuka keputusan arsitektur besar tanpa bukti konflik atau kebutuhan nyata.

## Priority Model
- P0 = rule inti, tidak boleh dilanggar tanpa keputusan eksplisit
- P1 = workflow enforcement dan architecture alignment
- P2 = delivery preference dan hygiene pendukung

## Operational Bootstrap for GPT
Sebelum menjawab, GPT wajib memastikan:
1. apa fakta yang benar-benar ada
2. apa tujuan step saat ini
3. apa scope in dan scope out
4. rule P0 apa yang mengikat
5. apakah data cukup untuk melanjutkan
6. bila data tidak cukup, berhenti di GAP

## Module Map
- `01_DECISION_POLICY.md`
- `10_CORE/11_BLUEPRINT_FIRST.md`
- `10_CORE/12_STEP_BY_STEP_EXECUTION.md`
- `10_CORE/13_PROOF_AND_PROGRESS.md`
- `40_ARCHITECTURE/44_AUDIT_AND_DOD.md`
- `60_STACK/61_GO_RULES.md`

## Non-Negotiable Behavior
- Dilarang mengarang fakta.
- Dilarang mengklaim progress tanpa proof.
- Dilarang langsung lompat ke implementasi bila blueprint belum jelas.
- Dilarang menjadikan format output lebih penting daripada correctness teknis.
- Dilarang menyamakan proposal dengan eksekusi selesai.

## Conflict Reminder
Jika ada konflik:
1. dahulukan P0
2. dahulukan aturan yang lebih spesifik
3. dahulukan correctness architecture dan boundary
4. jika data kurang, berhenti di GAP
