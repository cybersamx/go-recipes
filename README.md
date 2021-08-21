# Collection of Golang Recipes

[![Go Report Card](https://goreportcard.com/badge/github.com/cybersamx/go-recipes)](https://goreportcard.com/report/github.com/cybersamx/go-recipes)

I started this project as I was learning Go. Over time, it evolved into a handy reference for anyone to look up Go recipes and design patterns.

## Guides

* [Guide](docs/guide.md) - handy guide to Go - the common knowledge, tips, gotchas, etc.
* [Go Module](docs/module.md) - Go module.
* [Go Concurrency](docs/concurrency.md) - Go concurrency.

Go-related guides and docs:

* [Goenv](docs/goenv.md) - guide to using goenv, a version manager for Go.
* [Unicode](docs/unicode.md) - a primer on Unicode.
* [Go time format](docs/time-format.md) - intro to go time format.

## Recipes

Here's a collection of the recipes:

* [Collection](collection) - recipes on Go collections.
  * [Array Passing](collection/array-passing) - recipe that highlights how an array is passed as a value to a function and what we can do to modify the original array.
  * [Map](collection/map) - basic map operations.
  * [Slice](collection/slice) - basic slice operations.
* [Concurrency](concurrency) - Go concurrency.
  * [Channel](concurrency/channel)
    * [Basic Channel](concurrency/channel/basic) - Basic channel example.
    * [Buffered Channel](concurrency/channel/buffered) - Buffered channel example.
  * [WaitGroup](concurrency/waitgroup) - a recipe that demonstrates the use of `sync.WaitGroup` to help synchronize the execution of multiple concurrent goroutines.
* [Configurations](config) - load configurations into an app using a load order of default values > config file > environment variables > CLI arguments.
* [Cookie](cookie) - a recipe for working with http cookies. Includes a write-up on the architecture of http cookie and its key attributes.
* [Context](context) - recipes on `context` package.
  * [Deadline](context/deadline) - timeout using `context.WithDeadline` function.
  * [Timeout](context/timeout) - timeout using `context.WithTimeout` function.
* [Dependency injection](di) - recipes that shows different ways to implementing dependency injection design pattern.
  * [No DI framework](no-framework) - Basic DI implemented using Go, no third-party DI framework.
  * [Dig framework](dig) - how to leverage the Dig DI framework that injects dependencies at runtime.
  * [Wire framework](wire) - how to leverage the Wire DI framework that generate code that injects dependencies at compile time.
* [Ephemeral SQL data model](ephemeral-sql-data) - a recipe implementing a garbage collector that removes expired records in an SQL database in the background.
* [Fake vs mock in unit testing](fake-mock) - a recipe that shows the use of fake and mock using the `golang/mock` (aka Gomock) and `testify/mock` packages in a Go unit test.
* [File I/O](io) - recipes on file input/output operations.
  * [Reader](reader) - a recipe that shows the different ways of reading file: incrementally and all-at-once.
  * [Writer](writer) - a recipe for writing a file
* [GraphQL](graphql) - recipes for using GraphQL in Go.
  * [GraphQL client](graphql/client) -a simple GraphQL client using the machinebox/graphql package.
  * [GraphQL server](graphql/server) - a simple GraphQL server using the 99designs/gqlgen package/codegen.
* [gRPC](grpc) - recipes on gRPC.
  * [Hello world](grpc/hello-world) - a simple gRPC server-client setup.
  * [gRPC test](grpc/test) - a recipe that implements unit tests for a gRPC server.
* [HTTP](http) - recipes for using the web.
  * [Client](client) - a recipe that implements a simple HTTP client that sends a GET request.
  * Server
    * [Advanced template](http/server/advanced-template) - a web application using Go templating and Material Design Components for the real world.
    * [JWT](http/server/jwt) - a recipe that shows the handling of simple web form submission and a simple JWT-based authentication using the `dgrijalva/jwt-go` package.
    * [HTML template](http/server/html-template) - a recipe that shows server-side http content rendering using the `html/template` package.
    * [Static web server](http/server/static) - a recipe that implements a simple http server serving static content.
* [Kafka pubsub](kafka-pubsub) - a recipe that implements a simple pubsub setup in Kakfa using the `ThreeDotsLabs/watermill` package.
* [Long polling](long-poll) - a simple long-polling implementation.
* [Microservice](microservice) - recipes that implement microservice using the `go-kit/kit` package.
* [MongoDB](mongo) - recipes for working with MongoDB using the official driver.
  * [Basic](mongo/basic) - basic operations with Mongo.
  * [Expiring Data](mongo/expiring-data) - remove expired data using MongoDB TTL indexing.
  * [Mongo Operations](mongo-ops) - Recipe on running Mongo commands and other queries.
  * [Pre-defined schema](mongo/schema) - a recipe similar to the [basic mongo recipe](mongo/basic), but it's implemented using a pre-defined schema.
* [Postgres](postgres) - recipes on postgres and general SQL operations.
  * [ORM](postgres/orm-sql) - recipes for working with PostgreSQL using 3 popular frameworks: `sql`, `go-xorm/xorm`, and `jinzhu/gorm` packages.
* [Pulsar pubsub](pulsar-pubsub) - a recipe that implements a simple pubsub in Apache Pulsar.
* [Redis](redis) - recipes for working with Redis using the `mediocregopher/radix` and `go-redis/redis` drivers.
  * [Counter](redis/counter) - a recipe that implements a global atomic counter that showcases basic operations in Redis.
  * [Sessions](redis/sessions) - a recipe that implements the serialization of complex data to redis using GOB and ephemeral sessions in Redis by setting up Redis to remove expired content.
  * [Authentication](auth) - authentication in Redis.
* [Reflection](reflect) - Go runtime reflection.
  * [Deep Equal](reflect/equality) - a recipe on Go (deep) equality operation using the `reflect.DeepEqual` to compare 2 values.
  * [Print Fields](reflect/print-fields) - Given a value of a custom `struct` type, this recipe prints out the field names and values.
  * [Type check](reflect/typecheck) - a recipe that highlights the different ways to check the type of a value at runtime.
* [Retry](retry) - a recipe that implements retry, which can be useful for a reliable connection to a service in the network.
  * [No-package](retry/no-package) - a recipe using retry without a third-party framework.
  * [Retry](retry/retry) - a recipe using the `flowchartsman/retry` package.
* [Scheduler](scheduler) - a recipe that highlights ways to schedule code to be run periodically.
* [Serialization](serialization) - a recipes on serialization.
  * [CSV](serialization/csv) - a recipe for csv serialization.
  * [JSON](serialization/json) - a recipe for json serialization.
* [String](string) - strings and characters in Go.
  * [Count Iterate](string/count-iterate) - simple recipe that explores how we count and iterate over a string by rune and byte.
  * [Regular Expression](string/regexp) - a recipe on using regular expression in Go.
  * [Unicode](string/unicode) - basic unicode representation and usage.
* [Time](time) - recipes for manipulating time in Go.
  * [Parse](time/parse) - a recipe that parses a date string to a Go struct value.
  * [Print](time/print) - a recipe that prints a date value to different formats using layouts.
* [WebAssembly](wasm) - WebAssembly recipes.
  * [Simple](wasm/simple) - a simple WebAssembly that prints "hello world" on the browser console.
* [Web Scaper](web-scraper) - a recipe that extracts the price of a stock from MarketWatch.com using the `gocolly/colly` package.
* [Webhook](webhook) - a simple webhook implementation.

Other somewhat Go-related recipes:

* [Go Docker images](long-poll) - the long-poll recipe also contain an extensive example of building Go application Docker images using various base Docker images ubuntu, alpine, scratch, and distroless.

## Setup

The [Makefile](Makefile) in the project root directory calls the Makefile each recipe. To build all recipes run:

```bash
$ make build
```

To remove the built binaries and build cache, run:

```bash
$ make clean
```

## Credits and Reference

* [Azer Makefile](https://github.com/azer/go-makefile-example/blob/master/Makefile)
