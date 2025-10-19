package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=== Maps in Go ===")

	// 1. Creating maps
	fmt.Println("1. Creating Maps:")

	// Using make
	ages := make(map[string]int)
	fmt.Printf("Empty map: %v\n", ages)

	// Using literal
	scores := map[string]int{
		"Alice":   85,
		"Bob":     90,
		"Charlie": 78,
	}
	fmt.Printf("Map literal: %v\n\n", scores)

	// 2. Adding and updating
	fmt.Println("2. Adding and Updating:")
	ages["Alice"] = 25
	ages["Bob"] = 30
	fmt.Printf("After adding: %v\n", ages)

	ages["Alice"] = 26 // update
	fmt.Printf("After updating Alice: %v\n\n", ages)

	// 3. Getting values
	fmt.Println("3. Getting Values:")
	age := ages["Alice"]
	fmt.Printf("Alice's age: %d\n", age)

	unknownAge := ages["Unknown"]
	fmt.Printf("Unknown person's age: %d (zero value)\n\n", unknownAge)

	// 4. Check if key exists
	fmt.Println("4. Check Key Existence:")
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice exists with age: %d\n", age)
	}

	if age, exists := ages["Unknown"]; exists {
		fmt.Printf("Unknown exists with age: %d\n", age)
	} else {
		fmt.Println("Unknown doesn't exist")
	}
	fmt.Println()

	// 5. Delete key
	fmt.Println("5. Delete Key:")
	fmt.Printf("Before delete: %v\n", ages)
	delete(ages, "Bob")
	fmt.Printf("After deleting Bob: %v\n\n", ages)

	// 6. Map length
	fmt.Println("6. Map Length:")
	fmt.Printf("Number of entries: %d\n\n", len(scores))

	// 7. Iterating over map
	fmt.Println("7. Iterating Over Map:")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}
	fmt.Println()

	// 8. Iterate keys only
	fmt.Println("8. Keys Only:")
	for name := range scores {
		fmt.Printf("- %s\n", name)
	}
	fmt.Println()

	// 9. Map with struct values
	fmt.Println("9. Map with Struct Values:")
	type Student struct {
		Name  string
		Grade string
		Score int
	}

	students := map[string]Student{
		"S001": {Name: "Alice", Grade: "A", Score: 90},
		"S002": {Name: "Bob", Grade: "B", Score: 85},
		"S003": {Name: "Charlie", Grade: "A", Score: 92},
	}

	for id, student := range students {
		fmt.Printf("%s: %s (Grade: %s, Score: %d)\n",
			id, student.Name, student.Grade, student.Score)
	}
	fmt.Println()

	// 10. Nested maps
	fmt.Println("10. Nested Maps:")
	grades := map[string]map[string]int{
		"Alice": {
			"Math":    90,
			"English": 85,
			"Science": 88,
		},
		"Bob": {
			"Math":    85,
			"English": 90,
			"Science": 82,
		},
	}

	for student, subjects := range grades {
		fmt.Printf("%s's grades:\n", student)
		for subject, grade := range subjects {
			fmt.Printf("  %s: %d\n", subject, grade)
		}
	}
	fmt.Println()

	// 11. Practical examples
	fmt.Println("11. Practical Examples:")

	// Word frequency count
	text := "hello world hello go world go go"
	frequencies := wordFrequency(text)
	fmt.Println("Word frequencies:")
	for word, count := range frequencies {
		fmt.Printf("%s: %d\n", word, count)
	}
	fmt.Println()

	// Group by first letter
	words := []string{"apple", "banana", "avocado", "berry", "apricot"}
	grouped := groupByFirstLetter(words)
	fmt.Println("Grouped by first letter:")
	for letter, group := range grouped {
		fmt.Printf("%s: %v\n", letter, group)
	}
	fmt.Println()

	// 12. Two-sum problem
	fmt.Println("12. Two-Sum Problem:")
	nums := []int{2, 7, 11, 15}
	target := 9
	indices := twoSum(nums, target)
	if len(indices) == 2 {
		fmt.Printf("Indices: [%d, %d] (values: %d + %d = %d)\n",
			indices[0], indices[1], nums[indices[0]], nums[indices[1]], target)
	}
	fmt.Println()

	// 13. Merge maps
	fmt.Println("13. Merge Maps:")
	map1 := map[string]int{"A": 1, "B": 2}
	map2 := map[string]int{"B": 3, "C": 4}
	merged := mergeMaps(map1, map2)
	fmt.Printf("Map1: %v\n", map1)
	fmt.Printf("Map2: %v\n", map2)
	fmt.Printf("Merged: %v\n\n", merged)

	// 14. Sorted iteration
	fmt.Println("14. Sorted Map Iteration:")
	inventory := map[string]int{
		"Orange": 5,
		"Apple":  3,
		"Banana": 7,
		"Grape":  2,
	}

	// Get sorted keys
	keys := make([]string, 0, len(inventory))
	for k := range inventory {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Sorted inventory:")
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, inventory[k])
	}
}

// Word frequency counter
func wordFrequency(text string) map[string]int {
	frequencies := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		frequencies[word]++
	}

	return frequencies
}

// Group strings by first letter
func groupByFirstLetter(words []string) map[string][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		if len(word) > 0 {
			firstLetter := string(word[0])
			groups[firstLetter] = append(groups[firstLetter], word)
		}
	}

	return groups
}

// Two-sum problem using map
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[num] = i
	}

	return []int{}
}

// Merge two maps
func mergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	for k, v := range map1 {
		result[k] = v
	}

	for k, v := range map2 {
		result[k] = v // map2 overwrites map1 if key exists
	}

	return result
}
