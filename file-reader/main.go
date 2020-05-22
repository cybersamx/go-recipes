// Modified from the official gRPC example: google.golang.org/grpc/examples/helloworld/helloworld.

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// naivePrint prints the content of the file directly using ReadAll in 1 pass.
func naivePrint(reader io.Reader) (io.ReadCloser, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))

	return ioutil.NopCloser(bytes.NewReader(data)), nil
}

// efficientPrint breaks the reading of a stream by chunks. This is useful for
// reading big files.
func efficientPrint(reader io.Reader) (io.ReadCloser, error) {
	data := make([]byte, 256)
	for {
		n, err := reader.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading file: %v", err)
		}

		if n > 0 {
			fmt.Print(string(data[:n]))
		}
	}

	return ioutil.NopCloser(bytes.NewReader(data)), nil
}

func main() {
	// Create a file, which is an "implementation" of io.Reader.
	filename := "testdata.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %s: %v", filename, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("--- Naive Print ---")
	closer, err := naivePrint(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- Efficient Print ---")
	_, err = efficientPrint(closer)
	if err != nil {
		log.Fatal(err)
	}
}
