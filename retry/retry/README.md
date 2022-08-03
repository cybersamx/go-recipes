# Retry Using Retry

An example of using the [flowchartsman/retry](https://github.com/flowchartsman/retry) package to retry connecting to MongoDB after failed attempts. The package implements an [exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff) algorithm that multiplicatively decreases the retry rate, allowing for a more suited use in network retransmission.

We deliberately don't want to connect to the database to test retry.

## Setup

1. Run the application. Since we don't have mongo running in the background, the program will retry 5 times before it finally exits.

   ```bash
   $ make run
   ```

1. If you want to start mongo so that the retry works, run the following.

   ```bash
   $ docker-compose up
   ```

## Reference and Credits

* [Github: MongoDB Driver](https://github.com/mongodb/mongo-go-driver)
* [Flowchartsman Retry Package](https://github.com/flowchartsman/retry)
