package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"pert-5/handlers"
	"pert-5/middlewares"
)

const PORT = 3178

func main() {

	mux := http.NewServeMux()

	// Soal No 7A. Implementasikan endpoint /api/tasks untuk method GET dan POST.


	// Soal No 7B. Implementasikan endpoint /api/tasks/{id} untuk method PUT dan DELETE.


	mux.Handle("/", http.FileServer(http.Dir("./static")))

	// UBAH PORT 8080 MENGGUNAKAN 4 DIGIT NPM TERAKHIR
	log.Printf("Server running at http://localhost:%d\n", PORT)           // contoh-> http://localhost:2233
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), middlewares.Logger(mux))) // contoh-> :2233
}
