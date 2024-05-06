package main


import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3}
	var v1, v2, v3, v4 int

	for i := range s {
		v1 = s[i]
		fmt.Printf("v1: %v\n", v1)
	OUR_LABEL:
		for j := range s {
			v2 = s[j]
			fmt.Printf("\tv2: %v\n", v2)
			for k := range s {
				v3 = s[k]
				fmt.Printf("\t\tv3: %v\n", v3)
				for n := range s {
					v4 = s[n]
					fmt.Printf("\t\t\tv4: %v\n", v4)
					if n == 1 {
						break OUR_LABEL
					}
				}
			}
		}

	}

}
