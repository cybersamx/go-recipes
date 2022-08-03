# Simple Static HTTP Server

This example shows a web server serving static content, which can be implemented using:

* The `http` package in the standard Go library.
* The third-party `gin-gonic/gin` package.

## Setup

To run the static web server implemented using the `http` package:

1. Run the server.

   ```bash
   $ make run-http
   ```

1. Launch a web browser and navigate to <http://localhost:8000>.

To run the static web server implemented using the `gin-gonic/gin` package:

1. Run the server

   ```bash
   $ make run-gin
   ```

1. Launch a web browser and navigate to <http://localhost:8000>.
