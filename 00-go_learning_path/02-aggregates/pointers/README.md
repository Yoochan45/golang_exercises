# Pointers

## Summary
Pointer adalah variable yang menyimpan alamat memory dari variable lain. Pointers memungkinkan kita untuk memodifikasi nilai variable dan menghindari copying data besar.

## Pointer Basics

### Address operator (&)
Mendapatkan alamat memory dari variable:
```go
x := 42
ptr := &x  // ptr holds the address of x
```

### Dereference operator (*)
Mengakses nilai dari alamat yang ditunjuk pointer:
```go
x := 42
ptr := &x
fmt.Println(*ptr)  // 42
*ptr = 100         // modify value through pointer
```

### Declaring pointers
```go
var ptr *int       // pointer to int (nil)
var p *Person      // pointer to Person struct
```

## Creating Pointers

### Using address operator
```go
x := 42
ptr := &x
```

### Using new
```go
ptr := new(int)  // allocates memory, returns *int
*ptr = 42
```

## Pointer with Functions

### Pass by value (copy)
```go
func increment(x int) {
    x++  // only modifies local copy
}
```

### Pass by pointer (reference)
```go
func increment(x *int) {
    *x++  // modifies original value
}
```

## Pointers and Structs

### Accessing struct fields through pointer
```go
type Person struct {
    Name string
    Age  int
}

p := &Person{Name: "Alice", Age: 25}
p.Name = "Bob"  // automatic dereferencing
(*p).Name = "Bob"  // explicit (same as above)
```

## Nil Pointers
```go
var ptr *int
if ptr == nil {
    fmt.Println("pointer is nil")
}
```

## Pointer to Pointer
```go
x := 42
ptr := &x
pptr := &ptr  // **int
```

## Pointer Arithmetic
Go does NOT support pointer arithmetic (unlike C):
```go
ptr++  // NOT ALLOWED in Go
```

## Important Notes
- Go is "pass by value" language
- Pointers are used to avoid copying large data
- Automatic dereferencing for struct fields
- Cannot get pointer to literals: `&42` is invalid
- Nil pointer dereference causes panic
- Use pointers for structs in functions if you need to modify them

## When to Use Pointers?

### Use pointers when:
- Function needs to modify the argument
- Struct is large (avoid copying)
- Need to return "not found" (can return nil)
- Implementing methods that modify receiver

### Don't use pointers when:
- Small data types (int, bool, etc)
- Value semantics make more sense
- Thread safety is a concern

## Exercises

1. Buat function swap yang menukar nilai dua variable menggunakan pointers
2. Buat function untuk modify struct field via pointer
3. Implementasi linked list node dengan pointer
4. Buat function yang return pointer (or nil if not found)
5. Compare pointer vs value semantics dengan benchmark
6. Praktik nil pointer check

## Common Patterns

### Nil check
```go
if ptr != nil {
    fmt.Println(*ptr)
}
```

### Return pointer from function
```go
func createPerson(name string) *Person {
    return &Person{Name: name}
}
```

### Pointer receiver method
```go
func (p *Person) UpdateName(name string) {
    p.Name = name
}
```