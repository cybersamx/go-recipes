package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func writeJSONResponse(w http.ResponseWriter, payload interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	jsonData, err := json.Marshal(payload)
	if err == nil {
		w.Write(jsonData)
	}

	return err
}

func getRandomEvent() string {
	events := []string{"empty", "full", "new", "removed"}
	index := rand.Intn(3)
	return events[index]
}

// An event will surface anytime between 1 to 10 seconds and is then sent to the client.
// If time for event to surface > 5 seconds, timeout and return a response with no content.

func RootHandler(w http.ResponseWriter, r *http.Request) {
	wait := rand.Intn(9) + 1
	waitChan := time.Tick(time.Duration(wait) * time.Second) // Wait this long to pick an event
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), time.Duration(5) * time.Second)

	log.Printf("Received request and waiting %d seconds to event to emit", wait)
	select {
	case <- r.Context().Done():
		log.Println("request canceled")
		w.WriteHeader(http.StatusNotModified)
	case <- timeoutCtx.Done():
		log.Println("time out")
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
	err := http.ListenAndServe(":8000", http.HandlerFunc(RootHandler))
	log.Fatal(err)
}