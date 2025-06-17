# Common Design Patterns in Go

## Overview

Design patterns are proven solutions to common software design problems. This guide covers the most commonly used design patterns in Go, including creational, structural, and behavioral patterns with practical examples.

## Creational Patterns

### Singleton Pattern

Ensures a class has only one instance and provides global access to it.

```go
package main

import (
    "fmt"
    "sync"
)

type Database struct {
    connection string
}

type Singleton struct {
    db *Database
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{
            db: &Database{connection: "database_connection"},
        }
    })
    return instance
}

func (s *Singleton) GetDatabase() *Database {
    return s.db
}

func main() {
    instance1 := GetInstance()
    instance2 := GetInstance()
    
    fmt.Printf("Instance 1: %p\n", instance1)
    fmt.Printf("Instance 2: %p\n", instance2)
    fmt.Printf("Same instance: %t\n", instance1 == instance2)
}
```

### Factory Pattern

Creates objects without specifying their exact class.

```go
package main

import "fmt"

// Product interface
type Product interface {
    GetName() string
    GetPrice() float64
}

// Concrete products
type Book struct {
    name  string
    price float64
}

func (b *Book) GetName() string {
    return b.name
}

func (b *Book) GetPrice() float64 {
    return b.price
}

type Electronics struct {
    name  string
    price float64
}

func (e *Electronics) GetName() string {
    return e.name
}

func (e *Electronics) GetPrice() float64 {
    return e.price
}

// Factory
type ProductFactory struct{}

func (pf *ProductFactory) CreateProduct(productType string, name string, price float64) Product {
    switch productType {
    case "book":
        return &Book{name: name, price: price}
    case "electronics":
        return &Electronics{name: name, price: price}
    default:
        return nil
    }
}

func main() {
    factory := &ProductFactory{}
    
    book := factory.CreateProduct("book", "Go Programming", 29.99)
    electronics := factory.CreateProduct("electronics", "Laptop", 999.99)
    
    fmt.Printf("Book: %s - $%.2f\n", book.GetName(), book.GetPrice())
    fmt.Printf("Electronics: %s - $%.2f\n", electronics.GetName(), electronics.GetPrice())
}
```

### Builder Pattern

Constructs complex objects step by step.

```go
package main

import "fmt"

type Computer struct {
    CPU       string
    RAM       string
    Storage   string
    Graphics  string
    Monitor   string
}

type ComputerBuilder struct {
    computer *Computer
}

func NewComputerBuilder() *ComputerBuilder {
    return &ComputerBuilder{
        computer: &Computer{},
    }
}

func (cb *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
    cb.computer.CPU = cpu
    return cb
}

func (cb *ComputerBuilder) SetRAM(ram string) *ComputerBuilder {
    cb.computer.RAM = ram
    return cb
}

func (cb *ComputerBuilder) SetStorage(storage string) *ComputerBuilder {
    cb.computer.Storage = storage
    return cb
}

func (cb *ComputerBuilder) SetGraphics(graphics string) *ComputerBuilder {
    cb.computer.Graphics = graphics
    return cb
}

func (cb *ComputerBuilder) SetMonitor(monitor string) *ComputerBuilder {
    cb.computer.Monitor = monitor
    return cb
}

func (cb *ComputerBuilder) Build() *Computer {
    return cb.computer
}

func main() {
    computer := NewComputerBuilder().
        SetCPU("Intel i7").
        SetRAM("16GB").
        SetStorage("512GB SSD").
        SetGraphics("NVIDIA RTX 3080").
        SetMonitor("27\" 4K").
        Build()
    
    fmt.Printf("Computer: %+v\n", computer)
}
```

### Abstract Factory Pattern

Creates families of related objects.

```go
package main

import "fmt"

// Abstract products
type Button interface {
    Render()
}

type Checkbox interface {
    Render()
}

// Concrete products for Windows
type WindowsButton struct{}

func (wb *WindowsButton) Render() {
    fmt.Println("Rendering Windows button")
}

type WindowsCheckbox struct{}

func (wc *WindowsCheckbox) Render() {
    fmt.Println("Rendering Windows checkbox")
}

// Concrete products for Mac
type MacButton struct{}

func (mb *MacButton) Render() {
    fmt.Println("Rendering Mac button")
}

type MacCheckbox struct{}

func (mc *MacCheckbox) Render() {
    fmt.Println("Rendering Mac checkbox")
}

// Abstract factory
type GUIFactory interface {
    CreateButton() Button
    CreateCheckbox() Checkbox
}

// Concrete factories
type WindowsFactory struct{}

func (wf *WindowsFactory) CreateButton() Button {
    return &WindowsButton{}
}

func (wf *WindowsFactory) CreateCheckbox() Checkbox {
    return &WindowsCheckbox{}
}

type MacFactory struct{}

func (mf *MacFactory) CreateButton() Button {
    return &MacButton{}
}

func (mf *MacFactory) CreateCheckbox() Checkbox {
    return &MacCheckbox{}
}

// Application
type Application struct {
    factory GUIFactory
}

func NewApplication(factory GUIFactory) *Application {
    return &Application{factory: factory}
}

func (app *Application) CreateUI() {
    button := app.factory.CreateButton()
    checkbox := app.factory.CreateCheckbox()
    
    button.Render()
    checkbox.Render()
}

func main() {
    // Windows application
    windowsApp := NewApplication(&WindowsFactory{})
    windowsApp.CreateUI()
    
    // Mac application
    macApp := NewApplication(&MacFactory{})
    macApp.CreateUI()
}
```

## Structural Patterns

### Adapter Pattern

Allows incompatible interfaces to work together.

```go
package main

import "fmt"

// Target interface
type PaymentProcessor interface {
    ProcessPayment(amount float64) error
}

// Adaptee (existing system)
type LegacyPaymentSystem struct{}

func (lps *LegacyPaymentSystem) Charge(amount float64, currency string) error {
    fmt.Printf("Legacy system charging %.2f %s\n", amount, currency)
    return nil
}

// Adapter
type LegacyPaymentAdapter struct {
    legacySystem *LegacyPaymentSystem
}

func NewLegacyPaymentAdapter() *LegacyPaymentAdapter {
    return &LegacyPaymentAdapter{
        legacySystem: &LegacyPaymentSystem{},
    }
}

func (lpa *LegacyPaymentAdapter) ProcessPayment(amount float64) error {
    return lpa.legacySystem.Charge(amount, "USD")
}

// Modern payment system
type ModernPaymentSystem struct{}

func (mps *ModernPaymentSystem) ProcessPayment(amount float64) error {
    fmt.Printf("Modern system processing payment of %.2f USD\n", amount)
    return nil
}

// Client code
func ProcessPayment(processor PaymentProcessor, amount float64) {
    if err := processor.ProcessPayment(amount); err != nil {
        fmt.Printf("Payment failed: %v\n", err)
    } else {
        fmt.Println("Payment processed successfully")
    }
}

func main() {
    // Use modern system
    modernProcessor := &ModernPaymentSystem{}
    ProcessPayment(modernProcessor, 100.0)
    
    // Use legacy system through adapter
    legacyAdapter := NewLegacyPaymentAdapter()
    ProcessPayment(legacyAdapter, 50.0)
}
```

### Bridge Pattern

Decouples abstraction from implementation.

```go
package main

import "fmt"

// Implementation interface
type DrawingAPI interface {
    DrawCircle(x, y, radius int)
    DrawRectangle(x, y, width, height int)
}

// Concrete implementations
type DrawingAPI1 struct{}

func (api *DrawingAPI1) DrawCircle(x, y, radius int) {
    fmt.Printf("API1: Drawing circle at (%d,%d) with radius %d\n", x, y, radius)
}

func (api *DrawingAPI1) DrawRectangle(x, y, width, height int) {
    fmt.Printf("API1: Drawing rectangle at (%d,%d) with width %d and height %d\n", x, y, width, height)
}

type DrawingAPI2 struct{}

func (api *DrawingAPI2) DrawCircle(x, y, radius int) {
    fmt.Printf("API2: Drawing circle at (%d,%d) with radius %d\n", x, y, radius)
}

func (api *DrawingAPI2) DrawRectangle(x, y, width, height int) {
    fmt.Printf("API2: Drawing rectangle at (%d,%d) with width %d and height %d\n", x, y, width, height)
}

// Abstraction
type Shape interface {
    Draw()
    ResizeByPercentage(percentage int)
}

// Refined abstraction
type CircleShape struct {
    x, y, radius int
    drawingAPI   DrawingAPI
}

func NewCircleShape(x, y, radius int, drawingAPI DrawingAPI) *CircleShape {
    return &CircleShape{
        x:          x,
        y:          y,
        radius:     radius,
        drawingAPI: drawingAPI,
    }
}

func (c *CircleShape) Draw() {
    c.drawingAPI.DrawCircle(c.x, c.y, c.radius)
}

func (c *CircleShape) ResizeByPercentage(percentage int) {
    c.radius = c.radius * percentage / 100
}

type RectangleShape struct {
    x, y, width, height int
    drawingAPI          DrawingAPI
}

func NewRectangleShape(x, y, width, height int, drawingAPI DrawingAPI) *RectangleShape {
    return &RectangleShape{
        x:          x,
        y:          y,
        width:      width,
        height:     height,
        drawingAPI: drawingAPI,
    }
}

func (r *RectangleShape) Draw() {
    r.drawingAPI.DrawRectangle(r.x, r.y, r.width, r.height)
}

func (r *RectangleShape) ResizeByPercentage(percentage int) {
    r.width = r.width * percentage / 100
    r.height = r.height * percentage / 100
}

func main() {
    api1 := &DrawingAPI1{}
    api2 := &DrawingAPI2{}
    
    circle1 := NewCircleShape(1, 2, 3, api1)
    circle2 := NewCircleShape(5, 7, 11, api2)
    
    rectangle1 := NewRectangleShape(1, 2, 3, 4, api1)
    rectangle2 := NewRectangleShape(5, 7, 11, 13, api2)
    
    circle1.Draw()
    circle2.Draw()
    rectangle1.Draw()
    rectangle2.Draw()
    
    circle1.ResizeByPercentage(50)
    circle1.Draw()
}
```

### Composite Pattern

Composes objects into tree structures.

```go
package main

import "fmt"

// Component interface
type FileSystemComponent interface {
    GetName() string
    GetSize() int
    Display(indent string)
}

// Leaf
type File struct {
    name string
    size int
}

func NewFile(name string, size int) *File {
    return &File{name: name, size: size}
}

func (f *File) GetName() string {
    return f.name
}

func (f *File) GetSize() int {
    return f.size
}

func (f *File) Display(indent string) {
    fmt.Printf("%sFile: %s, Size: %d bytes\n", indent, f.name, f.size)
}

// Composite
type Directory struct {
    name      string
    children  []FileSystemComponent
}

func NewDirectory(name string) *Directory {
    return &Directory{
        name:     name,
        children: make([]FileSystemComponent, 0),
    }
}

func (d *Directory) GetName() string {
    return d.name
}

func (d *Directory) GetSize() int {
    totalSize := 0
    for _, child := range d.children {
        totalSize += child.GetSize()
    }
    return totalSize
}

func (d *Directory) Display(indent string) {
    fmt.Printf("%sDirectory: %s, Total Size: %d bytes\n", indent, d.name, d.GetSize())
    for _, child := range d.children {
        child.Display(indent + "  ")
    }
}

func (d *Directory) Add(component FileSystemComponent) {
    d.children = append(d.children, component)
}

func (d *Directory) Remove(component FileSystemComponent) {
    for i, child := range d.children {
        if child == component {
            d.children = append(d.children[:i], d.children[i+1:]...)
            break
        }
    }
}

func main() {
    root := NewDirectory("root")
    
    documents := NewDirectory("documents")
    pictures := NewDirectory("pictures")
    
    file1 := NewFile("readme.txt", 1024)
    file2 := NewFile("report.pdf", 2048)
    file3 := NewFile("photo.jpg", 3072)
    
    documents.Add(file1)
    documents.Add(file2)
    pictures.Add(file3)
    
    root.Add(documents)
    root.Add(pictures)
    
    root.Display("")
}
```

### Decorator Pattern

Adds behavior to objects dynamically.

```go
package main

import "fmt"

// Component interface
type Coffee interface {
    GetCost() float64
    GetDescription() string
}

// Concrete component
type SimpleCoffee struct{}

func (sc *SimpleCoffee) GetCost() float64 {
    return 2.0
}

func (sc *SimpleCoffee) GetDescription() string {
    return "Simple coffee"
}

// Base decorator
type CoffeeDecorator struct {
    coffee Coffee
}

func (cd *CoffeeDecorator) GetCost() float64 {
    return cd.coffee.GetCost()
}

func (cd *CoffeeDecorator) GetDescription() string {
    return cd.coffee.GetDescription()
}

// Concrete decorators
type MilkDecorator struct {
    CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
    return &MilkDecorator{
        CoffeeDecorator: CoffeeDecorator{coffee: coffee},
    }
}

func (md *MilkDecorator) GetCost() float64 {
    return md.coffee.GetCost() + 0.5
}

func (md *MilkDecorator) GetDescription() string {
    return md.coffee.GetDescription() + ", milk"
}

type SugarDecorator struct {
    CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
    return &SugarDecorator{
        CoffeeDecorator: CoffeeDecorator{coffee: coffee},
    }
}

func (sd *SugarDecorator) GetCost() float64 {
    return sd.coffee.GetCost() + 0.2
}

func (sd *SugarDecorator) GetDescription() string {
    return sd.coffee.GetDescription() + ", sugar"
}

type WhipDecorator struct {
    CoffeeDecorator
}

func NewWhipDecorator(coffee Coffee) *WhipDecorator {
    return &WhipDecorator{
        CoffeeDecorator: CoffeeDecorator{coffee: coffee},
    }
}

func (wd *WhipDecorator) GetCost() float64 {
    return wd.coffee.GetCost() + 0.7
}

func (wd *WhipDecorator) GetDescription() string {
    return wd.coffee.GetDescription() + ", whip"
}

func main() {
    coffee := &SimpleCoffee{}
    fmt.Printf("Cost: $%.2f, Description: %s\n", coffee.GetCost(), coffee.GetDescription())
    
    coffeeWithMilk := NewMilkDecorator(coffee)
    fmt.Printf("Cost: $%.2f, Description: %s\n", coffeeWithMilk.GetCost(), coffeeWithMilk.GetDescription())
    
    coffeeWithMilkAndSugar := NewSugarDecorator(coffeeWithMilk)
    fmt.Printf("Cost: $%.2f, Description: %s\n", coffeeWithMilkAndSugar.GetCost(), coffeeWithMilkAndSugar.GetDescription())
    
    coffeeWithEverything := NewWhipDecorator(coffeeWithMilkAndSugar)
    fmt.Printf("Cost: $%.2f, Description: %s\n", coffeeWithEverything.GetCost(), coffeeWithEverything.GetDescription())
}
```

## Behavioral Patterns

### Observer Pattern

Defines a one-to-many dependency between objects.

```go
package main

import "fmt"

// Observer interface
type Observer interface {
    Update(temperature, humidity, pressure float64)
}

// Subject interface
type Subject interface {
    RegisterObserver(observer Observer)
    RemoveObserver(observer Observer)
    NotifyObservers()
}

// Concrete subject
type WeatherData struct {
    observers    []Observer
    temperature  float64
    humidity     float64
    pressure     float64
}

func NewWeatherData() *WeatherData {
    return &WeatherData{
        observers: make([]Observer, 0),
    }
}

func (wd *WeatherData) RegisterObserver(observer Observer) {
    wd.observers = append(wd.observers, observer)
}

func (wd *WeatherData) RemoveObserver(observer Observer) {
    for i, obs := range wd.observers {
        if obs == observer {
            wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
            break
        }
    }
}

func (wd *WeatherData) NotifyObservers() {
    for _, observer := range wd.observers {
        observer.Update(wd.temperature, wd.humidity, wd.pressure)
    }
}

func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
    wd.temperature = temperature
    wd.humidity = humidity
    wd.pressure = pressure
    wd.NotifyObservers()
}

// Concrete observers
type CurrentConditionsDisplay struct {
    temperature float64
    humidity    float64
}

func (ccd *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
    ccd.temperature = temperature
    ccd.humidity = humidity
    ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
    fmt.Printf("Current conditions: %.1f°F and %.1f%% humidity\n", ccd.temperature, ccd.humidity)
}

type StatisticsDisplay struct {
    temperatures []float64
    humidities   []float64
}

func (sd *StatisticsDisplay) Update(temperature, humidity, pressure float64) {
    sd.temperatures = append(sd.temperatures, temperature)
    sd.humidities = append(sd.humidities, humidity)
    sd.Display()
}

func (sd *StatisticsDisplay) Display() {
    if len(sd.temperatures) == 0 {
        return
    }
    
    tempSum := 0.0
    humSum := 0.0
    
    for _, temp := range sd.temperatures {
        tempSum += temp
    }
    for _, hum := range sd.humidities {
        humSum += hum
    }
    
    avgTemp := tempSum / float64(len(sd.temperatures))
    avgHum := humSum / float64(len(sd.humidities))
    
    fmt.Printf("Avg/Max/Min temperature = %.1f/%.1f/%.1f\n", avgTemp, sd.temperatures[len(sd.temperatures)-1], sd.temperatures[0])
    fmt.Printf("Avg humidity = %.1f%%\n", avgHum)
}

func main() {
    weatherData := NewWeatherData()
    
    currentDisplay := &CurrentConditionsDisplay{}
    statisticsDisplay := &StatisticsDisplay{}
    
    weatherData.RegisterObserver(currentDisplay)
    weatherData.RegisterObserver(statisticsDisplay)
    
    weatherData.SetMeasurements(80, 65, 30.4)
    weatherData.SetMeasurements(82, 70, 29.2)
    weatherData.SetMeasurements(78, 90, 29.2)
}
```

### Strategy Pattern

Defines a family of algorithms and makes them interchangeable.

```go
package main

import "fmt"

// Strategy interface
type PaymentStrategy interface {
    Pay(amount float64) error
}

// Concrete strategies
type CreditCardPayment struct {
    cardNumber string
    cvv        string
    expiry     string
}

func NewCreditCardPayment(cardNumber, cvv, expiry string) *CreditCardPayment {
    return &CreditCardPayment{
        cardNumber: cardNumber,
        cvv:        cvv,
        expiry:     expiry,
    }
}

func (ccp *CreditCardPayment) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using credit card ending in %s\n", amount, ccp.cardNumber[len(ccp.cardNumber)-4:])
    return nil
}

type PayPalPayment struct {
    email string
}

func NewPayPalPayment(email string) *PayPalPayment {
    return &PayPalPayment{email: email}
}

func (ppp *PayPalPayment) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using PayPal account %s\n", amount, ppp.email)
    return nil
}

type BitcoinPayment struct {
    walletAddress string
}

func NewBitcoinPayment(walletAddress string) *BitcoinPayment {
    return &BitcoinPayment{walletAddress: walletAddress}
}

func (bp *BitcoinPayment) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using Bitcoin wallet %s\n", amount, bp.walletAddress)
    return nil
}

// Context
type ShoppingCart struct {
    paymentStrategy PaymentStrategy
}

func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{}
}

func (sc *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
    sc.paymentStrategy = strategy
}

func (sc *ShoppingCart) Checkout(amount float64) error {
    if sc.paymentStrategy == nil {
        return fmt.Errorf("no payment strategy set")
    }
    return sc.paymentStrategy.Pay(amount)
}

func main() {
    cart := NewShoppingCart()
    
    // Pay with credit card
    cart.SetPaymentStrategy(NewCreditCardPayment("1234567890123456", "123", "12/25"))
    cart.Checkout(100.0)
    
    // Pay with PayPal
    cart.SetPaymentStrategy(NewPayPalPayment("user@example.com"))
    cart.Checkout(50.0)
    
    // Pay with Bitcoin
    cart.SetPaymentStrategy(NewBitcoinPayment("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"))
    cart.Checkout(25.0)
}
```

### Command Pattern

Encapsulates a request as an object.

```go
package main

import "fmt"

// Command interface
type Command interface {
    Execute()
    Undo()
}

// Receiver
type Light struct {
    location string
}

func NewLight(location string) *Light {
    return &Light{location: location}
}

func (l *Light) On() {
    fmt.Printf("%s light is on\n", l.location)
}

func (l *Light) Off() {
    fmt.Printf("%s light is off\n", l.location)
}

// Concrete commands
type LightOnCommand struct {
    light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
    return &LightOnCommand{light: light}
}

func (loc *LightOnCommand) Execute() {
    loc.light.On()
}

func (loc *LightOnCommand) Undo() {
    loc.light.Off()
}

type LightOffCommand struct {
    light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
    return &LightOffCommand{light: light}
}

func (loc *LightOffCommand) Execute() {
    loc.light.Off()
}

func (loc *LightOffCommand) Undo() {
    loc.light.On()
}

// Invoker
type RemoteControl struct {
    onCommands  []Command
    offCommands []Command
    undoCommand Command
}

func NewRemoteControl() *RemoteControl {
    return &RemoteControl{
        onCommands:  make([]Command, 7),
        offCommands: make([]Command, 7),
    }
}

func (rc *RemoteControl) SetCommand(slot int, onCommand, offCommand Command) {
    rc.onCommands[slot] = onCommand
    rc.offCommands[slot] = offCommand
}

func (rc *RemoteControl) OnButtonWasPushed(slot int) {
    if rc.onCommands[slot] != nil {
        rc.onCommands[slot].Execute()
        rc.undoCommand = rc.onCommands[slot]
    }
}

func (rc *RemoteControl) OffButtonWasPushed(slot int) {
    if rc.offCommands[slot] != nil {
        rc.offCommands[slot].Execute()
        rc.undoCommand = rc.offCommands[slot]
    }
}

func (rc *RemoteControl) UndoButtonWasPushed() {
    if rc.undoCommand != nil {
        rc.undoCommand.Undo()
    }
}

func main() {
    remote := NewRemoteControl()
    
    livingRoomLight := NewLight("Living Room")
    kitchenLight := NewLight("Kitchen")
    
    livingRoomLightOn := NewLightOnCommand(livingRoomLight)
    livingRoomLightOff := NewLightOffCommand(livingRoomLight)
    kitchenLightOn := NewLightOnCommand(kitchenLight)
    kitchenLightOff := NewLightOffCommand(kitchenLight)
    
    remote.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
    remote.SetCommand(1, kitchenLightOn, kitchenLightOff)
    
    remote.OnButtonWasPushed(0)
    remote.OffButtonWasPushed(0)
    remote.OnButtonWasPushed(1)
    remote.OffButtonWasPushed(1)
    remote.UndoButtonWasPushed()
}
```

### Template Method Pattern

Defines the skeleton of an algorithm in a method.

```go
package main

import "fmt"

// Abstract class
type BeverageTemplate interface {
    PrepareRecipe()
    BoilWater()
    PourInCup()
    Hook() // Optional hook
}

// Base template
type BaseBeverage struct{}

func (bb *BaseBeverage) PrepareRecipe() {
    bb.BoilWater()
    bb.Brew()
    bb.PourInCup()
    if bb.CustomerWantsCondiments() {
        bb.AddCondiments()
    }
    bb.Hook()
}

func (bb *BaseBeverage) BoilWater() {
    fmt.Println("Boiling water")
}

func (bb *BaseBeverage) PourInCup() {
    fmt.Println("Pouring into cup")
}

func (bb *BaseBeverage) CustomerWantsCondiments() bool {
    return true // Default implementation
}

func (bb *BaseBeverage) Hook() {
    // Default empty implementation
}

// Abstract methods that must be implemented
func (bb *BaseBeverage) Brew() {
    // This should be overridden
}

func (bb *BaseBeverage) AddCondiments() {
    // This should be overridden
}

// Concrete implementations
type Coffee struct {
    BaseBeverage
}

func (c *Coffee) Brew() {
    fmt.Println("Dripping coffee through filter")
}

func (c *Coffee) AddCondiments() {
    fmt.Println("Adding sugar and milk")
}

type Tea struct {
    BaseBeverage
}

func (t *Tea) Brew() {
    fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
    fmt.Println("Adding lemon")
}

func (t *Tea) CustomerWantsCondiments() bool {
    return false // Override to not add condiments
}

func main() {
    coffee := &Coffee{}
    tea := &Tea{}
    
    fmt.Println("Making coffee:")
    coffee.PrepareRecipe()
    
    fmt.Println("\nMaking tea:")
    tea.PrepareRecipe()
}
```

## Summary

Design patterns provide:

- **Creational patterns**: Object creation mechanisms
- **Structural patterns**: Object composition and relationships
- **Behavioral patterns**: Communication between objects

Key patterns covered:
1. **Singleton**: Single instance with global access
2. **Factory**: Object creation without specifying exact class
3. **Builder**: Step-by-step object construction
4. **Abstract Factory**: Families of related objects
5. **Adapter**: Incompatible interface compatibility
6. **Bridge**: Abstraction-implementation decoupling
7. **Composite**: Tree structure composition
8. **Decorator**: Dynamic behavior addition
9. **Observer**: One-to-many dependency
10. **Strategy**: Interchangeable algorithms
11. **Command**: Request encapsulation
12. **Template Method**: Algorithm skeleton

Key points to remember:
1. Use patterns to solve specific problems, not for their own sake
2. Go's simplicity often makes some patterns unnecessary
3. Focus on composition over inheritance
4. Use interfaces for flexibility
5. Keep implementations simple and readable
6. Consider the trade-offs of each pattern
7. Adapt patterns to Go's idioms
8. Use patterns to improve code organization
9. Document when and why patterns are used
10. Refactor to patterns when needed

Understanding design patterns enables you to write more maintainable, flexible, and reusable code.

## Dependency Injection Patterns in Go

#### Overview

Dependency Injection (DI) is a technique for achieving Inversion of Control (IoC) between components and their dependencies. In Go, DI is typically achieved through interfaces, constructor functions, and sometimes with the help of DI containers or code generation tools. Go's simplicity and preference for explicitness make manual DI patterns idiomatic and easy to maintain.

#### 1. Constructor Injection

The most common and idiomatic pattern in Go. Dependencies are passed as parameters to constructors.

```go
type UserRepository interface {
    GetUser(id int) (*User, error)
}

type EmailService interface {
    SendEmail(to, subject, body string) error
}

type UserService struct {
    repo   UserRepository
    email  EmailService
}

func NewUserService(repo UserRepository, email EmailService) *UserService {
    return &UserService{repo: repo, email: email}
}
```

**Usage:**
```go
repo := NewPostgresUserRepository()
email := NewSMTPEmailService()
service := NewUserService(repo, email)
```

#### 2. Interface Injection

Define dependencies as interfaces and pass implementations at runtime, often as fields or setter methods.

```go
type Logger interface {
    Log(msg string)
}

type Service struct {
    logger Logger
}

func (s *Service) SetLogger(logger Logger) {
    s.logger = logger
}
```

#### 3. Functional Options

Use functional options to inject dependencies and configuration.

```go
type Server struct {
    port   int
    logger Logger
}

type Option func(*Server)

func WithPort(port int) Option {
    return func(s *Server) { s.port = port }
}

func WithLogger(logger Logger) Option {
    return func(s *Server) { s.logger = logger }
}

func NewServer(opts ...Option) *Server {
    s := &Server{port: 8080}
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

**Usage:**
```go
srv := NewServer(WithPort(9090), WithLogger(myLogger))
```

#### 4. Dependency Injection Containers

For large projects, you may use a DI container (e.g., Uber's dig, Facebook's inject, or Google's wire for compile-time DI).

**Example with Uber's dig:**
```go
import "go.uber.org/dig"

func BuildContainer() *dig.Container {
    c := dig.New()
    c.Provide(NewPostgresUserRepository)
    c.Provide(NewSMTPEmailService)
    c.Provide(NewUserService)
    return c
}
```

**Example with Google Wire (compile-time DI):**
```go
// +build wireinject

import "github.com/google/wire"

func InitializeUserService() *UserService {
    wire.Build(NewPostgresUserRepository, NewSMTPEmailService, NewUserService)
    return nil
}
```

#### 5. Manual Wiring

For most Go projects, manual wiring is preferred for its clarity and simplicity. Compose dependencies explicitly in your main function.

#### Best Practices

- Prefer constructor injection for clarity and testability.
- Use interfaces to decouple implementations.
- Avoid global state and singletons unless necessary.
- Use DI containers only for large, complex projects.
- Document dependencies clearly.

#### Summary Table:

| Pattern                | Pros                        | Cons                        | Use Case                |
|------------------------|-----------------------------|-----------------------------|-------------------------|
| Constructor Injection  | Simple, testable, explicit  | Boilerplate in large apps   | Most Go projects        |
| Interface Injection    | Flexible, decoupled         | Can be less explicit        | Pluggable components    |
| Functional Options     | Configurable, extensible    | Can be overused             | Config-heavy components |
| DI Container           | Automated, scalable         | More complex, less explicit | Large codebases         |
| Manual Wiring          | Explicit, clear             | Tedious for many deps       | Small/medium projects   |

#### In summary:  
Dependency injection in Go is best achieved through explicit constructor injection and interfaces. For larger projects, consider DI containers or code generation tools, but always favor clarity and simplicity. 