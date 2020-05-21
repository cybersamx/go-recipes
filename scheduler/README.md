# Recurring Task

Scheduling task, such as cleaning up expired resources, rotating keys, periodically is a common requirement for any system program. This recipe highlights the various ways of running an recurring task in the background in Go:

* **Strategy 1** - Using `time.Sleep()`.
* **Strategy 2** - Using `time.Ticker` #1.
* **Strategy 3** - Using `time.Ticker` #2.
* **Strategy 4** - Using `time.After()`.

## Setup

1. Run the program

   ```bash
   $ make run
   ```
