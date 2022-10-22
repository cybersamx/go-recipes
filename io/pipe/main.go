package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
	"unicode/utf8"
)

func readJSON() {
	reader, writer := io.Pipe()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Gets executed first.
		// We close to tell the reader that we have finished writing the data.
		defer writer.Close()

		fmt.Println("Started writing json message")
		if err := json.NewEncoder(writer).Encode(map[string]string{"message": "hello world"}); err != nil {
			log.Fatalln("Json encoding error:", err)
		}
		fmt.Println("Finished writing json message")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer reader.Close()

		var msg map[string]string
		if err := json.NewDecoder(reader).Decode(&msg); err != nil {
			log.Fatalln("Json encoding error:", err)
		}
		fmt.Println("Finished reading json message")

		fmt.Println("Content of json message:", msg)
	}()

	wg.Wait()
}

func readRune() {
	reader, writer := io.Pipe()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer writer.Close()

		str := "hello world"

		fmt.Println("Started writing runes")

		// rune = int32
		// byte = int8
		runebuf := make([]byte, utf8.UTFMax)
		for _, r := range str {
			utf8.EncodeRune(runebuf, r)
			if _, err := writer.Write(runebuf); err != nil {
				log.Fatalln("Write problem:", err)
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Finished writing runes")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer reader.Close()

		buf := make([]byte, utf8.UTFMax)
		for {
			// PipeWriter doesn't expose a channel for when a Write is performed. So we rely
			// on periodic check on the pipeline using time.Sleep.
			time.Sleep(200 * time.Millisecond)
			_, err := reader.Read(buf)
			switch {
			case err == io.EOF:
				fmt.Println("Finished reading runes")
				return
			case err != nil:
				log.Fatalln("Read problem:", err)
				return
			}
			fmt.Println(string(buf))
		}
	}()

	wg.Wait()
}

func main() {
	readJSON()
	readRune()
}
