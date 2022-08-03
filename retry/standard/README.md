# Retry

An example of using a standard simple pattern to retry connecting to MongoDB after failed attempts.

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
* [Gotchas of Defer in Go](https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01)
