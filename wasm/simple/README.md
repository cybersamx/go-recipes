# WebAssembly Hello World

An example of developing a WebAssembly that prints "hello world" in the browser console.

## Setup

1. Build the WebAssembly (wasm) and serve the wasm on a web server.

   ```bash
   $ make
   ```

1. Test the wasm by launching a web browser and navigate to <http://localhost:8080>. See the console for the output of "hello world".

## Notes

1. See the [Makefile](Makefile) to how to build a wasm.

## Reference and Credits

* [Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)
