package models

import (
	"github.com/jinzhu/gorm"
)

// ParkLocation model
type ParkLocation struct {
	gorm.Model
	Name string
	Code uint `gorm:"unique;not null"`
}

// Init for ParkLocation
func (p ParkLocation) Init() (parklocation ParkLocation,err error) {
	parklocation = ParkLocation{}
	DB.First(&parklocation,p.ID)
	if parklocation.ID == 0 {
		if err = DB.Create(&p).Error;err != nil{
			return
		}
		return p,nil
	}
	return
}