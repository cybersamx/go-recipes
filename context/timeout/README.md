# Timeout Context

A simple demonstration of using `context.WithTimeout()` to implement some sort of timeout.

This program will run a task for x seconds with timeout for y.
* If x > y, the task will be canceled via context cancel.
* If x < y, the task will complete.

The timeout feature in package `context` enables a powerful design pattern for executing tasks concurrently with some timeout capabilities. If the context times out sooner than the background task, the main program will be notified via a channel, allowing the program regain control of execution flow.

## Setup

1. Run the programs.

   ```bash
   $ make run
   ```

## Reference

* [Godoc: Context](https://godoc.org/context)
