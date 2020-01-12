package main

import (
	"github.com/cybersamx/go-recipes/random/rand"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"syreclabs.com/go/faker"
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

	defer func(){
		fatal("can't close connection to database", db.Close())
	}()

	// Get the last index.
	last := getLastUserID(db.DB())

	// Randomly generate and write fake data to the database.
	for i := last + 1; i <= n + last; i++ {
		user := getUser(i)

		fatal("can't insert users", db.Create(&user).Error)

		for j := 0; j < 5; j++ {
			restaurant := getRestaurant(i, past, now)

			fatal("can't insert restaurants", db.Create(&restaurant).Error)
		}
	}
}

func updateDataGORM(n int) {
	// Connect to the database.
	db, err := getGORMEngine()
	if err != nil {
		fatal("can't initialize xorm db", err)
	}

	defer func(){
		fatal("can't close connection to database", db.Close())
	}()

	// Get the first index.
	first := getFirstUserID(db.DB())

	// Randomly generate and write fake data to the database.
	for i := first; i < n; i++ {
		user := User{
			ID:   i,
			Name: faker.Name().Name(),
			Age:  rand.RandomIntRange(14, 80),
		}

		fatal("can't update users", db.Save(&user).Error)
	}
}

func selectDataGORM(n int) {
	// Connect to the database.
	db, err := getGORMEngine()
	if err != nil {
		fatal("can't initialize xorm db", err)
	}

	defer func(){
		fatal("can't close connection to database", db.Close())
	}()

	// Get the first index.
	first := getFirstUserID(db.DB())

	// Randomly generate and write fake data to the database.
	for i := first; i < n; i++ {
		var user User

		// TODO: Try using First() and see if there's any difference.
		fatal("can't select users", db.Where("id = ?", i).Take(&user).Error)
	}
}

