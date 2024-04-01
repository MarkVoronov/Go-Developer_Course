package main

import "fmt"

const (
	sku1 = iota
	sku2
	sku3
	sku4
	sku5
)

func main() {

	fmt.Printf(
		"Значение пяти глобальных констант: %v %v %v %v %v",
		sku1,
		sku2,
		sku3,
		sku4,
		sku5,
	)

}
