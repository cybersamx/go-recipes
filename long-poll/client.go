package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	errServerTimeout = errors.New("server timeout")
	errServerError   = errors.New("server error")
	errClientError   = errors.New("client error")
)

func callServer() (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8000", nil)
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

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return "", errClientError
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return string(body), nil
	case http.StatusNotModified:
		return "", errServerTimeout
	default:
		return "", errServerError
	}
}

func main() {
	for {
		callTime := time.Now()
		log.Print("request: calling server")
		result, err := callServer()

		diff := time.Now().Sub(callTime)
		if err == nil {
			log.Printf("respoonse: received event after %.1fs: %s", diff.Seconds(), result)
		} else if err == errServerTimeout {
			log.Printf("response: server timeout after %.1fs", diff.Seconds())
		} else {
			log.Fatalf("error: server error")
			break
		}
	}
}
