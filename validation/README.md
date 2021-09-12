# Validation

The [validator package](https://github.com/go-playground/validator) provides validation against a variable value or against the values of the fields in a struct by declaring a "tag" to instruct the engine how we want to validate the value. This recipe shows how we can use this package to perform value validation.

There are 2 parts in this recipe.

* `main.go` - sample code of using the validator package.
* `*_test.go` - a library of test cases testing the different validation tags.

## Setup

1. Run the sample code.

   ```bash
   $ make
   ```

1. Test all validation  tags.

   ```bash
   $ make test
   ```

## Reference

* [Validator package](https://github.com/go-playground/validator)
