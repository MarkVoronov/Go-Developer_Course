package main

import (
	"fmt"
	"unsafe"
)

type square int

func main() {
	// 5.2
	fmt.Println("Задание 5.2")
	str := "моя переменная"
	fmt.Printf("Значение переменной: %v, ее адрес: %v\n", str, &str)

	// 5.3
	fmt.Println("Задание 5.3")
	ptr := &str
	*ptr = "переменная для задания 5.3"
	fmt.Println(str)

	// 5.4 Создадим группу переменных и проанализируем их адреса
	fmt.Println("Задание 5.4")
	var a, b, c, d int
	addr_a := uintptr(unsafe.Pointer(&a))
	addr_b := uintptr(unsafe.Pointer(&b))
	addr_c := uintptr(unsafe.Pointer(&c))
	addr_d := uintptr(unsafe.Pointer(&d))

	delta1 := addr_b - addr_a
	delta2 := addr_c - addr_b
	delta3 := addr_d - addr_c
	fmt.Printf("Адреса переменных: %v %v %v %v\n", &a, &b, &c, &d)
	fmt.Printf("Адрес в десятичной системе :\n %v\n %v\n %v\n %v\n", addr_a, addr_b, addr_c, addr_d)
	fmt.Printf("Разница в адресах :  %v %v %v\n", delta1, delta2, delta3)

	// 5.5
	fmt.Println("Задание 5.5")
	someNumber := 8
	change(&someNumber)
	fmt.Println(someNumber)

	// 5.6
	fmt.Println("Задание 5.6")
	var s square = 25
	fmt.Printf("Значение переменной s типа square: %v\n", s)

	// 5.7
	fmt.Println("Задание 5.7")
	s = 30
	s += 15
	fmt.Printf("Значение переменной s после увеличения на 15: %v\n", s)

	// 5.8
	fmt.Println("Задание 5.8")
	s = 34
	s += 10
	fmt.Printf("Значение переменной по задаче 5.8: %s", s.representation())

}

// 5.5
func change(param *int) {
	*param = *param + 3
}

// 5.8
func (s square) representation() string {

	return fmt.Sprintf("%v м²", s)

}
