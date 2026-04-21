# Decision Policy

## Status
Dokumen ini adalah conflict protocol utama untuk AI yang bekerja di `pos-go`.

## Tujuan
Menetapkan cara mengambil keputusan saat:
- ada benturan antar rule
- ada benturan antara kecepatan dan correctness
- ada data yang belum cukup
- ada godaan untuk mengisi kekosongan dengan asumsi

## Rule Hierarchy
1. ADR accepted mengalahkan dokumen lain.
2. BLUEPRINT mengalahkan workflow dan preferensi delivery.
3. STRUCTURE dan AI_RULES mengalahkan convenience implementasi.
4. P0 mengalahkan P1.
5. P1 mengalahkan P2.
6. Aturan yang lebih spesifik mengalahkan aturan yang lebih umum.
7. Bukti nyata mengalahkan dugaan.

## P0 Modules
- `10_CORE/11_BLUEPRINT_FIRST.md`
- `10_CORE/12_STEP_BY_STEP_EXECUTION.md`

## P1 Modules
- `10_CORE/13_PROOF_AND_PROGRESS.md`
- `40_ARCHITECTURE/44_AUDIT_AND_DOD.md`
- `60_STACK/61_GO_RULES.md`

## Mandatory Decision Sequence
Setiap kali mengambil keputusan, GPT wajib urut seperti ini:
1. identifikasi fakta yang terbukti
2. identifikasi tujuan step aktif
3. identifikasi scope in dan scope out
4. identifikasi rule P0 yang relevan
5. identifikasi dampak ke boundary dan contract
6. identifikasi apakah data cukup
7. bila data tidak cukup, tandai GAP dan stop perluasan klaim

## GAP Rule
Jika data belum cukup:
- tulis apa yang belum diketahui
- tulis kenapa kekurangan itu menghambat keputusan
- jangan isi dengan tebakan
- jangan menyamarkan GAP seolah fakta

## Forbidden Shortcuts
- Tidak boleh mengklaim status repo tanpa bukti.
- Tidak boleh mengklaim file sudah benar tanpa inspeksi atau output.
- Tidak boleh mengklaim test pass tanpa output test.
- Tidak boleh mengklaim requirement user bila user belum menyatakannya.
- Tidak boleh menaikkan progress jika belum ada proof.

## Stop Conditions
GPT harus berhenti dan menyatakan GAP jika:
- source of truth tidak jelas
- blueprint belum cukup untuk implementasi
- proof yang dibutuhkan belum ada
- keputusan baru akan mengubah arah fondasi tanpa bukti kuat
