package main

import "fmt"

func main() {

	fmt.Println("Задача 7.1")

	var m [5]string
	m[0] = "Аля"
	m[1] = "Таня"
	m[2] = "Оля"
	m[3] = "Галя"
	m[4] = "Арфеня"
	fmt.Println(m)

	fmt.Println("Задача 7.2")
	var food [4]string = [4]string{"яблоко", "груша", "слива", "абрикос"}
	fmt.Println(food)

	fmt.Println("Задача 7.3")
	food2 := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	food2[2] = "персик"
	fmt.Println(food2)

	fmt.Println("Задача 7.4")
	s1 := []int{5, 2, 8, 3, 1, 9}
	fmt.Println(s1)

	fmt.Println("Задача 7.5")
	s2 := make([]int, 0, 10) // пустой срез емкостью 10
	fmt.Println(s2)

	fmt.Println("Задача 7.6")
	s3 := make([]int, 0, 10)
	s3 = append(s3, 4, 1, 8, 9)
	fmt.Println(s3)

	fmt.Println("Задача 7.7")
	s4 := []int{1, 2, 3}
	s5 := []int{4, 5, 6}

	result := append(s4, s5...)
	fmt.Println(result)

	fmt.Println("Задача 7.8")
	s6 := []int{1, 2, 3, 4, 5, 6}

	s6 = append(s6[:3], s6[4:]...)
	fmt.Println(s6)
	
}
