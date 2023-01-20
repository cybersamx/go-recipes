package main

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	// If we want to use the pq driver, rename the driver to 'postgres` and use SAN
	// in the cert's CN.
	//_ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const driver = "pgx"

func noTLSConnect() (*sql.DB, error) {
	db, err := sql.Open(driver, "host=localhost port=5432 user=postgres password=password dbname=postgres")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func tlsConnect() (*sql.DB, error) {
	// Let's connect to postgres using tls.

	// TODO: The dsn for postgres works now. However, we still can't use sslmode=verify-full. Self-signed CA is a pain. Using sslmode=verify-ca for now.
	db, err := sql.Open(driver, "host=localhost port=5432 user=postgres sslmode=verify-ca sslrootcert=docker/postgresql/ca-cert.pem sslcert=docker/postgresql/client-cert.pem sslkey=docker/postgresql/client-key.pem")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if !ok {
			fmt.Println(pgErr)
		}

		return nil, err
	}

	return db, nil
}

func main() {
	db, err := noTLSConnect()
	if err != nil {
		fmt.Println(err)
		fmt.Println("This above error is expected when connecting to postgres without tls. The program will now connect to postgres using tls.")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	tdb, err := tlsConnect()
	if err != nil {
		panic(err)
	}
	defer func() {
		if tdb != nil {
			tdb.Close()
		}
	}()

	fmt.Println("Successfully connected to postgres using tls")
}
