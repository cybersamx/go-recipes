# Merge Matching Fields Recursively from Map to Struct

Suppose we have a map[string]any. We want to copy the content of that map to a struct if the keys in the map matches the tags in the struct.

This example is an extension of the [simpler example](../merge-fields) of just mapping non-nested fields in a struct. This example traverse through all nested fields in the map and struct value.

## Setup

1. Run the program

   ```bash
   $ make run
   ```
