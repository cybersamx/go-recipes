# Regular Expression in Go

A regular expression recipe in Go.

## Notes

* It's a best practice to use backtick for regexp pattern. String enclosed by backticks are more "literal," which makes it easier to write - eg. backlash isn't interpreted as an escape.

  ```go
  patternInDQuote   := "[a-zA-Z0-9\\]+"
  patternInBacktick := `[a-zA-Z0-9\]+`
  ```

## Setup

1. Run the program

   ```bash
   $ make
   ```

## Reference

* [Go package: regexp](https://golang.org/pkg/regexp/)
* [Go package: regexp/syntax](https://golang.org/pkg/regexp/syntax/)
