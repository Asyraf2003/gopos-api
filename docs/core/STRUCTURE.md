# STRUCTURE

## Tujuan
Dokumen ini mengunci layout repo dan kontrak antar layer agar perubahan tetap presisi, mudah diaudit, dan tidak menimbulkan efek berantai yang tidak perlu.

## Baseline aktif
- bahasa utama: Go
- HTTP adapter: Echo
- database utama: PostgreSQL
- arsitektur: hexagonal / ports and adapters
- mode produk aktif: API-first
- auth target: Google login + token-based API flow
- provider cloud dan KMS belum aktif di fase ini, tetapi boundary harus siap untuk pertumbuhan ke arah itu

## Layout repo
### `cmd/`
Entrypoint binary.
- `cmd/api`: API utama

### `internal/config`
Load dan validasi konfigurasi runtime.

### `internal/app/bootstrap`
Wiring runtime minimum dan dependency assembly.

### `internal/modules/<module>/`
Struktur standar modul:
- `domain/` : entity, value object, invariants, semantic errors
- `ports/` : interface dependency / contract
- `usecase/` : orchestration flow
- `transport/http/` : HTTP adapter per modul

### `internal/transport/http/`
HTTP core lintas modul:
- `middleware/`

### `internal/platform/`
Adapter teknis eksternal:
- `postgres/`
- `google/`
- `token/jwt/`
- `state/memory/`
- `crypto/` (future)
- `kms/` (future)

### `migrations/`
Perubahan schema PostgreSQL.

### `docs/`
Pengendali keputusan, workflow, dan quality gate.

## Contracts antar layer
### Domain
- fokus pada model inti dan invariants
- tidak boleh import platform, transport/http, atau vendor SDK
- third-party package di domain dilarang

### Ports
- berisi interface
- boleh import stdlib dan domain modul sendiri
- third-party package diminimalkan

### Usecase
- mengorkestrasi flow bisnis
- boleh import domain + ports
- tidak boleh import platform atau transport/http

### Module transport/http
- hanya mapping request/response dan panggil usecase
- tidak boleh import platform

### Platform
- implementasi nyata untuk DB, provider, token, state, crypto, KMS
- tidak boleh import transport/http

### Bootstrap
- wiring dependency
- boleh import config, platform, usecase, transport/http
- tidak boleh memindahkan aturan bisnis ke bootstrap

## File hygiene
- target file Go <= 100 baris
- exception boleh bila file adalah glue/contract penting
- file yang melewati target harus punya alasan singkat yang jelas di komentar kerja atau docs terkait

## Protected contracts `pos-go`
Saat ini kontrak yang harus dijaga:
- `cmd/api/main.go` sebagai entry runtime API
- `internal/config/*` sebagai sumber config aktif
- `internal/app/bootstrap/*` sebagai wiring runtime
- kontrak ports di `internal/modules/*/ports/*`

Kontrak HTTP publik akan ditambahkan setelah route auth dan health benar-benar dikunci.

## Aturan perubahan
Jika boundary, layout, atau kontrak modul berubah:
- update dokumen ini
- update blueprint bila scope ikut berubah
- buat ADR bila perubahan menyentuh keputusan arsitektur utama
