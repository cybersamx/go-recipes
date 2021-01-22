# gRPC Hello World

An example of gRPC client and server. It is a modification of [the official gRPC example](google.golang.org/grpc/examples/helloworld/helloworld).

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
   $ protoc -I . hello.proto --go_out=plugins=grpc:$(PROTO_DIR) \
     --go_opt=paths=source_relative
   ```

1. Run both server and client concurrently.

   ```bash
   $ go run main.go
   ```

   You can also run the following command that combines steps 2 and 3.

   ```bash
   $ make install  # Optional: run this if you haven't installed protoc-gen-go
   $ make run
   ```

## Reference and Credits

* [gRPC Homepage](https://grpc.io/)
