# 📝 Article Service API

Sebuah proyek API service yang dibangun menggunakan **Golang**, **Redis**, dan **PostgreSQL** untuk mengelola artikel pribadi yang ditulis dalam format **Markdown**. Semua data artikel disimpan secara efisien dan responsif, memanfaatkan **Redis** sebagai caching layer utama.

---

## ⚙️ Teknologi yang Digunakan

- **Golang** → bahasa utama untuk API service
- **PostgreSQL** → penyimpanan data utama untuk artikel dan tag
- **Redis** → caching service untuk mempercepat akses data
- **Markdown** → format penulisan artikel
- **UUID (google/uuid)** → sebagai primary key pada semua entitas

---

## 🧩 Arsitektur & Rencana Jangka Panjang

Proyek ini dirancang sebagai bagian dari **arsitektur microservices**. Di masa depan, `Article Service` akan menjadi satu dari beberapa service mandiri yang saling berkomunikasi melalui API atau message queue. Service lain mungkin mencakup:

- **User Service**
- **Auth Service (SSO/API Key)**
- **Comment Service**
- **Analytics/Logging Service**

Setiap service akan dikembangkan secara independen, memungkinkan scalability dan maintainability yang lebih baik.

---

## 📦 Fitur Utama

- Menyimpan artikel (judul, slug, deskripsi, isi markdown, dan tag).
- Mendukung struktur tag dengan slug unik.
- Menggunakan Redis sebagai cache layer utama.
- Penulisan artikel fleksibel dengan format Markdown.
- Autentikasi ringan menggunakan API Key (tanpa login).

---

## 🔁 Flow Operasi `GET /article/:slug`

```
Client Request
     ↓
Cek Redis (key: article:{slug})
     ↓
[Jika ditemukan] → return response dari Redis
     ↓
[Jika tidak ditemukan]
→ Ambil data dari PostgreSQL
→ Simpan hasil ke Redis (TTL default: 10 menit)
→ Return ke client
```

---

## 🗃️ Struktur Data

### `Article`

```go
type Article struct {
	ID          uuid.UUID
	Title       string
	Slug        string
	Description string
	Body        string // markdown
	Tags        []Tag
	CreatedAt   string
	UpdatedAt   string
}
```

### `Tag`

```go
type Tag struct {
	ID        uuid.UUID
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
```

---

## 🚀 Menjalankan dengan Docker Compose

> Pastikan Redis & PostgreSQL tersedia secara lokal atau via Docker.

```bash
docker-compose up --build
```

---

## 🛠️ Rencana Pengembangan (TODO)

- [ ] Endpoint `POST /articles` → tambah artikel baru
- [ ] Endpoint `PUT /articles/:slug` → edit artikel
- [ ] Endpoint `DELETE /articles/:slug` → hapus artikel & invalidate Redis
- [ ] Pencarian artikel berdasarkan tag
- [ ] Integrasi dengan Markdown renderer (untuk preview HTML)
- [ ] Logging & tracing (opentelemetry, loki, grafana)
- [ ] Pemisahan ke dalam service-service independen (microservices architecture)

---

## 🧠 Catatan

- API ini bersifat **pribadi**, ditujukan untuk konsumsi penulis sendiri.
- Untuk saat ini tidak ada sistem login, cukup autentikasi menggunakan **API Key**.
- Semua artikel akan diformat Markdown dan bisa dikonversi ke HTML saat frontend rendering.

---
