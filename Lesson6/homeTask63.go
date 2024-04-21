package main


import "fmt"

type contact struct {
	ID           int
	Number, Date string
}

func (w contact) show() string {
	return fmt.Sprintf("Договор № %s от %s", w.Number, w.Date)
}

func main() {

	w := contact{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Println(w.show())

}
