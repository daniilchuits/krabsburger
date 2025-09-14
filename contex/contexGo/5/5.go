package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	name, _ := ctx.Value(ctxKey("name")).(string)
	if name == "" {
		name = "cho"
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("%s is working...\n", name)
		case <-ctx.Done():
			fmt.Printf("context %s is canceled or timedout (%v)\n", name, ctx.Err())
			return
		}
	}
}

type ctxKey string

func main() {
	var wg sync.WaitGroup
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	ctxTimeout, cancleTimeout := context.WithTimeout(ctx, 3*time.Second)
	defer cancleTimeout()

	ctx1 := context.WithValue(ctxTimeout, ctxKey("name"), "file 1")
	ctx2 := context.WithValue(ctxTimeout, ctxKey("name"), "file 2")
	ctx3 := context.WithValue(ctxTimeout, ctxKey("name"), "file 3")

	wg.Add(3)
	go Worker(ctx1, &wg)
	go Worker(ctx2, &wg)
	go Worker(ctx3, &wg)

	// time.Sleep(time.Second)
	// cancle()// отменяем если нужно
	wg.Wait()
}
