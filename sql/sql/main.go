package main

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	time.Sleep(2000)
	id := insertSQL(randLat(), randLong())
	loc := selectSQL(id)
	fmt.Println(loc)

	updateSQL(id, randLat(), randLong())
	loc = selectSQL(id)
	fmt.Println(loc)
}
