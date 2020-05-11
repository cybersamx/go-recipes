# Connecting to MongoDB

An example of how you would connect to MongoDB and perform basic CRUD operations.

## Requirements

* MongoDB version 4 and above.

## Setup

1. Launch a shell session and start MongoDB (server) via Docker:

   ```bash
   $ docker-compose up
   ```

1. You can connect to Mongo via the console:

   ```bash
   $ docker exec -it mongo mongo -u nobody -p secrets go-recipes
   ```

1. You can press CTRL-C and CTRL-D on the `mongo-server` and `mongo-cli` respectively to stop the containers. Don't forget to remove the containers as well.

   ```bash
   $ # This should remove both mongo-server and mongo-cli
   $ docker-compose down
   ```
   
1. Run Go programL

   ```bash
   $ go run ./main.go
   ```

## Notes

* We deliberately hard code the password to connect to MongoDB to keep the sample code simple and easy to understand. Do not do this in your code, especially production code.

## Reference and Credits

* [MongoDB Connection String URI](https://docs.mongodb.com/manual/reference/connection-string/)
* [MongoDB Go Driver API Guide](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
* [Stackoverflow: How to Create a DB for MongoDB](https://stackoverflow.com/questions/42912755/how-to-create-a-db-for-mongodb-container-on-start-up)
