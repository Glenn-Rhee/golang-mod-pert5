package main

import (
	"log"
	"net/http"
	"strings"

	"pert-5/handlers"
	"pert-5/middlewares"
)

func main() {

	mux := http.NewServeMux()

	// Soal No 7A. Implementasikan endpoint /api/tasks untuk method GET dan POST.
	mux.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method == http.MethodGet {
			handlers.GetAll(w, r)
		} else if r.Method == http.MethodPost {
			handlers.Create(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// Soal No 7B. Implementasikan endpoint /api/tasks/{id} untuk method PUT dan DELETE.
	mux.HandleFunc("/api/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/api/tasks/") {
			switch r.Method {
			case http.MethodPut:
				handlers.Update(w, r)
			case http.MethodDelete:
				handlers.Delete(w, r)
			default:
				http.NotFound(w, r)
			}
		}
	})

	mux.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", middlewares.Logger(mux)))
}
