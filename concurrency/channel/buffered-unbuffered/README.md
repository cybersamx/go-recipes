# Buffered vs Unbuffered Channels

The example shows the difference between buffered and unbuffered channels.

## Channel Capacity and Length

Capacity of a channel is defined as the number of values a channel can hold. Length is defined as the number of values that are currently enqueued in the channel buffer. For example:

```go
ch := make(chan int, 5)
ch <- 1
ch <- 2
// Output: cap: 5, len: 2
fmt.Printf("cap: %d, len: %d\n", cap(ch), len(ch))
```

## Buffered vs Unbuffered

There are 2 types of channels:

* Buffered channel - Asynchronous channel. This means `make(chan int, n)` where `n` >= 1.
  * A goroutine that writes to a sending channel will be blocked until there is room available in the buffer.
  * A goroutine reading a value from a receiving channel will be blocked until a value is placed into the channel.
* Unbuffered channel - Synchronous channel. It has the length and capacity of 0 ie. `make(chan int)` is the same as `make(chan int, 0)`. This means that channel always full and the associated goroutine will be blocked after writing to the channel unless the receiving channel reads the value.

## Notes on Concurrency

See more notes written about Concurrency [here](../../../docs/concurrency.md).

## Setup

1. Run the program.

   ```bash
   $ make run
   ```

## Reference

* [Go: Channels](https://golang.org/doc/effective_go.html#channels)
