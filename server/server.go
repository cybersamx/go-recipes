package main

import (
	"context"
	"encoding/json"
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

	log.Printf("Received request and waiting %.1fs for event to emit", waitDuration.Seconds())
	select {
	case <- r.Context().Done():
		log.Printf("request canceled")
		w.WriteHeader(http.StatusNotModified)
	case <- timeoutCtx.Done():
		log.Printf("time out after %.1f", timeoutDuration.Seconds())
		w.WriteHeader(http.StatusNotModified)
	case <- waitChan:
		event := getRandomEvent()
		w.WriteHeader(http.StatusOK)
		payload := map[string]string{"event": event}
		writeJSONResponse(w, payload)
		timeoutCancel()
	}
}

func main() {
	err := http.ListenAndServe(":8000", http.HandlerFunc(rootHandler))
	log.Fatal(err)
}