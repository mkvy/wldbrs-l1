package main

import (
	"context"
	"fmt"
	"time"
)

// для реализации sleep будем использовать timeout context, возвращаем ничего
func sleepCtx(c context.Context, t time.Duration) {
	ctx, _ := context.WithTimeout(c, t)
	select {
	case <-ctx.Done():
		return
	}
}

func main() {
	ctx := context.Background()
	//задаем время в секундах
	timeSec := time.Duration(5)
	fmt.Printf("Start waiting\n")
	sleepCtx(ctx, timeSec*time.Second)
	fmt.Printf("Waited for %d seconds\n", timeSec)
}
