package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Word       string
	Wovels     int
	Consonants int
}

func consumer(jobs chan string, results chan Result, wovels []rune, wg *sync.WaitGroup) {
	for w := range jobs {
		defer wg.Done()
		runs := []rune(w)
		glasnie := 0

		for _, runes := range runs {
			for _, r := range wovels {
				if r == runes {
					glasnie++
				}
			}
		}
		con := len(w) - glasnie
		results <- Result{Word: w, Wovels: glasnie, Consonants: con}
	}
}

func main() {
	words := []string{
		"concurrency", "goroutine", "channel",
		"synchronization", "parallelism", "mutex",
	}
	wovels := []rune{'a', 'e', 'i', 'o', 'u'}
	jobs := make(chan string, len(words))
	results := make(chan Result, len(words))
	numWorkers := 3
	var wg sync.WaitGroup

	for _, word := range words { // producer
		wg.Add(1)
		jobs <- word
	}
	close(jobs)

	for i := 0; i < numWorkers; i++ {
		go consumer(jobs, results, wovels, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("Word: %s | Vowels: %d | Consonants: %d\n", r.Word, r.Wovels, r.Consonants)
	}
}
