# ADR 0009: Debug Auth Lane Strategy

## Status
accepted

## Context
Setelah live auth lane selesai dari login Google sampai refresh/logout runtime proof, proyek membutuhkan debug auth lane agar skenario authorization dapat diuji cepat tanpa login Google berulang.

Debug lane ini harus:
- hanya aktif di environment debug/local
- tidak aktif di production
- tetap memakai contract auth internal yang sama semaksimal mungkin
- mendukung skenario role dasar seperti `base`, `cashier`, dan `admin`

## Decision
Dipilih strategi debug auth lane berbasis debug-only HTTP endpoints yang digate oleh config explicit.

Flow minimum:
1. endpoint debug auth hanya diregister jika debug mode aktif
2. endpoint dapat membuat session debug dan issue token memakai issuer yang sama
3. endpoint dapat memilih principal minimum untuk skenario uji
4. hasil token dari debug lane tetap melewati middleware auth normal saat dipakai ke protected route

## Rules
- debug lane tidak boleh aktif default di production
- debug lane wajib digate oleh config explicit
- debug lane hanya untuk local/development verification
- token hasil debug lane harus kompatibel dengan verifier dan middleware normal
- jalur debug dan live harus berbagi contract principal/session sebanyak mungkin

## Options considered
### Opsi A - login Google terus untuk semua test
Kelebihan:
- paling mendekati live flow

Kekurangan:
- lambat
- merepotkan untuk test matrix role/permission
- tidak efisien untuk development

### Opsi B - debug-only auth endpoints
Kelebihan:
- cepat
- mudah untuk test role matrix
- tetap bisa reuse issuer/verifier/middleware yang sama

Kekurangan:
- perlu pagar config yang disiplin

### Opsi C - hardcode bypass di middleware
Kelebihan:
- cepat dibuat

Kekurangan:
- berisiko
- merusak boundary auth
- mudah bocor ke environment yang salah

## Consequences
### Positif
- skenario auth bisa diuji cepat
- live lane tidak terganggu
- permission dan policy lebih mudah diverifikasi

### Negatif
- perlu config gate tambahan
- perlu route debug khusus dan test tambahan

## Follow-up
- tambah config explicit untuk debug auth lane
- buat ADR/contract endpoint debug minimum
- implement debug issue token endpoint
- tambah runtime proof untuk base/cashier/admin scenario
