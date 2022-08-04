package main

import (
	"fmt"
)

func main() {
	// --- Buffered ---

	buf := make(chan int, 3)

	go func() {
		buf <- 1
		buf <- 2
		buf <- 3
		fmt.Printf("cap: %d, len: %d\n", cap(buf), len(buf))
	}()

	// 2 separate goroutines are reading and writing to the channel.
	fmt.Println(<-buf)
	fmt.Println(<-buf)
	fmt.Println(<-buf)

	// Buffered channel capacity and length
	fmt.Printf("cap: %d, len: %d\n", cap(buf), len(buf))

	// Close a channel
	close(buf)

	// --- Unbuffered ---

	unbuf := make(chan int)

	go func() {
		fmt.Printf("cap: %d, len: %d\n", cap(unbuf), len(unbuf))
		unbuf <- 100
	}()

	// This works because we have a concurrent goroutine writing to the channel and
	// a reader that will eventually read from that channel.
	fmt.Println(<-unbuf)

	// The following won't work because the unbuf channel will never get read so the
	// main routine will be blocked forever after unbuf <- 101 leading to a deadlock.
	fmt.Println("Deadlock error is expected.")
	unbuf <- 101
}
