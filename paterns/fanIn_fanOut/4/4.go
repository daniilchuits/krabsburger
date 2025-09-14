package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Url string
	Len int
}

func consumer(jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	for url := range jobs {
		// обёртка на одно задание — чтобы гарантировать вызов wg.Done()
		func(u string) {
			defer wg.Done()
			time.Sleep(time.Second)
			results <- Result{Url: u, Len: len(u)}
		}(url)
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://yandex.ru",
		"https://golang.org",
		"https://openai.com",
		"https://github.com",
	}
	jobs := make(chan string)
	results := make(chan Result, len(urls))
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		go consumer(jobs, results, &wg)
	}

	for _, url := range urls { // producer
		wg.Add(1)
		jobs <- url
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results { // читатель из общего канала
		fmt.Printf("Url: %-20s | Length: %d\n", r.Url, r.Len)
	}
}
