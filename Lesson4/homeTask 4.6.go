package main

import "fmt"

var acc int = 1

func main() {
	fmt.Println("1: 1") // первое число - это 1, выведем его
	fibo(0, 1)
}

func fibo(a int, b int) {
	acc++
	if acc > 23 {
		return
	}

	c := a + b
	fmt.Printf("%d: %d\n", acc, c)
	fibo(b, c)
}
