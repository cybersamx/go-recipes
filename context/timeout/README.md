# Context Timeout

A simple example of using `context.WithTimeout()` to implement some sort of timeout.

* Task completes if time to complete task < context timeout.
* Task timeout via `context.Done` if time to complete task > context timeout.

## Setup

1. Run the programs.

   ```bash
   $ make run
   ```

## Reference

* [Godoc: Context](https://godoc.org/context)
