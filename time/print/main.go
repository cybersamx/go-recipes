package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Get the full list of timezones: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	loc, err := time.LoadLocation("US/Pacific")
	if err != nil {
		log.Fatal(err)
	}
	date := time.Date(2020, 7, 26, 13, 47, 56, 0, loc)

	// Ad-hoc layouts.

	layouts := []string {
		"Mon, 2 Jan 2006 15:04:05 MST",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05-0700",
		"2 Jan 2006 15:04:05",
		"2006-01-02",
		"20060102",
		"January 02, 2006",
		"02 Jan, 2006",
		"01/02/06",
		"01/02/2006",
		"Monday",
		"Mon",
		"15:04",
		"15:04:05",
		"3:04 PM",
		"03:04:05 PM",
	}

	fmt.Println("Ad hoc layouts")

	for _, layout := range layouts {
		fmt.Println(date.Format(layout))
	}

	// There are also pre-defined layouts.

	predefines := []string {
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
	}

	fmt.Println()
	fmt.Println("Pre-defined layouts")

	for _, predefine := range predefines {
		fmt.Println(date.Format(predefine))
	}
}
