package main

import (
	"fmt"
	hellomodule_v0 "hellomodule_v0"
	hellomodule_v1 "hellomodule_v1"
)

func main() {
	fmt.Println(hellomodule_v0.Hello())
	fmt.Println(hellomodule_v1.Hello())
}
