package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int

	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()

	// Создадим еще одну горутину, которая вперед main горутины отправит сигналы в канал stop.
	// Поскольку используется небуферизованный канал, вторая горутина будет заблокирована после записи очередного значения. Поможем ей
	// Для этого прочитаем записанное ей значение, чтобы она так же смогла завершиться.
	// Остается проблема завершить текущую горутину, но как это сделать я не придумал
	go func() {

		stop <- struct{}{}
		stop <- struct{}{}
		for range ch {
		}
	}()

	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}
