package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/cybersamx/go-recipes/webhook/utils"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

var (
	ErrClientError   = errors.New("client error")
)

const (
	defaultTimeout = 5 * time.Second
)

func getRandomEvent() string {
	events := []string{"create", "update", "remove", "clear"}
	index := rand.Intn(3)
	return events[index]
}

func callServer(ticker *time.Ticker) {
	for range ticker.C {
		event := getRandomEvent()

		log.Printf("sending event %s to the webhook\n", event)

		client := http.Client{}
		req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/webhook", nil)
		if err != nil {
			log.Fatal(err)
			continue
		}

		req.Header.Add("Accept", "application/json; charset=UTF-8")
		reqPayload := utils.Event{Event: getRandomEvent()}
		jsonData, err := json.Marshal(reqPayload)
		if err != nil {
			log.Fatal(err)
			continue
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))

		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			log.Fatal(err)
			continue
		}

		log.Printf("Response: %s", string(body))
	}
}

func main() {
	ticker := time.NewTicker(defaultTimeout)

	go callServer(ticker)

	// Terminates the main goroutine without main() returning.
	// This effectively runs the main() forever.
	runtime.Goexit()
}