package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DBName : nome do banco
const DBName = "yawoen.db"

func deleteDatabase() {

	_, currentFilePath, _, _ := runtime.Caller(0)
	dirpath := path.Dir(currentFilePath)
	dbPath := dirpath + `/` + DBName

	if _, err := os.Stat(dbPath); err == nil {

		var err = os.Remove(dbPath)
		if err != nil {
			return
		}

	}

}

func initialMigration() {
	db, err := gorm.Open("sqlite3", DBName)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Company{})
}

func loadSetupDataFromCsv() {
	var listCompany Companies

	setupCsv, err := os.Open("file/q1_catalog.csv")

	if err != nil {
		log.Printf("Error on load data from CSV. ")
		return
	}
	defer setupCsv.Close()

	reader := csv.NewReader(setupCsv)
	reader.Comma = ';'

	records, err := reader.ReadAll()

	for i, record := range records {

		if i == 0 {
			// Pular linha do header
			continue
		}

		c := Company{
			Name:       strings.ToUpper(record[0]),
			AddressZip: record[1]}

		listCompany = append(listCompany, c)
	}

	if !CCompany.addCompanies(listCompany){
		log.Fatalln("Error on add itens to database!")
	}

}
