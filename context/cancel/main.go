package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	timer := time.NewTimer(5 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("timer started")
		<-timer.C
		fmt.Println("time is up, canceling the context")
		cancel()
	}()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(ctx context.Context, n int) {
			defer wg.Done()

			fmt.Printf("spawned gorountine #%d\n", n)

			// Blocks until a signal is received in the channel. Calling the cancel
			// function causes the cancellation signal to be propagated through Done()
			// channel.
			<-ctx.Done()
			fmt.Printf("goroutine #%d received done signal\n", n)
		}(ctx, i)
	}

	wg.Wait()
}
