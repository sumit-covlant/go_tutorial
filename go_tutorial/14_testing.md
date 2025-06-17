# Go Testing

## Overview

Go has a built-in testing framework that makes it easy to write and run tests. The testing package provides a simple yet powerful way to test your Go code, including unit tests, benchmarks, and examples.

## Basic Testing

### Test File Structure

Test files in Go follow a specific naming convention and structure:

```go
// File: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

### Running Tests

```bash
# Run all tests in current package
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run specific test
go test -run TestAdd

# Run tests and generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Function Naming

```go
// Test functions must start with "Test"
func TestFunctionName(t *testing.T) {
    // Test implementation
}

// Benchmark functions must start with "Benchmark"
func BenchmarkFunctionName(b *testing.B) {
    // Benchmark implementation
}

// Example functions must start with "Example"
func ExampleFunctionName() {
    // Example implementation
}
```

## Unit Testing

### Basic Unit Test

```go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    expected := 2
    if result != expected {
        t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
    }
}
```

### Table-Driven Tests

Table-driven tests are a common pattern in Go for testing multiple scenarios:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 5, 5},
        {"mixed signs", -2, 3, 1},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Add(test.a, test.b)
            if result != test.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    test.a, test.b, result, test.expected)
            }
        })
    }
}
```

### Testing Error Conditions

```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name        string
        a, b        int
        expected    int
        expectError bool
    }{
        {"valid division", 10, 2, 5, false},
        {"division by zero", 10, 0, 0, true},
        {"zero dividend", 0, 5, 0, false},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := Divide(test.a, test.b)
            
            if test.expectError {
                if err == nil {
                    t.Errorf("Expected error but got none")
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if result != test.expected {
                    t.Errorf("Divide(%d, %d) = %d; want %d", 
                        test.a, test.b, result, test.expected)
                }
            }
        })
    }
}
```

### Testing with Subtests

Subtests allow you to group related tests and run them individually:

```go
func TestCalculator(t *testing.T) {
    t.Run("addition", func(t *testing.T) {
        result := Add(2, 3)
        if result != 5 {
            t.Errorf("Add(2, 3) = %d; want 5", result)
        }
    })

    t.Run("subtraction", func(t *testing.T) {
        result := Subtract(5, 3)
        if result != 2 {
            t.Errorf("Subtract(5, 3) = %d; want 2", result)
        }
    })

    t.Run("multiplication", func(t *testing.T) {
        result := Multiply(4, 3)
        if result != 12 {
            t.Errorf("Multiply(4, 3) = %d; want 12", result)
        }
    })
}
```

## Benchmark Testing

### Basic Benchmark

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = "hello" + " " + "world"
    }
}
```

### Benchmark with Setup

```go
func BenchmarkProcessData(b *testing.B) {
    // Setup (run once)
    data := generateTestData(1000)
    
    // Reset timer to exclude setup time
    b.ResetTimer()
    
    // Benchmark loop
    for i := 0; i < b.N; i++ {
        ProcessData(data)
    }
}
```

### Benchmark with Different Input Sizes

```go
func BenchmarkProcessData(b *testing.B) {
    benchmarks := []struct {
        name string
        size int
    }{
        {"small", 10},
        {"medium", 100},
        {"large", 1000},
    }

    for _, bm := range benchmarks {
        b.Run(bm.name, func(b *testing.B) {
            data := generateTestData(bm.size)
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                ProcessData(data)
            }
        })
    }
}
```

### Running Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkAdd

# Run benchmarks with memory allocation info
go test -bench=. -benchmem

# Run benchmarks for specific time
go test -bench=. -benchtime=5s
```

## Example Tests

Example tests serve as both tests and documentation:

```go
func ExampleAdd() {
    result := Add(2, 3)
    fmt.Println(result)
    // Output: 5
}

func ExampleCalculator_Add() {
    calc := Calculator{}
    result := calc.Add(2, 3)
    fmt.Println(result)
    // Output: 5
}

func ExampleCalculator_Add_multiple() {
    calc := Calculator{}
    fmt.Println(calc.Add(1, 2))
    fmt.Println(calc.Add(3, 4))
    fmt.Println(calc.Add(5, 6))
    // Output:
    // 3
    // 7
    // 11
}
```

## Testing Utilities

### Test Helpers

```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper() // Marks this function as a test helper
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func assertError(t *testing.T, err error) {
    t.Helper()
    if err == nil {
        t.Error("expected error but got none")
    }
}

func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
}

func TestWithHelpers(t *testing.T) {
    result, err := Divide(10, 2)
    assertNoError(t, err)
    assertEqual(t, result, 5)
}
```

### Test Setup and Teardown

```go
func TestMain(m *testing.M) {
    // Setup
    setup()
    
    // Run tests
    code := m.Run()
    
    // Teardown
    teardown()
    
    // Exit with test result code
    os.Exit(code)
}

func setup() {
    // Initialize test environment
    fmt.Println("Setting up test environment")
}

func teardown() {
    // Clean up test environment
    fmt.Println("Cleaning up test environment")
}
```

## Mocking and Stubbing

### Interface-Based Testing

```go
type DataStore interface {
    Get(id string) (string, error)
    Set(id, value string) error
}

type UserService struct {
    store DataStore
}

func (s *UserService) GetUserName(id string) (string, error) {
    return s.store.Get(id)
}

// Mock implementation for testing
type MockDataStore struct {
    data map[string]string
}

func (m *MockDataStore) Get(id string) (string, error) {
    if value, exists := m.data[id]; exists {
        return value, nil
    }
    return "", fmt.Errorf("user not found")
}

func (m *MockDataStore) Set(id, value string) error {
    m.data[id] = value
    return nil
}

func TestUserService_GetUserName(t *testing.T) {
    mockStore := &MockDataStore{
        data: map[string]string{"1": "Alice", "2": "Bob"},
    }
    
    service := &UserService{store: mockStore}
    
    tests := []struct {
        name        string
        id          string
        expected    string
        expectError bool
    }{
        {"existing user", "1", "Alice", false},
        {"non-existing user", "3", "", true},
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := service.GetUserName(test.id)
            
            if test.expectError {
                if err == nil {
                    t.Error("Expected error but got none")
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if result != test.expected {
                    t.Errorf("Expected %s, got %s", test.expected, result)
                }
            }
        })
    }
}
```

### Using testify/mock

```go
import (
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
)

type MockDataStore struct {
    mock.Mock
}

func (m *MockDataStore) Get(id string) (string, error) {
    args := m.Called(id)
    return args.String(0), args.Error(1)
}

func TestUserService_GetUserName_WithTestify(t *testing.T) {
    mockStore := new(MockDataStore)
    service := &UserService{store: mockStore}
    
    // Set up expectations
    mockStore.On("Get", "1").Return("Alice", nil)
    mockStore.On("Get", "2").Return("", fmt.Errorf("user not found"))
    
    // Test successful case
    result, err := service.GetUserName("1")
    assert.NoError(t, err)
    assert.Equal(t, "Alice", result)
    
    // Test error case
    result, err = service.GetUserName("2")
    assert.Error(t, err)
    assert.Empty(t, result)
    
    // Verify all expectations were met
    mockStore.AssertExpectations(t)
}
```

## Integration Testing

### Testing HTTP Handlers

```go
func TestUserHandler_GetUser(t *testing.T) {
    // Create a test server
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"name": "Alice"})
    }))
    defer ts.Close()
    
    // Make request to test server
    resp, err := http.Get(ts.URL)
    if err != nil {
        t.Fatalf("Failed to make request: %v", err)
    }
    defer resp.Body.Close()
    
    // Assert response
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }
    
    var result map[string]string
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        t.Fatalf("Failed to decode response: %v", err)
    }
    
    if result["name"] != "Alice" {
        t.Errorf("Expected name 'Alice', got %s", result["name"])
    }
}
```

### Testing Database Operations

```go
func TestUserRepository_Create(t *testing.T) {
    // Use test database
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("Failed to open test database: %v", err)
    }
    defer db.Close()
    
    // Create tables
    _, err = db.Exec(`
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL
        )
    `)
    if err != nil {
        t.Fatalf("Failed to create table: %v", err)
    }
    
    repo := &UserRepository{db: db}
    user := &User{Name: "Alice", Email: "alice@example.com"}
    
    // Test creation
    err = repo.Create(user)
    if err != nil {
        t.Errorf("Failed to create user: %v", err)
    }
    
    // Verify user was created
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ?", "Alice").Scan(&count)
    if err != nil {
        t.Fatalf("Failed to query database: %v", err)
    }
    
    if count != 1 {
        t.Errorf("Expected 1 user, got %d", count)
    }
}
```

## Test Coverage

### Coverage Analysis

```bash
# Generate coverage report
go test -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out

# View coverage in terminal
go tool cover -func=coverage.out
```

### Coverage Thresholds

```go
func TestMain(m *testing.M) {
    // Run tests
    code := m.Run()
    
    // Check coverage
    if code == 0 {
        if testing.CoverMode() != "" {
            coverage := testing.Coverage()
            if coverage < 0.8 {
                fmt.Printf("Coverage %.2f%% is below threshold of 80%%\n", coverage*100)
                code = 1
            }
        }
    }
    
    os.Exit(code)
}
```

## Testing Best Practices

### 1. Test Organization

```go
// Group related tests
func TestUserService(t *testing.T) {
    t.Run("create user", func(t *testing.T) {
        // Test user creation
    })
    
    t.Run("get user", func(t *testing.T) {
        // Test user retrieval
    })
    
    t.Run("update user", func(t *testing.T) {
        // Test user update
    })
    
    t.Run("delete user", func(t *testing.T) {
        // Test user deletion
    })
}
```

### 2. Test Data Management

```go
// Use test fixtures
func loadTestData(t *testing.T) []User {
    data, err := os.ReadFile("testdata/users.json")
    if err != nil {
        t.Fatalf("Failed to load test data: %v", err)
    }
    
    var users []User
    if err := json.Unmarshal(data, &users); err != nil {
        t.Fatalf("Failed to unmarshal test data: %v", err)
    }
    
    return users
}

func TestProcessUsers(t *testing.T) {
    users := loadTestData(t)
    // Use test data in tests
}
```

### 3. Parallel Testing

```go
func TestParallel(t *testing.T) {
    t.Parallel() // Run this test in parallel with other tests
    
    // Test implementation
}

func TestParallelGroup(t *testing.T) {
    tests := []struct {
        name string
        input int
        expected int
    }{
        {"test1", 1, 2},
        {"test2", 2, 4},
        {"test3", 3, 6},
    }
    
    for _, test := range tests {
        test := test // Capture loop variable
        t.Run(test.name, func(t *testing.T) {
            t.Parallel() // Run subtests in parallel
            result := Double(test.input)
            if result != test.expected {
                t.Errorf("Double(%d) = %d; want %d", 
                    test.input, result, test.expected)
            }
        })
    }
}
```

### 4. Test Cleanup

```go
func TestWithCleanup(t *testing.T) {
    // Setup
    tempFile, err := os.CreateTemp("", "test_*.txt")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    
    // Cleanup
    t.Cleanup(func() {
        os.Remove(tempFile.Name())
    })
    
    // Test implementation
}
```

## Common Testing Patterns

### 1. Golden File Testing

```go
func TestProcessData_Golden(t *testing.T) {
    input := "test input"
    result := ProcessData(input)
    
    goldenFile := "testdata/expected_output.txt"
    expected, err := os.ReadFile(goldenFile)
    if err != nil {
        t.Fatalf("Failed to read golden file: %v", err)
    }
    
    if string(result) != string(expected) {
        t.Errorf("Output doesn't match golden file")
        // Optionally update golden file
        if *updateGolden {
            os.WriteFile(goldenFile, result, 0644)
        }
    }
}
```

### 2. Property-Based Testing

```go
func TestAdd_Properties(t *testing.T) {
    // Commutative property: a + b = b + a
    for i := 0; i < 100; i++ {
        a := rand.Intn(1000)
        b := rand.Intn(1000)
        if Add(a, b) != Add(b, a) {
            t.Errorf("Add is not commutative: Add(%d, %d) != Add(%d, %d)", 
                a, b, b, a)
        }
    }
    
    // Identity property: a + 0 = a
    for i := 0; i < 100; i++ {
        a := rand.Intn(1000)
        if Add(a, 0) != a {
            t.Errorf("Add identity property failed: Add(%d, 0) != %d", a, a)
        }
    }
}
```

### 3. Fuzz Testing

```go
func FuzzAdd(f *testing.F) {
    f.Add(2, 3)
    f.Fuzz(func(t *testing.T, a, b int) {
        result := Add(a, b)
        // Check that result is reasonable
        if result < a && b > 0 {
            t.Errorf("Add(%d, %d) = %d, expected >= %d", a, b, result, a)
        }
    })
}
```

## Summary

Go testing provides:

- **Built-in testing framework**: No external dependencies required
- **Unit testing**: Test individual functions and methods
- **Benchmark testing**: Measure performance
- **Example testing**: Documentation and testing combined
- **Table-driven tests**: Test multiple scenarios efficiently
- **Mocking support**: Test with dependencies
- **Coverage analysis**: Measure test coverage
- **Parallel testing**: Run tests concurrently

Key points to remember:
1. Use descriptive test names
2. Write table-driven tests for multiple scenarios
3. Test both success and error cases
4. Use interfaces for mocking
5. Measure and maintain good test coverage
6. Use subtests for organization
7. Write benchmarks for performance-critical code
8. Use examples for documentation

Understanding testing is essential for writing reliable, maintainable Go code. 