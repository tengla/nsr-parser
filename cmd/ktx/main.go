package main

import (
	"context"
	"fmt"
	"time"
)

func mychan(ctx context.Context) chan byte {
	chars := "\t\nabc"
	ch := make(chan byte)
	go func() {
		t, _ := ctx.Deadline()
		for i := 0; i < len(chars); i++ {
			ch <- chars[i]
			time.Sleep(time.Second)
			if time.Now().After(t) {
				break
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	fmt.Println("main")
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*5))
	defer cancel()
	ints := mychan(ctx)
	for i := range ints {
		fmt.Println(i)
	}
}

func init() {
	fmt.Println("init")
}
