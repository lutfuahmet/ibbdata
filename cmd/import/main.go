package main

import (
    "github.com/tealeg/xlsx"
	"ibbdata/app/models"
	"log"
)

const fileName = "parks.xlsx"

func main() {
	importFile()
}

func importFile() {
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// 2 sheet kontrolu, ilk sheet metadata, ikinc sheet park dataları
	if len(xlFile.Sheets) < 2 {
		return
	}

	dataSheet := xlFile.Sheets[1]

	var countMap = map[int]int{}

	for i,row := range dataSheet.Rows {
			if i == 0 {
				continue // skip header row
			}
			log.Println("row : len cells",i,len(row.Cells))
			count := len(row.Cells)
			if countMap[count] == 0 {
				countMap[count] = i
			}
	}
	
	log.Println(countMap)
}

func row2Park(row *xlsx.Row) (park models.Park,err error) {
	cells := row.Cells
	park = models.Park{}
	parkID,_ := cells[0].Int()
	park.ID = uint(parkID)
	park.Name = cells[1].String()

	lokasyonID,_ := cells[2].Int()
	lokasyonKodu,_ := cells[3].Int()
	lokasyonAdi := cells[4].String()

	parkLokasyon := models.ParkLocation{}
	parkLokasyon.ID = uint(lokasyonID)
	parkLokasyon.Code = uint(lokasyonKodu)
	parkLokasyon.Name = lokasyonAdi

	parkLokasyon = parkLokasyon.Init()
	park.ParkLocation = parkLokasyon
	park.ParkLocationID = parkLokasyon.ID

	parkTypeID,_ := cells[5].Int()
	parkTypeName := cells[6].String()

	parkType := models.ParkType{}
	parkType.ID = uint(parkTypeID)
	parkType.Name = parkTypeName

	parkType = parkType.Init()
	park.ParkType = parkType
	park.ParkTypeID = parkType.ID

	capacity,_ := cells[7].Int()
	park.Capacity = uint(capacity)
	park.WorkingHours = cells[8].String()
	
	zoneID , _ := cells[9].Int()
	zoneName := cells[10].String()
	parkZone := models.ParkZone{}
	parkZone.ID = uint(zoneID)
	parkZone.Name = zoneName

	parkZone = parkZone.Init()
	park.ParkZone = parkZone
	park.ParkZoneID = parkZone.ID

	subZoneID,_ := cells[11].Int()
	subZoneName := cells[12].String()
	parkSubZone := models.ParkSubZone{}
	parkSubZone.ID = uint(subZoneID)
	parkSubZone.Name = subZoneName
	parkSubZone.ParkZoneID = park.ParkZoneID

	

	return 
}