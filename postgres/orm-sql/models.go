package main

import (
	"time"
)

// Make sure the model here and the schema defined in db-init.sql are consistent.

type User struct {
	ID          int           `xorm:"INT PK 'id' AUTOINCR" gorm:"type:INT;PRIMARY_KEY;AUTO_INCREMENT"`
	Name        string        `xorm:"VARCHAR(64)"          gorm:"type:VARCHAR(64)"`
	Age         int           `xorm:"INT"                  gorm:"type:INT"`
	Restaurants []*Restaurant `xorm:"-"                    gorm:"-"`
}

type Restaurant struct {
	ID        int       `xorm:"INT PK 'id' AUTOINCR" gorm:"type:INT;PRIMARY_KEY;AUTO_INCREMENT"`
	UserID    int       `xorm:"INT"                  gorm:"type:INT"`
	VisitedAt time.Time `xorm:"TIMESTAMP"            gorm:"type:TIMESTAMP"`
	Name      string    `xorm:"VARCHAR(64)"          gorm:"type:VARCHAR(64)"`
	NumSeats  int       `xorm:"INT"                  gorm:"type:INT"`
	Latitude  float32   `xorm:"FLOAT"                gorm:"type:FLOAT"`
	Longitude float32   `xorm:"FLOAT"                gorm:"type:FLOAT"`
}

// --- XORM ---

func (u *User) TableName() string {
	return "users"
}

func (r *Restaurant) TableName() string {
	return "restaurants"
}
