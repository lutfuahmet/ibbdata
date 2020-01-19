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

	for i,row := range dataSheet.Rows {
			if i == 0 {
				continue // skip header row
			}
			park,err := models.Row2Park(row)
			if err != nil{
				log.Printf("error %s, row %d",err.Error(),i + 1)
				continue
			}
			log.Println("park created succesfully")
			log.Println(park)
	}

}
