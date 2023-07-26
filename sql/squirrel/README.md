# Squirrel

An example of connecting to a postgres database using the `sqlx` and `squirrel` packages

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

   db=# SELECT * FROM users;
    id  |         created_at         |         updated_at         |      email      | username | first_name | last_name
   -----+----------------------------+----------------------------+-----------------+----------+------------+-----------
    123 | 2023-07-25 00:04:13.652829 | 2023-07-25 00:04:13.652829 | jon@example.com | jonny    | Jon        | Murphy
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

* [Sqlx homepage](https://jmoiron.github.io/sqlx/)
* [Github: Squirrel](https://github.com/Masterminds/squirrel)
