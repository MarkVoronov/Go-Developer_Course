package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter int32 // Определим тип, с которым смогут работать сразу несколько горутин

func (c *counter) Increment() {
	atomic.AddInt32((*int32)(c), 1)
}

func (c *counter) Decrement() {
	atomic.AddInt32((*int32)(c), -1)
}

func main() {

	wg := sync.WaitGroup{}

	var count counter
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Increment()
		}()
	}

	for i := 0; i < 500; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Decrement()
		}()
	}

	wg.Wait()
	fmt.Println(count)

}
