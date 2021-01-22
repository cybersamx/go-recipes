# Retries on MongoDB

An example of how you would keep retrying to connect to MongoDB using the [flowchartsman/retry](https://github.com/flowchartsman/retry) package.

It has a neat interface and implements an [exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff) algorithm that multiplicatively decrease the retry rate for more suited for network retransmission.

## Setup

1. Run the application. Since we don't have mongo running in the background, the program will retry 5 times before it finally exits.

   ```bash
   $ make run
   ```

1. On a separate shell, starts mongo.

   ```bash
   $ docker-compose up
   ```

1. Go back to the original shell and run the application again. Now the program should connect to the datastore.

   ```bash
   $ make run
   ```

## Reference and Credits

* [Github: MongoDB Driver](https://github.com/mongodb/mongo-go-driver)
* [Flowchartsman Retry Package](https://github.com/flowchartsman/retry)
