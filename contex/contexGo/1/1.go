package main

import (
	"context"
	"fmt"
	"time"
)

func Worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker [%s] finished: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("[%s] working...\n", name)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel1()

	ctx2, cancel2 := context.WithCancel(context.Background())

	go Worker(ctx1, "timeout")
	go Worker(ctx2, "cancel")

	time.Sleep(time.Second)
	cancel2()

	time.Sleep(2 * time.Second)
}
