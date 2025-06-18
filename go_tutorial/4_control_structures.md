# Go Control Structures

## Overview

Go has a minimal set of control structures compared to other languages. The main control structures are:
- `if` and `if else` statements with `else` statements
- `switch` statements
- `for` loops (no `while` or `do-while` in Go)
- `break` and `continue` statements

## If Statements

### Basic If Statement

```go
if condition {
    // code to execute if condition is true
}
```

**Example:**
```go
age := 18
if age >= 18 {
    fmt.Println("You are an adult")
}
```

### If-Else Statement

```go
if condition {
    // code to execute if condition is true
} else {
    // code to execute if condition is false
}
```

**Example:**
```go
age := 16
if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are a minor")
}
```

### If-Else If-Else Chain

```go
if condition1 {
    // code for condition1
} else if condition2 {
    // code for condition2
} else if condition3 {
    // code for condition3
} else {
    // default code
}
```

**Example:**
```go
score := 85
if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else if score >= 70 {
    fmt.Println("Grade: C")
} else if score >= 60 {
    fmt.Println("Grade: D")
} else {
    fmt.Println("Grade: F")
}
```

### If with Initialization Statement

Go allows you to initialize a variable in the if statement:

```go
if initialization; condition {
    // code to execute if condition is true
}
```

**Example:**
```go
if age := getUserAge(); age >= 18 {
    fmt.Printf("User is %d years old and can vote\n", age)
} else {
    fmt.Printf("User is %d years old and cannot vote\n", age)
}
```

**Multiple initialization:**
```go
if user, err := getUser(); err == nil {
    fmt.Printf("Welcome, %s!\n", user.Name)
} else {
    fmt.Printf("Error: %v\n", err)
}
```

### If with Short Variable Declaration

```go
if value := someFunction(); value > 0 {
    fmt.Printf("Positive value: %d\n", value)
}
```

## Switch Statements

### Basic Switch Statement

```go
switch expression {
case value1:
    // code for value1
case value2:
    // code for value2
default:
    // default code
}
```

**Example:**
```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("TGIF!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Midweek")
}
```

### Switch with Expression

```go
switch {
case condition1:
    // code for condition1
case condition2:
    // code for condition2
default:
    // default code
}
```

**Example:**
```go
age := 25
switch {
case age < 13:
    fmt.Println("Child")
case age < 20:
    fmt.Println("Teenager")
case age < 65:
    fmt.Println("Adult")
default:
    fmt.Println("Senior")
}
```

### Switch with Initialization

```go
switch initialization; expression {
case value1:
    // code for value1
case value2:
    // code for value2
default:
    // default code
}
```

**Example:**
```go
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("macOS")
case "linux":
    fmt.Println("Linux")
case "windows":
    fmt.Println("Windows")
default:
    fmt.Printf("Other: %s\n", os)
}
```

### Switch with Fallthrough

By default, Go switch statements don't fall through. Use `fallthrough` to continue to the next case:

```go
switch n := 2; n {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}
// Output: Two, Three
```

### Type Switch

Switch on the type of an interface:

```go
var value interface{} = "hello"
switch v := value.(type) {
case string:
    fmt.Printf("String: %s\n", v)
case int:
    fmt.Printf("Integer: %d\n", v)
case bool:
    fmt.Printf("Boolean: %t\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

## For Loops

Go has only one loop construct: the `for` loop. It can be used in several ways.

### Traditional For Loop

```go
for initialization; condition; increment {
    // loop body
}
```

**Example:**
```go
for i := 0; i < 5; i++ {
    fmt.Printf("%d ", i)
}
// Output: 0 1 2 3 4
```

### While-Style Loop

```go
for condition {
    // loop body
}
```

**Example:**
```go
count := 0
for count < 5 {
    fmt.Printf("%d ", count)
    count++
}
// Output: 0 1 2 3 4
```

### Infinite Loop

```go
for {
    // loop body
    if condition {
        break
    }
}
```

**Example:**
```go
count := 0
for {
    fmt.Printf("%d ", count)
    count++
    if count >= 5 {
        break
    }
}
// Output: 0 1 2 3 4
```

### For-Range Loop

Iterate over arrays, slices, maps, strings, and channels:

```go
for index, value := range collection {
    // loop body
}
```

**Array/Slice:**
```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

**Map:**
```go
person := map[string]string{
    "name": "John",
    "city": "New York",
}
for key, value := range person {
    fmt.Printf("Key: %s, Value: %s\n", key, value)
}
```

**String:**
```go
message := "Hello"
for index, char := range message {
    fmt.Printf("Index: %d, Character: %c\n", index, char)
}
```

**Ignore index or value:**
```go
numbers := []int{1, 2, 3, 4, 5}

// Ignore index
for _, value := range numbers {
    fmt.Printf("Value: %d\n", value)
}

// Ignore value
for index := range numbers {
    fmt.Printf("Index: %d\n", index)
}
```

## Break and Continue

### Break Statement

Exits the innermost loop or switch statement:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Exit loop when i equals 5
    }
    fmt.Printf("%d ", i)
}
// Output: 0 1 2 3 4
```

**Break with labels:**
```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer // Break out of both loops
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
// Output: (0,0) (0,1) (0,2) (1,0)
```

### Continue Statement

Skips the rest of the current iteration and continues with the next:

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Printf("%d ", i)
}
// Output: 1 3 5 7 9
```

**Continue with labels:**
```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            continue outer // Skip to next iteration of outer loop
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
```

## Nested Control Structures

### Nested If Statements

```go
age := 25
income := 50000

if age >= 18 {
    if income >= 30000 {
        fmt.Println("Eligible for loan")
    } else {
        fmt.Println("Income too low")
    }
} else {
    fmt.Println("Too young")
}
```

### Nested Loops

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("(%d,%d) ", i, j)
    }
    fmt.Println()
}
// Output:
// (1,1) (1,2) (1,3)
// (2,1) (2,2) (2,3)
// (3,1) (3,2) (3,3)
```

### Switch Inside Loop

```go
for i := 1; i <= 5; i++ {
    switch i {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    default:
        fmt.Printf("Number: %d\n", i)
    }
}
```

## Common Patterns

### Input Validation

```go
func validateAge(age int) string {
    if age < 0 {
        return "Invalid age"
    } else if age < 13 {
        return "Child"
    } else if age < 20 {
        return "Teenager"
    } else if age < 65 {
        return "Adult"
    } else {
        return "Senior"
    }
}
```

### Menu System

```go
func showMenu() {
    fmt.Println("1. Add user")
    fmt.Println("2. Delete user")
    fmt.Println("3. List users")
    fmt.Println("4. Exit")
}

func processChoice(choice int) {
    switch choice {
    case 1:
        fmt.Println("Adding user...")
    case 2:
        fmt.Println("Deleting user...")
    case 3:
        fmt.Println("Listing users...")
    case 4:
        fmt.Println("Exiting...")
    default:
        fmt.Println("Invalid choice")
    }
}
```

### Error Handling

```go
func processData(data string) error {
    if data == "" {
        return fmt.Errorf("data cannot be empty")
    }
    
    if len(data) < 5 {
        return fmt.Errorf("data too short")
    }
    
    // Process data...
    return nil
}
```

### Loop with Early Exit

```go
func findNumber(numbers []int, target int) (int, bool) {
    for i, num := range numbers {
        if num == target {
            return i, true
        }
    }
    return -1, false
}
```

## Best Practices

### 1. Use Switch for Multiple Conditions

Instead of long if-else chains, use switch:

```go
// Good
switch {
case age < 13:
    return "Child"
case age < 20:
    return "Teenager"
case age < 65:
    return "Adult"
default:
    return "Senior"
}

// Avoid long if-else chains
if age < 13 {
    return "Child"
} else if age < 20 {
    return "Teenager"
} else if age < 65 {
    return "Adult"
} else {
    return "Senior"
}
```

### 2. Use For-Range for Iteration

```go
// Good
for index, value := range items {
    fmt.Printf("Item %d: %v\n", index, value)
}

// Avoid
for i := 0; i < len(items); i++ {
    fmt.Printf("Item %d: %v\n", i, items[i])
}
```

### 3. Use Break for Early Exit

```go
// Good
for _, item := range items {
    if item == target {
        found = true
        break
    }
}

// Avoid
found = false
for _, item := range items {
    if item == target {
        found = true
    }
}
```

### 4. Use Initialization in Control Structures

```go
// Good
if value, err := someFunction(); err == nil {
    // Use value
}

// Avoid
value, err := someFunction()
if err == nil {
    // Use value
}
```

### 5. Keep Loops Simple

```go
// Good - simple loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Avoid - complex nested logic
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        for j := 0; j < i; j++ {
            if j%3 == 0 {
                fmt.Println(i, j)
            }
        }
    }
}
```

This comprehensive guide covers all the essential control structures in Go. Practice with these patterns to become comfortable with Go's control flow! 