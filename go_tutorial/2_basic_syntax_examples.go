package main

import (
	"fmt"
	"os"
	"strings"
)

// Package-level constants
const (
	AppName    = "BasicSyntaxExamples"
	Version    = "1.0.0"
	MaxRetries = 3
)

// Package-level variables
var (
	debug = false
	port  = 8080
)

// init function runs before main
func init() {
	fmt.Println("Initializing application...")
}

// Basic function with no parameters
func greet() {
	fmt.Println("Hello from Go!")
}

// Function with parameters
func greetPerson(name string) {
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

// Function with named return values
func divideAndRemainder(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return // Named return values
}

// Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function demonstrating defer
func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close() // This will run when function exits

	fmt.Printf("Processing file: %s\n", filename)
	// Simulate file processing
	return nil
}

// Function demonstrating error handling pattern
func readConfig() (string, error) {
	// Simulate reading configuration
	config := "app.config"
	if config == "" {
		return "", fmt.Errorf("configuration file not found")
	}
	return config, nil
}

// Function demonstrating multiple assignment
func getUserInfo() (string, int, bool) {
	return "John Doe", 25, true
}

// Function demonstrating blank identifier
func getCoordinates() (float64, float64, error) {
	// Simulate getting coordinates
	return 40.7128, -74.0060, nil
}

func main() {
	fmt.Printf("=== %s v%s ===\n", AppName, Version)
	fmt.Println()

	// Basic function calls
	greet()
	greetPerson("Alice")
	fmt.Println()

	// Working with return values
	result := add(10, 5)
	fmt.Printf("10 + 5 = %d\n", result)
	fmt.Println()

	// Multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", quotient)
	}

	// Error handling
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()

	// Named return values
	q, r := divideAndRemainder(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", q, r)
	fmt.Println()

	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
	fmt.Println()

	// Multiple assignment
	name, age, isActive := getUserInfo()
	fmt.Printf("User: %s, Age: %d, Active: %t\n", name, age, isActive)
	fmt.Println()

	// Blank identifier example
	lat, _, err := getCoordinates()
	if err != nil {
		fmt.Printf("Error getting coordinates: %v\n", err)
	} else {
		fmt.Printf("Latitude: %.4f\n", lat)
	}
	fmt.Println()

	// Error handling pattern
	config, err := readConfig()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
	} else {
		fmt.Printf("Config loaded: %s\n", config)
	}
	fmt.Println()

	// Defer example
	err = processFile("example.txt")
	if err != nil {
		fmt.Printf("File processing error: %v\n", err)
	}
	fmt.Println()

	// String operations
	message := "Hello, Go Programming!"
	fmt.Printf("Original: %s\n", message)
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(message))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(message))
	fmt.Printf("Length: %d\n", len(message))
	fmt.Println()

	// Conditional logic
	userAge := 25
	if userAge >= 18 {
		fmt.Println("User is an adult")
	} else {
		fmt.Println("User is a minor")
	}

	// Switch statement
	switch userAge {
	case 0, 1, 2:
		fmt.Println("Toddler")
	case 3, 4, 5:
		fmt.Println("Preschooler")
	case 6, 7, 8, 9, 10, 11, 12:
		fmt.Println("Child")
	case 13, 14, 15, 16, 17:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}
	fmt.Println()

	// For loop
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop as while
	fmt.Println("Counting down from 5:")
	count := 5
	for count > 0 {
		fmt.Printf("%d ", count)
		count--
	}
	fmt.Println()

	// For loop with range (we'll see more of this with slices)
	fmt.Println("Iterating over string characters:")
	for i, char := range "Go" {
		fmt.Printf("Index: %d, Character: %c\n", i, char)
	}
	fmt.Println()

	// Break and continue
	fmt.Println("Using break and continue:")
	for i := 1; i <= 10; i++ {
		if i == 3 {
			continue // Skip 3
		}
		if i == 8 {
			break // Stop at 8
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("\n=== Program completed successfully ===")
}
