# Timed Repetition

An example where a message is printed every few seconds until we interrupt the program with a SIGTERM by calling the context cancel. The delay between prints is implemented with `time.After`.

## Setup

1. Run the programs.

   ```bash
   $ make run
   ```
