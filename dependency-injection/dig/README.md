# Dependency Injection in Go Using Dig

Basic DI using the Dig framework.

## Overview

Here is a simple implementation of a database-driven API server. The [pkg](pkg) directory contains the typical components that support such application. Here's a breakdown:

* `models.go` - The main data model `City` and its corresponding Factory  and  support functions. 
* `datastore.go` - The persistence layer representing the data store. It contains the `DataStore` and its corresponding Factory and data access functions.
* `server.go` - The server that serves an RESTful endpoint over HTTP. It contains the `HTTPServer` and its corresponding Factory and route handler functions. 
* `settings.go` - Settings the runtime settings that we want to pass to the above components during construction and initialization at runtime.

Use [no-framework sample code](../no-framework) as the point of reference. For this recipe we simply drop the Dig package into the project. The source code for the above has not changed at all.

We added a new file [container.go](pkg/container.go) to define the dependencies and let Dig figure the dependency graph and perform dependency injection at runtime using Go reflection. Also, we reduce the code in [main.go](main.go).

Now the code is cleaner, has less code, and decoupled.  

## Setup

1. Run the program

   ```bash
   $ go run main.go
   ```
   
2. On another shell

   ```bash
   $ curl localhost:8000
   ```
   
## Reference and Credits

* [GoDoc: Dig](https://godoc.org/go.uber.org/dig)
* [Github: Dig](https://github.com/uber-go/dig)
* [Dependency Injection in Go](https://blog.drewolson.org/dependency-injection-in-go)
