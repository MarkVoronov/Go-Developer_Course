package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	Tenat    string `json:"tenat"`
}

func main() {
	str := ` {"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`
	var contract Contract
	err := json.Unmarshal([]byte(str), &contract)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", contract)
}
