package main

import (
	"fmt"
	"time"
)

func main() {
	stream := make(chan string, 2)

	go func() {
		// Sleeps for 3 seconds before sending a message to the concurrency.
		fmt.Println("Start goroutine - Sleep for 3 seconds")
		time.Sleep(3 * time.Second)
		fmt.Println("Send a message to the concurrency")
		stream <- "Hello"
	}()

	// A receiving concurrency will block until there's a message in the receiving concurrency.
	msg := <- stream
	fmt.Println("Receive a message from the concurrency: ", msg)
}
