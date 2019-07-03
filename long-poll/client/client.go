package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	ErrServerTimeout = errors.New("server timeout")
	ErrServerError   = errors.New("server error")
	ErrClientError   = errors.New("client error")
)

func callServer() (error, string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
		return ErrClientError, ""
	}
	req.Header.Add("Accept", "application/json; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return ErrClientError, ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return ErrClientError, ""
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return nil, string(body)
	case http.StatusNotModified:
		return ErrServerTimeout, ""
	default:
		return ErrServerError, ""
	}
}

func main() {
	for {
		callTime := time.Now()
		log.Print("REQUEST:  Calling server")
		err, result := callServer()

		diff := time.Now().Sub(callTime)
		if err == nil {
			log.Printf("RESPONSE: Received event after %.1fs: %s", diff.Seconds(), result)
		} else if err == ErrServerTimeout {
			log.Printf("RESPONSE: Server timeout after %.1fs", diff.Seconds())
		} else {
			log.Fatalf("PANIC: server error")
			break
		}
	}
}