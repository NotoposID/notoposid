# Service Layer (`service.go`)

Dalam arsitektur **Clean Architecture** atau pola **Controller-Service-Repository**, file `service.go` berfungsi sebagai tempat untuk menyimpan **Business Logic (Aturan Bisnis)**.

Pemisahan tanggung jawab ini dilakukan agar setiap layer memiliki fokus yang jelas dan mudah untuk dipelihara.

## Pembagian Tanggung Jawab

### `handler.go`
Berfungsi untuk menangani urusan HTTP, seperti:

- Membaca request body
- Validasi parameter request
- Mengatur response JSON
- Mengatur status code HTTP

### `repository.go`
Berfungsi untuk menangani akses database, seperti:

- Query SQL / ORM
- Insert data
- Update data
- Delete data
- Get data

### `service.go`
Berfungsi untuk menangani aturan bisnis dan alur logika aplikasi.

# Contoh Use Case `service.go` pada Modul Products

## 1. Pengecekan Duplikasi SKU / Barcode

Sebelum memanggil `repository.Create()`, service akan melakukan pengecekan apakah:

- SKU sudah digunakan
- Barcode sudah digunakan

dalam tenant yang sama.

Jika sudah digunakan, maka service akan mengembalikan error:

```text
"SKU sudah terdaftar"
```

## 2. Auto Generate SKU

Jika user tidak menginputkan SKU, maka service dapat otomatis membuat SKU berdasarkan:

- Singkatan kategori
- Nomor urut produk

Contoh:

```text
BEV-0012
```

## 3. Validasi Harga (Pricing Logic)

Service dapat menerapkan validasi bisnis seperti:

- Harga jual (`Price`) tidak boleh lebih kecil dari harga modal (`CostPrice`)
- Jika harga jual lebih kecil, maka proses pembuatan produk ditolak

Contoh validasi:

```go
if req.Price < req.CostPrice {
    return errors.New("harga jual tidak boleh lebih kecil dari harga modal")
}
```

## 4. Integrasi dengan Modul Lain (Event Publishing)

Setelah produk berhasil dibuat, service dapat mengirim event/notifikasi ke sistem lain melalui:

- Kafka
- RabbitMQ
- NATS
- Message Queue lainnya

Tujuannya agar modul lain seperti:

- Analytics
- Inventory
- Stock Management

mengetahui bahwa ada produk baru yang harus diproses lebih lanjut.

# Kesimpulan

`service.go` adalah pusat dari aturan bisnis aplikasi.

Dengan memisahkan logic ke dalam service layer:

- Code menjadi lebih rapi
- Mudah di-maintain
- Mudah di-test (unit testing)
- Tidak mencampur logic HTTP dan Database
- Mempermudah pengembangan fitur ke depan


# Pemilihan GORM dengan SQL Raw

Pemilihan **GORM** (Object-Relational Mapping / ORM) dibandingkan **Raw SQL** pada proyek seperti ini didasarkan pada beberapa alasan utama yang berfokus pada produktivitas dan pemeliharaan kode (meskipun raw SQL secara teknis memang bisa lebih cepat):

### 1. Kecepatan Pengembangan & Produktivitas
Dengan GORM, operasi standar (CRUD) bisa dilakukan hanya dengan 1-2 baris kode tanpa harus menulis sintaks SQL yang panjang secara manual. 
- **Raw SQL:** `INSERT INTO users (tenant_id, name, email) VALUES ($1, $2, $3)`
- **GORM:** `r.db.Create(&user)`
Hal ini membuat penulisan fitur menjadi jauh lebih cepat.

### 2. Relasi & *Preloading* yang Mudah
Pada modul produk, sebuah `Product` memiliki relasi ke `Tenant` dan `Category`. Dengan GORM, kita cukup menuliskan relasi di *struct* dan memanggil `.Preload("Category")`. Jika menggunakan Raw SQL, kita harus melakukan `JOIN` secara manual dan melakukan *mapping* kolom satu per satu ke dalam struct *nested* di Go, yang rentan terjadi *error* pengetikan (*typo*).

### 3. Keamanan (*SQL Injection*)
GORM secara otomatis memproses seluruh parameter inputan sebagai *parameterized query* (mencegah *SQL Injection*). Walaupun di library raw SQL bawaan Go (`database/sql`) kita juga bisa melakukannya dengan parameter `$1, $2`, menggunakan GORM secara *default* sudah melindungi kita jika terjadi keteledoran.

### 4. Fitur *Soft Delete* & *Auto Migration*
- **Soft Delete**: Cukup menambahkan `gorm.DeletedAt`, maka setiap kali kita memanggil fungsi `Delete`, GORM tidak akan benar-benar menghapus datanya dari database, melainkan hanya mengisi kolom `deleted_at`. (Lihat modul `users`).
- **Auto Migration**: Kita tidak perlu pusing menulis *file migration* SQL secara manual di awal. GORM dapat membaca *struct* kita dan langsung membuat/merombak struktur tabel database menyesuaikan tipe data yang ada.

### 5. Kode yang Bersih (*Maintainability*)
Banyaknya query teks Raw SQL yang disebar dalam bahasa Go bisa membuat kode terlihat berantakan. Menggunakan konsep *struct* yang sangat natural dengan bahasa Go membuat arsitekturnya mudah dibaca, dikembangkan, dan dipahami oleh *programmer* Go lainnya.

---
**Kapan Sebaiknya Menggunakan Raw SQL?**
Meski GORM luar biasa untuk operasi standar, **Raw SQL** tetap digunakan apabila kita butuh performa tingkat tinggi (optimasi maksimal), menangani *query* laporan agregasi yang teramat rumit (*complex sub-queries*, *window functions*), atau menggunakan fitur spesifik bawaan database PostgreSQL yang tidak didukung secara instan oleh GORM. (GORM sendiri mendukung eksekusi Raw SQL jika memang sedang dibutuhkan: `db.Raw("SELECT...").Scan(&result)`).