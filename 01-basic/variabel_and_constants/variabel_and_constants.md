# Variables and Constants

## Summary
Variables adalah tempat penyimpanan data yang nilainya dapat berubah. Constants adalah nilai tetap yang tidak dapat diubah setelah dideklarasikan.

## Variable Declaration

### 1. Dengan tipe data eksplisit
```go
var name string = "John"
var age int = 25
```

### 2. Dengan type inference
```go
var name = "John"  // Go otomatis tahu ini string
var age = 25       // Go otomatis tahu ini int
```

### 3. Short declaration (hanya di dalam function)
```go
name := "John"
age := 25
```

### 4. Multiple declaration
```go
var x, y, z int
var a, b = 1, "hello"
```

## Constants

Constants dideklarasikan dengan keyword `const`:
```go
const Pi = 3.14159
const AppName = "MyApp"
```

### Typed vs Untyped Constants
```go
const TypedInt int = 100        // typed
const UntypedInt = 100          // untyped (lebih fleksibel)
```

## Zero Values
Variabel tanpa nilai awal memiliki zero value:
- `0` untuk numeric types
- `false` untuk boolean
- `""` (empty string) untuk string
- `nil` untuk pointers, slices, maps, channels, functions, interfaces

## Syntax Rules
- Variable names harus dimulai dengan huruf atau underscore
- Case-sensitive: `name` berbeda dengan `Name`
- Exported names dimulai dengan huruf kapital: `Name` (public), `name` (private)

## Exercises

1. Deklarasikan variable `name`, `age`, `isStudent` dengan nilai masing-masing
2. Ubah nilai variable dan print hasilnya
3. Buat constant untuk nilai PI dan gunakan untuk menghitung luas lingkaran
4. Deklarasikan multiple variables dalam satu line
5. Coba compile code dengan unused variable (akan error!)

## Common Mistakes
- Unused variables akan menyebabkan compilation error
- Cannot reassign constants
- Short declaration `:=` hanya bisa digunakan di dalam function