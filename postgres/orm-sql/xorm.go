package main

import (
	"github.com/cybersamx/go-recipes/random/rand"
	"time"
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
	// Setup for random generation.
	rand.Seed()
	now := time.Now()
	past := now.Add(-3 * 365 * 24 * time.Hour)

	// Connect to the database.
	engine, err := getXORMEngine()
	if err != nil {
		fatal("can't initialize xorm engine", err)
	}

	session := engine.NewSession()
	defer session.Close()

	// Get the last index.
	last := getLastUserID(engine.DB().DB)

	// Randomly generate and write fake data to the database.
	for i := last + 1; i <= n + last; i++ {
		user := getUser(i)

		_, err := session.Insert(&user)
		if err != nil {
			fatal("can't insert users", err)
		}

		for j := 0; j < 5; j++ {
			restaurant := getRestaurant(i, past, now)

			_, err = session.Insert(&restaurant)
			if err != nil {
				fatal("can't insert restaurants", err)
			}
		}
	}
}

