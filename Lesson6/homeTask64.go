package main

import "fmt"

type contact struct {
	Addss, Phone string
}

type user struct {
	ID   int
	Name string
	contact
}
type employee struct {
	ID   int
	Name string
	contact
}

func main() {

	u := user{
		ID:   1,
		Name: "Petr",
		contact: contact{
			Addss: "USSR",
			Phone: "8904444",
		},
	}

	e := employee{
		ID:   2,
		Name: "Yana",
		contact: contact{
			Addss: "Emirates",
			Phone: "555-5678",
		},
	}

	fmt.Println(u.Addss, u.Phone)
	fmt.Println(e.Addss, e.Phone)

}
