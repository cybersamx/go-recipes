package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	port = 8000
)

var events = []string{"empty", "full", "new", "removed"}

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func writeJSONResponse(w http.ResponseWriter, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(payload)
	if err != nil {
		return err
	}

	return err
}

func getRandomEvent() string {
	index := rand.Intn(3)
	return events[index]
}

func respondWithEvent(w http.ResponseWriter) {
	event := getRandomEvent()
	w.WriteHeader(http.StatusOK)
	payload := map[string]string{"event": event}
	if err := writeJSONResponse(w, payload); err != nil {
		log.Printf("server error: %v", err)
	}
}

// An event will surface anytime between 1 and 10 seconds and is then sent to the client.
// If time for event to surface > 5 seconds, timeout and return a response with no content.

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	timeout := time.Duration(5) * time.Second
	wait := time.Duration(rand.Intn(9)+1) * time.Second
	if wait < 0 {
		log.Print("server received request and return an event back to the sender immediately")
		respondWithEvent(w)
		return
	}
	waitChan := time.Tick(wait) // Wait this long to pick an event
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Printf("server received request and waits %.1fs to emit an event back to the sender", wait.Seconds())
	select {
	case <-r.Context().Done():
		log.Print("server request canceled")
		http.Error(w, "request canceled", http.StatusNotModified)
	case <-ctx.Done():
		log.Printf("server timed out after %.1f", timeout.Seconds())
		http.Error(w, fmt.Sprintf("server timed out after %.1f", timeout.Seconds()), http.StatusNotModified)
	case <-waitChan:
		respondWithEvent(w)
	}
}

func ListenForMessages() {
	log.Printf("server starting web server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(rootHandler))
	log.Fatal(err)
}
