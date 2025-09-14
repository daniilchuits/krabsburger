package main

import (
	"fmt"
	"strings"
	"sync"
)

type Result struct {
	Word   string
	Vowels int
	Sogl   int
}

func Consumer(jobs chan string, results chan Result, wg *sync.WaitGroup, glasnie []rune) {
	for r := range jobs {
		defer wg.Done()
		vowels := 0
		total := 0
		for _, runes := range r {
			for _, gl := range glasnie {
				if gl == runes {
					vowels++
				}
			}
			total++
		}
		results <- Result{
			Word:   r,
			Vowels: vowels,
			Sogl:   total - vowels,
		}
	}
}

func main() {
	text := []string{
		"Go is expressive concise clean and efficient.",
		"Concurrency is not parallelism.",
		"Channels orchestrate communication.",
	}
	glasnie := []rune{'e', 'u', 'i', 'o', 'a'}
	var wg sync.WaitGroup

	var wrdss []string
	for _, sentances := range text {
		words := strings.Fields(sentances)
		for _, word := range words {
			wrdss = append(wrdss, word)
		}
	}
	jobs := make(chan string, len(wrdss))
	results := make(chan Result, len(wrdss))

	for _, word := range wrdss {
		cleanWord := strings.Trim(word, ".")
		jobs <- cleanWord
		wg.Add(1)
	}
	close(jobs)

	for i := 0; i < 3; i++ {
		go Consumer(jobs, results, &wg, glasnie)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("Слово: %-15s | Глассных: %-3d | Согласных: %-2d\n", r.Word, r.Vowels, r.Sogl)
	}
}
