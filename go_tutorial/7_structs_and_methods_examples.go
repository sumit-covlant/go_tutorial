package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

func main() {
	fmt.Println("=== Go Structs and Methods Examples ===\n")

	// Basic struct operations
	basicStructOperations()

	// Struct methods
	structMethods()

	// Struct composition
	structComposition()

	// Method overriding
	methodOverriding()

	// Multiple embedding
	multipleEmbedding()

	// Interface implementation
	interfaceImplementation()

	// Struct tags
	structTags()

	// Constructor functions
	constructorFunctions()

	// Common patterns
	commonPatterns()

	// Best practices
	bestPractices()

	// Performance considerations
	performanceConsiderations()
}

func basicStructOperations() {
	fmt.Println("1. Basic Struct Operations")
	fmt.Println("--------------------------")

	// Basic struct declaration
	type Point struct {
		X, Y int
	}

	// Creating struct instances
	point1 := Point{X: 10, Y: 20}
	fmt.Printf("Point1: %+v\n", point1)

	point2 := Point{30, 40}
	fmt.Printf("Point2: %+v\n", point2)

	// Zero-value struct
	var point3 Point
	fmt.Printf("Zero value Point: %+v\n", point3)

	// Accessing and modifying fields
	point1.X = 15
	point1.Y = 25
	fmt.Printf("Modified Point1: %+v\n", point1)

	// Anonymous structs
	person := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("Anonymous struct: %+v\n", person)
	fmt.Println()
}

func structMethods() {
	fmt.Println("2. Struct Methods")
	fmt.Println("-----------------")

	// Circle with methods
	circle := Circle{Radius: 5}
	fmt.Printf("Circle radius: %.2f\n", circle.Radius)
	fmt.Printf("Circle area: %.2f\n", circle.Area())
	fmt.Printf("Circle perimeter: %.2f\n", circle.Perimeter())

	// Modifying through pointer receiver
	circle.SetRadius(10)
	fmt.Printf("After setting radius: %.2f\n", circle.Radius)
	fmt.Printf("New area: %.2f\n", circle.Area())

	// Point with value vs pointer receivers
	point := Point{X: 3, Y: 4}
	fmt.Printf("Point: %+v\n", point)
	fmt.Printf("Distance: %.2f\n", point.Distance())

	// Value receiver doesn't modify original
	point.Move(1, 1)
	fmt.Printf("After Move (value receiver): %+v\n", point)

	// Pointer receiver modifies original
	pointPtr := &point
	pointPtr.MovePointer(1, 1)
	fmt.Printf("After MovePointer (pointer receiver): %+v\n", point)
	fmt.Println()
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

func (c *Circle) SetRadius(radius float64) {
	c.Radius = radius
}

type Point struct {
	X, Y int
}

func (p Point) Distance() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p Point) Move(dx, dy int) {
	p.X += dx // This modifies the copy, not the original
	p.Y += dy
}

func (p *Point) MovePointer(dx, dy int) {
	p.X += dx // This modifies the original
	p.Y += dy
}

func structComposition() {
	fmt.Println("3. Struct Composition")
	fmt.Println("----------------------")

	// Basic embedding
	dog := Dog{
		Animal: Animal{Name: "Buddy", Age: 3},
		Breed:  "Golden Retriever",
	}

	fmt.Printf("Dog: %+v\n", dog)
	fmt.Printf("Dog name: %s\n", dog.Name)              // Inherited field
	fmt.Printf("Dog description: %s\n", dog.Describe()) // Inherited method
	fmt.Printf("Dog breed: %s\n", dog.Breed)            // Own field

	// Nested structs
	employee := Employee{
		Name: "John Doe",
		ID:   12345,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
	}

	fmt.Printf("Employee: %+v\n", employee)
	fmt.Printf("Employee city: %s\n", employee.Address.City)
	fmt.Println()
}

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Describe() string {
	return fmt.Sprintf("%s is %d years old", a.Name, a.Age)
}

type Dog struct {
	Animal // Embedded struct
	Breed  string
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Employee struct {
	Name    string
	ID      int
	Address Address // Nested struct
}

func methodOverriding() {
	fmt.Println("4. Method Overriding")
	fmt.Println("---------------------")

	// Animals with different sounds
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden",
	}

	cat := Cat{
		Animal: Animal{Name: "Whiskers"},
		Color:  "Orange",
	}

	fmt.Printf("Dog sound: %s\n", dog.MakeSound())
	fmt.Printf("Cat sound: %s\n", cat.MakeSound())

	// Using interface
	animals := []AnimalInterface{dog, cat}
	for _, animal := range animals {
		fmt.Printf("%s says: %s\n", animal.GetName(), animal.MakeSound())
	}
	fmt.Println()
}

type AnimalInterface interface {
	MakeSound() string
	GetName() string
}

func (a Animal) MakeSound() string {
	return "Some sound"
}

func (a Animal) GetName() string {
	return a.Name
}

type Cat struct {
	Animal
	Color string
}

func (d Dog) MakeSound() string {
	return "Woof!"
}

func (c Cat) MakeSound() string {
	return "Meow!"
}

func multipleEmbedding() {
	fmt.Println("5. Multiple Embedding")
	fmt.Println("----------------------")

	// ReaderWriter with multiple embedding
	rw := ReaderWriter{
		Reader: Reader{Name: "Alice"},
		Writer: Writer{Name: "Alice"},
	}

	fmt.Printf("ReaderWriter read: %s\n", rw.Read())
	fmt.Printf("ReaderWriter write: %s\n", rw.Write())

	// Interface implementation through embedding
	processReaderWriter(rw)
	fmt.Println()
}

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

func processReaderWriter(rw ReaderWriter) {
	fmt.Printf("Processing: %s and %s\n", rw.Read(), rw.Write())
}

func interfaceImplementation() {
	fmt.Println("6. Interface Implementation")
	fmt.Println("----------------------------")

	// Shapes implementing Shape interface
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		fmt.Printf("Shape area: %.2f, perimeter: %.2f\n",
			shape.Area(), shape.Perimeter())
	}

	// Using interface methods
	printShapeInfo(circle)
	printShapeInfo(rectangle)
	fmt.Println()
}

type Shape interface {
	Area() float64
	Perimeter() float64
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

func structTags() {
	fmt.Println("7. Struct Tags")
	fmt.Println("---------------")

	// User with struct tags
	user := User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret123",
		Created:  time.Now(),
	}

	fmt.Printf("User: %+v\n", user)

	// Print JSON tags
	printTags()
	fmt.Println()
}

type User struct {
	ID       int       `json:"id" xml:"id"`
	Name     string    `json:"name" xml:"name"`
	Email    string    `json:"email" xml:"email"`
	Password string    `json:"-" xml:"-"` // Don't include in JSON/XML
	Created  time.Time `json:"created_at" xml:"created"`
}

func printTags() {
	user := User{}
	t := reflect.TypeOf(user)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s, JSON tag: %s\n", field.Name, field.Tag.Get("json"))
	}
}

func constructorFunctions() {
	fmt.Println("8. Constructor Functions")
	fmt.Println("-------------------------")

	// Using constructor functions
	person1 := NewPerson("Alice", 30)
	fmt.Printf("Person1: %+v\n", person1)

	person2, err := NewPersonWithValidation("", -5)
	if err != nil {
		fmt.Printf("Error creating person2: %v\n", err)
	} else {
		fmt.Printf("Person2: %+v\n", person2)
	}

	person3, err := NewPersonWithValidation("Bob", 25)
	if err != nil {
		fmt.Printf("Error creating person3: %v\n", err)
	} else {
		fmt.Printf("Person3: %+v\n", person3)
	}
	fmt.Println()
}

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	if age < 0 {
		age = 0
	}
	return &Person{
		Name: name,
		Age:  age,
	}
}

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

func commonPatterns() {
	fmt.Println("9. Common Patterns")
	fmt.Println("-------------------")

	// Builder pattern
	person := NewPersonBuilder().
		Name("Alice").
		Age(30).
		City("New York").
		Build()

	fmt.Printf("Built person: %+v\n", person)

	// Factory pattern
	dog, _ := NewAnimal("dog")
	cat, _ := NewAnimal("cat")

	fmt.Printf("Dog sound: %s\n", dog.MakeSound())
	fmt.Printf("Cat sound: %s\n", cat.MakeSound())

	// String method for debugging
	point := Point{X: 3, Y: 4}
	fmt.Printf("Point: %s\n", point)
	fmt.Println()
}

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
	// Note: Person struct doesn't have City field, but we can add it
	// For demonstration, we'll just print it
	fmt.Printf("Setting city to: %s\n", city)
	return pb
}

func (pb *PersonBuilder) Build() Person {
	return pb.person
}

type AnimalFactory interface {
	MakeSound() string
}

type DogFactory struct{}
type CatFactory struct{}

func (d DogFactory) MakeSound() string { return "Woof!" }
func (c CatFactory) MakeSound() string { return "Meow!" }

func NewAnimal(animalType string) (AnimalFactory, error) {
	switch animalType {
	case "dog":
		return DogFactory{}, nil
	case "cat":
		return CatFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown animal type: %s", animalType)
	}
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

func bestPractices() {
	fmt.Println("10. Best Practices")
	fmt.Println("-------------------")

	// Use meaningful field names
	user := User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret",
		Created:  time.Now(),
	}
	fmt.Printf("User with meaningful names: %+v\n", user)

	// Group related fields
	employee := Employee{
		Name: "John",
		ID:   123,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
	}
	fmt.Printf("Employee with grouped fields: %+v\n", employee)

	// Use pointer receivers appropriately
	counter := &Counter{count: 0}
	counter.Increment()
	counter.Increment()
	fmt.Printf("Counter value: %d\n", counter.GetCount())

	// Implement String() method
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %s\n", rect)
	fmt.Println()
}

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c Counter) GetCount() int {
	return c.count
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

func performanceConsiderations() {
	fmt.Println("11. Performance Considerations")
	fmt.Println("-------------------------------")

	// Struct field order optimization
	optimized := OptimizedStruct{
		A: 1, B: 2, C: 3, D: 4, E: 5,
	}
	fmt.Printf("Optimized struct: %+v\n", optimized)

	// Method receiver choice
	smallPoint := Point{X: 3, Y: 4}
	fmt.Printf("Small point distance: %.2f\n", smallPoint.Distance())

	largeStruct := &LargeStruct{Data: [1000]int{1, 2, 3}}
	largeStruct.Process()

	fmt.Println("Performance considerations completed.")
	fmt.Println()
}

type OptimizedStruct struct {
	A int64 // 8 bytes
	B int64 // 8 bytes
	C int32 // 4 bytes
	D int16 // 2 bytes
	E int8  // 1 byte
}

type LargeStruct struct {
	Data [1000]int
}

func (ls *LargeStruct) Process() {
	fmt.Printf("Processing large struct with %d elements\n", len(ls.Data))
}
