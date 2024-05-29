package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count--
			mu.Unlock()
		}()
	}

	wg.Wait()

	mu.Lock()
	fmt.Println(count)
	mu.Unlock()
}
