package main

import "time"

type User struct {
    ID         string
	Username   string
	Email      string
	City       string
	Age        int
    CreatedAt  time.Time
	ModifiedAt time.Time
}
