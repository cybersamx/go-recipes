# Timeout Context

A simple demonstration of using `context.WithTimeout()` and `context.WithDeadline()` to implement some sort of timeout.

The timeout feature in package `context` enables a powerful design pattern for executing tasks concurrently with some timeout capabilities. If the context times out sooner than the background task, the main program will be notified via a channel, allowing the program regain control of execution flow.

There are 2 programs and they are essentially the same with the exception that the context in`timeout` is created using `context.WithTimeout()` and the context in `deadline` is creating suing `context.WithDeadline()`.

## Setup

1. Run the programs. 

   ```bash
   $ go run deadline/main.go
   $ go run timeout/main.go
   ```

## Reference

* [Godoc: Context](https://godoc.org/context)
