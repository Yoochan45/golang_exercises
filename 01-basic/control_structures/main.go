package main

import "fmt"

func main() {
	fmt.Printf("=== Control Structures ===\n")

	// 1. Basic if-else
	fmt.Println("1. If-Else:")
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}
	fmt.Println()

	// 2. If-else if-else
	fmt.Println("2. If-Else If-Else:")
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}
	fmt.Println()

	// 3. If with initialization
	fmt.Println("3. If with Initialization:")
	if remainder := 10 % 2; remainder == 0 {
		fmt.Println("10 is even")
	} else {
		fmt.Println("10 is odd")
	}
	// remainder tidak accessible di sini
	fmt.Println()

	// 4. Switch - basic
	fmt.Println("4. Switch Statement:")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}
	fmt.Println()

	// 5. Switch with multiple cases
	fmt.Println("5. Switch with Multiple Cases:")
	month := "January"
	switch month {
	case "December", "January", "February":
		fmt.Println("Winter")
	case "March", "April", "May":
		fmt.Println("Spring")
	case "June", "July", "August":
		fmt.Println("Summer")
	default:
		fmt.Println("Autumn")
	}
	fmt.Println()

	// 6. Switch without condition (like if-else chain)
	fmt.Println("6. Switch without Condition:")
	temperature := 28
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature < 20:
		fmt.Println("Cold")
	case temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
	fmt.Println()

	// 7. For loop - classic
	fmt.Println("7. Classic For Loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// 8. For loop - while style
	fmt.Println("8. While-Style Loop:")
	count := 1
	for count <= 5 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Printf("\n")

	// 9. For loop with break
	fmt.Println("9. Loop with Break:")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break // keluar dari loop
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// 10. For loop with continue
	fmt.Println("10. Loop with Continue:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// 11. Nested loops
	fmt.Println("11. Nested Loops (Multiplication Table):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d×%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()

	// 12. FizzBuzz example
	fmt.Println("12. FizzBuzz (1-20):")
	fizzBuzz(20)
}

// FizzBuzz function
func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
