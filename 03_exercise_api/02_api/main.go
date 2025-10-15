package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// --------------------
// Struct data JSON
// --------------------
type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

// --------------------
// Handler Function
// --------------------
// Handler menerima request masuk dan menyiapkan response.
// Implementasi http.HandlerFunc menerima:
// - http.ResponseWriter -> untuk menulis response ke client
// - *http.Request -> data request dari client
func profileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set header supaya client tahu ini JSON
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON dari request body
	var p Profile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400 Bad Request
		json.NewEncoder(w).Encode(ErrorMessage{Message: "Invalid JSON"})
		return
	}

	// Validasi data
	if p.Name == "" || p.Age <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Message: "Invalid input data"})
		return
	}

	// Response sukses
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Profile created: %s (%d years old)", p.Name, p.Age),
	})
}

// --------------------
// Panic Handler
// --------------------
// Fungsi ini akan menangkap panic sehingga server tidak crash
func panicHandler(w http.ResponseWriter, r *http.Request, v interface{}) {
	log.Println("Panic terjadi:", v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorMessage{Message: "Internal Server Error"})
}

// --------------------
// Main function
// --------------------
func main() {
	// Buat router httprouter
	router := httprouter.New()

	// Set PanicHandler
	router.PanicHandler = panicHandler

	// Tambahkan endpoint POST /profile
	router.POST("/profile", profileHandler)

	// --------------------
	// Bisa juga pakai ServeMux bawaan net/http untuk multiple endpoint
	// --------------------
	mux := http.NewServeMux()
	mux.Handle("/api/", router) // semua route router akan di-mount ke /api/

	// --------------------
	// Buat struct Server
	// --------------------
	server := &http.Server{
		Addr:    ":8080", // port server
		Handler: mux,     // semua request akan diteruskan ke mux
	}

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
