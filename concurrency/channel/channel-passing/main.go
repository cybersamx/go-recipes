package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 6; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
			fmt.Printf("Send %d\n", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// A range on the channel will read continuously from the channel until it is closed.
		for i := range ch {
			// A receiving channel will block until it receives a message.
			time.Sleep(2 * time.Second)
			fmt.Printf("Received %d\n", i)
		}
	}()

	wg.Wait()
	fmt.Println("Done")
}
