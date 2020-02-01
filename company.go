package main

// Company : estrutura segundo dado originario do csv
type Company struct {
	ID         int64  `json:"_id"`
	Name       string `json:"name"`
	AddressZip string `json:"zip"`
	Website    string `json:"website"`
}

// Companies : slice da estrutura
type Companies []Company
