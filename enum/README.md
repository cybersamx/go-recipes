# Enum

An example showing the implementation of enum.

## Basics

This example includes a method for returning the enum value as string. In the `String` method, we create an array of
string values and use the enum value, which is an integer type to index the actual string value.

See [Stackoverflow: how to make go print enum fields as string](https://stackoverflow.com/questions/41480543/how-to-make-go-print-enum-fields-as-string/62291060#62291060) for details. You can also implement the method using `switch` to return the string values per Rishabh's comment.

## Using Stringer

We can use also use a codegen tool `stringer` to generate the method `String()` of an enum type definition for us. See the type `MyEnumValImplict`. We need to do the following:

```shell
go install golang.org/x/tools/cmd/stringer@latest
go generate
go run main.go myenumimplicit_string.go
```

Or simply run `make run`.

## References

* [Stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
