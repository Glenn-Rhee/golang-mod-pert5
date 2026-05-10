package handlers

// Soal No 3. Import package encoding/json, errors, net/http, pert-5/models, strconv, strings, sync, time

// In-memory storage
var (
	mu     sync.Mutex
	tasks  []models.Task
	nextID int
)

func init() {
	// Soal No 4. Inisialisasikan 2 data task awal yang berisi ID, Title, Description, Status, dan CreatedAt.
	tasks = []models.Task{
		
	}
	nextID = 4
}

// Soal No. 5. Inisialisasi validStatuses map[string]bool


// GET /api/tasks
// Soal No 6A. Implementasikan handler GetAll untuk mengambil semua task.


// POST /api/tasks
// Mengimplementasikan handler Create untuk membuat task baru.
func Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Format JSON tidak valid", nil)
		return
	}

	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "title wajib diisi", nil)
		return
	}
	if input.Status == "" {
		input.Status = "pending"
	}
	if !validStatuses[input.Status] {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "status tidak valid, gunakan: pending | in-progress | done", nil)
		return
	}

	mu.Lock()

	// Soal No 6B. Lengkapi kode berikut untuk membuat objek Task baru menggunakan struct models.Task.
	// Isi field ID dengan nextID, Title dengan input.Title, Description dengan input.Description,
	// Status dengan input.Status, dan CreatedAt dengan waktu saat ini.

}

// PUT /api/tasks/{id}
// Mengimplementasikan handler Update untuk memperbarui task.
func Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "ID tidak valid", nil)
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Format JSON tidak valid", nil)
		return
	}

	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	// Soal No 6C. Buatlah validasi untuk memastikan field Title tidak kosong.


	mu.Lock()
	defer mu.Unlock()

	updatedTask, err := func() (models.Task, error) {
		for i, t := range tasks {
			if t.ID == id {
				if input.Title != "" {
					tasks[i].Title = input.Title
				}
				if input.Description != "" {
					tasks[i].Description = input.Description
				}
				if input.Status != "" {
					tasks[i].Status = input.Status
				}
				return tasks[i], nil
			}
		}
		return models.Task{}, errors.New("task tidak ditemukan")
	}()

	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, "error", err.Error(), nil)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "success", "Task berhasil diupdate", updatedTask)
}

// DELETE /api/tasks/{id}
// Mengimplementasikan handler Delete untuk menghapus task.
func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/tasks/")
	id, err := strconv.Atoi(idStr)

	// Soal No 6D. Buatlah validasi untuk memastikan nilai ID yang diterima valid.


	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			utils.WriteJSON(w, http.StatusOK, "success", "Task berhasil dihapus", t)
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, "error", "task tidak ditemukan", nil)
}
