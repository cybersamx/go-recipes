package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
