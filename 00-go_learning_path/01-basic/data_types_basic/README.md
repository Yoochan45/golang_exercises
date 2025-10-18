# Basic Data Types

## Summary
Go adalah statically typed language dengan berbagai tipe data built-in untuk numeric, string, boolean, dan lainnya.

## Numeric Types

### Integer Types
- `int8`, `int16`, `int32`, `int64` - signed integers
- `uint8`, `uint16`, `uint32`, `uint64` - unsigned integers
- `int`, `uint` - platform dependent (32 or 64 bit)
- `byte` - alias untuk `uint8`
- `rune` - alias untuk `int32` (Unicode code point)

### Floating Point
- `float32` - 32-bit floating point
- `float64` - 64-bit floating point (default)

### Complex Numbers
- `complex64` - complex numbers dengan float32 parts
- `complex128` - complex numbers dengan float64 parts

## String Type
String adalah sequence of bytes (immutable).
```go
var message string = "Hello, Go!"
```

## Boolean Type
Boolean hanya memiliki dua nilai: `true` atau `false`.
```go
var isActive bool = true
```

## Type Conversion
Go tidak melakukan implicit type conversion. Harus explicit:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

## Syntax Examples
```go
// Integer
var age int = 25

// Float
var price float64 = 99.99

// String
var name string = "John"

// Boolean
var isActive bool = true

// Type conversion
var x int = 10
var y float64 = float64(x)
```

## Exercises

1. Deklarasikan variable dengan berbagai tipe numeric dan print nilainya
2. Konversi int ke float64 dan sebaliknya
3. Buat string dengan karakter spesial dan emoji
4. Hitung luas persegi panjang dengan float64
5. Eksplorasi batas nilai untuk int8, int16, int32

## Common Pitfalls
- Integer division truncates: `5 / 2 = 2` (not 2.5)
- String adalah immutable - tidak bisa mengubah karakter individual
- Tidak ada automatic type conversion - harus explicit cast