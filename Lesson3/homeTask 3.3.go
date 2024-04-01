package main

import "fmt"

const myConst = 1

func main() {
	const myConst = 2
	fmt.Printf("Значение локальной константы, затеняющей глобальную: %v", myConst)

}
