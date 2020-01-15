package pkg

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const cPrefix = "\033[31;1;4mclient\033[0m"

var (
	errServerTimeout = errors.New("server timeout")
	errServerError   = errors.New("server error")
	errClientError   = errors.New("client error")
)

func callServer() (retEvent string, retErr error) {
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
	if err != nil {
		log.Fatal(err)
		return "", errClientError
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			retEvent = ""
			retErr = err
		}
	}()


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
		log.Printf("%s calling server to receive an event", cPrefix)
		result, err := callServer()

		diff := time.Now().Sub(callTime)
		if err == nil {
			log.Printf("%s received event after %.1fs: %s", cPrefix, diff.Seconds(), result)
		} else if err == errServerTimeout {
			log.Printf("%s received server timeout after %.1fs", cPrefix, diff.Seconds())
		} else {
			log.Fatalf("%s received server error", cPrefix)
		}
	}
}
