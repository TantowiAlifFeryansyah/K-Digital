# 🖼️ Image Cropper via Border Detection (Go)

Proyek ini adalah implementasi take-home test untuk **memotong gambar secara otomatis berdasarkan border hitam**.
Hasil akhirnya hanya menyisakan area **dalam border + garis border hitam**, sedangkan area di luar border akan terhapus.

---

## ⚙️ Cara Menjalankan

1. Clone repository
    - git clone git@github.com:TantowiAlifFeryansyah/K-Digital.git
    - cd image-cropper
2. Inisialisasi Go modules
    - go mod tidy
3. Jalankan program
    - go run main.go

---

## 🎯 Tujuan

- Input: `image.png` (diletakkan di folder `uploads/`)
- Output: `output.png` (hasil potongan gambar, hanya bagian dalam border + border hitam)
- Metode: Deteksi border dengan cara membaca piksel hitam (`R=0, G=0, B=0`).

---

## 📂 Struktur Folder

```bash
.
├── cmd/
│   └── main.go
├── constants/
│   └── response.go
├── entity/
│   └── border.go
├── handler/
│   └── crop_handler.go
├── repository/
│   └── crop_repository.go
├── service/
│   └── crop_service.go
├── uploads/
│   └── image.png   # gambar input
├── go.mod
├── go.sum
└── README.md
