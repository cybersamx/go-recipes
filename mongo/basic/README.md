# Connecting to MongoDB

An example of how you would connect to MongoDB and perform basic CRUD operations.

There's another recipe that performs the same task but it's implemented using a pre-defined schema that helps the mapping of MongoDB documents to Go data structures. The alternate recipe can be found [here](../schema).

## Setup

1. To run everything in 1 command ie. spin up the mongo container, run the go program, and finally tear down the mongo container, just run the following:

   ```bash
   $ make
   ```

To do everything step-by-step.

1. You can connect to Mongo via the console:

   ```bash
   $ docker exec -it mongo mongo -u nobody -p secrets go-recipes
   ```

1. You can press CTRL-C and CTRL-D on the `mongo-server` and `mongo-cli` respectively to stop the containers. Don't forget to remove the containers as well.

   ```bash
   $ # This should remove both mongo-server and mongo-cli
   $ docker-compose down
   ```

1. Run Go program.

   ```bash
   $ go run ./main.go
   ```

1. Alternatively, you can run everything with just 2 commands.

   ```bash
   $ make
   $ make teardown    # Run this to remove container
   ```

## Reference and Credits

* [MongoDB Connection String URI](https://docs.mongodb.com/manual/reference/connection-string/)
* [MongoDB Go Driver API Guide](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
* [Stackoverflow: How to Create a DB for MongoDB](https://stackoverflow.com/questions/42912755/how-to-create-a-db-for-mongodb-container-on-start-up)
