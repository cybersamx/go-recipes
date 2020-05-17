# Dependency Injection Using Wire

This recipe is a modification of the [no-framework recipe](../no-framework). Here are the changes if you are doing this manually from scratch.

* No need to touch the `pkg` package.
* Add the `wire` package to `main`.
* Download and install the `wire` runtime tool that is needed to generate the code. ie. `go get github.com/google/wire/cmd/wire`.
* Add a file called [wire.go](wire.go) that wires all the dependencies together using the `wire.Build()` function. Running the `wire` command in the shell will generate code [wire_gen.go](wire_gen.go) from [wire.go](wire.go).

Compare to the original code, the code is cleaner and more decoupled.

## Setup

> **NOTE:**
> If you run this the first time, you may not see a response from the program for a few seconds. The program is performing a static analysis of your code to generate the source code.
