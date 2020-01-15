package main

import (
	"github.com/cybersamx/go-recipes/long-poll/pkg"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Handle system signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// Execute the server and client concurrently concurrently.
	go pkg.ListenForMessages()
	go pkg.SendMessages()

	sig := <-sigChan
	log.Printf("received signal %d, terminating...\n", sig)
	os.Exit(0)
}
