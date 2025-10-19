package main

import "fmt"

func main() {
	fmt.Println("=== Arrays in Go ===")

	// 1. Basic array declaration
	fmt.Println("1. Array Declaration:")
	var numbers [5]int
	fmt.Printf("Empty array: %v (length: %d)\n\n", numbers, len(numbers))

	// 2. Array initialization
	fmt.Println("2. Array Initialization:")
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	scores := [5]int{85, 90, 78, 92, 88}
	fmt.Printf("Fruits: %v\n", fruits)
	fmt.Printf("Scores: %v\n\n", scores)

	// 3. Let compiler count size
	fmt.Println("3. Auto-sized Array:")
	colors := [...]string{"Red", "Green", "Blue", "Yellow"}
	fmt.Printf("Colors: %v (length: %d)\n\n", colors, len(colors))

	// 4. Partial initialization
	fmt.Println("4. Partial Initialization:")
	partial := [5]int{1, 2, 3} // sisanya = 0
	fmt.Printf("Partial: %v\n\n", partial)

	// 5. Specific index initialization
	fmt.Println("5. Specific Index Init:")
	sparse := [5]int{0: 10, 2: 20, 4: 30}
	fmt.Printf("Sparse: %v\n\n", sparse)

	// 6. Accessing elements
	fmt.Println("6. Accessing Elements:")
	fmt.Printf("First fruit: %s\n", fruits[0])
	fmt.Printf("Last fruit: %s\n", fruits[len(fruits)-1])
	fruits[1] = "Mango"
	fmt.Printf("Modified fruits: %v\n\n", fruits)

	// 7. Iterating with for loop
	fmt.Println("7. For Loop Iteration:")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("Score %d: %d\n", i+1, scores[i])
	}
	fmt.Println()

	// 8. Iterating with range
	fmt.Println("8. Range Iteration:")
	for index, value := range scores {
		fmt.Printf("Index %d: %d\n", index, value)
	}
	fmt.Println()

	// 9. Range without index
	fmt.Println("9. Range (Value Only):")
	for _, color := range colors {
		fmt.Printf("- %s\n", color)
	}
	fmt.Println()

	// 10. Multi-dimensional array
	fmt.Println("10. Multi-dimensional Array:")
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	// 11. Array comparison
	fmt.Println("11. Array Comparison:")
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %t\n\n", arr1 == arr3)

	// 12. Array is value type (copy)
	fmt.Println("12. Array Copy Behavior:")
	original := [3]int{1, 2, 3}
	copied := original // creates a copy
	copied[0] = 99
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copied: %v\n\n", copied)

	// 13. Practical examples
	fmt.Println("13. Practical Examples:")

	// Calculate average
	average := calculateAverage(scores)
	fmt.Printf("Average score: %.2f\n", average)

	// Find max and min
	max, min := findMaxMin(scores)
	fmt.Printf("Max: %d, Min: %d\n", max, min)

	// Reverse array
	reversed := reverseArray(scores)
	fmt.Printf("Original: %v\n", scores)
	fmt.Printf("Reversed: %v\n\n", reversed)

	// 14. Days of week example
	fmt.Println("14. Real-world Example:")
	weekdays := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println("Workdays:")
	for i, day := range weekdays {
		fmt.Printf("%d. %s\n", i+1, day)
	}
}

// Calculate average of array
func calculateAverage(arr [5]int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

// Find maximum and minimum in array
func findMaxMin(arr [5]int) (int, int) {
	max := arr[0]
	min := arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return max, min
}

// Reverse array
func reverseArray(arr [5]int) [5]int {
	var reversed [5]int
	for i, value := range arr {
		reversed[len(arr)-1-i] = value
	}
	return reversed
}
