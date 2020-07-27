## Microservice

This is a collection of recipes for implementing microservices with Go-Kit.

* [Simple](simple) - a simple recipe of how to use Go-Kit to implement a REST-based microservice.

## Why Go-Kit?

Go-Kit is a platform-agnostic framework of libraries and programming model that enable you to develop scalable, extensive, and reliable microservices. Go-Kit doesn't attempt to lock you in an infrastructure or platform. Its libraries enable you to easily integrate into the hosting infrastructure and platform dependencies.

Go-Kit uses a layering architecture for the development of its microservices, which not only enable a pragmatic abstraction of the codebase, but also provide an flexible way of extending the microservices. These are the 3 layers that are integral to a Go-Kit microservice programming model:

* Service layer
* Endpoint layer
* Transport layer

## Why Not Go-Kit?

My thoughts on Go-Kit have evolved recently. Go-Kit is a comprehensive framework in building services and microservices. It does a good job in abstracting the different parts of the project from each other, but at the expense of requiring the developer to adhere to the framework's strict programming model. It's a bit heavy and require you to implement and wrap a lot of code for almost every aspect in a service or microservice project. This can be restrictive; often leading to overhead (with much boilerplate code) and adding unnecessary complexity to the project for certain use cases.

If we are going to implement a simple REST service, just use one of the popular third-party packages out there, like `gin` or `gorilla`. If we are going to implement microservices, just use GRPC. In either case, you can still apply good design patterns and best practices without resorting to a framework.

# References

* [Go-Kit](https://gokit.io/)
