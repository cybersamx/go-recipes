package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/flowchartsman/retry"
)

const (
	maxRetries   = 3
	initialDelay = 1 * time.Second
	maxDelay     = 6 * time.Second
	httpMaxDelay = 5
)

func main() {
	rand.Seed(time.Now().Unix())

	wg := sync.WaitGroup{}
	wg.Add(2)

	srv := http.Server{
		Addr: ":8000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello")
		}),
	}

	go func() {
		defer wg.Done()

		// Simulate delay.
		random := time.Duration(rand.Int()%httpMaxDelay) + 2
		log.Printf("Launcing http server in %d seconds\n", random)
		time.Sleep(random * time.Second)

		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server failed to launch: %v\n", err)
		}
	}()

	go func() {
		defer wg.Done()
		defer srv.Shutdown(context.Background())

		retrier := retry.NewRetrier(maxRetries, initialDelay, maxDelay)
		i := 0
		err := retrier.Run(func() error {
			log.Printf("Attempt %d: connecting to http server\n", i)
			i++
			client := http.DefaultClient
			res, err := client.Get("http://localhost:8000")
			if err != nil {
				return err // Triggers another retry
			}

			log.Println("Connection success!")
			fmt.Println(res)
			return nil
		})

		if err != nil {
			log.Printf("Exiting after %d retries\n", i)
		}
	}()

	wg.Wait()
}
