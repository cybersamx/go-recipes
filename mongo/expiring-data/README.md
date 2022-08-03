# TTL Collection

There are times when we need to remove data in a Mongo collection after a specific time. This example shows how we can remove expired data using [MongoDB TTL indexing](https://www.mongodb.com/docs/manual/core/index-ttl/).

## Notes

1. MongoDB runs the TTLMonitor on a separate thread that looks for TTL indexes in all collections.
1. By default, the TTLMonitor runs every 60 seconds. Read [here](http://hassansin.github.io/working-with-mongodb-ttl-index) for more information.
1. To enable TTLMonitor in our MongoDB service, the file [setup-ttl.js](./setup-ttl.js) is created and mounted on the Mongo container. The file will be ingested and executed by Mongo during boot-up to enable TTLMonitor.
1. It's important that the time field of the TTL index is set to is of type `time.Time`, any other data types will not trigger the object to be removed after expiration.
1. Since the code includes dropping a collection, we added an `dbAdmin` role to the user - see [create-user.js](./create-user.js).

## Setup

1. Launch a shell session and start MongoDB (server) via Docker:

   ```bash
   $ docker-compose up
   ```

1. Run Go program.

   ```bash
   $ go run ./main.go
   ```

1. You can connect to Mongo via the console and see that a token object that is set to expire in a minute, was inserted to the collection `tokens`.

   ```bash
   $ docker exec -it mongo mongo -u nobody -p secrets go-recipes
   > db.tokens.find()
   { "_id" : ObjectId("5ffff0bd4130100b939e597d"), "value" : "SomeSecretValue", "expireAt" : ISODate("2021-01-14T07:21:29.593Z") }
   > // Wait for a little more than 60 seconds
   > db.tokens.find()
   > // The object has expired and removed
   ```

1. Tear down the mongo docker container when done.

   ```bash
   $ docker-compose down
   ```

1. Alternatively, you can run everything with just 2 commands.

   ```bash
   $ make run
   $ make teardown    # Run this to remove containers
   ```

## Reference and Credits

* [MongoDB Go Driver API Guide](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
* [Expire Data from Collections by Setting TTL](https://docs.mongodb.com/manual/tutorial/expire-data/)
* [Working with MongoDB TTL Index](http://hassansin.github.io/working-with-mongodb-ttl-index)
