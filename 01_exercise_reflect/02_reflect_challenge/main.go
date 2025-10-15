package main

// Challenge: Reflect Challenge
// Sebagai seorang Back-end Developer, kamu ditugaskan untuk membantu project terbaru Avengers Corp, yaitu membuat program integrasi dan pusat data kejahatan untuk para anggota Avengers.
// Kamu diminta untuk membuat sebuah fungsi/method yang dapat digunakan secara global, yang berfungsi untuk melakukan validasi terhadap sebuah object struct berdasarkan tag nya.
// Tag yang perlu di-handle adalah sebagai berikut
// required (field harus memiliki value, mampu menghandle field dengan type string dan int)
// max & min (field dengan tipe data integer harus memiliki value sesuai dengan min / max yang ditentukan)
// maxLen & minLen (field dengan tipe data string harus memiliki length sesuai dengan minLen / maxLen yang ditentukan)
// email (field harus memiliki value dengan format yang sesuai email, bisa gunakan explorasi regex pada golang https://gobyexample.com/regular-expressions)

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

// Struct untuk validasi
type User struct {
	Name  string `validate:"required,minLen=3,maxLen=50"`
	Age   int    `validate:"required,min=1,max=100"`
	Email string `validate:"required,email"`
}

// Fungsi untuk melakukan validasi berdasarkan tag
func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("validate")

		if tag == "" {
			continue
		}

		rules := parseTag(tag)
		for rule, param := range rules {
			switch rule {
			case "required":
				if isEmpty(value) {
					return fmt.Errorf("%s is required", field.Name)
				}
			case "min":
				if value.Kind() == reflect.Int && value.Int() < toInt(param) {
					return fmt.Errorf("%s must be at least %s", field.Name, param)
				}
			case "max":
				if value.Kind() == reflect.Int && value.Int() > toInt(param) {
					return fmt.Errorf("%s must be at most %s", field.Name, param)
				}
			case "minLen":
				if value.Kind() == reflect.String && len(value.String()) < int(toInt(param)) {
					return fmt.Errorf("%s length must be at least %s", field.Name, param)
				}
			case "maxLen":
				if value.Kind() == reflect.String && len(value.String()) > int(toInt(param)) {
					return fmt.Errorf("%s length must be at most %s", field.Name, param)
				}
			case "email":
				if value.Kind() == reflect.String &&
					!isValidEmail(value.String()) {
					return fmt.Errorf("%s must be a valid email", field.Name)
				}
			}
		}
	}

	return nil
}

// Fungsi untuk memeriksa apakah nilai kosong
func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int:
		return v.Int() == 0
	}
	return false
}

// Fungsi untuk mengonversi string ke int
func toInt(s string) int64 {
	var i int64
	fmt.Sscanf(s, "%d", &i)
	return i
}

// Fungsi untuk memeriksa format email menggunakan regex
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Fungsi untuk memparsing tag menjadi map
func parseTag(tag string) map[string]string {
	rules := make(map[string]string)
	parts := regexp.MustCompile(`,`).Split(tag, -1)
	for _, part := range parts {
		if kv := regexp.MustCompile(`=`).Split(part, 2); len(kv) == 2 {
			rules[kv[0]] = kv[1]
		} else {
			rules[kv[0]] = ""
		}
	}
	return rules
}

// Fungsi utama
func main() {
	var user User
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input name: ")
	name, _ := reader.ReadString('\n')
	user.Name = strings.TrimSpace(name)

	fmt.Print("Input email:")
	email, _ := reader.ReadString('\n')
	user.Email = strings.TrimSpace(email)

	fmt.Print("Input age: ")
	fmt.Scan(&user.Age)
	if err := ValidateStruct(user); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation successful")
	}
}
