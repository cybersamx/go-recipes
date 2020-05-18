# Ephemeral data models in SQL

A recipe that shows the implementation of an ephemeral data model in Go. We create a SQL data store that will garbage collect every n seconds to check and remove any data objects in the store that have expired.

The recipe uses the time.Ticker to garbage collect every GCInterval.

## Setup

1. Run the program

   ```bash
   $ make
   ```
