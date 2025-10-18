package main

import "fmt"

// Constants - dapat dideklarasikan di level package
const (
	AppName    = "Learning Go"
	AppVersion = "1.0.0"
	Pi         = 3.14159
)

func main() {
	fmt.Println("=== Variables and Constants ===")

	// 1. Variable declaration dengan tipe eksplisit
	var name string = "John Doe"
	var age int = 25
	var isStudent bool = true

	fmt.Println("1. Explicit type declaration:")
	fmt.Printf("Name: %s, Age: %d, Is Student: %t\n\n", name, age, isStudent)

	// 2. Variable declaration dengan type inference
	var city = "Jakarta"
	var temperature = 32.5

	fmt.Println("2. Type inference:")
	fmt.Printf("City: %s, Temperature: %.1f°C\n\n", city, temperature)

	// 3. Short declaration (paling sering digunakan)
	country := "Indonesia"
	population := 270_000_000 // underscore untuk readability

	fmt.Println("3. Short declaration:")
	fmt.Printf("Country: %s, Population: %d\n\n", country, population)

	// 4. Multiple variable declaration
	var x, y, z int = 1, 2, 3
	a, b, c := "A", "B", "C"

	fmt.Println("4. Multiple declarations:")
	fmt.Printf("Numbers: %d, %d, %d\n", x, y, z)
	fmt.Printf("Letters: %s, %s, %s\n\n", a, b, c)

	// 5. Zero values
	var emptyString string
	var emptyInt int
	var emptyBool bool

	fmt.Println("5. Zero values:")
	fmt.Printf("String: '%s', Int: %d, Bool: %t\n\n", emptyString, emptyInt, emptyBool)

	// 6. Constants
	fmt.Println("6. Constants:")
	fmt.Printf("App: %s v%s\n", AppName, AppVersion)
	radius := 5.0
	area := Pi * radius * radius
	fmt.Printf("Circle area (radius=%.1f): %.2f\n\n", radius, area)

	// 7. Reassigning variables
	name = "Jane Doe"
	age = 30
	fmt.Println("7. After reassignment:")
	fmt.Printf("Name: %s, Age: %d\n\n", name, age)

	// 8. Demonstrasi scope
	demonstrateScope()
}

func demonstrateScope() {
	// Variable di dalam function hanya accessible di dalam function
	localVar := "I'm local to this function"
	fmt.Println("8. Variable scope:")
	fmt.Println(localVar)
}
