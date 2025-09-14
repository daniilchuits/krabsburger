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
		j = strings.TrimSpace(j)
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
	var text string = "Go is expressive concise clean and efficient Concurrency is not parallelism Channels orchestrate communication"
	jobs := make(chan string, len(text))
	var words []string
	var wg sync.WaitGroup

	words = strings.Fields(text)

	for _, word := range words {
		wg.Add(1)
		jobs <- word
	}
	results := make(chan Result, len(jobs))

	for i := 0; i < 3; i++ {
		go Consumer(jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	mapa := make(map[Result]int)
	for r := range results {
		mapa[r]++
	}

	for word, count := range mapa {
		fmt.Printf("Word: %-15s | Count: %d | Letters: %d\n", word.Word, count, word.Letters)
	}

	totalSim := 0
	totalSl := 0
	for word := range mapa {
		totalSim += word.Letters
		totalSl++
	}

	var srDlina float64 = float64(totalSim) / float64(totalSl)

	fmt.Printf("Средняя длина слов: %.2f", srDlina)

}
