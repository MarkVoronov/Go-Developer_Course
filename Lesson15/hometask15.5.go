package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type temperature struct {
	mu  sync.RWMutex
	val string
}

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

func (c *temperature) ReadTemp() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.val
}

func (c *temperature) ChangeTemp(v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val = v
}

func PrintTemperature(meteo Meteo) {
	fmt.Println("Текущая температура: ", meteo.ReadTemp())
}

func main() {

	var tempEnv = &temperature{}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		defer wg.Done()
		for i := 1; i < 100; i++ {

			str := strconv.Itoa(i) + " градусов С"
			tempEnv.ChangeTemp(str)
			time.Sleep(time.Microsecond)
		}
	}()

	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			currentTemp := tempEnv.ReadTemp()
			fmt.Println(currentTemp)
		}()
	}

	wg.Wait()
	PrintTemperature(tempEnv)

}
