# Arrays

## Summary
Array adalah koleksi elemen dengan tipe yang sama dan ukuran tetap (fixed size). Ukuran array adalah bagian dari tipenya.

## Array Declaration

### Basic declaration
```go
var numbers [5]int  // array dengan 5 elemen int, semua bernilai 0
```

### Declaration with initialization
```go
var numbers = [5]int{1, 2, 3, 4, 5}
numbers := [5]int{1, 2, 3, 4, 5}
```

### Let compiler count the size
```go
numbers := [...]int{1, 2, 3, 4, 5}  // compiler akan hitung ukurannya
```

### Partial initialization
```go
numbers := [5]int{1, 2, 3}  // sisanya = 0
```

### Specific index initialization
```go
numbers := [5]int{0: 10, 2: 20, 4: 30}  // index 1,3 = 0
```

## Accessing Array Elements
```go
numbers := [5]int{1, 2, 3, 4, 5}
first := numbers[0]   // 1
last := numbers[4]    // 5
numbers[2] = 10       // ubah element index 2
```

## Array Length
```go
numbers := [5]int{1, 2, 3, 4, 5}
length := len(numbers)  // 5
```

## Iterating Over Arrays

### Using for loop with index
```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

### Using range
```go
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Hanya value
for _, value := range numbers {
    fmt.Println(value)
}
```

## Multi-dimensional Arrays
```go
var matrix [3][3]int  // 3×3 matrix

// With initialization
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

## Array vs Slice
**Arrays**:
- Fixed size
- Size adalah bagian dari tipe: `[5]int` ≠ `[10]int`
- Passed by value (copy)
- Jarang digunakan langsung

**Slices** (akan dipelajari next):
- Dynamic size
- More flexible
- Passed by reference
- Lebih umum digunakan

## Important Notes
- Array size tidak bisa diubah setelah deklarasi
- Array adalah value type (copy on assignment)
- Array dengan ukuran berbeda adalah tipe berbeda
- Index dimulai dari 0

## Exercises

1. Buat array untuk menyimpan nama 5 hari kerja
2. Hitung rata-rata dari array angka
3. Temukan nilai terbesar dan terkecil dalam array
4. Reverse array (balik urutannya)
5. Buat matrix 3×3 dan hitung jumlah diagonal
6. Copy array ke array lain dan modifikasi copy-nya

## Common Patterns

### Initialize with loop
```go
var numbers [10]int
for i := 0; i < len(numbers); i++ {
    numbers[i] = i + 1
}
```

### Sum all elements
```go
sum := 0
for _, value := range numbers {
    sum += value
}
```