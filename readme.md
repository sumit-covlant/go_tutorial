# Go Learning Roadmap

## 🟢 Beginner Level: Language Basics

### 1. Introduction to Go

* History, use cases, and advantages
* Installing Go and setting up the environment

### 2. Basic Syntax

* Hello World
* Code structure (`main`, `package`, `import`)
* Comments

### 3. Data Types & Variables

* Primitive types: `int`, `float`, `string`, `bool`
* Zero values
* Constants (`const`)
* Type inference with `:=`

### 4. Control Structures

* `if`, `else`
* `switch`
* `for` loops (no `while` in Go)
* `break`, `continue`

### 5. Functions

* Function declaration
* Multiple return values
* Named return values
* Variadic functions
* `defer`

## 🟡 Intermediate Level: Core Concepts

### 6. Pointers

* Value vs reference
* Working with memory addresses

### 7. Structs and Methods

* Creating structs
* Adding methods to structs
* Struct embedding (composition)

### 8. Arrays, Slices, and Maps

* Fixed arrays
* Slices (dynamic arrays)
* Maps (hash tables/dictionaries)

### 9. Packages & Modules

* Creating and importing packages
* Go Modules (`go.mod`, `go.sum`)

### 10. Interfaces

* Interface types
* Implementing interfaces
* Empty interface (`interface{}`)

### 11. Error Handling

* `error` type
* Custom error types
* `errors` and `fmt.Errorf`
* `panic` and `recover`

## 🟠 Advanced Topics

### 12. Concurrency

* Goroutines
* Channels
* Buffered vs unbuffered channels
* Select statement
* Mutexes and WaitGroups (`sync` package)

### 13. File Handling & IO

* Reading/writing files
* Working with `io`, `os`, `bufio`

### 14. Testing

* Writing unit tests with `testing` package
* Benchmarking
* Table-driven tests
* Mocks (with `testify` or `gomock`)

### 15. Standard Library Deep Dive

* `fmt`, `strings`, `math`, `time`, `strconv`, etc.
* JSON handling (`encoding/json`)

## 🔵 Project-Level & Ecosystem Skills

### 16. Web Development with Go

* Using `net/http`
* Routers (e.g., `gorilla/mux`, `chi`, `gin`)
* Middleware
* Templates and HTML rendering
* REST APIs

### 17. Database Interaction

* SQL with `database/sql` + `pq`/`mysql`
* ORMs like `GORM`, `sqlx`

### 18. Go Tooling

* `go run`, `go build`, `go test`, `go vet`, `go fmt`, `go mod tidy`
* Linters (`golangci-lint`)

### 19. Deployment

* Cross-compiling
* Dockerizing Go apps
* Building CLI apps

### 20. Common Design Patterns

* Factory, Singleton, Strategy (in Go idioms)
* Dependency injection patterns in Go

### 21. Generics (Go 1.18+)

## 🧠 Optional Deep Dives

* Reflection (`reflect` package)
* Context API for managing goroutine lifecycle
* Performance profiling (`pprof`)
* Go Assembly (for low-level enthusiasts)
