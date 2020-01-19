package main

import (
	 "github.com/qor/admin"
	  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main(){
	Admin := admin.New(&admin.AdminConfig{DB: models.DB})
}