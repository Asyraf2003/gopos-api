# ADR 0002: Auth Minimum Data Model

## Status
accepted

## Context
Setelah fondasi runtime API dan koneksi PostgreSQL terbukti hidup, proyek membutuhkan model data auth minimum yang cukup kecil untuk fase awal tetapi cukup stabil untuk berkembang ke:
- login Google
- access token / refresh token flow
- session persistence
- penambahan security enhancement di fase berikutnya

Model ini tidak boleh terlalu kecil sampai cepat mentok, dan tidak boleh terlalu besar sampai membebani fase fondasi.

## Decision
Dipilih model data auth minimum berikut:

### 1. accounts
Tujuan:
- menjadi identitas account internal aplikasi
- menjadi referensi utama untuk session dan external identity

Field minimum:
- id
- email
- created_at
- updated_at

### 2. auth_identities
Tujuan:
- memetakan account internal ke identity provider eksternal seperti Google

Field minimum:
- id
- account_id
- provider
- subject
- email
- email_verified
- meta_json
- created_at
- updated_at

Constraint minimum:
- unique (provider, subject)

### 3. auth_sessions
Tujuan:
- menyimpan sesi login server-side dan refresh token hash

Field minimum:
- id
- account_id
- refresh_token_hash
- expires_at
- revoked_at
- meta_json
- created_at

Constraint minimum:
- foreign key account_id -> accounts.id

## Options considered
### Opsi A - minimal auth foundation
- accounts
- auth_identities
- auth_sessions

Kelebihan:
- cukup untuk Google login dan session flow
- tetap kecil
- mudah tumbuh

Kekurangan:
- audit dan trust belum masuk

### Opsi B - super minimal user table only
Kelebihan:
- cepat dibuat

Kekurangan:
- cepat mentok
- identity provider dan session jadi bercampur

### Opsi C - mature auth model penuh
Kelebihan:
- future-ready sangat jauh

Kekurangan:
- terlalu besar untuk fase sekarang

## Consequences
### Positif
- fondasi auth menjadi stabil
- Google login bisa masuk tanpa mengubah model inti
- session persistence punya tempat yang jelas

### Negatif
- perlu migration SQL sejak awal
- ada sedikit boilerplate ekstra dibanding model super minimal

## Follow-up
- buat migration SQL untuk ketiga tabel
- buat ports repository/session store yang sesuai
- implementasikan postgres adapter bertahap
