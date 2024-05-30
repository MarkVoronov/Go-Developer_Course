package main

import (
	"context"
	"fmt"
	"sync"
)

type contextKey string

const (
	key1 contextKey = "some key 1"
	key2 contextKey = "some key 2"
)

func main() {

	ctx := context.WithValue(context.Background(), key1, "some value 1")
	ctx = context.WithValue(ctx, key2, "some value 2")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(ctx.Value(key1))
		fmt.Println(ctx.Value(key2))

	}()
	wg.Wait()
}
