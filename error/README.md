# Error Wrapping

Most programs in the real world involves a chain of nested function calls. This means that an error may have to be passed several levels up before it finally gets handled.

Go 1.13 introduces error wrapping that allows each function to wrap additional context using the `fmt.Errorf()` function using the `%w` verb to wrap the root error as `error`. As the error bubbles up, we wrap the error with texts to provide additional contexts at subsequent levels.

For some open-source Go projects (with legacy pre-1.13 Go code), we may see the third-party [pkg/errors](https://github.com/pkg/errors) package.

| pkg/errors                | Go 1.13+                           |
|---------------------------|------------------------------------|
| errors.Wrap(err, message) | fmt.Errorf("%s: %w", message, err) |
| errors.Unwrap(err)        | errors.Unwrap(err)                 |
