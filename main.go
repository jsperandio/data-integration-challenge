package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// CCompany : Corresponde ao controller para chamada de metodos da entidade company
var CCompany companyController

func main() {

	initialDataLoad()
	handleRequests()

}

func initialDataLoad() {

	deleteDatabase()
	initialMigration()
	loadSetupDataFromCsv()
}

func handleRequests() {

	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/companies", getAll).Methods(http.MethodGet)
	api.HandleFunc("/companies/load", post).Methods(http.MethodPost)
	api.HandleFunc("/companies/find/{name}/{addresszip}", get).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))

}

func get(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	paramName := params["name"]
	paramAddressZip := params["addresszip"]

	localizedCompanies := CCompany.get(strings.ToUpper(paramName), paramAddressZip)

	data, _ := json.Marshal(localizedCompanies)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return

}

func getAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	listCompany := CCompany.getAllCompanies()

	if len(listCompany) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listCompany)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Can't find any Companies on Database"}`))
	}

}

func post(w http.ResponseWriter, r *http.Request) {

	reader := csv.NewReader(r.Body)
	reader.Comma = ';'
	records, err := reader.ReadAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Error on file load"}`))
		log.Fatalln("Error on file load", err)
		return
	}
	if len(records) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Empty file"}`))
		// log.Fatalln("Empty file!")
		return
	}

	var listCompany Companies

	for i, record := range records {

		if i == 0 {
			// Pular linha do header
			continue
		}

		c := Company{
			Name:       strings.ToUpper(record[0]),
			AddressZip: record[1],
			Website:    strings.ToLower(record[2])}
		listCompany = append(listCompany, c)
	}

	if CCompany.mergeCompanies(listCompany) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Merged Data"}`))

	} else {
		log.Fatalln("Error on add itens to database!")
		w.WriteHeader(http.StatusInternalServerError)
	}

}
