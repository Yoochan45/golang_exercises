# Functions

## Summary
Functions adalah blok kode yang dapat digunakan kembali. Di Go, functions adalah first-class citizens - bisa disimpan dalam variabel, dijadikan parameter, atau dikembalikan dari function lain.

## Basic Function Syntax

### Function tanpa parameter dan return value
```go
func sayHello() {
    fmt.Println("Hello!")
}
```

### Function dengan parameter
```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

### Function dengan return value
```go
func add(a int, b int) int {
    return a + b
}

// Shorthand untuk parameter dengan tipe sama
func multiply(a, b int) int {
    return a * b
}
```

### Multiple return values
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Named return values
```go
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return  // naked return
}
```

## Variadic Functions
Function yang menerima jumlah parameter variable:
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
sum(1, 2, 3)
sum(1, 2, 3, 4, 5)
```

## Anonymous Functions & Closures

### Anonymous function
```go
func() {
    fmt.Println("I'm anonymous!")
}()  // langsung dijalankan
```

### Function as variable
```go
add := func(a, b int) int {
    return a + b
}
result := add(3, 5)
```

### Closure
Function yang "mengingat" variabel dari scope luarnya:
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Higher-Order Functions
Function yang menerima function sebagai parameter:
```go
func apply(numbers []int, fn func(int) int) []int {
    result := make([]int, len(numbers))
    for i, n := range numbers {
        result[i] = fn(n)
    }
    return result
}
```

## Defer Statement
`defer` menunda eksekusi function hingga function pemanggil selesai:
```go
func readFile() {
    file, _ := os.Open("file.txt")
    defer file.Close()  // akan dijalankan sebelum function return
    
    // proses file...
}
```

## Exercises

1. Buat function `calculator` dengan 4 operasi dasar (+, -, *, /)
2. Buat function `findMax` yang mengembalikan nilai terbesar dari slice
3. Implementasi function `fibonacci(n int) int` 
4. Buat function yang mengembalikan nama lengkap dan inisial
5. Buat closure untuk membuat generator ID unik
6. Implementasi higher-order function untuk filter slice

## Common Patterns

### Error handling pattern
```go
func doSomething() error {
    if err := validate(); err != nil {
        return err
    }
    // logic...
    return nil
}
```

### Multiple defer (LIFO order)
```go
func example() {
    defer fmt.Println("3")
    defer fmt.Println("2")
    defer fmt.Println("1")
}  // Output: 1, 2, 3
```

## Best Practices
- Function names harus descriptive
- Keep functions small and focused (single responsibility)
- Return errors instead of panic
- Use named returns untuk readability di function kompleks
- Defer untuk cleanup operations (close file, unlock mutex, etc)