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
