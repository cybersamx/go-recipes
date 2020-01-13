package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	records = [][]string{
		{"1", "LA-0664", "34.19462", "-118.58843", "NB  De Soto  FS  Vanowen -NE", "L.A. Valley West"},
		{"2", "LA-0682", "34.20065", "-118.6234", "SB  Fallbrook  FS  Sherman Way -SW", "L.A. Valley West"},
		{"3", "LA-0683", "34.20065", "-118.6234", "SB  Fallbrook  FS  Sherman Way -SW", "L.A. Valley West"},
		{"4", "LA-0712", "34.15238", "-118.60487", "EB  Mulholland D  NS  Topanga -SW", "L.A. Valley West"},
	}
)

func fatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v\n", msg, err)
	}
}

func writeToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, os.ModePerm)
	fatal("can't open file", err)

	defer func() {
		file.Close()
	}()

	writer := csv.NewWriter(file)

	for _, record := range records {
		fatal("can't write to file buffer", writer.Write(record))
	}

	writer.Flush()
	fatal("can't flush to file", writer.Error())

	log.Println("Done writing to file...")
}

func readFromFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}

	defer func() {
		file.Close()
	}()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6

	for {
		record, err := reader.Read()
		if err == io.EOF {
			// Must check for EOF.
			break
		} else if err != nil {
			log.Fatalf("can't read file: %v", err)
		}

		fmt.Println(record)
	}
}

func main() {
	filename := "data.csv"
	writeToFile(filename)
	readFromFile(filename)
}
