package models

import "github.com/jinzhu/gorm"

// District model
type District struct {
	gorm.Model
	Name string
}

// Init for ParkZone
func (p District) Init() (district District,err error) {
	district = District{}
	DB.First(&district,p.ID)
	if district.ID == 0 {
		if err = DB.Create(&p).Error;err != nil {
			return
		}
		return p,nil
	}
	return
}