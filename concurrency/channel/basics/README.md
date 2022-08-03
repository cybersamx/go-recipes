# Channel Basics

Here are the basics on Go channels.

## Channel Operations

3 types of operations that can be performed on channels.

* In stream (send) - Send a value into a channel ie. `ch <- variable`.
* Out stream (receive) - Receive a value from of a channel ie. `variable <- ch`.
  * Goroutines with sending channels may block forever if there are no receiving channels.
* Close a channel
  * Pushing a value into a closed channel will cause a panic.
  * The value read from a closed channel will be a zero value.
  * Close the sending channel, not the receiving channel as this ensures proper control of channel to prevent race condition and panics (as send to a closed channel cause a panic and a receive from a closed channel returns zero value).

## Channel Types

Capacity of a channel is defined as the number of values a channel can hold. Length is defined as the number of values that are currently enqueued in the channel buffer. For example:

```go
ch := make(chan int, 5)
ch <- 1
ch <- 2
// Output: cap: 5, len: 2
fmt.Printf("cap: %d, len: %d\n", cap(ch), len(ch))
```

2 types of channels:

* Unbuffered channel - Synchronous channel. Only one operation can be performed at a time. It has the length and capacity of 0 ie. `make(chan int)` is the same as `make(chan int, 0)`.
* Buffered channel - Asynchronous channel. This means `make(chan int, n)` where `n` >= 1.

## Notes on Concurrency

See more notes written about Concurrency [here](../../../docs/concurrency.md).

## Setup

1. Run the program.

   ```bash
   $ make run
   ```

## Reference

* [Go: Channels](https://golang.org/doc/effective_go.html#channels)
