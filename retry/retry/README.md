# Retry Using Retry

An example of using the [flowchartsman/retry](https://github.com/flowchartsman/retry) package for setting a retry mechanism in the app for reconnecting the client to a http server.

The package implements an [exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff) algorithm that multiplicatively decreases the retry rate.

## Setup

1. Run the application.

   ```bash
   $ make run
   ```

## Reference and Credits

* [Flowchartsman Retry Package](https://github.com/flowchartsman/retry)
