package main

import (
	"context"
	"fmt"
	"time"
)

func mychan(ctx context.Context) chan int {
	count := 0
	ch := make(chan int)
	go func() {
		t, _ := ctx.Deadline()
		for {
			ch <- count
			count += 1
			time.Sleep(time.Second)
			if count > 10 || time.Now().After(t) {
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
