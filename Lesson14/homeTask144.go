package main

import "fmt"

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	go func() {
		<-ch
		stop <- struct{}{}
	}()
	close(ch) // Закроем канал, тогда горутина сможет прочитать нулевое значение и перейти к следующей строке
	<-stop
	fmt.Println("happy end")
}
