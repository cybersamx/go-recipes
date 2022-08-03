# SQL

An example of connecting to a postgres database using the `sql` package.

## Requirements

* Docker
* Docker Compose

We will be using Docker to spin up a Postgres instance.

## Setup

1. Launch a new shell session and start Postgres:

   ```bash
   $ docker-compose up
   ```

1. You can connect to Postgres from the shell:

   ```bash
   $ docker exec -it postgres psql postgres://pguser:password@localhost/db
   psql (12.11)
   Type "help" for help.

   db=# SELECT * FROM locations;
                     id                  |         updated_at         | latitude  | longitude
   --------------------------------------+----------------------------+-----------+------------
    3fc58bb9-e40b-4b32-9140-08ce178e4c06 | 2022-07-31 22:45:26.599539 | -37.79973 | -31.711046

   ```

1. Run the program

   ```bash
   $ go run .
   ```

1. When all is done, teardown Postgres.

   ```bash
   $ docker-compose down
   ```

## Reference and Credits

* [Golang: Package database/sql](https://golang.org/pkg/database/sql/)
