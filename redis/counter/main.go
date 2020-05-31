package main

import (
	"github.com/mediocregopher/radix/v3"
	"log"
)

const (
	counterKey = "counter"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client, err := radix.NewPool("tcp", "localhost:6379", 10)
	checkErr(err)
	defer func() {
		err := client.Close()
		checkErr(err)
	}()

	// --- Test the connection ---
	var str string
	err = client.Do(radix.Cmd(&str, "PING"))
	checkErr(err)

	// The response should be simply "PONG".
	log.Printf("connected to redis? %t", str == "PONG")

	// --- Global counter using Cmd ---

	// Start off by initializing the global counter to 0.
	err = client.Do(radix.Cmd(nil, "SET", counterKey, "0"))
	checkErr(err)

	// Get the value.
	err = client.Do(radix.Cmd(&str, "GET", counterKey))
	checkErr(err)
	log.Println(str)

	// We increment the counter by 1.
	err = client.Do(radix.Cmd(nil, "INCR", counterKey))
	checkErr(err)

	// Get the value again, it should be incremented by 1.
	err = client.Do(radix.Cmd(&str, "GET", counterKey))
	checkErr(err)
	log.Println(str)

	// --- Global counter using FlatCmd ---

	// Previously, we are passing the param as string because we are using the Cmd function
	// which only accepts string types. The following functions accept int values and
	// automatically flatten the the values via the FlatCmd function.

	// Start off by initializing the global counter to 0.
	err = client.Do(radix.FlatCmd(nil, "SET", counterKey, 0))
	checkErr(err)

	// Get the value.
	var counter int
	err = client.Do(radix.FlatCmd(&counter, "GET", counterKey))
	checkErr(err)
	log.Println(counter)

	// We increment the counter by 1.
	err = client.Do(radix.FlatCmd(nil, "INCR", counterKey))
	checkErr(err)

	// Get the value again, it should be incremented by 1.
	err = client.Do(radix.FlatCmd(&counter, "GET", counterKey))
	checkErr(err)
	log.Println(counter)
}
