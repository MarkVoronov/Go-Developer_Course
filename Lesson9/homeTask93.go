package main

import "fmt"

var typeEat map[string]string

func main() {

	typeEat = map[string]string{

		"груша":    "фрукт",
		"яблоко":   "фрукт",
		"апельсин": "фрукт",
		"тыква":    "овощ",
		"огурец":   "овощ",
		"помидор":  "овощ",
	}

	var temp string
	fmt.Println("Введите название еды:")
	fmt.Scanf("%s\n", &temp)
	checkFood(temp)
}

func checkFood(name string) {

	var txtMsg string

	value, status := typeEat[name]

	switch {
	case !status:
		txtMsg = "что-то странное"
	case value == "фрукт":
		txtMsg = "это фрукт"
	case value == "овощ":
		txtMsg = "это овощ"
	}
	fmt.Println(txtMsg)
}
