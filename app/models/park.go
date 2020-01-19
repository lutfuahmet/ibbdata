package models

import (
	"github.com/jinzhu/gorm"
)

// Park model
type Park struct {
	gorm.Model
	Name string
	ParkLocation ParkLocation
	ParkLocationID uint
	ParkType ParkType
	ParkTypeID uint
	ParkZone ParkZone
	ParkZoneID uint
	ParkSubZone ParkSubZone
	ParkSubZoneID uint
	Latitude float64
	Longitude float64
	MonthlyFee float64
	Tariffs string `gorm:"type:text;"`
	FreeParkingTime uint
	IsParkContinuePoint bool
	Address `gorm:"type:text;"`
}

// ParkLocation model
type ParkLocation struct {
	gorm.Model
	Name string
	Code string `gorm:"unique;not null"`
}

// ParkType model
type ParkType struct {
	gorm.Model
	Name string
}

// ParkZone model
type ParkZone struct {
	gorm.Model
	Name string
}

// ParkSubZone model
type ParkSubZone struct {
	gorm.Model
	Name string
	ParkZone ParkZone
	ParkZoneID uint
}

