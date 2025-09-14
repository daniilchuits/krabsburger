package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	Num      int
	NaSebya  int
	WorkerID int
	Count    int
}

func Consumer(ctx context.Context, id int, jobs <-chan int, results chan<- Result, cwg *sync.WaitGroup) {
	defer cwg.Done()
	count := 0
	defer func() {
		results <- Result{Num: 0, NaSebya: 0, WorkerID: id, Count: count}
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[Worker %d] Контекст завершился\n", id)
			return
		case num := <-jobs:
			time.Sleep(time.Duration(rand.Intn(401)+100) * time.Millisecond)

			orig := num
			sqwr := num * num

			select {
			case <-ctx.Done():
				fmt.Printf("[Worker %d] Контекст завершился (canceled)\n", id)
				return
			case results <- Result{Num: orig, NaSebya: sqwr, WorkerID: id}:
				count++
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rand.Seed(time.Now().UnixNano())
	jobs := make(chan int, 100)
	results := make(chan Result, 100)
	var cwg sync.WaitGroup
	const numCons = 5

	for i := 0; i < numCons; i++ {
		cwg.Add(1)
		go Consumer(ctx, i+1, jobs, results, &cwg)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				jobs <- rand.Intn(100) + 1
			}
		}
	}()

	go func() {
		cwg.Wait()
		close(results)
	}()

	for r := range results {
		if r.Num == 0 {
			fmt.Printf("worker %d did %d jobs\n", r.WorkerID, r.Count)
		} else {
			fmt.Printf("Num: %d - Kvadrat: %d - Worker %d\n",
				r.Num, r.NaSebya, r.WorkerID)
		}
	}

}
