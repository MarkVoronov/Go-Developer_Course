package main

import "fmt"

func main0() {

	msg := newMessage("Hello")
	fmt.Println(msg)

}

func newMessage(v string) string {

	return fmt.Sprintf("%s, Go!", v)

}
