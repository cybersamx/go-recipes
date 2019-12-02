package main

import "time"

// User represents a user account for an application.
type User struct {
    ID         string
	Username   string
	Email      string
	City       string
	Age        int
    CreatedAt  time.Time
	ModifiedAt time.Time
}
