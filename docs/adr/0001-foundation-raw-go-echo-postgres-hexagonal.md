# ADR 0001: Foundation Raw Go Echo PostgreSQL Hexagonal

## Status
accepted

## Context
Proyek `pos-go` dibuat sebagai fondasi API POS baru yang tidak lagi bergantung pada struktur proyek lama yang terlalu besar untuk fase belajar dan fondasi awal. Diperlukan arsitektur yang:
- cukup kecil untuk dipahami
- cukup disiplin untuk dikembangkan jangka panjang
- siap menerima auth Google, JWT, dan security enhancement bertahap
- tidak memaksa refactor besar saat domain POS mulai tumbuh

## Decision
Dipilih keputusan berikut:
- bahasa utama: Go
- HTTP adapter: Echo
- database utama: PostgreSQL
- gaya arsitektur: hexagonal / ports and adapters
- struktur proyek dibuat modular dan ringan
- auth dibangun bertahap dari kontrak minimum
- detail provider eksternal ditempatkan di platform adapter
- domain tidak boleh bergantung ke Echo, PostgreSQL driver, Google SDK, atau KMS SDK

## Options considered
- layered biasa tanpa port/adaptor
- menyalin penuh struktur proyek lama
- hexagonal modular ramping

## Consequences
### Positif
- fondasi lebih tahan terhadap perubahan
- auth, postgres, google, dan crypto bisa ditumbuhkan tanpa bongkar total
- batas tanggung jawab lebih jelas

### Negatif
- ada boilerplate awal
- implementasi awal lebih lambat daripada menaruh semua logika di handler

### Trade-off
Dipilih disiplin kontrak dari awal, tetapi implementasi fitur tetap dibuat bertahap agar tidak overbuilt.

## Follow-up
- hidupkan runtime API minimum
- hidupkan PostgreSQL connection layer
- bentuk auth contracts minimum
- lanjutkan auth flow Google versi API-first
