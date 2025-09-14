package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Text     string
	WorkerID int
	Duration time.Duration
}

func Producer(jobs chan string, pwg *sync.WaitGroup, count int, id int) {
	defer pwg.Done()
	for i := 1; i <= count; i++ {
		text := fmt.Sprintf("Toy_%d_fromProducer_%d", i, id)
		jobs <- text
	}
}

func Consumer(jobs chan string, results chan Result, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for j := range jobs {
		timeNow := time.Now()
		time.Sleep(time.Second)
		j += fmt.Sprint(" - assembled by Worker")
		results <- Result{
			Text:     j,
			WorkerID: id,
			Duration: time.Since(timeNow),
		}
	}
}

func main() {
	jobs := make(chan string, 100)
	results := make(chan Result, 100)
	var pwg sync.WaitGroup
	var wg sync.WaitGroup

	producerCount := []int{1, 1, 5}
	for id, count := range producerCount {
		pwg.Add(1)
		go Producer(jobs, &pwg, count, id+1)
	}

	go func() {
		pwg.Wait()
		close(jobs)
	}()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Consumer(jobs, results, &wg, i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("%s %d, for %s\n", r.Text, r.WorkerID, r.Duration)
	}
}
