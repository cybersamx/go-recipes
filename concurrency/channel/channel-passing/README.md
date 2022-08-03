# Channel Message Passing

This is a simple example of declaring and sending/receiving messages between 2 goroutines over a buffered channel.

## Buffered vs Unbuffered

There are 2 types of channels:

* Unbuffered channel - Synchronous channel. Only one operation can be performed at a time. It has the length and capacity of 0 ie. `make(chan int)` is the same as `make(chan int, 0)`.
  * When a value is sent to a sending channel, the control on the sender will be blocked until we read from the channel.
  * When a value is received from a receiving channel, the control on the reader will be blocked until we write to the channel.
* Buffered channel - Asynchronous channel. This means `make(chan int, n)` where `n` >= 1.
  * When there's a buffer in the sending, the sender won't be blocked until the buffer is fully filled.

In [main.go](main.go), we can change the statement to declare either a buffered or unbuffered channel.

```go
// When we declare the channel as unbuffered.
stream := make(chan int)
```

The program will produce the following output:

```
Send 0
Received 0
Send 1
Received 1
Send 2
Received 2
Send 3
Received 3
Send 4
Received 4
Send 5
Received 5
Done
```

As you can see the messages are synchronized as send and receive operations are synchronous.

Let's now declare the channel as buffered (with cap of 3) and see what happens.

```go
// When we declare the channel as unbuffered.
stream := make(chan int, 3)
```

And here's the output:

```
Send 0
Received 0
Send 1
Send 2
Send 3
Received 1
Send 4
Received 2
Send 5
Received 3
Received 4
Received 5
Done
```

## Deadlock

In [main.go](main.go), we have a goroutine that sends out messages to the channel and a goroutine that receives messages from that channel. In the receiving goroutine, we use the following code to read from the channel.

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
