# Collection of Golang Recipes

[![Go Report Card](https://goreportcard.com/badge/github.com/cybersamx/go-recipes)](https://goreportcard.com/report/github.com/cybersamx/go-recipes)

I started this project as I was learning Go. Over time, it evolved into a handy reference for me to look up Go recipes and design patterns.

Here's a collection of the recipes:

* [Context](context) - recipes on the `context` package.
* [Dependency injection](dependency-injection) - recipes that shows different ways to implementing dependency injection design pattern using vanilla Go and with open source frameworks dig and wire.
* [Ephemeral SQL data model](ephemeral-sql-data) - a recipe implementing a garbage collector that removes expired records in an SQL database in the background.
* [Fake vs mock in unit testing](fake-mock) - a recipe that shows the use of fake and mock in a Go unit test.
* [File reader](file-reader) - recipes that show the different ways of reading file: incremental and full.
* [gRPC](grpc) - a simple gRPC server-client setup.
* [Kafka pubsub](kafka-pubsub) - a recipe that implements a simple pubsub setup in Kakfa using the watermill framework.
* [Long polling](long-poll) - a simple long-polling implementation.
* [Microservice](microservice) - recipes that implement microservice using the go-kit framework.
* [MongoDB](mongo) - recipes for working with MongoDB using the official driver.
* [Postgres](postgres) - recipes for working with PostgreSQL using 3 popular frameworks sql, xorm, and gorm.
* [Retry](retry) - a recipe that implements retry, which can be useful for connecting to a service in the network reliably.
* [Serialization](serialization) - recipes on serialization.
* [Web cookie](cookie) - a recipe for working with web cookies.
* [Web form](web-form) - a recipe explores working with the rendering of a simple web form using html/template and handling of form submissions.
* [Webhook](webhook) - a simple webhook implemenation.

Other somewhat Go-related recipes:

* [Go Docker images](long-poll) - the long-poll recipe also contain an extensive example of building Go application Docker images using common base Docker images ubuntu, alpine, scratch, and distroless.
