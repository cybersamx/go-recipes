package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	errServerTimeout = errors.New("server timeout")
	errServerError   = errors.New("server error")
	errClientError   = errors.New("client error")
)

func callServer() (string, error) {
	client := http.Client{}
	url := fmt.Sprintf("http://localhost:%d", port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return "", errClientError
	}
	req.Header.Add("Accept", "application/json; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", errClientError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", errClientError
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return string(body), nil
	case http.StatusNotModified:
		return "", errServerTimeout
	default:
		return "", errServerError
	}
}

func SendMessages() {
	for {
		callTime := time.Now()
		log.Print("client calling server to receive an event")
		result, err := callServer()

		diff := time.Now().Sub(callTime)
		if err == nil {
			log.Printf("server received event after %.1fs: %s", diff.Seconds(), result)
		} else if err == errServerTimeout {
			log.Printf("server received server timeout after %.1fs", diff.Seconds())
		} else {
			log.Fatal("server received server error")
		}
	}
}
