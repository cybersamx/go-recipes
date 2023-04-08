# Cancel Context

Let's spawn several long-running goroutines. We will pass the same context to each of these goroutines. With a call to the cancel function, we should be able to stop the execution of all goroutines.

## Setup

1. Run the programs.

   ```bash
   $ make run
   ```

## Reference

* [Godoc: Context](https://godoc.org/context)
