# Global, Atomic Counter in Redis

The 2 most mature and commonly used Redis drivers in Go are:

* [gomodule/redigo](https://github.com/gomodule/redigo)
* [mediocregopher/radix](https://github.com/mediocregopher/radix)

A full list of Redis drivers for Go is available [here](http://redis.io/clients#go).

In this recipe, we will implement a global, [atomic](https://en.wikipedia.org/wiki/Atomicity_(database_systems)) counter using Redis using the [Radix](https://github.com/mediocregopher/radix) driver. Why not just implement a counter locally in the Go program instead. Sure, we can do that if it's just 1 client running a single thread. But in a concurrent context where we have multiple goroutines or in a distributed system, we need to have a system to support an atomic increment operation across the network. Redis provides support for a counter via the `INCR` command.

## Setup

1. Start redis.

   ```bash
   $ docker-compose up
   ```

1. In the another shell, you can connect to Redis via the CLI tool:

   ```bash
   $ docker-compose exec redis redis-cli
   ```

1. Run Go program.

   ```bash
   $ go run ./main.go
   ```

1. Shut down and remove the container when you are done.

   ```bash
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
