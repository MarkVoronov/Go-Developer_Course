package main

import (
	"encoding/xml"
	"fmt"
)

type contract struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date,omitempty"`
	Landlord string `xml:"-"`
	tenat    string `xml:"tenat"`
}
type contracts struct {
	List []contract `xml:"contract"`
}

func main() {

	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		tenat:    "Шура Балаганов",
	}

	d := contracts{}
	d.List = append(d.List, c)

	res, err := xml.Marshal(d)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(xml.Header, string(res))
}
