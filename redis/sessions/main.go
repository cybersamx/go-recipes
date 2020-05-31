 package main

 import (
	 "bytes"
	 "encoding/gob"
	 "fmt"
	 "github.com/google/uuid"
	 "github.com/mediocregopher/radix/v3"
	 "github.com/mediocregopher/radix/v3/resp"
	 "log"
	 "time"
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

func getKey(session *Session) string {
	return fmt.Sprintf("sessions.%s", session.SessionID)
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

	log.Println("session:              ", session)

	// Set the object in sessions
	var buffer bytes.Buffer // bytes.Buffer implements io.Reader and io.Writer.
	enc := gob.NewEncoder(&buffer)
	err = enc.Encode(session)
	checkErr(err)

	// See Radix documentation for more info NewLenReader and FlatCmd.
	// https://godoc.org/github.com/mediocregopher/radix#FlatCmd
	reader := resp.NewLenReader(&buffer, int64(buffer.Len()))
	err = client.Do(radix.FlatCmd(nil, "HSET", "sessions", session.SessionID, reader))
	checkErr(err)

	// Set timeout for the sessions (of type hash).
	err = client.Do(radix.FlatCmd(nil, "EXPIRE", "sessions", 2))
	checkErr(err)

	// Get the object.
	err = client.Do(radix.FlatCmd(&buffer, "HGET", "sessions", session.SessionID))
	checkErr(err)
	var foundSession Session
	dec := gob.NewDecoder(&buffer)
	err = dec.Decode(&foundSession)
	checkErr(err)

	log.Println("session in redis:     ", foundSession)

	time.Sleep(3 * time.Second)

	// Get the object after timeout.
	var timeoutSession Session
	var tbuffer bytes.Buffer
	mn := radix.MaybeNil{
		Rcv: &tbuffer,
	}
	err = client.Do(radix.FlatCmd(&mn, "HGET", "sessions", session.SessionID))
	if err != nil {
		log.Fatal(err)
	} else if mn.Nil {
		log.Println("session after timeout: nil")
		return
	}

	dec2 := gob.NewDecoder(&tbuffer)
	err = dec2.Decode(&timeoutSession)
	checkErr(err)
	log.Println("session after timeout:", timeoutSession)
}
