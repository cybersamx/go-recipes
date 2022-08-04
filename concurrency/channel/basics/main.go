package main

import (
	"fmt"
)

func main() {
	num := make(chan int, 5)

	// Channel capacity and length
	num <- 1

	// Read the channel and bring the cap bac to 0.
	// <-num returns a second value indicating whether the channel is closed.
	n, ok := <-num
	fmt.Println("When the channel isn't closed, ok =", ok)
	fmt.Println("Channel returns", n) // Closed channel will yield zero value

	// Write to the channel (within cap limit).
	for i := 0; i < 3; i++ {
		num <- i
	}

	// Read from the channel.
	for i := 0; i < 3; i++ {
		<-num
	}

	fmt.Printf("cap: %d, len: %d\n", cap(num), len(num))

	// Close a channel
	close(num)
	n, ok = <-num
	fmt.Println("When the channel is now closed, ok =", ok)
	fmt.Println("Channel returns", n) // Closed channel will yield zero value

	// Writing a value in closed channel will cause a panic.
	fmt.Println("Panic is expected.")
	num <- 100
}
