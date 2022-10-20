# Merge Matching Fields Recursively from Map to Struct

Suppose we have a map[string]any. We want to copy the content of that map to a struct if the keys in the map matches the tags in the struct.

This example is an extension of the [simpler example](../merge-fields) of just mapping non-nested fields in a struct. This example traverse through all nested fields in the map and struct value.

> **Note**
>
> Reflection is tricky business as we attempt to find out the type metadata of an object at runtime. So we can't fully rely on static types and the compiler to do the checking for us. There are many edge cases that are not captured in this example as we want to focus on the pertinent code on reflection. Please don't just copy the example without understanding what the code and include additional checks and logic for handling edge cases.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Credits

* [Golang Project: indirect function for getting the base type using reflect](https://github.com/golang/go/blob/master/src/html/template/content.go)
* [Spf13/cast for proper casting code logic in Go](https://github.com/spf13/cast)

## Reference

* [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
