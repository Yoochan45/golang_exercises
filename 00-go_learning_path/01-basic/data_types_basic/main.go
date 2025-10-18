package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("=== Basic Data Types ===")

	// 1. Integer types
	fmt.Println("1. Integer Types:")
	var age int = 25
	var population int64 = 270_000_000
	var byteValue byte = 255 // byte = uint8 (0-255)
	var runeValue rune = '🔥' // rune = int32 (Unicode)

	fmt.Printf("Age: %d (type: %s)\n", age, reflect.TypeOf(age))
	fmt.Printf("Population: %d (type: %s)\n", population, reflect.TypeOf(population))
	fmt.Printf("Byte: %d (type: %s)\n", byteValue, reflect.TypeOf(byteValue))
	fmt.Printf("Rune: %c (value: %d, type: %s)\n\n", runeValue, runeValue, reflect.TypeOf(runeValue))

	// 2. Floating point types
	fmt.Println("2. Floating Point Types:")
	var price float64 = 99.99
	var temperature float32 = 36.5

	fmt.Printf("Price: %.2f (type: %s)\n", price, reflect.TypeOf(price))
	fmt.Printf("Temperature: %.1f°C (type: %s)\n\n", temperature, reflect.TypeOf(temperature))

	// 3. String type
	fmt.Println("3. String Type:")
	var name string = "John Doe"
	var message = "Hello, 世界! World" // Go mendukung Unicode

	fmt.Printf("Name: %s (type: %s)\n", name, reflect.TypeOf(name))
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("String length: %d bytes\n\n", len(message))

	// 4. Boolean type
	fmt.Println("4. Boolean Type:")
	var isStudent bool = true
	var hasLicense = false

	fmt.Printf("Is Student: %t (type: %s)\n", isStudent, reflect.TypeOf(isStudent))
	fmt.Printf("Has License: %t\n\n", hasLicense)

	// 5. Type conversion
	fmt.Println("5. Type Conversion:")
	var intValue int = 42
	var floatValue float64 = float64(intValue)
	var backToInt int = int(floatValue)

	fmt.Printf("Original int: %d\n", intValue)
	fmt.Printf("Converted to float64: %.2f\n", floatValue)
	fmt.Printf("Back to int: %d\n\n", backToInt)

	// 6. Integer division vs float division
	fmt.Println("6. Division Behavior:")
	intDivision := 5 / 2
	floatDivision := 5.0 / 2.0
	mixedDivision := float64(5) / float64(2)

	fmt.Printf("Integer division 5/2: %d\n", intDivision)
	fmt.Printf("Float division 5.0/2.0: %.1f\n", floatDivision)
	fmt.Printf("Mixed division (cast): %.1f\n\n", mixedDivision)

	// 7. Complex numbers
	fmt.Println("7. Complex Numbers:")
	var complexNum complex128 = complex(3, 4) // 3 + 4i

	fmt.Printf("Complex: %v\n", complexNum)
	fmt.Printf("Real: %.1f, Imaginary: %.1f\n\n", real(complexNum), imag(complexNum))

	// 8. Practical example: Calculate circle area
	fmt.Println("8. Practical Example - Circle Area:")
	radius := 5.0
	area := math.Pi * radius * radius
	circumference := 2 * math.Pi * radius

	fmt.Printf("Radius: %.1f\n", radius)
	fmt.Printf("Area: %.2f\n", area)
	fmt.Printf("Circumference: %.2f\n\n", circumference)

	// 9. Min and max values
	fmt.Println("9. Min/Max Values:")
	fmt.Printf("int8 min: %d, max: %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16 min: %d, max: %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32 min: %d, max: %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("uint8 max: %d\n", math.MaxUint8)
}
