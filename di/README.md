# Dependency Injection in Go

The underlying application is a simple implementation of a database-driven API server. We are using dependency injection (DI) to load and initialize dependencies when starting the application.

The 3 recipes explore the different ways of DI:

* [No DI framework](no-framework) - Basic DI implemented using Go, no third-party DI framework.
* [Using the Dig framework](dig) - how to leverage the Dig DI framework that injects dependencies at runtime. See [the recipe](cmd/dig) for details.
* [Using the Wire framework](wire) - how to leverage the Wire DI framework that generate code that injects dependencies at compile time. See [the recipe](cmd/wire) for details.

All 3 recipes use the shared [pkg](pkg) package containing all the pertinent dependencies needed to run the application.

* `models.go` - The main data model `City` and its corresponding Factory  and  support functions.
* `datastore.go` - The persistence layer representing the data store. It contains the `DataStore` and its corresponding Factory and data access functions.
* `server.go` - The server that serves an RESTful endpoint over HTTP. It contains the `HTTPServer` and its corresponding Factory and route handler functions.
* `settings.go` - Settings the runtime settings that we want to pass to the above components during construction and initialization at runtime.

## Setup

1. Run the program

   ```bash
   $ make run-no-framework
   $ make run-dig
   $ make run-wire
   ```

2. On another shell

   ```bash
   $ curl localhost:8000
   ```

# Reference

* [GoDoc: Dig](https://godoc.org/go.uber.org/dig)
* [Github: Dig](https://github.com/uber-go/dig)
* [Github: Wire](https://github.com/google/wire)
* [Dependency Injection in Go](https://blog.drewolson.org/dependency-injection-in-go)
* [Go Dependency Injection with Wire](https://blog.drewolson.org/go-dependency-injection-with-wire)
