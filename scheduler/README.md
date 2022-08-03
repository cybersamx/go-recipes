# Recurring Task

Scheduling tasks, such as cleaning up expired resources, rotating keys, is a common operation for any system program. This example highlights the various ways of running  recurring tasks in the background in Go:

* **Strategy 1** - Using `time.Sleep()`.
* **Strategy 2** - Using `time.Ticker` #1.
* **Strategy 3** - Using `time.Ticker` #2.
* **Strategy 4** - Using `time.After()`.

## Setup

1. Run the program

   ```bash
   $ make run
   ```
