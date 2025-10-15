package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// --------------------
// Struct untuk JSON
// --------------------
type Message struct {
	Message string `json:"message"`
}

type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// --------------------
// Handler untuk root "/"
// --------------------
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// set header JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// kirim response JSON
	json.NewEncoder(w).Encode(Message{Message: "Server is running!"})
}

// --------------------
// Handler GET /hello?name=Yoo
// --------------------
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name") // ambil query param "name"
	if name == "" {
		name = "Guest"
	}

	json.NewEncoder(w).Encode(Message{Message: fmt.Sprintf("Hello %s!", name)})
}

// --------------------
// Handler POST /profile
// --------------------
func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// decode body JSON
	var p Profile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Invalid JSON body"})
		return
	}

	// validasi data
	if p.Name == "" || p.Age <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Name or Age is invalid"})
		return
	}

	// sukses
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// --------------------
// Handler GET /age/:age (dynamic path example)
// --------------------
func ageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ageStr := r.URL.Path[len("/age/"):] // ambil sisa path sebagai age
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Invalid age"})
		return
	}

	json.NewEncoder(w).Encode(Message{Message: fmt.Sprintf("Your age is %d", age)})
}

// --------------------
// MAIN: setup server
// --------------------
func main() {
	// buat multiplexer
	mux := http.NewServeMux()

	// daftar endpoint
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/profile", profileHandler)
	mux.HandleFunc("/age/", ageHandler) // contoh endpoint dinamis /age/25

	// jalankan server di port 8080
	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
