package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, name string, duration time.Duration, wg *sync.WaitGroup, done chan struct{}, made []<-chan struct{}) {
	defer wg.Done()

	for _, ma := range made {
		select {
		case <-ma:
		case <-ctx.Done():
			fmt.Printf("[%s] canceled due to dependency or timeout\n", name)
			return
		}
	}

	select {
	case <-time.After(duration):
		fmt.Printf("%s downloaded\n", name)
		if done != nil {
			close(done)
		}
		return
	case <-ctx.Done():
		fmt.Printf("%s won't be downloaded\n", name)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	done1 := make(chan struct{})
	done2 := make(chan struct{})
	done3 := make(chan struct{})
	done4 := make(chan struct{})

	wg.Add(5)
	go Worker(ctx, "file 1", 1*time.Second, &wg, done1, nil)
	go Worker(ctx, "file 2", 2*time.Second, &wg, done2, nil)
	go Worker(ctx, "file 3", 2*time.Second, &wg, done3, []<-chan struct{}{done1, done2})
	go Worker(ctx, "file 4", 3*time.Second, &wg, done4, []<-chan struct{}{done2})
	go Worker(ctx, "file 5", 3*time.Second, &wg, nil, []<-chan struct{}{done3, done4})
	wg.Wait()

	fmt.Println("All workers finished or canceled")

}
