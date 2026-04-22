# Handoff: Automation Stage 1

## Status
accepted

## Scope
Handoff ini menutup pekerjaan **otomasi repo tahap 1** untuk repo Go x PostgreSQL, dengan auth core diparkir sementara untuk dilanjutkan di chat baru.

Fokus yang sudah dikerjakan:
- make-based developer automation
- audit runner dasar
- security runner dasar
- database migration automation
- git push automation guard
- baseline output contract centralization awal
- repo police dasar untuk file size

Fokus yang sengaja **diparkir**:
- penyelesaian admin authorization API minimum
- runtime proof admin authz
- refactor semua file oversized sampai screening hijau penuh
- audit error policy
- audit time policy
- audit structure policy

## Keputusan yang sudah dikunci

### 1. Auth core live lane diparkir dalam kondisi kuat
Auth core live lane dianggap cukup kuat untuk diparkir sementara.

Yang sudah terbukti:
- login Google
- callback
- JWT issue / verify
- `/api/me`
- logout
- refresh rotation
- refresh token lama invalid

### 2. Screening dijadikan pintu utama disiplin repo
Screening sekarang menjadi jalur utama pengecekan repo.

Urutan screening saat ini:
1. `go test ./...`
2. `audit_ai_rules.sh`
3. `audit_file_size.sh`
4. `audit_security_gosec.sh`

### 3. File size policy berlaku keras
Rule yang dipakai:
- semua file code default maksimum 100 baris
- file non-kritikal yang lebih dari 100 baris wajib dipecah
- file kritikal hanya boleh lolos jika memenuhi dual control:
  1. path ada di registry exception
  2. file punya marker header `audit:allow-oversize`

### 4. Output contract API diarahkan ke folder presentation
Baseline awal sudah dibuat:
- `internal/presentation/http/id/system/me.go`

Artinya output API user-facing ke depan diarahkan terpusat, tidak dibiarkan liar di handler.

### 5. Jalur schema + reference auth seed yang canonical
Untuk repo ini, jalur canonical sekarang adalah:
- `make db-migrate`

Tidak dibuat `seed-auth` terpisah karena migration `.up.sql` yang ada sudah sekaligus memegang schema dan reference auth seed.

Akun Google live tetap **bukan seed**.

## Deliverables yang sudah jadi

### Makefile targets
Target yang sudah tersedia:
- `make help`
- `make test`
- `make audit-ai-rules`
- `make audit-file-size`
- `make security-gosec`
- `make audit-all`
- `make screening`
- `make build`
- `make run`
- `make auth-start`
- `make db-status`
- `make db-adopt-existing`
- `make db-migrate`
- `make git-status`
- `make push MSG="..."`

### Scripts
Script yang sudah tersedia:
- `scripts/audit_ai_rules.sh`
- `scripts/audit_file_size.sh`
- `scripts/audit_security_gosec.sh`
- `scripts/audit_all.sh`
- `scripts/db_status.sh`
- `scripts/db_migrate.sh`
- `scripts/db_adopt_existing.sh`
- `scripts/git_push.sh`

### Config / policy files
- `scripts/config/file_size_allowlist.txt`

### Presentation baseline
- `internal/presentation/http/id/system/me.go`

## Verification proof yang sudah ada

### Security
`make security-gosec` lulus dengan hasil:
- files scanned: 53
- issues: 0

### Screening
`make screening` sudah berjalan end-to-end.

Kondisi terakhir:
- `go test ./...` lulus
- `audit_ai_rules.sh` lulus
- `audit_security_gosec.sh` lulus
- gagal hanya di `audit_file_size.sh`

Artinya screening runner sudah valid, tapi repo belum lolos file-size discipline.

### Database automation
Sudah dibuktikan:
- `make db-status` bisa membaca `.env`
- `make db-adopt-existing` berhasil menandai migration lama sebagai applied
- `make db-migrate` sesudah adopt bersifat idempotent
- `make db-status` final menunjukkan semua migration `0001..0005` adalah `APPLIED`

### Push guard
`make push` tanpa `MSG` sudah menolak dengan benar:
- `[FAIL] MSG is required`

Artinya guard push minimum bekerja.

## Status operasional saat handoff

### Auth
- keseluruhan auth: **89%**
- auth core live: **100%**
- admin authz minimum: **74%**
- status: **parked**

### Audit / repo discipline
- keseluruhan audit/disiplin repo: **64%**
- AI rules audit: **100%**
- file-size audit engine: **100%**
- file-size compliance repo: **45%**
- output contract centralization: **35%**
- error policy audit: **0%**
- time policy audit: **0%**
- structure policy audit: **0%**
- aggregate audit runner: **85%**

### Security
- keseluruhan keamanan: **88%**
- auth runtime security: **90%**
- `gosec` automation: **100%**
- security baseline runner: **90%**

### Dev automation
- keseluruhan otomasi dev: **90%**
- DB automation: **100%**
- screening usable: **92%**

## Blocker aktif saat handoff

Satu blocker utama yang masih membuat `make screening` merah adalah **file-size compliance**.

File yang masih terdeteksi oversized:
- `internal/modules/auth/usecase/google_callback.go`
- `internal/platform/google/oidc_google.go`
- `internal/platform/postgres/account_role_remover_integration_test.go`
- `internal/platform/postgres/auth_account_identity_repository.go`
- `internal/platform/postgres/auth_account_identity_repository_integration_test.go`
- `internal/platform/postgres/principal_resolver.go`
- `internal/platform/postgres/principal_resolver_integration_test.go`
- `internal/platform/postgres/refresh_session_repository_integration_test.go`
- `internal/platform/token/jwt/issuer.go`
- `internal/platform/token/jwt/verifier_test.go`
- `internal/transport/http/middleware/authn_test.go`
- `internal/transport/http/middleware/authz_test.go`

## Rekomendasi kerja setelah handoff

### Rekomendasi utama
Prioritas paling bersih setelah handoff ini:
1. **chat baru untuk refactor file oversized sampai `make screening` hijau penuh**
2. baru setelah itu **chat baru untuk melanjutkan auth / admin authz minimum**
3. baru setelah screening dan auth cukup kuat, pindah ke API bisnis

### Alasan
- screening runner sudah jadi, jadi sekarang disiplin repo bisa dipaksakan dengan bukti
- kalau langsung pindah bisnis saat file-size audit masih merah, repo akan tumbuh di atas fondasi yang belum rapi
- sebagai solo dev, biaya menjaga repo kecil tetap waras jauh lebih murah daripada memperbaiki repo besar yang sudah keburu kusut

## Operating procedure yang direkomendasikan

Urutan kerja harian:
1. `make test`
2. `make security-gosec`
3. `make screening`
4. `make db-status`
5. `make db-migrate`
6. `make git-status`
7. `make push MSG="..."`

## Catatan penting
Chat ini dianggap selesai untuk **otomasi repo tahap 1**, tetapi **bukan** penutupan total proyek.

Penutupan total baru layak jika:
- auth + role minimum mencapai 100%
- audit pack mencapai 100%
- security runner + baseline mencapai 100%
- screening hijau penuh
- baru sesudah itu masuk handoff final dan pindah penuh ke domain bisnis
