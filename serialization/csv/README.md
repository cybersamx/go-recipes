# Basic CSV Parsing

A simple recipe for reading and parsing a CSV file.

The `encoding/csv` package in the standard Go library provides the parser you need to parse CSV. Use the function `Read()` to read and parse a record from an instance of `Reader`. Make sure you handle the returned error for `io.EOF`, which in my opinon isn't really an error but a state condition. 

## Setup

1. Run the program

   ```bash
   $ go run main.go
   ```
   
## Reference and Credits

* [LA Bus Stop Benches](http://geohub.lacity.org/datasets/bus-stop-benches) - This recipe uses a partial public dataset from the City of LA. 
* [Godoc: csv](https://godoc.org/encoding/csv)
