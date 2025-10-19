package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Structs in Go ===")

	// 1. Defining and creating structs
	fmt.Println("1. Basic Struct:")
	// using package-level Person type (defined below)
	// Note: avoid declaring a Person type inside main to allow package-level
	// functions and methods to accept the same Person type.

	person1 := Person{Name: "Alice", Age: 25}
	person2 := Person{"Bob", 30} // positional (not recommended)

	fmt.Printf("Person 1: %+v\n", person1)
	fmt.Printf("Person 2: %+v\n\n", person2)

	// 2. Accessing and modifying fields
	fmt.Println("2. Accessing Fields:")
	fmt.Printf("Name: %s\n", person1.Name)
	fmt.Printf("Age: %d\n", person1.Age)

	person1.Age = 26
	fmt.Printf("Updated age: %d\n\n", person1.Age)

	// 3. Zero value struct
	fmt.Println("3. Zero Value Struct:")
	var person3 Person
	fmt.Printf("Zero value: %+v\n\n", person3)

	// 4. Using new
	fmt.Println("4. Using new:")
	person4 := new(Person)
	person4.Name = "Charlie"
	person4.Age = 35
	fmt.Printf("New person: %+v\n\n", person4)

	// 5. Complex struct
	fmt.Println("5. Complex Struct:")
	type Employee struct {
		ID       int
		Name     string
		Email    string
		Salary   float64
		IsActive bool
	}

	emp := Employee{
		ID:       1001,
		Name:     "Alice Johnson",
		Email:    "alice@company.com",
		Salary:   75000.00,
		IsActive: true,
	}

	fmt.Printf("Employee: %+v\n\n", emp)

	// 6. Anonymous structs
	fmt.Println("6. Anonymous Struct:")
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Config: %+v\n\n", config)

	// 7. Nested structs
	fmt.Println("7. Nested Structs:")
	type Address struct {
		Street  string
		City    string
		ZipCode string
	}

	type PersonWithAddress struct {
		Name    string
		Age     int
		Address Address
	}

	personAddr := PersonWithAddress{
		Name: "Bob",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "10001",
		},
	}

	fmt.Printf("Person: %+v\n", personAddr)
	fmt.Printf("City: %s\n\n", personAddr.Address.City)

	// 8. Embedded structs (composition)
	fmt.Println("8. Embedded Structs:")
	type BasicInfo struct {
		Name string
		Age  int
	}

	type Student struct {
		BasicInfo // embedded
		StudentID string
		GPA       float64
	}

	student := Student{
		BasicInfo: BasicInfo{
			Name: "Alice",
			Age:  20,
		},
		StudentID: "S12345",
		GPA:       3.8,
	}

	fmt.Printf("Student: %+v\n", student)
	fmt.Printf("Name (promoted): %s\n\n", student.Name) // direct access

	// 9. Comparing structs
	fmt.Println("9. Comparing Structs:")
	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{Name: "Alice", Age: 25}
	p3 := Person{Name: "Bob", Age: 30}

	fmt.Printf("p1 == p2: %t\n", p1 == p2)
	fmt.Printf("p1 == p3: %t\n\n", p1 == p3)

	// 10. Slice of structs
	fmt.Println("10. Slice of Structs:")
	type Book struct {
		Title  string
		Author string
		Year   int
	}

	books := []Book{
		{Title: "1984", Author: "George Orwell", Year: 1949},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925},
	}

	fmt.Println("Books:")
	for i, book := range books {
		fmt.Printf("%d. %s by %s (%d)\n", i+1, book.Title, book.Author, book.Year)
	}
	fmt.Println()

	// 11. Map of structs
	fmt.Println("11. Map of Structs:")
	employees := map[string]Employee{
		"E001": {ID: 1001, Name: "Alice", Salary: 75000},
		"E002": {ID: 1002, Name: "Bob", Salary: 80000},
	}

	for id, emp := range employees {
		fmt.Printf("%s: %s ($%.2f)\n", id, emp.Name, emp.Salary)
	}
	fmt.Println()

	// 12. Struct with methods
	fmt.Println("12. Struct Methods:")

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n\n", rect.Perimeter())

	// 13. Struct tags (for JSON)
	fmt.Println("13. Struct Tags (JSON):")
	type User struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Age       int       `json:"age,omitempty"`
		CreatedAt time.Time `json:"created_at"`
		Password  string    `json:"-"` // never exported
	}

	user := User{
		Name:      "Alice",
		Email:     "alice@example.com",
		Age:       25,
		CreatedAt: time.Now(),
		Password:  "secret123",
	}

	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("JSON:\n%s\n\n", jsonData)

	// 14. Constructor pattern
	fmt.Println("14. Constructor Pattern:")
	newPerson := NewPerson("David", 28)
	fmt.Printf("New person: %+v\n\n", newPerson)

	// 15. Pointer to struct
	fmt.Println("15. Pointer to Struct:")
	p := &Person{Name: "Eve", Age: 22}
	fmt.Printf("Before: %+v\n", p)
	updateAge(p, 23)
	fmt.Printf("After: %+v\n", p)
}

// Method for Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Constructor function
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

type Person struct {
	Name string
	Age  int
}

// Function that modifies struct via pointer
func updateAge(p *Person, newAge int) {
	p.Age = newAge
}
