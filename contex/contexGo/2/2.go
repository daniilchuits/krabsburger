package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, name string, duration time.Duration, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(duration):
		if ctx.Err() == nil {
			fmt.Printf("%s done\n", name)
		}
		select {
		case result <- name:
		default:
		}
	case <-ctx.Done():
		fmt.Printf("%s cancled or timeout\n\n", name)
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	result := make(chan string, 2)

	wg.Add(3)
	go Worker(ctx, "worker 1", 5*time.Second, result, &wg)
	go Worker(ctx, "worker 2", 2*time.Second, result, &wg)
	go Worker(ctx, "worker 3", time.Second, result, &wg)

	var results []string
	for i := 0; i < 2; i++ {
		r := <-result
		results = append(results, r)
	}
	cancel()

	wg.Wait()

	for i, place := range results {
		fmt.Printf("%s came in the %d place\n", place, i+1)
	}
	time.Sleep(time.Second)
}
