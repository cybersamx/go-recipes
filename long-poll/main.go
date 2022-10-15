package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/cybersamx/go-recipes/long-poll/pkg"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	enableClient, _ := strconv.ParseBool(os.Getenv("ENABLE_CLIENT"))
	if enableClient {
		time.Sleep(2 * time.Second)
		go pkg.SendMessages()
	} else {
		go pkg.ListenForMessages()
	}

	sig := <-sigChan
	log.Printf("received signal %d, terminating...\n", sig)
	os.Exit(0)
}
