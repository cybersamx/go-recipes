# Ephemeral data models in SQL

An example of an ephemeral data model in SQL. The data store will garbage collect every n seconds to check and remove which data objects have expired.

The recipe uses the time.Ticker to garbage collect every GCInterval.

## Setup

1. Run the program

   ```bash
   $ go run *.go
   ```
   