# Go Data Types & Variables

## Primitive Data Types

Go has several built-in primitive data types that form the foundation of the language.

### Integer Types

```go
// Signed integers
var a int     // Platform-dependent (32 or 64 bits)
var b int8    // 8-bit signed integer (-128 to 127)
var c int16   // 16-bit signed integer (-32768 to 32767)
var d int32   // 32-bit signed integer (-2147483648 to 2147483647)
var e int64   // 64-bit signed integer

// Unsigned integers
var f uint    // Platform-dependent (32 or 64 bits)
var g uint8   // 8-bit unsigned integer (0 to 255)
var h uint16  // 16-bit unsigned integer (0 to 65535)
var i uint32  // 32-bit unsigned integer (0 to 4294967295)
var j uint64  // 64-bit unsigned integer

// Special integer types
var k byte    // Alias for uint8
var l rune    // Alias for int32 (represents Unicode code points)
```

**Common Usage:**
- Use `int` for most general-purpose integer operations
- Use `int64` for large numbers or when you need specific size
- Use `byte` for binary data
- Use `rune` for Unicode characters

### Floating-Point Types

```go
var a float32 // 32-bit floating-point number
var b float64 // 64-bit floating-point number (default)

// Examples
var pi float64 = 3.14159
var temperature float32 = 98.6
```

**Important Notes:**
- `float64` is the default floating-point type
- Floating-point arithmetic is not exact (use `math` package for precise calculations)
- Use `float64` unless you specifically need `float32` for memory optimization

### String Type

```go
var message string = "Hello, Go!"

// Multi-line strings
var longMessage string = `This is a
multi-line string
using backticks`

// String concatenation
var firstName = "John"
var lastName = "Doe"
var fullName = firstName + " " + lastName

// String length
length := len(fullName) // Returns number of bytes, not characters

// Accessing characters (returns bytes)
firstChar := fullName[0] // Returns byte value, not character
```

**String Characteristics:**
- Strings are immutable (cannot be changed after creation)
- Strings are UTF-8 encoded by default
- Use `rune` for individual Unicode characters
- Use `strings` package for string manipulation

### Boolean Type

```go
var isActive bool = true
var isComplete bool = false

// Boolean operations
var result bool = true && false  // AND
var result2 bool = true || false // OR
var result3 bool = !true         // NOT
```

## Zero Values

Every type in Go has a default "zero value" that is assigned when a variable is declared without initialization:

```go
var a int     // 0
var b float64 // 0.0
var c string  // "" (empty string)
var d bool    // false
var e *int    // nil (for pointers)
var f []int   // nil (for slices)
var g map[string]int // nil (for maps)
var h chan int // nil (for channels)
var i interface{} // nil (for interfaces)
```

**Zero Value Examples:**
```go
package main

import "fmt"

func main() {
    var i int
    var f float64
    var s string
    var b bool
    
    fmt.Printf("int zero value: %d\n", i)        // 0
    fmt.Printf("float64 zero value: %f\n", f)    // 0.000000
    fmt.Printf("string zero value: '%s'\n", s)   // ''
    fmt.Printf("bool zero value: %t\n", b)       // false
}
```

## Variable Declaration

### 1. Explicit Declaration

```go
// Single variable
var name string = "John"

// Multiple variables of same type
var age, height int = 25, 180

// Multiple variables of different types
var (
    firstName string = "John"
    lastName  string = "Doe"
    age       int    = 25
    isActive  bool   = true
)
```

### 2. Type Inference

```go
// Go infers the type from the value
var name = "John"        // string
var age = 25            // int
var height = 180.5      // float64
var isActive = true     // bool

// Multiple variables with type inference
var firstName, lastName = "John", "Doe"
var age, height = 25, 180.5
```

### 3. Short Variable Declaration (`:=`)

```go
// Short declaration (most common in functions)
name := "John"
age := 25
height := 180.5
isActive := true

// Multiple variables
firstName, lastName := "John", "Doe"
age, height := 25, 180.5

// Reassignment (must use =, not :=)
age = 26
```

**Short Declaration Rules:**
- Can only be used inside functions
- At least one variable must be new
- Cannot be used for package-level variables

### 4. Blank Identifier (`_`)

```go
// Ignoring return values
_, err := someFunction()
if err != nil {
    // Handle error
}

// Ignoring specific values in multiple assignment
name, _, age := getUserInfo() // Ignore middle value
```

## Constants

Constants are values that cannot be changed after declaration.

### Basic Constants

```go
const Pi = 3.14159
const MaxRetries = 3
const AppName = "MyApp"

// Multiple constants
const (
    StatusOK    = 200
    StatusError = 500
    Timeout     = 30
)
```

### Typed vs Untyped Constants

```go
// Untyped constants (more flexible)
const Pi = 3.14159
const MaxRetries = 3

// Typed constants
const Pi float64 = 3.14159
const MaxRetries int = 3
```

### Constant Expressions

```go
const (
    // Arithmetic
    Sum = 1 + 2
    Product = 3 * 4
    
    // String concatenation
    Greeting = "Hello" + " " + "World"
    
    // Bitwise operations
    Flag1 = 1 << 0  // 1
    Flag2 = 1 << 1  // 2
    Flag3 = 1 << 2  // 4
)

// Using iota for enumerated constants
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)

// iota with expressions
const (
    KB = 1 << (10 * iota)  // 1 << (10 * 0) = 1
    MB                      // 1 << (10 * 1) = 1024
    GB                      // 1 << (10 * 2) = 1048576
    TB                      // 1 << (10 * 3) = 1073741824
)
```

## Type Conversion

Go requires explicit type conversion (no implicit conversion).

```go
// Integer conversions
var i int = 42
var f float64 = float64(i)
var u uint = uint(i)

// Float to integer (truncates decimal part)
var f2 float64 = 3.14
var i2 int = int(f2) // 3

// String conversions
var num int = 42
var str string = string(num) // Converts to Unicode character
var str2 string = fmt.Sprintf("%d", num) // "42"

// Using strconv package for proper string conversions
import "strconv"

str3, _ := strconv.Atoi("42")        // string to int
str4 := strconv.Itoa(42)             // int to string
str5, _ := strconv.ParseFloat("3.14", 64) // string to float64
```

## Type Aliases and Custom Types

### Type Aliases (Go 1.9+)

```go
type MyInt = int // Alias - same type, different name
var x MyInt = 42
var y int = x // No conversion needed
```

### Custom Types

```go
type MyInt int // New type based on int
var x MyInt = 42
var y int = int(x) // Conversion required

// Custom types with methods
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
    return Celsius((f - 32) * 5 / 9)
}
```

## Practical Examples

### Variable Scoping

```go
package main

import "fmt"

// Package-level variables
var globalVar = "I'm global"

func main() {
    // Function-level variables
    localVar := "I'm local"
    
    {
        // Block-level variables
        blockVar := "I'm in a block"
        fmt.Println(blockVar)
    }
    // blockVar is not accessible here
    
    fmt.Println(localVar)
    fmt.Println(globalVar)
}
```

### Working with Different Types

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Integer operations
    var a, b int = 10, 3
    fmt.Printf("Sum: %d, Difference: %d, Product: %d, Quotient: %d\n",
        a+b, a-b, a*b, a/b)
    
    // Floating-point operations
    var x, y float64 = 10.5, 3.2
    fmt.Printf("Sum: %.2f, Product: %.2f\n", x+y, x*y)
    
    // String operations
    firstName := "John"
    lastName := "Doe"
    fullName := firstName + " " + lastName
    fmt.Printf("Full name: %s, Length: %d\n", fullName, len(fullName))
    
    // Boolean logic
    isAdult := true
    hasLicense := false
    canDrive := isAdult && hasLicense
    fmt.Printf("Can drive: %t\n", canDrive)
    
    // Type conversion
    ageStr := "25"
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        fmt.Println("Error converting string to int")
    } else {
        fmt.Printf("Age: %d\n", age)
    }
}
```

### Constants in Practice

```go
package main

import "fmt"

const (
    // Application constants
    AppName    = "MyGoApp"
    Version    = "1.0.0"
    MaxRetries = 3
    
    // HTTP status codes
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
    
    // Time constants
    DefaultTimeout = 30
    MaxTimeout     = 300
)

// Using iota for flags
const (
    FlagRead = 1 << iota  // 1
    FlagWrite             // 2
    FlagExecute           // 4
)

func main() {
    fmt.Printf("App: %s v%s\n", AppName, Version)
    fmt.Printf("Max retries: %d\n", MaxRetries)
    
    // Using flags
    permissions := FlagRead | FlagWrite
    fmt.Printf("Permissions: %d\n", permissions)
    
    if permissions&FlagRead != 0 {
        fmt.Println("Read permission granted")
    }
}
```

This comprehensive guide covers all the essential concepts of data types and variables in Go. Practice with these examples to become comfortable with Go's type system! 