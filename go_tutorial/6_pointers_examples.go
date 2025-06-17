package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go Pointers Examples ===\n")

	// Basic pointer operations
	basicPointerOperations()

	// Pointer declaration and initialization
	pointerDeclaration()

	// Pointer operations
	pointerOperations()

	// Pointers and functions
	pointersAndFunctions()

	// Pointers to different types
	pointersToDifferentTypes()

	// Nil pointers
	nilPointers()

	// Common pointer patterns
	commonPointerPatterns()

	// Pointers and slices
	pointersAndSlices()

	// Pointers and maps
	pointersAndMaps()

	// Best practices
	bestPractices()

	// Performance considerations
	performanceConsiderations()
}

func basicPointerOperations() {
	fmt.Println("1. Basic Pointer Operations")
	fmt.Println("---------------------------")

	var x int = 42
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Memory address of x: %p\n", &x)

	var ptr *int = &x
	fmt.Printf("Value of ptr: %p\n", ptr)
	fmt.Printf("Value pointed to by ptr: %d\n", *ptr)

	// Modify value through pointer
	*ptr = 100
	fmt.Printf("New value of x: %d\n", x)
	fmt.Println()
}

func pointerDeclaration() {
	fmt.Println("2. Pointer Declaration and Initialization")
	fmt.Println("------------------------------------------")

	// Declare a pointer to int
	var ptr1 *int
	fmt.Printf("Zero value of pointer: %v\n", ptr1)

	// Declare and initialize a pointer
	var x int = 42
	var ptr2 *int = &x
	fmt.Printf("ptr2 points to: %d\n", *ptr2)

	// Short declaration
	ptr3 := &x
	fmt.Printf("ptr3 points to: %d\n", *ptr3)

	// Pointers to different types
	var i int = 42
	var f float64 = 3.14
	var s string = "hello"
	var b bool = true

	var ptrInt *int = &i
	var ptrFloat *float64 = &f
	var ptrString *string = &s
	var ptrBool *bool = &b

	fmt.Printf("int: %d\n", *ptrInt)
	fmt.Printf("float: %.2f\n", *ptrFloat)
	fmt.Printf("string: %s\n", *ptrString)
	fmt.Printf("bool: %t\n", *ptrBool)
	fmt.Println()
}

func pointerOperations() {
	fmt.Println("3. Pointer Operations")
	fmt.Println("---------------------")

	// Modifying values through pointers
	var x int = 42
	fmt.Printf("Before: %d\n", x)
	modifyValue(&x)
	fmt.Printf("After: %d\n", x)

	// Comparing pointers
	var y int = 42
	ptr1 := &x
	ptr2 := &y
	ptr3 := &x

	fmt.Printf("ptr1 == ptr2: %t\n", ptr1 == ptr2) // false (different addresses)
	fmt.Printf("ptr1 == ptr3: %t\n", ptr1 == ptr3) // true (same address)
	fmt.Printf("ptr1 == nil: %t\n", ptr1 == nil)   // false
	fmt.Println()
}

func modifyValue(ptr *int) {
	*ptr = 100
	fmt.Printf("Inside function: %d\n", *ptr)
}

func pointersAndFunctions() {
	fmt.Println("4. Pointers and Functions")
	fmt.Println("-------------------------")

	var x int = 42

	// Pass by value
	fmt.Printf("Before modifyByValue: %d\n", x)
	modifyByValue(x)
	fmt.Printf("After modifyByValue: %d\n", x)

	// Pass by reference
	fmt.Printf("Before modifyByReference: %d\n", x)
	modifyByReference(&x)
	fmt.Printf("After modifyByReference: %d\n", x)

	// Returning pointers
	ptr := createPointer()
	fmt.Printf("Returned pointer value: %d\n", *ptr)

	// Function parameters with pointers
	counter := &Counter{count: 0}
	counter.Increment()
	counter.Increment()
	fmt.Printf("Counter value: %d\n", counter.GetCount())
	fmt.Println()
}

func modifyByValue(x int) {
	x = 100
	fmt.Printf("Inside modifyByValue: %d\n", x)
}

func modifyByReference(x *int) {
	*x = 100
	fmt.Printf("Inside modifyByReference: %d\n", *x)
}

func createPointer() *int {
	x := 42
	return &x
}

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) GetCount() int {
	return c.count
}

func pointersToDifferentTypes() {
	fmt.Println("5. Pointers to Different Types")
	fmt.Println("-------------------------------")

	// Pointers to arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var ptr *[5]int = &arr

	fmt.Printf("Array: %v\n", *ptr)
	fmt.Printf("First element: %d\n", (*ptr)[0])

	// Modify through pointer
	(*ptr)[0] = 100
	fmt.Printf("Modified array: %v\n", arr)

	// Pointers to structs
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 30}
	var personPtr *Person = &person

	fmt.Printf("Person: %+v\n", *personPtr)
	fmt.Printf("Name: %s\n", (*personPtr).Name) // or personPtr.Name
	fmt.Println()
}

func nilPointers() {
	fmt.Println("6. Nil Pointers")
	fmt.Println("----------------")

	var ptr *int = nil
	fmt.Printf("ptr is nil: %t\n", ptr == nil)

	// Safe dereferencing
	safeDereference(ptr)

	var x int = 42
	var ptr2 *int = &x
	safeDereference(ptr2)
	fmt.Println()
}

func safeDereference(ptr *int) {
	if ptr != nil {
		fmt.Printf("Value: %d\n", *ptr)
	} else {
		fmt.Println("Pointer is nil")
	}
}

func commonPointerPatterns() {
	fmt.Println("7. Common Pointer Patterns")
	fmt.Println("---------------------------")

	// Optional parameters
	processData("test", nil)

	customTimeout := 60 * time.Second
	processData("test", &customTimeout)

	// Returning multiple values with pointers
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", *result)
	}

	_, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	}

	// Efficient data structures
	head := createLinkedList()
	printList(head)
	fmt.Println()
}

func processData(data string, timeout *time.Duration) {
	defaultTimeout := 30 * time.Second
	if timeout == nil {
		timeout = &defaultTimeout
	}
	fmt.Printf("Processing '%s' with timeout: %v\n", data, *timeout)
}

func divide(a, b int) (result *float64, err error) {
	if b == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	res := float64(a) / float64(b)
	return &res, nil
}

type Node struct {
	Value int
	Next  *Node
}

func createLinkedList() *Node {
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	return head
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func pointersAndSlices() {
	fmt.Println("8. Pointers and Slices")
	fmt.Println("-----------------------")

	// Understanding slice pointers
	var slice []int = []int{1, 2, 3, 4, 5}
	var ptr *[]int = &slice

	fmt.Printf("Slice: %v\n", *ptr)

	// Modify slice through pointer
	(*ptr)[0] = 100
	fmt.Printf("Modified slice: %v\n", slice)

	// When to use pointers with slices
	slice2 := []int{1, 2, 3}

	modifySlice(slice2)
	fmt.Printf("After modifySlice: %v\n", slice2)

	appendToSlice(&slice2, 4)
	fmt.Printf("After appendToSlice: %v\n", slice2)
	fmt.Println()
}

func modifySlice(slice []int) {
	slice[0] = 100 // This modifies the original slice
}

func appendToSlice(slicePtr *[]int, value int) {
	*slicePtr = append(*slicePtr, value)
}

func pointersAndMaps() {
	fmt.Println("9. Pointers and Maps")
	fmt.Println("---------------------")

	// Understanding map pointers
	var m map[string]int = map[string]int{"a": 1, "b": 2}
	var ptr *map[string]int = &m

	fmt.Printf("Map: %v\n", *ptr)

	// Modify map through pointer
	(*ptr)["c"] = 3
	fmt.Printf("Modified map: %v\n", m)

	// When to use pointers with maps
	m2 := map[string]int{"a": 1, "b": 2}

	modifyMap(m2)
	fmt.Printf("After modifyMap: %v\n", m2)

	replaceMap(&m2)
	fmt.Printf("After replaceMap: %v\n", m2)
	fmt.Println()
}

func modifyMap(m map[string]int) {
	m["new"] = 42 // This modifies the original map
}

func replaceMap(mapPtr *map[string]int) {
	*mapPtr = map[string]int{"replaced": 1}
}

func bestPractices() {
	fmt.Println("10. Best Practices")
	fmt.Println("------------------")

	// Use pointers sparingly
	counter := 0
	incrementCounter(&counter)
	fmt.Printf("Counter after increment: %d\n", counter)

	// Check for nil pointers
	err := processPointer(nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Use pointers for large structs
	largeStruct := &LargeStruct{Data: [1000]int{1, 2, 3}}
	processLargeStruct(largeStruct)

	// Pointer receivers for methods
	rect := &Rectangle{Width: 10, Height: 5}
	rect.SetWidth(15)
	fmt.Printf("Rectangle area: %.2f\n", rect.Area())
	fmt.Println()
}

func incrementCounter(counter *int) {
	*counter++
}

func processPointer(ptr *int) error {
	if ptr == nil {
		return fmt.Errorf("pointer is nil")
	}
	*ptr = 42
	return nil
}

type LargeStruct struct {
	Data [1000]int
}

func processLargeStruct(data *LargeStruct) {
	fmt.Printf("Processing large struct with %d elements\n", len(data.Data))
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) SetWidth(width float64) {
	r.Width = width
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func performanceConsiderations() {
	fmt.Println("11. Performance Considerations")
	fmt.Println("-------------------------------")

	// Small values: pass by value
	result := processSmallValue(42)
	fmt.Printf("Small value result: %d\n", result)

	// Large values: pass by pointer
	largeStruct := &LargeStruct{Data: [1000]int{1, 2, 3}}
	processLargeStruct(largeStruct)

	fmt.Println("Performance considerations completed.")
	fmt.Println()
}

func processSmallValue(x int) int {
	return x * 2
}
