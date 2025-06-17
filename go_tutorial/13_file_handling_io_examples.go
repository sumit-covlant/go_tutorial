package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// This file demonstrates Go file handling and I/O concepts

func main() {
	fmt.Println("=== Go File Handling & I/O Examples ===\n")

	// Basic file operations
	basicFileOperations()

	// Reading files
	readingFileExamples()

	// Writing files
	writingFileExamples()

	// File information
	fileInformationExamples()

	// Directory operations
	directoryOperations()

	// File copying and moving
	fileCopyingAndMoving()

	// Temporary files
	temporaryFileExamples()

	// JSON file handling
	jsonFileHandling()

	// CSV file handling
	csvFileHandling()

	// Best practices
	bestPracticesExamples()

	// Common file operations
	commonFileOperations()
}

// Basic file operations
func basicFileOperations() {
	fmt.Println("1. Basic File Operations")
	fmt.Println("------------------------")

	// Opening files
	fmt.Println("Opening files:")
	openFileExample()

	// File modes
	fmt.Println("\nFile modes:")
	fileModesExample()

	// File permissions
	fmt.Println("\nFile permissions:")
	filePermissionsExample()
	fmt.Println()
}

// Open file example
func openFileExample() {
	// Open file for reading
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		// Create a sample file for demonstration
		createSampleFile()
		file, err = os.Open("example.txt")
		if err != nil {
			fmt.Printf("Error opening file after creation: %v\n", err)
			return
		}
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}

// Create sample file for examples
func createSampleFile() {
	content := "Hello, World!\nThis is a sample file.\nLine 3\nLine 4\nLine 5"
	err := os.WriteFile("example.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating sample file: %v\n", err)
	} else {
		fmt.Println("Sample file created: example.txt")
	}
}

// File modes example
func fileModesExample() {
	// Different file opening modes
	modes := []struct {
		name string
		flag int
	}{
		{"Read only", os.O_RDONLY},
		{"Write only", os.O_WRONLY},
		{"Read and write", os.O_RDWR},
		{"Create if doesn't exist", os.O_CREATE},
		{"Append to file", os.O_APPEND},
		{"Truncate file", os.O_TRUNC},
	}

	for _, mode := range modes {
		fmt.Printf("- %s: %d\n", mode.name, mode.flag)
	}
}

// File permissions example
func filePermissionsExample() {
	permissions := []struct {
		name string
		mode os.FileMode
	}{
		{"ReadWrite", 0666},
		{"ReadWriteExec", 0777},
		{"ReadOnly", 0444},
		{"Owner read/write, others read", 0644},
	}

	for _, perm := range permissions {
		fmt.Printf("- %s: %o\n", perm.name, perm.mode)
	}
}

// Reading files
func readingFileExamples() {
	fmt.Println("2. Reading Files")
	fmt.Println("----------------")

	// Read entire file
	fmt.Println("Reading entire file:")
	readEntireFileExample()

	// Read file line by line
	fmt.Println("\nReading file line by line:")
	readFileLineByLineExample()

	// Read with buffer
	fmt.Println("\nReading with buffer:")
	readWithBufferExample()

	// Read specific bytes
	fmt.Println("\nReading specific bytes:")
	readSpecificBytesExample()
	fmt.Println()
}

// Read entire file example
func readEntireFileExample() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("File content:\n%s\n", string(data))
}

// Read file line by line example
func readFileLineByLineExample() {
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

// Read with buffer example
func readWithBufferExample() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 10) // 10-byte buffer

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading: %v\n", err)
			return
		}

		fmt.Printf("Read %d bytes: '%s'\n", n, string(buffer[:n]))
	}
}

// Read specific bytes example
func readSpecificBytesExample() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Seek to specific position
	offset, err := file.Seek(10, 0) // Seek 10 bytes from beginning
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

	fmt.Printf("Read %d bytes: '%s'\n", n, string(buffer[:n]))
}

// Writing files
func writingFileExamples() {
	fmt.Println("3. Writing Files")
	fmt.Println("----------------")

	// Write entire file
	fmt.Println("Writing entire file:")
	writeEntireFileExample()

	// Write with buffer
	fmt.Println("\nWriting with buffer:")
	writeWithBufferExample()

	// Append to files
	fmt.Println("\nAppending to files:")
	appendToFileExample()
	fmt.Println()
}

// Write entire file example
func writeEntireFileExample() {
	content := "Hello, World!\nThis is a test file.\nWritten by Go program."

	err := os.WriteFile("output.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("File written successfully: output.txt")
}

// Write with buffer example
func writeWithBufferExample() {
	file, err := os.Create("buffered_output.txt")
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
		"Line 4",
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

	fmt.Println("Buffered file written successfully: buffered_output.txt")
}

// Append to file example
func appendToFileExample() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("%s - Application started\n", timestamp)

	_, err = file.WriteString(logEntry)
	if err != nil {
		fmt.Printf("Error writing: %v\n", err)
		return
	}

	fmt.Println("Log entry appended successfully: log.txt")
}

// File information
func fileInformationExamples() {
	fmt.Println("4. File Information")
	fmt.Println("-------------------")

	// Get file info
	fmt.Println("Getting file info:")
	getFileInfoExample()

	// Check file existence
	fmt.Println("\nChecking file existence:")
	checkFileExistenceExample()
	fmt.Println()
}

// Get file info example
func getFileInfoExample() {
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

// Check file existence example
func checkFileExistenceExample() {
	files := []string{"example.txt", "nonexistent.txt"}

	for _, filename := range files {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Printf("File '%s' does not exist\n", filename)
		} else {
			fmt.Printf("File '%s' exists\n", filename)
		}
	}
}

// Directory operations
func directoryOperations() {
	fmt.Println("5. Directory Operations")
	fmt.Println("----------------------")

	// Read directory contents
	fmt.Println("Reading directory contents:")
	readDirectoryContentsExample()

	// Create directories
	fmt.Println("\nCreating directories:")
	createDirectoriesExample()

	// Walk directory tree
	fmt.Println("\nWalking directory tree:")
	walkDirectoryTreeExample()
	fmt.Println()
}

// Read directory contents example
func readDirectoryContentsExample() {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	fmt.Println("Current directory contents:")
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

// Create directories example
func createDirectoriesExample() {
	// Create single directory
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
	} else {
		fmt.Println("Directory created: newdir")
	}

	// Create nested directories
	err = os.MkdirAll("parent/child/grandchild", 0755)
	if err != nil {
		fmt.Printf("Error creating nested directories: %v\n", err)
	} else {
		fmt.Println("Nested directories created: parent/child/grandchild")
	}
}

// Walk directory tree example
func walkDirectoryTreeExample() {
	fmt.Println("Walking current directory:")
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

// File copying and moving
func fileCopyingAndMoving() {
	fmt.Println("6. File Copying and Moving")
	fmt.Println("--------------------------")

	// Copy files
	fmt.Println("Copying files:")
	copyFileExample()

	// Move files
	fmt.Println("\nMoving files:")
	moveFileExample()
	fmt.Println()
}

// Copy file example
func copyFileExample() {
	// Create source file
	sourceContent := "This is the source file content."
	err := os.WriteFile("source.txt", []byte(sourceContent), 0644)
	if err != nil {
		fmt.Printf("Error creating source file: %v\n", err)
		return
	}

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

	fmt.Printf("Copied %d bytes from source.txt to destination.txt\n", bytesWritten)
}

// Move file example
func moveFileExample() {
	// Create a file to move
	content := "This file will be moved."
	err := os.WriteFile("oldname.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating file to move: %v\n", err)
		return
	}

	err = os.Rename("oldname.txt", "newname.txt")
	if err != nil {
		fmt.Printf("Error renaming file: %v\n", err)
		return
	}

	fmt.Println("File renamed successfully: oldname.txt -> newname.txt")
}

// Temporary files
func temporaryFileExamples() {
	fmt.Println("7. Temporary Files")
	fmt.Println("------------------")

	// Create temporary file
	fmt.Println("Creating temporary file:")
	createTemporaryFileExample()

	// Create temporary directory
	fmt.Println("\nCreating temporary directory:")
	createTemporaryDirectoryExample()
	fmt.Println()
}

// Create temporary file example
func createTemporaryFileExample() {
	// Create temporary file
	tempFile, err := os.CreateTemp("", "prefix_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up
	defer tempFile.Close()

	fmt.Printf("Temporary file: %s\n", tempFile.Name())

	// Write to temporary file
	_, err = tempFile.WriteString("Temporary content")
	if err != nil {
		fmt.Printf("Error writing: %v\n", err)
		return
	}

	fmt.Println("Temporary file written successfully")
}

// Create temporary directory example
func createTemporaryDirectoryExample() {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "tempdir_*")
	if err != nil {
		fmt.Printf("Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir) // Clean up

	fmt.Printf("Temporary directory: %s\n", tempDir)

	// Create a file in the temporary directory
	tempFile := filepath.Join(tempDir, "tempfile.txt")
	err = os.WriteFile(tempFile, []byte("Temporary file content"), 0644)
	if err != nil {
		fmt.Printf("Error creating file in temp directory: %v\n", err)
		return
	}

	fmt.Println("Temporary directory and file created successfully")
}

// JSON file handling
func jsonFileHandling() {
	fmt.Println("8. JSON File Handling")
	fmt.Println("---------------------")

	// Write JSON to file
	fmt.Println("Writing JSON to file:")
	writeJSONToFileExample()

	// Read JSON from file
	fmt.Println("\nReading JSON from file:")
	readJSONFromFileExample()

	// Read JSON array
	fmt.Println("\nReading JSON array:")
	readJSONArrayExample()
	fmt.Println()
}

// Person struct for JSON examples
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

// Write JSON to file example
func writeJSONToFileExample() {
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
	encoder.SetIndent("", "  ") // Pretty print

	err = encoder.Encode(person)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}

	fmt.Println("JSON written successfully: person.json")
}

// Read JSON from file example
func readJSONFromFileExample() {
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

// Read JSON array example
func readJSONArrayExample() {
	// Create JSON array file
	people := []Person{
		{Name: "Alice", Age: 30, City: "New York"},
		{Name: "Bob", Age: 25, City: "Los Angeles"},
		{Name: "Charlie", Age: 35, City: "Chicago"},
	}

	file, err := os.Create("people.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(people)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}

	// Read JSON array
	file, err = os.Open("people.json")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var readPeople []Person
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&readPeople)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	for i, person := range readPeople {
		fmt.Printf("Person %d: %+v\n", i+1, person)
	}
}

// CSV file handling
func csvFileHandling() {
	fmt.Println("9. CSV File Handling")
	fmt.Println("--------------------")

	// Write CSV files
	fmt.Println("Writing CSV files:")
	writeCSVFileExample()

	// Read CSV files
	fmt.Println("\nReading CSV files:")
	readCSVFileExample()
	fmt.Println()
}

// Write CSV file example
func writeCSVFileExample() {
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

	fmt.Println("CSV written successfully: data.csv")
}

// Read CSV file example
func readCSVFileExample() {
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

// Best practices examples
func bestPracticesExamples() {
	fmt.Println("10. Best Practices Examples")
	fmt.Println("---------------------------")

	// Always close files
	fmt.Println("Always close files:")
	alwaysCloseFilesExample()

	// Check for errors
	fmt.Println("\nCheck for errors:")
	checkForErrorsExample()

	// Use buffered I/O for large files
	fmt.Println("\nUse buffered I/O for large files:")
	useBufferedIOExample()

	// Handle large files efficiently
	fmt.Println("\nHandle large files efficiently:")
	handleLargeFilesExample()
	fmt.Println()
}

// Always close files example
func alwaysCloseFilesExample() {
	// Good: Use defer to ensure file is closed
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Always close the file

	fmt.Println("File opened and will be closed automatically")
}

// Check for errors example
func checkForErrorsExample() {
	// Always check for errors
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}

// Use buffered I/O example
func useBufferedIOExample() {
	// Create source and destination files
	sourceContent := "This is the source content for buffered I/O example."
	err := os.WriteFile("source_buffered.txt", []byte(sourceContent), 0644)
	if err != nil {
		fmt.Printf("Error creating source file: %v\n", err)
		return
	}

	source, err := os.Open("source_buffered.txt")
	if err != nil {
		fmt.Printf("Error opening source: %v\n", err)
		return
	}
	defer source.Close()

	destination, err := os.Create("destination_buffered.txt")
	if err != nil {
		fmt.Printf("Error creating destination: %v\n", err)
		return
	}
	defer destination.Close()

	// Use buffered copy
	bytesWritten, err := io.Copy(destination, source)
	if err != nil {
		fmt.Printf("Error copying: %v\n", err)
		return
	}

	fmt.Printf("Buffered copy completed: %d bytes\n", bytesWritten)
}

// Handle large files efficiently example
func handleLargeFilesExample() {
	// Create a large file for demonstration
	largeContent := ""
	for i := 0; i < 1000; i++ {
		largeContent += fmt.Sprintf("Line %d: This is a large file content for demonstration.\n", i)
	}

	err := os.WriteFile("large_file.txt", []byte(largeContent), 0644)
	if err != nil {
		fmt.Printf("Error creating large file: %v\n", err)
		return
	}

	// Process large file in chunks
	file, err := os.Open("large_file.txt")
	if err != nil {
		fmt.Printf("Error opening large file: %v\n", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 4096) // 4KB buffer
	totalBytes := 0
	chunks := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading: %v\n", err)
			return
		}

		totalBytes += n
		chunks++
		// Process buffer[:n] here
	}

	fmt.Printf("Processed large file: %d bytes in %d chunks\n", totalBytes, chunks)
}

// Common file operations
func commonFileOperations() {
	fmt.Println("11. Common File Operations")
	fmt.Println("--------------------------")

	// File monitoring
	fmt.Println("File monitoring:")
	fileMonitoringExample()

	// Safe file operations
	fmt.Println("\nSafe file operations:")
	safeFileOperationsExample()
	fmt.Println()
}

// File monitoring example
func fileMonitoringExample() {
	// Create a file to monitor
	err := os.WriteFile("monitor.txt", []byte("Initial content"), 0644)
	if err != nil {
		fmt.Printf("Error creating file to monitor: %v\n", err)
		return
	}

	var lastModTime time.Time

	// Monitor for a few seconds
	for i := 0; i < 3; i++ {
		fileInfo, err := os.Stat("monitor.txt")
		if err != nil {
			fmt.Printf("Error checking file: %v\n", err)
			time.Sleep(time.Second)
			continue
		}

		if !lastModTime.IsZero() && fileInfo.ModTime().After(lastModTime) {
			fmt.Printf("File monitor.txt was modified at %v\n", fileInfo.ModTime())
		}

		lastModTime = fileInfo.ModTime()
		time.Sleep(time.Second)

		// Modify file after first check
		if i == 0 {
			err = os.WriteFile("monitor.txt", []byte("Modified content"), 0644)
			if err != nil {
				fmt.Printf("Error modifying file: %v\n", err)
			}
		}
	}
}

// Safe file operations example
func safeFileOperationsExample() {
	// Safe file writing with atomic operation
	tempFile := "temp_output.txt"
	finalFile := "final_output.txt"

	// Write to temporary file first
	content := "This is safe content written atomically."
	err := os.WriteFile(tempFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing to temp file: %v\n", err)
		return
	}

	// Atomic move to final location
	err = os.Rename(tempFile, finalFile)
	if err != nil {
		fmt.Printf("Error moving file: %v\n", err)
		// Clean up temp file
		os.Remove(tempFile)
		return
	}

	fmt.Println("File written safely using atomic operation")
}

// Additional utility functions
func processChunk(data []byte) {
	// Process data chunk
	// This is a placeholder for actual processing logic
	_ = data
}

// Clean up function
func cleanup() {
	// Clean up temporary files created during examples
	files := []string{
		"output.txt",
		"buffered_output.txt",
		"log.txt",
		"source.txt",
		"destination.txt",
		"oldname.txt",
		"newname.txt",
		"person.json",
		"people.json",
		"data.csv",
		"source_buffered.txt",
		"destination_buffered.txt",
		"large_file.txt",
		"monitor.txt",
		"final_output.txt",
	}

	for _, file := range files {
		os.Remove(file)
	}

	// Clean up directories
	os.RemoveAll("newdir")
	os.RemoveAll("parent")
}
