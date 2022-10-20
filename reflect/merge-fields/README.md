# Merge Matching Fields from Map to Struct

Suppose we have a map[string]any. We want to copy the content of that map to a struct if the keys in the map matches the tags in the struct.

This just a simple example of mapping to a pre-defined struct value and doesn't support mapping the nested fields. Look at [merge fields recursive](../merge-fields-recursive) for nested mapping on any struct type.

> **Note**
>
> Reflection is tricky business as we attempt to find out the type metadata of an object at runtime. So we can't fully rely on static types and the compiler to do the checking for us. There are many edge cases that are not captured in this example as we want to focus on the pertinent code on reflection. Please don't just copy the example without understanding what the code and include additional checks and logic for handling edge cases.

## Notes

* To get the type of a struct value, we must pass the struct as a value to `reflect.TypeOf` function.
* To set a field in struct using `reflect`, the field must be exported. This means that the field name must be capitalized.
* If we want to set a struct value, we must pass the pointer of the struct value to `reflect.ValueOf` function.

## Setup

1. Run the program

   ```bash
   $ make run
   ```

## Reference

* [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
