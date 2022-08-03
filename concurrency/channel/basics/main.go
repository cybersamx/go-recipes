package main

import (
	"fmt"
)

func main() {
	num := make(chan int, 5)

	go func() {
		num <- 99
	}()

	// The receiving channel will block until it receives a message.
	msg := <-num
	fmt.Println(msg)

	// Channel capacity and length
	num <- 1
	fmt.Printf("cap: %d, len: %d\n", cap(num), len(num))
	<-num // Channel queue (len) is back to 0

	num <- 1
	num <- 2
	num <- 3
	fmt.Printf("cap: %d, len: %d\n", cap(num), len(num))

	// Close a channel
	close(num)
	fmt.Println(<-num) // Closed channel will yield zero value

	// Writing a value in closed channel will cause a panic.
	fmt.Println("Panic is expected.")
	num <- 100
}
