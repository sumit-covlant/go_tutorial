# Go Error Handling

## Overview

Error handling in Go is explicit and designed to be simple yet powerful. Unlike exceptions in other languages, Go uses return values to indicate errors, making error handling a first-class concern in the language.

## The error Interface

### Basic Error Interface

The `error` interface is the foundation of error handling in Go:

```go
type error interface {
    Error() string
}
```

### Simple Error Usage

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("nonexistent.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Use file...
}
```

## Creating Errors

### 1. Using errors.New()

```go
import "errors"

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Result: %d\n", result)
}
```

### 2. Using fmt.Errorf()

```go
import "fmt"

func validateAge(age int) error {
    if age < 0 {
        return fmt.Errorf("age cannot be negative, got %d", age)
    }
    if age > 150 {
        return fmt.Errorf("age cannot exceed 150, got %d", age)
    }
    return nil
}

func main() {
    err := validateAge(-5)
    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
    }
}
```

### 3. Custom Error Types

```go
type ValidationError struct {
    Field   string
    Message string
    Value   interface{}
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s (value: %v)", 
        e.Field, e.Message, e.Value)
}

func validateUser(name string, age int) error {
    if name == "" {
        return ValidationError{
            Field:   "name",
            Message: "cannot be empty",
            Value:   name,
        }
    }
    
    if age < 0 {
        return ValidationError{
            Field:   "age",
            Message: "cannot be negative",
            Value:   age,
        }
    }
    
    return nil
}
```

## Error Handling Patterns

### 1. Early Return Pattern

```go
func processUser(name string, age int) error {
    // Validate name
    if name == "" {
        return errors.New("name cannot be empty")
    }
    
    // Validate age
    if age < 0 {
        return errors.New("age cannot be negative")
    }
    
    // Process user...
    fmt.Printf("Processing user: %s, age: %d\n", name, age)
    return nil
}
```

### 2. Error Wrapping (Go 1.13+)

```go
import (
    "errors"
    "fmt"
)

func readConfig() error {
    err := readFile("config.json")
    if err != nil {
        return fmt.Errorf("failed to read config: %w", err)
    }
    return nil
}

func main() {
    err := readConfig()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        
        // Unwrap the error
        if unwrapped := errors.Unwrap(err); unwrapped != nil {
            fmt.Printf("Unwrapped error: %v\n", unwrapped)
        }
    }
}
```

### 3. Error Checking with errors.Is()

```go
import (
    "errors"
    "io"
    "os"
)

var ErrNotFound = errors.New("not found")

func findUser(id string) (*User, error) {
    // Simulate not found
    if id == "nonexistent" {
        return nil, ErrNotFound
    }
    return &User{ID: id}, nil
}

func main() {
    user, err := findUser("nonexistent")
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            fmt.Println("User not found")
        } else {
            fmt.Printf("Unexpected error: %v\n", err)
        }
        return
    }
    
    fmt.Printf("Found user: %+v\n", user)
}
```

### 4. Type Assertions with errors.As()

```go
import "errors"

type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return ValidationError{Field: "age", Message: "cannot be negative"}
    }
    return nil
}

func main() {
    err := validateAge(-5)
    if err != nil {
        var valErr ValidationError
        if errors.As(err, &valErr) {
            fmt.Printf("Validation error on field %s: %s\n", 
                valErr.Field, valErr.Message)
        } else {
            fmt.Printf("Unexpected error: %v\n", err)
        }
    }
}
```

## Common Error Handling Patterns

### 1. Function with Multiple Return Values

```go
func divideAndModulo(a, b int) (quotient, remainder int, err error) {
    if b == 0 {
        return 0, 0, errors.New("division by zero")
    }
    return a / b, a % b, nil
}

func main() {
    q, r, err := divideAndModulo(17, 5)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("17 ÷ 5 = %d remainder %d\n", q, r)
}
```

### 2. Error with Context

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    defer file.Close()
    
    // Process file...
    return nil
}
```

### 3. Conditional Error Handling

```go
func processData(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }
    
    if len(data) > 1000 {
        return fmt.Errorf("data too large: %d bytes", len(data))
    }
    
    // Process data...
    return nil
}
```

## Error Types and Categories

### 1. Sentinel Errors

```go
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrInvalidInput = errors.New("invalid input")
)

func findUser(id string) (*User, error) {
    if id == "" {
        return nil, ErrInvalidInput
    }
    
    // Simulate not found
    if id == "nonexistent" {
        return nil, ErrNotFound
    }
    
    return &User{ID: id}, nil
}
```

### 2. Error Types

```go
type NotFoundError struct {
    Resource string
    ID       string
}

func (e NotFoundError) Error() string {
    return fmt.Sprintf("%s with id %s not found", e.Resource, e.ID)
}

type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

func findUser(id string) (*User, error) {
    if id == "" {
        return nil, ValidationError{Field: "id", Message: "cannot be empty"}
    }
    
    if id == "nonexistent" {
        return nil, NotFoundError{Resource: "user", ID: id}
    }
    
    return &User{ID: id}, nil
}
```

### 3. Wrapped Errors

```go
func processUser(id string) error {
    user, err := findUser(id)
    if err != nil {
        return fmt.Errorf("failed to process user %s: %w", id, err)
    }
    
    // Process user...
    return nil
}

func main() {
    err := processUser("nonexistent")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        
        // Check for specific error types
        var notFound NotFoundError
        if errors.As(err, &notFound) {
            fmt.Printf("Not found: %s\n", notFound.Error())
        }
    }
}
```

## Error Handling Best Practices

### 1. Always Check Errors

```go
// Good: Always check errors
file, err := os.Open("file.txt")
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}
defer file.Close()

// Bad: Ignoring errors
file, _ := os.Open("file.txt")  // Don't do this!
```

### 2. Return Errors, Don't Panic

```go
// Good: Return errors
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Bad: Panicking
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")  // Don't do this!
    }
    return a / b
}
```

### 3. Add Context to Errors

```go
// Good: Add context
func readConfig() error {
    err := readFile("config.json")
    if err != nil {
        return fmt.Errorf("failed to read config file: %w", err)
    }
    return nil
}

// Bad: No context
func readConfig() error {
    err := readFile("config.json")
    if err != nil {
        return err  // Loses context
    }
    return nil
}
```

### 4. Use Custom Error Types for Complex Errors

```go
type DatabaseError struct {
    Operation string
    Table     string
    Err       error
}

func (e DatabaseError) Error() string {
    return fmt.Sprintf("database error during %s on table %s: %v", 
        e.Operation, e.Table, e.Err)
}

func (e DatabaseError) Unwrap() error {
    return e.Err
}

func insertUser(user *User) error {
    err := db.Insert(user)
    if err != nil {
        return DatabaseError{
            Operation: "insert",
            Table:     "users",
            Err:       err,
        }
    }
    return nil
}
```

## Error Handling in Different Contexts

### 1. HTTP Handlers

```go
func handleGetUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "missing user id", http.StatusBadRequest)
        return
    }
    
    user, err := findUser(id)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            http.Error(w, "user not found", http.StatusNotFound)
            return
        }
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(user)
}
```

### 2. Database Operations

```go
func getUserByID(id string) (*User, error) {
    var user User
    err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Name, &user.Email)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrNotFound
        }
        return nil, fmt.Errorf("database error: %w", err)
    }
    
    return &user, nil
}
```

### 3. File Operations

```go
func readFile(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, fmt.Errorf("file %s does not exist", filename)
        }
        return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    defer file.Close()
    
    data, err := io.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
    }
    
    return data, nil
}
```

## Error Logging

### 1. Structured Error Logging

```go
import "log"

func processRequest(req *Request) error {
    err := validateRequest(req)
    if err != nil {
        log.Printf("Request validation failed: %v", err)
        return fmt.Errorf("invalid request: %w", err)
    }
    
    // Process request...
    return nil
}
```

### 2. Error with Stack Trace

```go
import (
    "fmt"
    "runtime"
)

func logError(err error) {
    _, file, line, _ := runtime.Caller(1)
    log.Printf("Error at %s:%d: %v", file, line, err)
}

func processData(data []byte) error {
    if len(data) == 0 {
        err := errors.New("empty data")
        logError(err)
        return err
    }
    
    // Process data...
    return nil
}
```

## Testing Error Handling

### 1. Testing Error Returns

```go
import "testing"

func TestDivide(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
        hasError bool
    }{
        {10, 2, 5, false},
        {10, 0, 0, true},
        {0, 5, 0, false},
    }
    
    for _, test := range tests {
        result, err := divide(test.a, test.b)
        
        if test.hasError {
            if err == nil {
                t.Errorf("Expected error for %d / %d", test.a, test.b)
            }
        } else {
            if err != nil {
                t.Errorf("Unexpected error for %d / %d: %v", test.a, test.b, err)
            }
            if result != test.expected {
                t.Errorf("Expected %d, got %d", test.expected, result)
            }
        }
    }
}
```

### 2. Testing Custom Error Types

```go
func TestValidationError(t *testing.T) {
    err := ValidationError{Field: "age", Message: "cannot be negative"}
    
    if err.Field != "age" {
        t.Errorf("Expected field 'age', got %s", err.Field)
    }
    
    if err.Message != "cannot be negative" {
        t.Errorf("Expected message 'cannot be negative', got %s", err.Message)
    }
    
    expectedMsg := "validation error on field age: cannot be negative"
    if err.Error() != expectedMsg {
        t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
    }
}
```

## Common Pitfalls

### 1. Ignoring Errors

```go
// Bad: Ignoring errors
file, _ := os.Open("file.txt")
defer file.Close()

// Good: Handle errors
file, err := os.Open("file.txt")
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}
defer file.Close()
```

### 2. Overly Generic Error Messages

```go
// Bad: Generic error
if err != nil {
    return errors.New("error occurred")
}

// Good: Specific error
if err != nil {
    return fmt.Errorf("failed to read configuration file: %w", err)
}
```

### 3. Not Using Error Wrapping

```go
// Bad: Losing context
func process() error {
    err := readFile("config.json")
    if err != nil {
        return err  // Loses context
    }
    return nil
}

// Good: Adding context
func process() error {
    err := readFile("config.json")
    if err != nil {
        return fmt.Errorf("failed to process configuration: %w", err)
    }
    return nil
}
```

## Summary

Error handling in Go provides:

- **Explicit Error Handling**: Errors are returned as values, not thrown
- **Simple Interface**: Just implement the `error` interface
- **Rich Context**: Add context with error wrapping
- **Type Safety**: Use custom error types for complex scenarios
- **Composability**: Combine and wrap errors as needed

Key points to remember:
1. Always check returned errors
2. Add context to errors when wrapping them
3. Use custom error types for complex error scenarios
4. Return errors, don't panic
5. Use `errors.Is()` and `errors.As()` for error checking
6. Test error handling thoroughly
7. Log errors appropriately

Understanding error handling is essential for writing robust, reliable Go code that gracefully handles failure scenarios. 