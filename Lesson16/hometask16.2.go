package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()
			fmt.Println("Запущена горутина", j)
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Стоп горутина", j)
					return
				default:
					time.Sleep(10 * time.Microsecond)
				}
			}

		}()
	}
	wg.Wait()
}
