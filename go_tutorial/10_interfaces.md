# Go Interfaces

## Overview

Interfaces in Go are a powerful feature that enables polymorphism and code decoupling. They define a set of method signatures that types can implement, allowing you to write code that works with any type that satisfies the interface.

## What are Interfaces?

An interface is a type that defines a set of method signatures. Any type that implements all the methods of an interface automatically satisfies that interface.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

## Interface Declaration

### Basic Interface

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}
```

### Interface with Multiple Methods

```go
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}

type ReadWriteCloser interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
}
```

## Interface Implementation

### Implicit Implementation

In Go, interfaces are implemented implicitly. A type implements an interface by implementing all its methods.

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Circle automatically implements Shape interface
```

### Complete Example

```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}
    
    printShapeInfo(circle)     // Area: 78.54, Perimeter: 31.42
    printShapeInfo(rectangle)  // Area: 24.00, Perimeter: 20.00
}
```

## Interface Composition

### Embedding Interfaces

Interfaces can be composed by embedding other interfaces.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Compose interfaces
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

### Interface Composition Example

```go
type Animal interface {
    MakeSound() string
    GetName() string
}

type Walker interface {
    Walk() string
}

type Swimmer interface {
    Swim() string
}

type FlyingAnimal interface {
    Animal
    Walker
    Fly() string
}

type Dog struct {
    Name string
}

func (d Dog) MakeSound() string {
    return "Woof!"
}

func (d Dog) GetName() string {
    return d.Name
}

func (d Dog) Walk() string {
    return "Walking on four legs"
}

// Dog implements Animal and Walker interfaces
```

## Empty Interface

### What is an Empty Interface?

The empty interface `interface{}` (or `any` in Go 1.18+) has no methods and is satisfied by any type.

```go
func PrintAnything(v interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func main() {
    PrintAnything(42)           // Value: 42, Type: int
    PrintAnything("hello")      // Value: hello, Type: string
    PrintAnything(true)         // Value: true, Type: bool
    PrintAnything([]int{1,2,3}) // Value: [1 2 3], Type: []int
}
```

### Type Assertions

```go
func processValue(v interface{}) {
    // Type assertion
    if str, ok := v.(string); ok {
        fmt.Printf("String: %s\n", str)
    } else if num, ok := v.(int); ok {
        fmt.Printf("Number: %d\n", num)
    } else {
        fmt.Printf("Unknown type: %T\n", v)
    }
}

// Type switch
func processValueWithSwitch(v interface{}) {
    switch val := v.(type) {
    case string:
        fmt.Printf("String: %s\n", val)
    case int:
        fmt.Printf("Number: %d\n", val)
    case bool:
        fmt.Printf("Boolean: %t\n", val)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

## Interface Best Practices

### 1. Keep Interfaces Small

```go
// Good: Small, focused interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Avoid: Large interface with many methods
type BigInterface interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
    Flush() error
    Seek(offset int64, whence int) (int64, error)
    // ... many more methods
}
```

### 2. Define Interfaces Where They Are Used

```go
// Good: Define interface in the package that uses it
package client

type DataProcessor interface {
    Process(data []byte) error
}

func ProcessData(processor DataProcessor, data []byte) error {
    return processor.Process(data)
}
```

### 3. Use Interface Composition

```go
// Compose small interfaces into larger ones
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

## Common Interface Patterns

### 1. The io.Reader Interface

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Example implementation
type StringReader struct {
    data string
    pos  int
}

func NewStringReader(data string) *StringReader {
    return &StringReader{data: data, pos: 0}
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
    if sr.pos >= len(sr.data) {
        return 0, io.EOF
    }
    
    n = copy(p, sr.data[sr.pos:])
    sr.pos += n
    return n, nil
}
```

### 2. The fmt.Stringer Interface

```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    fmt.Println(person) // Alice (30 years old)
}
```

### 3. The error Interface

```go
type error interface {
    Error() string
}

// Custom error type
type ValidationError struct {
    Field string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return ValidationError{Field: "age", Message: "age cannot be negative"}
    }
    if age > 150 {
        return ValidationError{Field: "age", Message: "age cannot exceed 150"}
    }
    return nil
}
```

## Interface vs Concrete Types

### When to Use Interfaces

```go
// Use interfaces when you want to:
// 1. Accept any type that implements certain behavior
func ProcessData(processor DataProcessor, data []byte) error {
    return processor.Process(data)
}

// 2. Test with mocks
type MockProcessor struct{}

func (m MockProcessor) Process(data []byte) error {
    return nil
}

// 3. Provide multiple implementations
type JSONProcessor struct{}
type XMLProcessor struct{}

func (j JSONProcessor) Process(data []byte) error {
    // Process JSON
    return nil
}

func (x XMLProcessor) Process(data []byte) error {
    // Process XML
    return nil
}
```

### When to Use Concrete Types

```go
// Use concrete types when:
// 1. You need specific behavior
type Database struct {
    host string
    port int
}

func (db *Database) Connect() error {
    // Specific database connection logic
    return nil
}

// 2. Performance is critical
func ProcessNumbers(numbers []int) int {
    // Direct slice operations are faster than interface calls
    sum := 0
    for _, n := range numbers {
        sum += n
    }
    return sum
}
```

## Interface Design Patterns

### 1. Strategy Pattern

```go
type PaymentStrategy interface {
    Pay(amount float64) error
}

type CreditCardPayment struct{}
type PayPalPayment struct{}
type BankTransferPayment struct{}

func (c CreditCardPayment) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using credit card\n", amount)
    return nil
}

func (p PayPalPayment) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using PayPal\n", amount)
    return nil
}

func (b BankTransferPayment) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using bank transfer\n", amount)
    return nil
}

type PaymentProcessor struct {
    strategy PaymentStrategy
}

func (pp *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
    pp.strategy = strategy
}

func (pp *PaymentProcessor) ProcessPayment(amount float64) error {
    return pp.strategy.Pay(amount)
}
```

### 2. Factory Pattern

```go
type Animal interface {
    MakeSound() string
}

type Dog struct{}
type Cat struct{}
type Bird struct{}

func (d Dog) MakeSound() string { return "Woof!" }
func (c Cat) MakeSound() string { return "Meow!" }
func (b Bird) MakeSound() string { return "Tweet!" }

func NewAnimal(animalType string) (Animal, error) {
    switch animalType {
    case "dog":
        return Dog{}, nil
    case "cat":
        return Cat{}, nil
    case "bird":
        return Bird{}, nil
    default:
        return nil, fmt.Errorf("unknown animal type: %s", animalType)
    }
}
```

### 3. Observer Pattern

```go
type Observer interface {
    Update(message string)
}

type Subject interface {
    Attach(observer Observer)
    Detach(observer Observer)
    Notify(message string)
}

type NewsAgency struct {
    observers []Observer
}

func (na *NewsAgency) Attach(observer Observer) {
    na.observers = append(na.observers, observer)
}

func (na *NewsAgency) Detach(observer Observer) {
    // Implementation to remove observer
}

func (na *NewsAgency) Notify(message string) {
    for _, observer := range na.observers {
        observer.Update(message)
    }
}

type NewsChannel struct {
    name string
}

func (nc NewsChannel) Update(message string) {
    fmt.Printf("%s received news: %s\n", nc.name, message)
}
```

## Interface Testing

### Testing with Interfaces

```go
type DataStore interface {
    Get(id string) (string, error)
    Set(id, value string) error
}

type MemoryStore struct {
    data map[string]string
}

func NewMemoryStore() *MemoryStore {
    return &MemoryStore{data: make(map[string]string)}
}

func (m *MemoryStore) Get(id string) (string, error) {
    if value, exists := m.data[id]; exists {
        return value, nil
    }
    return "", fmt.Errorf("key not found: %s", id)
}

func (m *MemoryStore) Set(id, value string) error {
    m.data[id] = value
    return nil
}

// Service that uses DataStore
type UserService struct {
    store DataStore
}

func NewUserService(store DataStore) *UserService {
    return &UserService{store: store}
}

func (us *UserService) GetUserName(id string) (string, error) {
    return us.store.Get(id)
}

// Test with mock
type MockStore struct {
    data map[string]string
}

func (m *MockStore) Get(id string) (string, error) {
    return m.data[id], nil
}

func (m *MockStore) Set(id, value string) error {
    m.data[id] = value
    return nil
}

func TestUserService(t *testing.T) {
    mockStore := &MockStore{data: map[string]string{"1": "Alice"}}
    service := NewUserService(mockStore)
    
    name, err := service.GetUserName("1")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if name != "Alice" {
        t.Errorf("Expected Alice, got %s", name)
    }
}
```

## Interface Performance

### Interface Overhead

```go
// Interface calls have a small overhead compared to direct method calls
type Calculator interface {
    Add(a, b int) int
}

type SimpleCalculator struct{}

func (sc SimpleCalculator) Add(a, b int) int {
    return a + b
}

// Direct call (faster)
calc := SimpleCalculator{}
result := calc.Add(5, 3)

// Interface call (slightly slower)
var calcInterface Calculator = SimpleCalculator{}
result = calcInterface.Add(5, 3)
```

### Reducing Interface Overhead

```go
// Use concrete types when performance is critical
func ProcessNumbersDirect(numbers []int) int {
    sum := 0
    for _, n := range numbers {
        sum += n
    }
    return sum
}

// Use interfaces when flexibility is needed
func ProcessNumbersInterface(processor NumberProcessor, numbers []int) int {
    sum := 0
    for _, n := range numbers {
        sum = processor.Add(sum, n)
    }
    return sum
}
```

## Common Standard Library Interfaces

### 1. io.Reader and io.Writer

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Example usage
func CopyData(src io.Reader, dst io.Writer) error {
    buffer := make([]byte, 1024)
    for {
        n, err := src.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        
        _, err = dst.Write(buffer[:n])
        if err != nil {
            return err
        }
    }
    return nil
}
```

### 2. fmt.Stringer

```go
type Stringer interface {
    String() string
}

type Point struct {
    X, Y int
}

func (p Point) String() string {
    return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}
```

### 3. sort.Interface

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

type Person struct {
    Name string
    Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }
    
    sort.Sort(ByAge(people))
    fmt.Printf("Sorted by age: %+v\n", people)
}
```

## Summary

Interfaces in Go provide:

- **Polymorphism**: Write code that works with multiple types
- **Decoupling**: Separate interface from implementation
- **Testability**: Easy to mock and test
- **Flexibility**: Change implementations without changing client code
- **Composition**: Build complex interfaces from simple ones

Key points to remember:
1. Interfaces are implemented implicitly
2. Keep interfaces small and focused
3. Define interfaces where they are used
4. Use interface composition for complex interfaces
5. Use concrete types when performance is critical
6. Leverage standard library interfaces
7. Use interfaces for testing and mocking

Understanding interfaces is crucial for writing flexible, maintainable, and testable Go code. 