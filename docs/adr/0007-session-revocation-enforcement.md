# ADR 0007: Session Revocation Enforcement

## Status
accepted

## Context
Setelah access token verification, principal resolution, dan protected route runtime proof selesai, proyek membutuhkan strategi yang jelas untuk menangani session revoke dan logout.

Tanpa enforcement tambahan, access token yang sudah diterbitkan akan tetap valid sampai expiry walaupun session sudah dianggap logout atau revoked.

Untuk POS yang membutuhkan kontrol akses yang lebih ketat dan audit yang lebih jujur, perilaku ini perlu diputuskan secara eksplisit.

## Decision
Dipilih strategi DB-backed session revocation enforcement pada request path yang dilindungi.

Flow minimum:
1. access token diverifikasi
2. session_id diambil dari token claim
3. persistence layer memeriksa session masih aktif
4. jika session tidak ada atau revoked, request ditolak
5. jika session aktif, principal dilanjutkan ke permission enforcement

## Rule
- session dianggap aktif jika row session ada dan `revoked_at IS NULL`
- enforcement dilakukan pada middleware auth request
- logout dan revoke nanti akan mengubah status session di persistence layer
- access token tidak dianggap cukup hanya karena signature valid

## Options considered
### Opsi A - stateless access token until expiry
Kelebihan:
- sederhana
- tidak ada query tambahan pada request protected

Kekurangan:
- logout tidak langsung efektif
- revoke session tidak langsung memutus akses token aktif

### Opsi B - DB-backed session check
Kelebihan:
- logout/revoke benar-benar bermakna
- lebih cocok untuk POS dan audit
- sederhana untuk fase sekarang

Kekurangan:
- ada query tambahan pada request protected

### Opsi C - cache-assisted session check
Kelebihan:
- performa lebih baik pada skala besar

Kekurangan:
- kompleksitas naik
- belum perlu untuk fase sekarang

## Consequences
### Positif
- logout dan revoke menjadi efektif
- permission enforcement berjalan di atas session yang masih aktif
- audit auth lifecycle lebih jujur

### Negatif
- perlu port dan adapter session status checker
- request protected mendapat satu query tambahan

## Follow-up
- buat port session status checker minimum
- buat adapter PostgreSQL untuk cek active session
- integrasikan ke middleware auth
- tambah test untuk request dengan session revoked
