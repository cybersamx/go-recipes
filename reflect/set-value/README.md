# Set a Slice to Struct Field

Let's suppose we have a slice or a struct with a field of type slice, and we want to assign to the slice with
another slice or append with an item. This example shows how we can do that with Go's `reflect` package.

> **Note**
>
> Reflection is tricky business as we attempt to find out the type metadata of an object at runtime. So we can't fully rely on static types and the compiler to do the checking for us. There are many edge cases that are not captured in this example as we want to focus on the pertinent code on reflection. Please don't just copy the example without understanding what the code and include additional checks and logic for handling edge cases.

## Setting a Variable

When we are setting a `reflect.Value` object (or rval) with a value, the rval must be addressable. This means
that we should be initializing a rval with a pointer and not with a value, so that we can mutate the actual
referenced value and not the copied value.

```go
var n int

// The following will not work.
// Panic: reflect.Value.Set using unaddressable value
namesVal := reflect.ValueOf(n)
namesVal.Set(reflect.ValueOf(4))

// The following will work:
// We need get a reflect.Value of the pointer to the variable n and then expose the
// actual reflect.Value using the method Elem().
namesVal := reflect.ValueOf(&n).Elem()
namesVal.Set(reflect.ValueOf(4))
```

## Setting the Struct Field

What if we want to set a field in a struct?

```go
type Person struct {
	Age int
}

person := Person{
	Age: 15,
}
```

We first need to get a `reflect.Value` of the Person pointer so that we can actually set the actual Person object and
not the copy. Through the `reflect` package, we use the method `FieldByName()` to retrieve the `reflect.Value` of
the `Age` field. Finally, we  call `Addr()` on the field rval object to get its address. In other words:

```go
fieldVal := reflect.ValueOf(&person).Elem().FieldByName("Age").Addr().Elem()
fieldVal.Set(reflect.ValueOf(30))
fmt.Println(person) // It should print 30
```

Note that we have to call `Elem()` twice in this context. Once on the struct itself since we are using a pointer
to the struct. Second, on the field itself, since we called `Addr()`.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference

* [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
