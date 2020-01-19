package models

import (
	"github.com/jinzhu/gorm"
)


// ParkZone model
type ParkZone struct {
	gorm.Model
	Name string
}

// Init for ParkZone
func (p ParkZone) Init() (parkZone ParkZone,err error) {
	parkZone = ParkZone{}
	DB.First(&parkZone,p.ID)
	if parkZone.ID == 0 {
		if err = DB.Create(&p).Error;err!=nil {
			return
		}
		return p,nil
	}
	return
}