# Custom Marshaling

A simple recipe for custom JSON marshaling and unmarshaling in Go.

## Use Case

We use the following struct that represent a user account that is used in a real-world API.

```go
type User struct {
	ID        string
	Email     string
	Password  string
}
```

Let's suppose the application accepts a JSON object with the field password for authentication but doesn't return a `User` object with the password field. You can't really do that with tags `-` or `omitempty`. The former drops the tagged field from during both unmarshaling and marshaling. While the `omitempty` tag only omits the field when it's empty during marshaling.

The workaround is to "implement" the `Marshaler` interface by writing the `MarshalJSON()`.

There are a few solutions to this problem. They all involve "implementing" the `MarshalJSON()` in `Marshaler` interface.

* [Omit password field by reflection solution](user_reflect.go) - Use the `reflect` package to transform the User struct into a `map[string]interface{}`, dropping the password field completely during transformation before finally passing the resultant `map` to json.Marshal() for seralization. This is the most tedious solution but also the cleanest as the `password` field is completely omitted.
* [Custom type for the password field solution](user_custom_type.go) - Define a new type called `SensitiveString` to represent password. This new type implements the `MarshalJSON()` function that always return a empty `[]byte` slice, hence disabling the serialization of the string value to JSON. However, it doesn't drop the field entirely. You will still see `password=""` in the json object.
* [Type alias and selective field copy solution](user_alias_type.go) - Create a type alias inside the `MarshalJSON()` function, copy all fields (except `Password`) in the `User` object over to an object of thew new alias User type before serializing the alias User object to json. You will still see `password=""` in the json object.

In the recipe, we also define additional fields `Note` and `Tags` to validate that json option tags `-` and `omitempty` are correctly parsed by our custom reflection code.

## Setup

1. Test the program

   ```bash
   $ make
   ```

## Reference

* [Godoc: json](https://godoc.org/encoding/json)
