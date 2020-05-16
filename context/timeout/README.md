# Timeout Context

A simple demonstration of using `context.WithTimeout()a` to implement some sort of timeout.

The `timeout` support in package `context` enables a powerful design pattern for executing tasks concurrently with some timeout capabilities. If the context times out sooner than the background task, the main program will regain control after it has been notified through a channel.

## Setup

1. Run the program

   ```bash
   $ go run main.go
   ```

## Reference

* [Godoc: Context](https://godoc.org/context)
