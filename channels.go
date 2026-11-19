// 3. Channels: Communication between goroutines
// Now comes one of the most important Go concepts.

// You already know how to start goroutines.

// But imagine this situation:

// Goroutine 1 calculates a number.

// Goroutine 2 needs that number.

// How will Goroutine 1 send the result to Goroutine 2?

// We use channels.

// A channel is a communication mechanism that allows goroutines to send and receive values.

// Think of it like a pipe....

package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		fmt.Println("Sending Values....")
		ch <- 100
		fmt.Println("Completed Sending...")

	}()

	time.Sleep(time.Second)

	fmt.Println("Waiting for Receving...")

	result := <-ch

	fmt.Printf("Received %d", result)
}
