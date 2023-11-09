package main

import (
	"fmt"
	"sync"
	"time"
)

// simpleGoroutine starts a goroutine that prints a number.
func simpleGoroutine(number int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this goroutine is done
	fmt.Printf("Number: %d\n", number)
}

func main() {
	var wg sync.WaitGroup

	// Start multiple goroutines.
	for i := 0; i < 10; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go simpleGoroutine(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("All goroutines completed.")
}

// This function demonstrates sending and receiving data through a channel.
func channelDemo() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Data from goroutine" // Send data to the channel
	}()

	data := <-ch // Receive data from the channel
	fmt.Println(data)
}

// Call channelDemo to see channels in action
func init() {
	go channelDemo()
}

