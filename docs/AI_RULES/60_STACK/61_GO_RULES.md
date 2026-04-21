# P1 - Go Rules

## Tujuan
Menjaga hygiene implementasi Go tetap konsisten.

## Aturan
- Satu folder = satu package.
- Jaga ukuran file tetap terkontrol; bila melewati batas internal project, harus ada alasan jelas.
- Patuhi boundary dan import discipline.
- Jangan campur domain, transport, dan persistence tanpa jalur yang sah.
- `gofmt` wajib untuk file yang berubah.
- `go test ./...` wajib lulus untuk perubahan yang menyentuh kode Go.
- `go vet ./...` dijalankan bila perubahan sudah menyentuh fondasi runtime atau contract penting.
