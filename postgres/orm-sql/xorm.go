package main

import (
	"xorm.io/core"
	"xorm.io/xorm"
)

func getXORMEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(dialect, getDSN())
	if err != nil {
		return nil, err
	}

	engine.SetMapper(core.GonicMapper{})

	return engine, nil
}

func insertDataXORM(n int) {
	// Connect to the database.
	engine, err := getXORMEngine()
	if err != nil {
		fatal("can't initialize xorm engine", err)
	}

	session := engine.NewSession()
	defer session.Close()

	// Get the last index.
	last := getLastBusStopID(engine.DB().DB)

	// Randomly generate and write fake data to the database.
	for i := last + 1; i <= n + last; i++ {
		user := getBusStop(i)
		_, err := session.Insert(&user)
		if err != nil {
			fatal("can't insert bus stops", err)
		}
	}
}

func updateDataXORM(n int) {
	// Connect to the database.
	engine, err := getXORMEngine()
	if err != nil {
		fatal("can't initialize xorm engine", err)
	}

	session := engine.NewSession()
	defer session.Close()

	// Get the first index.
	first := getFirstBusStopID(engine.DB().DB)

	// Randomly generate and write fake data to the database.
	for i := first; i < n; i++ {
		busStop := getBusStopForUpdate(i)
		_, err = engine.ID(i).Update(&busStop)
		if err != nil {
			fatal("can't update bus stops", err)
		}
	}
}

func selectDataXORM(n int) {
	// Connect to the database.
	engine, err := getXORMEngine()
	if err != nil {
		fatal("can't initialize xorm engine", err)
	}

	session := engine.NewSession()
	defer session.Close()

	// Get the first index.
	first := getFirstBusStopID(engine.DB().DB)

	// Randomly generate and write fake data to the database.
	for i := first; i < n; i++ {
		var busStop BusStop

		_, err = engine.ID(i).Get(&busStop)
		if err != nil {
			fatal("can't select bus stops", err)
		}
	}
}

