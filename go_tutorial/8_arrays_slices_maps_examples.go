package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Go Arrays, Slices, and Maps Examples ===\n")

	// Arrays
	arrayExamples()

	// Slices
	sliceExamples()

	// Maps
	mapExamples()

	// Common patterns
	commonPatterns()

	// Performance considerations
	performanceConsiderations()

	// Best practices
	bestPractices()
}

func arrayExamples() {
	fmt.Println("1. Arrays")
	fmt.Println("----------")

	// Array declaration and initialization
	var arr1 [5]int
	fmt.Printf("Zero-value array: %v\n", arr1)

	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array literal: %v\n", arr2)

	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Short declaration: %v\n", arr3)

	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Compiler-counted array: %v, length: %d\n", arr4, len(arr4))

	arr5 := [5]int{1: 10, 3: 30}
	fmt.Printf("Specific elements: %v\n", arr5)

	// Accessing array elements
	fmt.Printf("First element: %d\n", arr3[0])
	fmt.Printf("Last element: %d\n", arr3[4])

	// Modify elements
	arr3[0] = 100
	fmt.Printf("Modified array: %v\n", arr3)

	// Array properties
	arr6 := [3]int{1, 2, 3}
	arr7 := arr6 // Creates a copy
	arr7[0] = 100
	fmt.Printf("Original: %v\n", arr6)
	fmt.Printf("Copy: %v\n", arr7)

	// Arrays are comparable
	arr8 := [3]int{1, 2, 3}
	arr9 := [3]int{1, 2, 3}
	fmt.Printf("Arrays equal: %t\n", arr8 == arr9)

	// Multi-dimensional arrays
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Matrix: %v\n", matrix)
	fmt.Printf("Element at [1][2]: %d\n", matrix[1][2])

	// Iterating over arrays
	fmt.Println("Iterating over array:")
	for i, value := range arr3 {
		fmt.Printf("arr3[%d] = %d\n", i, value)
	}
	fmt.Println()
}

func sliceExamples() {
	fmt.Println("2. Slices")
	fmt.Println("----------")

	// Slice declaration and initialization
	var slice1 []int
	fmt.Printf("Zero-value slice: %v, len: %d, cap: %d, nil: %t\n",
		slice1, len(slice1), cap(slice1), slice1 == nil)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice literal: %v, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))

	slice3 := make([]int, 5)
	fmt.Printf("Make slice: %v, len: %d, cap: %d\n", slice3, len(slice3), cap(slice3))

	slice4 := make([]int, 3, 5)
	fmt.Printf("Make slice with capacity: %v, len: %d, cap: %d\n", slice4, len(slice4), cap(slice4))

	// From array
	arr := [5]int{1, 2, 3, 4, 5}
	slice5 := arr[1:4]
	slice6 := arr[:3]
	slice7 := arr[2:]
	slice8 := arr[:]

	fmt.Printf("Slice from array [1:4]: %v\n", slice5)
	fmt.Printf("Slice from array [:3]: %v\n", slice6)
	fmt.Printf("Slice from array [2:]: %v\n", slice7)
	fmt.Printf("Slice from array [:]: %v\n", slice8)

	// Slice internals
	fmt.Printf("Original array: %v\n", arr)
	fmt.Printf("Slice: %v, len: %d, cap: %d\n", slice5, len(slice5), cap(slice5))

	// Modifying slices
	slice5[0] = 100
	fmt.Printf("Modified slice: %v\n", slice5)
	fmt.Printf("Original array: %v\n", arr)

	// Extend slice (within capacity)
	slice5 = slice5[:4]
	fmt.Printf("Extended slice: %v\n", slice5)

	// Appending to slices
	slice := []int{}
	fmt.Printf("Empty slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Printf("After append 1: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	slice = append(slice, 2, 3, 4)
	fmt.Printf("After append 2,3,4: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	var slice9 []int
	slice9 = []int{5, 6}
	slice = append(slice, slice9...)
	fmt.Printf("After append slice9: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	// Slice operations
	fmt.Printf("First element: %d\n", slice[0])
	fmt.Printf("Last element: %d\n", slice[len(slice)-1])

	fmt.Printf("slice[1:3]: %v\n", slice[1:3])
	fmt.Printf("slice[:3]: %v\n", slice[:3])
	fmt.Printf("slice[2:]: %v\n", slice[2:])

	// Copy slices
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Printf("Copied slice: %v\n", sliceCopy)

	// Common slice patterns
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Remove element at index
	removed := removeElement(numbers, 2)
	fmt.Printf("After removing index 2: %v\n", removed)

	// Remove last element
	removedLast := removeLast(numbers)
	fmt.Printf("After removing last: %v\n", removedLast)

	// Remove first element
	removedFirst := removeFirst(numbers)
	fmt.Printf("After removing first: %v\n", removedFirst)

	// Filtering
	evenNumbers := filterEven(numbers)
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	filtered := filter(numbers, func(n int) bool { return n > 5 })
	fmt.Printf("Numbers > 5: %v\n", filtered)

	// Mapping
	doubled := double(numbers)
	fmt.Printf("Doubled numbers: %v\n", doubled)

	mapped := mapSlice(numbers, func(n int) int { return n * n })
	fmt.Printf("Squared numbers: %v\n", mapped)

	fmt.Println()
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func removeLast(slice []int) []int {
	return slice[:len(slice)-1]
}

func removeFirst(slice []int) []int {
	return slice[1:]
}

func filterEven(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

func filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func double(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num * 2
	}
	return result
}

func mapSlice(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result
}

func mapExamples() {
	fmt.Println("3. Maps")
	fmt.Println("--------")

	// Map declaration and initialization
	var m1 map[string]int
	fmt.Printf("Zero-value map: %v, nil: %t\n", m1, m1 == nil)

	m2 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	fmt.Printf("Map literal: %v\n", m2)

	m3 := make(map[string]int)
	fmt.Printf("Make map: %v\n", m3)

	m4 := make(map[string]int, 10)
	fmt.Printf("Make map with capacity: %v\n", m4)

	// Map operations
	m := make(map[string]int)

	// Insert or update
	m["apple"] = 1
	m["banana"] = 2
	m["apple"] = 3 // Updates existing key
	fmt.Printf("After insertions: %v\n", m)

	// Access values
	fmt.Printf("apple: %d\n", m["apple"])
	fmt.Printf("banana: %d\n", m["banana"])

	// Check if key exists
	value, exists := m["cherry"]
	if exists {
		fmt.Printf("cherry: %d\n", value)
	} else {
		fmt.Println("cherry not found")
	}

	// Delete key
	delete(m, "banana")
	fmt.Printf("After delete: %v\n", m)

	// Get length
	fmt.Printf("Map length: %d\n", len(m))

	// Map properties
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map1 // Creates a reference, not a copy
	map2["c"] = 3
	fmt.Printf("map1: %v\n", map1)
	fmt.Printf("map2: %v\n", map2)

	// Iterating over maps
	fruits := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	fmt.Println("Iterating over map:")
	for key, value := range fruits {
		fmt.Printf("%s: %d\n", key, value)
	}

	fmt.Println("Keys only:")
	for key := range fruits {
		fmt.Printf("Key: %s\n", key)
	}

	// Common map patterns
	text := "hello world hello go world"
	wordCounts := countWords(text)
	fmt.Printf("Word counts: %v\n", wordCounts)

	// Grouping data
	people := []Person{
		{Name: "Alice", Age: 30, City: "New York"},
		{Name: "Bob", Age: 25, City: "Los Angeles"},
		{Name: "Charlie", Age: 35, City: "New York"},
		{Name: "Diana", Age: 28, City: "Chicago"},
	}

	groups := groupByCity(people)
	fmt.Printf("People grouped by city: %v\n", groups)

	// Set implementation
	set := NewSet()
	set.Add("apple")
	set.Add("banana")
	set.Add("apple") // Duplicate
	fmt.Printf("Set contains apple: %t\n", set.Contains("apple"))
	fmt.Printf("Set contains orange: %t\n", set.Contains("orange"))
	fmt.Printf("Set size: %d\n", set.Size())

	set.Remove("apple")
	fmt.Printf("After removing apple, size: %d\n", set.Size())

	// Nested maps and slices
	nestedMaps := map[string]map[string]int{
		"fruits": {
			"apple":  1,
			"banana": 2,
		},
		"vegetables": {
			"carrot":  3,
			"lettuce": 4,
		},
	}
	fmt.Printf("Nested maps: %v\n", nestedMaps)

	sliceOfMaps := []map[string]int{
		{"a": 1, "b": 2},
		{"c": 3, "d": 4},
	}
	fmt.Printf("Slice of maps: %v\n", sliceOfMaps)
	fmt.Println()
}

type Person struct {
	Name string
	Age  int
	City string
}

func countWords(text string) map[string]int {
	words := strings.Fields(text)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func groupByCity(people []Person) map[string][]Person {
	groups := make(map[string][]Person)

	for _, person := range people {
		groups[person.City] = append(groups[person.City], person)
	}

	return groups
}

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

func commonPatterns() {
	fmt.Println("4. Common Patterns")
	fmt.Println("------------------")

	// Slice patterns
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Reverse slice
	reversed := make([]int, len(numbers))
	for i, j := 0, len(numbers)-1; i < len(numbers); i, j = i+1, j-1 {
		reversed[i] = numbers[j]
	}
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Reversed: %v\n", reversed)

	// Find maximum
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Printf("Maximum: %d\n", max)

	// Map patterns
	scores := map[string]int{
		"Alice":   85,
		"Bob":     92,
		"Charlie": 78,
		"Diana":   95,
	}

	// Find highest scorer
	var topStudent string
	var topScore int
	for student, score := range scores {
		if score > topScore {
			topScore = score
			topStudent = student
		}
	}
	fmt.Printf("Top student: %s with score %d\n", topStudent, topScore)

	// Check if all values are positive
	allPositive := true
	for _, score := range scores {
		if score <= 0 {
			allPositive = false
			break
		}
	}
	fmt.Printf("All scores positive: %t\n", allPositive)

	// Array patterns
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Sum of diagonal
	sum := 0
	for i := 0; i < len(matrix); i++ {
		sum += matrix[i][i]
	}
	fmt.Printf("Sum of diagonal: %d\n", sum)

	// Transpose matrix
	transposed := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transposed[i][j] = matrix[j][i]
		}
	}
	fmt.Printf("Transposed matrix: %v\n", transposed)
	fmt.Println()
}

func performanceConsiderations() {
	fmt.Println("5. Performance Considerations")
	fmt.Println("------------------------------")

	// Pre-allocate slices
	slice := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("Pre-allocated slice length: %d, capacity: %d\n", len(slice), cap(slice))

	// Pre-allocate maps
	m := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key%d", i)] = i
	}
	fmt.Printf("Pre-allocated map size: %d\n", len(m))

	// Reuse slices to avoid allocations
	var buffer []byte
	for i := 0; i < 10; i++ {
		buffer = buffer[:0] // Reset slice, keep capacity
		buffer = append(buffer, fmt.Sprintf("data%d", i)...)
		fmt.Printf("Buffer length: %d, capacity: %d\n", len(buffer), cap(buffer))
	}

	// Use arrays for small, fixed-size collections
	var days [7]string = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("Days of week: %v\n", days)

	fmt.Println()
}

func bestPractices() {
	fmt.Println("6. Best Practices")
	fmt.Println("------------------")

	// Choose the right collection
	var numbers []int
	numbers = append(numbers, 1, 2, 3)
	fmt.Printf("Dynamic slice: %v\n", numbers)

	var cache map[string]interface{} = make(map[string]interface{})
	cache["key1"] = "value1"
	cache["key2"] = 42
	fmt.Printf("Cache: %v\n", cache)

	// Initialize properly
	slice := make([]int, 0, 10)
	m := make(map[string]int, 100)
	fmt.Printf("Properly initialized slice: len=%d, cap=%d\n", len(slice), cap(slice))
	fmt.Printf("Properly initialized map: size=%d\n", len(m))

	// Handle nil collections
	processSlice(nil)
	processMap(nil)

	// Safe map operations
	safeMap := make(map[string]int)
	safeMap["key1"] = 1

	// Always check if key exists
	if value, exists := safeMap["key1"]; exists {
		fmt.Printf("Key exists: %d\n", value)
	} else {
		fmt.Println("Key not found")
	}

	if value, exists := safeMap["key2"]; exists {
		fmt.Printf("Key exists: %d\n", value)
	} else {
		fmt.Println("Key not found")
	}

	// Use default values
	value := safeMap["key2"] // Returns zero value if key doesn't exist
	fmt.Printf("Value (or zero): %d\n", value)

	fmt.Println("Best practices completed.")
	fmt.Println()
}

func processSlice(slice []int) {
	if slice == nil {
		slice = make([]int, 0)
	}
	fmt.Printf("Processing slice with %d elements\n", len(slice))
}

func processMap(m map[string]int) {
	if m == nil {
		m = make(map[string]int)
	}
	fmt.Printf("Processing map with %d elements\n", len(m))
}
