# data-integration-challenge
RESTful API to perform merge with companies data.

Data description

| Name (upper case) | Address Zip (five digit text) | Website (lower case) |
| ------ | ------ | ------ |
|  NAME | 12345 | http://www.company.com |


##### Endpoints

Listening port : 8000.

| Name | Path | Method | Content-Type | Description |
| ------ | ------ | ------ | ------ | ------ |
| List companies| /v1/companies | GET | application/json | Retrieve all stored companies in database. |
| Merge companies | /v1/companies/load | POST | multipart/form-data | Parses a CSV file and merge its data with the existent records. If the record doesn't exist, it'll be discarded. The file must be extension "csv".|
| Find company | /v1/companies/find/{name}/{addresszip} | GET | application/json | Retrieve information of a company, Supported parameters: name (part of the company's name) AND addresszip(the entire zip code of the company) |





### Set up application

Steps to run this application :
- Download and Install Go : [Download](https://golang.org/dl/)
- Get third party libraries
- Build and run the application

```sh
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get github.com/mattn/go-sqlite3
go build
./data-integration-challenge
```

During application startup, a CSV file (q1_catalog.csv) located in the `file` folder will be analyzed and used to create a database in Sqlite and subsequently inserted its records.

##### Tests

To run tests on application, execute the command on directory:
```sh
go test -v
```
