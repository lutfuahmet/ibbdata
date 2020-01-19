package models

import "github.com/jinzhu/gorm"

// District model
type District struct {
	gorm.Model
	Name string
}