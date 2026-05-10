# NOTOPOS AI - Local Development Guide

Panduan ini akan membantu Anda menjalankan **NOTOPOS AI** di mesin lokal untuk keperluan pengembangan.

## 🛠️ Persiapan Sistem
Pastikan Anda sudah menginstal:
- **Node.js** (v20 atau lebih baru)
- **Go** (v1.21 atau lebih baru)
- **Docker & Docker Compose** (untuk Database & Redis)
- **pnpm** (`npm install -g pnpm`)

---

## 🚀 Langkah Menjalankan Project

### 1. Setup Environment
Salin file `.env.example` menjadi `.env` di masing-masing folder aplikasi:

```bash
# Backend
cp apps/api/.env.example apps/api/.env

# Frontend
cp apps/web/.env.example apps/web/.env
```

### 2. Jalankan Database & Redis (Docker)
Kita menggunakan Docker untuk menjalankan PostgreSQL (dengan `pgvector`) dan Redis agar tidak perlu menginstal manual di OS.

```bash
docker compose up -d db redis
```

### 3. Jalankan Project secara Bersamaan
Gunakan perintah `dev` di root folder untuk menjalankan Frontend (Next.js) dan Backend (Go) sekaligus menggunakan **Turborepo**.

```bash
pnpm dev
```

Aplikasi akan berjalan di:
- **Frontend**: [http://localhost:3000](http://localhost:3000)
- **Backend**: [http://localhost:8000](http://localhost:8000)
- **Health Check API**: [http://localhost:8000/health](http://localhost:8000/health)

---

## 📂 Struktur Penting (Local)
- `apps/api`: Source code backend Golang.
- `apps/web`: Source code frontend Next.js.
- `docker-compose.yml`: Konfigurasi database lokal.

---

## 🧪 Testing API
Anda bisa menggunakan Postman atau `curl` untuk mengecek apakah backend sudah terhubung ke database:
```bash
curl http://localhost:8000/health
```

Jika muncul `"status": "healthy"`, berarti sistem sudah siap!
