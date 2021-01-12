# Basic Example

This is a simple example of declaring a buffered channel and sending a message through it.

The example as shows that all receiving channels (unbuffered or buffered) will block until there's data to receive. Consequently, in this context, we don't need to use WaitGroup and let channel handles the synchronization for us.

## Notes on Concurrency

See more notes written about Concurrency [here](../../../docs/concurrency.md).

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference

* [Go: Channels](https://golang.org/doc/effective_go.html#channels)
