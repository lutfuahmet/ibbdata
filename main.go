package main

import (
	"github.com/qor/admin"
	"log"
	"ibbdata/app/models"
	"net/http"
)

var panel *admin.Admin

func main(){
	panel = admin.New(&admin.AdminConfig{DB: models.DB})
	panel.AddResource(&models.Park{})

	mux := http.NewServeMux()
	// Mount admin to the mux
  	panel.MountTo("/", mux)

  	log.Fatal(http.ListenAndServe(":3000", mux))
}