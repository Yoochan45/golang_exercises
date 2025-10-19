# Maps

## Summary
Map adalah koleksi pasangan key-value yang tidak berurutan. Map adalah reference type dan sangat efisien untuk lookup data.

## Creating Maps

### Using make
```go
ages := make(map[string]int)
```

### Using map literal
```go
ages := map[string]int{
    "Alice": 25,
    "Bob": 30,
}
```

### Empty map literal
```go
ages := map[string]int{}
```

## Map Operations

### Add or update
```go
ages["Charlie"] = 35
```

### Get value
```go
age := ages["Alice"]  // 25
age := ages["Unknown"]  // 0 (zero value if not exists)
```

### Check if key exists
```go
age, exists := ages["Alice"]
if exists {
    fmt.Println(age)
}
```

### Delete key
```go
delete(ages, "Alice")
```

### Get length
```go
count := len(ages)
```

## Iterating Over Maps
```go
for key, value := range ages {
    fmt.Printf("%s: %d\n", key, value)
}

// Keys only
for key := range ages {
    fmt.Println(key)
}

// Values only
for _, value := range ages {
    fmt.Println(value)
}
```

## Map with Struct Values
```go
type Person struct {
    Name string
    Age  int
}

people := map[string]Person{
    "emp1": {Name: "Alice", Age: 25},
    "emp2": {Name: "Bob", Age: 30},
}
```

## Nested Maps
```go
grades := map[string]map[string]int{
    "Alice": {"Math": 90, "English": 85},
    "Bob":   {"Math": 85, "English": 90},
}
```

## Important Notes
- Maps are reference types
- Map iteration order is **random** (not guaranteed)
- Zero value of map is `nil` - cannot add to nil map
- Maps are not safe for concurrent use (need sync.Map or mutex)
- Keys must be comparable types (can use ==)
- Cannot take address of map element

## Exercises

1. Buat map untuk menyimpan student scores dan hitung rata-rata
2. Count word frequency dalam sebuah text
3. Group slice of strings by first character
4. Implementasi two-sum problem menggunakan map
5. Merge dua maps
6. Invert map (swap keys and values)

## Common Patterns

### Check and initialize
```go
if ages == nil {
    ages = make(map[string]int)
}
```

### Safe access
```go
if age, ok := ages["Alice"]; ok {
    fmt.Println(age)
} else {
    fmt.Println("Not found")
}
```

### Count occurrences
```go
counts := make(map[string]int)
for _, item := range items {
    counts[item]++
}
```