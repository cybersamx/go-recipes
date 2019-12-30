# Basic Dependency Injection in Go

Basic DI implemented using Go, no third-party DI framework.

## Overview

Here is a simple implementation of a database-driven API server. The [pkg](pkg) directory contains the typical components that support such application. Here's a breakdown:

* `models.go` - The main data model `City` and its corresponding Factory  and  support functions. 
* `datastore.go` - The persistence layer representing the data store. It contains the `DataStore` and its corresponding Factory and data access functions.
* `server.go` - The server that serves an RESTful endpoint over HTTP. It contains the `HTTPServer` and its corresponding Factory and route handler functions. 
* `settings.go` - Settings the runtime settings that we want to pass to the above components during construction and initialization at runtime.

## Setup

1. Run the program

   ```bash
   $ go run main.go
   ```
   
2. On another shell

   ```bash
   $ curl localhost:8000
   ```
