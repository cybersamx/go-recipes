# Authentication

This is a simple recipe that performs authentication in Redis using the [Radix](https://github.com/mediocregopher/radix) driver.

## Setup

1. Start MongoDB (server) via Docker Compose:

   ```bash
   $ docker-compose up
   ```

1. In the another shell, you can connect to Redis via the CLI tool:

   ```bash
   $ docker-compose exec redis redis-cli
    127.0.0.1:6379> AUTH secrets
   ```

   Note: `secrets` is the password. See <docker-compose.yaml>.

1. Shut down and remove the container when you are done.

   ```bash
   $ # This should remove both redis-server and redis-cli
   $ docker-compose down
   ```

1. Run Go program.

   ```bash
   $ go run ./main.go
   ```

1. Alternatively, you can run everything with just 2 commands.

   ```bash
   $ make
   $ make teardown    # Run this to remove the container
   ```

## Reference and Credits

* [Tutorial: Design and implementation of a simple Twitter clone using PHP and the Redis key-value store](https://redis.io/topics/twitter-clone)
* [GoDoc: mediocreogopher/radix](ttps://godoc.org/github.com/mediocregopher)
* [Redis: AUTH password](https://redis.io/commands/auth)
