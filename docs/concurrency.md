# Concurrency

This is a general guide to concurrency in Go. This doc is broken into 2 key parts:

* **Corountine** - A coroutine is a generic term for a software construct that allows a block of code to be executed concurrently. In Go, a coroutine is called goroutine.
* **Synchronization** - Syncrhonization involves the coordination by the execution system in coordinating the sharing of system resources (data and execution scheduling). Go offers 2 ways for programmers to control synchronization:
  * **Traditional way** - Using more traditional means like concurrency counter, mutex, etc to coordinate coroutines execution and manage shared resources. The `sync` package provides such support.
  * **CSP way** - Out of the box, Go supports concurrency through a technique called Communicating Sequential Processes (or CSP). Instead of synchronizing by controlling access to memory, Go uses message passing between multiple coroutines through the use of a primitive called `channel`.

# Goroutine

To run something concurrently in go, just preface a function (and anonymous function works too) with the keyword `go`.

## Notes

* The main function ie. `func main() { ... }` is a goroutine. In other words, there's always at least 1 goroutine is a Go program.

# Channel

Channels are the conduits that enable message passing between concurrent goroutines.

All receiving channels will block the current goroutine until there's data in the channel to receive. See next section for details.

## Unbuffered Channel vs Buffered Channel

This recipe contrasts an unbuffered channel vs a buffered channel.

|                        | Unbuffered Channel | Buffered Channel |
|------------------------|--------------------|------------------|
| Type of Communication  | Synchronous        | Asynchronous     |
| Receiver               | Receiver blocks until there's data to receive | Receiver blocks until there's data to receive |
| Sender                 | The sender blocks until the receiver has received the value ||

To declare a buffered channel, simply declare with a capacity.

```go
myChan := make(chan string, 5)
```

For unbuffered, just create a channel without the second argument. Also `make(chan string)` is the same as `make(chan string, 0)`.

## Sending vs Receiving Channel

We can declare a unidirectional channel (receive-only or send-only) this way:

```go
// Receive-only concurrency
recvChan := make(<-chan string)

// Send-only concurrency
sendChan := make(chan<- string)
```

But in most declarations, a bidirectional channel declaration is used. So the following is accepted by convention.

```go
myChan := make(chan string)
```

### Receiving Values from Channel

There are 2 ways to receive a value from a channel:

* Using the `<-` operator.
* USing the `range` keyword.

```go
msg := <-myChan
fmt.Println(msg)
```

```go
msg := range myChan {
    fmt.Println(msg)
}
```

**Note**: Unlike the conventional range that is used in iteration, range over a channel only returns a single argument back to the caller.

# Sync Package

## WaitGroup

WaitGroup is a like a counter to keep track of the number of coroutines currently running. The `Add` method increments the counter, `Done` decrements the counter, and `Wait` blocks the current goroutine untils the counter is 0.
