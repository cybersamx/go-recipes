# Dependency Injection Using Dig

This example uses Dig as the DI framework. It is based on [no-framework recipe](../no-framework) with the following
differences.

* Add the `dig` package into `main`.
* Add a new file [container.go](container.go). Use this file to define the dependencies and let dig figure the dependency graph and perform dependency injection at runtime using Go reflection.
