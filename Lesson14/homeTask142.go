package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go func() {
		msg := "Привет, строковый канал!"
		ch <- msg
	}()
	fmt.Println(<-ch)

}
