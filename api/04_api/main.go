package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// --------------------
// Struct Profile
// --------------------
type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// --------------------
// Handler POST /profile
// --------------------
func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Pastikan method POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form (untuk akses FormValue dari x-www-form-urlencoded)
	r.ParseForm()
	nameForm := r.FormValue("name")
	ageForm := r.FormValue("age")

	fmt.Println("=== FORM VALUE ===")
	spew.Dump(map[string]string{"name": nameForm, "age": ageForm})

	// Parse JSON body
	var p Profile
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	fmt.Println("=== JSON BODY ===")
	spew.Dump(p) // Dump struct ke console

	// Response ke client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp := map[string]interface{}{
		"message": "Profile created",
		"profile": p,
	}
	json.NewEncoder(w).Encode(resp)
}

// --------------------
// Handler GET /
// --------------------
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Server is running.")
}

// --------------------
// Main Server
// --------------------
func main() {
	// Buat ServeMux untuk multiple endpoints
	mux := http.NewServeMux()

	// Endpoint GET /
	mux.HandleFunc("/", helloHandler)

	// Endpoint POST /profile
	mux.HandleFunc("/profile", profileHandler)

	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)

	// Jalankan server
	log.Fatal(http.ListenAndServe(port, mux))
}
