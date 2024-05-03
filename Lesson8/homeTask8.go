package main

import (
	"fmt"
)

func main() {

	fmt.Println("Задание 8.1")
	type emptyStructure struct{}
	myMap := map[string]emptyStructure{
		"слон":    emptyStructure{},
		"бегемот": emptyStructure{},
		"носорог": emptyStructure{},
		"лев":     emptyStructure{},
	}
	fmt.Println(myMap)

	fmt.Println("Задание 8.2")

	myMap2 := make(map[string]int, 4)
	myMap2["слон"] = 3
	myMap2["бегемот"] = 0
	myMap2["носорог"] = 5
	myMap2["лев"] = 1

	values := []string{"слон", "бегемот", "осьминог"}

	for _, value := range values {
		count, ok := myMap2[value]
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", value, count, ok)
	}

	fmt.Println("Задание 8.3")
	myMap3 := map[string]emptyStructure{
		"слон":    emptyStructure{},
		"бегемот": emptyStructure{},
		"носорог": emptyStructure{},
		"лев":     emptyStructure{},
	}
	delete(myMap3, "бегемот")
	fmt.Println(myMap3)

	fmt.Println("Задание 8.4")
	myMap4 := map[string]emptyStructure{
		"слон":    emptyStructure{},
		"бегемот": emptyStructure{},
		"носорог": emptyStructure{},
		"лев":     emptyStructure{},
	}
	myMap4["выдра"] = emptyStructure{}
	fmt.Println(myMap4)

	fmt.Println("Задание 8.5")

	myMap5 := make(map[string]int, 5)
	myMap5["слон"] = 3
	myMap5["бегемот"] = 0
	myMap5["носорог"] = 5
	myMap5["лев"] = 1

	myMap5["бегемот"] = 2
	fmt.Println(myMap5)
}
