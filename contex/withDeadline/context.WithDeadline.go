package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func rabota(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("тИК")
		case <-ctx.Done():
			fmt.Println("End:", ctx.Err())
			return
		}
	}
}

func main() {
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancle := context.WithDeadline(context.Background(), deadline)
	defer cancle()

	wg.Add(1)
	go rabota(ctx)
	wg.Wait()
}
