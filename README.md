# Task Management API

REST API sederhana untuk manajemen task, dibangun dengan **Go** dan terintegrasi dengan frontend berbasis **LitElement**. Semua logika bisnis dan akses data ditangani langsung di dalam layer handler.

---

## Struktur Project

```
pert-5-v2/
├── go.mod
├── main.go                ← entry point & routing mapping URL ke handler
├── middlewares/
│   └── logger.go          ← middleware logging request
├── handlers/
│   └── task_handler.go    ← semua logika (storage, validasi, CRUD)
├── models/
│   └── task.go            ← definisi struct data
├── utils/
│   └── response.go        ← utility untuk standardisasi JSON response
└── static/
    ├── index.html
    └── lit-all.min.js
```

---

## Arsitektur

Project ini menggunakan arsitektur sederhana — semua logika bisnis, validasi, dan akses data dihandle langsung di dalam **handler**:

```
Router (main.go) → Middleware → Handler → Model
```

| Layer           | File                       | Tanggung Jawab                                                       |
| --------------- | -------------------------- | -------------------------------------------------------------------- |
| **Model**       | `models/task.go`           | Definisi struct `Task`                                               |
| **Handler**     | `handlers/task_handler.go` | In-memory storage, validasi input, logika CRUD, format JSON response |
| **Middleware**  | `middlewares/logger.go`    | Intercept request untuk logging                                      |
| **Utils**       | `utils/response.go`        | Utility untuk standardisasi JSON response                            |
| **Main/Router** | `main.go`                  | Entry point — inisialisasi routing, middleware & menjalankan server  |

---

## Teknologi

- **Backend** : Go (net/http) — tanpa framework eksternal
- **Frontend** : LitElement (Web Component)
- **Storage** : In-memory (tidak menggunakan database)
- **Format** : JSON

---

## Cara Menjalankan

### Prasyarat

- Go versi 1.21 atau lebih baru

### Setup

```bash
# 1. Clone atau buat folder project
mkdir pert-5 && cd pert-5

# 2. Inisialisasi Go module
go mod init pert-5

# 3. Jalankan server
go run main.go
```

Server berjalan di: `http://localhost:8080`

---

## Model Data

```go
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
}
```

### Status yang tersedia

| Status        | Keterangan                 |
| ------------- | -------------------------- |
| `pending`     | Belum dikerjakan (default) |
| `in-progress` | Sedang dikerjakan          |
| `done`        | Selesai                    |

---

## API Endpoints

### Base URL

```
http://localhost:8080
```

### Format Response

Semua response menggunakan format standar:

```json
{
  "status": "success | error",
  "message": "Pesan response",
  "data": { ... },
  "timestamp": "2026-02-27 07:00:00"
}
```

---

### 1. Get All Tasks

```
GET /api/tasks
```

**Response `200 OK`:**

```json
{
  "status": "success",
  "message": "Data retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Belajar Go",
      "description": "Mempelajari dasar-dasar bahasa Go",
      "status": "pending",
      "created_at": "2026-02-27T07:00:00Z"
    }
  ],
  "timestamp": "2026-02-27 07:00:00"
}
```

---

### 2. Create Task

```
POST /api/tasks
Content-Type: application/json
```

**Request Body:**

```json
{
  "title": "Judul Task",
  "description": "Deskripsi task (opsional)",
  "status": "pending"
}
```

> `status` bersifat opsional, default: `pending`

**Response `201 Created`:**

```json
{
  "status": "success",
  "message": "Task berhasil dibuat",
  "data": {
    "id": 3,
    "title": "Judul Task",
    "description": "Deskripsi task",
    "status": "pending",
    "created_at": "2026-02-27T07:00:00Z"
  },
  "timestamp": "2026-02-27 07:00:00"
}
```

---

### 3. Update Task

```
PUT /api/tasks/{id}
Content-Type: application/json
```

**Request Body:**

```json
{
  "title": "Judul Baru",
  "description": "Deskripsi baru",
  "status": "in-progress"
}
```

> Semua field bersifat opsional (partial update). `created_at` tidak ikut berubah.

**Response `200 OK`:**

```json
{
  "status": "success",
  "message": "Task berhasil diupdate",
  "data": {
    "id": 3,
    "title": "Judul Baru",
    "description": "Deskripsi baru",
    "status": "in-progress",
    "created_at": "2026-02-27T07:00:00Z"
  },
  "timestamp": "2026-02-27 07:00:00"
}
```

---

### 4. Delete Task

```
DELETE /api/tasks/{id}
```

**Response `200 OK`:**

```json
{
  "status": "success",
  "message": "Task berhasil dihapus",
  "data": {
    "id": 3,
    "title": "Judul Task",
    "description": "Deskripsi task",
    "status": "in-progress",
    "created_at": "2026-02-27T07:00:00Z"
  },
  "timestamp": "2026-02-27 07:00:00"
}
```

---

### HTTP Status Code

| Code  | Keterangan                       |
| ----- | -------------------------------- |
| `200` | OK — request berhasil            |
| `201` | Created — data berhasil dibuat   |
| `400` | Bad Request — input tidak valid  |
| `404` | Not Found — data tidak ditemukan |
| `405` | Method Not Allowed               |

---

## Frontend

Antarmuka web tersedia di `http://localhost:8080` dengan fitur:

- **Dashboard** — statistik jumlah task per status
- **Tambah Task** — form inline untuk membuat task baru
- **Filter** — tampilkan task berdasarkan status
- **Edit Task** — modal popup untuk mengubah task
- **Hapus Task** — hapus task dengan notifikasi
- **Toast Notifikasi** — feedback sukses/error

### Menjalankan Frontend Offline (tanpa internet)

Download file Lit dan simpan ke folder `static/`:

```
https://cdn.jsdelivr.net/gh/lit/dist@3/all/lit-all.min.js
→ simpan sebagai static/lit-all.min.js
```

Lalu ubah import map di `static/index.html`:

```html
<script type="importmap">
  {
    "imports": {
      "lit": "/lit-all.min.js"
    }
  }
</script>
```

---

## Catatan

- Data **tidak persisten** — akan reset setiap kali server di-restart karena menggunakan in-memory storage
- Cocok digunakan untuk **pembelajaran** dan **demo**
- Untuk production, pindahkan logika akses data dari `handlers/task_handler.go` ke layer terpisah (misalnya repository dengan PostgreSQL atau MySQL)

---

## Dibuat untuk

**Pertemuan 5 — REST API Development & Frontend Integration (Lit UI)**  
Web Application Development with Go
