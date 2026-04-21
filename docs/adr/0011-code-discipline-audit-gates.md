# ADR 0011: Code Discipline Audit Gates

## Status
accepted

## Context
Proyek membutuhkan kontrol disiplin kode yang aktif, bukan hanya dokumentasi atau review manual.

Tujuan utamanya:
- mencegah spaghetti code yang hanya dibungkus struktur hexagonal
- menjaga file tetap kecil dan mudah dibaca
- menjaga pola error, waktu, dan struktur tetap seragam
- menyediakan polisi repo yang bisa dijalankan berulang saat jumlah file bertambah besar

Kondisi saat ini:
- sudah ada audit AI rules
- belum ada audit disiplin code structure
- belum ada kontrol default max line per file
- belum ada kontrol terpusat untuk pattern error, time, dan boundary

## Decision
Dipilih pendekatan modular audit gates berbasis script repository.

Audit gates minimum dibagi per kategori agar:
- kecil
- mudah diaudit
- mudah diperluas
- tidak mencampur banyak rule dalam satu script besar

Kategori audit minimum:
1. file size audit
2. error policy audit
3. time policy audit
4. structure policy audit
5. aggregate audit runner

## Rule Set

### 1. Default file size limit
- file code default maksimum 100 baris
- file yang masih dapat dipecah harus dipecah
- exception hanya boleh melalui allowlist eksplisit

### 1a. Test file policy
- file test tetap mengikuti limit default 100 baris
- exception untuk file test tidak boleh otomatis
- jika benar-benar perlu, file test masuk allowlist exception terpisah
- allowlist test harus diperlakukan sebagai hutang yang terlihat dan dapat dibersihkan bertahap

### 2. Allowed exceptions
Exception hanya untuk file yang secara alami berperan sebagai:
- bootstrap
- wiring
- routes
- provider
- config entrypoint
- file sensitif lain yang memang lebih aman tetap utuh

Exception bukan berarti bebas berantakan.
Exception tetap harus rapi dan auditable.

### 2a. Critical file exception marker
File yang dikecualikan dari screening ukuran wajib memenuhi dua syarat sekaligus:
1. path file tercatat di registry exception terpusat
2. file memiliki marker khusus di header file yang menjelaskan alasan exception

Tujuan dual control ini:
- exception terlihat dari file itu sendiri
- exception tetap bisa diaudit dari satu tempat terpusat
- exception tidak bisa lolos hanya karena salah satu sisi lupa diperiksa

### 3. Error policy
- error response HTTP harus seragam
- tidak boleh ada handler yang seenaknya menulis format sukses/error sendiri di luar contract yang disepakati
- mapping error bisnis harus konsisten dan dapat diaudit

### 4. Time policy
- penggunaan `time` harus konsisten
- tidak boleh tersebar liar tanpa pola yang jelas
- rule detail akan dikunci dan diaudit eksplisit agar perilaku waktu dapat diprediksi

### 5. Structure policy
- boundary antar layer harus konsisten
- handler tidak boleh mengambil tanggung jawab business rule yang bukan miliknya
- adapter tidak boleh diam-diam memegang policy bisnis
- folder/package harus tetap kecil dan jelas

## Options considered

### Opsi A - documentation only
Kelebihan:
- cepat

Kekurangan:
- tidak enforce
- tidak cukup untuk repo besar

### Opsi B - modular audit scripts
Kelebihan:
- enforceable
- kecil
- bisa berkembang bertahap
- cocok untuk repo Go dan bisa ditiru ke Laravel

Kekurangan:
- perlu maintain allowlist dan rules

### Opsi C - full custom analyzer from start
Kelebihan:
- sangat kuat

Kekurangan:
- terlalu besar untuk langkah awal
- tidak perlu untuk fase sekarang

## Consequences

### Positif
- repo punya polisi disiplin yang nyata
- potensi spaghetti dapat ditekan lebih awal
- pembacaan code tetap murah walaupun jumlah file bertambah

### Negatif
- perlu maintain rules dan allowlist
- perlu review berkala saat ada exception baru

## Follow-up
- buat file konfigurasi allowlist audit
- buat `scripts/audit_file_size.sh`
- buat `scripts/audit_error_policy.sh`
- buat `scripts/audit_time_policy.sh`
- buat `scripts/audit_structure_policy.sh`
- buat `scripts/audit_all.sh`
- integrasikan ke DoD / CI setelah stabil
