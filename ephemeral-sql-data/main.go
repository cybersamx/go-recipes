package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

const (
	sqliteFile = "test.db"
	gcSeconds = 5
)

func main() {
	fmt.Printf("connect to the database.\n")
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		log.Fatalf("failed to open database: %v\n", err)
	}

	fmt.Printf("instantiate a store that garbage collects every %d seconds.\n", gcSeconds)
	config := Config{
		GCInterval: time.Duration(gcSeconds) * time.Second,
	}
	ts, err := NewStore(db, config)
	if err != nil {
		log.Fatalf("failed to create a store: %v\n", err)
		os.Exit(1)
	}
	time.Sleep(500 * time.Millisecond)

	expiresIn := 4
	fmt.Printf("create an ephemeral entity that expires in %d seconds.\n", expiresIn)
	id, err := ts.Create(time.Now().Add(time.Duration(expiresIn) * time.Second))
	if err != nil {
		log.Fatalf("failed to create a entity: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("sleeps past the expiration.\n")
	time.Sleep(8 * time.Second)
	entity, err := ts.Get(id)
	if err != nil {
		log.Fatalf("failed to retrieve the entity: %v\n", err)
		os.Exit(1)
	}

	if entity == nil {
		fmt.Printf("no entity was retrieved.\n")
	}
}
