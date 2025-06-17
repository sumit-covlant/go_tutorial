# Optional Deep Dives: Advanced Go Topics

## 🎯 Overview

This guide covers advanced Go topics that are typically explored after mastering the core language concepts. These topics are essential for building high-performance, production-ready applications and understanding Go's internals.

---

## 1. Reflection (`reflect` package)

### What is Reflection?

Reflection is the ability of a program to examine, introspect, and modify its own structure and behavior at runtime. Go's `reflect` package provides this capability, allowing you to work with types and values dynamically.

### Key Concepts

**Type vs Value:**
- `reflect.Type`: Represents the type information
- `reflect.Value`: Represents the actual value

**Kinds:**
- Basic types: `Int`, `Float`, `String`, `Bool`
- Composite types: `Struct`, `Array`, `Slice`, `Map`, `Chan`
- Reference types: `Ptr`, `Interface`, `Func`

### Basic Usage

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x int = 42
    v := reflect.ValueOf(x)
    t := reflect.TypeOf(x)
    
    fmt.Printf("Type: %v\n", t)
    fmt.Printf("Value: %v\n", v)
    fmt.Printf("Kind: %v\n", v.Kind())
    fmt.Printf("Can set: %v\n", v.CanSet())
}
```

### Working with Structs

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func inspectStruct(s interface{}) {
    v := reflect.ValueOf(s)
    t := v.Type()
    
    fmt.Printf("Struct: %s\n", t.Name())
    
    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldType := t.Field(i)
        
        fmt.Printf("  %s (%s) = %v\n", 
            fieldType.Name, 
            fieldType.Type, 
            field.Interface())
        
        // Get struct tags
        if tag := fieldType.Tag.Get("json"); tag != "" {
            fmt.Printf("    JSON tag: %s\n", tag)
        }
    }
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    inspectStruct(person)
}
```

### Dynamic Function Calls

```go
func add(a, b int) int {
    return a + b
}

func callFunctionDynamically() {
    // Get function value
    fn := reflect.ValueOf(add)
    
    // Prepare arguments
    args := []reflect.Value{
        reflect.ValueOf(10),
        reflect.ValueOf(20),
    }
    
    // Call function
    results := fn.Call(args)
    
    fmt.Printf("Result: %v\n", results[0].Interface())
}
```

### Creating Values Dynamically

```go
func createSliceDynamically() {
    // Create a slice type
    sliceType := reflect.SliceOf(reflect.TypeOf(0))
    
    // Create a slice value
    slice := reflect.MakeSlice(sliceType, 0, 10)
    
    // Append values
    slice = reflect.Append(slice, reflect.ValueOf(1))
    slice = reflect.Append(slice, reflect.ValueOf(2))
    slice = reflect.Append(slice, reflect.ValueOf(3))
    
    fmt.Printf("Slice: %v\n", slice.Interface())
}
```

### Setting Values

```go
func setValueDynamically() {
    var x int = 10
    v := reflect.ValueOf(&x).Elem() // Get addressable value
    
    if v.CanSet() {
        v.SetInt(42)
        fmt.Printf("New value: %d\n", x)
    }
}
```

### Advanced Patterns

**Generic JSON Marshaler:**
```go
func marshalAny(v interface{}) ([]byte, error) {
    val := reflect.ValueOf(v)
    
    switch val.Kind() {
    case reflect.Struct:
        return json.Marshal(v)
    case reflect.Map:
        return json.Marshal(v)
    case reflect.Slice:
        return json.Marshal(v)
    default:
        return nil, fmt.Errorf("unsupported type: %v", val.Kind())
    }
}
```

**Dependency Injection Container:**
```go
type Container struct {
    services map[reflect.Type]interface{}
}

func (c *Container) Register(service interface{}) {
    t := reflect.TypeOf(service)
    c.services[t] = service
}

func (c *Container) Resolve(target interface{}) error {
    v := reflect.ValueOf(target)
    if v.Kind() != reflect.Ptr {
        return fmt.Errorf("target must be a pointer")
    }
    
    t := v.Elem().Type()
    if service, exists := c.services[t]; exists {
        v.Elem().Set(reflect.ValueOf(service))
        return nil
    }
    
    return fmt.Errorf("service not found: %v", t)
}
```

### Best Practices

**When to Use Reflection:**
- Generic libraries and frameworks
- Serialization/deserialization
- Configuration systems
- Testing frameworks
- Plugin systems

**When NOT to Use Reflection:**
- Performance-critical code
- Simple type conversions
- When static typing is sufficient
- When you can use interfaces instead

**Performance Considerations:**
- Reflection is slower than direct operations
- Cache `reflect.Type` and `reflect.Value` when possible
- Use type assertions when you know the type
- Consider code generation as an alternative

---

## 2. Context API for Managing Goroutine Lifecycle

### What is Context?

Context is a standard way to carry request-scoped values, cancellation signals, and deadlines across API boundaries and between processes. It's essential for managing goroutine lifecycles and preventing goroutine leaks.

### Key Concepts

**Context Interface:**
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

**Context Hierarchy:**
- `context.Background()`: Root context
- `context.TODO()`: Placeholder context
- Derived contexts: `WithCancel`, `WithTimeout`, `WithDeadline`, `WithValue`

### Basic Usage

**Creating Contexts:**
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Background context
    ctx := context.Background()
    
    // Context with cancellation
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    
    // Context with timeout
    ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    // Context with deadline
    deadline := time.Now().Add(10 * time.Second)
    ctx, cancel = context.WithDeadline(ctx, deadline)
    defer cancel()
    
    // Context with values
    ctx = context.WithValue(ctx, "user_id", "123")
}
```

### Cancellation Patterns

**Simple Cancellation:**
```go
func worker(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %s: cancelled\n", name)
            return
        default:
            fmt.Printf("Worker %s: working...\n", name)
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    
    // Start workers
    go worker(ctx, "A")
    go worker(ctx, "B")
    
    // Cancel after 3 seconds
    time.Sleep(3 * time.Second)
    cancel()
    
    // Wait for workers to finish
    time.Sleep(1 * time.Second)
}
```

**Timeout Pattern:**
```go
func fetchData(ctx context.Context, url string) (string, error) {
    // Simulate HTTP request
    select {
    case <-time.After(2 * time.Second):
        return "data", nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    data, err := fetchData(ctx, "https://api.example.com")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Data: %s\n", data)
}
```

### Context Values

**Passing Request-Scoped Data:**
```go
type userKey struct{}

func withUser(ctx context.Context, userID string) context.Context {
    return context.WithValue(ctx, userKey{}, userID)
}

func getUser(ctx context.Context) (string, bool) {
    user, ok := ctx.Value(userKey{}).(string)
    return user, ok
}

func processRequest(ctx context.Context) {
    if userID, ok := getUser(ctx); ok {
        fmt.Printf("Processing request for user: %s\n", userID)
    } else {
        fmt.Println("No user in context")
    }
}

func main() {
    ctx := context.Background()
    ctx = withUser(ctx, "user123")
    processRequest(ctx)
}
```

### Advanced Patterns

**Request Tracing:**
```go
type traceKey struct{}

func withTrace(ctx context.Context, traceID string) context.Context {
    return context.WithValue(ctx, traceKey{}, traceID)
}

func logWithTrace(ctx context.Context, message string) {
    if traceID, ok := ctx.Value(traceKey{}).(string); ok {
        fmt.Printf("[%s] %s\n", traceID, message)
    } else {
        fmt.Printf("[NO_TRACE] %s\n", message)
    }
}

func handleRequest(ctx context.Context) {
    logWithTrace(ctx, "Starting request")
    
    // Simulate work
    time.Sleep(100 * time.Millisecond)
    
    logWithTrace(ctx, "Request completed")
}
```

**Database Transaction Context:**
```go
type txKey struct{}

func withTransaction(ctx context.Context, tx *sql.Tx) context.Context {
    return context.WithValue(ctx, txKey{}, tx)
}

func getTransaction(ctx context.Context) (*sql.Tx, bool) {
    tx, ok := ctx.Value(txKey{}).(*sql.Tx)
    return tx, ok
}

func executeQuery(ctx context.Context, query string) error {
    tx, ok := getTransaction(ctx)
    if !ok {
        return fmt.Errorf("no transaction in context")
    }
    
    _, err := tx.ExecContext(ctx, query)
    return err
}
```

**HTTP Middleware:**
```go
func timeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()
            
            r = r.WithContext(ctx)
            next.ServeHTTP(w, r)
        })
    }
}

func userMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := r.Header.Get("X-User-ID")
        if userID != "" {
            ctx := context.WithValue(r.Context(), "user_id", userID)
            r = r.WithContext(ctx)
        }
        
        next.ServeHTTP(w, r)
    })
}
```

### Best Practices

**Context Guidelines:**
- Always pass context as the first parameter
- Use context for cancellation, timeouts, and request-scoped values
- Don't store context in structs
- Use context keys as unexported types
- Cancel contexts when done

**Common Mistakes:**
```go
// ❌ Don't store context in structs
type Service struct {
    ctx context.Context
}

// ✅ Pass context as parameter
func (s *Service) DoSomething(ctx context.Context) error {
    // Use ctx here
}

// ❌ Don't ignore context cancellation
func worker(ctx context.Context) {
    for {
        // Do work without checking ctx.Done()
    }
}

// ✅ Always check for cancellation
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Do work
        }
    }
}
```

---

## 3. Performance Profiling (`pprof`)

### What is pprof?

pprof is Go's built-in profiling tool that helps you analyze the performance characteristics of your Go programs. It can profile CPU usage, memory allocation, goroutines, and more.

### Types of Profiling

**CPU Profiling:** Measures CPU usage and identifies bottlenecks
**Memory Profiling:** Tracks memory allocation and identifies memory leaks
**Goroutine Profiling:** Shows goroutine usage and potential leaks
**Block Profiling:** Identifies blocking operations
**Mutex Profiling:** Shows mutex contention

### Basic Usage

**HTTP Server with Profiling:**
```go
package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    // Start profiling server
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Your application code
    for {
        expensiveOperation()
        time.Sleep(1 * time.Second)
    }
}

func expensiveOperation() {
    // Simulate expensive work
    time.Sleep(100 * time.Millisecond)
    
    // Simulate memory allocation
    data := make([]byte, 1024*1024)
    _ = data
}
```

### Programmatic Profiling

**CPU Profiling:**
```go
import (
    "os"
    "runtime/pprof"
    "time"
)

func cpuProfile() {
    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    if err := pprof.StartCPUProfile(f); err != nil {
        log.Fatal(err)
    }
    defer pprof.StopCPUProfile()
    
    // Your code here
    expensiveOperation()
}

func expensiveOperation() {
    for i := 0; i < 1000000; i++ {
        // Simulate CPU-intensive work
        _ = i * i
    }
}
```

**Memory Profiling:**
```go
func memoryProfile() {
    f, err := os.Create("memory.prof")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    // Your code here
    allocateMemory()
    
    if err := pprof.WriteHeapProfile(f); err != nil {
        log.Fatal(err)
    }
}

func allocateMemory() {
    var data [][]byte
    for i := 0; i < 1000; i++ {
        data = append(data, make([]byte, 1024))
    }
}
```

### Advanced Profiling

**Custom Profiling Labels:**
```go
import "runtime/pprof"

func labeledProfiling() {
    // Add labels to profiling data
    pprof.Do(context.Background(), pprof.Labels("operation", "database"), func(ctx context.Context) {
        // This code will be labeled in the profile
        databaseOperation()
    })
}

func databaseOperation() {
    time.Sleep(100 * time.Millisecond)
}
```

**Goroutine Profiling:**
```go
func goroutineProfile() {
    f, err := os.Create("goroutine.prof")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    if err := pprof.Lookup("goroutine").WriteTo(f, 0); err != nil {
        log.Fatal(err)
    }
}
```

### Analyzing Profiles

**Using go tool pprof:**
```bash
# CPU profile
go tool pprof cpu.prof

# Memory profile
go tool pprof memory.prof

# Web interface
go tool pprof -http=:8080 cpu.prof

# Top functions
go tool pprof -top cpu.prof

# Graph view
go tool pprof -web cpu.prof
```

**Interactive Commands:**
```
(pprof) top          # Show top functions by CPU/memory
(pprof) list func    # Show source code for function
(pprof) web          # Open web interface
(pprof) traces       # Show call traces
(pprof) help         # Show all commands
```

### Performance Optimization Examples

**Optimizing CPU Usage:**
```go
// Before optimization
func slowFunction() {
    for i := 0; i < 1000000; i++ {
        result := expensiveCalculation(i)
        _ = result
    }
}

// After optimization
func fastFunction() {
    // Pre-allocate slice
    results := make([]int, 1000000)
    
    // Use batch processing
    for i := 0; i < 1000000; i += 1000 {
        batch := results[i:i+1000]
        processBatch(batch)
    }
}
```

**Memory Optimization:**
```go
// Before optimization
func memoryLeak() {
    var data []*LargeObject
    for i := 0; i < 1000; i++ {
        obj := &LargeObject{Data: make([]byte, 1024*1024)}
        data = append(data, obj)
    }
    // data is never cleared
}

// After optimization
func memoryEfficient() {
    for i := 0; i < 1000; i++ {
        obj := &LargeObject{Data: make([]byte, 1024*1024)}
        processObject(obj)
        // obj goes out of scope and can be GC'd
    }
}
```

### Continuous Profiling

**Production Profiling:**
```go
func startContinuousProfiling() {
    go func() {
        ticker := time.NewTicker(30 * time.Second)
        defer ticker.Stop()
        
        for {
            select {
            case <-ticker.C:
                // Take memory profile every 30 seconds
                f, err := os.Create(fmt.Sprintf("memory_%d.prof", time.Now().Unix()))
                if err != nil {
                    log.Printf("Failed to create profile: %v", err)
                    continue
                }
                
                if err := pprof.WriteHeapProfile(f); err != nil {
                    log.Printf("Failed to write profile: %v", err)
                }
                f.Close()
            }
        }
    }()
}
```

### Best Practices

**Profiling Guidelines:**
- Profile in production-like environments
- Use representative workloads
- Profile both CPU and memory
- Use custom labels for better analysis
- Monitor goroutine count and stack traces

**Common Performance Issues:**
- Memory leaks from goroutines
- Excessive memory allocation
- CPU bottlenecks in hot paths
- Blocking operations
- Inefficient algorithms

---

## 4. Go Assembly (for low-level enthusiasts)

### What is Go Assembly?

Go Assembly is a low-level programming language that allows you to write assembly code directly in Go programs. It's used for performance-critical sections, platform-specific optimizations, and interfacing with hardware.

### Key Concepts

**Assembly Files:**
- Use `.s` extension
- Platform-specific (different for AMD64, ARM, etc.)
- Integrated with Go build system

**Assembly Syntax:**
- Go-specific assembly syntax
- Different from traditional assembly
- Platform-independent mnemonics

### Basic Assembly

**Simple Function:**
```go
// main.go
package main

import "fmt"

//go:noescape
func add(x, y int64) int64

func main() {
    result := add(10, 20)
    fmt.Printf("Result: %d\n", result)
}
```

```assembly
// add_amd64.s
TEXT ·add(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    ADDQ BX, AX
    MOVQ AX, ret+16(FP)
    RET
```

### Platform-Specific Assembly

**AMD64 Assembly:**
```assembly
// math_amd64.s
#include "textflag.h"

TEXT ·Multiply(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    IMULQ BX, AX
    MOVQ AX, ret+16(FP)
    RET

TEXT ·Divide(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    CQO                    // Sign extend RAX into RDX
    IDIVQ BX
    MOVQ AX, ret+16(FP)
    RET
```

**ARM64 Assembly:**
```assembly
// math_arm64.s
#include "textflag.h"

TEXT ·Multiply(SB), NOSPLIT, $0-24
    MOVD a+0(FP), R0
    MOVD b+8(FP), R1
    MUL R0, R1, R0
    MOVD R0, ret+16(FP)
    RET

TEXT ·Divide(SB), NOSPLIT, $0-24
    MOVD a+0(FP), R0
    MOVD b+8(FP), R1
    SDIV R0, R1, R0
    MOVD R0, ret+16(FP)
    RET
```

### Advanced Assembly

**SIMD Operations (AMD64):**
```assembly
// simd_amd64.s
#include "textflag.h"

TEXT ·AddVectors(SB), NOSPLIT, $0-48
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ result+16(FP), CX
    MOVQ length+24(FP), DX
    
    // Load 4 floats at a time
    MOVUPS (AX), X0
    MOVUPS (BX), X1
    ADDPS X1, X0
    MOVUPS X0, (CX)
    
    RET
```

**String Operations:**
```assembly
// string_amd64.s
#include "textflag.h"

TEXT ·StringLength(SB), NOSPLIT, $0-24
    MOVQ s_base+0(FP), AX
    MOVQ s_len+8(FP), BX
    
    // Use REP SCASB for fast string length
    MOVQ AX, DI
    MOVQ BX, CX
    MOVB $0, AL
    REPNE SCASB
    
    SUBQ AX, DI
    MOVQ DI, ret+16(FP)
    RET
```

### Assembly Macros and Directives

**Common Directives:**
```assembly
#include "textflag.h"     // Include flag definitions
TEXT ·FunctionName(SB), NOSPLIT, $frame_size-args_size
NOSPLIT                    // Don't check for stack split
RODATA                     // Read-only data section
DATA                       // Data section
GLOBL                      // Global symbol
```

**Conditional Assembly:**
```assembly
// +build amd64

TEXT ·OptimizedFunction(SB), NOSPLIT, $0-16
    // AMD64-specific implementation
    RET

// +build arm64

TEXT ·OptimizedFunction(SB), NOSPLIT, $0-16
    // ARM64-specific implementation
    RET
```

### Performance Optimization

**Optimized Math Functions:**
```assembly
// math_optimized_amd64.s
#include "textflag.h"

TEXT ·FastSqrt(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), X0
    SQRTSD X0, X0
    MOVQ X0, ret+8(FP)
    RET

TEXT ·FastAbs(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), X0
    ANDPD $0x7FFFFFFFFFFFFFFF, X0  // Clear sign bit
    MOVQ X0, ret+8(FP)
    RET
```

**Memory Operations:**
```assembly
// memory_amd64.s
#include "textflag.h"

TEXT ·Memcpy(SB), NOSPLIT, $0-32
    MOVQ dst+0(FP), DI
    MOVQ src+8(FP), SI
    MOVQ n+16(FP), CX
    
    // Use REP MOVSB for fast memory copy
    REP MOVSB
    
    MOVQ dst+0(FP), AX
    MOVQ AX, ret+24(FP)
    RET
```

### Assembly Best Practices

**Guidelines:**
- Only use assembly when necessary
- Profile before and after optimization
- Test on all target platforms
- Document assembly code thoroughly
- Use Go's build constraints for platform-specific code

**Common Patterns:**
```assembly
// Function prologue
TEXT ·FunctionName(SB), NOSPLIT, $0-24
    // Save registers if needed
    PUSHQ BP
    MOVQ SP, BP
    
    // Function body
    MOVQ arg+0(FP), AX
    
    // Restore registers
    POPQ BP
    RET
```

**Error Handling:**
```assembly
TEXT ·SafeDivide(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    
    // Check for division by zero
    TESTQ BX, BX
    JZ divide_by_zero
    
    CQO
    IDIVQ BX
    MOVQ AX, ret+16(FP)
    RET
    
divide_by_zero:
    MOVQ $0, ret+16(FP)
    RET
```

### Integration with Go

**Build Tags:**
```go
// +build amd64

package math

//go:noescape
func FastSqrt(x float64) float64

// +build !amd64

package math

func FastSqrt(x float64) float64 {
    // Fallback implementation
    return math.Sqrt(x)
}
```

**Assembly Functions:**
```go
// Declare assembly functions
//go:noescape
func add(x, y int64) int64

//go:noescape
func multiply(x, y int64) int64

//go:noescape
func divide(x, y int64) int64

//go:noescape
func fastSqrt(x float64) float64
```

### Advanced Examples

**Cryptographic Functions:**
```assembly
// crypto_amd64.s
#include "textflag.h"

TEXT ·SHA256Block(SB), NOSPLIT, $0-32
    MOVQ state+0(FP), DI
    MOVQ data+8(FP), SI
    MOVQ blocks+16(FP), DX
    
    // Use SHA256 instructions if available
    MOVL $1, AX
    CPUID
    TESTL $(1<<29), ECX  // Check SHA bit
    JZ fallback
    
    // SHA256 implementation
    // ... (complex SHA256 assembly)
    RET
    
fallback:
    // Fallback to Go implementation
    RET
```

**Vector Operations:**
```assembly
// vector_amd64.s
#include "textflag.h"

TEXT ·VectorAdd(SB), NOSPLIT, $0-48
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ result+16(FP), CX
    MOVQ length+24(FP), DX
    
    // Process 4 floats at a time
    MOVUPS (AX), X0
    MOVUPS (BX), X1
    ADDPS X1, X0
    MOVUPS X0, (CX)
    
    // Process remaining elements
    SUBQ $4, DX
    JLE done
    
    ADDQ $16, AX
    ADDQ $16, BX
    ADDQ $16, CX
    JMP loop
    
done:
    RET
```

### Debugging Assembly

**Debugging Techniques:**
- Use `go tool objdump` to disassemble
- Add debug symbols with `-gcflags="-S"`
- Use `go tool compile -S` to see generated assembly
- Profile assembly functions

**Example Debugging:**
```bash
# Disassemble binary
go tool objdump -s FunctionName binary

# Show assembly during compilation
go build -gcflags="-S" main.go

# Profile assembly function
go test -bench=BenchmarkFunction -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

---

## 🎯 When to Use These Advanced Topics

### Reflection
- **Use when:** Building generic libraries, serialization frameworks, testing tools
- **Avoid when:** Performance-critical code, simple type conversions
- **Examples:** JSON marshaling, dependency injection, configuration systems

### Context API
- **Use when:** Managing request lifecycles, cancellation, timeouts, request-scoped data
- **Avoid when:** Storing application state, passing data through many layers
- **Examples:** HTTP servers, database operations, gRPC services

### Performance Profiling
- **Use when:** Optimizing performance, debugging memory leaks, identifying bottlenecks
- **Avoid when:** Premature optimization, simple applications
- **Examples:** High-performance services, memory-intensive applications

### Go Assembly
- **Use when:** Performance-critical code, platform-specific optimizations, hardware interfaces
- **Avoid when:** General application code, when Go is fast enough
- **Examples:** Cryptographic functions, math libraries, SIMD operations

---

## 🚀 Getting Started Checklist

### Reflection
- [ ] Understand `reflect.Type` and `reflect.Value`
- [ ] Practice struct inspection and modification
- [ ] Learn dynamic function calls
- [ ] Build a simple serialization library

### Context API
- [ ] Master context creation and cancellation
- [ ] Practice timeout and deadline patterns
- [ ] Learn context value passing
- [ ] Build HTTP middleware with context

### Performance Profiling
- [ ] Set up HTTP profiling server
- [ ] Practice CPU and memory profiling
- [ ] Learn to analyze profiles with `go tool pprof`
- [ ] Optimize a performance-critical function

### Go Assembly
- [ ] Write a simple assembly function
- [ ] Learn platform-specific assembly syntax
- [ ] Practice SIMD operations
- [ ] Build a performance-optimized math library

---

## 📚 Additional Resources

### Reflection
- [Go reflect package documentation](https://golang.org/pkg/reflect/)
- [The Laws of Reflection](https://blog.golang.org/laws-of-reflection)
- [Reflection examples](https://github.com/golang/go/wiki/InterfaceSlice)

### Context API
- [Go context package documentation](https://golang.org/pkg/context/)
- [Context and Cancellation](https://blog.golang.org/context)
- [Context examples](https://github.com/golang/go/wiki/Context)

### Performance Profiling
- [Go pprof documentation](https://golang.org/pkg/runtime/pprof/)
- [Profiling Go Programs](https://blog.golang.org/profiling-go-programs)
- [pprof examples](https://github.com/google/pprof)

### Go Assembly
- [Go assembly documentation](https://golang.org/doc/asm)
- [Assembly examples](https://github.com/golang/go/wiki/Assembly)
- [Platform-specific assembly](https://golang.org/doc/asm#x86)

---

This guide covers the most advanced topics in Go programming. These concepts are typically explored after mastering the core language features and are essential for building high-performance, production-ready applications. Remember to use these tools judiciously and always measure the impact of your optimizations. 