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
		case <-ctx.Done():
			fmt.Println("End:", ctx.Err())
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Working")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg.Add(1)
	go rabota(ctx)

	wg.Wait()
	<-ctx.Done()
}
