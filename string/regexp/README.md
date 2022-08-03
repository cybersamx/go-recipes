# Regular Expression in Go

Regular expression in Go.

## Notes

* It's a best practice to use backtick for regexp pattern. String enclosed by backticks are more "literal," which makes it easier to write as certain characters like backlash isn't interpreted as an escape.

  ```go
  patternInDQuote   := "[a-zA-Z0-9\\]+" // Double quote
  patternInBacktick := `[a-zA-Z0-9\]+`  // Backtick
  ```

## Setup

1. Run the program.

   ```bash
   $ make run
   ```

## Reference

* [Go package: regexp](https://golang.org/pkg/regexp/)
* [Go package: regexp/syntax](https://golang.org/pkg/regexp/syntax/)
