package main

import (
	"fmt"
)

type square int

func main() {

	var s square = 34
	s += 10
	fmt.Println(s)

}

func (s square) String() string {
	return fmt.Sprintf("%d м²", s)
}
