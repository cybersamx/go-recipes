# Dependency Injection in Go Using Wire

Go DI using the Wire framework.

## Overview

Here is a simple implementation of a database-driven API server. The [pkg](pkg) directory contains the typical components that support such application. Here's a breakdown:

* `models.go` - The main data model `City` and its corresponding Factory  and  support functions. 
* `datastore.go` - The persistence layer representing the data store. It contains the `DataStore` and its corresponding Factory and data access functions.
* `server.go` - The server that serves an RESTful endpoint over HTTP. It contains the `HTTPServer` and its corresponding Factory and route handler functions. 
* `settings.go` - Settings the runtime settings that we want to pass to the above components during construction and initialization at runtime.

Use [no-framework sample code](../no-framework) as the point of reference. For this recipe we simply drop the Wire package into the project. The source code for the above has not changed at all.

We added a file called [wire.go](pkg/wire.go) that wires all the dependencies together using the `wire.Build()` function. This will generate a [wire_gen.go](pkg/wire_gen.go). Modify [main.go](main.go) accordingly. Those are the changes needed.

Download the `wire` binary and run it to generate construction functions using DI. You just need to run `wire` the first time. For any subsequent changes you made to your source code, just run `go generate` to regenerate the DI source code. For more details, read the official [user guide](https://github.com/google/wire/blob/master/docs/guide.md), [tutorial](https://github.com/google/wire/blob/master/_tutorial/README.md), and [best practices doc](https://github.com/google/wire/blob/master/docs/best-practices.md).

Now the code is cleaner, has less code, and decoupled.  

## Setup

1. Install wire. **NOTE:** if you run this the first, you may not see a response from the program for a few seconds. The program is performing a static analysis of your code to generate the source code. I wish the author would change this future releases.

   ```bash
   $ go get github.com/google/wire/cmd/wire
   $ wire
   $ # It may take a few seconds.
   ```

1. Run the program

   ```bash
   $ go run main.go
   ```
   
2. On another shell

   ```bash
   $ curl localhost:8000
   ```
   
## Reference and Credits

* [Github: Wire](https://github.com/google/wire)
* [Go Dependency Injection with Wire](https://blog.drewolson.org/go-dependency-injection-with-wire)
