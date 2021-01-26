package main

import (
	"fmt"
	"time"
)

var tcases = []struct {
	input  string
	layout string
}{
	{
		input:  "26 Jul 20 13:47 PDT",
		layout: "02 Jan 06 15:04 MST",
	},
	{
		input:  "26 Jul 20 13:47 PDT",
		layout: time.RFC822, // Same as above
	},
	{
		input:  "26 Jul 20 13:47 -0700",
		layout: "02 Jan 06 15:04 -0700",
	},
	{
		input:  "26 Jul 20 13:47 -0700",
		layout: time.RFC822Z, // Same as above
	},
	{
		input:  "Sun, 26 Jul 2020 13:47:56 PDT",
		layout: "Mon, 2 Jan 2006 15:04:05 MST",
	},
	{
		input:  "Sun, 26 Jul 2020 13:47:56 PDT",
		layout: time.RFC1123, // Same as above
	},
	{
		input:  "Sun, 26 Jul 2020 13:47:56 -0700",
		layout: "Mon, 2 Jan 2006 15:04:05 -0700",
	},
	{
		input:  "Sun, 26 Jul 2020 13:47:56 -0700",
		layout: time.RFC1123Z, // Same as above
	},
	{
		input:  "2020-07-26T13:47:56-07:00",
		layout: "2006-01-02T15:04:05Z07:00",
	},
	{
		input:  "2020-07-26T13:47:56-07:00",
		layout: time.RFC3339, // Same as above
	},
	{
		input:  "2020-07-26T13:47:56.093906-07:00",
		layout: "2006-01-02T15:04:05.999999999Z07:00",
	},
	{
		input:  "2020-07-26T13:47:56.093906-07:00",
		layout: time.RFC3339Nano, // Same as above
	},
}

func main() {
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	refDate := time.Date(2020, 7, 26, 13, 47, 56, 93906, loc)
	before := refDate.Add(-time.Minute)
	after := refDate.Add(time.Minute)

	for _, tcase := range tcases {
		date, err := time.Parse(tcase.layout, tcase.input)
		if err != nil {
			panic(err)
		}

		if date.After(before) && date.Before(after) {
			fmt.Println("Parse successful:", tcase.layout)
		} else {
			fmt.Println("Parse failure:", tcase.layout)
		}
	}
}
