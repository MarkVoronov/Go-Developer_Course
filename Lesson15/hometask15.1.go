package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		i := i
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Горутина %d\n", i)
		}(i) // захватываем горутину через параметр функции
	}

	wg.Wait()

}
