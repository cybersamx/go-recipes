# Retries on MongoDB

An example of how you would keep retrying to connect to MongoDB using the [flowchartsman/retry](https://github.com/flowchartsman/retry) package.

It has a neat interface and implements an [exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff) algorithm that multiplicatively decrease the retry rate for more suited for network retransmission.

## Setup

1. Run the application

   ```bash
   $ make run
   ```

## Reference and Credits

* [Github: MongoDB Driver](https://github.com/mongodb/mongo-go-driver)
* [Flowchartsman Retry Package](https://github.com/flowchartsman/retry)
