# Go Learning Path - From Zero to REST API

Repositori pembelajaran komprehensif untuk menguasai Go dari dasar hingga REST API development.

## 📚 Deskripsi

Repo ini dibuat untuk membantu siapa saja yang ingin belajar Go programming dari nol hingga mahir membuat REST API. Materi disusun secara terstruktur dan progresif dengan contoh kode yang praktis.


## 🗂️ Struktur Pembelajaran

```
golang_exercises/
│
├── 00-go_learning_path/          # 🆕 Tutorial lengkap dari zero
│   ├── 01-basic/                 # Fundamental Go
│   ├── 02-aggregates/            # Arrays, Slices, Maps, Structs
│   ├── 03-oop_style/             # Methods, Interfaces, Embedding
│   ├── 04-advanced/              # Error handling, Concurrency, Testing
│   ├── 05-files_io/              # File operations, JSON, ENV
│   ├── 06-web_basics/            # HTTP Server/Client, Routing
│   ├── 07-rest_api/              # CRUD, Middleware, DB, Documentation
│   └── 08-final_project/         # Mini REST API + Auth + Docker
│
├── 01_exercise_reflect/          # ✅ Materi tambahan: Reflect & struct tags
├── 02_exercise_api/              # ✅ Materi tambahan: Web server basics
└── 03_challenge_api/             # 🔄 On progress
```

> **Catatan**: Folder `01_exercise_reflect/` dan `02_exercise_api/` adalah materi tambahan/referensi. Bisa diabaikan jika fokus ke learning path utama (`00-go_learning_path/`). Folder ini akan dihapus setelah materi utama selesai.

## 🎯 Status Materi

### ✅ Tersedia (Available)
- **01-basic**: Go commands, variables, data types, control structures *(sebagian)*
- **Materi Tambahan**: Reflect & API basics di folder terpisah

### 🔄 Dalam Pengembangan (On Progress)
- 01-basic: Functions
- 02-aggregates: Arrays, Slices, Maps, Structs, Pointers
- 03-oop_style: Methods, Interfaces, Embedding, Polymorphism
- 04-advanced: Error handling, Concurrency, Channels, Testing
- 05-files_io: File operations, JSON, Environment variables
- 06-web_basics: HTTP server/client, Templates, Routing
- 07-rest_api: CRUD, Middleware, Database, Documentation
- 08-final_project: Complete REST API with Auth & Docker

## 🚀 Quick Start

### Prerequisites
```bash
# Install Go (1.21+)
go version

# Clone repository
git clone <repo-url>
cd golang_exercises
```

### Memulai Pembelajaran

#### Untuk Pemula (Mulai dari Nol)
```bash
# Mulai dari basics
cd 00-go_learning_path/01-basic/go_commands
cat README.md        # Baca penjelasan konsep
go run main.go       # Jalankan contoh kode

# Lanjut ke folder berikutnya
cd ../variables_and_constants
cat README.md
go run main.go
```

#### Menjalankan Materi Tambahan (Opsional)
```bash
# Reflect examples
cd 01_exercise_reflect/01_reflect
go run main.go

# API examples
cd 02_exercise_api/01_api
go run main.go
```

## 📖 Cara Belajar yang Efektif

### Langkah-langkah Belajar
1. **Baca README.md** di setiap folder untuk memahami konsep
2. **Jalankan main.go** untuk melihat output
3. **Modifikasi kode** - ubah nilai, tambah logic, experiment!
4. **Kerjakan exercises** yang ada di README
5. **Pindah ke folder berikutnya** setelah paham

### Urutan Pembelajaran
```
01-basic (Week 1-2)
  → go_commands
  → variables_and_constants
  → data_types_basic
  → control_structures
  → functions

02-aggregates (Week 2-3)
  → arrays
  → slices
  → maps
  → structs
  → pointers

03-oop_style (Week 3-4)
  → methods
  → interfaces
  → embedding
  → polymorphism

04-advanced (Week 4-5)
  → error_handling
  → concurrency_goroutines
  → channels
  → packages_modules
  → testing

05-files_io (Week 5-6)
  → reading_files
  → writing_files
  → json_handling
  → environment_variables

06-web_basics (Week 6-7)
  → http_server
  → http_client
  → templates
  → routing_basics

07-rest_api (Week 7-8)
  → crud_example
  → routing_with_mux
  → middleware
  → struct_validation
  → database_connection_postgres
  → env_config
  → error_handling_api
  → api_documentation

08-final_project (Week 8+)
  → mini_rest_api_project
  → authentication_jwt
  → modular_architecture
  → docker_setup
  → testing_api
```

## 🎓 Apa yang Akan Dipelajari

### Phase 1: Fundamentals
- Setup Go environment & tooling
- Variabel, tipe data, konstanta
- Control flow (if, switch, for)
- Functions & scope
- Data structures (array, slice, map, struct)
- Pointers & memory management

### Phase 2: Intermediate
- Object-oriented patterns di Go
- Error handling best practices
- Concurrency dengan goroutines
- Channel communication
- Package management
- Unit testing

### Phase 3: Web & API
- HTTP server dari scratch
- Routing & middleware
- JSON handling
- REST API design
- CRUD operations
- Database integration (PostgreSQL)

### Phase 4: Production Ready
- Authentication & authorization (JWT)
- API documentation (Swagger)
- Environment configuration
- Error contracts
- Logging & monitoring
- Containerization (Docker)
- Load testing

## 📝 Konvensi Repo

Setiap folder pembelajaran berisi:
- **`README.md`** - Penjelasan konsep, syntax, dan exercises
- **`main.go`** - Kode contoh lengkap dengan komentar detail
- **`go.mod`** - Module definition (jika butuh dependencies)

## 🛠️ Tools & Libraries

```bash
# Core
go 1.21+

# Web frameworks
github.com/gin-gonic/gin
github.com/labstack/echo/v4
github.com/gorilla/mux

# Database
github.com/lib/pq          # PostgreSQL driver
gorm.io/gorm               # ORM

# Testing & docs
github.com/stretchr/testify
github.com/swaggo/swag

# Tools
k6                         # Load testing
docker                     # Containerization
```

## 🤝 Contributing

Repo ini open untuk kontribusi! Jika kamu menemukan:
- ❌ Bug atau error dalam kode
- 📝 Penjelasan yang kurang jelas
- 💡 Ide untuk materi tambahan
- ✨ Perbaikan atau enhancement

Silakan:
1. Fork repo ini
2. Buat branch baru (`git checkout -b feature/improvement`)
3. Commit changes (`git commit -m 'Add: explanation for closures'`)
4. Push ke branch (`git push origin feature/improvement`)
5. Buat Pull Request

Atau cukup buat **issue** untuk diskusi atau request materi.

## 📌 Tips Belajar Go

- ✅ **Unused variables = compilation error** - Go sangat strict!
- ✅ **go fmt** - Format kode otomatis, gunakan sebelum commit
- ✅ **go mod tidy** - Jalankan setelah tambah/hapus dependencies
- ✅ **Read error messages** - Go error messages sangat informatif
- ✅ **Experiment** - Ubah kode, break things, learn from errors!
- ✅ **Practice** - Kerjakan exercises di setiap README

## 🔗 Resources Tambahan

### Official Documentation
- [Go Documentation](https://go.dev/doc/)
- [Go Tour](https://go.dev/tour/) - Interactive tutorial
- [Effective Go](https://go.dev/doc/effective_go) - Best practices

### Learning Resources
- [Go by Example](https://gobyexample.com/)
- [Go Playground](https://go.dev/play/) - Test code online
- [A Tour of Go](https://go.dev/tour/welcome/1)

### Community
- [Go Forum](https://forum.golangbridge.org/)
- [Reddit r/golang](https://reddit.com/r/golang)
- [Gophers Slack](https://gophers.slack.com/)

## 📧 Contact & Support

**Author**: Aisiya Qutwatunnada

## 📄 License

Repository ini free untuk digunakan untuk tujuan pembelajaran.

---

**Happy Learning Go! 🚀**

*"The best way to learn is by doing"*

---
**Dibuat oleh**: Aisiya Qutwatunnada

Last updated: October 2025 | Made with ❤️ for Go learners
