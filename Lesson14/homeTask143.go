package main

import (
	"fmt"
)

func main() {

	ch := make(chan string, 4)
	go func() {
		ch <- "Привет"
		ch <- "буферизованный канал!"
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}

}
