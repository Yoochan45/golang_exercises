# Structs

## Summary
Struct adalah tipe data composite yang mengelompokkan fields dengan tipe berbeda. Struct adalah building block untuk membuat custom types di Go.

## Defining Structs

### Basic struct
```go
type Person struct {
    Name string
    Age  int
}
```

### Struct with different types
```go
type Employee struct {
    ID        int
    Name      string
    Salary    float64
    IsActive  bool
}
```

## Creating Struct Instances

### Using struct literal
```go
p1 := Person{Name: "Alice", Age: 25}
p2 := Person{"Bob", 30}  // positional (not recommended)
```

### Using new keyword
```go
p := new(Person)  // returns *Person with zero values
```

### Zero value
```go
var p Person  // all fields have zero values
```

## Accessing Fields
```go
person := Person{Name: "Alice", Age: 25}
fmt.Println(person.Name)  // Alice
person.Age = 26           // modify field
```

## Anonymous Structs
```go
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  25,
}
```

## Embedded Structs
```go
type Address struct {
    Street string
    City   string
}

type Person struct {
    Name    string
    Address Address  // nested
}
```

## Struct Embedding (Composition)
```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person       // embedded (promoted fields)
    Company string
}

emp := Employee{
    Person:  Person{Name: "Alice", Age: 25},
    Company: "TechCorp",
}
fmt.Println(emp.Name)  // direct access to embedded field
```

## Struct Tags
Metadata untuk fields (used by encoding/json, etc):
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
    Age   int    `json:"-"`  // ignored
}
```

## Comparing Structs
```go
p1 := Person{Name: "Alice", Age: 25}
p2 := Person{Name: "Alice", Age: 25}
fmt.Println(p1 == p2)  // true if all fields are comparable
```

## Struct Methods
Will be covered in detail in `03-oop_style/methods/`:
```go
func (p Person) Greet() {
    fmt.Printf("Hello, I'm %s\n", p.Name)
}
```

## Important Notes
- Struct fields with capital letter are exported (public)
- Struct fields with lowercase are unexported (private to package)
- Structs are value types (copied on assignment)
- Use pointers to structs for large structs or when you need to modify
- Empty struct `struct{}` occupies zero bytes

## Exercises

1. Buat struct `Book` dengan fields: Title, Author, Year, Pages
2. Buat struct `Rectangle` dan function untuk hitung luas dan keliling
3. Buat struct `Student` dengan embedded `Person` struct
4. Implementasi struct dengan method untuk calculate age from birth year
5. Buat slice of structs dan sort berdasarkan field tertentu
6. Practice JSON encoding/decoding dengan struct tags

## Common Patterns

### Constructor function
```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}
```

### Pointer to struct
```go
func updateAge(p *Person, newAge int) {
    p.Age = newAge
}
```