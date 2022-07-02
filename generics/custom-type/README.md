# Custom Type in Go Generics

A recipe on using a custom type in Go generics.

## Basics

In this example, we "constrain" the generic type T to the interface `Animal`. So any types that implement the behaviors ie. `Talk` and `Species` functions can be passed to the function `Print` which accepts the T type `Animal`.

With Go support for generics, we no longer need to declare `interface{}` as a type for passing a value to a function.

## Reference

* [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
