# MongoDB Change Stream

MongoDB Change Stream allows application to subscribe to data chanages in the collection or database in an MongoDB instance.

For Change Stream to work, MongoDB must run in replica set mode. So this recipe includes Docker setup for running MongoDB replica set locally.

## Requirements

* MongoDB version 4 and above.

## Setup

1. Launch a shell session and start MongoDB (server) via Docker:

   ```bash
   $ docker-compose up
   ```

1. Run Go program:

   ```bash
   $ go run ./main.go
   ```

1. Alternatively, you can run everything with the following command:

   ```bash
   $ make
   $ CTRL-C
   $ # After you are done with the program, you should stop/remove the container.
   $ docker-compose down
   ```

## Notes

* We deliberately hard code the password to connect to MongoDB to keep the sample code simple and easy to understand. Do not do this in your code, especially production code.

## Reference and Credits

* [MongoDB Change Stream](https://docs.mongodb.com/manual/changeStreams/)
* [MongoDB Connection String URI](https://docs.mongodb.com/manual/reference/connection-string/)
* [MongoDB Go Driver API Guide](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
