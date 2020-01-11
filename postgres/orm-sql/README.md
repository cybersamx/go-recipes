# ORM vs SQL

Recipe to show how to insert, update, and select with the standard `sql` package as well as with popular ORM packages such as `xorm`. The recipe also demonstrates the use of go benchmark to shows the comparative performance between the different packages. 

## Requirements

* Docker
* Docker Compose

This recipe uses Postgres. We will be using Docker to spin up an instance of Postgres for the running of this recipe.

## Setup

1. Launch a shell session and start Postgres via Docker:

   ```bash
   $ docker-compose up
   ```

1. You can connect to Postgres from the shell:

   ```bash
   $ docker exec -it postgres psql postgres://pguser:password@localhost/db
   ```

1. Run the program

   ```bash
   $ go test -bench=.
   ```

1. When all is done, teardown Postgres.

   ```bash
   $ docker-compose down
   ```

## Reference and Credits

* [Golang: Package testing](https://golang.org/pkg/testing/)
