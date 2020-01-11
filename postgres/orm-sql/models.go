package main

import (
	"time"
)

// Make sure the model here and the schema defined in db-init.sql are consistent.

type User struct {
	ID          int           `json:"id"   xorm:"INT PK 'id' AUTOINCR"`
	Name        string        `json:"name" xorm:"VARCHAR(64)"`
	Age         int           `json:"age"  xorm:"INT"`
	Restaurants []*Restaurant `json:"restaurants" xorm:"-"`
}

type Restaurant struct {
	ID        int       `json:"id"         xorm:"INT PK 'id' AUTOINCR"`
	UserID    int       `json:"user_id"    xorm:"INT"`
	VisitedAt time.Time `json:"visited_at" xorm:"TIMESTAMP"`
	Name      string    `json:"name"       xorm:"VARCHAR(64)"`
	NumSeats  int       `json:"num_seats"  xorm:"INT"`
	Latitude  float32   `json:"latitude"   xorm:"FLOAT"`
	Longitude float32   `json:"longitude"  xorm:"FLOAT"`
}

// --- XORM ---

func (u *User) TableName() string {
	return "users"
}

func (r *Restaurant) TableName() string {
	return "restaurants"
}
