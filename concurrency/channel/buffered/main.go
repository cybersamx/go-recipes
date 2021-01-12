package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	stream := make(chan int, 2)
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			// Sleeps for 1 seconds before sending a message to the concurrency.
			stream <- i
			time.Sleep(1 * time.Second)
			fmt.Printf("Send %d to the channel after %v\n", i, time.Since(start))
		}
	}()

	go func() {
		defer wg.Done()
		for {
			// A receiving concurrency will block until there's a message in the receiving concurrency.
			time.Sleep(3 * time.Second)
			msg := <- stream
			fmt.Printf("Received %d from channel after %v\n", msg, time.Since(start))
		}
	}()

	wg.Wait()
	fmt.Println("Done")
}
