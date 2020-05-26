// Modified from the official gRPC example: google.golang.org/grpc/examples/helloworld/helloworld.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	content = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse molestie purus vitae vestibulum rutrum. Mauris mollis pellentesque aliquam."
)

// naiveWrite prints the content of the file directly using ReadAll in 1 pass.
func naiveWrite(writer io.Writer) (io.Writer, error) {
	n, err := writer.Write([]byte(content))
	log.Printf("wrote %d bytes\n", n)
	if err != nil {
		return nil, err
	}

	return writer, nil
}

func main() {
	// Create a file, which is an "implementation" of io.Reader.
	filename := "testdata.txt"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("can't open file %s: %v", filename, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("--- file Write ---")
	writer, err := naiveWrite(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- bufio Write ---")
	_, err = naiveWrite(bufio.NewWriter(writer))
	if err != nil {
		log.Fatal(err)
	}
}
