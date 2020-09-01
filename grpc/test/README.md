# gRPC Test

An example of writing tests for a gRPC server.

## Setup

1. Install protobuf compiler and the plugin for Go.

   ```bash
   $ brew install protobuf
   $ go get github.com/golang/protobuf/protoc-gen-go
   $ # Run the following if you haven't done so already.
   $ export PATH="$PATH:$(go env GOPATH)/bin"
   ```

1. Run the tests.

   ```bash
   $ make test
   ```

## Reference and Credits

* [gRPC Homepage](https://grpc.io/)
