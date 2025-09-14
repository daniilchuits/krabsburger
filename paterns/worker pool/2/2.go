package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Word       string
	Length     int
	Vowels     int
	Consonants int
}

func Worker(jobs chan string, results chan Result, wg *sync.WaitGroup, glasniye []rune) {
	defer wg.Done()
	for j := range jobs {
		vowels := 0
		for _, run := range j {
			for _, gl := range glasniye {
				if run == gl {
					vowels++
					break
				}
			}
		}
		con := len(j) - vowels
		results <- Result{
			Word:       j,
			Length:     len(j),
			Vowels:     vowels,
			Consonants: con,
		}
	}
}

func main() {
	words := []string{
		"Moscow",
		"Berlin",
		"Tokyo",
		"Amsterdam",
		"Copenhagen",
		"Lisbon",
		"Warsaw",
		"Madrid",
		"Rome",
		"Vienna",
		"Paris",
		"Prague",
		"London",
		"Helsinki",
		"Dublin",
	}
	jobs := make(chan string)
	results := make(chan Result)
	var wg sync.WaitGroup
	const numWorkers = 3
	const numJobs = 3
	glasniye := []rune{'e', 'u', 'i', 'o', 'a'}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(jobs, results, &wg, glasniye)
	}

	for _, word := range words {
		jobs <- word
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("Word: %-10s | Length: %-2d | Vowels: %d | Consonants: %d\n", r.Word, r.Length, r.Vowels, r.Consonants)
	}

}
