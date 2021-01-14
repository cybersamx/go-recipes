package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func insertDataSQL(n int) {
	// Connect to the database.
	db, err := sql.Open(dialect, getDSN())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		fatal("can't close the connection to the database", db.Close())
	}()

	// Get the last index.
	last := getLastBusStopID(db)

	// Randomly generate and write fake data to the database.
	for i := last + 1; i <= n+last; i++ {
		busStop := getBusStop(i)
		_, err := db.Exec("INSERT INTO bus_stops(id, updated_at, number, latitude, longitude, siteats, city_site) VALUES($1, $2, $3, $4, $5, $6, $7)",
			busStop.ID, busStop.UpdatedAt, busStop.Number, busStop.Latitude, busStop.Longitude, busStop.SiteATS, busStop.CitySite)
		if err != nil {
			fatal("can't insert bus stops", err)
		}
	}
}

func updateDataSQL(n int) {
	// Connect to the database.
	db, err := sql.Open(dialect, getDSN())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		fatal("can't close the connection to the database", db.Close())
	}()

	// Get the first index.
	first := getFirstBusStopID(db)

	// Randomly generate and write fake data to the database.
	for i := first; i < n; i++ {
		busStop := getBusStopForUpdate(i)
		_, err := db.Exec("UPDATE bus_stops SET updated_at = $1, number = $2, latitude = $3, longitude = $4, siteats = $5, city_site = $6 WHERE id = $7",
			busStop.UpdatedAt,
			busStop.Number,
			busStop.Latitude,
			busStop.Longitude,
			busStop.SiteATS,
			busStop.CitySite,
			i)
		if err != nil {
			fatal("can't update bus stops", err)
		}
	}
}

func selectDataSQL(n int) {
	// Connect to the database.
	db, err := sql.Open(dialect, getDSN())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		fatal("can't close the connection to the database", db.Close())
	}()

	// Get the first index.
	first := getFirstBusStopID(db)

	// Randomly generate and write fake data to the database.
	for i := first; i < n; i++ {
		rows, err := db.Query("SELECT id, updated_at, number, latitude, longitude, siteats, city_site FROM bus_stops WHERE id = $1", i)
		if err != nil {
			fatal("can't query from bus stops", err)
		}

		var busStop BusStop
		for rows.Next() {
			fatal("can't scan from a row", rows.Scan(&busStop.ID, &busStop.UpdatedAt, &busStop.Number, &busStop.Latitude, &busStop.Longitude, &busStop.SiteATS, &busStop.CitySite))
		}

		fatal("can't close row", rows.Close())
	}
}
