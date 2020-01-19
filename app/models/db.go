package models

import (
  	"github.com/jinzhu/gorm"
	"log"
	"github.com/qor/validations"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB ...
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "/tmp/ibbdata.db")
	if err != nil {
		log.Fatal(err)
	}
	
	DB.SingularTable(true)
	validations.RegisterCallbacks(DB)
	autoMigrate()
}

func autoMigrate() {
	DB.AutoMigrate(&Park{},&ParkLocation{},&ParkType{},&ParkZone{},&ParkSubZone{},&District{})
}