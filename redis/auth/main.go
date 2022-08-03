package main

import (
	"log"
	"time"

	"github.com/mediocregopher/radix/v3"
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
	// Create a custom function of type ConnFunc or ClientFunc to pass it to NewPool to initialize
	// the connection to Redis.
	connHandler := func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(30*time.Second), // Additional configuration.
			radix.DialAuthPass("secrets"),     // Set password.
		)
	}

	client, err := radix.NewPool("tcp", "localhost:6379", 10, radix.PoolConnFunc(connHandler))
	checkErr(err)
	defer func() {
		err := client.Close()
		checkErr(err)
	}()

	// --- Test the connection ---
	var str string
	err = client.Do(radix.Cmd(&str, "PING"))
	checkErr(err)
	log.Println(str)

	// Start off by initializing the global counter to 0.
	err = client.Do(radix.FlatCmd(nil, "SET", counterKey, 7))
	checkErr(err)

	// Get the value.
	var counter int
	err = client.Do(radix.FlatCmd(&counter, "GET", counterKey))
	checkErr(err)
	log.Println(counter)
}
