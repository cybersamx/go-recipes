# Dependency Injection Using Wire

This example uses Dig as the DI framework. It is based on [no-framework recipe](../no-framework) with the following
differences.

* Add the `wire` package to `main`.
* Download and install the `wire` runtime tool that is needed to generate the code. ie. `go get github.com/google/wire/cmd/wire`.
* Add a file called [wire.go](wire.go) that wires all the dependencies together using the `wire.Build()` function. Running the `wire` command in the shell will generate code [wire_gen.go](wire_gen.go) from [wire.go](wire.go).

## Setup

> **NOTE:**
> If you run this the first time, you may not see a response from the program for a few seconds. The program is performing a static analysis of your code to generate the source code.
