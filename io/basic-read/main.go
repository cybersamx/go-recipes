package main

import (
	"bytes"
	"fmt"
	"io"
)

func read() {
	str := "hello world"

	// Read with reader.Read()
	reader := bytes.NewBufferString(str)
	buf := make([]byte, len(str))
	_, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read:", string(buf))

	// Read with ReadFrom() in bytes.Buffer
	reader = bytes.NewBufferString(str)
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ReadFrom:", buffer.String())

	// Read with io.ReadAll().
	reader = bytes.NewBufferString(str)
	allbuf, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ReadAll:", string(allbuf))

	// Read with io.ReadFull().
	reader = bytes.NewBufferString(str)
	fullbuf := make([]byte, len(str))
	_, err = io.ReadFull(reader, fullbuf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ReadFull:", string(allbuf))
}

func main() {
	read()
}
