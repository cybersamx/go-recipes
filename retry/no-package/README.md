# Retries on MongoDB

An example of how you would keep retrying to connect to MongoDB.

For this example, we deliberately don't want to connect to the database to test retry.

## Setup

1. Run the application

   ```bash
   $ make run
   ```

## Reference and Credits

* [Github: MongoDB Driver](https://github.com/mongodb/mongo-go-driver)
* [Gotchas of Defer in Go](https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01)
