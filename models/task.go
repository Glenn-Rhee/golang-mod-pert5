package models

import "time"

// Soal No 2. Buat sebuah struct bernama Task yang merepresentasikan data tugas.
// Struct ini memiliki field: ID, Title, Description, Status, dan CreatedAt.
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
