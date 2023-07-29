package main

import (
	"context"
	"fmt"
	"time"
)

const (
	ctxTimeout   = 3 * time.Second
	taskATimeout = 2 * time.Second
	taskBTimeout = 5 * time.Second
)

// runTask runs a "task" with 2 timers, one is timer.After() based and the other is context.WithTimeout based.
// The function waits and listens for the signal in the select statement. Whichever timer finishes will complete
// the function.
func runTask(ctx context.Context, task string, d time.Duration) {
	now := time.Now()
	fmt.Println("Current time:   ", now)

	select {
	case <-time.After(d):
		fmt.Println(task, "done at:   ", time.Now())
	case <-ctx.Done():
		fmt.Println(task, "timeout at:", time.Now())
	}
}

func main() {
	fmt.Println("Context timeout in", ctxTimeout)

	// Wait for the task A to finish.
	ctxA, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	fmt.Println("Task A to run for ", taskATimeout, " - this task will complete as the task time < context timeout")
	go runTask(ctxA, "Task A", taskATimeout)

	// Wait for the task B to finish.
	ctxB, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	fmt.Println("Task B to run for ", taskBTimeout, " - this task will not complete as the task time > context timeout")
	go runTask(ctxB, "Task B", taskBTimeout)

	<-ctxA.Done()
	<-ctxB.Done()
}
