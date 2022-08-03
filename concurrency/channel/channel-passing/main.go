package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	stream := make(chan int, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(ch chan<- int) {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
			fmt.Printf("Send %d\n", i)
		}
		close(ch)
	}(stream)

	wg.Add(1)
	go func(ch <-chan int) {
		defer wg.Done()
		// A range on the channel will read continuously from the channel until it is closed.
		for i := range ch {
			// A receiving channel will block until it receives a message.
			time.Sleep(2 * time.Second)
			fmt.Printf("Received %d\n", i)
		}
	}(stream)

	wg.Wait()
	fmt.Println("Done")
}
