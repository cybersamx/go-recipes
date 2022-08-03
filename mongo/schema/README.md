# Connecting to MongoDB

An example of how you would connect to MongoDB and perform basic CRUD operations. In this example, we use a pre-defined schema that facilitates the mapping of MongoDB documents to Go data structures

There's [another example](../simple) that performs the same operation but without using a pre-defined schema.

## Setup

Here's a sequence of steps to run this project.

1. Launch a shell session and start MongoDB (server) via Docker:

   ```bash
   $ docker-compose up
   ```

1. You can connect to Mongo via docker interactive tty interface.

   ```bash
   $ docker exec -it mongo mongo -u nobody -p secrets go-recipes
   ```

1. Run Go program.

   ```bash
   $ go run ./main.go
   ```

1. When you are done, just shut mongo down.

   ```bash
   $ docker-compose down
   ```

1. Alternatively, you can run everything with just 2 commands.

   ```bash
   $ make
   $ make teardown    # Run this to remove the container
   ```

## Reference and Credits

* [MongoDB Connection String URI](https://docs.mongodb.com/manual/reference/connection-string/)
* [MongoDB Go Driver API Guide](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
* [Stackoverflow: How to Create a DB for MongoDB](https://stackoverflow.com/questions/42912755/how-to-create-a-db-for-mongodb-container-on-start-up)
