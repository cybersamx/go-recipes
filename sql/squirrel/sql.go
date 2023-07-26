package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const (
	username = "pguser"
	password = "password"
	host     = "localhost"
	port     = 5432
	database = "db"
)

type User struct {
	ID        int        `db:"id"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Age       *int       `db:"age"`
	Email     string     `db:"email"`
	Username  string     `db:"username"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
}

func postgresDSN() string {
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

func clearTable(db *sqlx.DB) {
	result, err := db.Exec("DELETE FROM users")
	if err != nil {
		fatal("can't execute sql", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		fatal("driver doesn't support RowsAffected", err)
	}
	if n == 0 {
		fatal("no rows inserted", err)
	}
}

func insertSQL(db *sqlx.DB, id int, email, username, firstName, lastName string) {
	now := time.Now()

	// We deliberately not set some of the fields.
	user := User{
		ID:        id,
		CreatedAt: &now,
		Email:     email,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
	}

	stmtBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	stmt := stmtBuilder.
		Insert("users").
		Columns("id", "created_at", "updated_at", "email", "username").
		Values(user.ID, user.CreatedAt, user.UpdatedAt, user.Email, user.Username)

	query, args, err := stmt.ToSql()
	if err != nil {
		fatal("can't convert to sql", err)
	}

	result, err := db.Exec(query, args...)
	if err != nil {
		fatal("can't execute sql", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		fatal("driver doesn't support RowsAffected", err)
	}
	if n == 0 {
		fatal("no rows inserted", err)
	}
}

func selectSQL(db *sqlx.DB, id int) *User {
	stmtBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	stmt := stmtBuilder.
		Select("id", "created_at", "updated_at", "email", "username").
		From("users").
		Where("id = ?", id)

	query, args, err := stmt.ToSql()
	if err != nil {
		fatal("can't convert to sql", err)
	}

	var user User
	err = db.Get(&user, query, args...)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil
	case err != nil:
		fatal("can't get data from the datastore", err)
	}

	return &user
}
