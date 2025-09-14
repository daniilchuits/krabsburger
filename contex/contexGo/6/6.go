package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Worker(ctx context.Context, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()

	name, _ := ctx.Value(ctxKey("name")).(string)
	if name == "" {
		name = "zalupa"
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	timer := time.NewTimer(duration)
	defer timer.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("%s is working...\n", name)
		case <-timer.C:
			fmt.Printf("%s done\n", name)
			return
		case <-ctx.Done():
			fmt.Printf("%s is canceled or timed out(%v)\n", name, ctx.Err())
			return
		}
	}
}

var wg sync.WaitGroup

type ctxKey string

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx0, cancel0 := context.WithCancel(context.Background())
	defer cancel0()

	ctxTime, cancelTime := context.WithTimeout(ctx0, 5*time.Second)
	defer cancelTime()

	ctx1 := context.WithValue(ctxTime, ctxKey("name"), "file_1")
	ctx2 := context.WithValue(ctxTime, ctxKey("name"), "file_2")
	ctx3 := context.WithValue(ctxTime, ctxKey("name"), "file_3")
	ctx4 := context.WithValue(ctxTime, ctxKey("name"), "file_4")
	ctx5 := context.WithValue(ctxTime, ctxKey("name"), "file_5")

	wg.Add(5)
	go Worker(ctx1, &wg, time.Duration(rand.Intn(2)+4)*time.Second)
	go Worker(ctx2, &wg, time.Duration(rand.Intn(2)+4)*time.Second)
	go Worker(ctx3, &wg, time.Duration(rand.Intn(2)+4)*time.Second)
	go Worker(ctx4, &wg, time.Duration(rand.Intn(2)+4)*time.Second)
	go Worker(ctx5, &wg, time.Duration(rand.Intn(2)+4)*time.Second)
	wg.Wait()
}
