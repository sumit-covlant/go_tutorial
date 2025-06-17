# Go Concurrency

## Overview

Concurrency is one of Go's most powerful features. Go provides goroutines for lightweight concurrent execution and channels for communication between goroutines. The combination of goroutines and channels makes it easy to write concurrent programs that are both efficient and safe.

## Goroutines

### What are Goroutines?

Goroutines are lightweight threads managed by the Go runtime. They are much cheaper than OS threads and can be created in large numbers.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Start a goroutine
    go sayHello("world")
    
    // Main function continues
    fmt.Println("Hello from main")
    
    // Wait for goroutine to complete
    time.Sleep(time.Second)
}

func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

### Creating Goroutines

```go
// Simple goroutine
go functionName()

// Goroutine with anonymous function
go func() {
    fmt.Println("Anonymous goroutine")
}()

// Goroutine with parameters
go func(name string) {
    fmt.Printf("Hello, %s!\n", name)
}("Alice")
```

### Goroutine Lifecycle

```go
func main() {
    // Goroutine starts immediately
    go worker("worker1")
    
    // Main function continues
    fmt.Println("Main function")
    
    // If main exits, all goroutines are terminated
    time.Sleep(time.Second)
}

func worker(name string) {
    fmt.Printf("%s starting\n", name)
    time.Sleep(500 * time.Millisecond)
    fmt.Printf("%s done\n", name)
}
```

## Channels

### What are Channels?

Channels are typed conduits for communication between goroutines. They provide a safe way to share data between concurrent operations.

```go
// Create a channel
ch := make(chan int)

// Send data to channel
ch <- 42

// Receive data from channel
value := <-ch
```

### Channel Types

#### 1. Unbuffered Channels

```go
// Unbuffered channel (synchronous)
ch := make(chan int)

// Sender blocks until receiver is ready
go func() {
    ch <- 42  // Blocks until someone receives
    fmt.Println("Sent 42")
}()

// Receiver blocks until sender is ready
value := <-ch  // Blocks until someone sends
fmt.Printf("Received: %d\n", value)
```

#### 2. Buffered Channels

```go
// Buffered channel (asynchronous)
ch := make(chan int, 3)  // Buffer size of 3

// Can send up to 3 values without blocking
ch <- 1
ch <- 2
ch <- 3

// Fourth send would block
// ch <- 4  // This would block

// Receive values
fmt.Println(<-ch)  // 1
fmt.Println(<-ch)  // 2
fmt.Println(<-ch)  // 3
```

### Channel Operations

#### 1. Send and Receive

```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Hello from goroutine"
    }()
    
    message := <-ch
    fmt.Println(message)
}
```

#### 2. Channel Direction

```go
// Send-only channel
func sendOnly(ch chan<- int) {
    ch <- 42
    // <-ch  // Compile error: cannot receive from send-only channel
}

// Receive-only channel
func receiveOnly(ch <-chan int) {
    value := <-ch
    fmt.Println(value)
    // ch <- 42  // Compile error: cannot send to receive-only channel
}

// Bidirectional channel
func bidirectional(ch chan int) {
    ch <- 42
    value := <-ch
    fmt.Println(value)
}
```

#### 3. Closing Channels

```go
func main() {
    ch := make(chan int)
    
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch)  // Close the channel
    }()
    
    // Receive until channel is closed
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
    }
}
```

### Channel Patterns

#### 1. Producer-Consumer Pattern

```go
func main() {
    ch := make(chan int)
    
    // Producer
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
            time.Sleep(100 * time.Millisecond)
        }
        close(ch)
    }()
    
    // Consumer
    for value := range ch {
        fmt.Printf("Consumed: %d\n", value)
    }
}
```

#### 2. Worker Pool Pattern

```go
func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for i := 0; i < 3; i++ {
        go worker(i, jobs, results)
    }
    
    // Send jobs
    for i := 0; i < 10; i++ {
        jobs <- i
    }
    close(jobs)
    
    // Collect results
    for i := 0; i < 10; i++ {
        result := <-results
        fmt.Printf("Result: %d\n", result)
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(100 * time.Millisecond)
        results <- job * 2
    }
}
```

## Select Statement

### Basic Select

The `select` statement allows a goroutine to wait on multiple channel operations.

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "one"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "two"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Printf("Received: %s\n", msg1)
        case msg2 := <-ch2:
            fmt.Printf("Received: %s\n", msg2)
        }
    }
}
```

### Select with Default

```go
func main() {
    ch := make(chan string)
    
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    default:
        fmt.Println("No message received")
    }
}
```

### Select with Timeout

```go
func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "result"
    }()
    
    select {
    case result := <-ch:
        fmt.Printf("Received: %s\n", result)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout")
    }
}
```

## Sync Package

### WaitGroup

`WaitGroup` is used to wait for a collection of goroutines to finish.

```go
import "sync"

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        wg.Add(1)  // Increment counter
        go func(id int) {
            defer wg.Done()  // Decrement counter when done
            worker(id)
        }(i)
    }
    
    wg.Wait()  // Wait for all goroutines to finish
    fmt.Println("All workers completed")
}

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}
```

### Mutex

`Mutex` provides mutual exclusion for shared resources.

```go
import "sync"

type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *Counter) GetCount() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

func main() {
    counter := &Counter{}
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Final count: %d\n", counter.GetCount())
}
```

### RWMutex

`RWMutex` allows multiple readers or a single writer.

```go
type DataStore struct {
    mu    sync.RWMutex
    data  map[string]string
}

func (ds *DataStore) Set(key, value string) {
    ds.mu.Lock()
    defer ds.mu.Unlock()
    ds.data[key] = value
}

func (ds *DataStore) Get(key string) (string, bool) {
    ds.mu.RLock()
    defer ds.mu.RUnlock()
    value, exists := ds.data[key]
    return value, exists
}

func main() {
    store := &DataStore{data: make(map[string]string)}
    var wg sync.WaitGroup
    
    // Writers
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            store.Set(fmt.Sprintf("key%d", id), fmt.Sprintf("value%d", id))
        }(i)
    }
    
    // Readers
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            value, exists := store.Get(fmt.Sprintf("key%d", id%5))
            if exists {
                fmt.Printf("Read: %s\n", value)
            }
        }(i)
    }
    
    wg.Wait()
}
```

### Once

`Once` ensures a function is called only once.

```go
import "sync"

var (
    once sync.Once
    instance *Singleton
)

type Singleton struct {
    data string
}

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{data: "initialized"}
        fmt.Println("Singleton initialized")
    })
    return instance
}

func main() {
    // Multiple calls to GetInstance
    for i := 0; i < 5; i++ {
        go func(id int) {
            instance := GetInstance()
            fmt.Printf("Goroutine %d got instance: %s\n", id, instance.data)
        }(i)
    }
    
    time.Sleep(time.Second)
}
```

## Context Package

### Basic Context Usage

Context is used to carry deadlines, cancellation signals, and request-scoped values.

```go
import "context"

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    go func() {
        time.Sleep(2 * time.Second)
        cancel()  // Cancel the context
    }()
    
    select {
    case <-ctx.Done():
        fmt.Println("Context cancelled")
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout")
    }
}
```

### Context with Timeout

```go
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    go func() {
        time.Sleep(2 * time.Second)
        fmt.Println("Work completed")
    }()
    
    select {
    case <-ctx.Done():
        fmt.Println("Context timeout")
    }
}
```

### Context with Values

```go
func main() {
    ctx := context.WithValue(context.Background(), "user", "alice")
    
    go workerWithContext(ctx)
    
    time.Sleep(time.Second)
}

func workerWithContext(ctx context.Context) {
    user := ctx.Value("user").(string)
    fmt.Printf("Working for user: %s\n", user)
    
    select {
    case <-ctx.Done():
        fmt.Println("Context cancelled")
    case <-time.After(500 * time.Millisecond):
        fmt.Println("Work completed")
    }
}
```

## Common Concurrency Patterns

### 1. Fan-Out, Fan-In

```go
func main() {
    numbers := generate(1, 2, 3, 4, 5)
    
    // Fan-out: distribute work across multiple goroutines
    c1 := square(numbers)
    c2 := square(numbers)
    
    // Fan-in: combine results
    result := merge(c1, c2)
    
    for value := range result {
        fmt.Printf("Result: %d\n", value)
    }
}

func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

func merge(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for value := range c {
                out <- value
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}
```

### 2. Pipeline Pattern

```go
func main() {
    numbers := generateNumbers(1, 2, 3, 4, 5)
    squared := squareNumbers(numbers)
    filtered := filterEven(squared)
    
    for result := range filtered {
        fmt.Printf("Pipeline result: %d\n", result)
    }
}

func generateNumbers(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func squareNumbers(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

func filterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            if n%2 == 0 {
                out <- n
            }
        }
    }()
    return out
}
```

### 3. Rate Limiting

```go
func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)
    
    limiter := time.Tick(200 * time.Millisecond)
    
    for req := range requests {
        <-limiter  // Rate limit
        fmt.Printf("Processing request %d\n", req)
    }
}
```

## Concurrency Best Practices

### 1. Avoid Goroutine Leaks

```go
// Bad: Goroutine leak
func badExample() {
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    // Goroutine is blocked forever
}

// Good: Proper cleanup
func goodExample() {
    ch := make(chan int)
    go func() {
        defer close(ch)
        ch <- 42
    }()
    
    value := <-ch
    fmt.Println(value)
}
```

### 2. Use Buffered Channels Appropriately

```go
// Use buffered channels when you know the number of sends
func workerPool(jobs []int) {
    results := make(chan int, len(jobs))  // Buffer for all results
    
    for _, job := range jobs {
        go func(j int) {
            results <- process(j)
        }(job)
    }
    
    for i := 0; i < len(jobs); i++ {
        fmt.Println(<-results)
    }
}
```

### 3. Handle Channel Closing

```go
// Only the sender should close a channel
func sender(ch chan<- int) {
    defer close(ch)
    for i := 0; i < 5; i++ {
        ch <- i
    }
}

// Receivers should check if channel is closed
func receiver(ch <-chan int) {
    for value := range ch {
        fmt.Println(value)
    }
    // Channel is automatically closed when loop exits
}
```

### 4. Use Context for Cancellation

```go
func workerWithCancellation(ctx context.Context, ch <-chan int) {
    for {
        select {
        case value := <-ch:
            fmt.Printf("Processing: %d\n", value)
        case <-ctx.Done():
            fmt.Println("Worker cancelled")
            return
        }
    }
}
```

## Common Pitfalls

### 1. Race Conditions

```go
// Bad: Race condition
var counter int

func increment() {
    counter++  // Race condition
}

// Good: Use mutex
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (sc *SafeCounter) Increment() {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    sc.count++
}
```

### 2. Deadlocks

```go
// Bad: Deadlock
func deadlock() {
    ch := make(chan int)
    ch <- 42  // Blocks forever
    value := <-ch
    fmt.Println(value)
}

// Good: Use goroutines
func noDeadlock() {
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    value := <-ch
    fmt.Println(value)
}
```

### 3. Goroutine Leaks

```go
// Bad: Goroutine leak
func leak() {
    ch := make(chan int)
    go func() {
        for {
            ch <- 1
        }
    }()
    // Goroutine runs forever
}

// Good: Proper cleanup
func noLeak() {
    ch := make(chan int)
    done := make(chan bool)
    
    go func() {
        defer close(ch)
        for {
            select {
            case ch <- 1:
            case <-done:
                return
            }
        }
    }()
    
    // Do some work
    time.Sleep(time.Second)
    
    // Signal cleanup
    close(done)
}
```

## Summary

Go concurrency provides:

- **Goroutines**: Lightweight threads for concurrent execution
- **Channels**: Safe communication between goroutines
- **Select**: Multi-channel operations
- **Sync Package**: Synchronization primitives
- **Context**: Request-scoped values and cancellation

Key points to remember:
1. Use goroutines for concurrent execution
2. Use channels for communication between goroutines
3. Use `select` for multi-channel operations
4. Use `WaitGroup` to wait for goroutines to finish
5. Use `Mutex` for shared resource protection
6. Use `Context` for cancellation and timeouts
7. Avoid race conditions and deadlocks
8. Clean up goroutines to prevent leaks

Understanding concurrency is essential for writing efficient, scalable Go programs that can handle multiple operations simultaneously. 