# Simple Static Web Server

This recipe shows a web server serving static content. Two implementations are included in the recipe, they are:

* The `http` package in the standard Go library.
* The `gin-gonic/gin` package.

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

> **Notes**
>
> When running the program via `make run-http` or `make run-gin`, the program uses the directory `public` in the project root dir as the default root content directory. You can also pass the directory path to the program by running `static-http -dir [directory]`.
