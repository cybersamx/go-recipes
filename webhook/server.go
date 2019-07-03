package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cybersamx/go-recipes/webhook/utils"
	"io/ioutil"
	"log"
	"net/http"
)

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
	event := utils.Event{}
	err = json.Unmarshal(body, &event)
	if err != nil {
		http.Error(w, "problem parsing the request json", http.StatusBadRequest)
		return
	}

	log.Printf("Received event: %v", event.Event)

	resPayload := utils.Messasge{Message: fmt.Sprintf("Received %s", event.Event)}
	if err := utils.WriteJSONResponse(w, resPayload); err != nil {
		http.Error(w, "error generating the response json", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	port := 8000
	log.Printf("Starting web server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}