# Simple Static HTTP Server

This recipe shows a web server serving static content. The sample code uses 2 third-party frameworks:

* The `http` package in the standard Go library.
* The open source `gin-gonic/gin` package.

## Setup

To run the static web server implemented using the `http` package, run:

1. Run the server

   ```bash
   $ make run-http
   ```

1. Launch a web browser and navigate to <http://localhost:8000>.

To run the static web server implemented using the `gin-gonic/gin` package, run:

1. Run the server

   ```bash
   $ make run-gin
   ```

1. Launch a web browser and navigate to <http://localhost:8000>.
