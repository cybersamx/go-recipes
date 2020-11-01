package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mediocregopher/radix/v3"
	"github.com/mediocregopher/radix/v3/resp"
)

const (
	expiry = 2  // Seconds
	sleep  = 3  // Seconds
)

// Session represents a session of an application.
type Session struct {
	SessionID  string
	UserID     uint
	Username   string
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func gobEncode(val interface{}) (*bytes.Buffer, error) {
	// bytes.Buffer implements io.Reader and io.Writer.
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}

	return &buf, nil
}

func gobDecode(buf *bytes.Buffer, val interface{}) error {
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(val); err != nil {
		return err
	}

	return nil
}

func main() {
	client, err := radix.NewPool("tcp", "localhost:6379", 10)
	checkErr(err)
	defer func() {
		err := client.Close()
		checkErr(err)
	}()

	// If you want to set an object to redis, you would need to serialized to a flat string.
	// And deserialized the string back to an object when you retrieve it from redis.
	session := Session{
		SessionID: uuid.New().String(),
		UserID: 12345,
		Username: "john2000",
	}

	log.Println("session created: ", session)

	// Set the object in sessions
	buffer, err := gobEncode(&session)
	checkErr(err)

	// See Radix documentation for more info NewLenReader and FlatCmd.
	// https://godoc.org/github.com/mediocregopher/radix#FlatCmd
	reader := resp.NewLenReader(buffer, int64(buffer.Len()))
	err = client.Do(radix.FlatCmd(nil, "SET", session.SessionID, reader))
	checkErr(err)

	// Set timeout for the sessions (of type hash).
	log.Printf("Set sessions to expire in %d seconds", expiry)
	err = client.Do(radix.FlatCmd(nil, "EXPIRE", session.SessionID, expiry))
	checkErr(err)

	// Get the object.
	err = client.Do(radix.FlatCmd(buffer, "GET", session.SessionID))
	checkErr(err)
	var foundSession Session
	err = gobDecode(buffer, &foundSession)
	checkErr(err)

	log.Println("session in redis:", foundSession)

	log.Printf("Wait %d seconds for session to expire", sleep)
	time.Sleep(sleep * time.Second)

	// Get the object after timeout.
	var tbuffer bytes.Buffer
	mn := radix.MaybeNil{
		Rcv: &tbuffer,
	}
	err = client.Do(radix.FlatCmd(&mn, "GET",  session.SessionID))
	if err != nil {
		log.Fatalf("session should have expired but we are able to fetch object %v from redis", session)
	} else if mn.Nil {
		log.Println("success: the session expired in redis, we didn't fetch the object from redis")
		return
	}
}
