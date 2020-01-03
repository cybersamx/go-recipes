package main

import (
	"context"
	"fmt"
	"time"
)

type result struct {
	value string
}

const timeout = 2 * time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// Call cancel when done or we'll have a context leak.
	defer cancel()

	ch := make(chan result, 1)

	// Perform the task asynchronously.
	go func() {
		// Simulate work.
		// NOTE: Change the time to exceed the timeout and see the result.
		time.Sleep(1 * time.Second)

		ch <- result{value: time.Now().String()}
	}()

	// Wait for the task to finish.
	select {
	case res := <-ch:
		fmt.Printf("task done: %s\n", res.value)
	case <-ctx.Done():
		fmt.Printf("task canceled/timeout")
	}
}
