package main

import (
	"fmt"
	"reflect"
)

// --------------------
// PENJELASAN UMUM
// --------------------
// Reflection itu cara program "mengintip" struktur dan informasi variabel saat runtime.
// Di Go, reflection bisa dipakai untuk:
// - Melihat tipe data variabel
// - Mengambil nilai variabel
// - Mengecek pointer atau field dalam struct
// Package bawaan Go untuk reflection adalah "reflect".

// --------------------
// MAIN
// --------------------
func main() {
	fmt.Println("===EXAMPLE ONE===")
	reflectOneBasic()
	fmt.Println("-----")
	fmt.Println("===EXAMPLE TWO===")
	reflectTwo()
	fmt.Println("-----")
	fmt.Println("===EXAMPLE THREE===")
	reflectThree()
	fmt.Println("===EXAMPLE FOUR===")
	refletFour()
}

// --------------------
// Fungsi 1: Basic Reflection
// --------------------
// - reflect.ValueOf() → ambil nilai (value) variabel dalam bentuk reflect.Value
// - reflect.TypeOf() → ambil tipe data (type) variabel dalam bentuk reflect.Type
// - Kind() → cek tipe dasar (int, string, struct, dll)
func reflectOneBasic() {
	var number float64 = 23.42
	var reflectValue = reflect.ValueOf(number) // ambil reflect.Value dari variabel

	fmt.Println("Type variabel: ", reflectValue.Type()) // tampilkan tipe data

	// cek tipe dasar, ambil nilai sesuai tipenya
	if reflectValue.Kind() == reflect.Float64 {
		fmt.Println("Nilai variabel: ", reflectValue.Float())
	}

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel: ", reflectValue.Int())
	}
}

// --------------------
// Fungsi 2: Reflection pada Struct
// --------------------
// Reflection pada struct:
//   - NumField() → ambil jumlah field/properti dalam struct
//   - Field(i int) → ambil informasi field tertentu (reflect.StructField) berdasarkan index
//     Hasil reflect.StructField bisa digunakan untuk melihat:
//   - Nama field
//   - Tipe data
//   - Tag (nanti di reflectThree)
type User struct {
	Email    string
	FullName string
}

func reflectTwo() {
	newUser := User{
		Email:    "Example@email.com",
		FullName: "Example Name",
	}

	userValue := reflect.ValueOf(newUser)
	fmt.Println("Value struct:", userValue) // tampilkan semua value

	userType := reflect.TypeOf(newUser)
	fmt.Println("Type struct:", userType) // tampilkan tipe struct

	fmt.Println("Jumlah field:", userType.NumField()) // jumlah field

	// loop semua field untuk lihat info masing-masing
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fmt.Println("Field index", i, ":", field.Name, "(", field.Type, ")")
	}
}

// --------------------
// Fungsi 3: Reflection dengan Tag
// --------------------
// Tags di struct bisa dipakai untuk rules/metadata, misal validasi.
// Kita bisa ambil value tag dengan Tag.Get("nama_tag")
type Users struct {
	Email    string `required:"true"`
	FullName string `required:"true" min:"5" max:"15"`
}

func reflectThree() {
	newUser := Users{
		Email:    "Example@mail.com",
		FullName: "Example Name",
	}

	userType := reflect.TypeOf(newUser)

	// ambil tag di field index 1 (FullName)
	fmt.Println(userType.Field(1).Tag.Get("required"))  // true
	fmt.Println(userType.Field(1).Tag.Get("min"))       // 5
	fmt.Println(userType.Field(1).Tag.Get("max"))       // 15
	fmt.Println(userType.Field(1).Tag.Get("notExists")) // kosong, tag tidak ada
}

func ValidateStructs(s interface{}) error {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(s).Field(i).Interface()
			if value == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}
	}
	return nil
}

func refletFour() {
	newUser := Users{
		Email:    "yaya@mail.com",
		FullName: "yoochan",
	}

	err := ValidateStructs(newUser)
	fmt.Println(err, "<--- error")
}
