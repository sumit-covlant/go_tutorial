# Go Packages & Modules

## Overview

Packages and modules are fundamental concepts in Go that enable code organization, reusability, and dependency management. Packages group related code together, while modules provide a way to manage dependencies and versioning.

## Packages

### What are Packages?

A package is a collection of Go source files in the same directory that are compiled together. Packages provide a way to organize and reuse code, and they control visibility of functions, types, and variables.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Package Declaration

Every Go source file must begin with a package declaration:

```go
package packagename
```

**Important**: The package name is the same as the last element of the import path.

```go
// File: math/calculator.go
package math

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return b - a
}
```

### Package Types

#### 1. Main Package

The `main` package is special - it defines a standalone executable program.

```go
package main

import "fmt"

func main() {
    fmt.Println("This is an executable program")
}
```

#### 2. Library Packages

All other packages are library packages that can be imported by other code.

```go
// File: utils/stringutils.go
package utils

import "strings"

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ToUpperCase(s string) string {
    return strings.ToUpper(s)
}
```

### Package Structure

A typical Go package structure:

```
myproject/
├── main.go
├── utils/
│   ├── stringutils.go
│   └── mathutils.go
├── models/
│   └── user.go
└── handlers/
    └── userhandler.go
```

### Package Naming Conventions

```go
// Good package names (short, clear, no underscores)
package math
package utils
package user
package http

// Avoid these
package string_utils  // Use stringutils
package my_package    // Use mypackage
package pkg           // Too short
package verylongpackagename  // Too long
```

### Package Visibility

Go uses capitalization to control visibility:

```go
package utils

// Exported (public) - starts with uppercase
func PublicFunction() string {
    return "This is public"
}

// Unexported (private) - starts with lowercase
func privateFunction() string {
    return "This is private"
}

// Exported variable
var PublicVariable = "public"

// Unexported variable
var privateVariable = "private"

// Exported type
type PublicStruct struct {
    PublicField   string  // Exported field
    privateField  int     // Unexported field
}
```

## Importing Packages

### Basic Import

```go
package main

import "fmt"
import "math"

func main() {
    fmt.Println("Hello")
    fmt.Printf("Pi: %.2f\n", math.Pi)
}
```

### Multiple Imports

```go
package main

import (
    "fmt"
    "math"
    "strings"
)

func main() {
    fmt.Println(strings.ToUpper("hello"))
    fmt.Printf("Square root of 16: %.2f\n", math.Sqrt(16))
}
```

### Import Aliases

```go
package main

import (
    "fmt"
    m "math"
    str "strings"
)

func main() {
    fmt.Println(str.ToUpper("hello"))
    fmt.Printf("Pi: %.2f\n", m.Pi)
}
```

### Dot Import

```go
package main

import . "fmt"

func main() {
    // Can use Println directly without fmt prefix
    Println("Hello, World!")
}
```

**Warning**: Dot imports are generally discouraged as they can cause naming conflicts.

### Blank Import

```go
package main

import (
    "fmt"
    _ "image/png"  // Import for side effects only
)

func main() {
    fmt.Println("PNG format registered")
}
```

### Import Paths

```go
// Standard library packages
import "fmt"
import "math"
import "strings"

// Third-party packages
import "github.com/gorilla/mux"
import "golang.org/x/crypto/bcrypt"

// Local packages (relative to module root)
import "myproject/utils"
import "myproject/models"
```

## Creating Your Own Packages

### Step 1: Create Package Directory

```bash
mkdir -p myproject/utils
mkdir -p myproject/models
```

### Step 2: Create Package Files

```go
// File: myproject/utils/stringutils.go
package utils

import "strings"

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ToUpperCase(s string) string {
    return strings.ToUpper(s)
}

func ToLowerCase(s string) string {
    return strings.ToLower(s)
}
```

```go
// File: myproject/models/user.go
package models

import "time"

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email string) *User {
    return &User{
        Name:      name,
        Email:     email,
        CreatedAt: time.Now(),
    }
}

func (u *User) GetFullName() string {
    return u.Name
}
```

### Step 3: Use Your Packages

```go
// File: myproject/main.go
package main

import (
    "fmt"
    "myproject/utils"
    "myproject/models"
)

func main() {
    // Use utils package
    reversed := utils.Reverse("hello")
    fmt.Printf("Reversed: %s\n", reversed)
    
    // Use models package
    user := models.NewUser("Alice", "alice@example.com")
    fmt.Printf("User: %+v\n", user)
}
```

## Modules

### What are Modules?

A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root. Modules provide dependency management and versioning.

### Creating a Module

```bash
# Create a new directory
mkdir myproject
cd myproject

# Initialize a module
go mod init myproject
```

This creates a `go.mod` file:

```go
module myproject

go 1.21
```

### Module Structure

```
myproject/
├── go.mod
├── go.sum
├── main.go
├── utils/
│   └── stringutils.go
├── models/
│   └── user.go
└── handlers/
    └── userhandler.go
```

### Adding Dependencies

```bash
# Add a dependency
go get github.com/gorilla/mux

# Add a specific version
go get github.com/gorilla/mux@v1.8.0

# Add multiple dependencies
go get github.com/gorilla/mux github.com/gorilla/sessions
```

The `go.mod` file will be updated:

```go
module myproject

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/gorilla/sessions v1.2.1
)
```

### Using Dependencies

```go
package main

import (
    "fmt"
    "net/http"
    
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    
    http.ListenAndServe(":8080", r)
}
```

### Module Commands

```bash
# Initialize a new module
go mod init myproject

# Add missing dependencies
go mod tidy

# Download dependencies
go mod download

# Verify dependencies
go mod verify

# Edit go.mod file
go mod edit

# List all dependencies
go list -m all

# Update dependencies
go get -u all
```

### Module Versioning

#### Semantic Versioning

Go modules use semantic versioning (semver):

- `v1.2.3` - Major.Minor.Patch
- `v1.2.3-pre` - Pre-release
- `v1.2.3+metadata` - Build metadata

#### Version Selection

```go
// go.mod
require (
    github.com/gorilla/mux v1.8.0  // Exact version
    github.com/gorilla/sessions v1.2.1  // Exact version
)
```

#### Version Constraints

```go
// go.mod
require (
    github.com/gorilla/mux v1.8.0  // Exact version
    github.com/gorilla/sessions v1.2.1  // Exact version
    golang.org/x/crypto v0.0.0-20210921155107-089bfa567519  // Pseudo-version
)
```

### Working with Modules

#### Publishing a Module

```bash
# Tag your release
git tag v1.0.0
git push origin v1.0.0

# Or create a release on GitHub/GitLab
```

#### Using Your Published Module

```go
package main

import (
    "fmt"
    "github.com/yourusername/yourmodule/utils"
)

func main() {
    result := utils.SomeFunction()
    fmt.Println(result)
}
```

## Package Documentation

### Package Comments

```go
// Package utils provides utility functions for string manipulation.
package utils

import "strings"

// Reverse returns the string s with its characters in reverse order.
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// ToUpperCase converts the string s to uppercase.
func ToUpperCase(s string) string {
    return strings.ToUpper(s)
}
```

### Generating Documentation

```bash
# Generate documentation
go doc

# Generate documentation for a specific package
go doc utils

# Generate documentation for a specific function
go doc utils.Reverse

# Generate HTML documentation
godoc -http=:6060
```

## Testing Packages

### Unit Tests

```go
// File: utils/stringutils_test.go
package utils

import "testing"

func TestReverse(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello", "olleh"},
        {"", ""},
        {"a", "a"},
        {"123", "321"},
    }
    
    for _, test := range tests {
        result := Reverse(test.input)
        if result != test.expected {
            t.Errorf("Reverse(%q) = %q, want %q", test.input, result, test.expected)
        }
    }
}

func TestToUpperCase(t *testing.T) {
    result := ToUpperCase("hello")
    expected := "HELLO"
    if result != expected {
        t.Errorf("ToUpperCase('hello') = %q, want %q", result, expected)
    }
}
```

### Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests for a specific package
go test ./utils

# Run tests with coverage
go test -cover

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Benchmark Tests

```go
// File: utils/stringutils_test.go
func BenchmarkReverse(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Reverse("hello world")
    }
}
```

```bash
# Run benchmarks
go test -bench=.

# Run benchmarks with memory allocation info
go test -bench=. -benchmem
```

## Package Best Practices

### 1. Package Design

```go
// Good: Single responsibility
package math
package strings
package time

// Avoid: Multiple responsibilities
package utils  // Too generic
```

### 2. Package Naming

```go
// Good: Clear, descriptive names
package user
package auth
package database

// Avoid: Generic names
package helper
package common
package util
```

### 3. Package Organization

```
myproject/
├── cmd/           # Main applications
│   ├── server/
│   └── client/
├── internal/      # Private application code
│   ├── auth/
│   └── database/
├── pkg/           # Public library code
│   ├── utils/
│   └── models/
└── api/           # API definitions
    └── v1/
```

### 4. Package Dependencies

```go
// Good: Minimal dependencies
package utils

import "strings"  // Only what you need

// Avoid: Unnecessary dependencies
package utils

import (
    "fmt"
    "math"
    "strings"
    "time"
    // ... many more
)
```

### 5. Package Documentation

```go
// Package math provides mathematical utilities.
//
// This package includes functions for basic arithmetic operations,
// mathematical constants, and common mathematical functions.
package math

// Add returns the sum of two integers.
//
// Example:
//     result := Add(5, 3)  // result == 8
func Add(a, b int) int {
    return a + b
}
```

### 6. Package Testing

```go
// Always include tests with your packages
package math

import "testing"

func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Error("Add(2, 3) should equal 5")
    }
}
```

## Common Patterns

### 1. Package Initialization

```go
package database

import "database/sql"

var db *sql.DB

func init() {
    // Package initialization code
    var err error
    db, err = sql.Open("postgres", "connection_string")
    if err != nil {
        panic(err)
    }
}

func GetDB() *sql.DB {
    return db
}
```

### 2. Package Configuration

```go
package config

import "os"

type Config struct {
    DatabaseURL string
    Port        string
    Environment string
}

var defaultConfig = Config{
    DatabaseURL: "localhost:5432",
    Port:        "8080",
    Environment: "development",
}

func Load() *Config {
    config := defaultConfig
    
    if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
        config.DatabaseURL = dbURL
    }
    
    if port := os.Getenv("PORT"); port != "" {
        config.Port = port
    }
    
    if env := os.Getenv("ENVIRONMENT"); env != "" {
        config.Environment = env
    }
    
    return &config
}
```

### 3. Package Factories

```go
package logger

import "log"

type Logger struct {
    level string
}

func NewLogger(level string) *Logger {
    return &Logger{level: level}
}

func (l *Logger) Info(msg string) {
    if l.level == "info" || l.level == "debug" {
        log.Printf("[INFO] %s", msg)
    }
}

func (l *Logger) Error(msg string) {
    log.Printf("[ERROR] %s", msg)
}
```

## Summary

Packages and modules in Go provide:

- **Code Organization**: Group related functionality together
- **Reusability**: Share code across different projects
- **Dependency Management**: Handle external dependencies
- **Versioning**: Manage different versions of dependencies
- **Testing**: Organize and run tests effectively

Key points to remember:
1. Every Go file must have a package declaration
2. Use capitalization to control visibility
3. Import packages using their full path
4. Use modules for dependency management
5. Follow naming conventions and best practices
6. Include comprehensive documentation and tests
7. Keep packages focused and minimal

Understanding packages and modules is essential for writing maintainable, reusable Go code and managing project dependencies effectively. 