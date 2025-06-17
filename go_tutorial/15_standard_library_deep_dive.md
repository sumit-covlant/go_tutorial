# Go Standard Library Deep Dive

## Overview

Go's standard library is extensive and well-designed, providing a rich set of packages for common programming tasks. Understanding the standard library is crucial for writing efficient and idiomatic Go code.

## Core Packages

### fmt - Formatting and Printing

The `fmt` package provides formatted I/O with functions similar to C's printf and scanf.

```go
package main

import "fmt"

func main() {
    // Basic printing
    fmt.Println("Hello, World!")
    fmt.Printf("Value: %d\n", 42)
    
    // String formatting
    name := "Alice"
    age := 30
    formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
    fmt.Println(formatted)
    
    // Scanning input
    var input string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&input)
    fmt.Printf("Hello, %s!\n", input)
    
    // Error formatting
    err := fmt.Errorf("operation failed: %s", "connection timeout")
    fmt.Println(err)
}
```

### strings - String Manipulation

The `strings` package provides functions for string manipulation.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "  Hello, World!  "
    
    // Trimming whitespace
    trimmed := strings.TrimSpace(text)
    fmt.Printf("Trimmed: '%s'\n", trimmed)
    
    // Case conversion
    upper := strings.ToUpper(text)
    lower := strings.ToLower(text)
    fmt.Printf("Upper: %s\n", upper)
    fmt.Printf("Lower: %s\n", lower)
    
    // Splitting and joining
    parts := strings.Split("apple,banana,cherry", ",")
    fmt.Printf("Parts: %v\n", parts)
    
    joined := strings.Join(parts, " | ")
    fmt.Printf("Joined: %s\n", joined)
    
    // Searching and replacing
    contains := strings.Contains(text, "Hello")
    fmt.Printf("Contains 'Hello': %t\n", contains)
    
    replaced := strings.ReplaceAll(text, "World", "Go")
    fmt.Printf("Replaced: %s\n", replaced)
    
    // Prefix and suffix
    hasPrefix := strings.HasPrefix(text, "  Hello")
    hasSuffix := strings.HasSuffix(text, "!  ")
    fmt.Printf("Has prefix: %t, Has suffix: %t\n", hasPrefix, hasSuffix)
}
```

### strconv - String Conversions

The `strconv` package provides functions for converting strings to and from basic data types.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // String to number conversions
    i, err := strconv.Atoi("123")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Integer: %d\n", i)
    }
    
    f, err := strconv.ParseFloat("3.14", 64)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Float: %f\n", f)
    }
    
    b, err := strconv.ParseBool("true")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Boolean: %t\n", b)
    }
    
    // Number to string conversions
    intStr := strconv.Itoa(123)
    fmt.Printf("Int to string: %s\n", intStr)
    
    floatStr := strconv.FormatFloat(3.14, 'f', 2, 64)
    fmt.Printf("Float to string: %s\n", floatStr)
    
    boolStr := strconv.FormatBool(true)
    fmt.Printf("Bool to string: %s\n", boolStr)
}
```

### time - Time and Date

The `time` package provides functionality for measuring and displaying time.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Current time
    now := time.Now()
    fmt.Printf("Current time: %v\n", now)
    
    // Time components
    fmt.Printf("Year: %d\n", now.Year())
    fmt.Printf("Month: %s\n", now.Month())
    fmt.Printf("Day: %d\n", now.Day())
    fmt.Printf("Hour: %d\n", now.Hour())
    fmt.Printf("Minute: %d\n", now.Minute())
    fmt.Printf("Second: %d\n", now.Second())
    
    // Time formatting
    formatted := now.Format("2006-01-02 15:04:05")
    fmt.Printf("Formatted: %s\n", formatted)
    
    // Custom format
    custom := now.Format("January 2, 2006 at 3:04 PM")
    fmt.Printf("Custom: %s\n", custom)
    
    // Parsing time
    parsed, err := time.Parse("2006-01-02", "2023-12-25")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Parsed: %v\n", parsed)
    }
    
    // Time arithmetic
    tomorrow := now.AddDate(0, 0, 1)
    fmt.Printf("Tomorrow: %v\n", tomorrow)
    
    duration := time.Hour * 2
    future := now.Add(duration)
    fmt.Printf("In 2 hours: %v\n", future)
    
    // Time comparison
    if now.Before(tomorrow) {
        fmt.Println("Now is before tomorrow")
    }
    
    // Sleep
    fmt.Println("Sleeping for 1 second...")
    time.Sleep(time.Second)
    fmt.Println("Awake!")
}
```

### math - Mathematical Functions

The `math` package provides mathematical constants and functions.

```go
package main

import (
    "fmt"
    "math"
    "math/rand"
)

func main() {
    // Mathematical constants
    fmt.Printf("Pi: %f\n", math.Pi)
    fmt.Printf("E: %f\n", math.E)
    
    // Basic functions
    fmt.Printf("Abs(-5): %f\n", math.Abs(-5))
    fmt.Printf("Ceil(3.7): %f\n", math.Ceil(3.7))
    fmt.Printf("Floor(3.7): %f\n", math.Floor(3.7))
    fmt.Printf("Round(3.7): %f\n", math.Round(3.7))
    
    // Power and roots
    fmt.Printf("Pow(2, 3): %f\n", math.Pow(2, 3))
    fmt.Printf("Sqrt(16): %f\n", math.Sqrt(16))
    fmt.Printf("Cbrt(27): %f\n", math.Cbrt(27))
    
    // Trigonometric functions
    fmt.Printf("Sin(π/2): %f\n", math.Sin(math.Pi/2))
    fmt.Printf("Cos(0): %f\n", math.Cos(0))
    fmt.Printf("Tan(π/4): %f\n", math.Tan(math.Pi/4))
    
    // Logarithmic functions
    fmt.Printf("Log(10): %f\n", math.Log(10))
    fmt.Printf("Log10(100): %f\n", math.Log10(100))
    
    // Min and Max
    fmt.Printf("Min(3, 7): %f\n", math.Min(3, 7))
    fmt.Printf("Max(3, 7): %f\n", math.Max(3, 7))
    
    // Random numbers
    rand.Seed(time.Now().UnixNano())
    fmt.Printf("Random int: %d\n", rand.Intn(100))
    fmt.Printf("Random float: %f\n", rand.Float64())
}
```

## Collections and Data Structures

### sort - Sorting

The `sort` package provides primitives for sorting slices and user-defined collections.

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // Sorting integers
    numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
    sort.Ints(numbers)
    fmt.Printf("Sorted numbers: %v\n", numbers)
    
    // Sorting strings
    names := []string{"Charlie", "Alice", "Bob", "David"}
    sort.Strings(names)
    fmt.Printf("Sorted names: %v\n", names)
    
    // Sorting floats
    floats := []float64{3.14, 2.71, 1.41, 2.23}
    sort.Float64s(floats)
    fmt.Printf("Sorted floats: %v\n", floats)
    
    // Custom sorting
    people := []struct {
        Name string
        Age  int
    }{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }
    
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })
    fmt.Printf("Sorted by age: %v\n", people)
    
    // Reverse sorting
    sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
    fmt.Printf("Reverse sorted: %v\n", numbers)
    
    // Checking if sorted
    isSorted := sort.IntsAreSorted([]int{1, 2, 3, 4, 5})
    fmt.Printf("Is sorted: %t\n", isSorted)
}
```

### container - Container Data Structures

The `container` package provides heap, list, and ring data structures.

```go
package main

import (
    "container/heap"
    "container/list"
    "container/ring"
    "fmt"
)

// Heap example
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    // Heap example
    h := &IntHeap{2, 1, 5}
    heap.Init(h)
    heap.Push(h, 3)
    fmt.Printf("Heap minimum: %d\n", (*h)[0])
    
    // List example
    l := list.New()
    l.PushBack("first")
    l.PushFront("second")
    l.PushBack("third")
    
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Printf("List element: %v\n", e.Value)
    }
    
    // Ring example
    r := ring.New(3)
    for i := 0; i < r.Len(); i++ {
        r.Value = i
        r = r.Next()
    }
    
    r.Do(func(x interface{}) {
        fmt.Printf("Ring element: %v\n", x)
    })
}
```

## I/O and File Operations

### io - Basic I/O Interfaces

The `io` package provides basic interfaces to I/O primitives.

```go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    // Reading from strings
    reader := strings.NewReader("Hello, World!")
    buffer := make([]byte, 5)
    
    for {
        n, err := reader.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            break
        }
        fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
    }
    
    // Writing to strings
    var builder strings.Builder
    writer := io.Writer(&builder)
    
    data := []byte("Hello, Go!")
    n, err := writer.Write(data)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Wrote %d bytes: %s\n", n, builder.String())
    }
    
    // Copying between reader and writer
    src := strings.NewReader("Source data")
    dst := &strings.Builder{}
    
    bytesCopied, err := io.Copy(dst, src)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Copied %d bytes: %s\n", bytesCopied, dst.String())
    }
}
```

### bufio - Buffered I/O

The `bufio` package implements buffered I/O.

```go
package main

import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    // Buffered reading
    reader := strings.NewReader("line 1\nline 2\nline 3")
    scanner := bufio.NewScanner(reader)
    
    for scanner.Scan() {
        fmt.Printf("Line: %s\n", scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Buffered writing
    var builder strings.Builder
    writer := bufio.NewWriter(&builder)
    
    writer.WriteString("Hello")
    writer.WriteByte(' ')
    writer.WriteString("World")
    writer.WriteByte('!')
    
    writer.Flush() // Don't forget to flush!
    fmt.Printf("Written: %s\n", builder.String())
    
    // Reading with buffer
    reader2 := strings.NewReader("Hello, World!")
    bufReader := bufio.NewReader(reader2)
    
    // Read until delimiter
    line, err := bufReader.ReadString(',')
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Read until comma: %s\n", line)
    }
    
    // Read remaining
    remaining, err := bufReader.ReadString('\n')
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Remaining: %s\n", remaining)
    }
}
```

### os - Operating System Interface

The `os` package provides a platform-independent interface to operating system functionality.

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // Environment variables
    home := os.Getenv("HOME")
    fmt.Printf("Home directory: %s\n", home)
    
    os.Setenv("MY_VAR", "my_value")
    fmt.Printf("MY_VAR: %s\n", os.Getenv("MY_VAR"))
    
    // Working directory
    wd, err := os.Getwd()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Working directory: %s\n", wd)
    }
    
    // File operations
    tempFile, err := os.CreateTemp("", "example_*.txt")
    if err != nil {
        fmt.Printf("Error creating temp file: %v\n", err)
        return
    }
    defer os.Remove(tempFile.Name())
    
    // Write to file
    content := []byte("Hello, World!")
    _, err = tempFile.Write(content)
    if err != nil {
        fmt.Printf("Error writing to file: %v\n", err)
    }
    tempFile.Close()
    
    // Read from file
    data, err := os.ReadFile(tempFile.Name())
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    } else {
        fmt.Printf("File content: %s\n", string(data))
    }
    
    // File info
    info, err := os.Stat(tempFile.Name())
    if err != nil {
        fmt.Printf("Error getting file info: %v\n", err)
    } else {
        fmt.Printf("File size: %d bytes\n", info.Size())
        fmt.Printf("Is directory: %t\n", info.IsDir())
        fmt.Printf("Mode: %v\n", info.Mode())
    }
    
    // Directory operations
    entries, err := os.ReadDir(".")
    if err != nil {
        fmt.Printf("Error reading directory: %v\n", err)
    } else {
        fmt.Println("Directory entries:")
        for _, entry := range entries {
            fmt.Printf("  %s (dir: %t)\n", entry.Name(), entry.IsDir())
        }
    }
}
```

## Networking and HTTP

### net/http - HTTP Client and Server

The `net/http` package provides HTTP client and server implementations.

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

// HTTP Server
func startServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    
    http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        data := map[string]string{"message": "Hello, JSON!"}
        json.NewEncoder(w).Encode(data)
    })
    
    go http.ListenAndServe(":8080", nil)
    time.Sleep(time.Second) // Give server time to start
}

// HTTP Client
func httpClient() {
    // GET request
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response: %v\n", err)
        return
    }
    
    fmt.Printf("Response: %s\n", string(body))
    
    // POST request
    data := map[string]string{"name": "Alice", "age": "30"}
    jsonData, _ := json.Marshal(data)
    
    resp, err = http.Post("http://httpbin.org/post", 
        "application/json", strings.NewReader(string(jsonData)))
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, err = io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response: %v\n", err)
        return
    }
    
    fmt.Printf("POST response: %s\n", string(body))
}

func main() {
    startServer()
    httpClient()
}
```

### net/url - URL Parsing

The `net/url` package provides functions for parsing URLs and query parameters.

```go
package main

import (
    "fmt"
    "net/url"
)

func main() {
    // Parsing URLs
    u, err := url.Parse("https://example.com/path?param1=value1&param2=value2")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Scheme: %s\n", u.Scheme)
    fmt.Printf("Host: %s\n", u.Host)
    fmt.Printf("Path: %s\n", u.Path)
    fmt.Printf("Raw query: %s\n", u.RawQuery)
    
    // Parsing query parameters
    values, err := url.ParseQuery(u.RawQuery)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Param1: %s\n", values.Get("param1"))
    fmt.Printf("Param2: %s\n", values.Get("param2"))
    
    // Building URLs
    baseURL, _ := url.Parse("https://api.example.com")
    params := url.Values{}
    params.Set("api_key", "secret123")
    params.Set("format", "json")
    
    baseURL.RawQuery = params.Encode()
    fmt.Printf("Built URL: %s\n", baseURL.String())
    
    // URL encoding/decoding
    encoded := url.QueryEscape("Hello, World!")
    fmt.Printf("Encoded: %s\n", encoded)
    
    decoded, err := url.QueryUnescape(encoded)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Decoded: %s\n", decoded)
    }
}
```

## Encoding and Serialization

### encoding/json - JSON Encoding

The `encoding/json` package provides JSON encoding and decoding.

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Email   string   `json:"email,omitempty"`
    Hobbies []string `json:"hobbies"`
}

func main() {
    // Struct to JSON
    person := Person{
        Name:    "Alice",
        Age:     30,
        Hobbies: []string{"reading", "swimming"},
    }
    
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("JSON: %s\n", string(jsonData))
    
    // Pretty printing
    prettyJSON, err := json.MarshalIndent(person, "", "  ")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Pretty JSON:\n%s\n", string(prettyJSON))
    
    // JSON to struct
    jsonStr := `{"name":"Bob","age":25,"hobbies":["gaming","coding"]}`
    var newPerson Person
    
    err = json.Unmarshal([]byte(jsonStr), &newPerson)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Unmarshaled: %+v\n", newPerson)
    
    // JSON with unknown fields
    var data map[string]interface{}
    err = json.Unmarshal([]byte(jsonStr), &data)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Generic data: %+v\n", data)
    
    // JSON streaming
    decoder := json.NewDecoder(strings.NewReader(jsonStr))
    var streamedPerson Person
    err = decoder.Decode(&streamedPerson)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Streamed: %+v\n", streamedPerson)
}
```

### encoding/xml - XML Encoding

The `encoding/xml` package provides XML encoding and decoding.

```go
package main

import (
    "encoding/xml"
    "fmt"
)

type Book struct {
    XMLName xml.Name `xml:"book"`
    Title   string   `xml:"title"`
    Author  string   `xml:"author"`
    Year    int      `xml:"year"`
    Price   float64  `xml:"price"`
}

type Library struct {
    XMLName xml.Name `xml:"library"`
    Books   []Book   `xml:"book"`
}

func main() {
    // Struct to XML
    book := Book{
        Title:  "The Go Programming Language",
        Author: "Alan Donovan",
        Year:   2015,
        Price:  29.99,
    }
    
    xmlData, err := xml.Marshal(book)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("XML: %s\n", string(xmlData))
    
    // Pretty printing
    prettyXML, err := xml.MarshalIndent(book, "", "  ")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Pretty XML:\n%s\n", string(prettyXML))
    
    // XML to struct
    xmlStr := `<book><title>Effective Go</title><author>Go Team</author><year>2020</year><price>0.00</price></book>`
    var newBook Book
    
    err = xml.Unmarshal([]byte(xmlStr), &newBook)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Unmarshaled: %+v\n", newBook)
}
```

## Concurrency and Synchronization

### sync - Synchronization Primitives

The `sync` package provides basic synchronization primitives.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // WaitGroup
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All workers completed")
    
    // Mutex
    var mu sync.Mutex
    counter := 0
    
    for i := 0; i < 10; i++ {
        go func() {
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }
    
    time.Sleep(time.Second)
    fmt.Printf("Counter: %d\n", counter)
    
    // RWMutex
    var rwmu sync.RWMutex
    data := make(map[string]string)
    
    // Writers
    for i := 0; i < 3; i++ {
        go func(id int) {
            rwmu.Lock()
            data[fmt.Sprintf("key%d", id)] = fmt.Sprintf("value%d", id)
            rwmu.Unlock()
        }(i)
    }
    
    // Readers
    for i := 0; i < 5; i++ {
        go func(id int) {
            rwmu.RLock()
            fmt.Printf("Reader %d: data has %d items\n", id, len(data))
            rwmu.RUnlock()
        }(i)
    }
    
    time.Sleep(time.Second)
    
    // Once
    var once sync.Once
    setup := func() {
        fmt.Println("Setup called once")
    }
    
    for i := 0; i < 5; i++ {
        go once.Do(setup)
    }
    
    time.Sleep(time.Second)
}
```

### context - Context Package

The `context` package provides context for deadlines, cancellation, and request-scoped values.

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %s: context cancelled\n", name)
            return
        default:
            fmt.Printf("Worker %s: working...\n", name)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    // Context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    go worker(ctx, "A")
    go worker(ctx, "B")
    
    time.Sleep(3 * time.Second)
    
    // Context with deadline
    deadline := time.Now().Add(1 * time.Second)
    ctx2, cancel2 := context.WithDeadline(context.Background(), deadline)
    defer cancel2()
    
    go worker(ctx2, "C")
    time.Sleep(2 * time.Second)
    
    // Context with values
    ctx3 := context.WithValue(context.Background(), "user_id", "12345")
    ctx3 = context.WithValue(ctx3, "request_id", "req_67890")
    
    userID := ctx3.Value("user_id").(string)
    requestID := ctx3.Value("request_id").(string)
    
    fmt.Printf("User ID: %s, Request ID: %s\n", userID, requestID)
    
    // Context cancellation
    ctx4, cancel4 := context.WithCancel(context.Background())
    
    go func() {
        time.Sleep(1 * time.Second)
        cancel4() // Cancel after 1 second
    }()
    
    go worker(ctx4, "D")
    time.Sleep(2 * time.Second)
}
```

## Summary

The Go standard library provides:

- **Core packages**: fmt, strings, strconv, time, math
- **Collections**: sort, container (heap, list, ring)
- **I/O**: io, bufio, os
- **Networking**: net/http, net/url
- **Encoding**: encoding/json, encoding/xml
- **Concurrency**: sync, context

Key points to remember:
1. The standard library is comprehensive and well-designed
2. Most common programming tasks are covered
3. Packages follow consistent naming and API patterns
4. Documentation is excellent and includes examples
5. Performance is optimized for common use cases
6. Thread safety is clearly documented
7. Error handling is consistent across packages

Understanding the standard library is essential for writing idiomatic and efficient Go code. 