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
