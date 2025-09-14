package main

import (
	"fmt"
	"strings"
	"sync"
)

type Result struct {
	Word    string
	Letters int
}

func Consumer(jobs chan string, results chan Result, wg *sync.WaitGroup) {
	for j := range jobs {
		j = strings.ToLower(j)
		total := 0
		for range j {
			total++
		}
		results <- Result{
			Word:    j,
			Letters: total,
		}
		wg.Done()
	}
}

func main() {
	text := []string{
		"Go is expressive, concise, clean and efficient.",
		"Concurrency is not parallelism!",
		"Channels orchestrate communication.",
	}
	var wg sync.WaitGroup
	jobs := make(chan string, len(text)*10)

	for _, sent := range text {
		words := strings.Fields(sent)
		for _, word := range words {
			wg.Add(1)
			clearedWord := strings.Trim(word, ".,")
			jobs <- clearedWord
		}
	}
	close(jobs)
	results := make(chan Result, len(text)*10)

	for i := 0; i < 3; i++ {
		go Consumer(jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	mapa := make(map[string]int)
	for r := range results {
		mapa[r.Word]++
	}

	for word, count := range mapa {
		fmt.Printf("Word: %-15s | Count: %-2d | Letters: %d \n", word, count, len(word))
	}
}
