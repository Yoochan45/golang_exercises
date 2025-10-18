package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("=== Functions in Go ===\n")

	// 1. Basic function call
	fmt.Println("1. Basic Function:")
	sayHello()
	greet("Alice")
	fmt.Println()

	// 2. Function with return value
	fmt.Println("2. Function with Return:")
	addResult := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", addResult)
	product := multiply(4, 7)
	fmt.Printf("4 × 7 = %d\n\n", product)

	// 3. Multiple return values
	fmt.Println("3. Multiple Return Values:")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("10 ÷ 0 = Error: %s\n\n", err)
	}

	// 4. Named return values
	fmt.Println("4. Named Returns:")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle (5×3): Area=%.1f, Perimeter=%.1f\n\n", area, perimeter)

	// 5. Variadic functions
	fmt.Println("5. Variadic Functions:")
	fmt.Printf("Sum of 1,2,3: %d\n", sum(1, 2, 3))
	fmt.Printf("Sum of 1,2,3,4,5: %d\n", sum(1, 2, 3, 4, 5))

	// Passing slice to variadic function
	numbers := []int{10, 20, 30}
	fmt.Printf("Sum of slice: %d\n\n", sum(numbers...))

	// 6. Anonymous functions
	fmt.Println("6. Anonymous Functions:")
	func() {
		fmt.Println("I'm an anonymous function!")
	}()

	// Function as variable
	double := func(x int) int {
		return x * 2
	}
	fmt.Printf("Double of 5: %d\n\n", double(5))

	// 7. Closures
	fmt.Println("7. Closures:")
	counter := makeCounter()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n\n", counter())

	// 8. Higher-order functions
	fmt.Println("8. Higher-Order Functions:")
	nums := []int{1, 2, 3, 4, 5}
	squared := apply(nums, func(n int) int {
		return n * n
	})
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("Squared: %v\n\n", squared)

	// 9. Defer statement
	fmt.Println("9. Defer Statement:")
	demonstrateDefer()
	fmt.Println()

	// 10. Practical example - calculator
	fmt.Println("10. Practical Example - Calculator:")
	fmt.Printf("Addition: %.2f\n", calculator(10, 5, "+"))
	fmt.Printf("Subtraction: %.2f\n", calculator(10, 5, "-"))
	fmt.Printf("Multiplication: %.2f\n", calculator(10, 5, "*"))
	fmt.Printf("Division: %.2f\n", calculator(10, 5, "/"))
}

// Basic functions
func sayHello() {
	fmt.Println("Hello, World!")
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a, b int) int {
	return a + b
}

// Shorthand parameter declaration
func multiply(a, b int) int {
	return a * b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Closure - function yang mengembalikan function
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Higher-order function
func apply(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = fn(n)
	}
	return result
}

// Defer demonstration
func demonstrateDefer() {
	defer fmt.Println("Third (deferred)")
	defer fmt.Println("Second (deferred)")
	defer fmt.Println("First (deferred)")
	fmt.Println("Normal execution")
}

// Practical calculator function
func calculator(a, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		}
		return 0
	default:
		return 0
	}
}
