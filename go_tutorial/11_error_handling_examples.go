package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
)

// This file demonstrates Go error handling concepts

func main() {
	fmt.Println("=== Go Error Handling Examples ===\n")

	// Basic error handling examples
	basicErrorHandling()

	// Error creation examples
	errorCreationExamples()

	// Error handling patterns
	errorHandlingPatterns()

	// Error types and categories
	errorTypesAndCategories()

	// Error handling best practices
	errorHandlingBestPractices()

	// Error handling in different contexts
	errorHandlingInContexts()

	// Error logging
	errorLogging()

	// Testing error handling
	testingErrorHandling()

	// Common pitfalls
	commonPitfalls()
}

// Basic error handling examples
func basicErrorHandling() {
	fmt.Println("1. Basic Error Handling Examples")
	fmt.Println("---------------------------------")

	// Simple error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	// Error handling with file operations
	err = processFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}

	// Error handling with validation
	err = validateAge(25)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Age is valid")
	}

	err = validateAge(-5)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}
	fmt.Println()
}

// Simple division function with error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// File processing function
func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Process file...
	return nil
}

// Age validation function
func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age cannot be negative, got %d", age)
	}
	if age > 150 {
		return fmt.Errorf("age cannot exceed 150, got %d", age)
	}
	return nil
}

// Error creation examples
func errorCreationExamples() {
	fmt.Println("2. Error Creation Examples")
	fmt.Println("--------------------------")

	// Using errors.New()
	err1 := errors.New("simple error message")
	fmt.Printf("Simple error: %v\n", err1)

	// Using fmt.Errorf()
	err2 := fmt.Errorf("formatted error with value: %d", 42)
	fmt.Printf("Formatted error: %v\n", err2)

	// Custom error types
	err3 := ValidationError{
		Field:   "email",
		Message: "invalid format",
		Value:   "invalid-email",
	}
	fmt.Printf("Custom error: %v\n", err3)

	// Error with context
	err4 := fmt.Errorf("failed to process user data: %w", err3)
	fmt.Printf("Error with context: %v\n", err4)
	fmt.Println()
}

// Custom error type
type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %s: %s (value: %v)",
		e.Field, e.Message, e.Value)
}

// Error handling patterns
func errorHandlingPatterns() {
	fmt.Println("3. Error Handling Patterns")
	fmt.Println("---------------------------")

	// Early return pattern
	err := processUser("", -5)
	if err != nil {
		fmt.Printf("Process user error: %v\n", err)
	}

	// Error wrapping
	err = readConfig()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		// Unwrap the error
		if unwrapped := errors.Unwrap(err); unwrapped != nil {
			fmt.Printf("Unwrapped error: %v\n", unwrapped)
		}
	}

	// Error checking with errors.Is()
	user, err := findUser("nonexistent")
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("User not found")
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
	} else {
		fmt.Printf("Found user: %+v\n", user)
	}

	// Type assertions with errors.As()
	err = validateAge(-5)
	if err != nil {
		var valErr ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("Validation error on field %s: %s\n",
				valErr.Field, valErr.Message)
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
	}
	fmt.Println()
}

// Sentinel errors
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

// Early return pattern
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

// Error wrapping example
func readConfig() error {
	err := readFile("config.json")
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}
	return nil
}

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Read file...
	return nil
}

// User struct and find function
type User struct {
	ID   string
	Name string
}

func findUser(id string) (*User, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	// Simulate not found
	if id == "nonexistent" {
		return nil, ErrNotFound
	}

	return &User{ID: id, Name: "John Doe"}, nil
}

// Error types and categories
func errorTypesAndCategories() {
	fmt.Println("4. Error Types and Categories")
	fmt.Println("------------------------------")

	// Sentinel errors
	fmt.Println("Sentinel errors:")
	fmt.Printf("ErrNotFound: %v\n", ErrNotFound)
	fmt.Printf("ErrUnauthorized: %v\n", ErrUnauthorized)
	fmt.Printf("ErrInvalidInput: %v\n", ErrInvalidInput)

	// Error types
	fmt.Println("\nError types:")
	notFoundErr := NotFoundError{Resource: "user", ID: "123"}
	fmt.Printf("NotFoundError: %v\n", notFoundErr)

	// Wrapped errors
	fmt.Println("\nWrapped errors:")
	err := processUserWithWrapping("nonexistent")
	if err != nil {
		fmt.Printf("Wrapped error: %v\n", err)

		// Check for specific error types
		var notFound NotFoundError
		if errors.As(err, &notFound) {
			fmt.Printf("Not found: %s\n", notFound.Error())
		}
	}
	fmt.Println()
}

// Error types
type NotFoundError struct {
	Resource string
	ID       string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with id %s not found", e.Resource, e.ID)
}

// Wrapped error example
func processUserWithWrapping(id string) error {
	user, err := findUser(id)
	if err != nil {
		return fmt.Errorf("failed to process user %s: %w", id, err)
	}

	// Process user...
	fmt.Printf("Processing user: %+v\n", user)
	return nil
}

// Error handling best practices
func errorHandlingBestPractices() {
	fmt.Println("5. Error Handling Best Practices")
	fmt.Println("---------------------------------")

	// Always check errors
	fmt.Println("Always check errors:")
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	} else {
		defer file.Close()
		fmt.Println("File opened successfully")
	}

	// Return errors, don't panic
	fmt.Println("\nReturn errors, don't panic:")
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Safe division error: %v\n", err)
	} else {
		fmt.Printf("Safe division result: %d\n", result)
	}

	// Add context to errors
	fmt.Println("\nAdd context to errors:")
	err = readConfigWithContext()
	if err != nil {
		fmt.Printf("Config error with context: %v\n", err)
	}
	fmt.Println()
}

// Safe division function
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Config reading with context
func readConfigWithContext() error {
	err := readFile("config.json")
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	return nil
}

// Error handling in different contexts
func errorHandlingInContexts() {
	fmt.Println("6. Error Handling in Different Contexts")
	fmt.Println("----------------------------------------")

	// HTTP handler context
	fmt.Println("HTTP handler context:")
	handleGetUserExample()

	// Database operations context
	fmt.Println("\nDatabase operations context:")
	user, err := getUserByIDExample("123")
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}

	// File operations context
	fmt.Println("\nFile operations context:")
	data, err := readFileExample("nonexistent.txt")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	} else {
		fmt.Printf("File data: %s\n", string(data))
	}
	fmt.Println()
}

// HTTP handler example
func handleGetUserExample() {
	// Simulate HTTP request
	id := "nonexistent"

	if id == "" {
		fmt.Println("HTTP 400: missing user id")
		return
	}

	user, err := findUser(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("HTTP 404: user not found")
		} else {
			fmt.Println("HTTP 500: internal server error")
		}
		return
	}

	fmt.Printf("HTTP 200: user found - %+v\n", user)
}

// Database operation example
func getUserByIDExample(id string) (*User, error) {
	// Simulate database query
	if id == "nonexistent" {
		return nil, ErrNotFound
	}

	// Simulate database error
	if id == "error" {
		return nil, fmt.Errorf("database connection failed")
	}

	return &User{ID: id, Name: "John Doe"}, nil
}

// File operation example
func readFileExample(filename string) ([]byte, error) {
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

// Error logging
func errorLogging() {
	fmt.Println("7. Error Logging")
	fmt.Println("----------------")

	// Structured error logging
	fmt.Println("Structured error logging:")
	err := processRequestExample()
	if err != nil {
		fmt.Printf("Request processing failed: %v\n", err)
	}

	// Error with stack trace
	fmt.Println("\nError with stack trace:")
	err = processDataExample([]byte{})
	if err != nil {
		logError(err)
	}
	fmt.Println()
}

// Request processing example
func processRequestExample() error {
	err := validateRequestExample()
	if err != nil {
		fmt.Printf("Request validation failed: %v\n", err)
		return fmt.Errorf("invalid request: %w", err)
	}

	// Process request...
	return nil
}

func validateRequestExample() error {
	return errors.New("validation failed")
}

// Data processing example
func processDataExample(data []byte) error {
	if len(data) == 0 {
		err := errors.New("empty data")
		return err
	}

	// Process data...
	return nil
}

// Error logging with stack trace
func logError(err error) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("Error at %s:%d: %v\n", file, line, err)
}

// Testing error handling
func testingErrorHandling() {
	fmt.Println("8. Testing Error Handling")
	fmt.Println("-------------------------")

	// Test error returns
	fmt.Println("Testing error returns:")
	runErrorTests()

	// Test custom error types
	fmt.Println("\nTesting custom error types:")
	testCustomErrorTypes()
	fmt.Println()
}

// Error test cases
func runErrorTests() {
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
				fmt.Printf("✗ Expected error for %d / %d\n", test.a, test.b)
			} else {
				fmt.Printf("✓ Got expected error for %d / %d: %v\n", test.a, test.b, err)
			}
		} else {
			if err != nil {
				fmt.Printf("✗ Unexpected error for %d / %d: %v\n", test.a, test.b, err)
			} else if result != test.expected {
				fmt.Printf("✗ Expected %d, got %d\n", test.expected, result)
			} else {
				fmt.Printf("✓ %d / %d = %d\n", test.a, test.b, result)
			}
		}
	}
}

// Custom error type testing
func testCustomErrorTypes() {
	err := ValidationError{Field: "age", Message: "cannot be negative"}

	if err.Field != "age" {
		fmt.Printf("✗ Expected field 'age', got %s\n", err.Field)
	} else {
		fmt.Printf("✓ Field is correct: %s\n", err.Field)
	}

	if err.Message != "cannot be negative" {
		fmt.Printf("✗ Expected message 'cannot be negative', got %s\n", err.Message)
	} else {
		fmt.Printf("✓ Message is correct: %s\n", err.Message)
	}

	expectedMsg := "validation error on field age: cannot be negative (value: <nil>)"
	if err.Error() != expectedMsg {
		fmt.Printf("✗ Expected error message '%s', got '%s'\n", expectedMsg, err.Error())
	} else {
		fmt.Printf("✓ Error message is correct: %s\n", err.Error())
	}
}

// Common pitfalls
func commonPitfalls() {
	fmt.Println("9. Common Pitfalls")
	fmt.Println("------------------")

	// Ignoring errors
	fmt.Println("Ignoring errors:")
	fmt.Println("✗ file, _ := os.Open(\"file.txt\")  // Don't do this!")
	fmt.Println("✓ file, err := os.Open(\"file.txt\")")
	fmt.Println("  if err != nil {")
	fmt.Println("      return fmt.Errorf(\"failed to open file: %w\", err)")
	fmt.Println("  }")

	// Overly generic error messages
	fmt.Println("\nOverly generic error messages:")
	fmt.Println("✗ return errors.New(\"error occurred\")")
	fmt.Println("✓ return fmt.Errorf(\"failed to read configuration file: %w\", err)")

	// Not using error wrapping
	fmt.Println("\nNot using error wrapping:")
	fmt.Println("✗ return err  // Loses context")
	fmt.Println("✓ return fmt.Errorf(\"failed to process configuration: %w\", err)")

	// Demonstrate good error handling
	fmt.Println("\nGood error handling example:")
	err := goodErrorHandlingExample()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}

// Good error handling example
func goodErrorHandlingExample() error {
	// Step 1: Open file
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	// Step 2: Read file
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Step 3: Process data
	if n == 0 {
		return errors.New("config file is empty")
	}

	fmt.Printf("Successfully read %d bytes from config file\n", n)
	return nil
}

// Additional utility functions for demonstration
func divideAndModulo(a, b int) (quotient, remainder int, err error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return a / b, a % b, nil
}

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

// Database error type
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

// Simulate database operation
func insertUser(user *User) error {
	// Simulate database error
	if user.ID == "error" {
		return DatabaseError{
			Operation: "insert",
			Table:     "users",
			Err:       errors.New("connection timeout"),
		}
	}

	// Success
	return nil
}
