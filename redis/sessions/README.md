# Expiring Sessions

Web sessions are ephemeral. We can leverage the `EXPIRE` command in Redis to specify a key to be automatically removed from Redis when it expires. In this recipe, we explore how we can implement expiring web sessions in Redis with the following highlights:

* Use the [Radix](https://github.com/mediocregopher/radix) driver.
* Use serialization save complex types into Redis.

## Setup

1. Start MongoDB (server) via Docker Compose:

   ```bash
   $ docker-compose up
   ```

1. In the another shell, you can connect to Redis via the CLI tool:

   ```bash
   $ docker-compose exec redis redis-cli
   127.0.0.1:6379> keys sessions
   ```

1. Shut down and remove the container when you are done.

   ```bash
   $ # This should remove both mongo-server and mongo-cli
   $ docker-compose down
   ```

1. Run Go program.

   ```bash
   $ go run ./main.go
   ```

## Reference and Credits

* [Tutorial: Design and implementation of a simple Twitter clone using PHP and the Redis key-value store](https://redis.io/topics/twitter-clone)
* [GoDoc: mediocreogopher/radix](ttps://godoc.org/github.com/mediocregopher)
