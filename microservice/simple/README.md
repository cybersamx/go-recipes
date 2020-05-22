# Simple Monolithic REST Service

This is a simple REST service implemented with [Go-kit](https://github.com/go-kit/kit).

## Setup

1. Run the server

   ```bash
   $ make run
   ```

2. To test the server, run:

   ```bash
   $ curl http://localhost:8000/states/
   $ curl http://localhost:8000/states/ca
   ```

## Data

The data in `states.yaml` was taken from US Census - see <https://www.census.gov/data/datasets.html>.
