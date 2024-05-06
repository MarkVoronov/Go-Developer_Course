package main


import (
	"fmt"
)

var stockMap map[string]int

func main() {

	var temp string
	fmt.Println("Введите название фрукта:")
	fmt.Scanf("%s\n", &temp)

	num, status := fruitMarket(temp)
	if !status {
		fmt.Printf("Введенное название %s отсутствует в карте!", temp)
	} else {
		fmt.Printf("Число фруктов с названием %s: %d", temp, num)
	}

}
func fruitMarket(name string) (int, bool) {

	if len(stockMap) == 0 {

		stockMap = map[string]int{

			"апельсин": 5,
			"яблоки":   3,
			"сливы":    1,
			"груши":    0,
		}

	}

	num, status := stockMap[name]
	return num, status
}
