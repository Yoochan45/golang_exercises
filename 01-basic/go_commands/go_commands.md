# Go Commands - Getting Started

## Summary
Go menyediakan berbagai command line tools untuk mengelola project, compile, run, dan test aplikasi Go. Command ini adalah fondasi untuk bekerja dengan Go.

## Essential Go Commands

### 1. `go version`
Mengecek versi Go yang terinstall.
```bash
go version
```

### 2. `go mod init`
Membuat module baru (project Go).
```bash
go mod init nama-module
```

### 3. `go run`
Compile dan menjalankan program Go tanpa membuat binary.
```bash
go run main.go
```

### 4. `go build`
Compile program menjadi binary executable.
```bash
go build
go build -o nama-output
```

### 5. `go install`
Compile dan install binary ke $GOPATH/bin.
```bash
go install
```

### 6. `go get`
Download dan install dependencies.
```bash
go get github.com/package/name
```

### 7. `go mod tidy`
Menambahkan missing dependencies dan menghapus unused dependencies.
```bash
go mod tidy
```

### 8. `go test`
Menjalankan unit tests.
```bash
go test
go test -v
go test ./...
```

### 9. `go fmt`
Format kode sesuai Go style guide.
```bash
go fmt ./...
```

## Exercises

1. Buat module baru dengan nama `myapp`
2. Jalankan hello world program dengan `go run`
3. Build program menjadi binary
4. Cek dependencies dengan `go mod tidy`

## Next Steps
Lanjut ke `variables_and_constants` untuk belajar dasar-dasar pemrograman Go.