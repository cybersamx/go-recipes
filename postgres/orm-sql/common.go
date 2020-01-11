package main

import (
	"database/sql"
	"fmt"
	"github.com/cybersamx/go-recipes/random/rand"
	_ "github.com/lib/pq"
	"log"
	"syreclabs.com/go/faker"
	"time"
)

const (
	username = "pguser"
	password = "password"
	host     = "localhost"
	port     = 5432
	database = "db"
	dialect  = "postgres"
)

var (
	now = time.Now()
	past = now.Add(-3 * 365 * 24 * time.Hour)
)

func init() {
	log.Println("initializing a random seed...")
	rand.Seed()
}

// fatal is a simple way to dump all errors to log.Fatalf and exit from the program if an error occurs.
// Don't do this in production. This is strictly for the demo.
func fatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v\n", msg, err)
	}
}

func getDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		database)
}

func clearTables() {
	db, err := sql.Open(dialect, getDSN())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		fatal("can't close the connection to the database", db.Close())
	}()

	_, err = db.Exec("DELETE FROM restaurants")
	if err != nil {
		fatal("can't delete restaurants", err)
	}

	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		//fatal("can't rollback a transaction", tx.Rollback())
		fatal("can't delete users", err)
	}
}

func getFirstUserID(db *sql.DB) int {
	var count int
	if err := db.QueryRow("SELECT MIN(id) FROM users").Scan(&count); err != nil {
		return 0
	}
	return count
}

func getLastUserID(db *sql.DB) int {
	var count int
	if err := db.QueryRow("SELECT MAX(id) FROM users").Scan(&count); err != nil {
		return 0
	}
	return count
}

func getUser(id int) User {
	return User{
		ID: id,
		Name: faker.Name().Name(),
		Age: rand.RandomIntRange(14, 80),
	}
}

func getRestaurant(userID int, past, now time.Time) Restaurant {
	return Restaurant{
		UserID:    userID,
		VisitedAt: faker.Time().Between(past, now),
		Name:      faker.Company().Name(),
		NumSeats:  rand.RandomIntRange(8, 100),
		Latitude:  faker.Address().Latitude(),
		Longitude: faker.Address().Latitude(),
	}
}


