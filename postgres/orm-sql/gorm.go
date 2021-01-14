package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getGORMEngine() (*gorm.DB, error) {
	db, err := gorm.Open(dialect, getDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func insertDataGORM(n int) {
	// Connect to the database.
	db, err := getGORMEngine()
	if err != nil {
		fatal("can't initialize xorm db", err)
	}

	defer func() {
		fatal("can't close connection to database", db.Close())
	}()

	// Get the last index.
	last := getLastBusStopID(db.DB())

	// Randomly generate and write data to the database.
	for i := last + 1; i <= n+last; i++ {
		busStop := getBusStop(i)
		fatal("can't insert bus stop", db.Create(&busStop).Error)
	}
}

func updateDataGORM(n int) {
	// Connect to the database.
	db, err := getGORMEngine()
	if err != nil {
		fatal("can't initialize xorm db", err)
	}

	defer func() {
		fatal("can't close connection to database", db.Close())
	}()

	// Get the first index.
	first := getFirstBusStopID(db.DB())

	// Randomly generate and write fake data to the database.
	for i := first; i < n; i++ {
		busStop := getBusStopForUpdate(i)

		fatal("can't update bus stops", db.Save(&busStop).Error)
	}
}

func selectDataGORM(n int) {
	// Connect to the database.
	db, err := getGORMEngine()
	if err != nil {
		fatal("can't initialize xorm db", err)
	}

	defer func() {
		fatal("can't close connection to database", db.Close())
	}()

	// Get the first index.
	first := getFirstBusStopID(db.DB())

	// Randomly generate and write fake data to the database.
	for i := first; i < n; i++ {
		var busStop BusStop

		// TODO: Try using First() and see if there's any difference.
		fatal("can't select bus stops", db.Where("id = ?", i).Take(&busStop).Error)
	}
}
