package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Slices in Go ===")

	// 1. Creating slices
	fmt.Println("1. Creating Slices:")

	// From array
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4]
	fmt.Printf("From array: %v\n", slice1)

	// Using make
	slice2 := make([]int, 5)
	fmt.Printf("Using make: %v\n", slice2)

	// Using literal
	slice3 := []int{10, 20, 30}
	fmt.Printf("Using literal: %v\n", slice3)

	// Nil slice
	var slice4 []int
	fmt.Printf("Nil slice: %v (nil: %t)\n\n", slice4, slice4 == nil)

	// 2. Length and Capacity
	fmt.Println("2. Length and Capacity:")
	slice := make([]int, 3, 5)
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Length: %d, Capacity: %d\n\n", len(slice), cap(slice))

	// 3. Append operation
	fmt.Println("3. Append Operation:")
	numbers := []int{1, 2, 3}
	fmt.Printf("Original: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 4)
	fmt.Printf("After append 4: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 5, 6, 7)
	fmt.Printf("After append 5,6,7: %v (len=%d, cap=%d)\n\n", numbers, len(numbers), cap(numbers))

	// 4. Slicing operations
	fmt.Println("4. Slicing Operations:")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("original[2:5]: %v\n", original[2:5])
	fmt.Printf("original[:5]: %v\n", original[:5])
	fmt.Printf("original[5:]: %v\n", original[5:])
	fmt.Printf("original[:]: %v\n\n", original[:])

	// 5. Copy slices
	fmt.Println("5. Copy Slices:")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 99
	fmt.Printf("Source: %v\n", src)
	fmt.Printf("Destination: %v\n\n", dst)

	// 6. Slice reference behavior
	fmt.Println("6. Slice Reference Behavior:")
	base := []int{1, 2, 3, 4, 5}
	derived := base[1:4]
	fmt.Printf("Base: %v\n", base)
	fmt.Printf("Derived: %v\n", derived)

	derived[0] = 99 // modifies base too!
	fmt.Printf("After modifying derived[0]:\n")
	fmt.Printf("Base: %v\n", base)
	fmt.Printf("Derived: %v\n\n", derived)

	// 7. Iterating slices
	fmt.Println("7. Iterating Slices:")
	colors := []string{"Red", "Green", "Blue"}
	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}
	fmt.Println()

	// 8. Multi-dimensional slices
	fmt.Println("8. Multi-dimensional Slices:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()

	// 9. Variadic function with slices
	fmt.Println("9. Variadic Functions:")
	nums := []int{1, 2, 3, 4, 5}
	total := sum(nums...)
	fmt.Printf("Sum of %v: %d\n\n", nums, total)

	// 10. Practical examples
	fmt.Println("10. Practical Examples:")

	// Filter even numbers
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filterEven(data)
	fmt.Printf("Original: %v\n", data)
	fmt.Printf("Even numbers: %v\n", evens)

	// Remove duplicates
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	unique := removeDuplicates(duplicates)
	fmt.Printf("With duplicates: %v\n", duplicates)
	fmt.Printf("Unique: %v\n", unique)

	// Reverse slice
	toReverse := []int{1, 2, 3, 4, 5}
	reverse(toReverse)
	fmt.Printf("Reversed: %v\n\n", toReverse)

	// 11. Slice tricks
	fmt.Println("11. Slice Tricks:")

	// Remove element at index
	items := []string{"A", "B", "C", "D", "E"}
	fmt.Printf("Original: %v\n", items)
	items = removeAt(items, 2)
	fmt.Printf("After removing index 2: %v\n", items)

	// Insert element
	items = insertAt(items, 1, "X")
	fmt.Printf("After inserting 'X' at index 1: %v\n\n", items)

	// 12. Sorting slices
	fmt.Println("12. Sorting Slices:")
	unsorted := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("Unsorted: %v\n", unsorted)
	sort.Ints(unsorted)
	fmt.Printf("Sorted: %v\n", unsorted)

	names := []string{"Charlie", "Alice", "Bob"}
	fmt.Printf("Unsorted names: %v\n", names)
	sort.Strings(names)
	fmt.Printf("Sorted names: %v\n", names)
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// Filter even numbers
func filterEven(numbers []int) []int {
	result := []int{}
	for _, n := range numbers {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

// Remove duplicates
func removeDuplicates(slice []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, value := range slice {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}

	return result
}

// Reverse slice in-place
func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Remove element at index
func removeAt(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// Insert element at index
func insertAt(slice []string, index int, value string) []string {
	slice = append(slice[:index], append([]string{value}, slice[index:]...)...)
	return slice
}
