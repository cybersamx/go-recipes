package main

import (
	"database/sql"
	"fmt"
	"github.com/cybersamx/go-recipes/random/rand"
	_ "github.com/lib/pq"
	"log"
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

func init() {
	log.Println("initializing a random seed...")
	rand.Seed()
}

// fatal is a timeout way to dump all errors to log.Fatalf and exit from the program if an error occurs.
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

	_, err = db.Exec("DELETE FROM bus_stops")
	if err != nil {
		fatal("can't delete bus stops", err)
	}
}

func getFirstBusStopID(db *sql.DB) int {
	var count int
	if err := db.QueryRow("SELECT MIN(id) FROM bus_stops").Scan(&count); err != nil {
		return 0
	}
	return count
}

func getLastBusStopID(db *sql.DB) int {
	var count int
	if err := db.QueryRow("SELECT MAX(id) FROM bus_stops").Scan(&count); err != nil {
		return 0
	}
	return count
}

func getBusStop(id int) BusStop {
	date := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
	return BusStop{
		ID:        id,
		UpdatedAt: date,
		Number:    "LA-0664",
		Latitude:  34.19462,
		Longitude: -118.58843,
		SiteATS:   "NB  De Soto  FS  Vanowen -NE",
		CitySite:  "L.A. Valley West",
	}
}

func getBusStopForUpdate(id int) BusStop {
	date := time.Date(2001, 6, 6, 6, 0, 0, 0, time.UTC)
	return BusStop{
		ID:        id,
		UpdatedAt: date,
		Number:    "LA-0750",
		Latitude:  34.15238,
		Longitude: -118.60487,
		SiteATS:   "EB  Mulholland D  NS  Topanga -SW",
		CitySite:  "L.A. Valley West",
	}
}
