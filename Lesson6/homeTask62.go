package main


import "fmt"

func main() {

	type contact struct {
		ID           int
		Number, Date string
	}

	w := contact{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}

	fmt.Printf("{ID:%d Number:%s Date:%s}\n", w.ID, w.Number, w.Date)

}

