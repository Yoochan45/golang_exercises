package main

import "fmt"

func main() {
	fmt.Println("=== Pointers in Go ===")

	// 1. Basic pointer
	fmt.Println("1. Basic Pointer:")
	x := 42
	ptr := &x // ptr holds address of x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Pointer ptr: %p\n", ptr)
	fmt.Printf("Value at ptr: %d\n\n", *ptr)

	// 2. Modifying through pointer
	fmt.Println("2. Modifying Through Pointer:")
	fmt.Printf("Before: x = %d\n", x)
	*ptr = 100
	fmt.Printf("After *ptr = 100: x = %d\n\n", x)

	// 3. Declaring pointers
	fmt.Println("3. Declaring Pointers:")
	var p1 *int
	fmt.Printf("Nil pointer: %v (nil: %t)\n", p1, p1 == nil)

	p2 := new(int)
	*p2 = 42
	fmt.Printf("Using new: %v (value: %d)\n\n", p2, *p2)

	// 4. Pass by value vs pass by pointer
	fmt.Println("4. Pass by Value vs Pass by Pointer:")

	a := 10
	fmt.Printf("Original: a = %d\n", a)

	incrementValue(a)
	fmt.Printf("After incrementValue: a = %d (unchanged)\n", a)

	incrementPointer(&a)
	fmt.Printf("After incrementPointer: a = %d (modified)\n\n", a)

	// 5. Pointers with structs
	fmt.Println("5. Pointers with Structs:")

	person1 := Person{Name: "Alice", Age: 25}
	person2 := &Person{Name: "Bob", Age: 30}

	fmt.Printf("person1 (value): %+v\n", person1)
	fmt.Printf("person2 (pointer): %+v\n", person2)

	// Automatic dereferencing
	person2.Name = "Robert"
	fmt.Printf("Modified person2: %+v\n\n", person2)

	// 6. Struct method with pointer receiver
	fmt.Println("6. Pointer Receiver Methods:")
	p := Person{Name: "Charlie", Age: 35}
	fmt.Printf("Before: %+v\n", p)
	p.UpdateAge(36)
	fmt.Printf("After UpdateAge: %+v\n\n", p)

	// 7. Nil pointer check
	fmt.Println("7. Nil Pointer Check:")
	var ptr1 *int
	fmt.Printf("ptr1 is nil: %t\n", ptr1 == nil)

	// Assign a value conditionally to demonstrate nil check
	someValue := 99
	if false { // Change to true to assign
		ptr1 = &someValue
	}

	if ptr1 != nil {
		fmt.Printf("Value: %d\n", *ptr1)
	} else {
		fmt.Println("Cannot dereference nil pointer")
	}
	fmt.Println()

	// 8. Pointer to pointer
	fmt.Println("8. Pointer to Pointer:")
	value := 42
	ptr2 := &value
	pp := &ptr2

	fmt.Printf("value: %d\n", value)
	fmt.Printf("*ptr2: %d\n", *ptr2)
	fmt.Printf("**pp: %d\n\n", **pp)

	// 9. Returning pointer from function
	fmt.Println("9. Returning Pointer from Function:")
	newPerson := createPerson("David", 28)
	fmt.Printf("New person: %+v\n\n", newPerson)

	// 10. Swap function using pointers
	fmt.Println("10. Swap Function:")
	x1, y1 := 10, 20
	fmt.Printf("Before swap: x=%d, y=%d\n", x1, y1)
	swap(&x1, &y1)
	fmt.Printf("After swap: x=%d, y=%d\n\n", x1, y1)

	// 11. Slice and map are already reference types
	fmt.Println("11. Slice/Map (Already Reference Types):")
	numbers := []int{1, 2, 3}
	fmt.Printf("Before: %v\n", numbers)
	modifySlice(numbers)
	fmt.Printf("After: %v (modified without pointer)\n\n", numbers)

	// 12. Pointer in slice
	fmt.Println("12. Slice of Pointers:")
	people := []*Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	for i, person := range people {
		fmt.Printf("%d. %+v\n", i+1, person)
	}

	// Modify through pointer
	people[0].Age = 26
	fmt.Printf("\nAfter modifying people[0]: %+v\n\n", people[0])

	// 13. Pointer comparison
	fmt.Println("13. Pointer Comparison:")
	a1 := 42
	a2 := 42
	ptr1 = &a1
	ptr2 = &a2
	ptr3 := ptr1

	fmt.Printf("ptr1 == ptr2: %t (different addresses)\n", ptr1 == ptr2)
	fmt.Printf("ptr1 == ptr3: %t (same address)\n", ptr1 == ptr3)
	fmt.Printf("*ptr1 == *ptr2: %t (same values)\n\n", *ptr1 == *ptr2)

	// 14. Find function returning pointer
	fmt.Println("14. Find Function (Returns Pointer or Nil):")
	users := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	found := findPerson(users, "Bob")
	if found != nil {
		fmt.Printf("Found: %+v\n", found)
	} else {
		fmt.Println("Not found")
	}

	notFound := findPerson(users, "David")
	if notFound != nil {
		fmt.Printf("Found: %+v\n", notFound)
	} else {
		fmt.Println("David not found")
	}
}

// Pass by value - does not modify original
func incrementValue(x int) {
	x++
}

// Pass by pointer - modifies original
func incrementPointer(x *int) {
	*x++
}

// Struct type
type Person struct {
	Name string
	Age  int
}

// Pointer receiver method
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// Return pointer from function
func createPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// Swap two values using pointers
func swap(x, y *int) {
	*x, *y = *y, *x
}

// Slice is already reference type
func modifySlice(s []int) {
	s[0] = 999
}

// Find function returning pointer or nil
func findPerson(people []Person, name string) *Person {
	for i := range people {
		if people[i].Name == name {
			return &people[i]
		}
	}
	return nil
}
