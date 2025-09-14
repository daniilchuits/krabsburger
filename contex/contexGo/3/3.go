package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, name string, wg *sync.WaitGroup, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Printf("[%s] downloaded\n", name)
		wg.Done()
	case <-ctx.Done():
		fmt.Printf("[%s] canceled or timedout\n", name)
		fmt.Println("RESULT: process failed due to timeout")
		wg.Done()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wgFirstTwo sync.WaitGroup
	var wg sync.WaitGroup
	wgFirstTwo.Add(2)
	go Worker(ctx, "file 1", &wgFirstTwo, 2*time.Second)
	go Worker(ctx, "file 2", &wgFirstTwo, 3*time.Second)
	wgFirstTwo.Wait()
	wg.Add(1)
	go Worker(ctx, "file 3", &wg, 4*time.Second)
	wg.Wait()
}
