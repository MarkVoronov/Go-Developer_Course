package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	ch := make(chan string)
	for i := 0; i < 10; i++ { // Для запуска рабочих горутин
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ch:
					fmt.Println("Стоп горутина ", i)
					return
				default:
					fmt.Println("Сложные вычисления горутины ", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("Ой, все")
		close(ch)
	}()
	wg.Wait()
}
