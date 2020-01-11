package main

import (
	"database/sql"
	"github.com/cybersamx/go-recipes/random/rand"
	_ "github.com/lib/pq"
	"time"
)

func insertDataSQL(n int) {
	// Setup for random generation.
	rand.Seed()
	now := time.Now()
	past := now.Add(-3 * 365 * 24 * time.Hour)

	// Connect to the database.
	db, err := sql.Open(dialect, getDSN())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		fatal("can't close the connection to the database", db.Close())
	}()

	// Get the last index.
	last := getLastUserID(db)

	// Randomly generate and write fake data to the database.
	for i := last + 1; i <= n + last; i++ {
		user := getUser(i)

		_, err := db.Exec("INSERT INTO users(id, name, age) VALUES($1, $2, $3)",
			user.ID, user.Name, user.Age)
		if err != nil {
			fatal("can't insert users", err)
		}

		for j := 0; j < 5; j++ {
			restaurant := getRestaurant(i, past, now)

			_, err = db.Exec("INSERT INTO restaurants(user_id, visited_at, name, num_seats, latitude, longitude) VALUES($1, $2, $3, $4, $5, $6)",
				restaurant.UserID,
				restaurant.VisitedAt,
				restaurant.Name,
				restaurant.NumSeats,
				restaurant.Latitude,
				restaurant.Longitude)
			if err != nil {
				fatal("can't insert restaurants", err)
			}
		}
	}
}
