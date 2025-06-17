package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// This file demonstrates Go concurrency concepts

func main() {
	fmt.Println("=== Go Concurrency Examples ===\n")

	// Basic goroutine examples
	basicGoroutineExamples()

	// Channel examples
	channelExamples()

	// Select statement examples
	selectStatementExamples()

	// Sync package examples
	syncPackageExamples()

	// Context examples
	contextExamples()

	// Common concurrency patterns
	commonConcurrencyPatterns()

	// Best practices examples
	bestPracticesExamples()

	// Common pitfalls examples
	commonPitfallsExamples()
}

// Basic goroutine examples
func basicGoroutineExamples() {
	fmt.Println("1. Basic Goroutine Examples")
	fmt.Println("---------------------------")

	// Simple goroutine
	fmt.Println("Starting simple goroutine...")
	go sayHello("world")
	time.Sleep(100 * time.Millisecond)

	// Goroutine with anonymous function
	fmt.Println("\nStarting anonymous goroutine...")
	go func() {
		fmt.Println("Anonymous goroutine executing")
	}()
	time.Sleep(100 * time.Millisecond)

	// Goroutine with parameters
	fmt.Println("\nStarting goroutine with parameters...")
	go func(name string) {
		fmt.Printf("Hello, %s from goroutine!\n", name)
	}("Alice")
	time.Sleep(100 * time.Millisecond)

	// Multiple goroutines
	fmt.Println("\nStarting multiple goroutines...")
	for i := 1; i <= 3; i++ {
		go worker(fmt.Sprintf("worker%d", i))
	}
	time.Sleep(1 * time.Second)
	fmt.Println()
}

// Simple function for goroutines
func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Worker function
func worker(name string) {
	fmt.Printf("%s starting\n", name)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%s done\n", name)
}

// Channel examples
func channelExamples() {
	fmt.Println("2. Channel Examples")
	fmt.Println("-------------------")

	// Unbuffered channel
	fmt.Println("Unbuffered channel example:")
	unbufferedChannelExample()

	// Buffered channel
	fmt.Println("\nBuffered channel example:")
	bufferedChannelExample()

	// Channel direction
	fmt.Println("\nChannel direction example:")
	channelDirectionExample()

	// Closing channels
	fmt.Println("\nClosing channels example:")
	closingChannelsExample()

	// Producer-consumer pattern
	fmt.Println("\nProducer-consumer pattern:")
	producerConsumerExample()

	// Worker pool pattern
	fmt.Println("\nWorker pool pattern:")
	workerPoolExample()
	fmt.Println()
}

// Unbuffered channel example
func unbufferedChannelExample() {
	ch := make(chan int)

	// Sender goroutine
	go func() {
		fmt.Println("Sending value to channel...")
		ch <- 42
		fmt.Println("Value sent to channel")
	}()

	// Receiver (main goroutine)
	fmt.Println("Waiting to receive value...")
	value := <-ch
	fmt.Printf("Received: %d\n", value)
}

// Buffered channel example
func bufferedChannelExample() {
	ch := make(chan int, 3) // Buffer size of 3

	// Can send multiple values without blocking
	fmt.Println("Sending values to buffered channel...")
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("All values sent")

	// Receive values
	fmt.Println("Receiving values:")
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3
}

// Channel direction example
func channelDirectionExample() {
	ch := make(chan int)

	// Send-only function
	go sendOnly(ch)

	// Receive-only function
	receiveOnly(ch)
}

// Send-only channel function
func sendOnly(ch chan<- int) {
	fmt.Println("Sending value to send-only channel")
	ch <- 42
}

// Receive-only channel function
func receiveOnly(ch <-chan int) {
	fmt.Println("Receiving from receive-only channel")
	value := <-ch
	fmt.Printf("Received: %d\n", value)
}

// Closing channels example
func closingChannelsExample() {
	ch := make(chan int)

	// Sender goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch) // Close the channel
		fmt.Println("Channel closed")
	}()

	// Receiver
	fmt.Println("Receiving values until channel closes:")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

// Producer-consumer pattern
func producerConsumerExample() {
	ch := make(chan int)

	// Producer
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Producing: %d\n", i)
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	// Consumer
	for value := range ch {
		fmt.Printf("Consuming: %d\n", value)
	}
}

// Worker pool pattern
func workerPoolExample() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results
	for i := 0; i < 5; i++ {
		result := <-results
		fmt.Printf("Job result: %d\n", result)
	}
}

// Worker function for worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

// Select statement examples
func selectStatementExamples() {
	fmt.Println("3. Select Statement Examples")
	fmt.Println("----------------------------")

	// Basic select
	fmt.Println("Basic select example:")
	basicSelectExample()

	// Select with default
	fmt.Println("\nSelect with default example:")
	selectWithDefaultExample()

	// Select with timeout
	fmt.Println("\nSelect with timeout example:")
	selectWithTimeoutExample()

	// Select with multiple channels
	fmt.Println("\nSelect with multiple channels:")
	selectMultipleChannelsExample()
	fmt.Println()
}

// Basic select example
func basicSelectExample() {
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
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		}
	}
}

// Select with default example
func selectWithDefaultExample() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message received (default case)")
	}
}

// Select with timeout example
func selectWithTimeoutExample() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout occurred")
	}
}

// Select with multiple channels example
func selectMultipleChannelsExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received: %s\n", msg2)
		case <-done:
			fmt.Println("Done")
			return
		}
	}
}

// Sync package examples
func syncPackageExamples() {
	fmt.Println("4. Sync Package Examples")
	fmt.Println("------------------------")

	// WaitGroup example
	fmt.Println("WaitGroup example:")
	waitGroupExample()

	// Mutex example
	fmt.Println("\nMutex example:")
	mutexExample()

	// RWMutex example
	fmt.Println("\nRWMutex example:")
	rwMutexExample()

	// Once example
	fmt.Println("\nOnce example:")
	onceExample()
	fmt.Println()
}

// WaitGroup example
func waitGroupExample() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // Increment counter
		go func(id int) {
			defer wg.Done() // Decrement counter when done
			worker(fmt.Sprintf("worker%d", id))
		}(i)
	}

	fmt.Println("Waiting for all workers to complete...")
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers completed")
}

// Mutex example
func mutexExample() {
	counter := &Counter{}
	var wg sync.WaitGroup

	// Start multiple goroutines to increment counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.GetCount())
}

// Counter with mutex
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

// RWMutex example
func rwMutexExample() {
	store := &DataStore{data: make(map[string]string)}
	var wg sync.WaitGroup

	// Writers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			store.Set(key, value)
			fmt.Printf("Set %s = %s\n", key, value)
		}(i)
	}

	// Readers
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id%3)
			value, exists := store.Get(key)
			if exists {
				fmt.Printf("Read %s = %s\n", key, value)
			}
		}(i)
	}

	wg.Wait()
}

// DataStore with RWMutex
type DataStore struct {
	mu   sync.RWMutex
	data map[string]string
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

// Once example
func onceExample() {
	var once sync.Once
	var wg sync.WaitGroup

	// Multiple calls to GetInstance
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			instance := GetInstance()
			fmt.Printf("Goroutine %d got instance: %s\n", id, instance.data)
		}(i)
	}

	wg.Wait()
}

// Singleton with Once
var (
	once     sync.Once
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

// Context examples
func contextExamples() {
	fmt.Println("5. Context Examples")
	fmt.Println("-------------------")

	// Basic context usage
	fmt.Println("Basic context example:")
	basicContextExample()

	// Context with timeout
	fmt.Println("\nContext with timeout example:")
	contextWithTimeoutExample()

	// Context with values
	fmt.Println("\nContext with values example:")
	contextWithValuesExample()

	// Context cancellation
	fmt.Println("\nContext cancellation example:")
	contextCancellationExample()
	fmt.Println()
}

// Basic context example
func basicContextExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(2 * time.Second)
		cancel() // Cancel the context
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}

// Context with timeout example
func contextWithTimeoutExample() {
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

// Context with values example
func contextWithValuesExample() {
	ctx := context.WithValue(context.Background(), "user", "alice")

	go workerWithContext(ctx)

	time.Sleep(time.Second)
}

// Worker with context
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

// Context cancellation example
func contextCancellationExample() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling context...")
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context was cancelled")
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}

// Common concurrency patterns
func commonConcurrencyPatterns() {
	fmt.Println("6. Common Concurrency Patterns")
	fmt.Println("------------------------------")

	// Fan-out, Fan-in pattern
	fmt.Println("Fan-out, Fan-in pattern:")
	fanOutFanInExample()

	// Pipeline pattern
	fmt.Println("\nPipeline pattern:")
	pipelineExample()

	// Rate limiting pattern
	fmt.Println("\nRate limiting pattern:")
	rateLimitingExample()

	// Worker pool with context
	fmt.Println("\nWorker pool with context:")
	workerPoolWithContextExample()
	fmt.Println()
}

// Fan-out, Fan-in pattern
func fanOutFanInExample() {
	numbers := generate(1, 2, 3, 4, 5)

	// Fan-out: distribute work across multiple goroutines
	c1 := square(numbers)
	c2 := square(numbers)

	// Fan-in: combine results
	result := merge(c1, c2)

	for value := range result {
		fmt.Printf("Fan-out/Fan-in result: %d\n", value)
	}
}

// Generate numbers
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

// Square numbers
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

// Merge channels
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

// Pipeline pattern
func pipelineExample() {
	numbers := generateNumbers(1, 2, 3, 4, 5)
	squared := squareNumbers(numbers)
	filtered := filterEven(squared)

	for result := range filtered {
		fmt.Printf("Pipeline result: %d\n", result)
	}
}

// Generate numbers for pipeline
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

// Square numbers for pipeline
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

// Filter even numbers
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

// Rate limiting pattern
func rateLimitingExample() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter // Rate limit
		fmt.Printf("Processing request %d\n", req)
	}
}

// Worker pool with context
func workerPoolWithContextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	for i := 0; i < 3; i++ {
		go workerWithContext(ctx, i, jobs, results)
	}

	// Send jobs
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case jobs <- i:
				fmt.Printf("Sent job %d\n", i)
			case <-ctx.Done():
				fmt.Println("Context cancelled, stopping job sending")
				return
			}
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case result := <-results:
				fmt.Printf("Received result: %d\n", result)
			case <-ctx.Done():
				fmt.Println("Context cancelled, stopping result collection")
				return
			}
		}
	}()

	// Wait for context to be cancelled
	<-ctx.Done()
	fmt.Println("Worker pool example completed")
}

// Worker with context
func workerWithContext(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case job := <-jobs:
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(100 * time.Millisecond)
			results <- job * 2
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled\n", id)
			return
		}
	}
}

// Best practices examples
func bestPracticesExamples() {
	fmt.Println("7. Best Practices Examples")
	fmt.Println("--------------------------")

	// Avoid goroutine leaks
	fmt.Println("Avoiding goroutine leaks:")
	avoidGoroutineLeaksExample()

	// Use buffered channels appropriately
	fmt.Println("\nUsing buffered channels appropriately:")
	useBufferedChannelsExample()

	// Handle channel closing
	fmt.Println("\nHandling channel closing:")
	handleChannelClosingExample()

	// Use context for cancellation
	fmt.Println("\nUsing context for cancellation:")
	useContextForCancellationExample()
	fmt.Println()
}

// Avoid goroutine leaks example
func avoidGoroutineLeaksExample() {
	// Good: Proper cleanup
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			select {
			case ch <- i:
				fmt.Printf("Sent: %d\n", i)
			case <-done:
				fmt.Println("Goroutine cleanup")
				return
			}
		}
	}()

	// Receive values
	for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Printf("Received: %d\n", value)
	}

	// Signal cleanup
	close(done)
	time.Sleep(100 * time.Millisecond)
}

// Use buffered channels appropriately
func useBufferedChannelsExample() {
	jobs := []int{1, 2, 3, 4, 5}
	results := make(chan int, len(jobs)) // Buffer for all results

	for _, job := range jobs {
		go func(j int) {
			results <- process(j)
		}(job)
	}

	for i := 0; i < len(jobs); i++ {
		result := <-results
		fmt.Printf("Job result: %d\n", result)
	}
}

// Process function
func process(job int) int {
	time.Sleep(100 * time.Millisecond)
	return job * 2
}

// Handle channel closing example
func handleChannelClosingExample() {
	ch := make(chan int)

	// Sender
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Receiver
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println("Channel closed")
}

// Use context for cancellation example
func useContextForCancellationExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
				fmt.Printf("Sent: %d\n", i)
				time.Sleep(100 * time.Millisecond)
			case <-ctx.Done():
				fmt.Println("Sender cancelled")
				return
			}
		}
		close(ch)
	}()

	// Cancel after some time
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Cancelling context...")
		cancel()
	}()

	// Receive values
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

// Common pitfalls examples
func commonPitfallsExamples() {
	fmt.Println("8. Common Pitfalls Examples")
	fmt.Println("---------------------------")

	// Race condition example
	fmt.Println("Race condition example:")
	raceConditionExample()

	// Deadlock example
	fmt.Println("\nDeadlock example:")
	deadlockExample()

	// Goroutine leak example
	fmt.Println("\nGoroutine leak example:")
	goroutineLeakExample()

	// Safe alternatives
	fmt.Println("\nSafe alternatives:")
	safeAlternativesExample()
	fmt.Println()
}

// Race condition example
func raceConditionExample() {
	var counter int
	var wg sync.WaitGroup

	// Bad: Race condition
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // Race condition
		}()
	}

	wg.Wait()
	fmt.Printf("Counter (race condition): %d\n", counter)

	// Good: Use mutex
	safeCounter := &SafeCounter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeCounter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Safe counter: %d\n", safeCounter.GetCount())
}

// Safe counter
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count++
}

func (sc *SafeCounter) GetCount() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

// Deadlock example
func deadlockExample() {
	fmt.Println("Demonstrating deadlock prevention:")

	// Good: Use goroutines to prevent deadlock
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	value := <-ch
	fmt.Printf("Received: %d (no deadlock)\n", value)
}

// Goroutine leak example
func goroutineLeakExample() {
	fmt.Println("Demonstrating goroutine leak prevention:")

	// Good: Proper cleanup
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		defer close(ch)
		for {
			select {
			case ch <- 1:
				time.Sleep(100 * time.Millisecond)
			case <-done:
				fmt.Println("Goroutine cleanup")
				return
			}
		}
	}()

	// Do some work
	time.Sleep(200 * time.Millisecond)

	// Signal cleanup
	close(done)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Goroutine leak prevented")
}

// Safe alternatives example
func safeAlternativesExample() {
	fmt.Println("Safe concurrency patterns:")

	// Safe counter with atomic operations
	var atomicCounter int32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Use atomic operations for simple counters
			// atomic.AddInt32(&atomicCounter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("Atomic counter: %d\n", atomicCounter)

	// Safe channel usage
	safeChannelExample()
}

// Safe channel example
func safeChannelExample() {
	ch := make(chan int, 1) // Buffered channel

	// Safe send
	select {
	case ch <- 42:
		fmt.Println("Value sent successfully")
	default:
		fmt.Println("Channel full, value not sent")
	}

	// Safe receive
	select {
	case value := <-ch:
		fmt.Printf("Value received: %d\n", value)
	default:
		fmt.Println("No value available")
	}
}
