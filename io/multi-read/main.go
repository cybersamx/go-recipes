package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func reread() {
	str := "hello world"
	reader := bytes.NewBufferString(str)
	// Need a buffer that can hold the string, hence the same length as str.
	buf := make([]byte, len(str))

	_, err := reader.Read(buf)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(buf))

	// Reading it again returns nothing because the stream has been read.
	rebuf := make([]byte, len(str))
	reader.Read(rebuf)
	fmt.Println(string(rebuf))

	// This won't work either.
	retBuf, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(retBuf))

	// Instantiate a new reader to reread the content again.
	reader = bytes.NewBufferString(str)
	rebuf = make([]byte, len(str))
	reader.Read(rebuf)
	fmt.Println(string(rebuf))
}

func main() {
	reread()
}
