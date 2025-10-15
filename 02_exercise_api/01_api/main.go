package main

// Import package yang dibutuhkan
import (
	"encoding/json" // untuk encode/decode JSON
	"net/http"      // untuk HTTP server & request/response

	"github.com/julienschmidt/httprouter" // router yang ringan & cepat
)

// --------------------
// Struct untuk request body
// --------------------
type Profile struct {
	Name string `json:"name"` // json:"name" → key JSON akan bernama "name"
	Age  int    `json:"age"`  // json:"age" → key JSON akan bernama "age"
}

// --------------------
// Struct untuk error response
// --------------------
type ErrorMessage struct {
	Message string `json:"message"` // key JSON akan bernama "message"
}

// --------------------
// Fungsi utama
// --------------------
func main() {
	// buat router baru
	router := httprouter.New()

	// --------------------
	// POST /profile
	// --------------------
	// handler function untuk endpoint POST /profile
	router.POST("/profile", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// w = ResponseWriter → untuk kirim response ke client
		// r = *http.Request → request dari client
		// p = httprouter.Params → params dari URL (misal /user/:id)

		// --------------------
		// Parse JSON dari request body
		// --------------------
		var t Profile                              // buat variabel untuk menampung data JSON
		decoder := json.NewDecoder(r.Body)         // ambil body request sebagai JSON decoder
		if err := decoder.Decode(&t); err != nil { // decode ke struct Profile
			// jika ada error decode, kembalikan error ke client
			w.Header().Set("Content-Type", "application/json")                     // set header response
			w.WriteHeader(http.StatusBadRequest)                                   // 400 Bad Request
			json.NewEncoder(w).Encode(ErrorMessage{Message: "Invalid Body input"}) // kirim JSON error
			return
		}

		// --------------------
		// Validasi data
		// --------------------
		if t.Name == "" || t.Age <= 0 { // cek jika Name kosong atau Age <=0
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorMessage{Message: "Invalid Body input"})
			return
		}

		// --------------------
		// Jika valid, kirim response sukses
		// --------------------
		w.Header().Set("Content-Type", "application/json") // header harus JSON
		w.WriteHeader(http.StatusCreated)                  // 201 Created
		// encode response ke JSON
		// pakai map[string]string karena kita cuma kirim pesan
		json.NewEncoder(w).Encode(map[string]string{"message": "Success create Profile"})
	})

	// jalankan HTTP server di port 8080
	// router digunakan untuk handle semua route
	http.ListenAndServe(":8080", router)
}
