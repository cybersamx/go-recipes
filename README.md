# Collection of Golang Recipes

[![Go Report Card](https://goreportcard.com/badge/github.com/cybersamx/go-recipes)](https://goreportcard.com/report/github.com/cybersamx/go-recipes)

I started this project as I was learning Go. Over time, it evolved into a handy reference for anyone to look up Go recipes and design patterns.

## Guides

* [Guide](docs/guide.md) - handy guide to Go - the common knowledge, tips, gotchas, etc.
* [Go Module](docs/module.md) - Go module

Go-related guides and docs:

* [Goenv](docs/goenv.md) - guide to using goenv, a version manager for Go.
* [Unicode](docs/unicode.md) - a primer on Unicode.

## Recipes

Here's a collection of the recipes:

* [Context](context) - recipes on `context` package.
  * [Timeout](context/timeout) - timeout using `context` package.
* [Collection](collection) - recipes on Go collections.
  * [Array Passing](collection/array-passing) - recipe that highlights that array is passed as a value to a function and what we can do to modify the original array.
  * [Map](collection/map) - basic map operations.
  * [Slice](collection/slice) - basic slice operations.
* [Cookie](cookie) - a recipe for working with http cookies. Includes a write-up on the architecture of http cookie and its key attributes.
* [Dependency injection](di) - recipes that shows different ways to implementing dependency injection design pattern using vanilla Go and with open source frameworks dig and wire.
* [Ephemeral SQL data model](ephemeral-sql-data) - a recipe implementing a garbage collector that removes expired records in an SQL database in the background.
* [Fake vs mock in unit testing](fake-mock) - a recipe that shows the use of fake and mock using the `golang/mock` (aka Gomock) and `testify/mock` packages in a Go unit test.
* [File I/O](io) - recipes on file input/output operations.
  * [Reader](reader) - a recipe that shows the different ways of reading file: incrementally and all-at-once.
  * [Writer](writer) - a recipe for writing a file
* [GraphQL](graphql) - recipes for using GraphQL in Go.
  * [GraphQL client](graphql/client) -a simple GraphQL client using the machinebox/graphql package.
  * [GraphQL server](graphql/server) - a simple GraphQL server using the 99designs/gqlgen package/codegen.
* [gRPC](grpc) - a simple gRPC server-client setup.
* [Kafka pubsub](kafka-pubsub) - a recipe that implements a simple pubsub setup in Kakfa using the `ThreeDotsLabs/watermill` package.
* [Long polling](long-poll) - a simple long-polling implementation.
* [Microservice](microservice) - recipes that implement microservice using the `go-kit/kit` package.
* [MongoDB](mongo) - recipes for working with MongoDB using the official driver.
  * [Basic](mongo/basic) - basic operations with Mongo.
  * [Pre-defined schema](mongo/schema) - a recipe similar to the [basic mongo recipe](mongo/basic), but it's implemented using a pre-defined schema.
  * [Changestream](mongo/change-stream) - recipe using Mongo Changestream.
* [Postgres](postgres) - recipes on postgres and general SQL operations.
  * [ORM](postgres/orm-sql) - recipes for working with PostgreSQL using 3 popular frameworks: `sql`, `go-xorm/xorm`, and `jinzhu/gorm` packages.
* [Redis](redis) - recipes for working with Redis using the `mediocregopher/radix` driver.
  * [Counter](redis/counter) - a recipe that implements global atomic counter that showcases basic operations in Redis.
  * [Sessions](redis/sessions) - a recipe that implements ephemeral sessions in Redis by setting up Redis to remove expired content.
  * [Authentication](auth) - authentication in Redis.
* [Reflection](reflect) - Go runtime reflection.
  * [Deep Equal](equality) - recipe on Go (deep) equality operation using the `reflect.DeepEqual` to compare 2 values.
  * [Type check](typecheck) - recipe that explores ways to check the type of value at runtime.
* [Retry](retry) - a recipe that implements retry, which can be useful for connecting to a service in the network reliably.
  * [No-package](retry/no-package) - recipe using retry without a third-party framework.
  * [Retry](retry/retry) - recipe using the `flowchartsman/retry` package.
* [Scheduler](scheduler) - a recipe that highlights ways to schedule code to be run periodically.
* [Serialization](serialization) - recipes on serialization.
  * [CSV](serialization/csv) - recipe for csv serialization.
  * [JSON](serialization/json) - recipe for json serialization.
* [String](string) - strings and characters in Go.
  * [Count Iterate](string/count-iterate) - simple recipe that explores how we count and iterate over a string by rune and byte.
  * [Unicode](string/unicode) - basic unicode representation and usage.
* [Web](web) - recipes for using the web.
  * [JWT](web/jwt) - a recipe that shows the handling of simple web form submission and a simple JWT-based authentication using the `dgrijalva/jwt-go` package.
  * [HTML template](web/html-template) - a recipe that shows server-side web content rendering using the `html/template` package.
* [Webhook](webhook) - a simple webhook implementation.

Other somewhat Go-related recipes:

* [Go Docker images](long-poll) - the long-poll recipe also contain an extensive example of building Go application Docker images using various base Docker images ubuntu, alpine, scratch, and distroless.

## Credits and Reference

* [Azer Makefile](https://github.com/azer/go-makefile-example/blob/master/Makefile)
