# Slices

## Summary
Slice adalah abstraksi dari array yang lebih fleksibel. Slice memiliki ukuran dinamis dan merupakan reference type (tidak seperti array yang value type).

## Slice Structure
Slice memiliki 3 komponen:
- **Pointer** ke underlying array
- **Length** - jumlah elemen saat ini
- **Capacity** - maksimum elemen yang bisa ditampung

## Creating Slices

### From array
```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2, 3, 4]
```

### Using make
```go
slice := make([]int, 5)       // length=5, capacity=5
slice := make([]int, 5, 10)   // length=5, capacity=10
```

### Using literal
```go
slice := []int{1, 2, 3, 4, 5}
```

### Empty slice
```go
var slice []int  // nil slice
slice := []int{}  // empty slice (not nil)
```

## Slice Operations

### Accessing elements
```go
slice := []int{1, 2, 3, 4, 5}
first := slice[0]
last := slice[len(slice)-1]
```

### Append
```go
slice = append(slice, 6)           // add one element
slice = append(slice, 7, 8, 9)     // add multiple
slice = append(slice, other...)    // append another slice
```

### Slicing
```go
slice := []int{1, 2, 3, 4, 5}
slice[1:4]    // [2, 3, 4]
slice[:3]     // [1, 2, 3]
slice[2:]     // [3, 4, 5]
slice[:]      // copy all
```

### Copy
```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

### Length and Capacity
```go
len(slice)  // jumlah elemen
cap(slice)  // kapasitas
```

## Iterating Over Slices
```go
for index, value := range slice {
    fmt.Printf("%d: %d\n", index, value)
}
```

## Multi-dimensional Slices
```go
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

## Slice Tricks

### Remove element at index
```go
slice = append(slice[:i], slice[i+1:]...)
```

### Insert element at index
```go
slice = append(slice[:i], append([]int{x}, slice[i:]...)...)
```

### Filter slice
```go
result := slice[:0]
for _, v := range slice {
    if condition(v) {
        result = append(result, v)
    }
}
```

## Important Notes
- Slices are reference types (sharing underlying array)
- Modifying slice element modifies the underlying array
- `append` may create new underlying array if capacity exceeded
- nil slice vs empty slice behaves differently
- Be careful with slice expressions - they share memory

## Exercises

1. Buat slice dari 1-10 dan filter hanya angka genap
2. Implementasi function untuk remove duplicates dari slice
3. Reverse slice tanpa membuat slice baru
4. Merge dua sorted slices menjadi satu sorted slice
5. Split slice menjadi chunks dengan ukuran tertentu
6. Implementasi binary search pada sorted slice

## Common Patterns

### Initialize with values
```go
slice := make([]int, 0, 10)
for i := 1; i <= 5; i++ {
    slice = append(slice, i)
}
```

### Remove element
```go
// Remove at index i
slice = append(slice[:i], slice[i+1:]...)
```

### Check if empty
```go
if len(slice) == 0 {
    // empty
}
```