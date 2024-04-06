package main

import "fmt"

func main() {
	// 4.1
	hello1()

	// 4.2 Анонимная функция, которая сразу вызывается
	func() {
		fmt.Println("Hello, Go!")
	}()

	// 4.3 Определим анонимную функцию и запишем ее в переменную
	af := func() {
		fmt.Println("Hello, Go!")
	}
	hello2(af)

	// 4.4 Вызываем функцию hello3, которая возвращает анонимную функцию. Используем ту же переменную, потому что
	// тип у af - функция
	af = hello3()
	af()

	// 4.5 функция hell01 уже определна ранее, не буду дублировать
	defer fmt.Println("Завершение работы")
}

func hello1() {

	fmt.Println("Hello, Go!")

}

func hello2(v func()) {

	v()

}

func hello3() func() {
	return func() {
		fmt.Println("Hello, Go!")
	}
}
