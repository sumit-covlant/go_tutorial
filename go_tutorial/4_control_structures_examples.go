package main

import (
	"fmt"
	"runtime"
)

// Function to demonstrate if statements
func demonstrateIfStatements() {
	fmt.Println("=== If Statements ===")

	// Basic if statement
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// If-else statement
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

	// If with initialization
	if userAge := getUserAge(); userAge >= 18 {
		fmt.Printf("User is %d years old and can vote\n", userAge)
	} else {
		fmt.Printf("User is %d years old and cannot vote\n", userAge)
	}

	// If with multiple initialization
	if user, err := getUser(); err == nil {
		fmt.Printf("Welcome, %s!\n", user.Name)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()
}

// Function to demonstrate switch statements
func demonstrateSwitchStatements() {
	fmt.Println("=== Switch Statements ===")

	// Basic switch statement
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

	// Switch with expression
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

	// Switch with initialization
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

	// Switch with fallthrough
	fmt.Println("Fallthrough example:")
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

	// Type switch
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

	fmt.Println()
}

// Function to demonstrate for loops
func demonstrateForLoops() {
	fmt.Println("=== For Loops ===")

	// Traditional for loop
	fmt.Println("Traditional for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// While-style loop
	fmt.Println("While-style loop:")
	count := 0
	for count < 5 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Println("Infinite loop with break:")
	count = 0
	for {
		fmt.Printf("%d ", count)
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println()

	// For-range with slice
	fmt.Println("For-range with slice:")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// For-range with map
	fmt.Println("For-range with map:")
	person := map[string]string{
		"name": "John",
		"city": "New York",
		"job":  "Developer",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// For-range with string
	fmt.Println("For-range with string:")
	message := "Hello"
	for index, char := range message {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// Ignore index or value
	fmt.Println("Ignoring index:")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	fmt.Println("Ignoring value:")
	for index := range numbers {
		fmt.Printf("Index: %d\n", index)
	}

	fmt.Println()
}

// Function to demonstrate break and continue
func demonstrateBreakAndContinue() {
	fmt.Println("=== Break and Continue ===")

	// Break statement
	fmt.Println("Break example:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Exit loop when i equals 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Continue statement
	fmt.Println("Continue example (skip even numbers):")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Break with labels
	fmt.Println("Break with labels:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // Break out of both loops
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()

	// Continue with labels
	fmt.Println("Continue with labels:")
outer2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue outer2 // Skip to next iteration of outer loop
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()

	fmt.Println()
}

// Function to demonstrate nested control structures
func demonstrateNestedControlStructures() {
	fmt.Println("=== Nested Control Structures ===")

	// Nested if statements
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

	// Nested loops
	fmt.Println("Nested loops:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}

	// Switch inside loop
	fmt.Println("Switch inside loop:")
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

	fmt.Println()
}

// Function to demonstrate common patterns
func demonstrateCommonPatterns() {
	fmt.Println("=== Common Patterns ===")

	// Input validation
	fmt.Println("Input validation:")
	ages := []int{-5, 10, 25, 70}
	for _, age := range ages {
		result := validateAge(age)
		fmt.Printf("Age %d: %s\n", age, result)
	}

	// Menu system
	fmt.Println("Menu system:")
	choices := []int{1, 2, 3, 4, 5}
	for _, choice := range choices {
		processChoice(choice)
	}

	// Error handling
	fmt.Println("Error handling:")
	testData := []string{"", "short", "valid data"}
	for _, data := range testData {
		if err := processData(data); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Successfully processed: %s\n", data)
		}
	}

	// Loop with early exit
	fmt.Println("Loop with early exit:")
	numbers := []int{1, 2, 3, 4, 5}
	target := 3
	if index, found := findNumber(numbers, target); found {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found\n", target)
	}

	fmt.Println()
}

// Helper functions for demonstrations

func getUserAge() int {
	return 20
}

type User struct {
	Name string
}

func getUser() (User, error) {
	return User{Name: "Alice"}, nil
}

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

func findNumber(numbers []int, target int) (int, bool) {
	for i, num := range numbers {
		if num == target {
			return i, true
		}
	}
	return -1, false
}

// Function to demonstrate best practices
func demonstrateBestPractices() {
	fmt.Println("=== Best Practices ===")

	// Use switch for multiple conditions
	ages := []int{5, 15, 25, 70}
	for _, age := range ages {
		category := getAgeCategory(age)
		fmt.Printf("Age %d: %s\n", age, category)
	}

	// Use for-range for iteration
	fmt.Println("Using for-range:")
	items := []string{"apple", "banana", "cherry"}
	for index, value := range items {
		fmt.Printf("Item %d: %s\n", index, value)
	}

	// Use break for early exit
	fmt.Println("Using break for early exit:")
	target := "banana"
	found := false
	for _, item := range items {
		if item == target {
			found = true
			break
		}
	}
	fmt.Printf("Found %s: %t\n", target, found)

	// Use initialization in control structures
	fmt.Println("Using initialization in control structures:")
	if value, err := getValue(); err == nil {
		fmt.Printf("Got value: %d\n", value)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()
}

func getAgeCategory(age int) string {
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
}

func getValue() (int, error) {
	return 42, nil
}

func main() {
	fmt.Printf("=== %s ===\n", "Go Control Structures Examples")
	fmt.Println()

	demonstrateIfStatements()
	demonstrateSwitchStatements()
	demonstrateForLoops()
	demonstrateBreakAndContinue()
	demonstrateNestedControlStructures()
	demonstrateCommonPatterns()
	demonstrateBestPractices()

	fmt.Println("=== All control structure examples completed successfully ===")
}
