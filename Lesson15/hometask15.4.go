package main

import (
	"fmt"
	"sync"
)

func start() {
	fmt.Println("Некоторое сообщение из функции start")
}

var loadOnce sync.Once

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			loadOnce.Do(start)
			fmt.Printf("Запустилась горутина %d\n", i)
		}(i)
	}

	wg.Wait()

}
