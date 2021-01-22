# Mongo Operations

There are Mongo operations that are used to run commands against a MongoDB database. This recipe is a collection of useful Mongo operations that are set and executed from Go using [the official MongoDB Go driver](https://github.com/mongodb/mongo-go-driver).

## Notes

* Many MongoDB commands can be run using the `RunCommand` function in the `mongo` package. See [helpers.go](helpers.go) for an example.

## Setup

1. To run everything in 1 command ie. spin up the mongo container, run the go program, and finally tear down the mongo container, just run the following:

   ```bash
   $ make
   ```

To do everything step-by-step.

1. Launch a shell session and start MongoDB (server) via Docker:

   ```bash
   $ docker-compose up
   ```

1. Run Go program.

   ```bash
   $ go run ./main.go
   ```

1. Tear down the mongo docker container when done.

   ```bash
   $ docker-compose down
   ```

1. Alternatively, you can run everything with just 2 commands.

   ```bash
   $ make
   $ make teardown    # Run this to remove the container
   ```

## Reference and Credits

* [MongoDB Go Driver API Guide](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
* [Expire Data from Collections by Setting TTL](https://docs.mongodb.com/manual/tutorial/expire-data/)
* [Working with MongoDB TTL Index](http://hassansin.github.io/working-with-mongodb-ttl-index)
