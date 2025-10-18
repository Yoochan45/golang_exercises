package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func profileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var p Profile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Message: "Invalid JSON"})
		return
	}

	if p.Name == "" || p.Age <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Message: "Invalid input data"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Profile created: %s (%d years old)", p.Name, p.Age),
	})
}

func panicHandler(w http.ResponseWriter, r *http.Request, v interface{}) {
	log.Println("Panic terjadi:", v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorMessage{Message: "Internal server error"})
}

func main() {
	router := httprouter.New()
	router.PanicHandler = panicHandler
	router.POST("/profile", profileHandler)

	mux := http.NewServeMux()
	mux.Handle("/api/", router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
