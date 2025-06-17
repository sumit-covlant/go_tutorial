# Go Functions

## Overview

Functions are the building blocks of Go programs. They allow you to organize code into reusable, named blocks that can be called with different arguments.

## Basic Function Declaration

### Function Syntax

```go
func functionName(parameters) returnType {
    // function body
}
```

**Example:**
```go
func greet(name string) string {
    return "Hello, " + name + "!"
}
```

### Function with No Parameters

```go
func sayHello() {
    fmt.Println("Hello, World!")
}
```

### Function with No Return Value

```go
func printMessage(message string) {
    fmt.Println(message)
}
```

### Function with Multiple Parameters

```go
func add(a, b int) int {
    return a + b
}

func greetPerson(firstName, lastName string) string {
    return "Hello, " + firstName + " " + lastName + "!"
}
```

## Return Values

### Single Return Value

```go
func square(x int) int {
    return x * x
}

func getGreeting() string {
    return "Hello, World!"
}
```

### Multiple Return Values

Go functions can return multiple values, which is commonly used for error handling:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

**Usage:**
```go
result, err := divide(10, 2)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Result: %d\n", result)
}
```

### Multiple Return Values of Same Type

```go
func minMax(numbers []int) (int, int) {
    if len(numbers) == 0 {
        return 0, 0
    }
    
    min := numbers[0]
    max := numbers[0]
    
    for _, num := range numbers {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    
    return min, max
}
```

### Ignoring Return Values

Use the blank identifier `_` to ignore return values:

```go
// Ignore error
result, _ := divide(10, 2)

// Ignore first return value
_, err := divide(10, 0)

// Ignore all return values
divide(10, 2) // This will cause a compilation error if you don't use the values
```

## Named Return Values

Go allows you to name return values, which can make code more readable:

```go
func divideAndRemainder(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return // Named return values
}
```

**Benefits:**
- Self-documenting code
- Can be used in defer statements
- Clearer function signatures

**Example with error handling:**
```go
func getUserInfo(userID int) (name string, age int, err error) {
    if userID <= 0 {
        err = fmt.Errorf("invalid user ID: %d", userID)
        return
    }
    
    // Simulate getting user data
    name = "John Doe"
    age = 25
    return
}
```

## Variadic Functions

Variadic functions can accept a variable number of arguments:

### Basic Variadic Function

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

**Usage:**
```go
fmt.Println(sum(1, 2, 3))        // 6
fmt.Println(sum(1, 2, 3, 4, 5))  // 15
fmt.Println(sum())               // 0
```

### Variadic Function with Regular Parameters

```go
func join(separator string, strings ...string) string {
    if len(strings) == 0 {
        return ""
    }
    
    result := strings[0]
    for _, s := range strings[1:] {
        result += separator + s
    }
    return result
}
```

**Usage:**
```go
fmt.Println(join(", ", "apple", "banana", "cherry")) // apple, banana, cherry
```

### Passing Slice to Variadic Function

Use the `...` operator to pass a slice to a variadic function:

```go
numbers := []int{1, 2, 3, 4, 5}
total := sum(numbers...)
fmt.Printf("Sum: %d\n", total)
```

## Function Types

Functions are first-class values in Go, meaning they can be assigned to variables and passed as arguments:

### Function as Variable

```go
var operation func(int, int) int

operation = add
result := operation(5, 3) // 8

operation = func(a, b int) int {
    return a * b
}
result = operation(5, 3) // 15
```

### Function as Parameter

```go
func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    result := applyOperation(10, 5, add)
    fmt.Printf("Result: %d\n", result)
}
```

### Function as Return Value

```go
func getOperation(operationType string) func(int, int) int {
    switch operationType {
    case "add":
        return func(a, b int) int { return a + b }
    case "multiply":
        return func(a, b int) int { return a * b }
    default:
        return func(a, b int) int { return 0 }
    }
}
```

## Anonymous Functions (Closures)

### Basic Anonymous Function

```go
func main() {
    greet := func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }
    
    greet("Alice")
}
```

### Closure with Captured Variables

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
}
```

### Immediately Invoked Function Expression (IIFE)

```go
func main() {
    result := func(a, b int) int {
        return a + b
    }(5, 3)
    
    fmt.Printf("Result: %d\n", result) // 8
}
```

## Defer Statement

The `defer` statement schedules a function call to be run before the function returns:

### Basic Defer

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Will be called when function exits
    
    // Process file...
    return nil
}
```

### Multiple Defer Statements

Defer statements are executed in LIFO (Last In, First Out) order:

```go
func main() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    
    fmt.Println("Main function")
}
// Output:
// Main function
// Third defer
// Second defer
// First defer
```

### Defer with Arguments

Arguments to deferred functions are evaluated when the defer statement is executed:

```go
func main() {
    i := 1
    defer fmt.Printf("Deferred: %d\n", i)
    
    i = 2
    fmt.Printf("Current: %d\n", i)
}
// Output:
// Current: 2
// Deferred: 1
```

### Defer with Named Return Values

```go
func processData() (result string, err error) {
    defer func() {
        if err != nil {
            result = "" // Can modify named return values
        }
    }()
    
    // Process data...
    return "success", nil
}
```

## Function Recursion

### Basic Recursion

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

### Recursion with Memoization

```go
var memo = make(map[int]int)

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    
    if result, exists := memo[n]; exists {
        return result
    }
    
    memo[n] = fibonacci(n-1) + fibonacci(n-2)
    return memo[n]
}
```

## Function Overloading and Default Parameters

Go doesn't support function overloading or default parameters, but you can achieve similar functionality:

### Using Optional Parameters Pattern

```go
type Config struct {
    Timeout time.Duration
    Retries int
    Debug   bool
}

func NewConfig() *Config {
    return &Config{
        Timeout: 30 * time.Second,
        Retries: 3,
        Debug:   false,
    }
}

func processWithConfig(data string, config *Config) {
    if config == nil {
        config = NewConfig()
    }
    // Process with config...
}
```

### Using Variadic Functions for Optional Parameters

```go
func connect(host string, options ...string) {
    port := "8080" // default
    protocol := "http" // default
    
    for i := 0; i < len(options); i++ {
        switch options[i] {
        case "https":
            protocol = "https"
        case "8080", "443", "3000":
            port = options[i]
        }
    }
    
    fmt.Printf("Connecting to %s://%s:%s\n", protocol, host, port)
}
```

## Error Handling Patterns

### Return Error Pattern

```go
func readConfig(filename string) (string, error) {
    if filename == "" {
        return "", fmt.Errorf("filename cannot be empty")
    }
    
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("failed to read file: %w", err)
    }
    
    return string(data), nil
}
```

### Error Wrapping

```go
func processUser(userID int) error {
    user, err := getUser(userID)
    if err != nil {
        return fmt.Errorf("failed to get user %d: %w", userID, err)
    }
    
    err = validateUser(user)
    if err != nil {
        return fmt.Errorf("user %d validation failed: %w", userID, err)
    }
    
    return nil
}
```

## Best Practices

### 1. Keep Functions Small and Focused

```go
// Good - single responsibility
func validateEmail(email string) error {
    if email == "" {
        return fmt.Errorf("email cannot be empty")
    }
    if !strings.Contains(email, "@") {
        return fmt.Errorf("invalid email format")
    }
    return nil
}

// Avoid - multiple responsibilities
func processUser(user User) error {
    // Validation
    if user.Email == "" {
        return fmt.Errorf("email cannot be empty")
    }
    // Database operations
    err := db.Save(user)
    if err != nil {
        return err
    }
    // Email notification
    return sendWelcomeEmail(user.Email)
}
```

### 2. Use Meaningful Function Names

```go
// Good
func calculateTotalPrice(items []Item) float64
func validateUserInput(input string) error
func sendEmailNotification(to, subject, body string) error

// Avoid
func calc(items []Item) float64
func validate(input string) error
func send(to, subject, body string) error
```

### 3. Return Errors, Don't Panic

```go
// Good
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Avoid
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

### 4. Use Named Return Values Appropriately

```go
// Good - for simple functions
func minMax(numbers []int) (min, max int) {
    if len(numbers) == 0 {
        return
    }
    min, max = numbers[0], numbers[0]
    for _, n := range numbers[1:] {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    return
}

// Avoid - for complex functions
func processData(data []byte) (result string, processed bool, err error) {
    // Complex logic with multiple return points
    if len(data) == 0 {
        return "", false, fmt.Errorf("empty data")
    }
    // More complex logic...
    return "processed", true, nil
}
```

### 5. Use Defer for Cleanup

```go
// Good
func readFile(filename string) (string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()
    
    data, err := io.ReadAll(file)
    if err != nil {
        return "", err
    }
    
    return string(data), nil
}
```

### 6. Document Public Functions

```go
// CalculateArea calculates the area of a rectangle
// given its width and height.
func CalculateArea(width, height float64) float64 {
    return width * height
}
```

This comprehensive guide covers all the essential aspects of functions in Go. Practice with these patterns to become proficient in writing clean, maintainable Go code! 