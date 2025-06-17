package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

// Package-level constants using iota
const (
	// HTTP status codes
	StatusOK       = 200
	StatusNotFound = 404
	StatusError    = 500

	// File permissions using iota
	FlagRead    = 1 << iota // 1
	FlagWrite               // 2
	FlagExecute             // 4
)

// Size constants (using explicit values to avoid overflow)
const (
	KB int64 = 1024
	MB int64 = 1024 * 1024
	GB int64 = 1024 * 1024 * 1024
	TB int64 = 1024 * 1024 * 1024 * 1024
)

// Package-level variables
var (
	port = 8080
)

// Custom types
type Celsius float64
type Fahrenheit float64
type UserID int64

// Custom type with methods
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.1f°F", f)
}

// Function to demonstrate zero values
func demonstrateZeroValues() {
	fmt.Println("=== Zero Values ===")

	var i int
	var f float64
	var s string
	var b bool
	var p *int
	var slice []int
	var m map[string]int
	var ch chan int

	fmt.Printf("int zero value: %d\n", i)
	fmt.Printf("float64 zero value: %f\n", f)
	fmt.Printf("string zero value: '%s'\n", s)
	fmt.Printf("bool zero value: %t\n", b)
	fmt.Printf("pointer zero value: %v\n", p)
	fmt.Printf("slice zero value: %v\n", slice)
	fmt.Printf("map zero value: %v\n", m)
	fmt.Printf("channel zero value: %v\n", ch)
	fmt.Println()
}

// Function to demonstrate integer types
func demonstrateIntegerTypes() {
	fmt.Println("=== Integer Types ===")

	// Signed integers
	var a int = 42
	var b int8 = 127
	var c int16 = 32767
	var d int32 = 2147483647
	var e int64 = 9223372036854775807

	fmt.Printf("int: %d\n", a)
	fmt.Printf("int8: %d\n", b)
	fmt.Printf("int16: %d\n", c)
	fmt.Printf("int32: %d\n", d)
	fmt.Printf("int64: %d\n", e)

	// Unsigned integers
	var f uint = 42
	var g uint8 = 255
	var h uint16 = 65535
	var i uint32 = 4294967295
	var j uint64 = 18446744073709551615

	fmt.Printf("uint: %d\n", f)
	fmt.Printf("uint8: %d\n", g)
	fmt.Printf("uint16: %d\n", h)
	fmt.Printf("uint32: %d\n", i)
	fmt.Printf("uint64: %d\n", j)

	// Special types
	var k byte = 65  // Alias for uint8
	var l rune = 'A' // Alias for int32

	fmt.Printf("byte: %d (ASCII: %c)\n", k, k)
	fmt.Printf("rune: %d (Unicode: %c)\n", l, l)
	fmt.Println()
}

// Function to demonstrate floating-point types
func demonstrateFloatTypes() {
	fmt.Println("=== Floating-Point Types ===")

	var a float32 = 3.14159
	var b float64 = 3.141592653589793

	fmt.Printf("float32: %.5f\n", a)
	fmt.Printf("float64: %.15f\n", b)

	// Mathematical constants
	fmt.Printf("Pi (math.Pi): %.15f\n", math.Pi)
	fmt.Printf("E (math.E): %.15f\n", math.E)

	// Special values
	fmt.Printf("Positive infinity: %f\n", math.Inf(1))
	fmt.Printf("Negative infinity: %f\n", math.Inf(-1))
	fmt.Printf("NaN: %f\n", math.NaN())
	fmt.Println()
}

// Function to demonstrate string operations
func demonstrateStringTypes() {
	fmt.Println("=== String Types ===")

	// Basic strings
	message := "Hello, Go!"
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Length (bytes): %d\n", len(message))
	fmt.Printf("Length (runes): %d\n", utf8.RuneCountInString(message))

	// Multi-line strings
	multiLine := `This is a
multi-line string
using backticks`
	fmt.Printf("Multi-line: %s\n", multiLine)

	// String concatenation
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Printf("Full name: %s\n", fullName)

	// String with Unicode
	unicodeStr := "Hello, 世界! 🌍"
	fmt.Printf("Unicode string: %s\n", unicodeStr)
	fmt.Printf("Length (bytes): %d\n", len(unicodeStr))
	fmt.Printf("Length (runes): %d\n", utf8.RuneCountInString(unicodeStr))

	// Accessing characters
	fmt.Println("Characters in 'Hello':")
	for i, char := range "Hello" {
		fmt.Printf("Index %d: %c (Unicode: %d)\n", i, char, char)
	}
	fmt.Println()
}

// Function to demonstrate boolean operations
func demonstrateBooleanTypes() {
	fmt.Println("=== Boolean Types ===")

	// Basic boolean values
	isActive := true
	isComplete := false

	fmt.Printf("isActive: %t\n", isActive)
	fmt.Printf("isComplete: %t\n", isComplete)

	// Boolean operations
	a := true
	b := false

	fmt.Printf("a && b (AND): %t\n", a && b)
	fmt.Printf("a || b (OR): %t\n", a || b)
	fmt.Printf("!a (NOT): %t\n", !a)
	fmt.Printf("!b (NOT): %t\n", !b)

	// Boolean expressions
	age := 25
	isAdult := age >= 18
	hasLicense := true
	canDrive := isAdult && hasLicense

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is adult: %t\n", isAdult)
	fmt.Printf("Has license: %t\n", hasLicense)
	fmt.Printf("Can drive: %t\n", canDrive)
	fmt.Println()
}

// Function to demonstrate variable declaration
func demonstrateVariableDeclaration() {
	fmt.Println("=== Variable Declaration ===")

	// Explicit declaration
	var name string = "Alice"
	var age int = 25
	var height float64 = 165.5
	var isStudent bool = true

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.1f\n", height)
	fmt.Printf("Is student: %t\n", isStudent)

	// Type inference
	var city = "New York"
	var population = 8336817
	var temperature = 72.5
	var isCapital = false

	fmt.Printf("City: %s\n", city)
	fmt.Printf("Population: %d\n", population)
	fmt.Printf("Temperature: %.1f\n", temperature)
	fmt.Printf("Is capital: %t\n", isCapital)

	// Short declaration
	country := "USA"
	area := 9833517.0
	hasStates := true

	fmt.Printf("Country: %s\n", country)
	fmt.Printf("Area: %.0f\n", area)
	fmt.Printf("Has states: %t\n", hasStates)

	// Multiple variables
	var (
		firstName = "John"
		lastName  = "Smith"
		userAge   = 30
		userID    = 12345
	)

	fmt.Printf("User: %s %s (ID: %d, Age: %d)\n", firstName, lastName, userID, userAge)
	fmt.Println()
}

// Function to demonstrate constants
func demonstrateConstants() {
	fmt.Println("=== Constants ===")

	// Basic constants
	const Pi = 3.14159
	const MaxRetries = 3
	const AppName = "DataTypesExamples"

	fmt.Printf("Pi: %.5f\n", Pi)
	fmt.Printf("Max retries: %d\n", MaxRetries)
	fmt.Printf("App name: %s\n", AppName)

	// Typed constants
	const PiFloat64 float64 = 3.141592653589793
	const MaxUsers int = 1000

	fmt.Printf("Pi (float64): %.15f\n", PiFloat64)
	fmt.Printf("Max users: %d\n", MaxUsers)

	// Constant expressions
	const (
		Sum      = 1 + 2
		Product  = 3 * 4
		Greeting = "Hello" + " " + "World"
	)

	fmt.Printf("Sum: %d\n", Sum)
	fmt.Printf("Product: %d\n", Product)
	fmt.Printf("Greeting: %s\n", Greeting)

	// Using iota
	fmt.Printf("File permissions - Read: %d, Write: %d, Execute: %d\n", FlagRead, FlagWrite, FlagExecute)
	fmt.Printf("File sizes - KB: %d, MB: %d, GB: %d, TB: %d\n", KB, MB, GB, TB)
	fmt.Println()
}

// Function to demonstrate type conversion
func demonstrateTypeConversion() {
	fmt.Println("=== Type Conversion ===")

	// Integer conversions
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(i)

	fmt.Printf("int: %d -> float64: %f\n", i, f)
	fmt.Printf("int: %d -> uint: %d\n", i, u)

	// Float to integer (truncates)
	var f2 float64 = 3.14
	var i2 int = int(f2)

	fmt.Printf("float64: %.2f -> int: %d\n", f2, i2)

	// String conversions
	num := 42
	str1 := string(num)            // Converts to Unicode character
	str2 := fmt.Sprintf("%d", num) // Converts to string representation

	fmt.Printf("int: %d -> string (Unicode): %s\n", num, str1)
	fmt.Printf("int: %d -> string (decimal): %s\n", num, str2)

	// Using strconv package
	strNum := "42"
	intNum, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
	} else {
		fmt.Printf("string: %s -> int: %d\n", strNum, intNum)
	}

	floatStr := "3.14"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Printf("Error converting string to float: %v\n", err)
	} else {
		fmt.Printf("string: %s -> float64: %.2f\n", floatStr, floatNum)
	}

	// Convert back to string
	intBackToStr := strconv.Itoa(intNum)
	fmt.Printf("int: %d -> string: %s\n", intNum, intBackToStr)
	fmt.Println()
}

// Function to demonstrate custom types
func demonstrateCustomTypes() {
	fmt.Println("=== Custom Types ===")

	// Temperature conversions
	celsius := Celsius(25.0)
	fahrenheit := Fahrenheit(77.0)

	fmt.Printf("Temperature: %s = %s\n", celsius, celsius.ToFahrenheit())
	fmt.Printf("Temperature: %s = %s\n", fahrenheit, fahrenheit.ToCelsius())

	// User ID type
	userID := UserID(12345)
	fmt.Printf("User ID: %d\n", userID)

	// Type conversion required for custom types
	var regularInt int = 42
	var customInt UserID = UserID(regularInt)

	fmt.Printf("Regular int: %d\n", regularInt)
	fmt.Printf("Custom int: %d\n", customInt)
	fmt.Println()
}

// Function to demonstrate variable scoping
func demonstrateVariableScoping() {
	fmt.Println("=== Variable Scoping ===")

	// Package-level variable (global)
	fmt.Printf("Global port: %d\n", port)

	// Function-level variable
	localVar := "I'm local to this function"
	fmt.Printf("Local variable: %s\n", localVar)

	// Block-level variable
	{
		blockVar := "I'm in a block"
		fmt.Printf("Block variable: %s\n", blockVar)
	}
	// blockVar is not accessible here

	// Shadowing
	localPort := 9090 // Shadows the package-level port
	fmt.Printf("Local port: %d\n", localPort)
	fmt.Printf("Global port (still accessible): %d\n", port)
	fmt.Println()
}

func main() {
	fmt.Printf("=== %s ===\n", "Go Data Types & Variables Examples")
	fmt.Println()

	demonstrateZeroValues()
	demonstrateIntegerTypes()
	demonstrateFloatTypes()
	demonstrateStringTypes()
	demonstrateBooleanTypes()
	demonstrateVariableDeclaration()
	demonstrateConstants()
	demonstrateTypeConversion()
	demonstrateCustomTypes()
	demonstrateVariableScoping()

	fmt.Println("=== All examples completed successfully ===")
}
