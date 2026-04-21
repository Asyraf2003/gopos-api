# ADR 0008: Refresh Token Rotation Strategy

## Status
accepted

## Context
Setelah login, protected route, dan logout/revoke berhasil, API auth masih membutuhkan refresh token flow agar client tidak perlu melakukan login Google ulang setiap kali access token expired.

Strategi refresh harus:
- cukup aman untuk fase awal
- cocok dengan schema session yang sudah ada
- mudah diuji
- tidak menambah kompleksitas berlebihan

## Decision
Dipilih strategi refresh token rotation in-place pada session yang sama.

Flow minimum:
1. client mengirim refresh token
2. sistem menghitung hash refresh token
3. sistem mencari session aktif berdasarkan hash refresh token
4. jika session valid dan belum revoked, sistem menerbitkan:
   - access token baru
   - refresh token baru
5. sistem mengganti refresh_token_hash dan expires_at pada session yang sama
6. session_id tetap

## Rule
- refresh token mentah tidak disimpan di database
- yang disimpan tetap hash refresh token
- refresh hanya berlaku untuk session aktif dengan `revoked_at IS NULL`
- refresh token lama tidak boleh tetap valid setelah rotasi berhasil

## Options considered
### Opsi A - rotate in-place pada session yang sama
Kelebihan:
- sederhana
- cocok dengan schema sekarang
- mudah diuji

Kekurangan:
- histori rotasi tidak sekaya model session baru

### Opsi B - revoke session lama dan buat session baru
Kelebihan:
- audit lifecycle lebih detail
- lebih ketat

Kekurangan:
- kompleksitas lebih tinggi
- butuh plumbing tambahan

### Opsi C - stateless refresh
Kelebihan:
- sederhana di server

Kekurangan:
- tidak cocok dengan model session persistence sekarang
- kurang baik untuk revoke/logout

## Consequences
### Positif
- client bisa memperoleh access token baru tanpa login ulang
- logout/revoke tetap bermakna
- implementasi tetap kecil

### Negatif
- perlu port dan adapter refresh session lookup/update
- perlu endpoint refresh khusus

## Follow-up
- buat port refresh session repository minimum
- buat usecase refresh token
- buat adapter PostgreSQL untuk lookup by refresh token hash dan rotate
- buat endpoint `/api/auth/refresh`
- tambah test unit, integration, dan runtime proof
