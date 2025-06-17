# Go File Handling & I/O

## Overview

Go provides a comprehensive set of packages for file handling and I/O operations. The `os`, `io`, `bufio`, `encoding/json`, and other packages offer powerful tools for reading from and writing to files, network connections, and other data sources.

## Basic File Operations

### Opening Files

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Open file for reading
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Use file...
}
```

### File Modes

```go
// Different file opening modes
file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
defer file.Close()
```

Common file modes:
- `os.O_RDONLY`: Read only
- `os.O_WRONLY`: Write only
- `os.O_RDWR`: Read and write
- `os.O_CREATE`: Create if doesn't exist
- `os.O_APPEND`: Append to file
- `os.O_TRUNC`: Truncate file
- `os.O_EXCL`: Exclusive creation

### File Permissions

```go
// File permissions (Unix-style)
const (
    // Owner permissions
    OwnerRead  = 0400
    OwnerWrite = 0200
    OwnerExec  = 0100
    
    // Group permissions
    GroupRead  = 0040
    GroupWrite = 0020
    GroupExec  = 0010
    
    // Others permissions
    OthersRead  = 0004
    OthersWrite = 0002
    OthersExec  = 0001
)

// Common permission combinations
const (
    ReadWrite     = 0666  // rw-rw-rw-
    ReadWriteExec = 0777  // rwxrwxrwx
    ReadOnly      = 0444  // r--r--r--
)
```

## Reading Files

### Reading Entire File

```go
import (
    "fmt"
    "os"
)

func main() {
    // Read entire file into memory
    data, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("File content: %s\n", string(data))
}
```

### Reading File Line by Line

```go
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lineNumber := 1
    
    for scanner.Scan() {
        fmt.Printf("Line %d: %s\n", lineNumber, scanner.Text())
        lineNumber++
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}
```

### Reading with Buffer

```go
import (
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    buffer := make([]byte, 1024)
    
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Printf("Error reading: %v\n", err)
            return
        }
        
        fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
    }
}
```

### Reading Specific Bytes

```go
import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Seek to specific position
    offset, err := file.Seek(10, 0)  // Seek 10 bytes from beginning
    if err != nil {
        fmt.Printf("Error seeking: %v\n", err)
        return
    }
    
    fmt.Printf("Current position: %d\n", offset)
    
    // Read from current position
    buffer := make([]byte, 20)
    n, err := file.Read(buffer)
    if err != nil {
        fmt.Printf("Error reading: %v\n", err)
        return
    }
    
    fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
}
```

## Writing Files

### Writing Entire File

```go
import (
    "fmt"
    "os"
)

func main() {
    content := "Hello, World!\nThis is a test file."
    
    err := os.WriteFile("output.txt", []byte(content), 0644)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }
    
    fmt.Println("File written successfully")
}
```

### Writing with Buffer

```go
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    writer := bufio.NewWriter(file)
    
    lines := []string{
        "Line 1",
        "Line 2",
        "Line 3",
    }
    
    for _, line := range lines {
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            fmt.Printf("Error writing: %v\n", err)
            return
        }
    }
    
    // Flush buffer to ensure all data is written
    err = writer.Flush()
    if err != nil {
        fmt.Printf("Error flushing: %v\n", err)
        return
    }
    
    fmt.Println("File written successfully")
}
```

### Appending to Files

```go
import (
    "fmt"
    "os"
)

func main() {
    file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    logEntry := "2024-01-01 12:00:00 - Application started\n"
    
    _, err = file.WriteString(logEntry)
    if err != nil {
        fmt.Printf("Error writing: %v\n", err)
        return
    }
    
    fmt.Println("Log entry appended successfully")
}
```

## File Information

### Getting File Info

```go
import (
    "fmt"
    "os"
    "time"
)

func main() {
    fileInfo, err := os.Stat("example.txt")
    if err != nil {
        fmt.Printf("Error getting file info: %v\n", err)
        return
    }
    
    fmt.Printf("Name: %s\n", fileInfo.Name())
    fmt.Printf("Size: %d bytes\n", fileInfo.Size())
    fmt.Printf("Mode: %v\n", fileInfo.Mode())
    fmt.Printf("Modified: %v\n", fileInfo.ModTime())
    fmt.Printf("Is directory: %t\n", fileInfo.IsDir())
}
```

### Checking File Existence

```go
import (
    "fmt"
    "os"
)

func main() {
    // Check if file exists
    if _, err := os.Stat("example.txt"); os.IsNotExist(err) {
        fmt.Println("File does not exist")
    } else {
        fmt.Println("File exists")
    }
    
    // Alternative way
    if _, err := os.Stat("example.txt"); err == nil {
        fmt.Println("File exists")
    } else {
        fmt.Println("File does not exist")
    }
}
```

## Directory Operations

### Reading Directory Contents

```go
import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // Read directory contents
    entries, err := os.ReadDir(".")
    if err != nil {
        fmt.Printf("Error reading directory: %v\n", err)
        return
    }
    
    for _, entry := range entries {
        info, err := entry.Info()
        if err != nil {
            continue
        }
        
        if entry.IsDir() {
            fmt.Printf("Directory: %s\n", entry.Name())
        } else {
            fmt.Printf("File: %s (%d bytes)\n", entry.Name(), info.Size())
        }
    }
}
```

### Walking Directory Tree

```go
import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if info.IsDir() {
            fmt.Printf("Directory: %s\n", path)
        } else {
            fmt.Printf("File: %s (%d bytes)\n", path, info.Size())
        }
        
        return nil
    })
    
    if err != nil {
        fmt.Printf("Error walking directory: %v\n", err)
    }
}
```

### Creating Directories

```go
import (
    "fmt"
    "os"
)

func main() {
    // Create single directory
    err := os.Mkdir("newdir", 0755)
    if err != nil {
        fmt.Printf("Error creating directory: %v\n", err)
        return
    }
    
    // Create nested directories
    err = os.MkdirAll("parent/child/grandchild", 0755)
    if err != nil {
        fmt.Printf("Error creating nested directories: %v\n", err)
        return
    }
    
    fmt.Println("Directories created successfully")
}
```

## File Copying and Moving

### Copying Files

```go
import (
    "fmt"
    "io"
    "os"
)

func main() {
    source, err := os.Open("source.txt")
    if err != nil {
        fmt.Printf("Error opening source: %v\n", err)
        return
    }
    defer source.Close()
    
    destination, err := os.Create("destination.txt")
    if err != nil {
        fmt.Printf("Error creating destination: %v\n", err)
        return
    }
    defer destination.Close()
    
    bytesWritten, err := io.Copy(destination, source)
    if err != nil {
        fmt.Printf("Error copying: %v\n", err)
        return
    }
    
    fmt.Printf("Copied %d bytes\n", bytesWritten)
}
```

### Moving Files

```go
import (
    "fmt"
    "os"
)

func main() {
    err := os.Rename("oldname.txt", "newname.txt")
    if err != nil {
        fmt.Printf("Error renaming file: %v\n", err)
        return
    }
    
    fmt.Println("File renamed successfully")
}
```

## Temporary Files

### Creating Temporary Files

```go
import (
    "fmt"
    "os"
)

func main() {
    // Create temporary file
    tempFile, err := os.CreateTemp("", "prefix_*.txt")
    if err != nil {
        fmt.Printf("Error creating temp file: %v\n", err)
        return
    }
    defer os.Remove(tempFile.Name())  // Clean up
    defer tempFile.Close()
    
    fmt.Printf("Temporary file: %s\n", tempFile.Name())
    
    // Write to temporary file
    _, err = tempFile.WriteString("Temporary content")
    if err != nil {
        fmt.Printf("Error writing: %v\n", err)
        return
    }
}
```

### Creating Temporary Directories

```go
import (
    "fmt"
    "os"
)

func main() {
    // Create temporary directory
    tempDir, err := os.MkdirTemp("", "tempdir_*")
    if err != nil {
        fmt.Printf("Error creating temp directory: %v\n", err)
        return
    }
    defer os.RemoveAll(tempDir)  // Clean up
    
    fmt.Printf("Temporary directory: %s\n", tempDir)
}
```

## JSON File Handling

### Writing JSON to File

```go
import (
    "encoding/json"
    "fmt"
    "os"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}

func main() {
    person := Person{
        Name: "Alice",
        Age:  30,
        City: "New York",
    }
    
    // Write JSON to file
    file, err := os.Create("person.json")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")  // Pretty print
    
    err = encoder.Encode(person)
    if err != nil {
        fmt.Printf("Error encoding JSON: %v\n", err)
        return
    }
    
    fmt.Println("JSON written successfully")
}
```

### Reading JSON from File

```go
import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("person.json")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    var person Person
    decoder := json.NewDecoder(file)
    
    err = decoder.Decode(&person)
    if err != nil {
        fmt.Printf("Error decoding JSON: %v\n", err)
        return
    }
    
    fmt.Printf("Person: %+v\n", person)
}
```

### Reading JSON Array

```go
import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("people.json")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    var people []Person
    decoder := json.NewDecoder(file)
    
    err = decoder.Decode(&people)
    if err != nil {
        fmt.Printf("Error decoding JSON: %v\n", err)
        return
    }
    
    for i, person := range people {
        fmt.Printf("Person %d: %+v\n", i+1, person)
    }
}
```

## CSV File Handling

### Writing CSV Files

```go
import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("data.csv")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // Write header
    header := []string{"Name", "Age", "City"}
    err = writer.Write(header)
    if err != nil {
        fmt.Printf("Error writing header: %v\n", err)
        return
    }
    
    // Write data
    data := [][]string{
        {"Alice", "30", "New York"},
        {"Bob", "25", "Los Angeles"},
        {"Charlie", "35", "Chicago"},
    }
    
    for _, row := range data {
        err = writer.Write(row)
        if err != nil {
            fmt.Printf("Error writing row: %v\n", err)
            return
        }
    }
    
    fmt.Println("CSV written successfully")
}
```

### Reading CSV Files

```go
import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    
    // Read all records
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Printf("Error reading CSV: %v\n", err)
        return
    }
    
    for i, record := range records {
        if i == 0 {
            fmt.Printf("Header: %v\n", record)
        } else {
            fmt.Printf("Row %d: %v\n", i, record)
        }
    }
}
```

## File Handling Best Practices

### 1. Always Close Files

```go
// Good: Use defer to ensure file is closed
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always close the file
    
    // Use file...
    return nil
}
```

### 2. Check for Errors

```go
// Always check for errors
file, err := os.Open("example.txt")
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}
defer file.Close()
```

### 3. Use Buffered I/O for Large Files

```go
// Use buffered I/O for better performance
func copyFile(src, dst string) error {
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()
    
    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()
    
    // Use buffered copy
    _, err = io.Copy(destination, source)
    return err
}
```

### 4. Handle Large Files Efficiently

```go
// Process large files in chunks
func processLargeFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    buffer := make([]byte, 4096)  // 4KB buffer
    
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        
        // Process buffer[:n]
        processChunk(buffer[:n])
    }
    
    return nil
}
```

### 5. Use Context for Cancellation

```go
import "context"

func readFileWithContext(ctx context.Context, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    buffer := make([]byte, 1024)
    
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            n, err := file.Read(buffer)
            if err == io.EOF {
                return nil
            }
            if err != nil {
                return err
            }
            
            // Process buffer[:n]
        }
    }
}
```

## Common File Operations

### File Locking

```go
import (
    "fmt"
    "os"
    "syscall"
)

func lockFile(filename string) (*os.File, error) {
    file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    
    // Lock the file
    err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
    if err != nil {
        file.Close()
        return nil, err
    }
    
    return file, nil
}

func unlockFile(file *os.File) error {
    // Unlock the file
    err := syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
    if err != nil {
        return err
    }
    
    return file.Close()
}
```

### File Monitoring

```go
import (
    "fmt"
    "os"
    "time"
)

func monitorFile(filename string) {
    var lastModTime time.Time
    
    for {
        fileInfo, err := os.Stat(filename)
        if err != nil {
            fmt.Printf("Error checking file: %v\n", err)
            time.Sleep(time.Second)
            continue
        }
        
        if !lastModTime.IsZero() && fileInfo.ModTime().After(lastModTime) {
            fmt.Printf("File %s was modified at %v\n", filename, fileInfo.ModTime())
        }
        
        lastModTime = fileInfo.ModTime()
        time.Sleep(time.Second)
    }
}
```

## Summary

Go file handling provides:

- **Basic Operations**: Open, read, write, close files
- **Directory Operations**: List, create, walk directories
- **File Information**: Get file metadata and properties
- **Buffered I/O**: Efficient reading and writing
- **Format Support**: JSON, CSV, and other formats
- **Error Handling**: Comprehensive error checking
- **Performance**: Optimized for large files

Key points to remember:
1. Always close files using `defer`
2. Check for errors after every operation
3. Use buffered I/O for better performance
4. Handle large files in chunks
5. Use appropriate file permissions
6. Clean up temporary files
7. Use context for cancellation
8. Consider file locking for concurrent access

Understanding file handling is essential for building applications that need to read from and write to the file system efficiently and safely. 