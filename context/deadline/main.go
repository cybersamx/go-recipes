package main

import (
	"context"
	"fmt"
	"time"
)

type result struct {
	value string
}

const (
	ctxTimeout  = 2 * time.Second
	workTimeout = 5 * time.Second
)

func main() {
	future := time.Now().Add(ctxTimeout)
	ctx, cancel := context.WithDeadline(context.Background(), future)
	// Call cancel when done or we'll have a context leak.
	defer cancel()

	ch := make(chan result, 1)

	// Perform the task asynchronously.
	go func() {
		// Simulate work.
		// NOTE: Change the time to exceed the ctxTimeout and see the result.
		time.Sleep(workTimeout * time.Second)

		ch <- result{value: time.Now().String()}
	}()

	// Wait for the task to finish.
	select {
	case res := <-ch:
		fmt.Printf("task done: %s\n", res.value)
	case <-ctx.Done():
		fmt.Println("task canceled/ctxTimeout")
	}
}
