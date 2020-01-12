# ORM vs SQL

Recipe to show how to insert, update, and select with the standard `sql` package as well as with popular ORM packages such as `xorm`. The recipe also demonstrates the use of go benchmark to shows the comparative performance between the different packages. 

## Requirements

* Docker
* Docker Compose

This recipe uses Postgres. We will be using Docker to spin up an instance of Postgres for the running of this recipe.

## Setup

1. Launch a shell session and start Postgres via Docker:

   ```bash
   $ docker-compose up
   ```

1. You can connect to Postgres from the shell:

   ```bash
   $ docker exec -it postgres psql postgres://pguser:password@localhost/db
   ```

1. Run the program

   ```bash
   $ go test -bench=. -timeout=30m
   ```

1. When all is done, teardown Postgres.

   ```bash
   $ docker-compose down
   ```

## Benchmark Results

**Environment:** MacOS Mojave, Go 1.13.4, Mac Pro 3.5GHz Intel Xeon.

![Benchmark: Time taken to complete a call to the database](images/benchmark-speed.svg)

![Benchmark: Memory used to complete a call to the database](images/benchmark-memory.svg)

![Benchmark: Number of allocations made to complete a call to the database](images/benchmark-allocations.svg)

### Take-Aways

* The ORM packages are less resource efficient than the standard package `sql`.
* XORM performs surprisingly well in speed, almost as fast as the standard package `sql`.
* Speed performance for GORM isn't as good as the other 2 packages for insert and update operations. For select the 3 packages perform equally well.

## Caveats and Notes

Running benchmark takes patience and tweaking so that we can get the right dataset to make a good conclusion of the results. The above charts are just the initial benchmark results, I will make incremental improvements to the code.

* Benchmark is contextual and its result relative. It looks like insert is an expensive operation in terms of speed and memory. But if you delve into the code, the insert operation is comprised of insertions of 1 parent row and 5 child rows. While the update and select operations are associated with just 1 row. This explains why the insert result is multiple times those of update and select. Future work will include some form of normalization.
* Ideally we should isolate and run the benchmarks separately. Overhead from another benchmark may affect another benchmark. In the initial runs, there were definitely some spillover from previous runs. Using the `ResetTimer()` function seems to help and so far the benchmark runs have been quite consistent. May run a benchmark individually in the future.
* No transaction is used.

## Reference and Credits

* [Golang: Package testing](https://golang.org/pkg/testing/)
* [XORM](https://xorm.io/)
* [GORM](https://gorm.io/)
