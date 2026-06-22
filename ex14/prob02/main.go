package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, parentCancelFunc := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer parentCancelFunc()

	total := 0
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("合計: %d, 反復回数: %d, 終了理由: %s\n", total, i, ctx.Err())
			return
		default:
		}

		n := rand.Intn(100_000_000)
		if n == 1234 {
			fmt.Printf("合計: %d, 反復回数: %d, 終了理由: 成功！\n", total, i)
			return
		}
		total += n
		i++
	}
}
