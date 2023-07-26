package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Open("postgres", postgresDSN())
	if err != nil {
		fatal("can't connect to the database", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			fatal("can't close the connection to the database", err)
		}
	}()

	clearTable(db)

	insertSQL(db, 123, "jon@exapmple.com", "jonny", "Jon", "Murphy")
	insertSQL(db, 234, "linda@exapmple.com", "linda", "Linda", "Lee")

	user := selectSQL(db, 123)
	fmt.Println(user)
	user = selectSQL(db, 234)
	fmt.Println(user)
}
