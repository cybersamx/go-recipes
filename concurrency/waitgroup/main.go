package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Without WaitGroup, main doesn't wait for goroutines to exit and the goroutines
	// won't complete their execution to print the output.

	var wg sync.WaitGroup

	// Add 1 to tell wg that a goroutine is starting.
	wg.Add(1)
	go func() {
		// Done tells wg that the goroutine completes.
		defer wg.Done()
		fmt.Println("GoRoutine 1 Start - Process for 5 seconds")
		time.Sleep(5 * time.Second)
		fmt.Println("GoRoutine 1 End")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("GoRoutine 2 Start - Process for 3 seconds")
		time.Sleep(3 * time.Second)
		fmt.Println("GoRoutine 2 End")
	}()

	// Block the main routine until all goroutines completes.
	wg.Wait()

	fmt.Println("All goroutines completed...")
}
