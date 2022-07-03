# Type Parameters

## Type Parameters in Functions

A Go function can be declared as follows:

```go
func FunctionName(arg1 Type, arg2 Type, ...) {}
```

With Go generics, type parameters can now be applied to function declaration. Type parameters are encapsulated in the square brackets `[...]` and before the function parameters.

```go
func FunctionName[T1 TypeConstraint, T2 TypeContraint, ...](arg1 T, arg2 T,...) {}
```

## Type Parameters in Types

With Go 1.18, types can now be declared with type parameters.

```go
type StructName[T TypeConstraint, ...] struct {
	attribute1 T
}
```

## Reference

* [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
