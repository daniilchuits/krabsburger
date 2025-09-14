package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Result struct {
	Num     int
	IsPrime bool
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Writer(jobs chan int) {
	for i := 0; i < 20; i++ {
		a := rand.Intn(100) + 1
		jobs <- a
	}
	close(jobs)
}

func Worker(jobs chan int, results chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for a := range jobs {
		results <- Result{Num: a, IsPrime: IsPrime(a)}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make(chan int, 20)
	results := make(chan Result, 20)
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go Worker(jobs, results, &wg)
	}
	go Writer(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()
	prime := []int{}
	for r := range results {
		if r.IsPrime {
			prime = append(prime, r.Num)
		}
	}
	sort.Ints(prime)
	for _, k := range prime {
		fmt.Println(k)
	}
}
