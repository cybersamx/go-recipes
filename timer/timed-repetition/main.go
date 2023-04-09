package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func randDelay() time.Duration {
	return time.Duration(rand.Intn(9)+1) * time.Second
}

func repeatPrint(ctx context.Context) {
	go func() {
		n := randDelay()
		log.Printf("hello and sleep for %v", n)

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(n):
				n = randDelay()
				log.Printf("hello and sleep for %v", n)
			}
		}
	}()
}

func main() {
	sigChan := make(chan os.Signal)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	repeatPrint(ctx)

	<-sigChan
}
