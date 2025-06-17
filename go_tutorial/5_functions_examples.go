package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Basic function examples
func demonstrateBasicFunctions() {
	fmt.Println("=== Basic Functions ===")

	// Function with no parameters
	sayHello()

	// Function with parameters
	message := greet("Alice")
	fmt.Printf("Greeting: %s\n", message)

	// Function with multiple parameters
	result := add(10, 5)
	fmt.Printf("10 + 5 = %d\n", result)

	// Function with no return value
	printMessage("This is a test message")

	fmt.Println()
}

// Return values examples
func demonstrateReturnValues() {
	fmt.Println("=== Return Values ===")

	// Single return value
	squared := square(5)
	fmt.Printf("5 squared = %d\n", squared)

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

	// Multiple return values of same type
	min, max := minMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// Ignoring return values
	result, _ := divide(20, 4)
	fmt.Printf("20 / 4 = %d (ignored error)\n", result)

	fmt.Println()
}

// Named return values examples
func demonstrateNamedReturnValues() {
	fmt.Println("=== Named Return Values ===")

	// Basic named return values
	q, r := divideAndRemainder(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", q, r)

	// Named return values with error handling
	name, age, err := getUserInfo(123)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("User: %s, Age: %d\n", name, age)
	}

	// Error case
	_, _, err = getUserInfo(-1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()
}

// Variadic functions examples
func demonstrateVariadicFunctions() {
	fmt.Println("=== Variadic Functions ===")

	// Basic variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)

	// Variadic function with no arguments
	emptySum := sum()
	fmt.Printf("Sum of no numbers = %d\n", emptySum)

	// Variadic function with regular parameters
	joined := join(", ", "apple", "banana", "cherry")
	fmt.Printf("Joined: %s\n", joined)

	// Passing slice to variadic function
	numbers := []int{10, 20, 30, 40, 50}
	sliceSum := sum(numbers...)
	fmt.Printf("Sum of slice %v = %d\n", numbers, sliceSum)

	// Variadic function with different types
	printInfo("John", 25, "Developer", "New York")

	fmt.Println()
}

// Function types examples
func demonstrateFunctionTypes() {
	fmt.Println("=== Function Types ===")

	// Function as variable
	var operation func(int, int) int

	operation = add
	result := operation(5, 3)
	fmt.Printf("Add operation: 5 + 3 = %d\n", result)

	operation = multiply
	result = operation(5, 3)
	fmt.Printf("Multiply operation: 5 * 3 = %d\n", result)

	// Function as parameter
	result = applyOperation(10, 5, add)
	fmt.Printf("Applied add operation: 10 + 5 = %d\n", result)

	result = applyOperation(10, 5, multiply)
	fmt.Printf("Applied multiply operation: 10 * 5 = %d\n", result)

	// Function as return value
	addFunc := getOperation("add")
	result = addFunc(8, 4)
	fmt.Printf("Returned add function: 8 + 4 = %d\n", result)

	multiplyFunc := getOperation("multiply")
	result = multiplyFunc(8, 4)
	fmt.Printf("Returned multiply function: 8 * 4 = %d\n", result)

	fmt.Println()
}

// Anonymous functions and closures examples
func demonstrateAnonymousFunctions() {
	fmt.Println("=== Anonymous Functions & Closures ===")

	// Basic anonymous function
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	greet("Bob")

	// Immediately invoked function expression (IIFE)
	result := func(a, b int) int {
		return a + b
	}(5, 3)
	fmt.Printf("IIFE result: 5 + 3 = %d\n", result)

	// Closure with captured variables
	counter := createCounter()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())

	// Another counter instance
	counter2 := createCounter()
	fmt.Printf("Counter2: %d\n", counter2())

	// Closure with multiple captured variables
	adder := createAdder(10)
	fmt.Printf("Adder(5): %d\n", adder(5))
	fmt.Printf("Adder(3): %d\n", adder(3))

	fmt.Println()
}

// Defer examples
func demonstrateDefer() {
	fmt.Println("=== Defer Statements ===")

	// Basic defer
	fmt.Println("Basic defer:")
	deferExample()

	// Multiple defer statements (LIFO order)
	fmt.Println("Multiple defer statements:")
	multipleDeferExample()

	// Defer with arguments
	fmt.Println("Defer with arguments:")
	deferWithArguments()

	// Defer with file operations
	fmt.Println("Defer with file operations:")
	err := processFileExample("example.txt")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}

	// Defer with named return values
	fmt.Println("Defer with named return values:")
	result, err := processDataWithDefer()
	if err != nil {
		fmt.Printf("Process error: %v\n", err)
	} else {
		fmt.Printf("Process result: %s\n", result)
	}

	fmt.Println()
}

// Recursion examples
func demonstrateRecursion() {
	fmt.Println("=== Function Recursion ===")

	// Basic recursion - factorial
	n := 5
	fact := factorial(n)
	fmt.Printf("Factorial of %d = %d\n", n, fact)

	// Recursion with memoization - fibonacci
	fmt.Println("Fibonacci numbers:")
	for i := 0; i <= 10; i++ {
		fib := fibonacci(i)
		fmt.Printf("fibonacci(%d) = %d\n", i, fib)
	}

	// Recursive function with early exit
	fmt.Println("Finding number in array:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	found := binarySearch(numbers, target, 0, len(numbers)-1)
	fmt.Printf("Found %d: %t\n", target, found)

	fmt.Println()
}

// Error handling patterns
func demonstrateErrorHandling() {
	fmt.Println("=== Error Handling Patterns ===")

	// Return error pattern
	config, err := readConfig("config.txt")
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
	} else {
		fmt.Printf("Config loaded: %s\n", config)
	}

	// Error wrapping
	err = processUser(123)
	if err != nil {
		fmt.Printf("User processing error: %v\n", err)
	}

	// Multiple error checks
	err = validateAndProcess("test@example.com")
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}

	fmt.Println()
}

// Optional parameters patterns
func demonstrateOptionalParameters() {
	fmt.Println("=== Optional Parameters Patterns ===")

	// Using struct for optional parameters
	config1 := NewConfig()
	processWithConfig("data1", config1)

	config2 := &Config{
		Timeout: 60 * time.Second,
		Retries: 5,
		Debug:   true,
	}
	processWithConfig("data2", config2)

	// Using variadic functions for optional parameters
	connect("localhost")
	connect("localhost", "https")
	connect("localhost", "https", "443")

	fmt.Println()
}

// Helper functions

func sayHello() {
	fmt.Println("Hello, World!")
}

func greet(name string) string {
	return "Hello, " + name + "!"
}

func add(a, b int) int {
	return a + b
}

func printMessage(message string) {
	fmt.Printf("Message: %s\n", message)
}

func square(x int) int {
	return x * x
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

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

func divideAndRemainder(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

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

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

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

func printInfo(name string, age int, info ...string) {
	fmt.Printf("Name: %s, Age: %d", name, age)
	for _, item := range info {
		fmt.Printf(", %s", item)
	}
	fmt.Println()
}

func multiply(a, b int) int {
	return a * b
}

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

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

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func createAdder(initial int) func(int) int {
	sum := initial
	return func(value int) int {
		sum += value
		return sum
	}
}

func deferExample() {
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
	fmt.Println("This will be printed second")
}

func multipleDeferExample() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	fmt.Println("Main function")
}

func deferWithArguments() {
	i := 1
	defer fmt.Printf("Deferred: %d\n", i)

	i = 2
	fmt.Printf("Current: %d\n", i)
}

func processFileExample(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fmt.Printf("Processing file: %s\n", filename)
	// Simulate file processing
	return nil
}

func processDataWithDefer() (result string, err error) {
	defer func() {
		if err != nil {
			result = "" // Can modify named return values
		}
	}()

	// Simulate processing
	result = "processed data"
	return
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

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

func binarySearch(arr []int, target, left, right int) bool {
	if left > right {
		return false
	}

	mid := (left + right) / 2
	if arr[mid] == target {
		return true
	}

	if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1)
	}
	return binarySearch(arr, target, mid+1, right)
}

func readConfig(filename string) (string, error) {
	if filename == "" {
		return "", fmt.Errorf("filename cannot be empty")
	}

	// Simulate reading file
	if filename == "config.txt" {
		return "config data", nil
	}
	return "", fmt.Errorf("file not found: %s", filename)
}

type User struct {
	ID   int
	Name string
}

func getUser(userID int) (User, error) {
	if userID <= 0 {
		return User{}, fmt.Errorf("invalid user ID: %d", userID)
	}
	return User{ID: userID, Name: "John Doe"}, nil
}

func validateUser(user User) error {
	if user.Name == "" {
		return fmt.Errorf("user name cannot be empty")
	}
	return nil
}

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

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func validateAndProcess(email string) error {
	if err := validateEmail(email); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Process email...
	return nil
}

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
	fmt.Printf("Processing %s with timeout: %v, retries: %d, debug: %t\n",
		data, config.Timeout, config.Retries, config.Debug)
}

func connect(host string, options ...string) {
	port := "8080"     // default
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

func main() {
	fmt.Printf("=== %s ===\n", "Go Functions Examples")
	fmt.Println()

	demonstrateBasicFunctions()
	demonstrateReturnValues()
	demonstrateNamedReturnValues()
	demonstrateVariadicFunctions()
	demonstrateFunctionTypes()
	demonstrateAnonymousFunctions()
	demonstrateDefer()
	demonstrateRecursion()
	demonstrateErrorHandling()
	demonstrateOptionalParameters()

	fmt.Println("=== All function examples completed successfully ===")
}
