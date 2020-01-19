package models

import (
	"github.com/jinzhu/gorm"
)

// ParkType model
type ParkType struct {
	gorm.Model
	Name string
}

// Init for ParkType
func (p ParkType) Init() (parkType ParkType,err error) {
	parkType = ParkType{}
	DB.First(&parkType,p.ID)
	if parkType.ID == 0 {
		if err = DB.Create(&p).Error;err != nil {
			return
		}
		return p,nil
	}
	return
}