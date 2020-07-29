package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Message struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Text      string    `json:text`
	CreatedAt time.Time `json:created_at`
}

func printNode(node interface{}) {
	typ := reflect.TypeOf(node)
	if typ.Kind() == reflect.Map {
		if m, ok := node.(map[string]interface{}); ok {
			for k, v := range m {
				vTyp := reflect.TypeOf(v)
				if vTyp.Kind() == reflect.Map {
					printNode(v)
				} else {
					fmt.Println(k, ":", v)
				}
			}
		}
	}
}

func get() {
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("closing response from get()...")
		fmt.Println()
		if err := res.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Print the response body as string to stdout.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func post() {
	msg := Message{
		From:      "cybersamx",
		To:        "golang",
		Text:      "Go rocks!",
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("closing response from post()...")
		fmt.Println()
		if err := res.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Parse the response body into a semi-structure
	var obj map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		log.Fatal(err)
	}

	// Flat print the response body.
	printNode(obj)
}

func main() {
	get()
	post()
}
