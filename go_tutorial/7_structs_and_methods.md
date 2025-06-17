# Go Structs and Methods

## Overview

Structs in Go are composite data types that group together variables (fields) under a single name. Methods are functions that are associated with a specific type, allowing you to define behavior for your custom types. Together, structs and methods provide the foundation for object-oriented programming in Go.

## Structs

### What is a Struct?

A struct is a collection of fields, each with a name and a type. Structs are used to group related data together to form a single unit.

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

### Declaring Structs

```go
// Basic struct declaration
type Point struct {
    X, Y int
}

// Struct with different field types
type Rectangle struct {
    Width  float64
    Height float64
    Color  string
}

// Nested struct
type Address struct {
    Street  string
    City    string
    State   string
    ZipCode string
}

type Employee struct {
    Name    string
    ID      int
    Address Address  // Nested struct
}
```

### Creating Struct Instances

```go
// Method 1: Using struct literal
person1 := Person{
    Name: "Alice",
    Age:  30,
    City: "New York",
}

// Method 2: Using field names (recommended for clarity)
person2 := Person{
    Name: "Bob",
    Age:  25,
    City: "Los Angeles",
}

// Method 3: Using positional values (order matters)
person3 := Person{"Charlie", 35, "Chicago"}

// Method 4: Zero-value struct
var person4 Person
fmt.Printf("Zero value: %+v\n", person4) // {Name: Age:0 City:}
```

### Accessing Struct Fields

```go
person := Person{Name: "Alice", Age: 30, City: "New York"}

// Access fields using dot notation
fmt.Printf("Name: %s\n", person.Name)
fmt.Printf("Age: %d\n", person.Age)
fmt.Printf("City: %s\n", person.City)

// Modify fields
person.Age = 31
person.City = "Boston"
```

### Anonymous Structs

```go
// Create a struct without declaring a type
point := struct {
    X, Y int
}{
    X: 10,
    Y: 20,
}

fmt.Printf("Point: %+v\n", point)

// Anonymous struct as function parameter
func printPoint(p struct{ X, Y int }) {
    fmt.Printf("(%d, %d)\n", p.X, p.Y)
}
```

### Struct Tags

Struct tags provide metadata about struct fields, commonly used for JSON serialization, database mapping, and validation.

```go
type User struct {
    ID       int    `json:"id" db:"user_id"`
    Name     string `json:"name" db:"user_name" validate:"required"`
    Email    string `json:"email" db:"email" validate:"email"`
    Password string `json:"-" db:"password_hash"` // "-" means don't include in JSON
}

// Using struct tags with reflection
import "reflect"

func printTags() {
    user := User{}
    t := reflect.TypeOf(user)
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field: %s, JSON tag: %s\n", field.Name, field.Tag.Get("json"))
    }
}
```

## Methods

### What are Methods?

Methods are functions that are associated with a specific type. They allow you to define behavior for your custom types.

```go
type Circle struct {
    Radius float64
}

// Method with value receiver
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// Method with pointer receiver
func (c *Circle) SetRadius(radius float64) {
    c.Radius = radius
}
```

### Method Syntax

```go
func (receiver receiverType) methodName(parameters) returnType {
    // method body
}
```

### Value Receivers vs Pointer Receivers

#### Value Receivers

```go
type Point struct {
    X, Y int
}

// Value receiver - works on a copy of the struct
func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// Value receiver - cannot modify the original struct
func (p Point) Move(dx, dy int) {
    p.X += dx  // This modifies the copy, not the original
    p.Y += dy
}

func main() {
    point := Point{X: 3, Y: 4}
    fmt.Printf("Distance: %.2f\n", point.Distance()) // 5.00
    
    point.Move(1, 1)
    fmt.Printf("After move: %+v\n", point) // {X:3 Y:4} (unchanged)
}
```

#### Pointer Receivers

```go
type Point struct {
    X, Y int
}

// Pointer receiver - can modify the original struct
func (p *Point) Move(dx, dy int) {
    p.X += dx  // This modifies the original struct
    p.Y += dy
}

// Pointer receiver - can also be used for read-only operations
func (p *Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func main() {
    point := Point{X: 3, Y: 4}
    point.Move(1, 1)
    fmt.Printf("After move: %+v\n", point) // {X:4 Y:5} (changed)
}
```

### When to Use Value vs Pointer Receivers

#### Use Value Receivers When:
- The method doesn't need to modify the struct
- The struct is small (few fields)
- You want to work with a copy of the data

```go
type Rectangle struct {
    Width, Height float64
}

// Value receiver - read-only operation
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Value receiver - small struct
func (r Rectangle) IsSquare() bool {
    return r.Width == r.Height
}
```

#### Use Pointer Receivers When:
- The method needs to modify the struct
- The struct is large (many fields)
- You want to avoid copying the struct

```go
type BankAccount struct {
    AccountNumber string
    Balance       float64
    Owner         string
    Transactions  []Transaction
}

// Pointer receiver - modifies the struct
func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount
    b.Transactions = append(b.Transactions, Transaction{
        Type:   "deposit",
        Amount: amount,
    })
}

// Pointer receiver - large struct (avoid copying)
func (b *BankAccount) GetBalance() float64 {
    return b.Balance
}
```

### Method Chaining

```go
type StringBuilder struct {
    parts []string
}

func (sb *StringBuilder) Append(s string) *StringBuilder {
    sb.parts = append(sb.parts, s)
    return sb
}

func (sb *StringBuilder) String() string {
    return strings.Join(sb.parts, "")
}

func main() {
    sb := &StringBuilder{}
    result := sb.Append("Hello").Append(" ").Append("World").String()
    fmt.Println(result) // "Hello World"
}
```

## Struct Composition

### Embedding Structs

Go uses composition instead of inheritance. You can embed one struct within another to inherit its fields and methods.

```go
type Animal struct {
    Name string
    Age  int
}

func (a Animal) Describe() string {
    return fmt.Sprintf("%s is %d years old", a.Name, a.Age)
}

type Dog struct {
    Animal      // Embedded struct
    Breed string
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Buddy", Age: 3},
        Breed:  "Golden Retriever",
    }
    
    fmt.Println(dog.Name)           // "Buddy" (inherited field)
    fmt.Println(dog.Describe())     // "Buddy is 3 years old" (inherited method)
    fmt.Println(dog.Breed)          // "Golden Retriever"
}
```

### Method Overriding

```go
type Animal struct {
    Name string
    Age  int
}

func (a Animal) MakeSound() string {
    return "Some sound"
}

type Dog struct {
    Animal
    Breed string
}

// Override the MakeSound method
func (d Dog) MakeSound() string {
    return "Woof!"
}

type Cat struct {
    Animal
    Color string
}

// Override the MakeSound method
func (c Cat) MakeSound() string {
    return "Meow!"
}

func main() {
    dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden"}
    cat := Cat{Animal: Animal{Name: "Whiskers"}, Color: "Orange"}
    
    fmt.Println(dog.MakeSound())  // "Woof!"
    fmt.Println(cat.MakeSound())  // "Meow!"
}
```

### Multiple Embedding

```go
type Reader struct {
    Name string
}

func (r Reader) Read() string {
    return fmt.Sprintf("%s is reading", r.Name)
}

type Writer struct {
    Name string
}

func (w Writer) Write() string {
    return fmt.Sprintf("%s is writing", w.Name)
}

type ReaderWriter struct {
    Reader
    Writer
}

func main() {
    rw := ReaderWriter{
        Reader: Reader{Name: "Alice"},
        Writer: Writer{Name: "Alice"},
    }
    
    fmt.Println(rw.Read())   // "Alice is reading"
    fmt.Println(rw.Write())  // "Alice is writing"
}
```

### Interface Implementation Through Embedding

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write() string
}

type ReaderWriter interface {
    Reader
    Writer
}

type File struct {
    Name string
}

func (f File) Read() string {
    return fmt.Sprintf("Reading from %s", f.Name)
}

func (f File) Write() string {
    return fmt.Sprintf("Writing to %s", f.Name)
}

// File automatically implements ReaderWriter through embedding
func process(rw ReaderWriter) {
    fmt.Println(rw.Read())
    fmt.Println(rw.Write())
}
```

## Advanced Struct Features

### Private vs Public Fields

```go
type Person struct {
    Name    string  // Public (exported) - accessible from other packages
    age     int     // Private (unexported) - only accessible within the same package
    Address string  // Public
}

func (p *Person) SetAge(age int) {
    if age >= 0 {
        p.age = age  // Can access private field within the same package
    }
}

func (p Person) GetAge() int {
    return p.age  // Can access private field within the same package
}
```

### Struct Methods with Interfaces

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
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

### Constructor Functions

```go
type Person struct {
    Name string
    Age  int
}

// Constructor function
func NewPerson(name string, age int) *Person {
    if age < 0 {
        age = 0
    }
    return &Person{
        Name: name,
        Age:  age,
    }
}

// Constructor with validation
func NewPersonWithValidation(name string, age int) (*Person, error) {
    if name == "" {
        return nil, fmt.Errorf("name cannot be empty")
    }
    if age < 0 {
        return nil, fmt.Errorf("age cannot be negative")
    }
    
    return &Person{
        Name: name,
        Age:  age,
    }, nil
}

func main() {
    person1 := NewPerson("Alice", 30)
    person2, err := NewPersonWithValidation("", -5)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Common Patterns

### Builder Pattern

```go
type PersonBuilder struct {
    person Person
}

func NewPersonBuilder() *PersonBuilder {
    return &PersonBuilder{}
}

func (pb *PersonBuilder) Name(name string) *PersonBuilder {
    pb.person.Name = name
    return pb
}

func (pb *PersonBuilder) Age(age int) *PersonBuilder {
    pb.person.Age = age
    return pb
}

func (pb *PersonBuilder) City(city string) *PersonBuilder {
    pb.person.City = city
    return pb
}

func (pb *PersonBuilder) Build() Person {
    return pb.person
}

func main() {
    person := NewPersonBuilder().
        Name("Alice").
        Age(30).
        City("New York").
        Build()
    
    fmt.Printf("Person: %+v\n", person)
}
```

### Factory Pattern

```go
type Animal interface {
    MakeSound() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) MakeSound() string { return "Woof!" }
func (c Cat) MakeSound() string { return "Meow!" }

func NewAnimal(animalType string) (Animal, error) {
    switch animalType {
    case "dog":
        return Dog{}, nil
    case "cat":
        return Cat{}, nil
    default:
        return nil, fmt.Errorf("unknown animal type: %s", animalType)
    }
}

func main() {
    dog, _ := NewAnimal("dog")
    cat, _ := NewAnimal("cat")
    
    fmt.Println(dog.MakeSound())  // "Woof!"
    fmt.Println(cat.MakeSound())  // "Meow!"
}
```

## Best Practices

### 1. Use Meaningful Field Names

```go
// Good
type User struct {
    ID       int
    Username string
    Email    string
    Created  time.Time
}

// Avoid
type User struct {
    I int
    U string
    E string
    C time.Time
}
```

### 2. Group Related Fields

```go
type Employee struct {
    // Personal information
    Name     string
    Age      int
    Email    string
    
    // Work information
    ID       int
    Position string
    Salary   float64
    
    // Address information
    Street   string
    City     string
    State    string
    ZipCode  string
}
```

### 3. Use Pointer Receivers Appropriately

```go
type Counter struct {
    count int
}

// Use pointer receiver when modifying the struct
func (c *Counter) Increment() {
    c.count++
}

// Use value receiver when only reading
func (c Counter) GetCount() int {
    return c.count
}
```

### 4. Implement String() Method for Debugging

```go
type Point struct {
    X, Y int
}

func (p Point) String() string {
    return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

func main() {
    point := Point{X: 3, Y: 4}
    fmt.Println(point) // "Point(3, 4)"
}
```

### 5. Use Struct Tags for Serialization

```go
type User struct {
    ID       int       `json:"id" xml:"id"`
    Name     string    `json:"name" xml:"name"`
    Email    string    `json:"email" xml:"email"`
    Password string    `json:"-" xml:"-"` // Don't include in JSON/XML
    Created  time.Time `json:"created_at" xml:"created"`
}
```

## Performance Considerations

### 1. Struct Field Order

```go
// Good: Group fields by size to minimize padding
type OptimizedStruct struct {
    A int64   // 8 bytes
    B int64   // 8 bytes
    C int32   // 4 bytes
    D int16   // 2 bytes
    E int8    // 1 byte
    // Total: 24 bytes
}

// Bad: Fields not grouped by size
type UnoptimizedStruct struct {
    A int8    // 1 byte
    B int64   // 8 bytes (7 bytes padding)
    C int16   // 2 bytes (6 bytes padding)
    D int64   // 8 bytes (6 bytes padding)
    E int32   // 4 bytes (4 bytes padding)
    // Total: 32 bytes
}
```

### 2. Method Receiver Choice

```go
// Use value receiver for small structs
type Point struct {
    X, Y int
}

func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// Use pointer receiver for large structs
type LargeStruct struct {
    Data [1000]int
    // ... many fields
}

func (ls *LargeStruct) Process() {
    // Process data
}
```

## Summary

Structs and methods in Go provide:

- **Data Organization**: Group related data together
- **Behavior Definition**: Define methods for custom types
- **Composition**: Use embedding instead of inheritance
- **Type Safety**: Compile-time type checking
- **Performance**: Efficient memory layout and method calls

Key points to remember:
1. Use structs to group related data
2. Use methods to define behavior for your types
3. Choose between value and pointer receivers based on your needs
4. Use composition (embedding) instead of inheritance
5. Implement interfaces through method sets
6. Use struct tags for metadata and serialization
7. Follow naming conventions and best practices

Structs and methods are fundamental to Go programming and provide the building blocks for creating well-structured, maintainable code. 