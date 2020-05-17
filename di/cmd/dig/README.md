# Dependency Injection Using Dig

This recipe is a modification of the [no-framework recipe](../no-framework). Here are the changes if you are doing this manually from scratch.

* No need to touch the `pkg` package.
* Add the `dig` package into `main`.
* Add a new file [container.go](container.go). Use this file to define the dependencies and let dig figure the dependency graph and perform dependency injection at runtime using Go reflection.

Compare to the original code, the code is cleaner and more decoupled.
