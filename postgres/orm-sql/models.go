package main

import (
	"time"
)

// Make sure the model here and the schema defined in db-init.sql are consistent.

type BusStop struct {
	ID        int       `xorm:"int pk 'id' autoincr"    gorm:"type:int;primary_key;auto_increment;column:id"`
	UpdatedAt time.Time `xorm:"timestamp 'updated_at'"  gorm:"type:timestamp;column:updated_at"`
	Number    string    `xorm:"varchar(16)"             gorm:"type:varchar(16)"`
	Latitude  float32   `xorm:"float"                   gorm:"type:float"`
	Longitude float32   `xorm:"float"                   gorm:"type:float"`
	SiteATS   string    `xorm:"varchar(64) 'siteats'"   gorm:"type:varchar(64);column:siteats"`
	CitySite  string    `xorm:"varchar(48) 'city_site'" gorm:"type:varchar(48);column:city_site"`
}

// --- XORM ---

func (r *BusStop) TableName() string {
	return "bus_stops"
}
