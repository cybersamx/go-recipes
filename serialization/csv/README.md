# Basic CSV Parsing

This example covers:

* Reading from a csv file.
* Writing to a csv file.

The `encoding/csv` package in the standard Go library provides the parser you need to parse CSV. Use the function `Read()` to read and parse a record from an instance of `Reader`. Make sure you handle the returned state for `io.EOF`.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference and Credits

* [LA Bus Stop Benches](http://geohub.lacity.org/datasets/bus-stop-benches) - This recipe uses a partial public dataset from the City of LA.
* [Godoc: csv](https://godoc.org/encoding/csv)
