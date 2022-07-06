# Type Constraints

We use type constraints to set narrow the set of types for a generic function or types. This helps to define what set of types and constraints the associated type parameters should accept.

| Type Constraint     | ==  | !=  | <   | <=  | \>  | \>= | +   |
|---------------------|-----|-----|-----|-----|-----|-----|-----|
| any                 |     |     |     |     |     |     |     |
| comparable          | ✓   | ✓   |     |     |     |     |     |
| constraints.Ordered | ✓   | ✓   | ✓   | ✓   | ✓   | ✓   | ✓   |

## Type Approximation

We sometimes define a type like this:

```go
type MyInteger int
```

Furthermore, enum in Go relies on type definition. For example:

```go
type MyEnum int

const (
	MyEnumA MyEnum = iota
	MyEnumB
)
```

It's not going to work if our generic function is declared as follows:

```go
func foo[T int | int32 | int64](t T) {
	// ....
}
```

Use type approximation by prefixing the base type with ~, so that `MyInteger` and `MyEnum` will be accepted by the generic function.

```go
func foo[T ~int | ~int32 | ~int64](t T) {
	// ....
}
```

## Reference

* [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
* [Bitfield Consulting: Generics in Go](https://bitfieldconsulting.com/golang/generics)
* [GopherCon 2021: Generics by Robert Griesemer and Ian Taylor](https://www.youtube.com/watch?v=Pa_e9EeCdy8&t=567s)
