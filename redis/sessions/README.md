# Expiring Sessions

Web sessions are ephemeral. We can leverage the `EXPIRE` command in Redis to mark a key to be automatically removed from Redis when it expires. In this example, we explore how we can implement expiring web sessions in redis:

* Use the [Radix](https://github.com/mediocregopher/radix) driver.
* Use the [Go-Redis](https://redis.uptrace.dev/) driver.
* Use GOB serialization to save complex types into Redis.

## Setup

1. Start redis via docker compose:

   ```bash
   $ docker-compose up
   ```

1. In the another shell, you can connect to Redis via the CLI tool:

   ```bash
   $ docker-compose exec redis redis-cli
   127.0.0.1:6379> KEYS sessions
   ```

1. Run Go program.

   ```bash
   $ go run ./main.go
   ```

1. Shut down and remove the container when you are done.

   ```bash
   $ # This should remove both redis-server and redis-cli
   $ docker-compose down
   ```

1. Alternatively, you can run everything with just 2 commands.

   ```bash
   $ make run
   $ make teardown    # Run this to remove the container
   ```

## Reference and Credits

* [Tutorial: Design and implementation of a simple Twitter clone using PHP and the Redis key-value store](https://redis.io/topics/twitter-clone)
* [GoDoc: mediocreogopher/radix](ttps://godoc.org/github.com/mediocregopher)
* [Go-Redis homepage](https://redis.uptrace.dev/)
