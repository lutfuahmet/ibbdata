package models

import (
	"errors"
	"strings"
	"github.com/jinzhu/gorm"
)

// Validate Park
func (p Park) Validate(db *gorm.DB) {
	if p.ID == 0 || strings.TrimSpace(p.Name) == "" {
		db.AddError(errors.New("id and name can't be blank"))
	}
	if p.ParkLocationID == 0 {
		db.AddError(errors.New("parklocation can't be blank"))
	}
	if p.ParkTypeID == 0 {
		db.AddError(errors.New("parktype can't be blank"))
	}
	if p.ParkZoneID == 0 {
		db.AddError(errors.New("parkzone can't be blank"))
	}
	if p.ParkSubZoneID == 0 {
		db.AddError(errors.New("parksubzone can't be blank"))
	}
	if p.DistrictID == 0 {
		db.AddError(errors.New("district can't be blank"))
	}
}

// Validate ParkLocation
func (p ParkLocation) Validate(db *gorm.DB) {
	if p.ID == 0 || strings.TrimSpace(p.Name) == "" {
		db.AddError(errors.New("id and name can't be blank"))
	}
}

// Validate ParkType
func (p ParkType) Validate(db *gorm.DB) {
	if p.ID == 0 || strings.TrimSpace(p.Name) == "" {
		db.AddError(errors.New("id and name can't be blank"))
	}
}

// Validate ParkZone
func (p ParkZone) Validate(db *gorm.DB) {
	if p.ID == 0 || strings.TrimSpace(p.Name) == "" {
		db.AddError(errors.New("id and name can't be blank"))
	}
}

// Validate ParkSubZone
func (p ParkSubZone) Validate(db *gorm.DB) {
	if p.ID == 0 || strings.TrimSpace(p.Name) == "" || p.ParkZoneID == 0 {
		db.AddError(errors.New("id, name, parkzone can't be blank"))
	}
}

// Validate District
func (p District) Validate(db *gorm.DB) {
	if p.ID == 0 || strings.TrimSpace(p.Name) == "" {
		db.AddError(errors.New("id, name can't be blank"))
	}
}