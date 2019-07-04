package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var events = []string{"empty", "full", "new", "removed"}

func writeJSONResponse(w http.ResponseWriter, payload interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	jsonData, err := json.Marshal(payload)
	if err == nil {
		w.Write(jsonData)
	}

	return err
}

func getRandomEvent() string {
	index := rand.Intn(3)
	return events[index]
}

// An event will surface anytime between 1 to 10 seconds and is then sent to the client.
// If time for event to surface > 5 seconds, timeout and return a response with no content.

func rootHandler(w http.ResponseWriter, r *http.Request) {
	timeoutDuration := time.Duration(5) * time.Second
	waitDuration := time.Duration(rand.Intn(9) + 1) * time.Second
	waitChan := time.Tick(waitDuration) // Wait this long to pick an event
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), timeoutDuration)

	log.Printf("Received request and waiting %.1fs to emit event to the sender", waitDuration.Seconds())
	select {
	case <- r.Context().Done():
		log.Printf("request canceled")
		http.Error(w, "request canceled", http.StatusNotModified)
	case <- timeoutCtx.Done():
		log.Printf("time out after %.1f", timeoutDuration.Seconds())
		http.Error(w, fmt.Sprintf("time out after %.1f", timeoutDuration.Seconds()), http.StatusNotModified)
	case <- waitChan:
		event := getRandomEvent()
		w.WriteHeader(http.StatusOK)
		payload := map[string]string{"event": event}
		writeJSONResponse(w, payload)
		timeoutCancel()
	}
}

func main() {
	port := 8000
	log.Printf("Starting web server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(rootHandler))
	log.Fatal(err)
}