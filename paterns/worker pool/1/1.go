package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Url    string
	Length int
}

func Worker(id int, jobs chan string, results chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d is doing %s\n", id, j)
		time.Sleep(time.Second)
		results <- Result{
			Url:    j,
			Length: len(j),
		}
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://example.com",
		"https://github.com",
		"https://google.com",
		"https://openai.com",
		"https://wikipedia.org",
		"https://reddit.com",
		"https://stackoverflow.com",
		"https://news.ycombinator.com",
		"https://medium.com",
	}
	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))
	var wg sync.WaitGroup
	const numWorkers = 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, jobs, results, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("Line: %-30s | Length: %d\n", r.Url, r.Length)
	}
}
