package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/tealeg/xlsx"
)

// Park model
type Park struct {
	gorm.Model
	Name string `gorm:"not null"`
	ParkLocation ParkLocation
	ParkLocationID uint
	ParkType ParkType
	ParkTypeID uint
	Capacity uint
	WorkingHours string `gorm:"not null"`
	ParkZone ParkZone
	ParkZoneID uint
	ParkSubZone ParkSubZone
	ParkSubZoneID uint
	District District
	DistrictID uint
	Latitude float64
	Longitude float64
	MonthlyFee float64
	Tariffs string `gorm:"type:text;"`
	FreeParkingTime uint
	IsParkContinuePoint bool
	Address string `gorm:"type:text;"`
}

// GetAllParks => Parks
func GetAllParks() []Park {
	var parks []Park
	DB.Find(&parks)
	return parks
}

// Row2Park helper
func Row2Park(row *xlsx.Row) (park Park,err error) {
	park = Park{}
	cells := row.Cells

	if len(cells) < 24 {
		err = fmt.Errorf("missing cells :  %d",len(cells))
		return
	}

	parkID,_ := cells[0].Int()
	park.ID = uint(parkID)
	park.Name = cells[1].String()

	lokasyonID,_ := cells[2].Int()
	lokasyonKodu,_ := cells[3].Int()
	lokasyonAdi := cells[4].String()

	parkLokasyon := ParkLocation{}
	parkLokasyon.ID = uint(lokasyonID)
	parkLokasyon.Code = uint(lokasyonKodu)
	parkLokasyon.Name = lokasyonAdi

	parkLokasyon,err  = parkLokasyon.Init()
	if err != nil {
		return
	}
	park.ParkLocation = parkLokasyon
	park.ParkLocationID = parkLokasyon.ID

	parkTypeID,_ := cells[5].Int()
	parkTypeName := cells[6].String()

	parkType := ParkType{}
	parkType.ID = uint(parkTypeID)
	parkType.Name = parkTypeName

	parkType,err  = parkType.Init()
	if err != nil {
		return
	}
	park.ParkType = parkType
	park.ParkTypeID = parkType.ID

	capacity,_ := cells[7].Int()
	park.Capacity = uint(capacity)
	park.WorkingHours = cells[8].String()
	
	zoneID , _ := cells[9].Int()
	zoneName := cells[10].String()
	parkZone := ParkZone{}
	parkZone.ID = uint(zoneID)
	parkZone.Name = zoneName

	parkZone,err = parkZone.Init()
	if err != nil {
		return
	}
	park.ParkZone = parkZone
	park.ParkZoneID = parkZone.ID

	subZoneID,_ := cells[11].Int()
	subZoneName := cells[12].String()
	parkSubZone := ParkSubZone{}
	parkSubZone.ID = uint(subZoneID)
	parkSubZone.Name = subZoneName
	parkSubZone.ParkZoneID = park.ParkZoneID

	parkSubZone,err = parkSubZone.Init()
	if err != nil {
		return
	}
	park.ParkSubZone = parkSubZone
	park.ParkSubZoneID = parkSubZone.ID

	districtID,_ := cells[13].Int()
	districtName := cells[14].String()

	district := District{}
	district.ID = uint(districtID)
	district.Name = districtName
	district ,err = district.Init()
	if err != nil {
		return
	}

	park.District = district
	park.DistrictID = district.ID

	park.Address = cells[15].String()
	park.Longitude,_ = cells[18].Float()
	park.Latitude,_ = cells[19].Float()
	park.MonthlyFee,_ = cells[20].Float()

	freeParkingTime, _ := cells[21].Int()
	park.FreeParkingTime = uint(freeParkingTime)
	park.Tariffs = cells[22].String()
	park.IsParkContinuePoint = cells[23].Bool()

	err = DB.Create(&park).Error
	return 
}