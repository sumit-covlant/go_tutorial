package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// This file demonstrates packages and modules concepts
// In a real project, these would be in separate files and directories

func main() {
	fmt.Println("=== Go Packages & Modules Examples ===\n")

	// Demonstrate package concepts
	packageExamples()

	// Demonstrate module concepts
	moduleExamples()

	// Demonstrate package documentation
	packageDocumentationExamples()

	// Demonstrate testing
	testingExamples()

	// Demonstrate best practices
	bestPracticesExamples()
}

// Package examples
func packageExamples() {
	fmt.Println("1. Package Examples")
	fmt.Println("-------------------")

	// Using utility functions from "utils" package
	reversed := reverseString("hello")
	fmt.Printf("Reversed string: %s\n", reversed)

	uppercase := toUpperCase("hello")
	fmt.Printf("Uppercase string: %s\n", uppercase)

	// Using math functions from "math" package
	area := calculateCircleArea(5.0)
	fmt.Printf("Circle area: %.2f\n", area)

	perimeter := calculateRectanglePerimeter(4.0, 6.0)
	fmt.Printf("Rectangle perimeter: %.2f\n", perimeter)

	// Using models from "models" package
	user := createUser("Alice", "alice@example.com")
	fmt.Printf("User: %+v\n", user)

	// Demonstrate package visibility
	demoVisibility()
	fmt.Println()
}

// Utility functions (simulating utils package)
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

// Math functions (simulating math package)
func calculateCircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func calculateCirclePerimeter(radius float64) float64 {
	return 2 * math.Pi * radius
}

func calculateRectangleArea(width, height float64) float64 {
	return width * height
}

func calculateRectanglePerimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Models (simulating models package)
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func createUser(name, email string) *User {
	return &User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

func (u *User) GetFullName() string {
	return u.Name
}

// Demonstrate package visibility
func demoVisibility() {
	fmt.Println("Package visibility demonstration:")

	// Public function (exported)
	result := publicFunction()
	fmt.Printf("Public function result: %s\n", result)

	// Private function (unexported) - this would be in the same package
	// privateFunction() // This would work in the same package

	// Public variable
	fmt.Printf("Public variable: %s\n", publicVariable)

	// Private variable - this would be in the same package
	// fmt.Printf("Private variable: %s\n", privateVariable)
}

// Public (exported) function
func publicFunction() string {
	return "This is a public function"
}

// Private (unexported) function
func privateFunction() string {
	return "This is a private function"
}

// Public variable
var publicVariable = "public"

// Private variable
var privateVariable = "private"

// Module examples
func moduleExamples() {
	fmt.Println("2. Module Examples")
	fmt.Println("------------------")

	// Demonstrate dependency management concepts
	fmt.Println("Module structure:")
	fmt.Println("- go.mod (module definition)")
	fmt.Println("- go.sum (dependency checksums)")
	fmt.Println("- main.go (entry point)")
	fmt.Println("- pkg/ (public packages)")
	fmt.Println("- internal/ (private packages)")
	fmt.Println("- cmd/ (executables)")

	// Demonstrate versioning
	demoVersioning()

	// Demonstrate dependency management
	demoDependencyManagement()
	fmt.Println()
}

func demoVersioning() {
	fmt.Println("\nVersioning examples:")
	fmt.Println("- v1.2.3 (semantic versioning)")
	fmt.Println("- v1.2.3-pre (pre-release)")
	fmt.Println("- v1.2.3+metadata (build metadata)")
	fmt.Println("- v0.0.0-20210921155107-089bfa567519 (pseudo-version)")
}

func demoDependencyManagement() {
	fmt.Println("\nDependency management commands:")
	fmt.Println("- go mod init myproject")
	fmt.Println("- go get github.com/gorilla/mux")
	fmt.Println("- go mod tidy")
	fmt.Println("- go mod download")
	fmt.Println("- go mod verify")
}

// Package documentation examples
func packageDocumentationExamples() {
	fmt.Println("3. Package Documentation Examples")
	fmt.Println("---------------------------------")

	// Demonstrate package comments
	fmt.Println("Package comments should be:")
	fmt.Println("- Placed before package declaration")
	fmt.Println("- Start with 'Package packagename'")
	fmt.Println("- Provide overview of package functionality")

	// Demonstrate function documentation
	result := documentedFunction(5, 3)
	fmt.Printf("Documented function result: %d\n", result)

	// Demonstrate example usage
	fmt.Println("\nExample usage:")
	fmt.Println("```go")
	fmt.Println("result := Add(5, 3)  // result == 8")
	fmt.Println("```")
	fmt.Println()
}

// Documented function with comments
// Add returns the sum of two integers.
//
// Example:
//
//	result := Add(5, 3)  // result == 8
func documentedFunction(a, b int) int {
	return a + b
}

// Testing examples
func testingExamples() {
	fmt.Println("4. Testing Examples")
	fmt.Println("-------------------")

	// Demonstrate test structure
	fmt.Println("Test file structure:")
	fmt.Println("- Test files end with _test.go")
	fmt.Println("- Test functions start with Test")
	fmt.Println("- Benchmark functions start with Benchmark")
	fmt.Println("- Example functions start with Example")

	// Demonstrate test examples
	runTestExamples()

	// Demonstrate benchmark examples
	runBenchmarkExamples()
	fmt.Println()
}

func runTestExamples() {
	fmt.Println("\nRunning test examples:")

	// Test reverseString function
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"123", "321"},
	}

	for _, test := range testCases {
		result := reverseString(test.input)
		if result == test.expected {
			fmt.Printf("✓ reverseString(%q) = %q\n", test.input, result)
		} else {
			fmt.Printf("✗ reverseString(%q) = %q, expected %q\n", test.input, result, test.expected)
		}
	}
}

func runBenchmarkExamples() {
	fmt.Println("\nBenchmark examples:")
	fmt.Println("Benchmark functions measure performance:")
	fmt.Println("func BenchmarkReverse(b *testing.B) {")
	fmt.Println("    for i := 0; i < b.N; i++ {")
	fmt.Println("        reverseString(\"hello world\")")
	fmt.Println("    }")
	fmt.Println("}")
}

// Best practices examples
func bestPracticesExamples() {
	fmt.Println("5. Best Practices Examples")
	fmt.Println("---------------------------")

	// Demonstrate package design
	demoPackageDesign()

	// Demonstrate package naming
	demoPackageNaming()

	// Demonstrate package organization
	demoPackageOrganization()

	// Demonstrate package dependencies
	demoPackageDependencies()
	fmt.Println()
}

func demoPackageDesign() {
	fmt.Println("Package design principles:")
	fmt.Println("✓ Single responsibility")
	fmt.Println("  - math package: mathematical operations")
	fmt.Println("  - strings package: string manipulation")
	fmt.Println("  - time package: time operations")
	fmt.Println("✗ Multiple responsibilities")
	fmt.Println("  - utils package: too generic")
	fmt.Println("  - helper package: unclear purpose")
}

func demoPackageNaming() {
	fmt.Println("\nPackage naming conventions:")
	fmt.Println("✓ Good names:")
	fmt.Println("  - user")
	fmt.Println("  - auth")
	fmt.Println("  - database")
	fmt.Println("✗ Avoid:")
	fmt.Println("  - helper")
	fmt.Println("  - common")
	fmt.Println("  - util")
}

func demoPackageOrganization() {
	fmt.Println("\nPackage organization:")
	fmt.Println("myproject/")
	fmt.Println("├── cmd/           # Main applications")
	fmt.Println("│   ├── server/")
	fmt.Println("│   └── client/")
	fmt.Println("├── internal/      # Private application code")
	fmt.Println("│   ├── auth/")
	fmt.Println("│   └── database/")
	fmt.Println("├── pkg/           # Public library code")
	fmt.Println("│   ├── utils/")
	fmt.Println("│   └── models/")
	fmt.Println("└── api/           # API definitions")
	fmt.Println("    └── v1/")
}

func demoPackageDependencies() {
	fmt.Println("\nPackage dependencies:")
	fmt.Println("✓ Minimal dependencies:")
	fmt.Println("  import \"strings\"  // Only what you need")
	fmt.Println("✗ Unnecessary dependencies:")
	fmt.Println("  import (")
	fmt.Println("    \"fmt\"")
	fmt.Println("    \"math\"")
	fmt.Println("    \"strings\"")
	fmt.Println("    \"time\"")
	fmt.Println("    // ... many more")
	fmt.Println("  )")
}

// Common patterns examples
func commonPatternsExamples() {
	fmt.Println("6. Common Patterns Examples")
	fmt.Println("----------------------------")

	// Package initialization
	demoPackageInitialization()

	// Package configuration
	demoPackageConfiguration()

	// Package factories
	demoPackageFactories()
	fmt.Println()
}

func demoPackageInitialization() {
	fmt.Println("Package initialization pattern:")
	fmt.Println("```go")
	fmt.Println("package database")
	fmt.Println("")
	fmt.Println("var db *sql.DB")
	fmt.Println("")
	fmt.Println("func init() {")
	fmt.Println("    // Package initialization code")
	fmt.Println("    var err error")
	fmt.Println("    db, err = sql.Open(\"postgres\", \"connection_string\")")
	fmt.Println("    if err != nil {")
	fmt.Println("        panic(err)")
	fmt.Println("    }")
	fmt.Println("}")
	fmt.Println("```")
}

func demoPackageConfiguration() {
	fmt.Println("\nPackage configuration pattern:")
	fmt.Println("```go")
	fmt.Println("type Config struct {")
	fmt.Println("    DatabaseURL string")
	fmt.Println("    Port        string")
	fmt.Println("    Environment string")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("func Load() *Config {")
	fmt.Println("    // Load configuration from environment")
	fmt.Println("    return &Config{...}")
	fmt.Println("}")
	fmt.Println("```")
}

func demoPackageFactories() {
	fmt.Println("\nPackage factory pattern:")
	fmt.Println("```go")
	fmt.Println("type Logger struct {")
	fmt.Println("    level string")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("func NewLogger(level string) *Logger {")
	fmt.Println("    return &Logger{level: level}")
	fmt.Println("}")
	fmt.Println("```")
}

// Additional utility functions for demonstration
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return b - a
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a) / float64(b), nil
}
