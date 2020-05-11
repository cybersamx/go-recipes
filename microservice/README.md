## Microservice

This is a collection of recipes for implementing microservices with Go-Kit.

Given that microservice is a big and complex subject, we don't want to just cramp everything into a single project. Instead, we break the microservice recipe into smaller recipes starting with the recipe called `simple`. We gradually build upon each recipe with more scope and advanced constructs leading to the final recipe. Here's the order of the recipes to consume:

* [Simple](simple) - a simple recipe of how to use Go-Kit to implement a REST microservice or monolith service.
* Middleware - we add logging and tracing to the previous recipe.

## Why Go-Kit?

Go-Kit is a platform-agnostic framework of libraries and programming model that enable you to develop scalable, extensive, and reliable microservices. Go-Kit doesn't attempt to lock you in an infrastructure or platform. Its libraries enable you to easily integrate into the hosting infrastructure and platform dependencies.

Go-Kit uses a layering architecture for the development of its microservices, which not only enable a pragmatic abstraction of the codebase, but also provide an flexible way of extending the microservices. These are the 3 layers that are integral to a Go-Kit microservice programming model:

* Service layer 
* Endpoint layer
* Transport layer

# References

* [Go-Kit](https://gokit.io/)
