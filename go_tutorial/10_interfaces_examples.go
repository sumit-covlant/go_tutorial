package main

import (
	"fmt"
	"math"
	"sort"
)

// This file demonstrates Go interfaces concepts

func main() {
	fmt.Println("=== Go Interfaces Examples ===\n")

	// Basic interface examples
	basicInterfaceExamples()

	// Interface implementation examples
	interfaceImplementationExamples()

	// Interface composition examples
	interfaceCompositionExamples()

	// Empty interface examples
	emptyInterfaceExamples()

	// Interface best practices
	interfaceBestPractices()

	// Common interface patterns
	commonInterfacePatterns()

	// Interface vs concrete types
	interfaceVsConcreteTypes()

	// Interface design patterns
	interfaceDesignPatterns()

	// Interface testing
	interfaceTesting()

	// Interface performance
	interfacePerformance()

	// Standard library interfaces
	standardLibraryInterfaces()
}

// Basic interface examples
func basicInterfaceExamples() {
	fmt.Println("1. Basic Interface Examples")
	fmt.Println("---------------------------")

	// Define interfaces
	fmt.Println("Defining interfaces:")
	fmt.Println("type Shape interface {")
	fmt.Println("    Area() float64")
	fmt.Println("    Perimeter() float64")
	fmt.Println("}")

	// Demonstrate interface usage
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d: Area=%.2f, Perimeter=%.2f\n",
			i+1, shape.Area(), shape.Perimeter())
	}
	fmt.Println()
}

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle implements Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle implements Shape
type Triangle struct {
	SideA, SideB, SideC float64
}

func (t Triangle) Area() float64 {
	// Heron's formula
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// Interface implementation examples
func interfaceImplementationExamples() {
	fmt.Println("2. Interface Implementation Examples")
	fmt.Println("------------------------------------")

	// Demonstrate implicit implementation
	fmt.Println("Implicit implementation:")
	fmt.Println("- No explicit declaration needed")
	fmt.Println("- Just implement all methods")
	fmt.Println("- Circle automatically implements Shape")

	// Demonstrate multiple interface implementation
	dog := Dog{Name: "Buddy"}
	fmt.Printf("Dog: %s\n", dog.MakeSound())
	fmt.Printf("Dog walking: %s\n", dog.Walk())

	// Demonstrate interface satisfaction
	var animal Animal = dog
	var walker Walker = dog
	fmt.Printf("Animal sound: %s\n", animal.MakeSound())
	fmt.Printf("Walker: %s\n", walker.Walk())
	fmt.Println()
}

// Animal interface
type Animal interface {
	MakeSound() string
	GetName() string
}

// Walker interface
type Walker interface {
	Walk() string
}

// Swimmer interface
type Swimmer interface {
	Swim() string
}

// FlyingAnimal interface (composition)
type FlyingAnimal interface {
	Animal
	Walker
	Fly() string
}

// Dog implements Animal and Walker
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

// Interface composition examples
func interfaceCompositionExamples() {
	fmt.Println("3. Interface Composition Examples")
	fmt.Println("---------------------------------")

	// Demonstrate interface embedding
	fmt.Println("Interface composition:")
	fmt.Println("type ReadWriter interface {")
	fmt.Println("    Reader")
	fmt.Println("    Writer")
	fmt.Println("}")

	// Demonstrate composition usage
	bird := Bird{Name: "Sparrow"}
	fmt.Printf("Bird: %s\n", bird.MakeSound())
	fmt.Printf("Bird walking: %s\n", bird.Walk())
	fmt.Printf("Bird flying: %s\n", bird.Fly())

	// Demonstrate interface composition
	var flyingAnimal FlyingAnimal = bird
	fmt.Printf("Flying animal sound: %s\n", flyingAnimal.MakeSound())
	fmt.Printf("Flying animal walk: %s\n", flyingAnimal.Walk())
	fmt.Printf("Flying animal fly: %s\n", flyingAnimal.Fly())
	fmt.Println()
}

// Bird implements FlyingAnimal
type Bird struct {
	Name string
}

func (b Bird) MakeSound() string {
	return "Tweet!"
}

func (b Bird) GetName() string {
	return b.Name
}

func (b Bird) Walk() string {
	return "Hopping on two legs"
}

func (b Bird) Fly() string {
	return "Flying through the air"
}

// Empty interface examples
func emptyInterfaceExamples() {
	fmt.Println("4. Empty Interface Examples")
	fmt.Println("---------------------------")

	// Demonstrate empty interface
	printAnything(42)
	printAnything("hello")
	printAnything(true)
	printAnything([]int{1, 2, 3})

	// Demonstrate type assertions
	processValue(42)
	processValue("hello")
	processValue(true)

	// Demonstrate type switch
	processValueWithSwitch(42)
	processValueWithSwitch("hello")
	processValueWithSwitch(true)
	processValueWithSwitch(3.14)
	fmt.Println()
}

// Empty interface function
func printAnything(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

// Type assertion function
func processValue(v interface{}) {
	if str, ok := v.(string); ok {
		fmt.Printf("String: %s\n", str)
	} else if num, ok := v.(int); ok {
		fmt.Printf("Number: %d\n", num)
	} else {
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// Type switch function
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

// Interface best practices
func interfaceBestPractices() {
	fmt.Println("5. Interface Best Practices")
	fmt.Println("---------------------------")

	// Demonstrate small interfaces
	fmt.Println("Small interfaces:")
	reader := StringReader{data: "hello world", pos: 0}
	buffer := make([]byte, 5)
	n, err := reader.Read(buffer)
	fmt.Printf("Read %d bytes: %s, error: %v\n", n, buffer[:n], err)

	// Demonstrate interface composition
	fmt.Println("\nInterface composition:")
	readWriter := &StringReadWriter{data: "hello world", pos: 0}
	readWriter.Write([]byte(" goodbye"))
	fmt.Printf("Data: %s\n", readWriter.data)
	fmt.Println()
}

// Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer interface
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter interface (composition)
type ReadWriter interface {
	Reader
	Writer
}

// StringReader implements Reader
type StringReader struct {
	data string
	pos  int
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.pos >= len(sr.data) {
		return 0, fmt.Errorf("end of data")
	}

	n = copy(p, sr.data[sr.pos:])
	sr.pos += n
	return n, nil
}

// StringReadWriter implements ReadWriter
type StringReadWriter struct {
	data string
	pos  int
}

func (srw *StringReadWriter) Read(p []byte) (n int, err error) {
	if srw.pos >= len(srw.data) {
		return 0, fmt.Errorf("end of data")
	}

	n = copy(p, srw.data[srw.pos:])
	srw.pos += n
	return n, nil
}

func (srw *StringReadWriter) Write(p []byte) (n int, err error) {
	srw.data += string(p)
	return len(p), nil
}

// Common interface patterns
func commonInterfacePatterns() {
	fmt.Println("6. Common Interface Patterns")
	fmt.Println("----------------------------")

	// Demonstrate Stringer interface
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %s\n", person)

	// Demonstrate custom error interface
	err := ValidationError{Field: "age", Message: "cannot be negative"}
	fmt.Printf("Error: %s\n", err)

	// Demonstrate validation
	validationErr := validateAge(-5)
	if validationErr != nil {
		fmt.Printf("Validation error: %s\n", validationErr)
	}
	fmt.Println()
}

// Stringer interface implementation
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return ValidationError{Field: "age", Message: "cannot be negative"}
	}
	if age > 150 {
		return ValidationError{Field: "age", Message: "cannot exceed 150"}
	}
	return nil
}

// Interface vs concrete types
func interfaceVsConcreteTypes() {
	fmt.Println("7. Interface vs Concrete Types")
	fmt.Println("------------------------------")

	// Demonstrate when to use interfaces
	fmt.Println("When to use interfaces:")

	// Strategy pattern example
	processor := &PaymentProcessor{}

	processor.SetStrategy(CreditCardPayment{})
	processor.ProcessPayment(100.0)

	processor.SetStrategy(PayPalPayment{})
	processor.ProcessPayment(50.0)

	processor.SetStrategy(BankTransferPayment{})
	processor.ProcessPayment(200.0)

	// Demonstrate when to use concrete types
	fmt.Println("\nWhen to use concrete types:")
	numbers := []int{1, 2, 3, 4, 5}
	sum := processNumbersDirect(numbers)
	fmt.Printf("Sum of numbers: %d\n", sum)
	fmt.Println()
}

// Strategy pattern interfaces
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

// Concrete type function (faster)
func processNumbersDirect(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Interface design patterns
func interfaceDesignPatterns() {
	fmt.Println("8. Interface Design Patterns")
	fmt.Println("----------------------------")

	// Factory pattern
	fmt.Println("Factory pattern:")
	animals := []string{"dog", "cat", "bird"}
	for _, animalType := range animals {
		animal, err := NewAnimal(animalType)
		if err != nil {
			fmt.Printf("Error creating %s: %v\n", animalType, err)
			continue
		}
		fmt.Printf("%s says: %s\n", animalType, animal.MakeSound())
	}

	// Observer pattern
	fmt.Println("\nObserver pattern:")
	newsAgency := &NewsAgency{}
	channel1 := NewsChannel{name: "CNN"}
	channel2 := NewsChannel{name: "BBC"}

	newsAgency.Attach(channel1)
	newsAgency.Attach(channel2)

	newsAgency.Notify("Breaking news: Go interfaces are awesome!")
	fmt.Println()
}

// Factory pattern
type AnimalFactory interface {
	MakeSound() string
}

type DogFactory struct{}
type CatFactory struct{}
type BirdFactory struct{}

func (d DogFactory) MakeSound() string  { return "Woof!" }
func (c CatFactory) MakeSound() string  { return "Meow!" }
func (b BirdFactory) MakeSound() string { return "Tweet!" }

func NewAnimal(animalType string) (AnimalFactory, error) {
	switch animalType {
	case "dog":
		return DogFactory{}, nil
	case "cat":
		return CatFactory{}, nil
	case "bird":
		return BirdFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown animal type: %s", animalType)
	}
}

// Observer pattern
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

// Interface testing
func interfaceTesting() {
	fmt.Println("9. Interface Testing")
	fmt.Println("-------------------")

	// Demonstrate testing with interfaces
	fmt.Println("Testing with interfaces:")

	// Mock store for testing
	mockStore := &MockStore{data: map[string]string{"1": "Alice"}}
	service := NewUserService(mockStore)

	name, err := service.GetUserName("1")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("User name: %s\n", name)
	}

	// Test with non-existent user
	name, err = service.GetUserName("2")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("User name: %s\n", name)
	}
	fmt.Println()
}

// DataStore interface for testing
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

// UserService that uses DataStore
type UserService struct {
	store DataStore
}

func NewUserService(store DataStore) *UserService {
	return &UserService{store: store}
}

func (us *UserService) GetUserName(id string) (string, error) {
	return us.store.Get(id)
}

// Mock store for testing
type MockStore struct {
	data map[string]string
}

func (m *MockStore) Get(id string) (string, error) {
	if value, exists := m.data[id]; exists {
		return value, nil
	}
	return "", fmt.Errorf("key not found: %s", id)
}

func (m *MockStore) Set(id, value string) error {
	m.data[id] = value
	return nil
}

// Interface performance
func interfacePerformance() {
	fmt.Println("10. Interface Performance")
	fmt.Println("-------------------------")

	// Demonstrate interface overhead
	fmt.Println("Interface vs concrete type performance:")

	// Direct call (faster)
	calc := SimpleCalculator{}
	result := calc.Add(5, 3)
	fmt.Printf("Direct call result: %d\n", result)

	// Interface call (slightly slower)
	var calcInterface Calculator = SimpleCalculator{}
	result = calcInterface.Add(5, 3)
	fmt.Printf("Interface call result: %d\n", result)

	// Performance comparison
	numbers := []int{1, 2, 3, 4, 5}
	sumDirect := processNumbersDirect(numbers)
	sumInterface := processNumbersInterface(SimpleCalculator{}, numbers)
	fmt.Printf("Direct sum: %d, Interface sum: %d\n", sumDirect, sumInterface)
	fmt.Println()
}

// Calculator interface
type Calculator interface {
	Add(a, b int) int
}

type SimpleCalculator struct{}

func (sc SimpleCalculator) Add(a, b int) int {
	return a + b
}

// Interface function (slower)
func processNumbersInterface(processor Calculator, numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum = processor.Add(sum, n)
	}
	return sum
}

// Standard library interfaces
func standardLibraryInterfaces() {
	fmt.Println("11. Standard Library Interfaces")
	fmt.Println("--------------------------------")

	// Demonstrate sort.Interface
	fmt.Println("Sort.Interface example:")
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	fmt.Printf("Before sorting: %+v\n", people)
	sort.Sort(ByAge(people))
	fmt.Printf("After sorting by age: %+v\n", people)

	// Demonstrate custom sorting
	sort.Sort(ByName(people))
	fmt.Printf("After sorting by name: %+v\n", people)
	fmt.Println()
}

// ByAge implements sort.Interface for []Person based on Age field
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByName implements sort.Interface for []Person based on Name field
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
