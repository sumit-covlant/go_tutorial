# Go Basic Syntax

## Hello World Program

Let's start with the classic "Hello, World!" program in Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Breaking Down the Hello World Program

1. **`package main`**: Every Go program starts with a package declaration
   - `main` is a special package that defines an executable program
   - Other packages are typically named after their directory

2. **`import "fmt"`**: Imports the `fmt` package for formatted I/O
   - `fmt` provides functions like `Println`, `Printf`, `Sprintf`, etc.
   - Multiple imports can be grouped: `import ("fmt"; "os")`

3. **`func main()`**: The entry point of the program
   - Every executable Go program must have a `main` function
   - It's called automatically when the program starts

4. **`fmt.Println("Hello, World!")`**: Prints text to the console
   - `Println` adds a newline at the end
   - Use `Print` if you don't want a newline

## Code Structure

### Package Declaration

```go
package main        // Executable program
package mypackage   // Library package
```

**Rules:**
- Must be the first line (after comments)
- Package name is usually the same as the directory name
- `main` package creates an executable
- Other packages create libraries

### Import Statements

```go
// Single import
import "fmt"

// Multiple imports (recommended style)
import (
    "fmt"
    "os"
    "strings"
)

// Import with alias
import myfmt "fmt"

// Import without using the package name
import _ "database/sql/driver"

// Import all exported names directly
import . "fmt"  // Now you can use Println() instead of fmt.Println()
```

**Import Rules:**
- Unused imports cause compilation errors
- Use `go fmt` to automatically organize imports
- Use `go mod tidy` to clean up unused dependencies

### Function Declaration

```go
// Basic function
func greet() {
    fmt.Println("Hello!")
}

// Function with parameters
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

### Main Function

```go
func main() {
    // Program starts here
    fmt.Println("Program started")
    
    // Your code here
    
    fmt.Println("Program finished")
}
```

**Main Function Rules:**
- Must be in the `main` package
- No parameters or return values
- Only one `main` function per program

## Comments

Go supports three types of comments:

### 1. Single-line Comments

```go
// This is a single-line comment
fmt.Println("Hello") // Comment after code
```

### 2. Multi-line Comments

```go
/*
This is a multi-line comment.
It can span multiple lines.
Useful for longer explanations.
*/
```

### 3. Documentation Comments

```go
// Package example provides example functions
package example

// Greet returns a greeting message
// It takes a name parameter and returns a string
func Greet(name string) string {
    return "Hello, " + name + "!"
}
```

**Documentation Comments Rules:**
- Start with `//` (not `/*`)
- Must be immediately before the declaration
- Used by `go doc` and `godoc`
- Should be complete sentences

## File Organization

### Basic File Structure

```go
// File: main.go
package main

import (
    "fmt"
    "os"
)

// Constants
const (
    Version = "1.0.0"
    Author  = "John Doe"
)

// Global variables
var (
    debug = false
    port  = 8080
)

// Main function
func main() {
    fmt.Printf("Version: %s\n", Version)
    fmt.Printf("Author: %s\n", Author)
    
    if debug {
        fmt.Println("Debug mode enabled")
    }
    
    fmt.Printf("Server starting on port %d\n", port)
}
```

### Multiple Files in Same Package

**File 1: main.go**
```go
package main

import "fmt"

func main() {
    fmt.Println("Starting application...")
    result := calculate(10, 5)
    fmt.Printf("Result: %d\n", result)
}
```

**File 2: math.go**
```go
package main

// calculate performs basic arithmetic
func calculate(a, b int) int {
    return a + b
}
```

## Naming Conventions

### Exported vs Unexported

```go
// Exported (public) - starts with uppercase
func PublicFunction() {
    // Can be used by other packages
}

// Unexported (private) - starts with lowercase
func privateFunction() {
    // Only available within the same package
}

// Exported constant
const MaxSize = 100

// Unexported constant
const defaultTimeout = 30
```

### Variable and Function Names

```go
// Use camelCase for variables and functions
var userName string
var maxRetries int

func getUserInfo() {
    // Function implementation
}

// Use PascalCase for exported names
func GetUserInfo() {
    // Exported function
}

// Use UPPER_CASE for constants (optional)
const MAX_CONNECTIONS = 100
const DefaultTimeout = 30
```

## Running Go Programs

### Basic Commands

```bash
# Run a Go program
go run main.go

# Run multiple files
go run main.go math.go

# Build an executable
go build main.go

# Build and run
go build main.go && ./main

# Format code
go fmt main.go

# Check for common errors
go vet main.go
```

### Project Structure Example

```
myproject/
├── main.go
├── go.mod
├── go.sum
├── pkg/
│   ├── math/
│   │   └── calculator.go
│   └── utils/
│       └── helpers.go
└── cmd/
    └── server/
        └── main.go
```

## Common Patterns

### Error Handling

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    // Process file...
    return nil
}
```

### Initialization

```go
func init() {
    // This function runs before main()
    fmt.Println("Initializing...")
}

func main() {
    fmt.Println("Main function")
}
```

### Multiple Return Values

```go
func divideAndRemainder(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return // Named return values
}
```

This covers the fundamental syntax concepts you need to start writing Go programs. Practice with these concepts before moving on to more advanced topics! 