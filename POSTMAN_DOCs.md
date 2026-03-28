## 🧪 Panduan Testing (Postman)

### Setup Awal

1. Buat **Collection** baru → beri nama `Task API`
2. Tambahkan **Collection Variable**:

| Variable   | Value                   |
| ---------- | ----------------------- |
| `base_url` | `http://localhost:8080` |
| `task_id`  | _(diisi otomatis)_      |

---

### Test Case 1 — GET All Tasks

| Field  | Value                    |
| ------ | ------------------------ |
| Method | `GET`                    |
| URL    | `{{base_url}}/api/tasks` |

**Expected:** `200 OK`, `data` berupa array

```javascript
pm.test('Status code is 200', () => {
  pm.response.to.have.status(200);
});
pm.test('Response status is success', () => {
  pm.expect(pm.response.json().status).to.eql('success');
});
pm.test('Data is an array', () => {
  pm.expect(pm.response.json().data).to.be.an('array');
});
pm.test('Response time < 500ms', () => {
  pm.expect(pm.response.responseTime).to.be.below(500);
});
```

---

### Test Case 2 — POST Create Task (Berhasil)

| Field   | Value                            |
| ------- | -------------------------------- |
| Method  | `POST`                           |
| URL     | `{{base_url}}/api/tasks`         |
| Headers | `Content-Type: application/json` |

**Body:**

```json
{
  "title": "Belajar Mandarin",
  "description": "Belajar mandarin dari a - z",
  "status": "pending"
}
```

**Expected:** `201 Created`

```javascript
pm.test('Status code is 201', () => {
  pm.response.to.have.status(201);
});
pm.test('Task created with correct title', () => {
  const requestBody = JSON.parse(pm.request.body.raw);
  pm.expect(pm.response.json().data.title).to.eql(requestBody.title);
});
pm.test('Task has an ID', () => {
  pm.expect(pm.response.json().data.id).to.be.a('number');
});
pm.test('Task has created_at', () => {
  pm.expect(pm.response.json().data.created_at).to.not.be.empty;
});

// Simpan ID untuk test berikutnya
pm.collectionVariables.set('task_id', pm.response.json().data.id);
```

---

### Test Case 3 — POST Create Task (Status Default)

| Field  | Value                    |
| ------ | ------------------------ |
| Method | `POST`                   |
| URL    | `{{base_url}}/api/tasks` |

**Body:** _(tanpa field `status`)_

```json
{
  "title": "Task Tanpa Status",
  "description": "Cek apakah default status pending"
}
```

**Expected:** `201 Created`, `status` = `pending`

```javascript
pm.test('Status code is 201', () => {
  pm.response.to.have.status(201);
});
pm.test('Default status is pending', () => {
  pm.expect(pm.response.json().data.status).to.eql('pending');
});
```

---

### Test Case 4 — POST Create Task (Title Kosong)

**Body:**

```json
{
  "title": "",
  "description": "Tidak ada title"
}
```

**Expected:** `400 Bad Request`

```javascript
pm.test('Status code is 400', () => {
  pm.response.to.have.status(400);
});
pm.test('Error status returned', () => {
  pm.expect(pm.response.json().status).to.eql('error');
});
pm.test('Error message mentions title', () => {
  pm.expect(pm.response.json().message).to.include('title');
});
```

---

### Test Case 5 — POST Create Task (Status Tidak Valid)

**Body:**

```json
{
  "title": "Task Invalid Status",
  "status": "selesai"
}
```

**Expected:** `400 Bad Request`

```javascript
pm.test('Status code is 400', () => {
  pm.response.to.have.status(400);
});
pm.test('Error message mentions valid statuses', () => {
  pm.expect(pm.response.json().message).to.include('pending');
});
```

---

### Test Case 6 — PUT Update Task (Berhasil)

| Field   | Value                                |
| ------- | ------------------------------------ |
| Method  | `PUT`                                |
| URL     | `{{base_url}}/api/tasks/{{task_id}}` |
| Headers | `Content-Type: application/json`     |

**Body:**

```json
{
  "title": "Belajar Mandarin v2",
  "description": "Belajar Mandarin Intermediate",
  "status": "in-progress"
}
```

**Expected:** `200 OK`, data terupdate

```javascript
pm.test('Status code is 200', () => {
  pm.response.to.have.status(200);
});
pm.test('Title updated correctly', () => {
  pm.expect(pm.response.json().data.title).to.eql('Belajar Mandarin v2');
});
pm.test('Status updated correctly', () => {
  pm.expect(pm.response.json().data.status).to.eql('in-progress');
});
pm.test('created_at tidak berubah', () => {
  pm.expect(pm.response.json().data.created_at).to.not.be.empty;
});
```

---

### Test Case 7 — PUT Update Task (ID Tidak Ditemukan)

| Field  | Value                        |
| ------ | ---------------------------- |
| Method | `PUT`                        |
| URL    | `{{base_url}}/api/tasks/999` |

**Body:**

```json
{
  "title": "task tidak ditemukan",
  "status": "done"
}
```

**Expected:** `404 Not Found`

```javascript
pm.test('Status code is 404', () => {
  pm.response.to.have.status(404);
});
pm.test('Error status returned', () => {
  pm.expect(pm.response.json().status).to.eql('error');
});
```

---

### Test Case 8 — PUT Update Task (ID Tidak Valid)

| Field  | Value                        |
| ------ | ---------------------------- |
| Method | `PUT`                        |
| URL    | `{{base_url}}/api/tasks/abc` |

**Expected:** `400 Bad Request`

```javascript
pm.test('Status code is 400', () => {
  pm.response.to.have.status(400);
});
pm.test('Error mentions invalid ID', () => {
  pm.expect(pm.response.json().message).to.include('ID tidak valid');
});
```

---

### Test Case 9 — DELETE Task (Berhasil)

| Field  | Value                                |
| ------ | ------------------------------------ |
| Method | `DELETE`                             |
| URL    | `{{base_url}}/api/tasks/{{task_id}}` |

**Expected:** `200 OK`, data task yang dihapus dikembalikan

```javascript
pm.test('Status code is 200', () => {
  pm.response.to.have.status(200);
});
pm.test('Deleted task data returned', () => {
  pm.expect(pm.response.json().data).to.not.be.null;
  pm.expect(pm.response.json().data.id).to.be.a('number');
});
pm.test('Message confirms deletion', () => {
  pm.expect(pm.response.json().message).to.include('dihapus');
});
```

---

### Test Case 10 — DELETE Task (ID Tidak Ditemukan)

| Field  | Value                        |
| ------ | ---------------------------- |
| Method | `DELETE`                     |
| URL    | `{{base_url}}/api/tasks/999` |

**Expected:** `404 Not Found`

```javascript
pm.test('Status code is 404', () => {
  pm.response.to.have.status(404);
});
pm.test('Error status returned', () => {
  pm.expect(pm.response.json().status).to.eql('error');
});
```

---

### Ringkasan Test Case

| #   | Endpoint                | Skenario               | Expected |
| --- | ----------------------- | ---------------------- | -------- |
| 1   | GET `/api/tasks`        | Ambil semua task       | `200`    |
| 2   | POST `/api/tasks`       | Create berhasil        | `201`    |
| 3   | POST `/api/tasks`       | Status default pending | `201`    |
| 4   | POST `/api/tasks`       | Title kosong           | `400`    |
| 5   | POST `/api/tasks`       | Status tidak valid     | `400`    |
| 6   | PUT `/api/tasks/:id`    | Update berhasil        | `200`    |
| 7   | PUT `/api/tasks/999`    | ID tidak ditemukan     | `404`    |
| 8   | PUT `/api/tasks/abc`    | ID tidak valid         | `400`    |
| 9   | DELETE `/api/tasks/:id` | Delete berhasil        | `200`    |
| 10  | DELETE `/api/tasks/999` | ID tidak ditemukan     | `404`    |

> **Tips:** Jalankan semua test sekaligus menggunakan **Collection Runner**.
> Pastikan urutan test **dari atas ke bawah** karena `{{task_id}}` diisi otomatis oleh Test Case 2.
