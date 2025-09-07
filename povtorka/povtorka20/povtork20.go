package main

import (
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	Num     int
	IsPrime bool
}

func Writer(jobs chan int, wg *sync.WaitGroup, r rand.Rand) {
	for i:=0; i < 20; i++ {
		a := r.Intn(100)+1
		jobs <- a
		wg.Done()
	}
}

func Worker(jobs chan int, results chan Result){
	for a :=range jobs
}

func main() {
	r:=rand.Seed(rand.NewSource(time.now().UnixNano))
	jobs := make(chan int, 20)
	results := make(chan Result, 20)
	var wg sync.WaitGroup
	nums:=20

	wg.Add(nums)
	go Writer(jobs, &wg, r)

	go func ()  {
		wg.Wait()
		close(jobs)
	}
}