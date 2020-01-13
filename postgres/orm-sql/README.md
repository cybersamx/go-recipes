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
   $ go test -bench=.
   ```

1. When all is done, teardown Postgres.

   ```bash
   $ docker-compose down
   ```

## Benchmark Results

**Environment:** MacOS High Sierra, Go 1.12.4, Macbook Pro 2.5GHz Intel Core i7.

Here's the result of a benchmark run:

```text
BenchmarkInsertSQL_1-8       	     200	   8878037 ns/op	   25284 B/op	     299 allocs/op
BenchmarkInsertSQL_10-8      	      50	  34611079 ns/op	   35937 B/op	     579 allocs/op
BenchmarkInsertSQL_100-8     	       5	 274724846 ns/op	  142827 B/op	    3375 allocs/op
BenchmarkInsertSQL_1000-8    	       1	2208230151 ns/op	 1209352 B/op	   31289 allocs/op
BenchmarkUpdateSQL_1-8       	     200	   6230600 ns/op	   24091 B/op	     266 allocs/op
BenchmarkUpdateSQL_10-8      	      50	  27047491 ns/op	   34494 B/op	     545 allocs/op
BenchmarkUpdateSQL_100-8     	       5	 240848268 ns/op	  139310 B/op	    3342 allocs/op
BenchmarkUpdateSQL_1000-8    	       1	2596158384 ns/op	 1182920 B/op	   31237 allocs/op
BenchmarkSelectSQL_1-8       	     200	   6573060 ns/op	   24102 B/op	     266 allocs/op
BenchmarkSelectSQL_10-8      	     100	  24234866 ns/op	   39135 B/op	     680 allocs/op
BenchmarkSelectSQL_100-8     	      10	 186871591 ns/op	  190730 B/op	    4824 allocs/op
BenchmarkSelectSQL_1000-8    	       1	1909938927 ns/op	 1708216 B/op	   47124 allocs/op
BenchmarkInsertXORM_1-8      	     200	   9176330 ns/op	   43820 B/op	     577 allocs/op
BenchmarkInsertXORM_10-8     	      50	  29915722 ns/op	   88971 B/op	    1853 allocs/op
BenchmarkInsertXORM_100-8    	       5	 242410587 ns/op	  541910 B/op	   14614 allocs/op
BenchmarkInsertXORM_1000-8   	       1	2252374369 ns/op	 5072792 B/op	  142346 allocs/op
BenchmarkUpdateXORM_1-8      	     200	   6570845 ns/op	   33002 B/op	     310 allocs/op
BenchmarkUpdateXORM_10-8     	      50	  28260264 ns/op	   94748 B/op	    2044 allocs/op
BenchmarkUpdateXORM_100-8    	       5	 255407720 ns/op	  656480 B/op	   18157 allocs/op
BenchmarkUpdateXORM_1000-8   	       1	2286156692 ns/op	 6270136 B/op	  179288 allocs/op
BenchmarkSelectXORM_1-8      	     200	   6673537 ns/op	   32977 B/op	     309 allocs/op
BenchmarkSelectXORM_10-8     	     100	  23396516 ns/op	  100715 B/op	    2413 allocs/op
BenchmarkSelectXORM_100-8    	      10	 191402344 ns/op	  722137 B/op	   22217 allocs/op
BenchmarkSelectXORM_1000-8   	       1	1724595119 ns/op	 6932032 B/op	  220234 allocs/op
BenchmarkInsertGORM_1-8      	     100	  11585892 ns/op	   31173 B/op	     450 allocs/op
BenchmarkInsertGORM_10-8     	      30	  50307965 ns/op	   90517 B/op	    1997 allocs/op
BenchmarkInsertGORM_100-8    	       3	 421409432 ns/op	  685138 B/op	   17463 allocs/op
BenchmarkInsertGORM_1000-8   	       1	3986506467 ns/op	 6632504 B/op	  172331 allocs/op
BenchmarkUpdateGORM_1-8      	     200	   7376537 ns/op	   24539 B/op	     278 allocs/op
BenchmarkUpdateGORM_10-8     	      30	  43724413 ns/op	   86429 B/op	    1944 allocs/op
BenchmarkUpdateGORM_100-8    	       3	 424758456 ns/op	  705101 B/op	   18607 allocs/op
BenchmarkUpdateGORM_1000-8   	       1	4306743562 ns/op	 6882208 B/op	  185154 allocs/op
BenchmarkSelectGORM_1-8      	     200	   7182942 ns/op	   24533 B/op	     278 allocs/op
BenchmarkSelectGORM_10-8     	      50	  22354574 ns/op	   91347 B/op	    1809 allocs/op
BenchmarkSelectGORM_100-8    	       5	 202633672 ns/op	  758755 B/op	   17117 allocs/op
BenchmarkSelectGORM_1000-8   	       1	2188011084 ns/op	 7426904 B/op	  171054 allocs/op
```

A benchmark run is named with the following convention:

```text
Benchmark[SQL operator][Package name]_[Number of SQL calls]
```

* SQL operator = SQL operator that is used in the benchmark run. Can be one of the following: Insert, Update, and Select.
* Package name = Package used in the benchmark run. Can be one of the following: SQL, XORM, or GORM
* Number of SQL calls = Number of times the SQL operator is called in a single benchmark operation. Note this number is different from the benchmark operations ie. `b.N`. For I/O bound benchmark, it's often preferable to measure the atomic calls in large sample due to the overhead involved in setting a connection to the database or device, ie. 1 benchmark operation with a large sample of atomic calls. Indeed there are larger variation in the result for batch sizes of 1 and 10. So the result below are all based on the batch size of 1,000.

![Benchmark: Time taken to complete a call to the database](images/benchmark-speed.svg)

![Benchmark: Memory used to complete a call to the database](images/benchmark-memory.svg)

![Benchmark: Number of allocations made to complete a call to the database](images/benchmark-allocations.svg)

### Take-Aways

* The ORM packages are less memory efficient than the standard package `sql`.
* XORM performs surprisingly well in speed, almost as fast as the standard package `sql`.
* Speed performance for GORM isn't as good as the other 2 packages for insert and update. For select the 3 packages perform equally well.

## Caveats and Notes

Running benchmark takes patience and tweaking so that we can get the right dataset to make a good conclusion of the results. The above charts are just the initial benchmark results, I will make incremental improvements to the code.

* Ideally we should isolate and run the benchmarks separately. Overhead from another benchmark may affect another benchmark. In the initial runs, there were definitely some spillover from previous runs. Using the `ResetTimer()` function seems to help and so far the benchmark runs have been quite consistent. May run a benchmark individually in the future.
* No transactions.
* No joins.

## Reference and Credits

* [Golang: Package testing](https://golang.org/pkg/testing/)
* [Golang: Package database/sql](https://golang.org/pkg/database/sql/)
* [XORM](https://xorm.io/)
* [GORM](https://gorm.io/)
