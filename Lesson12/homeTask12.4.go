package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}
type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	// Третий вариант решения - это в самой функции проверять, а не nil ли нам передали
	if d == nil {
		fmt.Println("Ошибка пения: не передана утка")
		return ""
	}
	return d.voice
}
func main() {
	var d *Duck
	// Первый вариант решения. Присвоим переменной d конкретное значение. Тогда метод Sing() сможет отработать
	d = &Duck{
		voice: "кря",
	}

	// Второй вариант решения - это объявить переменную d не как указатель на Duck, а как Bird (не указатель). Тогда интерфейс будет равен nil
	//var d Bird

	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}

// Почему происходит паника?
// Потому что на принимающей стороне (функция Sing) находится интерфейс Bird, а передаем мы в него неинициализированную переменную типа Duck
// Поэтому интерфейс получает знание о типе переменной - это Duck, и проходит проверку на nil. То есть он уже не nil
// Затем вызывается функция Sing типа Duck, но она не может отработать без установленного свойства voice
