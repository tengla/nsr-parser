package main

import (
	"context"
	"fmt"
	"time"
)

func mychan(ctx context.Context) chan byte {
	chars := "abcdefghijklmnopqrstuvwxyz"
	ch := make(chan byte)
	go func() {
		t, _ := ctx.Deadline()
		for i := 0; i < len(chars); i++ {
			ch <- chars[i]
			time.Sleep(time.Millisecond * 50)
			if time.Now().After(t) {
				break
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*30))
	defer cancel()
	ints := mychan(ctx)
	for i := range ints {
		fmt.Printf("%c", i)
	}
}
