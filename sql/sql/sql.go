package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	username = "pguser"
	password = "password"
	host     = "localhost"
	port     = 5432
	database = "db"
	dialect  = "postgres"
)

type Location struct {
	ID        string
	UpdatedAt time.Time
	Latitude  float64
	Longitude float64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		database)
}

func fatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v\n", msg, err)
	}
}

func randRange(lower, upper int) float64 {
	return float64(lower) + rand.Float64()*float64(upper-lower)
}

func randLat() float64 {
	return randRange(-90, 90)
}

func randLong() float64 {
	return randRange(-180, 180)
}

func insertSQL(latitude, longitude float64) string {
	db, err := sql.Open(dialect, dsn())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			fatal("can't close the connection to the database", err)
		}
	}()

	id := uuid.New().String()
	_, err = db.Exec("INSERT INTO locations(id, updated_at, latitude, longitude) VALUES($1, $2, $3, $4)",
		id,
		time.Now(),
		latitude,
		longitude,
	)
	if err != nil {
		fatal("can't insert", err)
	}

	return id
}

func updateSQL(id string, latitude, longitude float64) {
	db, err := sql.Open(dialect, dsn())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			fatal("can't close the connection to the database", err)
		}
	}()

	_, err = db.Exec("UPDATE locations SET updated_at = $1, latitude = $2, longitude = $3 WHERE id = $4",
		time.Now(),
		latitude,
		longitude,
		id,
	)
	if err != nil {
		fatal("can't update", err)
	}
}

func selectSQL(id string) Location {
	// Connect to the database.
	db, err := sql.Open(dialect, dsn())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			fatal("can't close the connection to the database", err)
		}
	}()

	rows, err := db.Query("SELECT id, updated_at, latitude, longitude FROM locations WHERE id = $1", id)
	if err != nil {
		fatal("can't query from bus stops", err)
	}

	var loc Location
	for rows.Next() {
		fatal("can't scan from a row", rows.Scan(&loc.ID, &loc.UpdatedAt, &loc.Latitude, &loc.Longitude))
	}

	return loc
}

