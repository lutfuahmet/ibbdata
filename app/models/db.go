package models

import (
  	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB ...
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	DB.SingularTable(true)
	autoMigrate()
}

func autoMigrate() {
	DB.AutoMigrate(&Park{},&ParkLocation{},&ParkType{},&ParkZone{},&ParkSubZone{},&District{})
}