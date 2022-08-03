# Dependency Injection in Go

This is a simple RESTful API service. In the code, we are using dependency injection (DI) to load and
initialize dependencies when starting the application.

3 ways to inject dependencies:

* [No DI framework](no-framework) - Basic DI implemented using Go, no third-party DI framework.
* [The Dig framework](dig) - Use the Dig framework to inject dependencies at runtime. See [the example](cmd/dig)
  for details.
* [The Wire framework](wire) - Use the Wire framework to generate the DI code at compile time. See [the example]
  (cmd/wire) for details.

## Setup

1. Install the third party packages.

   ```bash
   $ make install
   ```

1. You can run of the following targets corresponding to a DI framework.

   ```bash
   $ make run-no-framework
   $ make run-dig
   $ make run-wire
   ```

1. On another shell.

   ```bash
   $ curl localhost:8000
   ```

# Reference

* [GoDoc: Dig](https://godoc.org/go.uber.org/dig)
* [Github: Dig](https://github.com/uber-go/dig)
* [Github: Wire](https://github.com/google/wire)
* [Dependency Injection in Go](https://blog.drewolson.org/dependency-injection-in-go)
* [Go Dependency Injection with Wire](https://blog.drewolson.org/go-dependency-injection-with-wire)
