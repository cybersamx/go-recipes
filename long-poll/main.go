package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	enableClient, _ := strconv.ParseBool(os.Getenv("ENABLE_CLIENT"))
	if enableClient {
		time.Sleep(2 * time.Second)
		go SendMessages()
	} else {
		go ListenForMessages()
	}

	sig := <-sigChan
	log.Printf("received signal %v, terminating...\n", sig)
}
