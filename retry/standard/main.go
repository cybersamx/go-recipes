package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	pingTimeout  = 1 * time.Second
	maxRetries   = 4
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

		// Simple retry.
		i := 0
		for {
			log.Printf("Attempt %d: connecting to http server\n", i)
			client := http.DefaultClient
			res, err := client.Get("http://localhost:8000")
			if err == nil {
				log.Println("Connection success!")
				fmt.Println(res)
				break
			}
			i++
			if i > maxRetries {
				log.Printf("Exiting after %d retries\n", i)
				os.Exit(1)
			}

			log.Printf("Failed, retry in %d seconds\n", pingTimeout/time.Second)
			time.Sleep(pingTimeout)
		}
	}()

	wg.Wait()
}
