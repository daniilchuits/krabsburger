package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func goroutWithCont(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Остановка: context canceled")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Работаю")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go goroutWithCont(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	wg.Wait()
}
