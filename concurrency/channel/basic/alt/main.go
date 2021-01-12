package main

import (
	"fmt"
	"time"
)

func main() {
	// Creates an buffered channel with a capacity of 2. This enables asynchronous channel
	// communication.
	stream := make(chan string, 2)
	start := time.Now()

	go func() {
		// Sleeps for 3 seconds before sending a message to the concurrency.
		msg := "Hello"
		fmt.Println("Start goroutine - Sleep for 3 seconds")
		time.Sleep(3 * time.Second)
		fmt.Printf("Send a message to the channel after %v: %s\n", time.Since(start), msg)
		stream <- msg
	}()

	// A receiving concurrency will block until there's a message in the receiving concurrency.
	// Here, we use range instead of <- to receive a message from a channel.
	for msg := range stream {
		fmt.Printf("Received a message from the concurrency after %v: %s\n", time.Since(start), msg)
	}
}
