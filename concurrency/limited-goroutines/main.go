package main

import (
	"log"
	"sync"
	"time"
)

const (
	workersCap = 3
)

func main() {
	// Need this to let the goroutines to complete execution before the program ends.
	var wg sync.WaitGroup

	// Simulate a queue for this program to consume its content concurrently. The
	// values are the simulated time (in seconds) to complete the processing.
	queue := []int{1, 3, 2, 4, 1, 1, 2, 3, 4, 2, 3, 1, 2, 2}

	// If channel is filled to the cap, the goroutine will be blocked until objects
	// are released from the channel.
	workersChan := make(chan struct{}, workersCap)

	for i, n := range queue {
		wg.Add(1)
		go func(i, n int) {
			defer wg.Done()

			// Fill the channel with an object.
			workersChan <- struct{}{}

			duration := time.Duration(n) * time.Second
			log.Printf("Started processing index %d for %s", i, duration.String())
			time.Sleep(duration)
			log.Printf("Finished processing index %d", i)

			// Consume the channel to release the object, subsequently allowing the
			// next queued object to be passed to the channel. This resume the paused
			// goroutine continue its execution.
			<-workersChan
		}(i, n)
	}

	wg.Wait()
}
