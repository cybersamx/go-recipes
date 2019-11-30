# Connecting to MongoDB

An example of how you would connect to MongoDB and perform basic CRUD operations.

## Requirements 

* MongoDB version 4 and above.

## Setup

1. Launch a shell session and start MongoDB (server) via Docker:

   ```bash
   $ docker-compose up mongo-server
   ```
   
1. Launch a separate shell session and run the mongo cli via Docker:

   ```bash
   $ docker-compose run mongo-cli
   ```
   
1. You can press CTRL-C and CTRL-D on the `mongo-server` and `mongo-cli` respectively to stop the containers. Don't forget to remove the containers as well.

   ```bash
   $ # This should remove both mongo-server and mongo-cli
   $ docker-compose down
   ```

## Notes

* We deliberately hard code the password and use the global root user to connect to MongoDB to keep the sample code simple and easy to understand. Do not do this in your code, especially production code.

## Further Reading

* [MongoDB Connection String URI](https://docs.mongodb.com/manual/reference/connection-string/)
