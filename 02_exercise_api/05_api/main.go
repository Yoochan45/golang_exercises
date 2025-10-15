package main

import (
	"encoding/json"
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

func main() {
	router := httprouter.New()
	router.POST("/profile", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var t Profile
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&t); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorMessage{Message: "Invalid body input"})
			return
		}

		if t.Name == "" || t.Age <= 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorMessage{Message: "Invalid body input"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Success create profile"})
	})
	http.ListenAndServe(":8080", router)
}
