# Go Tooling

## Overview

Go comes with a rich set of built-in tools that make development, testing, and deployment efficient and consistent. This guide covers the essential Go tools and how to use them effectively.

## Go Modules

### go mod init

Initialize a new Go module.

```bash
# Initialize a new module
go mod init myproject

# Initialize with specific module path
go mod init github.com/username/myproject

# Initialize in existing project
cd /path/to/project
go mod init myproject
```

**go.mod file:**
```go
module myproject

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/lib/pq v1.10.9
)

require (
    github.com/felixge/httpsnoop v1.0.3 // indirect
    github.com/gorilla/websocket v1.5.0 // indirect
)
```

### go mod tidy

Clean up and organize dependencies.

```bash
# Add missing dependencies and remove unused ones
go mod tidy

# Verify dependencies
go mod verify

# Download dependencies
go mod download
```

### go mod vendor

Vendor dependencies locally.

```bash
# Create vendor directory with all dependencies
go mod vendor

# Build using vendored dependencies
go build -mod=vendor

# Run tests using vendored dependencies
go test -mod=vendor
```

### go mod graph

View dependency graph.

```bash
# Show dependency graph
go mod graph

# Show why a package is needed
go mod why github.com/gorilla/mux

# Show module information
go mod edit -json
```

## Package Management

### go get

Download and install packages.

```bash
# Install latest version
go get github.com/gorilla/mux

# Install specific version
go get github.com/gorilla/mux@v1.8.0

# Install with specific version constraint
go get github.com/gorilla/mux@latest

# Update to latest version
go get -u github.com/gorilla/mux

# Update all dependencies
go get -u all

# Install for specific build tags
go get -tags=dev github.com/gorilla/mux

# Install and update go.mod
go get -d github.com/gorilla/mux
```

### go install

Install Go programs.

```bash
# Install a program
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install local program
go install ./cmd/myapp

# Install with specific version
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2

# Install with build tags
go install -tags=dev ./cmd/myapp
```

## Building and Running

### go build

Compile Go programs.

```bash
# Build current package
go build

# Build specific file
go build main.go

# Build with output name
go build -o myapp

# Build for specific OS/architecture
go build -o myapp-linux-amd64
GOOS=linux GOARCH=amd64 go build -o myapp

# Build with build tags
go build -tags=dev

# Build with race detector
go build -race

# Build with debug information
go build -gcflags="-N -l"

# Build with optimizations disabled
go build -gcflags="-N -l -S"

# Build with specific go.mod
go build -mod=readonly

# Build with vendor
go build -mod=vendor
```

### go run

Compile and run Go programs.

```bash
# Run main package
go run .

# Run specific file
go run main.go

# Run with arguments
go run main.go arg1 arg2

# Run with environment variables
go run -env=DEBUG=1 main.go

# Run with build tags
go run -tags=dev main.go

# Run with race detector
go run -race main.go

# Run multiple files
go run main.go utils.go
```

### go install

Install and build programs.

```bash
# Install current package
go install

# Install specific package
go install ./cmd/myapp

# Install with build tags
go install -tags=dev ./cmd/myapp

# Install with specific version
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Testing

### go test

Run tests.

```bash
# Run tests in current package
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run tests with coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run tests with race detector
go test -race

# Run specific test
go test -run TestFunctionName

# Run tests with timeout
go test -timeout=30s

# Run tests with build tags
go test -tags=integration

# Run tests in parallel
go test -parallel=4

# Run tests with short flag
go test -short

# Run benchmarks
go test -bench=.

# Run benchmarks with memory info
go test -bench=. -benchmem

# Run benchmarks for specific time
go test -bench=. -benchtime=5s

# Run tests with specific go.mod
go test -mod=readonly

# Run tests with vendor
go test -mod=vendor
```

### go test with examples

```bash
# Run example tests
go test -run Example

# Run example tests with output
go test -run Example -v

# Run specific example
go test -run ExampleFunctionName
```

## Code Quality Tools

### go fmt

Format Go code.

```bash
# Format current package
go fmt

# Format specific file
go fmt main.go

# Format all files in directory
go fmt ./...

# Format with specific options
go fmt -s -w main.go

# Check if files are formatted
go fmt -d main.go
```

### go vet

Analyze Go code for common mistakes.

```bash
# Analyze current package
go vet

# Analyze specific file
go vet main.go

# Analyze all packages
go vet ./...

# Analyze with specific checks
go vet -atomic -bool -buildtag -nilfunc main.go

# Analyze with all checks
go vet -all main.go

# Analyze with shadowing check
go vet -shadow main.go
```

### go lint

Run linters (requires golint).

```bash
# Install golint
go install golang.org/x/lint/golint@latest

# Run golint
golint

# Run golint on specific file
golint main.go

# Run golint on all packages
golint ./...

# Run golint with specific options
golint -set_exit_status ./...
```

### golangci-lint

Run multiple linters.

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run all linters
golangci-lint run

# Run on specific file
golangci-lint run main.go

# Run with specific linters
golangci-lint run --disable-all --enable=errcheck,gosimple

# Run with configuration
golangci-lint run --config=.golangci.yml

# Run with timeout
golangci-lint run --timeout=5m

# Run with specific go version
golangci-lint run --go=1.21
```

**Configuration file (.golangci.yml):**
```yaml
run:
  timeout: 5m
  go: "1.21"

linters:
  enable:
    - gofmt
    - goimports
    - govet
    - errcheck
    - gosimple
    - staticcheck
    - unused
    - misspell
    - gosec

linters-settings:
  govet:
    check-shadowing: true
  gosec:
    excludes:
      - G404

issues:
  exclude-rules:
    - path: _test\.go
      linters:
        - errcheck
```

## Documentation

### go doc

View documentation.

```bash
# View package documentation
go doc fmt

# View function documentation
go doc fmt.Printf

# View type documentation
go doc fmt.Stringer

# View with examples
go doc -ex fmt

# View source code
go doc -src fmt.Printf

# View all symbols
go doc -all fmt

# View in browser
go doc -http=:6060
```

### godoc

Serve documentation locally.

```bash
# Install godoc
go install golang.org/x/tools/cmd/godoc@latest

# Serve documentation
godoc -http=:6060

# Serve with specific workspace
godoc -http=:6060 -goroot=/path/to/go

# Serve with specific workspace mode
godoc -http=:6060 -workspace
```

## Code Generation

### go generate

Run code generation commands.

```bash
# Run all generate commands
go generate

# Run specific generate command
go generate ./...

# Run with specific tags
go generate -tags=dev
```

**Example with stringer:**
```go
//go:generate stringer -type=Pill
type Pill int

const (
    Placebo Pill = iota
    Aspirin
    Ibuprofen
    Paracetamol
    Acetaminophen = Paracetamol
)

// Run: go generate
```

**Example with protobuf:**
```go
//go:generate protoc --go_out=. --go_opt=paths=source_relative proto/*.proto
```

## Profiling and Debugging

### go tool pprof

Profile Go programs.

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof

# Block profiling
go test -blockprofile=block.prof -bench=.
go tool pprof block.prof

# Mutex profiling
go test -mutexprofile=mutex.prof -bench=.
go tool pprof mutex.prof

# HTTP profiling
go tool pprof http://localhost:6060/debug/pprof/heap
```

### go tool trace

Trace Go programs.

```bash
# Generate trace
go test -trace=trace.out -bench=.

# View trace
go tool trace trace.out

# Generate trace for specific test
go test -trace=trace.out -run TestFunctionName
```

## Cross-Compilation

### Building for different platforms

```bash
# Build for Linux AMD64
GOOS=linux GOARCH=amd64 go build -o myapp-linux-amd64

# Build for Windows AMD64
GOOS=windows GOARCH=amd64 go build -o myapp-windows-amd64.exe

# Build for macOS AMD64
GOOS=darwin GOARCH=amd64 go build -o myapp-darwin-amd64

# Build for ARM64
GOOS=linux GOARCH=arm64 go build -o myapp-linux-arm64

# Build for multiple platforms
for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        go build -o myapp-$GOOS-$GOARCH
    done
done
```

## Workspace Management

### go work

Manage Go workspaces.

```bash
# Initialize workspace
go work init

# Add modules to workspace
go work use ./module1
go work use ./module2

# Edit workspace
go work edit -replace=example.com/module=./local/module

# Sync workspace
go work sync
```

**go.work file:**
```go
go 1.21

use (
    ./module1
    ./module2
)

replace example.com/module => ./local/module
```

## Environment and Configuration

### Environment Variables

```bash
# Set Go environment
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# Set Go proxy
export GOPROXY=https://proxy.golang.org,direct

# Set Go private modules
export GOPRIVATE=*.internal.example.com,github.com/mycompany/*

# Set Go sum database
export GOSUMDB=sum.golang.org

# Set Go module mode
export GO111MODULE=on

# Set Go flags
export GOFLAGS=-mod=vendor
```

### go env

View and set Go environment.

```bash
# View all environment variables
go env

# View specific variable
go env GOPATH
go env GOROOT
go env GOPROXY

# Set environment variable
go env -w GOPROXY=https://proxy.golang.org,direct
go env -w GOPRIVATE=*.internal.example.com

# Unset environment variable
go env -u GOPROXY
```

## Build Constraints and Tags

### Build tags

```go
// +build linux,amd64

package main

// +build !windows

package main

//go:build linux && amd64

package main
```

### Using build tags

```bash
# Build with specific tags
go build -tags=dev
go build -tags=production
go build -tags=linux,amd64

# Test with specific tags
go test -tags=integration
go test -tags=unit

# Run with specific tags
go run -tags=dev main.go
```

## Dependency Management

### go mod edit

Edit go.mod file.

```bash
# Add require
go mod edit -require=github.com/gorilla/mux@v1.8.0

# Add replace
go mod edit -replace=github.com/gorilla/mux=./local/mux

# Add exclude
go mod edit -exclude=github.com/gorilla/mux@v1.7.0

# Add retract
go mod edit -retract=v1.0.0

# Add go version
go mod edit -go=1.21

# Add module path
go mod edit -module=myproject
```

### go mod download

Download dependencies.

```bash
# Download all dependencies
go mod download

# Download specific module
go mod download github.com/gorilla/mux

# Download with specific version
go mod download github.com/gorilla/mux@v1.8.0

# Download with specific go version
go mod download -x
```

## Code Analysis Tools

### go list

List packages and modules.

```bash
# List packages
go list ./...

# List with format
go list -f '{{.Dir}}' ./...

# List dependencies
go list -m all

# List specific module
go list -m github.com/gorilla/mux

# List with JSON output
go list -json ./...

# List test files
go list -f '{{.TestGoFiles}}' ./...
```

### go clean

Clean build artifacts.

```bash
# Clean build cache
go clean -cache

# Clean test cache
go clean -testcache

# Clean module cache
go clean -modcache

# Clean all caches
go clean -cache -testcache -modcache

# Clean specific files
go clean -i ./...
```

## Performance Tools

### go tool compile

Compile with specific options.

```bash
# Compile with optimizations
go tool compile -S main.go

# Compile with debug info
go tool compile -N -l main.go

# Compile with race detector
go tool compile -race main.go

# Compile with specific architecture
go tool compile -S -m main.go
```

### go tool link

Link with specific options.

```bash
# Link with debug info
go tool link -w main.o

# Link with specific output
go tool link -o myapp main.o

# Link with specific architecture
go tool link -H linux/amd64 main.o
```

## Summary

Go tooling provides:

- **Module management**: `go mod` for dependency management
- **Building**: `go build`, `go run`, `go install` for compilation
- **Testing**: `go test` for testing and benchmarking
- **Code quality**: `go fmt`, `go vet`, linters for code quality
- **Documentation**: `go doc`, `godoc` for documentation
- **Profiling**: `go tool pprof`, `go tool trace` for performance analysis
- **Cross-compilation**: Build for multiple platforms
- **Workspace management**: `go work` for multi-module development

Key points to remember:
1. Use `go mod` for dependency management
2. Run `go mod tidy` regularly
3. Use `go fmt` to maintain consistent formatting
4. Use `go vet` to catch common mistakes
5. Use linters for additional code quality checks
6. Use `go test` with coverage for testing
7. Use profiling tools for performance optimization
8. Use build tags for conditional compilation
9. Use environment variables for configuration
10. Use workspaces for multi-module development

Understanding Go tooling enables you to write, test, and deploy Go applications efficiently and consistently. 