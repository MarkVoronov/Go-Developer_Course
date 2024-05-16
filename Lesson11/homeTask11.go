package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	msg string
}

func (err *myFirstError) Error() string {
	return err.msg
}

func main() {

	fmt.Println("Задание 11.1")
	err := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err)
	err3 := fmt.Errorf("ошибка3:%w", err2)

	fmt.Println(err3)

	fmt.Println("Задание 11.2")
	wrapErr := errors.Unwrap(err3)
	fmt.Println(wrapErr)

	fmt.Println("Задание 11.3")
	if errors.Is(err3, err) {
		fmt.Println("В цепочке ошибок была ошибка 1")
	} else {
		fmt.Println("В цепочке ошибок отсутствует ошибка 1")
	}

	fmt.Println("Задание 11.4")
	var myFirstError *myFirstError
	if errors.As(err, &myFirstError) {
		fmt.Println("В цепочке ошибок была ошибка типа myFirstError")
	} else {
		fmt.Println("В цепочке ошибок отсутствует ошибка типа myFirstError")
	}
}
