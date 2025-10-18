# Control Structures

## Summary
Control structures mengontrol alur eksekusi program: percabangan (if-else, switch) dan perulangan (for loops).

## If-Else Statements

### Basic if
```go
if condition {
    // code
}
```

### If-else
```go
if condition {
    // code if true
} else {
    // code if false
}
```

### If with initialization
```go
if value := getValue(); value > 0 {
    // value only accessible in if block
}
```

## Switch Statements

### Basic switch
```go
switch value {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

### Switch without condition (like if-else chain)
```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

## For Loops

Go hanya memiliki `for` loop (tidak ada while loop).

### Classic for loop
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### While-style loop
```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

### Infinite loop
```go
for {
    // runs forever until break
}
```

### Range loop
```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

## Break and Continue

- `break` - keluar dari loop
- `continue` - skip ke iterasi berikutnya

## Exercises

1. Buat program yang mengecek apakah angka genap atau ganjil
2. Implementasi grade calculator dengan switch
3. Print angka 1-100, skip kelipatan 5
4. Buat nested loop untuk print matrix
5. Implementasi FizzBuzz (1-100)

## Common Patterns
- No parentheses around condition (but braces required)
- Multiple conditions dapat digabung dengan `&&` (and) atau `||` (or)
- Switch tidak perlu `break` (automatic)