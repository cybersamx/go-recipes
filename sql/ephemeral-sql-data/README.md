# Ephemeral data models in SQL

An example showing the implementation of an ephemeral data model in Go. We create a SQL data store that will garbage collect every n seconds to check and remove any expired data objects in the store.

The recipe uses the time.Ticker to garbage collect every GCInterval.

## Setup

1. Run the program.

   ```bash
   $ make run
   ```
