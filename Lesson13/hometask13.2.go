package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int
	Landlord string
	Tenat    string
}

func main() {
	ctr := contract{
		Number:   2,
		Landlord: "Остап",
		Tenat:    "Шура",
	}

	res, err := json.Marshal(ctr)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", string(res))

}
