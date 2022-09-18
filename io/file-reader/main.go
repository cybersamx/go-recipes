package main

import (
	"fmt"
	"io"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readAll(filename string) {
	fileReader, err := os.Open(filename)
	checkError(err)
	defer fileReader.Close()

	data, err := io.ReadAll(fileReader)
	checkError(err)
	fmt.Println(string(data))
}

func readIncremental(filename string) {
	fileReader, err := os.Open(filename)
	checkError(err)
	defer fileReader.Close()

	data := make([]byte, 16)
	for {
		n, err := fileReader.Read(data)
		if err == io.EOF {
			break
		}
		checkError(err)

		if n > 0 {
			fmt.Println(string(data[:n]))
		}
	}
}

func main() {
	readAll("testdata.txt")
	readIncremental("testdata.txt")
}
