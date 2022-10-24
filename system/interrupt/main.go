package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		log.Println("Simulate a background task that prints out a timestamp every 1-15 seconds")
		for {
			seconds := rand.Intn(14) + 1
			log.Printf("Sleep for another %d seconds\n", seconds)
			time.Sleep(time.Duration(seconds) * time.Second)
		}
	}()

	sig := <-sigChan
	log.Printf("received signal %v, terminating...\n", sig)
}
