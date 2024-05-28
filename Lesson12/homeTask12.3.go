package main

import "fmt"

type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}

type documenter interface {
	Format()
}

func main() {
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}
func Report(x documenter) {
	if x != nil {
		x.Format()
	}

}
