# Custom Type as a Type Constraint

An example of using a custom type as a type constraint in Go generics.

## Basics

With Go support for generics, we no longer need to declare `interface{}` as a type for passing a value to a function. We can now use type parameter and type constraint to define the set of types the generic function or type would accept.

In this example, we "constrain" the parameter type T to the interface `Animal`. So any types that implement the behaviors ie. `Talk` and `Species` functions can be passed to the generic function `Print` which accepts the type  `Animal`.

## Reference

* [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
