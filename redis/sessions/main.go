package main

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/mediocregopher/radix/v3"

	"github.com/cybersamx/go-recipes/redis/sessions/rdb"
)

const (
	expiry = 2 // Seconds
	sleep  = 3 // Seconds
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runRadix(session *rdb.Session) {
	log.Println("Running using the radix driver...")

	client, err := radix.NewPool("tcp", "localhost:6379", 10)
	checkErr(err)
	defer func() {
		err := client.Close()
		checkErr(err)
	}()

	log.Println("session created: ", session)

	log.Printf("Set sessions to expire in %d seconds", expiry)
	checkErr(rdb.RadixSetSession(client, session, expiry))

	// Get the object.
	var foundSession rdb.Session
	checkErr(rdb.RadixGetSession(client, &foundSession, session.SessionID))
	log.Println("Found session in redis:", foundSession)

	// Wait.
	log.Printf("Wait %d seconds for session to expire", sleep)
	time.Sleep(sleep * time.Second)

	// Get the object after timeout.
	foundSession = rdb.Session{}
	err = rdb.RadixGetSession(client, &foundSession, session.SessionID)
	if err == rdb.NotFoundErr {
		// Expected.
		log.Println("Success: the session expired in redis, we didn't fetch the object from redis")
	} else if err != nil {
		log.Fatalf("Unexpected failure: %v", err)
	} else {
		log.Fatal("Failure: session should have expired from redis")
	}

	log.Println()
}

func runRedis(session *rdb.Session) {
	log.Println("Running using the go-redis driver...")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer func() {
		err := client.Close()
		checkErr(err)
	}()

	log.Println("session created: ", session)

	log.Printf("Set sessions to expire in %d seconds", expiry)
	checkErr(rdb.RedisSetSession(client, session, expiry))

	// Get the object.
	var foundSession rdb.Session
	checkErr(rdb.RedisGetSession(client, &foundSession, session.SessionID))
	log.Println("Found session in redis:", foundSession)

	// Wait.
	log.Printf("Wait %d seconds for session to expire", sleep)
	time.Sleep(sleep * time.Second)

	// Get the object after timeout.
	foundSession = rdb.Session{}
	err := rdb.RedisGetSession(client, &foundSession, session.SessionID)
	if err == rdb.NotFoundErr {
		// Expected.
		log.Println("Success: the session expired in redis, we didn't fetch the object from redis")
	} else if err != nil {
		log.Fatalf("Unexpected failure: %v", err)
	} else {
		log.Fatal("Failure: session should have expired from redis")
	}

	log.Println()
}

func main() {
	// If we want to set an object to redis, we would need to serialized to bytes.
	// And deserialized the bytes back to an object after we fetch it from redis.
	session := rdb.Session{
		SessionID: uuid.New().String(),
		UserID:    12345,
		Username:  "john2000",
	}

	runRadix(&session)
	runRedis(&session)
}
