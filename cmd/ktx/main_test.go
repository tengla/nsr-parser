package main

import (
	"context"
	"testing"
	"time"
)

func TestKtx(t *testing.T) {
	t.Run("ktx", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(time.Second*30))
		defer cancel()
		ch := mychan(ctx)
		var b byte
		for c := range ch {
			b = c
		}
		if b != 'z' {
			t.Fatal("expected a 'z'")
		}
	})
}
