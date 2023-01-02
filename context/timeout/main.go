package main

import (
	"context"
	"log"
	"time"
)

func runTask(ctx context.Context, d time.Duration) {
	start := time.Now()

	// Timer to manage how long to simulate run the task.
	timer := time.NewTimer(d)
	defer timer.Stop()

	// This needs to be asynchronous so that we can receive the signals from the timer
	// and context Done() concurrently.
	go func() {
		// Simulate a long-running operation here.

		log.Printf("Task will run for %v\n", d)
	}()

	// We block and capture the first signal received. If we receive the timer signal,
	// the simulated run finishes before the context deadline. If we receive the Done()
	// signal, the simulated run ran over the deadline.
	select {
	case <-timer.C:
		log.Printf("Task ran for %v\n", time.Since(start))
		return
	case <-ctx.Done():
		log.Printf("Task ran longer than deadline or task has been canceled after %v\n", time.Since(start))
		log.Printf("Context error is %v\n", ctx.Err())
		return
	}
}

func run(d time.Duration) {
	const timeout = 2 * time.Second
	log.Printf("Task should not run more than %v\n", timeout)

	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer func() {
		log.Printf("Cancel called after %v\n", time.Since(start))
		cancel()
	}()

	runTask(ctx, d)
}

func main() {
	run(time.Second)

	log.Println()

	run(3 * time.Second)
}
