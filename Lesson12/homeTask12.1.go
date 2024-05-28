package main

import (
	"fmt"
	"log"
)

func main() {
	a := 1
	do(a)
}
func do(v any) {
	a := 10

	if sl, ok := v.(int); ok {
		a += sl
	} else {
		log.Fatalf("Невозможно привести переданное значение %v к int", v)
	}

	fmt.Println(a)
}
