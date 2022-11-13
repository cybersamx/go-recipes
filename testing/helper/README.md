# T.Helper() Function

What is T.Helper() and how to use it.

It's sometimes useful to write test helpers so that common operations can be run with a test helper. If an error is
raised in a test helper, eg. failed assertion, it would be more meaningful for Go to indicate the filename and the
line number of the function calling the test helper than the filename and line number of the test helper. To
accomplish this, we mark a test helper with `T.Helper()` at the start of the function.

## Reference

* [Testing.Helper doc](https://pkg.go.dev/testing#T.Helper)
