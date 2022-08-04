# Channel Message Passing

This is a simple example of declaring and sending/receiving messages between 2 goroutines over a buffered channel.

In this example, we set up a channel to send and receive values between 2 goroutines.

## Range over Channel

In [main.go](main.go), if the goroutines are waiting for an empty channel to be written and nothing will be written to the sending channel, we get a deadlock.

```go
go func(ch <-chan int) {
	for i := range ch {
		time.Sleep(2 * time.Second)
		fmt.Printf("Received %d\n", i)
	}
}
```

When we prefix a `range` to a channel, the goroutine will read from the channel continuously **until it is closed**.

So if we don't close the sender channel in the sending goroutine, the receiving goroutine will just keep waiting for message to be received or the channel to be closed. Since both conditions will never be met, we get a deadlock.

Try commenting `close(ch)` out, run the program, and see what happens.

```go
go func(ch chan<- int) {
	for i := 0; i < 6; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("Send %d\n", i)
	}
	close(ch)
}
```

## Notes on Concurrency

See more notes written about Concurrency [here](../../../docs/concurrency.md).

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference

* [Go: Channels](https://golang.org/doc/effective_go.html#channels)
