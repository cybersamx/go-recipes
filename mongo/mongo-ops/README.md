# Mongo Operations

There are Mongo operations that are used to run commands against a MongoDB database. This recipe is a collection of useful Mongo operations that are set and executed from Go using the [official MongoDB Go driver](https://github.com/mongodb/mongo-go-driver).

## Notes

* Many MongoDB commands can be run using the `RunCommand` function in the `mongo` package. See [helpers.go](helpers.go) for an example.

## Setup

Here's a sequence of steps to run this project.

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
   $ make run
   $ make teardown    # Run this to remove the container
   ```

## Reference and Credits

* [MongoDB Go Driver API Guide](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
* [Expire Data from Collections by Setting TTL](https://docs.mongodb.com/manual/tutorial/expire-data/)
* [Working with MongoDB TTL Index](http://hassansin.github.io/working-with-mongodb-ttl-index)
