package models

import (
	"github.com/jinzhu/gorm"
)

// ParkSubZone model
type ParkSubZone struct {
	gorm.Model
	Name string
	ParkZone ParkZone
	ParkZoneID uint
}

// Init for ParkZone
func (p ParkSubZone) Init() (parkSubZone ParkSubZone,err error) {
	parkSubZone = ParkSubZone{}
	DB.First(&parkSubZone,p.ID)
	if parkSubZone.ID == 0 {
		if err = DB.Create(&p).Error;err != nil {
			return
		}
		return p,nil
	}
	return
}