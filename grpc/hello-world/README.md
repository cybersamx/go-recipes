# gRPC

An example of gRPC using protobuf. It is a modification of [the official gRPC example](google.golang.org/grpc/examples/helloworld/helloworld).

## Setup

1. Install protobuf compiler and the plugin for Go.

   ```bash
   $ brew install protobuf
   $ go get github.com/golang/protobuf/protoc-gen-go
   $ # Run the following if you haven't done so already.
   $ export PATH="$PATH:$(go env GOPATH)/bin"
   ```

1. Generate code from the protobuf definition file.

   ```bash
   $ mkdir hello
   $ protoc -I . hello.proto --go_out=plugins=grpc:hello
   ```

1. Run both server and client concurrently.

   ```bash
   $ go run main.go
   ```

   You can also run the following command that combines steps 2 and 3.

   ```bash
   $ make run
   ```

## Reference and Credits

* [gRPC Homepage](https://grpc.io/)
