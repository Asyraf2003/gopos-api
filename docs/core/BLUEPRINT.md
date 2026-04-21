# BLUEPRINT

## Tujuan utama
Membangun fondasi API POS berbasis Go yang:
- raw Go tanpa starter kit
- memakai Echo sebagai adapter HTTP
- memakai PostgreSQL sebagai database utama
- memakai struktur hexagonal yang disiplin
- siap berkembang ke auth Google, JWT, dan keamanan data yang lebih kuat tanpa bongkar fondasi

## Scope fase aktif
Fase aktif saat ini hanya mencakup:
- fondasi runtime API
- konfigurasi aplikasi
- koneksi PostgreSQL
- health endpoint
- kerangka auth minimum
- governance docs

## Scope yang belum aktif
Belum aktif untuk fase ini:
- domain POS penuh
- workflow transaksi POS
- Google OAuth implementation penuh
- refresh session flow penuh
- audit sink lanjutan
- trust scoring / AAL
- encryption/KMS implementation nyata
- observability lanjutan
- deployment/hosting orchestration

## Prinsip desain
- API-first
- PostgreSQL-first
- hexagonal boundary wajib dijaga
- domain tidak bergantung ke transport, driver database, atau provider pihak ketiga
- Echo hanya adapter HTTP
- platform adapter menampung detail teknis eksternal
- auth harus tumbuh dari kontrak kecil yang jelas
- fitur keamanan lanjutan boleh ditambahkan bertahap tanpa merusak kontrak inti

## Boundary utama
- `cmd/` untuk entrypoint
- `internal/config` untuk config loading
- `internal/app/bootstrap` untuk wiring runtime minimum
- `internal/modules/*/domain` untuk entity dan aturan inti
- `internal/modules/*/ports` untuk kontrak
- `internal/modules/*/usecase` untuk orkestrasi bisnis
- `internal/modules/*/transport/http` untuk handler HTTP
- `internal/platform/*` untuk adapter teknis seperti postgres, google, jwt, state store, crypto
- `migrations/` untuk perubahan schema database
- `docs/` untuk pengendali keputusan dan proses

## Arah auth
Auth akan dibangun bertahap dengan komponen minimum:
- auth domain
- auth ports
- auth usecase
- auth HTTP transport
- google adapter
- jwt issuer
- state store
- postgres-backed persistence

## Arah keamanan data
Target jangka menengah:
- secret/config sensitif siap dipindah ke secret manager atau KMS-backed flow
- field sensitif dapat dienkripsi di adapter/platform layer
- domain tidak mengetahui detail KMS/encryption implementation

## Aturan perubahan
Blueprint hanya boleh diubah jika:
- ada bukti konflik teknis nyata
- ada scope baru yang resmi diaktifkan
- ada ADR accepted yang memaksa penyesuaian
