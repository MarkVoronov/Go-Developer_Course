package main

import "fmt"

func main() {

	const (
		sku1 = iota
		sku2
		sku3
		sku4
		sku5
	)
	fmt.Printf(
		"Значение пяти локальных констант: %v %v %v %v %v",
		sku1,
		sku2,
		sku3,
		sku4,
		sku5,
	)

}
