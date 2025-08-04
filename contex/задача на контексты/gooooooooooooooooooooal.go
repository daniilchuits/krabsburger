package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func rabota(ctx context.Context) {
	files := []string{"file1", "file2", "file3", "file4"}
	for _, f := range files {
		wg.Add(1)
		go func(ctx context.Context, fname string) {
			defer wg.Done()
			for {
				select {
				case <-time.After(time.Second):
					fmt.Printf("%s downloading\n", f)
				case <-ctx.Done():
					if v := ctx.Value("userID"); v != nil {
						fmt.Printf("[%s] stopped for userID: %v, reason: %s\n", f, v, ctx.Err())
						return
					}
				}
			}
		}(ctx, f)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "userID", 7)

	go rabota(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	wg.Wait()
}
