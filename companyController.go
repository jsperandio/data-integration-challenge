package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type companyController struct {
	iCompany Company
}

func (cc companyController) get(pName, pZip string) Companies {
	db, err := gorm.Open("sqlite3", DBName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var resultSet Companies

	db.Where("name LIKE ? AND address_zip = ?", "%"+pName+"%", pZip).Find(&resultSet)

	return resultSet

}

func (cc companyController) getAllCompanies() Companies {
	db, err := gorm.Open("sqlite3", DBName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var allCompanies Companies
	db.Find(&allCompanies)

	return allCompanies
}

func (cc companyController) addCompanies(listCompany Companies) bool {

	db, err := gorm.Open("sqlite3", DBName)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	for index := 0; index < len(listCompany); index++ {
		c := listCompany[index]
		db.Create(&c)
	}

	return true

}

func (cc companyController) mergeCompanies(listToInsert Companies) bool {
	db, err := gorm.Open("sqlite3", DBName)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	var c Company

	for i := 0; i < len(listToInsert); i++ {
		c = listToInsert[i]
		db.Model(&c).Where("name = ?", c.Name).Update("Website", c.Website)
	}

	return true
}
