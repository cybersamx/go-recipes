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
	x = 2 * time.Second
	y = 5 * time.Second
)

func main() {
	fmt.Println(
		`This program will run a task for x seconds with timeout for y.
If x > y, the task will be canceled via context cancel.
If x < y, the task will complete.
Change x and y in the code.`)

	future := time.Now().Add(x)
	ctx, cancel := context.WithDeadline(context.Background(), future)
	// Call cancel when we go out of scope, or we'll have a context leak.
	defer cancel()

	ch := make(chan result, 1)

	// Perform the task asynchronously.
	go func() {
		// Simulate work.
		// NOTE: Change the time to exceed the x and see the result.
		time.Sleep(y * time.Second)

		ch <- result{value: time.Now().String()}
	}()

	// Wait for the task to finish.
	select {
	case res := <-ch:
		fmt.Printf("task done: %s\n", res.value)
	case <-ctx.Done():
		fmt.Println("task canceled or timeout")
	}
}
