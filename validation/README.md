# Validation

The [validator package](https://github.com/go-playground/validator) provides validation against a variable value or against the values of the fields in a struct by declaring a "tag" that instruct the engine how we want to validate the value using some of the built-in rules. This recipe shows how we can use this package to perform value validation.

There are 2 parts in this recipe.

* `main.go` - sample code of using the validator package.
* `*_test.go` - a set of test cases that show the different validation tags - good sample code to see the extent/richness of supported validation rules that are provided by the package.

## Setup

1. Run the sample code.

   ```bash
   $ make run
   ```

1. Run (via test) to see the validation rules in action.

   ```bash
   $ make test
   ```

1. Run both at the same time.

   ```bash
   $ make
   ```

## Reference

* [Validator package](https://github.com/go-playground/validator)
* Some formats that are supported in the validation rules.
  * [ISO 3166](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes) - 2- and 3-letter country codes
  * [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) - Currency codes
  * [Timezone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
