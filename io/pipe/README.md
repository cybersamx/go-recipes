# Pipe

This example shows how we can establish a synchronous, in-memory pipe using `io.Pipe()` that allows content to be
written on one end of the pipe (`io.PipeWriter`) to the content being read on the other end of the pipe (`io.
PipeReader`). The `io.Pipe()` function returns a pair of reader and writer values for us to establish the pipe.

```go
reader, writer := io.Pipe()
```

Here are 2 examples:

* `readJSON` - Pass the reader/writer pair from `io.Pipe()` to `json.Decoder` and `json.Encoder` respectively. So
  when a value gets encoded to json, it would be written to the pipe writer, and simultaneously read by the pipe
  reader and decoded to the value.
* `readRune` - Write a rune asynchronously to the pipe writer and get read synchronously by the pipe reader.

## Reference and Credits

* [Go Package io](https://golang.org/pkg/io/)
