# Deadline Context

A simple demonstration of using `context.WithDeadline()` to implement some sort of timeout.

This program takes in arguments x, y and will run a task for x seconds with a timeout for y.
* If x > y, the task will be canceled via context cancel.
* If x < y, the task will complete.

## Setup

1. Run the programs.

   ```bash
   $ make run
   ```

## Reference

* [Godoc: Context](https://godoc.org/context)
