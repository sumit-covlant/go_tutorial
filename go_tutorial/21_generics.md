# Generics in Go (Go 1.18+)

## Overview

Generics, introduced in Go 1.18, allow you to write functions, types, and data structures that can operate on different types while maintaining type safety. Generics enable code reuse, reduce duplication, and improve expressiveness.

---

## Basic Syntax

### Type Parameters
A type parameter is a placeholder for a type, specified in square brackets `[]` after the function or type name.

```go
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```
- `T` is a type parameter.
- `any` is a built-in constraint (alias for `interface{}`) meaning any type.

### Calling a Generic Function
```go
PrintSlice([]int{1, 2, 3})
PrintSlice([]string{"a", "b", "c"})
```

---

## Generic Functions

```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

fmt.Println(Min(3, 5))         // int
fmt.Println(Min(2.5, 1.8))     // float64
fmt.Println(Min("foo", "bar")) // string
```

---

## Type Constraints

Constraints restrict the set of types that can be used as type parameters.

### Built-in Constraints
- `any` — any type
- `comparable` — types that support `==` and `!=`
- `constraints.Ordered` — types that support `<`, `>`, etc. (from `golang.org/x/exp/constraints` or `constraints` in Go 1.20+)

### Custom Constraints
```go
type Stringer interface {
    String() string
}

type HasLength interface {
    ~string | ~[]byte | ~[]rune | ~[]int
}

func PrintLength[T HasLength](v T) {
    fmt.Println(len(v))
}
```
- The `~` operator allows matching underlying types.

---

## Generic Types

### Generic Structs
```go
type Pair[T, U any] struct {
    First  T
    Second U
}

p := Pair[int, string]{First: 1, Second: "hello"}
```

### Generic Methods
```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    n := len(s.items)
    item := s.items[n-1]
    s.items = s.items[:n-1]
    return item
}

s := Stack[int]{}
s.Push(10)
s.Push(20)
fmt.Println(s.Pop()) // 20
```

---

## Type Sets and Union Types

Type sets define which types are allowed for a type parameter.

```go
type Number interface {
    ~int | ~int32 | ~float64
}

func Add[T Number](a, b T) T {
    return a + b
}

fmt.Println(Add(1, 2))       // int
fmt.Println(Add(1.5, 2.5))   // float64
```

---

## Constraints Package

Go 1.20+ includes the `constraints` package in the standard library:
```go
import "constraints"

func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

---

## Real-World Examples

### Generic Map Function
```go
func Map[T any, U any](in []T, f func(T) U) []U {
    out := make([]U, len(in))
    for i, v := range in {
        out[i] = f(v)
    }
    return out
}

squares := Map([]int{1, 2, 3}, func(x int) int { return x * x })
fmt.Println(squares) // [1 4 9]
```

### Generic Filter Function
```go
func Filter[T any](in []T, f func(T) bool) []T {
    out := make([]T, 0)
    for _, v := range in {
        if f(v) {
            out = append(out, v)
        }
    }
    return out
}

evens := Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
fmt.Println(evens) // [2 4]
```

### Generic Set
```go
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
    s := make(Set[T])
    for _, item := range items {
        s[item] = struct{}{}
    }
    return s
}

set := NewSet(1, 2, 3)
fmt.Println(set)
```

---

## Best Practices

- Use generics to reduce code duplication and improve type safety.
- Avoid overusing generics for simple cases; prefer concrete types when possible.
- Use constraints to ensure type safety and meaningful operations.
- Document type parameters and constraints clearly.
- Test generic code with multiple types.

---

## Limitations and Gotchas

- Type parameters cannot be used for struct fields without explicit generic types.
- Type inference works in many cases, but sometimes you must specify type parameters explicitly.
- Reflection with generics can be tricky.
- Generics may increase compile times and binary size in some cases.

---

## Summary

Generics in Go provide:
- Type-safe, reusable code for functions and types
- Reduced code duplication
- Expressive APIs for collections and algorithms

Key points:
1. Use type parameters and constraints for flexibility and safety
2. Prefer generics for data structures and algorithms
3. Avoid unnecessary complexity—use concrete types when generics are not needed
4. Leverage the `constraints` package for common constraints
5. Test and document your generic code

Generics are a powerful addition to Go, enabling you to write more flexible, reusable, and maintainable code. 