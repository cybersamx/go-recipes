# Channel Basics

Here are the basics on Go channels.

## Channel Operations

Here are the basic operations that can be performed on channels.

* Writing to stream (send) - Send a value into a channel ie. `ch <- variable`.
* Reading from stream (receive) - Receive a value from of a channel ie. `variable <- ch`.
  * Goroutines with sending channels may block forever if there are no receiving channels.
* Closing a channel
  * Pushing a value into a closed channel will cause a panic.
  * The value read from a closed channel will be a zero value.
  * Close the sending channel, not the receiving channel as this ensures proper control of channel to prevent race condition and panics (as writing to a closed channel cause a panic and a reading from a closed channel returns zero value).

## Setup

1. Run the program.

   ```bash
   $ make run
   ```

## Reference

* [Go: Channels](https://golang.org/doc/effective_go.html#channels)
