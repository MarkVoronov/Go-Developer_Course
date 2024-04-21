package main

import "fmt"

type contact struct {
	ID           int
	Number, Date string
}

func main() {

	w := contact{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Printf("{ID:%d Number:%s Date:%s}\n", w.ID, w.Number, w.Date)

}
