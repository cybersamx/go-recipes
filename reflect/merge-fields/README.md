# Merge Matching Fields from Map to Struct

Suppose we have a map[string]any. We want to copy the content of that map to a struct if a key in the map matches a tag in the struct.

## Notes

* To get the type of a struct value, we must pass the struct as a value to `reflect.TypeOf` function.
* To set a field in struct using `reflect`, the field must be exported. This means that the field name must be capitalized.
* If we want to set a struct value, we must pass the pointer of the struct value to `reflect.ValueOf` function.

## Setup

1. Run the program

   ```bash
   $ make run
   ```
