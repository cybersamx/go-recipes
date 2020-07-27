package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	datestr := "26 Jul 2020 13:47:56 PDT"
	layout := "2 Jan 2006 15:04:05 MST"

	date, err := time.Parse(layout, datestr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day:  ", date.Day())
	fmt.Println("Month:", date.Month())
	fmt.Println("Year: ", date.Year())
	fmt.Println("Hour: ", date.Hour())
	fmt.Println("Min:  ", date.Minute())
	fmt.Println("Sec:  ", date.Second())
	fmt.Println("Nano: ", date.Nanosecond())
	fmt.Println("TZ:   ", date.Location())
}
