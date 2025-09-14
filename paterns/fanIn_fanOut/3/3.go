package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Word       string
	Vowels     int
	Consonants int
}

func producer(jobs <-chan string, glasnie []rune, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for word := range jobs {
		runes := []rune(word)

		vowels := 0
		for _, run := range runes {
			for _, g := range glasnie {
				if run == g {
					vowels++
					break
				}
			}
		}
		consonants := len(runes) - vowels

		results <- Result{Word: word, Vowels: vowels, Consonants: consonants}
	}
}

func main() {
	words := []string{"qwe", "asd", "zxc", "rty", "fgh", "vbn", "uio", "jkl", "m,.", "asddasf", "asdqw", "ghrt", "fgdsprgte", "dqwdgreg", "ewfqsd", "qwdsadxcz"}
	glasnie := []rune{'e', 'y', 'u', 'i', 'o', 'a'}

	jobs := make(chan string, len(words))
	results := make(chan Result, len(words))
	var wg sync.WaitGroup

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go producer(jobs, glasnie, results, &wg)
	}

	go func() {
		for _, w := range words {
			jobs <- w
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("Слово: %-10s | Гласных: %d | Согласных: %d\n", r.Word, r.Vowels, r.Consonants)
	}
}
