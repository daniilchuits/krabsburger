package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Num int
	Fac int
}

func writer(jobs chan int) {
	for i := 1; i < 11; i++ {
		jobs <- i
	}
	close(jobs)
}

func factorial(jobs chan int, result chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		res := 1
		for i := 1; i <= n; i++ {
			res *= i
		}
		result <- Result{Num: n, Fac: res}
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan Result, 10)
	resultSlice := make([]Result, 10)

	var wg sync.WaitGroup

	numWorkers := 3
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go factorial(jobs, results, &wg)
	}

	go writer(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		resultSlice[r.Num-1] = r
	}

	for _, i := range resultSlice {
		fmt.Printf("Факториал числа %d - %d\n", i.Num, i.Fac)
	}
}
