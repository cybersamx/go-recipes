package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, payload interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	jsonData, err := json.Marshal(payload)
	if err == nil {
		w.Write(jsonData)
	}

	return err
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "problem reading request body", http.StatusBadRequest)
		return
	}

	// Restore the body back to the original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	event := Event{}
	err = json.Unmarshal(body, &event)
	if err != nil {
		http.Error(w, "problem parsing the request json", http.StatusBadRequest)
		return
	}

	log.Printf("received event: %v", event.Event)

	resPayload := Message{Message: fmt.Sprintf("Received %s", event.Event)}
	if err := writeJSONResponse(w, resPayload); err != nil {
		http.Error(w, "error generating the response json", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	port := 8000
	log.Printf("starting web server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}
