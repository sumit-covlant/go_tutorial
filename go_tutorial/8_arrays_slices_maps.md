# Go Arrays, Slices, and Maps

## Overview

Go provides three main collection types: arrays, slices, and maps. Each serves different purposes and has unique characteristics that make them suitable for specific use cases.

## Arrays

### What are Arrays?

Arrays in Go are fixed-length sequences of elements of the same type. Once created, their length cannot be changed.

```go
var arr [5]int  // Declares an array of 5 integers
```

### Array Declaration and Initialization

```go
// Method 1: Zero-value initialization
var arr1 [5]int
fmt.Printf("arr1: %v\n", arr1) // [0 0 0 0 0]

// Method 2: Array literal
var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
fmt.Printf("arr2: %v\n", arr2) // [1 2 3 4 5]

// Method 3: Short declaration
arr3 := [5]int{1, 2, 3, 4, 5}

// Method 4: Let compiler count elements
arr4 := [...]int{1, 2, 3, 4, 5}
fmt.Printf("arr4 length: %d\n", len(arr4)) // 5

// Method 5: Initialize specific elements
arr5 := [5]int{1: 10, 3: 30}
fmt.Printf("arr5: %v\n", arr5) // [0 10 0 30 0]
```

### Accessing Array Elements

```go
arr := [5]int{1, 2, 3, 4, 5}

// Access by index (zero-based)
fmt.Printf("First element: %d\n", arr[0])  // 1
fmt.Printf("Last element: %d\n", arr[4])   // 5

// Modify elements
arr[0] = 100
fmt.Printf("Modified array: %v\n", arr) // [100 2 3 4 5]

// Get array length
fmt.Printf("Array length: %d\n", len(arr)) // 5
```

### Array Properties

```go
// Arrays are value types
arr1 := [3]int{1, 2, 3}
arr2 := arr1  // Creates a copy

arr2[0] = 100
fmt.Printf("arr1: %v\n", arr1) // [1 2 3] (unchanged)
fmt.Printf("arr2: %v\n", arr2) // [100 2 3]

// Arrays are comparable
arr3 := [3]int{1, 2, 3}
arr4 := [3]int{1, 2, 3}
fmt.Printf("arr3 == arr4: %t\n", arr3 == arr4) // true

// Different sizes are different types
var arr5 [3]int
var arr6 [4]int
// arr5 = arr6  // Compilation error: cannot assign [4]int to [3]int
```

### Multi-dimensional Arrays

```go
// 2D array
var matrix [3][3]int = [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Access elements
fmt.Printf("Element at [1][2]: %d\n", matrix[1][2]) // 6

// Modify elements
matrix[0][0] = 100
fmt.Printf("Modified matrix: %v\n", matrix)

// 3D array
var cube [2][3][4]int
fmt.Printf("Cube dimensions: %dx%dx%d\n", len(cube), len(cube[0]), len(cube[0][0]))
```

### Iterating Over Arrays

```go
arr := [5]int{1, 2, 3, 4, 5}

// Method 1: Traditional for loop
for i := 0; i < len(arr); i++ {
    fmt.Printf("arr[%d] = %d\n", i, arr[i])
}

// Method 2: Range loop
for index, value := range arr {
    fmt.Printf("arr[%d] = %d\n", index, value)
}

// Method 3: Range loop (index only)
for index := range arr {
    fmt.Printf("Index: %d\n", index)
}

// Method 4: Range loop (value only)
for _, value := range arr {
    fmt.Printf("Value: %d\n", value)
}
```

## Slices

### What are Slices?

Slices are dynamic, flexible views into arrays. They provide a more convenient way to work with sequences of data compared to arrays.

```go
var slice []int  // Declares a slice of integers
```

### Slice Declaration and Initialization

```go
// Method 1: Zero-value slice (nil)
var slice1 []int
fmt.Printf("slice1: %v, len: %d, cap: %d, nil: %t\n", 
    slice1, len(slice1), cap(slice1), slice1 == nil) // [] 0 0 true

// Method 2: Slice literal
slice2 := []int{1, 2, 3, 4, 5}
fmt.Printf("slice2: %v, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))

// Method 3: Make function
slice3 := make([]int, 5)        // Length 5, capacity 5
slice4 := make([]int, 3, 5)     // Length 3, capacity 5

// Method 4: From array
arr := [5]int{1, 2, 3, 4, 5}
slice5 := arr[1:4]              // [2 3 4]
slice6 := arr[:3]               // [1 2 3]
slice7 := arr[2:]               // [3 4 5]
slice8 := arr[:]                // [1 2 3 4 5]
```

### Slice Internals

Slices have three components:
- **Pointer**: Points to the underlying array
- **Length**: Number of elements in the slice
- **Capacity**: Maximum number of elements from the start of the slice

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]

fmt.Printf("Array: %v\n", arr)                    // [1 2 3 4 5]
fmt.Printf("Slice: %v\n", slice)                  // [2 3 4]
fmt.Printf("Slice length: %d\n", len(slice))      // 3
fmt.Printf("Slice capacity: %d\n", cap(slice))    // 4

// Visual representation:
// Array:  [1][2][3][4][5]
// Slice:     [2][3][4]
//           ^     ^
//         pointer len=3, cap=4
```

### Modifying Slices

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]

// Modify slice elements
slice[0] = 100
fmt.Printf("Modified slice: %v\n", slice) // [100 3 4]
fmt.Printf("Original array: %v\n", arr)    // [1 100 3 4 5]

// Extend slice (within capacity)
slice = slice[:4]  // Extend to capacity
fmt.Printf("Extended slice: %v\n", slice) // [100 3 4 5]

// Cannot extend beyond capacity
// slice = slice[:5]  // Runtime panic: slice bounds out of range
```

### Appending to Slices

```go
// Start with empty slice
slice := []int{}

// Append single element
slice = append(slice, 1)
fmt.Printf("After append 1: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

// Append multiple elements
slice = append(slice, 2, 3, 4)
fmt.Printf("After append 2,3,4: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

// Append slice to slice
slice2 := []int{5, 6}
slice = append(slice, slice2...)
fmt.Printf("After append slice2: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

// Capacity growth pattern
var s []int
for i := 0; i < 10; i++ {
    s = append(s, i)
    fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
}
```

### Slice Operations

```go
slice := []int{1, 2, 3, 4, 5}

// Access elements
fmt.Printf("First element: %d\n", slice[0])
fmt.Printf("Last element: %d\n", slice[len(slice)-1])

// Slice operations
fmt.Printf("slice[1:3]: %v\n", slice[1:3])   // [2 3]
fmt.Printf("slice[:3]: %v\n", slice[:3])     // [1 2 3]
fmt.Printf("slice[2:]: %v\n", slice[2:])     // [3 4 5]
fmt.Printf("slice[:]: %v\n", slice[:])       // [1 2 3 4 5]

// Copy slices
slice2 := make([]int, len(slice))
copy(slice2, slice)
fmt.Printf("Copied slice: %v\n", slice2)
```

### Common Slice Patterns

#### 1. Removing Elements

```go
// Remove element at index
func removeElement(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

// Remove last element
func removeLast(slice []int) []int {
    return slice[:len(slice)-1]
}

// Remove first element
func removeFirst(slice []int) []int {
    return slice[1:]
}
```

#### 2. Filtering

```go
// Filter even numbers
func filterEven(numbers []int) []int {
    var result []int
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

// Filter with function
func filter(numbers []int, predicate func(int) bool) []int {
    var result []int
    for _, num := range numbers {
        if predicate(num) {
            result = append(result, num)
        }
    }
    return result
}
```

#### 3. Mapping

```go
// Double each element
func double(numbers []int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = num * 2
    }
    return result
}

// Map with function
func mapSlice(numbers []int, fn func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = fn(num)
    }
    return result
}
```

## Maps

### What are Maps?

Maps are unordered collections of key-value pairs. They provide fast lookup, insertion, and deletion operations.

```go
var m map[string]int  // Declares a map with string keys and int values
```

### Map Declaration and Initialization

```go
// Method 1: Zero-value map (nil)
var m1 map[string]int
fmt.Printf("m1: %v, nil: %t\n", m1, m1 == nil) // map[] true

// Method 2: Map literal
m2 := map[string]int{
    "apple":  1,
    "banana": 2,
    "cherry": 3,
}

// Method 3: Make function
m3 := make(map[string]int)
m4 := make(map[string]int, 10)  // Initial capacity hint

// Method 4: Short declaration
m5 := map[string]int{"a": 1, "b": 2}
```

### Map Operations

```go
m := make(map[string]int)

// Insert or update
m["apple"] = 1
m["banana"] = 2
m["apple"] = 3  // Updates existing key

// Access values
fmt.Printf("apple: %d\n", m["apple"])   // 3
fmt.Printf("banana: %d\n", m["banana"]) // 2

// Check if key exists
value, exists := m["cherry"]
if exists {
    fmt.Printf("cherry: %d\n", value)
} else {
    fmt.Println("cherry not found")
}

// Delete key
delete(m, "banana")
fmt.Printf("After delete: %v\n", m) // map[apple:3]

// Get length
fmt.Printf("Map length: %d\n", len(m)) // 1
```

### Map Properties

```go
// Maps are reference types
m1 := map[string]int{"a": 1, "b": 2}
m2 := m1  // Creates a reference, not a copy

m2["c"] = 3
fmt.Printf("m1: %v\n", m1) // map[a:1 b:2 c:3] (changed)
fmt.Printf("m2: %v\n", m2) // map[a:1 b:2 c:3]

// Maps are not comparable
// m1 == m2  // Compilation error: maps are not comparable
```

### Iterating Over Maps

```go
m := map[string]int{
    "apple":  1,
    "banana": 2,
    "cherry": 3,
}

// Iterate over key-value pairs
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Iterate over keys only
for key := range m {
    fmt.Printf("Key: %s\n", key)
}

// Note: Map iteration order is random
```

### Common Map Patterns

#### 1. Counting Occurrences

```go
func countWords(text string) map[string]int {
    words := strings.Fields(text)
    counts := make(map[string]int)
    
    for _, word := range words {
        counts[word]++
    }
    
    return counts
}

func main() {
    text := "hello world hello go world"
    counts := countWords(text)
    fmt.Printf("Word counts: %v\n", counts) // map[go:1 hello:2 world:2]
}
```

#### 2. Grouping Data

```go
type Person struct {
    Name string
    Age  int
    City string
}

func groupByCity(people []Person) map[string][]Person {
    groups := make(map[string][]Person)
    
    for _, person := range people {
        groups[person.City] = append(groups[person.City], person)
    }
    
    return groups
}
```

#### 3. Set Implementation

```go
type Set map[string]bool

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(item string) {
    s[item] = true
}

func (s Set) Remove(item string) {
    delete(s, item)
}

func (s Set) Contains(item string) bool {
    return s[item]
}

func (s Set) Size() int {
    return len(s)
}
```

### Nested Maps and Slices

```go
// Map of slices
m1 := map[string][]int{
    "even": {2, 4, 6, 8},
    "odd":  {1, 3, 5, 7},
}

// Map of maps
m2 := map[string]map[string]int{
    "fruits": {
        "apple":  1,
        "banana": 2,
    },
    "vegetables": {
        "carrot": 3,
        "lettuce": 4,
    },
}

// Slice of maps
slice := []map[string]int{
    {"a": 1, "b": 2},
    {"c": 3, "d": 4},
}
```

## Performance Considerations

### Arrays vs Slices

```go
// Arrays: Fixed size, stack allocation (small arrays)
var arr [100]int  // Allocated on stack

// Slices: Dynamic size, heap allocation
slice := make([]int, 100)  // Allocated on heap
```

### Map Performance

```go
// Pre-allocate maps when size is known
m := make(map[string]int, 1000)  // Capacity hint

// Use appropriate key types
// Good: strings, integers
// Avoid: slices, maps (not comparable)
```

### Memory Management

```go
// Reuse slices to avoid allocations
var buffer []byte
for i := 0; i < 1000; i++ {
    buffer = buffer[:0]  // Reset slice, keep capacity
    buffer = append(buffer, "data"...)
}
```

## Best Practices

### 1. Choose the Right Collection

```go
// Use arrays for fixed-size collections
var days [7]string = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

// Use slices for dynamic collections
var numbers []int
numbers = append(numbers, 1, 2, 3)

// Use maps for key-value lookups
var cache map[string]interface{} = make(map[string]interface{})
```

### 2. Initialize Properly

```go
// Good: Initialize with make
slice := make([]int, 0, 10)  // Pre-allocate capacity
m := make(map[string]int, 100)  // Pre-allocate capacity

// Avoid: Nil slices/maps (unless you want nil behavior)
var slice []int  // nil slice
var m map[string]int  // nil map
```

### 3. Handle Nil Collections

```go
func processSlice(slice []int) {
    if slice == nil {
        slice = make([]int, 0)
    }
    // Process slice
}

func processMap(m map[string]int) {
    if m == nil {
        m = make(map[string]int)
    }
    // Process map
}
```

### 4. Use Appropriate Capacity

```go
// Good: Pre-allocate when size is known
slice := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    slice = append(slice, i)
}

// Avoid: Frequent reallocations
var slice []int
for i := 0; i < 1000; i++ {
    slice = append(slice, i)  // May cause multiple reallocations
}
```

### 5. Safe Map Operations

```go
// Always check if key exists
if value, exists := m["key"]; exists {
    // Use value
} else {
    // Handle missing key
}

// Use default values
value := m["key"]  // Returns zero value if key doesn't exist
```

## Summary

Arrays, slices, and maps in Go provide:

- **Arrays**: Fixed-size, value types, good for small, known-size collections
- **Slices**: Dynamic, reference types, most commonly used collection type
- **Maps**: Key-value pairs, fast lookups, unordered collections

Key points to remember:
1. Arrays are fixed-size and value types
2. Slices are dynamic views into arrays
3. Maps provide fast key-value lookups
4. Use `make()` to initialize slices and maps
5. Always check map key existence
6. Pre-allocate capacity when size is known
7. Choose the right collection for your use case

Understanding these collection types and their characteristics will help you write efficient and maintainable Go code. 