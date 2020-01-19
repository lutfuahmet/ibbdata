package main

import (
    "github.com/tealeg/xlsx"
	"log"
)

const fileName = "ispark-otoparklarna-ait-bilgiler.xlsx"

func main() {
	importFile()
}

func importFile() {
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(xlFile.Sheets[0].Name)
}